/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Howard McKinney- Initial Contribution
Kim Letkeman - Initial Contribution
*/


// v3.0 HM 25 Feb 2016 Moved the asset state history code into a separate package.
// v3.0.1 HM 03 Mar 2016 Store the state history in descending order.  

package main

import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const STATEHISTORYKEY string = ".StateHistory"

type AssetStateHistory struct {
	AssetHistory []string `json:"assetHistory"`
}

// Create a new history entry in the ledger for an asset.,\
func createStateHistory(stub *shim.ChaincodeStub, assetID string, stateJSON string) error {

	var ledgerKey = assetID + STATEHISTORYKEY
	var assetStateHistory = AssetStateHistory{make([]string, 1)}
	assetStateHistory.AssetHistory[0] = stateJSON

	assetState, err := json.Marshal(&assetStateHistory)
	if err != nil {
		return err
	}

	return stub.PutState(ledgerKey, []byte(assetState))

}

// Update the ledger with new state history for an asset. States are stored in the ledger in descending order by timestamp.
func updateStateHistory(stub *shim.ChaincodeStub, assetID string, stateJSON string) error {

	var ledgerKey = assetID + STATEHISTORYKEY
	var historyBytes []byte
	var assetStateHistory AssetStateHistory
	
	historyBytes, err := stub.GetState(ledgerKey)
	if err != nil {
		return err
	}

	err = json.Unmarshal(historyBytes, &assetStateHistory)
	if err != nil {
		return err
	}

	var newSlice []string = make([]string, 0)
	newSlice = append(newSlice, stateJSON)
	newSlice = append(newSlice, assetStateHistory.AssetHistory...)
	assetStateHistory.AssetHistory = newSlice

	assetState, err := json.Marshal(&assetStateHistory)
	if err != nil {
		return err
	}

	return stub.PutState(ledgerKey, []byte(assetState))

}

// Delete an state history from the ledger.
func deleteStateHistory(stub *shim.ChaincodeStub, assetID string) error {

	var ledgerKey = assetID + STATEHISTORYKEY
	return stub.DelState(ledgerKey)

}

// Get the state history for an asset.
func readStateHistory(stub *shim.ChaincodeStub, assetID string) (AssetStateHistory, error) {

	var ledgerKey = assetID + STATEHISTORYKEY
	var assetStateHistory AssetStateHistory
	var historyBytes []byte

	historyBytes, err := stub.GetState(ledgerKey)
	if err != nil {
		return assetStateHistory, err
	}

	err = json.Unmarshal(historyBytes, &assetStateHistory)
	if err != nil {
		return assetStateHistory, err
	}

	return assetStateHistory, nil

}
