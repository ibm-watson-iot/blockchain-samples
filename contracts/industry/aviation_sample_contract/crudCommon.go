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

// v1 KL 08 Aug 2016 Separate crudUtils to their own module.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

// **************************************************
// CRUD utility functions
// **************************************************

// This is a multi-asset contract and so there is an *external* assetID that
// must be valid. What is returned is the *internal* assetID to be used to
// PUT and GET world state. This has a 2-char prefix denoting asset type.
// CRITICAL NOTE: The assetID inside the args or state maps is *never*
// adjusted so the state can always be returned directly to a caller.
func validateAssetID(caller string, eventName string, args interface{}) (string, error) {
	var assetID string
	var amargs ArgsMap
	var margs map[string]interface{}
	var found bool
	amargs, found = args.(ArgsMap)
	if !found {
		margs, found = args.(map[string]interface{})
		if !found {
			err := fmt.Errorf(caller+": validateAssetID: not passed a map, type is %T", args)
			log.Error(err)
			return "", err
		}
		amargs = ArgsMap(margs)
	}
	// is assetID present in common?
	assetID, found = getObjectAsString(amargs, ASSETID)
	if !found {
		// is assetID present in top level as parameter to e.g. read?
		assetID, found = getObjectAsString(amargs, "assetID")
	}
	if found {
		if assetID != "" {
			// we have something in assetID
			prefix, err := eventNameToAssetPrefix(eventName)
			if err != nil {
				err = errors.New(caller + ": validateAssetID: prefix not found, " + err.Error())
				log.Error(err)
				return "", err
			}
			log.Debug("validateAssetID: returning " + prefix + assetID)
			return prefix + assetID, nil
		}
		err := errors.New(caller + ": assetID is blank")
		log.Error(err)
		return "", err
	}
	// not found
	err := errors.New(caller + ": assetID is missing")
	log.Error(err)
	return "", err
}

// The events in this contract carry the assetID to which they pertain as a property
// of he event definition. The property is named as it appears in the raw data. So,
// for example, the maintanence event would carry an assemblyID property to be used
// as an external assetID and converted to an internal assetID for accessing the
// assembly's state.
func getEventAssetID(caller string, eventName string, assetIDProp string, event interface{}) (string, error) {
	var assetID string
	var amevent ArgsMap
	var mevent map[string]interface{}
	var found bool
	amevent, found = event.(ArgsMap)
	if !found {
		mevent, found = event.(map[string]interface{})
		if !found {
			err := fmt.Errorf(caller+": validateAssetID: not passed a map, type is %T", event)
			log.Error(err)
			return "", err
		}
		amevent = ArgsMap(mevent)
	}
	// is assetID present?
	assetID, found = getObjectAsString(amevent, assetIDProp)
	if !found {
		err := fmt.Errorf(caller + ": getEventAssetID: not found")
		log.Error(err)
		return "", err
	}
	if assetID != "" {
		// we have something in assetID
		prefix, err := eventNameToAssetPrefix(eventName)
		if err != nil {
			err = errors.New(caller + ": getEventAssetID: prefix not found, " + err.Error())
			log.Error(err)
			return "", err
		}
		log.Debug("getEventAssetID: returning " + prefix + assetID)
		return prefix + assetID, nil
	}
	err := errors.New(caller + ": assetID is blank")
	log.Error(err)
	return "", err
}

// In the multi-asset version of contract state, we no longer remember the asset list
// in memory, relying instead on retrieval from world state. This means, though, that
// this function expects the *internal* assetID, which prepends the assetName's
// 2-letter prefix.
func assetIsActive(stub *shim.ChaincodeStub, assetID string) ([]byte, error) {
	// no logging because it is used to determine presence or absence without
	// context
	stateBytes, err := stub.GetState(assetID)
	if err != nil {
		return nil, err
	}
	if len(stateBytes) == 0 {
		return nil, err
	}
	return stateBytes, nil
}

// Returns a map containing the JSON object represented by args[0]
func getUnmarshalledArgument(stub *shim.ChaincodeStub, caller string, args []string) (interface{}, error) {
	var event interface{}
	var err error

	if len(args) != 1 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	eventBytes := []byte(args[0])
	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		err = fmt.Errorf("%s failed to unmarshal arg: %s", caller, err)
		log.Error(err)
		return nil, err
	}

	if event == nil {
		err = fmt.Errorf("%s unmarshal arg created nil event", caller)
		log.Error(err)
		return nil, err
	}

	argsMap, found := event.(map[string]interface{})
	if !found {
		err := fmt.Errorf("%s arg is not a map shape", caller)
		log.Error(err)
		return nil, err
	}

	return argsMap, nil
}

