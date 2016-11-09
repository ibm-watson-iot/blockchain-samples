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

// v0.1 KL -- new iot chaincode platform

package cthistory

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Changing to prepend history key so that an asset's history is separated from it's
// current state.

// Logger for the cthistory package
var log = shim.NewLogger("hist")

// Enabled is false by default, import it into your main and set it to true
var Enabled bool

// STATEHISTORYKEY is used to separate history from current asset state and is prepended
// to the assetID
const STATEHISTORYKEY string = "HIST."

// AssetStateHistory is used to hold the array of states as strings.
type AssetStateHistory struct {
	AssetHistory []string `json:"assetHistory"`
}

// CreateStateHistory creates a new history entry in the ledger for an asset.
func CreateStateHistory(stub shim.ChaincodeStubInterface, assetID string, stateJSON string) error {

	var ledgerKey = STATEHISTORYKEY + assetID
	var assetStateHistory = AssetStateHistory{make([]string, 1)}
	assetStateHistory.AssetHistory[0] = stateJSON

	assetState, err := json.Marshal(&assetStateHistory)
	if err != nil {
		return err
	}

	return stub.PutState(ledgerKey, []byte(assetState))

}

// UpdateStateHistory adds a new state history for an asset. States are stored in the ledger
// in descending order by timestamp. AssetID is expected to by the *internal*
// assetID with a unique asset class prefix.
func UpdateStateHistory(stub shim.ChaincodeStubInterface, assetID string, stateJSON string) error {

	var ledgerKey = STATEHISTORYKEY + assetID
	var historyBytes []byte
	var assetStateHistory AssetStateHistory

	historyBytes, err := stub.GetState(ledgerKey)
	if err != nil {
		// assume that this is a new asset.
		return CreateStateHistory(stub, assetID, stateJSON)
	}

	err = json.Unmarshal(historyBytes, &assetStateHistory)
	if err != nil {
		// assume that history is corrupted, so reset.
		return CreateStateHistory(stub, assetID, stateJSON)
	}

	var newSlice = make([]string, 0)
	newSlice = append(newSlice, stateJSON)
	newSlice = append(newSlice, assetStateHistory.AssetHistory...)
	assetStateHistory.AssetHistory = newSlice

	assetState, err := json.Marshal(&assetStateHistory)
	if err != nil {
		return err
	}
	log.Debug("Update state history succedded for asset " + assetID)
	return stub.PutState(ledgerKey, []byte(assetState))

}

// DeleteStateHistory deletes all history for an asset from the ledger.
func DeleteStateHistory(stub shim.ChaincodeStubInterface, assetID string) error {
	var ledgerKey = STATEHISTORYKEY + assetID
	return stub.DelState(ledgerKey)
}

// ReadStateHistory gets the state history for an asset.
func ReadStateHistory(stub shim.ChaincodeStubInterface, assetID string) (AssetStateHistory, error) {

	var ledgerKey = STATEHISTORYKEY + assetID
	var assetStateHistory AssetStateHistory
	var historyBytes []byte

	historyBytes, err := stub.GetState(ledgerKey)
	if err != nil {
		return AssetStateHistory{}, err
	}

	err = json.Unmarshal(historyBytes, &assetStateHistory)
	if err != nil {
		return AssetStateHistory{}, err
	}

	return assetStateHistory, nil

}
