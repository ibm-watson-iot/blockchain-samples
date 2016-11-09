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

// v0.1 KL -- new iot chaincode platform

package ctasset

import (
	"encoding/json"
	"fmt"

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
const RECENTSTATESKEY string = "RECENTSTATES"

// RecentStates is statearray
type RecentStates AssetArray

// MaxRecentStates is an arbitrary limit on how many asset states we track across the
// entire contract
const MaxRecentStates int = 20

// GETRecentStatesFromLedger returns the unmarshaled recent states
func GETRecentStatesFromLedger(stub shim.ChaincodeStubInterface) (RecentStates, error) {
	var state = make(RecentStates, 0, MaxRecentStates)
	var err error
	recentStatesBytes, err := stub.GetState(RECENTSTATESKEY)
	if err != nil {
		err = fmt.Errorf("Failed to get recent states from world state: %s", err)
		log.Errorf(err.Error())
		return state, err
	}
	// this MUST be here
	if recentStatesBytes == nil || len(recentStatesBytes) == 0 {
		log.Debugf("GETRecentStatesFromLedger: returning empty recent states")
		return state, nil
	}
	err = json.Unmarshal(recentStatesBytes, &state)
	if err != nil {
		err = fmt.Errorf("Failed to unmarshall recent states: %s", err)
		log.Errorf(err.Error())
		return state, err
	}
	return state, nil
}

// PUTRecentStatesToLedger marshals and writes the recent states
func PUTRecentStatesToLedger(stub shim.ChaincodeStubInterface, state RecentStates) error {
	var recentStatesJSON []byte
	var err error
	recentStatesJSON, err = json.Marshal(state)
	if err != nil {
		log.Criticalf("Failed to marshal recent states: %s", err)
		return err
	}
	err = stub.PutState(RECENTSTATESKEY, recentStatesJSON)
	if err != nil {
		log.Criticalf("Failed to PUTSTATE recent states: %s", err)
		return err
	}
	return nil
}

// ClearRecentStates resets recent states to an empty array
func ClearRecentStates(stub shim.ChaincodeStubInterface) error {
	var rstates RecentStates
	rstates = make(RecentStates, 0, MaxRecentStates)
	return PUTRecentStatesToLedger(stub, rstates)
}

// PushRecentState pushes the state to the first entry, or moves it to
// the first entry if this asset already shows up
func (a *Asset) PushRecentState(stub shim.ChaincodeStubInterface) error {
	var err error

	rstate, err := GETRecentStatesFromLedger(stub)
	if err != nil {
		return err
	}

	// shift slice to the right
	assetPosn, err := findAssetInRecent(a.AssetKey, rstate)
	if err != nil {
		return err
	} else if assetPosn == -1 {
		// grow if not at capacity, since this one is new
		if len(rstate) < MaxRecentStates {
			rstate = rstate[0 : len(rstate)+1]
		}
		// shift it all since not found
		copy(rstate[1:], rstate[0:])
	} else {
		if len(rstate) > 1 {
			// shift over top of the same asset, can appear only once
			copy(rstate[1:], rstate[0:assetPosn])
		}
	}
	rstate[0] = *a
	log.Debugf("pushRecentStates succeeded for asset %s", a.AssetKey)
	return PUTRecentStatesToLedger(stub, rstate)
}

// RemoveAssetFromRecentStates is called when an asset is deleted
func RemoveAssetFromRecentStates(stub shim.ChaincodeStubInterface, assetID string) error {
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
		if len(rstate) > 0 {
			// shift slice to the left to close the hole left by the asset
			copy(rstate[assetPosn:], rstate[assetPosn+1:])
		}
		if len(rstate) > 0 {
			rstate = rstate[0 : len(rstate)-1]
		}
	}
	return PUTRecentStatesToLedger(stub, rstate)
}

func findAssetInRecent(assetID string, rstate RecentStates) (int, error) {
	// returns -1 to signify not found (or error)
	for i := 0; i < len(rstate); i++ {
		if rstate[i].AssetKey == assetID {
			log.Debugf("findAssetInRecent found assetID %s at position %d in recent states", assetID, i)
			return i, nil
		}
	}
	// not found
	log.Debugf("findAssetInRecent Did not find assetID %s in recent states", assetID)
	return -1, nil
}

// readRecentStates returns the marshaled recent states from the ledger
var readRecentStates = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	r, err := GETRecentStatesFromLedger(stub)
	if err != nil {
		return nil, err
	}
	return json.Marshal(r)
}

func init() {
	AddRoute("readRecentStates", "query", SystemClass, readRecentStates)
}