// Returns the world state represented by prefix + assetID unmarshalled.
func getUnmarshalledState(stub *shim.ChaincodeStub, caller string, assetID string) (interface{}, error) {
	var ledgerBytes interface{}

	assetBytes, err := assetIsActive(stub, assetID)
	if err != nil || assetBytes == nil || len(assetBytes) == 0 {
		err = fmt.Errorf("%s: asset %s does not exist: %s", caller, assetID, err)
		log.Error(err)
		return nil, err
	}

	// unmarshal the existing state from the ledger to theinterface
	err = json.Unmarshal(assetBytes, &ledgerBytes)
	if err != nil {
		log.Errorf("%s: assetID %s unmarshal failed: %s", caller, assetID, err)
		return nil, err
	}

	// assert the existing state as a map
	ledgerMap, found := ledgerBytes.(map[string]interface{})
	if !found {
		log.Errorf("%s: assetID %s LEDGER state is not a map shape", caller, assetID)
		return nil, err
	}

	return ledgerMap, nil
}

// Pushes state to the ledger using assetID, which is expected to be prefixed.
func putMarshalledState(stub *shim.ChaincodeStub, caller string, eventName string, assetID string, state interface{}) error {
	amstate, found := state.(ArgsMap)
	if !found {
		mstate, found := state.(map[string]interface{})
		if !found {
			err := fmt.Errorf("%s: passed putMarshalledState non-map, type is %T", caller, state)
			log.Error(err)
			return err
		}
		amstate = ArgsMap(mstate)
	}
	// Write the new state to the ledger
	stateJSON, err := json.Marshal(&amstate)
	if err != nil {
		err = fmt.Errorf("%s: event %s assetID %s marshal failed: %s", caller, eventName, assetID, err)
		log.Error(err)
		return err
	}

	err = stub.PutState(assetID, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("%s: event %s assetID %s PUTSTATE failed: %s", caller, eventName, assetID, err)
		log.Error(err)
		return err
	}

	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("%s: event %s assetID %s push recent states failed: %s", caller, eventName, assetID, err)
		log.Error(err)
		return err
	}

	// add history state
	err = updateStateHistory(stub, assetID, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("%s: event %s assetID %s push history failed: %s", caller, eventName, assetID, err)
		log.Error(err)
		return err
	}
	return nil
}

// Pushes state to the ledger using assetID, which is expected to be prefixed.
func removeOneAssetFromWorldState(stub *shim.ChaincodeStub, caller string, assetName string, assetID string) error {
	err := stub.DelState(assetID)
	if err != nil {
		err = fmt.Errorf("%s: %s assetID %s deletion failed", caller, assetName, assetID)
		log.Error(err)
		return err
	}
	err = removeAssetFromRecentStates(stub, assetID)
	if err != nil {
		err = fmt.Errorf("%s: %s assetID %s deletion failed", caller, assetName, assetID)
		log.Error(err)
		return err
	}
	err = deleteStateHistory(stub, assetID)
	if err != nil {
		err = fmt.Errorf("%s: %s assetID %s history deletion failed", caller, assetName, assetID)
		log.Error(err)
		return err
	}
	return nil
}

// Adds the current transaction timestamp into world state, replacing it if it was already there.
func addTXNTimestampToState(stub *shim.ChaincodeStub, caller string, state interface{}) (interface{}, error) {
	amstate, found := state.(ArgsMap)
	if !found {
		mstate, found := state.(map[string]interface{})
		if !found {
			err := fmt.Errorf("%s: passed addTXNTimestampToState non-map, type is %T", caller, state)
			log.Error(err)
			return nil, err
		}
		amstate = ArgsMap(mstate)
	}
	// add transaction uuid and timestamp
	amstate[TXNUUID] = stub.UUID
	txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("%s: error getting transaction timestamp: %s", caller, err)
		log.Error(err)
		return nil, err
	}
	txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
	amstate[TXNTIMESTAMP] = txntimestamp
	return amstate, nil
}

// Fills in the "lastevent" property of asset state. Note that redirectedBy is rarely used in
// multi-asset contracts.
func addLastEventToState(stub *shim.ChaincodeStub, caller string, args interface{}, state interface{}, redirectedBy string) interface{} {
	amstate, found := state.(ArgsMap)
	if !found {
		mstate, found := state.(map[string]interface{})
		if !found {
			err := fmt.Errorf("%s: passed addLastEventToState non-map, type is %T", caller, state)
			log.Error(err)
			return nil
		}
		amstate = ArgsMap(mstate)
	}
	// save the original event
	amstate["lastEvent"] = make(map[string]interface{})
	amstate["lastEvent"].(map[string]interface{})["function"] = caller
	amstate["lastEvent"].(map[string]interface{})["arg"] = args
	if len(redirectedBy) > 0 {
		amstate["lastEvent"].(map[string]interface{})["redirectedFromFunction"] = redirectedBy
	}
	return amstate
}

