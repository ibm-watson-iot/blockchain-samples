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
// v2 KL 02 Nov 2016 new package ctasset

package ctasset

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	a "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctalerts"
	cf "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctconfig"
	h "github.com/ibm-watson-iot/blockchain-samples/iotbase/cthistory"
	r "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctrecent"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctstate"
	"time"
)

// **************************************************
// CRUD utility functions
// **************************************************

// TXNTIMESTAMP is the JSON tag for transaction timestamps, which map directly onto the transaction in the blockchain
const TXNTIMESTAMP string = "txntimestamp"

// TXNUUID is the JSON tag for transaction UUIDs, which map directly onto the transaction in the blockchain
const TXNUUID string = "txnuuid"

// This is a multi-asset contract and so there is an *external* assetID that
// must be valid. What is returned is the *internal* assetID to be used to
// PUT and GET world state. This has a 2-char prefix denoting asset type.
// CRITICAL NOTE: The assetID inside the args or state maps is *never*
// adjusted so the state can always be returned directly to a caller.
func validateAssetID(caller string, eventName string, args interface{}) (string, error) {
	var assetID string
	var found bool
	amargs, found := st.AsMap(args)
	if !found {
		err := fmt.Errorf(caller+": validateAssetID: not passed a map, type is %T", args)
		log.Error(err)
		return "", err
	}
	// is assetID present in common?
	assetID, found = st.GetObjectAsString(amargs, "common.assetID")
	if !found {
		// is assetID present in top level as parameter to e.g. read?
		assetID, found = st.GetObjectAsString(amargs, "assetID")
	}
	if found {
		if assetID != "" {
			// we have something in assetID
			prefix, err := cf.EventNameToAssetPrefix(eventName)
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
	var found bool
	amevent, found := st.AsMap(event)
	if !found {
		err := fmt.Errorf(caller+": validateAssetID: not passed a map, type is %T", event)
		log.Error(err)
		return "", err
	}
	// is assetID present?
	assetID, found = st.GetObjectAsString(amevent, assetIDProp)
	if !found {
		err := fmt.Errorf(caller + ": getEventAssetID: not found")
		log.Error(err)
		return "", err
	}
	if assetID != "" {
		// we have something in assetID
		prefix, err := cf.EventNameToAssetPrefix(eventName)
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
		err = fmt.Errorf("assetIsActive failed for ID %s with error %s", assetID, err)
		log.Debug(err)
		return nil, err
	}
	// seems to be needed in recent builds, appears to happen after asset deletion
	if len(stateBytes) == 0 {
		err = fmt.Errorf("assetIsActive failed for ID %s with length of state == 0", assetID)
		log.Error(err)
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
	amstate, found := st.AsMap(state)
	if !found {
		err := fmt.Errorf("%s: passed putMarshalledState non-map, type is %T", caller, state)
		log.Error(err)
		return err
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

	err = r.PushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("%s: event %s assetID %s push recent states failed: %s", caller, eventName, assetID, err)
		log.Error(err)
		return err
	}

	// add history state
	err = h.UpdateStateHistory(stub, assetID, string(stateJSON))
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
	err = r.RemoveAssetFromRecentStates(stub, assetID)
	if err != nil {
		err = fmt.Errorf("%s: %s assetID %s deletion failed", caller, assetName, assetID)
		log.Error(err)
		return err
	}
	err = h.DeleteStateHistory(stub, assetID)
	if err != nil {
		err = fmt.Errorf("%s: %s assetID %s history deletion failed", caller, assetName, assetID)
		log.Error(err)
		return err
	}
	return nil
}

// Adds the current transaction timestamp into world state, replacing it if it was already there.
func addTXNTimestampToState(stub *shim.ChaincodeStub, caller string, state interface{}) (interface{}, error) {
	amstate, found := st.AsMap(state)
	if !found {
		err := fmt.Errorf("%s: passed addTXNTimestampToState non-map, type is %T", caller, state)
		log.Error(err)
		return nil, err
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
	amstate, found := st.AsMap(state)
	if !found {
		err := fmt.Errorf("%s: passed addLastEventToState non-map, type is %T", caller, state)
		log.Error(err)
		return nil
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
	amargsMap, found := st.AsMap(argsMap)
	if !found {
		err := fmt.Errorf("%s: passed handleAlertsAndRules a non-map as args, type is %T", caller, argsMap)
		log.Error(err)
		return nil, err
	}
	amstate, found := st.AsMap(state)
	if !found {
		err := fmt.Errorf("%s: passed handleAlertsAndRules non-map, type is %T", caller, state)
		log.Error(err)
		return nil, err
	}
	alerts := a.NewAlertStatus()
	al, found := st.GetObject(amstate, "alerts") // is there an existing alert state?
	if found {
		log.Debugf("updateAsset Found existing alerts state: %s", al)
		alerts = a.AlertStatusFromMap(al.(map[string]interface{}), alerts)
	}
	// important: rules need access to the entire calculated state
	amstate, alertactive, err := a.ExecuteRules(stub, eventName, &alerts, amstate, amargsMap)
	if err != nil {
		err = fmt.Errorf("%s: event %s has rules engine failure: %s", caller, eventName, err)
		log.Error(err)
		return nil, err
	}
	if alertactive {
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
func injectProps(state interface{}, qprops []QPropNV) (interface{}, error) {
	am, ok := st.AsMap(state)
	if !ok {
		err := fmt.Errorf("injectProps passed a non-map of type %T", state)
		log.Error(err)
		return state, err
	}
	for _, qp := range qprops {
		am, ok := st.PutObject(am, qp.QProp, qp.Value)
		if !ok {
			err := fmt.Errorf("injectProps->putObject failed to put %s:%s to state %#v", qp.QProp, qp.Value, am)
			log.Error(err)
			return am, err
		}
	}
	return am, nil
}

// ReadWorldState read everything in the database for debugging purposes ...
func ReadWorldState(stub *shim.ChaincodeStub) ([]byte, error) {
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

// DeleteWorldState clear everything out from the database for DEBUGGING purposes ...
func DeleteWorldState(stub *shim.ChaincodeStub) error {
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
	time.Sleep(300)
	// now obtain the current contract config and reinitialize the contract as if just deployed to save developer time
	cstate, err := st.GETContractStateFromLedger(stub)
	if err != nil {
		log.Debug("\n\n********** WORLD STATE REINITIALIZATION FAILED *************\nPlease kill the chaincode, restart, and send Deploy to reinitialize.\n\n")
	}
	st.InitializeContractState(stub, cstate.Version, cstate.Nickname)
	log.Debug("\n\n********** WORLD STATE REINITIALIZED *************\n\n")
	return nil
}
