/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Kim Letkeman - Initial Contribution
*/


// Recent state management
// v1   KL 21 Feb 2016 Initial split from contract
// v2   KL 11 Mar 2016 All state stored in the ledger.
// v3   KL 15 Mar 2016 read cleaned up and returns one level of escaping now

package main // sitting beside the main file for now

import (
	"encoding/json"
    "errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//***************************************************
//***************************************************
//* RECENT STATE MANAGEMENT 
//***************************************************
//***************************************************

// A state is simply a JSON encoded string, just as written to the ledger
// This module stores in memory only, as recent states make no sense after
// a restart. 

// RECENTSTATESKEY is used as key for recent states bucket
const RECENTSTATESKEY string = "RecentStatesKey"

// RecentStates is JSON encoded string slice 
type RecentStates struct {
    RecentStates []string `json:"recentStates"`
}

// AssetIDT is assetID as type, used for simple unmarshaling
type AssetIDT struct {
    ID string `json:"assetID"`
} 

// MaxRecentStates is an arbitrary limit on how many asset states we track across the 
// entire contract
const MaxRecentStates int = 20

// GETRecentStatesFromLedger returns the unmarshaled recent states
func GETRecentStatesFromLedger(stub *shim.ChaincodeStub) (RecentStates, error) {
    var state = RecentStates{make([]string, 0, MaxRecentStates)}
    var err error
	recentStatesBytes, err := stub.GetState(RECENTSTATESKEY)
	if err == nil { 
		err = json.Unmarshal(recentStatesBytes, &state.RecentStates)
		if err != nil {
            log.Noticef("Unmarshal failed for recent states: %s", err)
		}
	}
    // this MUST be here
    if state.RecentStates == nil || len(state.RecentStates) == 0 {
        state.RecentStates = make([]string, 0, MaxRecentStates)
    }
    log.Debugf("GETRecentStates returns: %#v", state)
    return state, nil 
}

// PUTRecentStatesToLedger marshals and writes the recent states
func PUTRecentStatesToLedger(stub *shim.ChaincodeStub, state RecentStates) (error) {
    var recentStatesJSON []byte
    var err error
    recentStatesJSON, err = json.Marshal(state.RecentStates)
    if err != nil {
        log.Criticalf("Failed to marshal recent states: %s", err)
        return err
    }
    err = stub.PutState(RECENTSTATESKEY, recentStatesJSON)
    if err != nil {
        log.Criticalf("Failed to PUTSTATE recent states: %s", err)
        return err
    } 
    log.Debugf("PUTRecentStates: %#v", state)
    return nil 
}

func clearRecentStates(stub *shim.ChaincodeStub) (error) {
    var rstates RecentStates
    rstates.RecentStates = make([]string, 0, MaxRecentStates)
    return PUTRecentStatesToLedger(stub, rstates)
}

func pushRecentState (stub *shim.ChaincodeStub, state string) (error) {
    var rstate RecentStates
    var err error
    var assetID string
    
    assetID, err = getAssetIDFromState(state)
    if err != nil {
        return err
    }
    rstate, err = GETRecentStatesFromLedger(stub)
    if err != nil {
        return err
    }
    
    // shift slice to the right
    assetPosn, err := findAssetInRecent(assetID, rstate) 
    if err != nil {
        return err
    } else if assetPosn == -1 {
        // grow if not at capacity, since this one is new
        if len(rstate.RecentStates) < MaxRecentStates {
            rstate.RecentStates = rstate.RecentStates[0 : len(rstate.RecentStates)+1]
        }
        // shift it all since not found
        copy(rstate.RecentStates[1:], rstate.RecentStates[0:])
    } else {
        if len(rstate.RecentStates) > 1 {
            // shift over top of the same asset, can appear only once
            copy(rstate.RecentStates[1:], rstate.RecentStates[0:assetPosn])
        }
    }
    rstate.RecentStates[0] = state
    return PUTRecentStatesToLedger(stub, rstate)
}

// typically called when an asset is deleted
func removeAssetFromRecentState (stub *shim.ChaincodeStub, assetID string) (error) {
    var rstate RecentStates
    var err error
    rstate, err = GETRecentStatesFromLedger(stub)
    if err != nil {
        return err
    }
    assetPosn, err := findAssetInRecent(assetID, rstate)
    if err != nil {
        return err
    } else if assetPosn == -1 {
        // nothing to do
        return nil
    } else {
        if len(rstate.RecentStates) > 0 {
            // shift slice to the left to close the hole left by the asset
            copy(rstate.RecentStates[assetPosn:], rstate.RecentStates[assetPosn+1:])
        }
        if len(rstate.RecentStates) > 0 {
            rstate.RecentStates = rstate.RecentStates[0 : len(rstate.RecentStates)-1]
        }
    }
    return PUTRecentStatesToLedger(stub, rstate)
}

func getAssetIDFromState(state string) (string, error) {
    var substate AssetIDT
    var err error
    err = json.Unmarshal([]byte(state), &substate)
    if err != nil {
        log.Errorf("getAssetIDFromState state unmarshal to AssetID failed: %s", err)
        return "", err
    }
    if len(substate.ID) == 0 {
        err = errors.New("getAssetIDFromState substate.common.assetID is blank")
        log.Error(err)
        return "", err
    }
    return substate.ID, nil 
}

func findAssetInRecent (assetID string, rstate RecentStates) (int, error) {
    // returns -1 to signify not found (or error)
    var err error
    var substate AssetIDT
    for i := 0; i < len(rstate.RecentStates); i++ {
        err = json.Unmarshal([]byte(rstate.RecentStates[i]), &substate)
        if err != nil {
            log.Errorf("findAssetInRecent JSON unmarshal of entry %d failed [%#v]", i, rstate.RecentStates[i])
            return -1, err
        }
        if substate.ID == assetID {
        	log.Debugf("findAssetInRecent found assetID %s at position %d in recent states", assetID, i)
            return i, nil
        }
    }
    // not found
    log.Debugf("findAssetInRecent Did not find assetID %s in recent states", assetID)
    return -1, nil
}

func readRecentStates(stub *shim.ChaincodeStub) ([]byte, error) {
	var err error
    var rstate RecentStates
    var rstateOut = make([]interface{}, 0, MaxRecentStates) 

	// Get the recent states from the ledger
    rstate, err = GETRecentStatesFromLedger(stub)
    if err != nil {
        return nil, err
    }
    for i := 0; i < len(rstate.RecentStates); i++ {
        var obj interface{}
        err = json.Unmarshal([]byte(rstate.RecentStates[i]), &obj)
        if err != nil {
            log.Errorf("findAssetInRecent JSON unmarshal of entry %d failed [%#v]", i, rstate.RecentStates[i])
            return nil, err
        }
        rstateOut = append(rstateOut, obj)
    }
    rsBytes, err := json.Marshal(rstateOut)
    if err != nil {
        log.Errorf("readRecentStates JSON marshal of result failed [%#v]", rstate.RecentStates)
        return nil, err
    }
	return rsBytes, nil
}