// Executes the rules engine and returns the updated state.
func handleAlertsAndRules(stub *shim.ChaincodeStub, caller string, eventName string, assetID string, argsMap interface{}, state interface{}) (interface{}, error) {
	amargsMap, found := argsMap.(ArgsMap)
	if !found {
		margsMap, found := argsMap.(map[string]interface{})
		if !found {
			err := fmt.Errorf("%s: passed handleAlertsAndRules a non-map as args, type is %T", caller, argsMap)
			log.Error(err)
			return nil, err
		}
		amargsMap = ArgsMap(margsMap)
	}
	amstate, found := state.(ArgsMap)
	if !found {
		mstate, found := state.(map[string]interface{})
		if !found {
			err := fmt.Errorf("%s: passed handleAlertsAndRules non-map, type is %T", caller, state)
			log.Error(err)
			return nil, err
		}
		amstate = ArgsMap(mstate)
	}
	// run the rules and raise or clear alerts
	alerts := newAlertStatus()
	a, found := getObject(amstate, "alerts") // is there an existing alert state?
	if found {
		// convert to an AlertStatus, which does not work by type assertion
		log.Debugf("updateAsset Found existing alerts state: %s", a)
		// complex types are all untyped interfaces, so require conversion to
		// the structure that is used, but not in the other direction as the
		// type is properly specified
		alerts.alertStatusFromMap(a.(map[string]interface{}))
	}
	// important: rules need access to the entire calculated state
	noncompliant, err := amstate.executeRules(stub, eventName, &alerts, amargsMap)
	if err != nil {
		err = fmt.Errorf("%s: event %s has rules engine failure: %s", caller, eventName, err)
		log.Error(err)
		return nil, err
	}
	if noncompliant {
		log.Debugf("%s: event %s assetID %s is noncompliant", caller, eventName, assetID)
		amstate["alerts"] = alerts
		amstate["compliant"] = false
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(amstate, "alerts")
		} else {
			amstate["alerts"] = alerts
		}
		amstate["compliant"] = true
	}
	return amstate, nil
}

// ********** property injection implementation
func injectProps(state interface{}, qprops []QualifiedPropertyNameValue) (interface{}, error) {
	am, ok := asMap(state)
	if !ok {
		err := fmt.Errorf("injectProps passed a non-map of type %T", state)
		log.Error(err)
		return state, err
	}
	for _, qp := range qprops {
		am, ok := putObject(am, qp.QProp, qp.Value)
		if !ok {
			err := fmt.Errorf("injectProps->putObject failed to put %s:%s to state %#v", qp.QProp, qp.Value, am)
			log.Error(err)
			return am, err
		}
	}
	return am, nil
}

// a generic implementation to read everything in the database for debugging purposes ...
func (t *SimpleChaincode) readWorldState(stub *shim.ChaincodeStub) ([]byte, error) {
	var err error
	var results map[string]interface{}
	var state interface{}

	iter, err := stub.RangeQueryState("", "")
	if err != nil {
		err = fmt.Errorf("readWorldState failed to get a range query iterator: %s", err)
		log.Error(err)
		return nil, err
	}
	defer iter.Close()
	results = make(map[string]interface{})
	for iter.HasNext() {
		assetID, assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("readWorldState iter.Next() failed: %s", err)
			log.Error(err)
			return nil, err
		}
		err = json.Unmarshal(assetBytes, &state)
		if err != nil {
			err = fmt.Errorf("readWorldState unmarshal failed: %s", err)
			log.Error(err)
			return nil, err
		}
		results[assetID] = state
	}

	resultsBytes, err := json.MarshalIndent(&results, "", "    ")
	if err != nil {
		err = fmt.Errorf("readWorldState failed to marshal results: %s", err)
		log.Error(err)
		return nil, err
	}

	return resultsBytes, nil
}

// a generic implementation to clear everything out from the database for DEBUGGING purposes ...
func (t *SimpleChaincode) deleteWorldState(stub *shim.ChaincodeStub) error {
	iter, err := stub.RangeQueryState("", "")
	if err != nil {
		err = fmt.Errorf("clearWorldState failed to get a range query iterator: %s", err)
		log.Error(err)
		return err
	}
	defer iter.Close()
	for iter.HasNext() {
		assetID, _, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("clearWorldState iter.Next() failed: %s", err)
			log.Error(err)
			return err
		}
		// Delete the key / asset from the ledger
		err = stub.DelState(assetID)
		if err != nil {
			log.Errorf("deleteAsset assetID %s failed DELSTATE", assetID)
			return err
		}
	}
	log.Debug("\n\n********** WORLD STATE CLEARED *************\n\n")
	return nil
}
