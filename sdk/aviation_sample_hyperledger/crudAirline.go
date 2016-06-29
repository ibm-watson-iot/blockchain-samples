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

/*

   The aviation contract models four crud interfaces and five objects. The four
   primary objects are airline, airplane, assembly and part. The object that is
   not meant to have its own crud interface is activity, which is an object that
   can be passed into any of the other crud interfaces as part of the event. So,
   for example, in this file the updateAirline function will handle a change to
   the basic state of the airline, and at the same time it can handle the addition
   of a new activity.

*/

// 17 June 2016 KL  Created.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
    "strings"
    "reflect"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// IOTCRUDIF is the receiver for airline specific functions
type IOTCRUDIF struct {
	StatePrefix string
	JSONTag     string
}

var airline = IOTCRUDIF{StatePrefix: "AL", JSONTag: "airline"}

//***************************************************
//***************************************************
//* AIRLINE CRUD INTERFACE
//***************************************************
//***************************************************

// ************************************
// createAirline
// ************************************
func (t *SimpleChaincode) createAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var argsMap ArgsMap
	var event interface{}
	var found bool
	var err error

	log.Info("Entering createAirline")

	// allowing 2 args because updateAirline is allowed to redirect when
	// airline does not exist
	if len(args) < 1 || len(args) > 2 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	airlineID = ""
	eventBytes := []byte(args[0])
	log.Debugf("createAirline arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("createAirline failed to unmarshal arg: %s", err)
		return nil, err
	}

	if event == nil {
		err = errors.New("createAirline unmarshal arg created nil event")
		log.Error(err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("createAirline arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("createAirline arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}

	found = airlineIsActive(stub, airlineID)
	if found {
		err := fmt.Errorf("createAirline arg airline %s already exists", airlineID)
		log.Error(err)
		return nil, err
	}

	// add transaction uuid and timestamp
	argsMap[TXNUUID] = stub.UUID
	txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("Error getting transaction timestamp: %s", err)
		log.Error(err)
		return nil, err
	}
	txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
	argsMap[TXNTIMESTAMP] = txntimestamp

	// copy incoming event to outgoing state
	// this contract respects the fact that createAirline can accept a partial state
	// as the moral equivalent of one or more discrete events
	// further: this contract understands that its schema has two discrete objects
	// that are meant to be used to send events: common, and custom
	stateOut := argsMap

	// save the original event
	stateOut["txnEvent"] = make(map[string]interface{})
	stateOut["txnEvent"].(map[string]interface{})["function"] = "createAirline"
	stateOut["txnEvent"].(map[string]interface{})["args"] = args[0]
	if len(args) == 2 {
		// in-band protocol for redirect
		stateOut["txnEvent"].(map[string]interface{})["redirectedFromFunction"] = args[1]
	}

	// run the rules and raise or clear alerts
	alerts := newAlertStatus()
	noncompliant, err := stateOut.executeRules(&alerts)
	if err != nil {
		err = fmt.Errorf("Rules engine failure: %s", err)
		log.Error(err)
		return nil, err
	}
	if noncompliant {
		log.Noticef("createAirline airlineID %s is noncompliant", airlineID)
		stateOut["alerts"] = alerts
		delete(stateOut, "compliant")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(stateOut, "alerts")
		} else {
			stateOut["alerts"] = alerts
		}
		stateOut["compliant"] = true
	}

	// marshal to JSON and write
	stateJSON, err := json.Marshal(&stateOut)
	if err != nil {
		err := fmt.Errorf("createAirline state for airlineID %s failed to marshal", airlineID)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	log.Infof("Putting new airline state %s to ledger", string(stateJSON))
	err = stub.PutState(airlineID, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("createAirline AirlineID %s PUTSTATE failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}
	log.Infof("createAirline AirlineID %s state successfully written to ledger: %s", airlineID, string(stateJSON))

	// add airline to contract state
	err = addAirlineToContractState(stub, airlineID)
	if err != nil {
		err := fmt.Errorf("createAirline airline %s failed to write airline state: %s", airlineID, err)
		log.Critical(err)
		return nil, err
	}

	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("createAirline AirlineID %s push to recentstates failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// save state history
	err = createStateHistory(stub, airlineID, string(stateJSON))
	if err != nil {
		err := fmt.Errorf("createAirline airline %s state history save failed: %s", airlineID, err)
		log.Critical(err)
		return nil, err
	}

	return nil, nil
}

// ************************************
// updateAirline
// ************************************
func (t *SimpleChaincode) updateAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var argsMap ArgsMap
	var event interface{}
	var ledgerMap ArgsMap
	var ledgerBytes interface{}
	var found bool
	var err error

	log.Info("Entering updateAirline")

	if len(args) != 1 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	airlineID = ""
	eventBytes := []byte(args[0])
	log.Debugf("updateAirline arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("updateAirline failed to unmarshal arg: %s", err)
		return nil, err
	}

	if event == nil {
		err = errors.New("createAirline unmarshal arg created nil event")
		log.Error(err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("updateAirline arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("updateAirline arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}
	log.Noticef("updateAirline found airlineID %s", airlineID)

	found = airlineIsActive(stub, airlineID)
	if !found {
		// redirect to createAirline with same parameter list
		if canCreateOnUpdate(stub) {
			log.Noticef("updateAirline redirecting airline %s to createAirline", airlineID)
			var newArgs = []string{args[0], "updateAirline"}
			return t.createAirline(stub, newArgs)
		}
		err = fmt.Errorf("updateAirline airline %s does not exist", airlineID)
		log.Error(err)
		return nil, err
	}

	// add transaction uuid and timestamp
	argsMap[TXNUUID] = stub.UUID
	txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("Error getting transaction timestamp: %s", err)
		log.Error(err)
		return nil, err
	}
	txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
	argsMap[TXNTIMESTAMP] = txntimestamp

	// **********************************
	// find the airline state in the ledger
	// **********************************
	log.Infof("updateAirline: retrieving airline %s state from ledger", airlineID)
	airlineBytes, err := stub.GetState(airlineID)
	if err != nil {
		log.Errorf("updateAirline airlineID %s GETSTATE failed: %s", airlineID, err)
		return nil, err
	}

	// unmarshal the existing state from the ledger to theinterface
	err = json.Unmarshal(airlineBytes, &ledgerBytes)
	if err != nil {
		log.Errorf("updateAirline airlineID %s unmarshal failed: %s", airlineID, err)
		return nil, err
	}

	// assert the existing state as a map
	ledgerMap, found = ledgerBytes.(map[string]interface{})
	if !found {
		log.Errorf("updateAirline airlineID %s LEDGER state is not a map shape", airlineID)
		return nil, err
	}

	// now add incoming map values to existing state to merge them
	// this contract respects the fact that updateAirline can accept a partial state
	// as the moral equivalent of one or more discrete events
	// further: this contract understands that its schema has two discrete objects
	// that are meant to be used to send events: common, and custom
	// ledger has to have common section
	stateOut := deepMerge(map[string]interface{}(argsMap),
		map[string]interface{}(ledgerMap))
	log.Debugf("updateAirline airlineID %s merged state: %s", airlineID, stateOut)

	// save the original event
	stateOut["txnEvent"] = make(map[string]interface{})
	stateOut["txnEvent"].(map[string]interface{})["function"] = "updateAirline"
	stateOut["txnEvent"].(map[string]interface{})["args"] = args[0]

	// handle compliance section
	alerts := newAlertStatus()
	a, found := stateOut["alerts"] // is there an existing alert state?
	if found {
		// convert to an AlertStatus, which does not work by type assertion
		log.Debugf("updateAirline Found existing alerts state: %s", a)
		// complex types are all untyped interfaces, so require conversion to
		// the structure that is used, but not in the other direction as the
		// type is properly specified
		alerts.alertStatusFromMap(a.(map[string]interface{}))
	}
	// important: rules need access to the entire calculated state
	noncompliant, err := ledgerMap.executeRules(&alerts)
	if err != nil {
		err = fmt.Errorf("Rules engine failure: %s", err)
		log.Error(err)
		return nil, err
	}
	if noncompliant {
		// true means noncompliant
		log.Noticef("updateAirline airlineID %s is noncompliant", airlineID)
		// update ledger with new state, if all clear then delete
		stateOut["alerts"] = alerts
		delete(stateOut, "compliant")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(stateOut, "alerts")
		} else {
			stateOut["alerts"] = alerts
		}
		stateOut["compliant"] = true
	}

	// Write the new state to the ledger
	stateJSON, err := json.Marshal(ledgerMap)
	if err != nil {
		err = fmt.Errorf("updateAirline AirlineID %s marshal failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	err = stub.PutState(airlineID, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAirline AirlineID %s PUTSTATE failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}
	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAirline AirlineID %s push to recentstates failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// add history state
	err = updateStateHistory(stub, airlineID, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAirline AirlineID %s push to history failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// NOTE: Contract state is not updated by updateAirline

	return nil, nil
}

// ************************************
// deleteAirline
// ************************************
func (t *SimpleChaincode) deleteAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var argsMap ArgsMap
	var event interface{}
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("Expecting one JSON state object with an airlineID")
		log.Error(err)
		return nil, err
	}

	airlineID = ""
	eventBytes := []byte(args[0])
	log.Debugf("deleteAirline arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("deleteAirline failed to unmarshal arg: %s", err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("deleteAirline arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("deleteAirline arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}

	found = airlineIsActive(stub, airlineID)
	if !found {
		err = fmt.Errorf("deleteAirline airlineID %s does not exist", airlineID)
		log.Error(err)
		return nil, err
	}

	// Delete the key / airline from the ledger
	err = stub.DelState(airlineID)
	if err != nil {
		log.Errorf("deleteAirline airlineID %s failed DELSTATE", airlineID)
		return nil, err
	}
	// remove airline from contract state
	err = removeAirlineFromContractState(stub, airlineID)
	if err != nil {
		err := fmt.Errorf("deleteAirline airline %s failed to remove airline from contract state: %s", airlineID, err)
		log.Critical(err)
		return nil, err
	}
	// save state history
	err = deleteStateHistory(stub, airlineID)
	if err != nil {
		err := fmt.Errorf("deleteAirline airline %s state history delete failed: %s", airlineID, err)
		log.Critical(err)
		return nil, err
	}
	// push the recent state
	err = removeAirlineFromRecentState(stub, airlineID)
	if err != nil {
		err := fmt.Errorf("deleteAirline airline %s recent state removal failed: %s", airlineID, err)
		log.Critical(err)
		return nil, err
	}

	return nil, nil
}

// ************************************
// deletePropertiesFromAirline
// ************************************
func (t *SimpleChaincode) deletePropertiesFromAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var argsMap ArgsMap
	var event interface{}
	var ledgerMap ArgsMap
	var ledgerBytes interface{}
	var found bool
	var err error
	var alerts AlertStatus

	if len(args) < 1 {
		err = errors.New("Not enough arguments. Expecting one JSON object with mandatory AirlineID and property name array")
		log.Error(err)
		return nil, err
	}
	eventBytes := []byte(args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Error("deletePropertiesFromAirline failed to unmarshal arg")
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("deletePropertiesFromAirline arg is not a map shape")
		log.Error(err)
		return nil, err
	}
	log.Debugf("deletePropertiesFromAirline arg: %+v", argsMap)

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("deletePropertiesFromAirline arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}

	found = airlineIsActive(stub, airlineID)
	if !found {
		err = fmt.Errorf("deletePropertiesFromAirline airlineID %s does not exist", airlineID)
		log.Error(err)
		return nil, err
	}

	// is there a list of property names?
	var qprops []interface{}
	qpropsBytes, found := getObject(argsMap, "qualPropsToDelete")
	if found {
		qprops, found = qpropsBytes.([]interface{})
		log.Debugf("deletePropertiesFromAirline qProps: %+v, Found: %+v, Type: %+v", qprops, found, reflect.TypeOf(qprops))
		if !found || len(qprops) < 1 {
			log.Errorf("deletePropertiesFromAirline airline %s qualPropsToDelete is not an array or is empty", airlineID)
			return nil, err
		}
	} else {
		log.Errorf("deletePropertiesFromAirline airline %s has no qualPropsToDelete argument", airlineID)
		return nil, err
	}

	// **********************************
	// find the airline state in the ledger
	// **********************************
	log.Infof("deletePropertiesFromAirline: retrieving airline %s state from ledger", airlineID)
	airlineBytes, err := stub.GetState(airlineID)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s GETSTATE failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// unmarshal the existing state from the ledger to the interface
	err = json.Unmarshal(airlineBytes, &ledgerBytes)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s unmarshal failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// assert the existing state as a map
	ledgerMap, found = ledgerBytes.(map[string]interface{})
	if !found {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s LEDGER state is not a map shape", airlineID)
		log.Error(err)
		return nil, err
	}

	// now remove properties from state, they are qualified by level
OUTERDELETELOOP:
	for p := range qprops {
		prop := qprops[p].(string)
		log.Debugf("deletePropertiesFromAirline AirlineID %s deleting qualified property: %s", airlineID, prop)
		// TODO Ugly, isolate in a function at some point
		if (CASESENSITIVEMODE && strings.HasSuffix(prop, airline.JSONTag)) ||
			(!CASESENSITIVEMODE && strings.HasSuffix(strings.ToLower(prop), strings.ToLower(airline.JSONTag))) {
			log.Warningf("deletePropertiesFromAirline AirlineID %s cannot delete protected qualified property: %s", airlineID, prop)
		} else {
			levels := strings.Split(prop, ".")
			lm := (map[string]interface{})(ledgerMap)
			for l := range levels {
				// lev is the name of a level
				lev := levels[l]
				if l == len(levels)-1 {
					// we're here, delete the actual property name from this level of the map
					levActual, found := findMatchingKey(lm, lev)
					if !found {
						log.Warningf("deletePropertiesFromAirline AirlineID %s property match %s not found", airlineID, lev)
						continue OUTERDELETELOOP
					}
					log.Debugf("deletePropertiesFromAirline AirlineID %s deleting %s", airlineID, prop)
					delete(lm, levActual)
				} else {
					// navigate to the next level object
					log.Debugf("deletePropertiesFromAirline AirlineID %s navigating to level %s", airlineID, lev)
					lmBytes, found := findObjectByKey(lm, lev)
					if found {
						lm, found = lmBytes.(map[string]interface{})
						if !found {
							log.Noticef("deletePropertiesFromAirline AirlineID %s level %s not found in ledger", airlineID, lev)
							continue OUTERDELETELOOP
						}
					}
				}
			}
		}
	}
	log.Debugf("updateAirline AirlineID %s final state: %s", airlineID, ledgerMap)

	// add transaction uuid and timestamp
	ledgerMap[TXNUUID] = stub.UUID
	txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("Error getting transaction timestamp: %s", err)
		log.Error(err)
		return nil, err
	}
	txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
	ledgerMap[TXNTIMESTAMP] = txntimestamp

	// save the original event
	ledgerMap["txnEvent"] = make(map[string]interface{})
	ledgerMap["txnEvent"].(map[string]interface{})["function"] = "deletePropertiesFromAirline"
	ledgerMap["txnEvent"].(map[string]interface{})["args"] = args[0]

	// handle compliance section
	alerts = newAlertStatus()
	a, found := ledgerMap["alerts"] // is there an existing alert state?
	if found {
		// convert to an AlertStatus, which does not work by type assertion
		log.Debugf("deletePropertiesFromAirline Found existing alerts state: %s", a)
		// complex types are all untyped interfaces, so require conversion to
		// the structure that is used, but not in the other direction as the
		// type is properly specified
		alerts.alertStatusFromMap(a.(map[string]interface{}))
	}
	// important: rules need access to the entire calculated state
	noncompliant, err := ledgerMap.executeRules(&alerts)
	if err != nil {
		err = fmt.Errorf("Rules engine failure: %s", err)
		log.Error(err)
		return nil, err
	}
	if noncompliant {
		// true means noncompliant
		log.Noticef("deletePropertiesFromAirline airlineID %s is noncompliant", airlineID)
		// update ledger with new state, if all clear then delete
		ledgerMap["alerts"] = alerts
		delete(ledgerMap, "compliant")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(ledgerMap, "alerts")
		} else {
			ledgerMap["alerts"] = alerts
		}
		ledgerMap["compliant"] = true
	}

	// Write the new state to the ledger
	stateJSON, err := json.Marshal(ledgerMap)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s marshal failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	err = stub.PutState(airlineID, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s PUTSTATE failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}
	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s push to recentstates failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// add history state
	err = updateStateHistory(stub, airlineID, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAirline AirlineID %s push to history failed: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	return nil, nil
}

// ************************************
// deleteAllAirlines
// ************************************
func (t *SimpleChaincode) deleteAllAirlines(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var err error

	if len(args) > 0 {
		err = errors.New("Too many arguments. Expecting none.")
		log.Error(err)
		return nil, err
	}

	aa, err := getActiveAirlines(stub)
	if err != nil {
		err = fmt.Errorf("deleteAllAirlines failed to get the active airlines: %s", err)
		log.Error(err)
		return nil, err
	}
	for i := range aa {
		airlineID = aa[i]

		// Delete the key / airline from the ledger
		err = stub.DelState(airlineID)
		if err != nil {
			err = fmt.Errorf("deleteAllAirlines arg %d airlineID %s failed DELSTATE", i, airlineID)
			log.Error(err)
			return nil, err
		}
		// remove airline from contract state
		err = removeAirlineFromContractState(stub, airlineID)
		if err != nil {
			err = fmt.Errorf("deleteAllAirlines airline %s failed to remove airline from contract state: %s", airlineID, err)
			log.Critical(err)
			return nil, err
		}
		// save state history
		err = deleteStateHistory(stub, airlineID)
		if err != nil {
			err := fmt.Errorf("deleteAllAirlines airline %s state history delete failed: %s", airlineID, err)
			log.Critical(err)
			return nil, err
		}
	}
	err = clearRecentStates(stub)
	if err != nil {
		err = fmt.Errorf("deleteAllAirlines clearRecentStates failed: %s", err)
		log.Error(err)
		return nil, err
	}
	return nil, nil
}

// ************************************
// readAirline
// ************************************
func (t *SimpleChaincode) readAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var argsMap ArgsMap
	var request interface{}
	var airlineBytes []byte
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	requestBytes := []byte(args[0])
	log.Debugf("readAirline arg: %s", args[0])

	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		log.Errorf("readAirline failed to unmarshal arg: %s", err)
		return nil, err
	}

	argsMap, found = request.(map[string]interface{})
	if !found {
		err := errors.New("readAirline arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("readAirline arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}

	found = airlineIsActive(stub, airlineID)
	if !found {
		err := fmt.Errorf("readAirline arg airline %s does not exist", airlineID)
		log.Error(err)
		return nil, err
	}

	// Get the state from the ledger
	airlineBytes, err = stub.GetState(airlineID)
	if err != nil {
		log.Errorf("readAirline airlineID %s failed GETSTATE", airlineID)
		return nil, err
	}

	return airlineBytes, nil
}

// ************************************
// readAllAirlines
// ************************************
func (t *SimpleChaincode) readAllAirlines(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineID string
	var err error
	var results []interface{}
	var state interface{}

	if len(args) > 0 {
		err = errors.New("readAllAirlines expects no arguments")
		log.Error(err)
		return nil, err
	}

	aa, err := getActiveAirlines(stub)
	if err != nil {
		err = fmt.Errorf("readAllAirlines failed to get the active airlines: %s", err)
		log.Error(err)
		return nil, err
	}
	results = make([]interface{}, 0, len(aa))
	for i := range aa {
		airlineID = aa[i]
		// Get the state from the ledger
		airlineBytes, err := stub.GetState(airlineID)
		if err != nil {
			// best efforts, return what we can
			log.Errorf("readAllAirlines airlineID %s failed GETSTATE", airlineID)
			continue
		} else {
			err = json.Unmarshal(airlineBytes, &state)
			if err != nil {
				// best efforts, return what we can
				log.Errorf("readAllAirlines airlineID %s failed to unmarshal", airlineID)
				continue
			}
			results = append(results, state)
		}
	}

	resultsStr, err := json.Marshal(results)
	if err != nil {
		err = fmt.Errorf("readallAirlines failed to marshal results: %s", err)
		log.Error(err)
		return nil, err
	}

	return []byte(resultsStr), nil
}

// ************************************
// readAirlineHistory
// ************************************
func (t *SimpleChaincode) readAirlineHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var airlineBytes []byte
	var airlineID string
	var argsMap ArgsMap
	var request interface{}
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("readAirlineHistory expects a JSON encoded object with airlineID and count")
		log.Error(err)
		return nil, err
	}

	requestBytes := []byte(args[0])
	log.Debugf("readAirlineHistory arg: %s", args[0])

	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		err = fmt.Errorf("readAirlineHistory failed to unmarshal arg: %s", err)
		log.Error(err)
		return nil, err
	}

	argsMap, found = request.(map[string]interface{})
	if !found {
		err := errors.New("readAirlineHistory arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is airlineID present or blank?
	airlineIDBytes, found := getObject(argsMap, airline.JSONTag)
	if found {
		airlineID, found = airlineIDBytes.(string)
		if !found || airlineID == "" {
			err := errors.New("readAirlineHistory arg does not include airlineID")
			log.Error(err)
			return nil, err
		}
	}

	found = airlineIsActive(stub, airlineID)
	if !found {
		err := fmt.Errorf("readAirlineHistory arg airline %s does not exist", airlineID)
		log.Error(err)
		return nil, err
	}

	// Get the history from the ledger
	stateHistory, err := readStateHistory(stub, airlineID)
	if err != nil {
		err = fmt.Errorf("readAirlineHistory airlineID %s failed readStateHistory: %s", airlineID, err)
		log.Error(err)
		return nil, err
	}

	// is count present?
	var olen int
	countBytes, found := getObject(argsMap, "count")
	if found {
		olen = int(countBytes.(float64))
	}
	if olen <= 0 || olen > len(stateHistory.AirlineHistory) {
		olen = len(stateHistory.AirlineHistory)
	}
	var hStatesOut = make([]interface{}, 0, olen)
	for i := 0; i < olen; i++ {
		var obj interface{}
		err = json.Unmarshal([]byte(stateHistory.AirlineHistory[i]), &obj)
		if err != nil {
			log.Errorf("readAirlineHistory JSON unmarshal of entry %d failed [%#v]", i, stateHistory.AirlineHistory[i])
			return nil, err
		}
		hStatesOut = append(hStatesOut, obj)
	}
	airlineBytes, err = json.Marshal(hStatesOut)
	if err != nil {
		log.Errorf("readAirlineHistory failed to marshal results: %s", err)
		return nil, err
	}

	return []byte(airlineBytes), nil
}
