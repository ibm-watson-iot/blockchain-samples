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

// v3.0 HM 25 Feb 2016 Moved the asset state history code into a separate package.
// v3.0.1 HM 03 Mar 2016 Store the state history in descending order.
// v3.0.2 KL 07 Mar 2016 Reduce memory garbage in updateStateHistory
// v3.0.3 KL             backported from original 3.1/4.0
// v4.0 KL 17 Mar 2016 Update version number to 4.0 for Hyperledger compatibility.
//                     Clean up lint issues.
// v4.3 KL August 2016 Remove activeAssets array as it is not useful in a multi-asset
//                     contract. World state is polled instead.
// v5.0 KL 02 Nov 2016 package the platform, this package is state related

package ctstate

import (
    "encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

// CONTRACTSTATEKEY is used to store contract state, including version, nickname and activeAssets
const CONTRACTSTATEKEY string = "ContractStateKey"

// ContractState struct defines contract state. Unlike the main contract maps, structs work fine
// for this fixed structure.
type ContractState struct {
    Version  string `json:"version"`
    Nickname string `json:"nickname"`
}

// Logger for the ctstate package
var log = shim.NewLogger("stat")

// GETContractStateFromLedger retrieves state from ledger and returns to caller
func GETContractStateFromLedger(stub shim.ChaincodeStubInterface) (ContractState, error) {
    var state = ContractState{}
    var err error
    contractStateBytes, err := stub.GetState(CONTRACTSTATEKEY)
    // minimum string is {"version":""} and version cannot be empty
    if err == nil && len(contractStateBytes) > 14 {
        // apparently, this blockchain instance is being reloaded, has the version changed?
        err = json.Unmarshal(contractStateBytes, &state)
        if err != nil {
            err = fmt.Errorf("Unmarshal failed for contract state: %s", err)
            log.Criticalf(err.Error())
            return ContractState{}, err
        }
    } else {
        // empty state already initialized
        log.Noticef("Initialized newly deployed contract state version %s", state.Version)
    }
    log.Debugf("GETContractState successful")
    return state, nil
}

// PUTContractStateToLedger writes a contract state into the ledger
func PUTContractStateToLedger(stub shim.ChaincodeStubInterface, state ContractState) error {
    var contractStateJSON []byte
    var err error
    contractStateJSON, err = json.Marshal(state)
    if err != nil {
        err = fmt.Errorf("Failed to marshal contract state: %s", err)
        log.Criticalf(err.Error())
        return err
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        err = fmt.Errorf("Failed to PUTSTATE contract state: %s", err)
        log.Criticalf(err.Error())
        return err
    }
    log.Debugf("PUTContractState: %#v", state)
    return nil
}

// InitializeContractState verifies the version passed by the deploy message, and
// writes an initial contract state into world state.
func InitializeContractState(stub shim.ChaincodeStubInterface, version string, nickname string) error {
    state, err := GETContractStateFromLedger(stub)
    if err != nil {
        err = fmt.Errorf("Initialize contract state failed to get contract state from ledger: %s", err)
        log.Errorf(err.Error())
        return err
    }
    if version != state.Version {
        log.Noticef("Contract version has changed from %s to %s", state.Version, version)
    }
    state.Version = version
    state.Nickname = nickname
    return PUTContractStateToLedger(stub, state)
}

// GetLedgerContractVersion returns the deployed contract version
func GetLedgerContractVersion(stub shim.ChaincodeStubInterface) (string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)
    if err != nil {
        return "", err
    }
    return state.Version, nil
}
