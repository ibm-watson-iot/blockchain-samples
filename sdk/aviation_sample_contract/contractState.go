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

package main

import (
	"encoding/json"
    "fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// MYVERSION Update for every change, use VX.X.X (Major, Minor, Fix). Suggest that we update
// Major for API break, Minor when adding a feature or behavior, Fix when fixing a bug.
// If the init comes in with the wrong major version, then  we might consider exiting with
// an error.
const MYVERSION string = "4.3"

// DEFAULTNICKNAME is used when a contract is initialized without giving it a nickname
const DEFAULTNICKNAME string = "AVIATION_SAMPLE" 

// CONTRACTSTATEKEY is used to store contract state, including version, nickname and activeAssets
const CONTRACTSTATEKEY string = "ContractStateKey"

// ContractState struct defines contract state. Unlike the main contract maps, structs work fine
// for this fixed structure.
type ContractState struct {
	Version      string           `json:"version"`
    Nickname     string           `json:"nickname"`
}

// GETContractStateFromLedger retrieves state from ledger and returns to caller
func GETContractStateFromLedger(stub *shim.ChaincodeStub) (ContractState, error) {
    var state = ContractState{ MYVERSION, DEFAULTNICKNAME }
    var err error
	contractStateBytes, err := stub.GetState(CONTRACTSTATEKEY)
    // minimum string is {"version":""} and version cannot be empty 
	if err == nil && len(contractStateBytes) > 14 {    
		// apparently, this blockchain instance is being reloaded, has the version changed?
		err = json.Unmarshal(contractStateBytes, &state)
		if err != nil {
            err = fmt.Errorf("Unmarshal failed for contract state: %s", err)
            log.Critical(err)
			return ContractState{}, err
		}
        if MYVERSION != state.Version {
            log.Noticef("Contract version has changed from %s to %s", state.Version, MYVERSION)
            state.Version = MYVERSION
        }
	} else {
        // empty state already initialized 
		log.Noticef("Initialized newly deployed contract state version %s", state.Version)
	}
    log.Debug("GETContractState successful")
    return state, nil 
}

// PUTContractStateToLedger writes a contract state into the ledger
func PUTContractStateToLedger(stub *shim.ChaincodeStub, state ContractState) (error) {
    var contractStateJSON []byte
    var err error
    contractStateJSON, err = json.Marshal(state)
    if err != nil {
        err = fmt.Errorf("Failed to marshal contract state: %s", err)
        log.Critical(err)
        return err
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        err = fmt.Errorf("Failed to PUTSTATE contract state: %s", err)
        log.Critical(err)
        return err
    } 
    log.Debugf("PUTContractState: %#v", state)
    return nil 
}

func initializeContractState(stub *shim.ChaincodeStub, version string, nickname string) (error) {
    var state ContractState
    var err error
    if version != MYVERSION {
        err = fmt.Errorf("Contract version: %s does not match version argument: %s", MYVERSION, version)
        log.Critical(err)
        return err
    }
    state, err = GETContractStateFromLedger(stub)
    if err != nil {
        return err
    }  
    if version != state.Version {
        log.Noticef("Contract version has changed from %s to %s", version, MYVERSION)
        // keep going, this is an update of version -- later this will
        // be handled by pulling state from the superseded contract version
    }
    state.Version = MYVERSION
    state.Nickname = nickname
    return PUTContractStateToLedger(stub, state)
}

func getLedgerContractVersion(stub *shim.ChaincodeStub) (string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return "", err
    }
    return state.Version, nil   
}

// In the multi-asset version of contract state, we no longer remember the asset list
// in memory, relying instead on retrieval from world state. This means, though, that
// this function expects the *internal* assetID, which prepends the assetName's 
// 2-letter prefix.
func assetIsActive(stub *shim.ChaincodeStub, assetID string) ([]byte, error) {
    stateBytes, err := stub.GetState(assetID)
    if err != nil { 
        err = fmt.Errorf("assetIsActive: assetID %s: %s %s", assetID, string(stateBytes), err)
        log.Error(err)
        return nil, err
    }
    return stateBytes, nil
}
