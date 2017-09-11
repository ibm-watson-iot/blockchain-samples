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
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var log = shim.NewLogger("iotcontractplatform")

// SetContractLogger allows the whole package to be loaded at startup and to share a
// single chaincode logger
func SetContractLogger(logger *shim.ChaincodeLogger) {
	log = logger
}

// CREATEONFIRSTUPDATEKEY is used to store can create on update status, which if true by default
const CREATEONFIRSTUPDATEKEY string = "IOTCP:CreateOnFirstUpdate"

// readWorldState read everything in the database for debugging purposes ...
var readWorldState ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var results map[string]interface{}
	var state interface{}

	iter, err := stub.GetStateByRange("", "")
	if err != nil {
		err = fmt.Errorf("readWorldState failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	results = make(map[string]interface{})
	for iter.HasNext() {
		//assetID, assetBytes, err := iter.Next()
		assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("readWorldState iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		err = json.Unmarshal(assetBytes.Value, &state)
		if err != nil {
			err = fmt.Errorf("readWorldState unmarshal failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		//results[state.assetID] = state
		results[assetBytes.Key] = state
		//append(everything.Assets, asset)
		//results[assetID] = stateappend(everything.Assets, state)
	}

	resultsBytes, err := json.MarshalIndent(&results, "", "    ")
	if err != nil {
		err = fmt.Errorf("readWorldState failed to marshal results: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}

	log.Debugf(string(resultsBytes))

	return resultsBytes, nil
}

// deleteWorldState clear everything out from the database for DEBUGGING purposes ...
var deleteWorldState ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// obtain the current contract config and reinitialize the contract later as if just
	// deployed (saves developer time)
	cstate, _ := GETContractStateFromLedger(stub)

	iter, err := stub.GetStateByRange("", "")
	if err != nil {
		err = fmt.Errorf("clearWorldState failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		//assetID, _, err := iter.Next()
		assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("clearWorldState iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		// Delete the key / asset from the ledger
		//err = stub.DelState(state.assetID)
		err = stub.DelState(assetBytes.Key)
		if err != nil {
			log.Errorf("deleteAsset assetID %s failed DELSTATE", assetBytes.Key)
			return nil, err
		}
	}
	log.Debugf("\n\n********** WORLD STATE CLEARED *************\n\n")
	if len(args) > 0 && args[0] == "reinit" {
		time.Sleep(300)
		InitializeContractState(stub, cstate.Version, cstate.Version, cstate.Nickname)
		log.Debugf("\n\n********** WORLD STATE REINITIALIZED *************\n\n")
	}
	return nil, nil
}

var setLoggingLevel ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	type LogLevelArg struct {
		Level string `json:"logLevel"`
	}
	var level LogLevelArg
	var err error
	if len(args) != 1 {
		err = errors.New("Incorrect number of arguments. Expecting a JSON encoded LogLevel.")
		log.Errorf(err.Error())
		return nil, err
	}
	err = json.Unmarshal([]byte(args[0]), &level)
	if err != nil {
		err = fmt.Errorf("setLoggingLevel failed to unmarshal arg: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}

	switch level.Level {
	case "DEBUG":
		log.SetLevel(shim.LogDebug)
	case "INFO":
		log.SetLevel(shim.LogInfo)
	case "NOTICE":
		log.SetLevel(shim.LogNotice)
	case "WARNING":
		log.SetLevel(shim.LogWarning)
	case "ERROR":
		log.SetLevel(shim.LogError)
	case "CRITICAL":
		log.SetLevel(shim.LogCritical)
	default:
		err = fmt.Errorf("setLoggingLevel failed with unknown arg: %s", level.Level)
		log.Errorf(err.Error())
		return nil, err
	}

	return nil, nil
}

// CreateOnFirstUpdate is a shared parameter structure for the use of
// the createonupdate feature
type CreateOnFirstUpdate struct {
	SetCreateOnFirstUpdate bool `json:"setCreateOnFirstUpdate"`
}

// ************************************
// setCreateOnFirstUpdate
// ************************************
var setCreateOnFirstUpdate ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var createOnFirstUpdate CreateOnFirstUpdate
	var err error
	if len(args) != 1 {
		err = errors.New("setCreateOnFirstUpdate expects a single parameter")
		log.Errorf(err.Error())
		return nil, err
	}
	err = json.Unmarshal([]byte(args[0]), &createOnFirstUpdate)
	if err != nil {
		err = fmt.Errorf("setCreateOnFirstUpdate failed to unmarshal arg: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	err = PUTcreateOnFirstUpdate(stub, createOnFirstUpdate)
	if err != nil {
		err = fmt.Errorf("setCreateOnFirstUpdate failed to PUT setting: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	return nil, nil
}

// PUTcreateOnFirstUpdate marshals the new setting and writes it to the ledger
func PUTcreateOnFirstUpdate(stub shim.ChaincodeStubInterface, createOnFirstUpdate CreateOnFirstUpdate) (err error) {
	createOnFirstUpdateBytes, err := json.Marshal(createOnFirstUpdate)
	if err != nil {
		err = errors.New("PUTcreateOnFirstUpdate failed to marshal")
		log.Errorf(err.Error())
		return err
	}
	err = stub.PutState(CREATEONFIRSTUPDATEKEY, createOnFirstUpdateBytes)
	if err != nil {
		err = fmt.Errorf("PUTSTATE createOnFirstUpdate failed: %s", err)
		log.Errorf(err.Error())
		return err
	}
	return nil
}

// CanCreateOnFirstUpdate retrieves the setting from the ledger and returns it to the calling function
func CanCreateOnFirstUpdate(stub shim.ChaincodeStubInterface) bool {
	var createOnFirstUpdate CreateOnFirstUpdate
	createOnFirstUpdateBytes, err := stub.GetState(CREATEONFIRSTUPDATEKEY)
	if err != nil {
		err = fmt.Errorf("GETSTATE for canCreateOnFirstUpdate failed: %s", err)
		log.Errorf(err.Error())
		return true // true is the default
	}
	err = json.Unmarshal(createOnFirstUpdateBytes, &createOnFirstUpdate)
	if err != nil {
		err = fmt.Errorf("canCreateOnFirstUpdate failed to marshal: %s", err)
		log.Errorf(err.Error())
		return true // true is the default
	}
	return createOnFirstUpdate.SetCreateOnFirstUpdate
}

func init() {
	AddRoute("deleteWorldState", "invoke", SystemClass, deleteWorldState)
	AddRoute("readWorldState", "query", SystemClass, readWorldState)
	AddRoute("setLoggingLevel", "invoke", SystemClass, setLoggingLevel)
	AddRoute("setCreateOnFirstUpdate", "invoke", SystemClass, setCreateOnFirstUpdate)
}
