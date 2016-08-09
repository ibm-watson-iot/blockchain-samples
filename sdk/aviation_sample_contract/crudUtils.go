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
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// **************************************************
// CRUD utility functions
// **************************************************

// This is a multi-asset contract and so there is an eventName in the schema that
// must be valid. All further processing is driven by eventName and the related
// assetName. Note that this property can be a pure event or a partial state for 
// an asset. This too is denoted in the schema and implemented in contractConfig.go 
func validateEventName(caller string, args ArgsMap) (string, error) {
    enBytes, found := getObject(args, EVENTNAME)
    if found {
        en, found := enBytes.(string)
        if found {
            if isEventName(en) {
                // we're good! return the eventName for convenience
                return en, nil
            }
            // what is this event?
            err := errors.New(caller + ": eventName is unknown: " + en)
            log.Error(err)
            return "", err
        }
        // not a string
        err := errors.New(caller + ": eventName is not a string shape")
        log.Error(err)
        return "", err
    }
    // not found
    err := errors.New(caller + ": eventName is missing")
    log.Error(err)
    return "", err
}

// Validation of eventName allowed both partial state asset events and pure
// events that are *about* assets. This validation of assetName allows only
// creatable assets to pass.  
func validateAssetName(caller string, args ArgsMap) (string, error) {
    enBytes, found := getObject(args, EVENTNAME)
    if found {
        en, found := enBytes.(string)
        if found {
            if isAssetName(en) {
                // we're good! return the eventName for convenience
                return en, nil
            }
            // what is this event?
            err := errors.New(caller + ": eventName is unknown: " + en)
            log.Error(err)
            return "", err
        }
        // not a string
        err := errors.New(caller + ": eventName is not a string shape")
        log.Error(err)
        return "", err
    }
    // not found
    err := errors.New(caller + ": eventName is missing")
    log.Error(err)
    return "", err
}

// This is a multi-asset contract and so there is an *external* assetID that
// must be valid. What is returned is the *internal* assetID to be used to 
// PUT and GET world state. This has a 2-char prefix denoting asset type.
// CRITICAL NOTE: The assetID inside the map is *never* adjusted so it can 
// always be returned directly to a caller.
func validateAssetID(caller string, eventName string, args ArgsMap) (string, error) {
    // is assetID present or blank?
    assetIDBytes, found := getObject(args, ASSETID)
    if found {
        assetID, found := assetIDBytes.(string) 
        if found {
            if assetID != "" {
                // we have something in assetID
                prefix, err := eventNameToAssetPrefix(eventName)
                if err != nil {
                    err = errors.New(caller + ": prefix not found, " + err.Error())
                    log.Error(err)
                    return "", err
                }
                return prefix + assetID, nil
            }
            err := errors.New(caller + ": assetID is blank")
            log.Error(err)
            return "", err
        }
        // not a string
        err := errors.New(caller + ": assetID is not a string shape")
        log.Error(err)
        return "", err
    }
    // not found
    err := errors.New(caller + ": assetID is missing")
    log.Error(err)
    return "", err
}

func getUnmarshalledArgument(stub *shim.ChaincodeStub, caller string, args []string) (ArgsMap, error) {
    var event interface{}
    var err error

	if len(args) != 1 {
        err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}
    
    log.Debugf("%1 arg: %s", caller, args[0])

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

func getUnmarshalledState(stub *shim.ChaincodeStub, caller string, eventName string, assetID string) (ArgsMap, error) {
    var ledgerBytes interface{}

    prefix, err := eventNameToAssetPrefix(eventName)
    if err != nil {
        err = fmt.Errorf("%s: eventName %s does not exist", caller, eventName)
        log.Error(err)
        return nil, err
    }

    if len(assetID) < 3 || assetID[0:2] != prefix {
        assetID = prefix + assetID
    }

    assetBytes, err := assetIsActive(stub, assetID)
    if err != nil {
        err = fmt.Errorf("%s: asset %s does not exist: %s", caller, assetID, err.Error())
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

func putMarshalledState(stub *shim.ChaincodeStub, caller string, eventName string, assetID string, state ArgsMap) (error) {
    // Write the new state to the ledger
    stateJSON, err := json.Marshal(state)
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

func addTXNTimestampToState(stub *shim.ChaincodeStub, caller string, state ArgsMap) (ArgsMap, error) {
    // add transaction uuid and timestamp
    state[TXNUUID] = stub.UUID
    txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("%s: error getting transaction timestamp: %s", caller, err)
        log.Error(err)
        return nil, err
	}
    txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
    state[TXNTIMESTAMP] = txntimestamp
    return state, nil
}

func addLastEventToState(stub *shim.ChaincodeStub, caller string, args ArgsMap, state ArgsMap, redirectedBy string) (ArgsMap) {
    // save the original event
    state["lastEvent"] = make(map[string]interface{})
    state["lastEvent"].(map[string]interface{})["function"] = "caller"
    state["lastEvent"].(map[string]interface{})["arg"] = args
    if len(redirectedBy) > 0 {
        state["lastEvent"].(map[string]interface{})["redirectedFromFunction"] = redirectedBy
    }
    return state
}

func handleAlertsAndRules(stub *shim.ChaincodeStub, caller string, eventName string, assetID string, argsMap ArgsMap, state ArgsMap) (ArgsMap, error) {
    // run the rules and raise or clear alerts
    alerts := newAlertStatus()

    noncompliant, err := state.executeRules(&alerts)
    if err != nil {
		err = fmt.Errorf("%s: event %s has rules engine failure: %s", caller, eventName, err)
        log.Error(err)
        return nil, err
    }
    if noncompliant {
        log.Noticef("%s: event %s assetID %s is noncompliant", caller, eventName, assetID)
        state["alerts"] = alerts
        delete(state, "compliant")
    } else {
        if alerts.AllClear() {
            // all false, no need to appear
            delete(state, "alerts")
        } else {
            state["alerts"] = alerts
        }
        state["compliant"] = true
    }
    return state, nil
}