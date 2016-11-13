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
// v0.2 KL -- dramatic reduction in memory and disk space by storing only the keys
//            and reading the states only when queried, raised limit to 100, added
//            range to query

package iotcontractplatform

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
// a begin.

// RECENTSTATESKEY is used as key for recent states bucket
const RECENTSTATESKEY string = "IOTCP.RecentStates"

// MaxRecentStates is an arbitrary limit on how many asset states we track across the
// entire contract
const MaxRecentStates int = 40

// RecentStates in world state
type RecentStates struct {
	States []string `json:"recentstates"`
}

// RecentStatesOut is query output format
type RecentStatesOut AssetArray

// GETRecentStatesFromLedger returns the unmarshaled recent states
func GETRecentStatesFromLedger(stub shim.ChaincodeStubInterface) (RecentStates, error) {
	var rstates = RecentStates{make([]string, 0, MaxRecentStates+1)}
	var err error
	recentStatesBytes, err := stub.GetState(RECENTSTATESKEY)
	if err != nil {
		err = fmt.Errorf("Failed to get recent states from world state: %s", err)
		log.Errorf(err.Error())
		return rstates, err
	}
	// this MUST be here
	if recentStatesBytes == nil || len(recentStatesBytes) == 0 {
		log.Debugf("GETRecentStatesFromLedger: returning empty recent states")
		return rstates, nil
	}
	err = json.Unmarshal(recentStatesBytes, &rstates)
	if err != nil {
		err = fmt.Errorf("Failed to unmarshall recent states: %s", err)
		log.Errorf(err.Error())
		return rstates, err
	}
	return rstates, nil
}

// PUTRecentStatesToLedger marshals and writes the recent states
func PUTRecentStatesToLedger(stub shim.ChaincodeStubInterface, rstates RecentStates) error {
	var recentStatesJSON []byte
	var err error
	recentStatesJSON, err = json.Marshal(rstates)
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
	rstates = RecentStates{make([]string, 0, MaxRecentStates)}
	return PUTRecentStatesToLedger(stub, rstates)
}

// PushRecentState pushes the state to the first entry, or moves it to
// the first entry if this asset already shows up
func (a *Asset) PushRecentState(stub shim.ChaincodeStubInterface) error {
	var err error

	rstates, err := GETRecentStatesFromLedger(stub)
	if err != nil {
		return err
	}

	// shift slice to the right
	assetPosn := findAssetInRecent(a.AssetKey, rstates)
	if assetPosn == -1 {
		// shift right
		rstates.States = append(rstates.States, a.AssetKey)
		copy(rstates.States[1:], rstates.States[0:])
	} else {
		// shift right to close the gap
		copy(rstates.States[1:], rstates.States[0:assetPosn])
	}
	// insert at the front
	rstates.States[0] = a.AssetKey
	log.Debugf("pushRecentStates succeeded for asset %s", a.AssetKey)
	return PUTRecentStatesToLedger(stub, rstates)
}

// RemoveAssetFromRecentStates is called when an asset is deleted
func (a *Asset) RemoveAssetFromRecentStates(stub shim.ChaincodeStubInterface) error {
	var rstates RecentStates
	var err error

	rstates, err = GETRecentStatesFromLedger(stub)
	if err != nil {
		return err
	}
	posn := findAssetInRecent(a.AssetKey, rstates)
	if posn >= 0 {
		rstates.States = append(rstates.States[:posn], rstates.States[posn+1:]...)
	}
	return PUTRecentStatesToLedger(stub, rstates)
}

func findAssetInRecent(assetID string, rstates RecentStates) int {
	// returns -1 to signify not found (or error)
	for i := 0; i < len(rstates.States); i++ {
		if rstates.States[i] == assetID {
			log.Debugf("findAssetInRecent found assetID %s at position %d in recent states", assetID, i)
			return i
		}
	}
	// not found
	log.Debugf("findAssetInRecent Did not find assetID %s in recent states", assetID)
	return -1
}

// readRecentStates returns the marshaled recent states from the ledger
var readRecentStates = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var begin, end, count int
	var found bool

	if len(args) > 0 {
		var arg map[string]interface{}
		eventBytes := []byte(args[0])
		err := json.Unmarshal(eventBytes, &arg)
		if err != nil {
			err = fmt.Errorf("readRecentStates: failed to unmarshal args[0] ' %s': %s", args[0], err)
			log.Error(err)
			return nil, err
		}
		begin, found = GetObjectAsInteger(&arg, "begin")
		if !found {
			begin = 0
		}
		end, found = GetObjectAsInteger(&arg, "end")
		if !found {
			end = MaxRecentStates - begin - 1
		}
	} else {
		end = MaxRecentStates - 1
	}

	r, err := GETRecentStatesFromLedger(stub)
	if err != nil {
		err = fmt.Errorf("readRecentStates: failed to get recent states from ledger: %s", err)
		log.Error(err)
		return nil, err
	}

	if begin >= len(r.States) {
		err := fmt.Errorf("readRecentStates: begin position %d beyond end of recent states, last state is position %d", begin, len(r.States)-1)
		log.Error(err)
		return nil, err
	}

	if end >= len(r.States) {
		end = len(r.States) - 1
	}

	count = end - begin + 1

	var rstatesout = make(RecentStatesOut, 0, count)
	for i := begin; i <= end; i++ {
		a, exists, err := GetAssetFromLedger(stub, r.States[i])
		if err != nil {
			err = fmt.Errorf("readRecentStates: failed to get asset from ledger: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		if !exists {
			err = fmt.Errorf("readRecentStates: recent asset state does not exist: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		rstatesout = append(rstatesout, a)
	}
	return json.Marshal(rstatesout)
}

func init() {
	AddRoute("readRecentStates", "query", SystemClass, readRecentStates)
}
