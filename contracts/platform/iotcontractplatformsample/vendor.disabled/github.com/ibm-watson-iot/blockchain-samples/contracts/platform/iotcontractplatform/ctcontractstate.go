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

package iotcontractplatform

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CONTRACTSTATEKEY is used to store contract state, including version, nickname and activeAssets
const CONTRACTSTATEKEY string = "IOTCP:ContractState"

// ContractState struct defines contract state. Unlike the main contract maps, structs work fine
// for this fixed structure.
type ContractState struct {
	Version  string `json:"version"`
	Nickname string `json:"nickname"`
}

// GETContractStateFromLedger retrieves state from ledger and returns to caller
func GETContractStateFromLedger(stub shim.ChaincodeStubInterface) (ContractState, error) {
	var err error
	var state ContractState
	contractStateBytes, err := stub.GetState(CONTRACTSTATEKEY)
	if err != nil {
		err = fmt.Errorf("GETContractStateFromLedger returned error: %s", err)
		log.Notice(err)
		return ContractState{}, err
	}
	if contractStateBytes == nil || len(contractStateBytes) == 0 {
		err = errors.New("GETContractStateFromLedger contract state does not exist")
		log.Notice(err)
		return ContractState{}, err
	}
	// this contract instance is being reloaded
	err = json.Unmarshal(contractStateBytes, &state)
	if err != nil {
		err = fmt.Errorf("Unmarshal failed for contract state: %s", err)
		log.Critical(err)
		return ContractState{}, err
	}
	return state, nil
}

// PUTContractStateToLedger writes a contract state into the ledger
func PUTContractStateToLedger(stub shim.ChaincodeStubInterface, state ContractState) error {
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
	return nil
}

// InitializeContractState sets version and nickname back to defaults
func InitializeContractState(stub shim.ChaincodeStubInterface, contractversion string, nicknamearg string, versionarg string) error {
	var state ContractState
	var err error
	if versionarg != contractversion {
		err = fmt.Errorf("Deploy version argument: %s MUST match contract version: %s", versionarg, contractversion)
		log.Critical(err)
		return err
	}
	state, err = GETContractStateFromLedger(stub)
	if err != nil {
		// assume new
		log.Noticef("Deployed contract version %s appears to be newly deployed", versionarg)
		state.Version = contractversion
		state.Nickname = nicknamearg
	} else {
		log.Noticef("Deployed contract version %s appears to be redeployed", versionarg)
	}
	if contractversion != state.Version {
		log.Noticef("Deployed contract version has changed from %s to %s", versionarg, contractversion)
	}
	return PUTContractStateToLedger(stub, state)
}

var readContractState = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 0 {
		err = errors.New("Too many arguments. Expecting none.")
		log.Errorf(err.Error())
		return nil, err
	}

	// Get the state from the ledger
	chaincodeBytes, err := stub.GetState(CONTRACTSTATEKEY)
	if err != nil {
		err = fmt.Errorf("readContractState failed GETSTATE: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}

	return chaincodeBytes, nil
}

var initContract = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var stateArg ContractState
	var err error

	log.Infof("Entering initContract with args: %+v", args)

	if len(args) != 2 {
		err = errors.New("initContract expects two arguments, a JSON object with version and nickname, and the contract version")
		log.Criticalf(err.Error())
		return nil, err
	}

	err = json.Unmarshal([]byte(args[0]), &stateArg)
	if err != nil {
		err = fmt.Errorf("initContract argument unmarshal failed: %s", err)
		log.Criticalf(err.Error())
		return nil, err
	}

	if stateArg.Nickname == "" {
		stateArg.Nickname = "IOTSampleContractII"
	}

	// the framework requires that args[1] be set by the router from the contract version passed from main
	err = InitializeContractState(stub, args[1], stateArg.Nickname, stateArg.Version)
	if err != nil {
		return nil, err
	}

	log.Infof("initContract - contract initialized")
	return nil, nil
}

func init() {
	AddRoute("readContractState", "query", SystemClass, readContractState)
	AddRoute("initContract", "deploy", SystemClass, initContract)
}
