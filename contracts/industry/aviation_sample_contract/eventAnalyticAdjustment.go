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

// v1 KL Aug 2016 Add analytics adjustment event

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func eventAnalyticAdjustment(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	event, err := getUnmarshalledArgument(stub, "eventAnalyticAdjustment", args)
	if err != nil {
		return nil, err
	}

	_, err = handleAssemblyAnalyticAdjustmentEvent(stub, event)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func handleAssemblyAnalyticAdjustmentEvent(stub *shim.ChaincodeStub, event interface{}) (interface{}, error) {
	assetID, err := getEventAssetID("handleAssemblyAnalyticAdjustmentEvent", "analyticAdjustment", "analyticAdjustment.assembly", event)
	if err != nil {
		return nil, err
	}

	state, err := getUnmarshalledState(stub, "handleAssemblyAnalyticAdjustmentEvent", assetID)
	if err != nil {
		return nil, err
	}

	// adjust the life limit according to instructions
	state, err = processAnalyticAdjustmentAction(stub, state, event, assetID)
	if err != nil {
		return nil, err
	}

	// state will be of type interface{} for use with crudUtils
	state, err = addTXNTimestampToState(stub, "handleAssemblyAnalyticAdjustmentEvent", state)
	if err != nil {
		return nil, err
	}

	state = addLastEventToState(stub, "handleAssemblyAnalyticAdjustmentEvent", event, state, "")

	state, err = handleAlertsAndRules(stub, "handleAssemblyAnalyticAdjustmentEvent", "analyticAdjustment", assetID, event, state)
	if err != nil {
		return nil, err
	}

	err = putMarshalledState(stub, "handleAssemblyAnalyticAdjustmentEvent", "analyticAdjustment", assetID, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func processAnalyticAdjustmentAction(stub *shim.ChaincodeStub, state interface{}, event interface{}, assetID string) (interface{}, error) {
	action, found := getObjectAsString(event, "analyticAdjustment.action")
	if !found {
		err := fmt.Errorf("processAnalyticAdjustmentAction: action property not found in event: %+v", event)
		log.Error(err)
		return nil, err
	}

	amount, found := getObjectAsNumber(event, "analyticAdjustment.amount")
	if !found {
		err := fmt.Errorf("processAnalyticAdjustmentAction: amount property not found in event: %+v", event)
		log.Error(err)
		return nil, err
	}

	// get the three counters and their adjusted variants

	cycles, found := getObjectAsNumber(state, "cycles")
	if !found {
		err := fmt.Errorf("processAnalyticAdjustmentAction: cycles property not found in state, possible out of order event: %+v", state)
		log.Error(err)
		return nil, err
	}

	adjustedCycles, found := getObjectAsNumber(state, "adjustedCycles")
	if !found {
		// first adjustment
		adjustedCycles = cycles
	}

	acc, found := getObjectAsNumber(state, "aCheckCounter")
	if !found {
		err := fmt.Errorf("processAnalyticAdjustmentAction: aCheckCounter property not found in state, possible out of order event: %+v", state)
		log.Error(err)
		return nil, err
	}
	acca, found := getObjectAsNumber(state, "aCheckCounterAdjusted")
	if !found {
		// first cycle count
		acca = acc
	}

	bcc, found := getObjectAsNumber(state, "bCheckCounter")
	if !found {
		err := fmt.Errorf("processAnalyticAdjustmentAction: bCheckCounter property not found in state, possible out of order event: %+v", state)
		log.Error(err)
		return nil, err
	}
	bcca, found := getObjectAsNumber(state, "bCheckCounterAdjusted")
	if !found {
		// first cycle count
		bcca = bcc
	}

	// adjust all of the adjusted variants

	switch action {
	case "adjustLifeLimit":
		adjustedCycles += amount
		acca += amount
		bcca += amount
	default:
		err := fmt.Errorf("processAnalyticAdjustmentAction: unknown action property: %s", action)
		log.Error(err)
		return nil, err
	}

	// put all of the adjusted variants into the state

	state, ok := putObject(state, "adjustedCycles", adjustedCycles)
	if !ok {
		err := fmt.Errorf("processAnalyticAdjustmentAction: adjustedCycles property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	state, ok = putObject(state, "aCheckCounterAdjusted", acca)
	if !ok {
		err := fmt.Errorf("processAnalyticAdjustmentAction: aCheckCounterAdjusted property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	state, ok = putObject(state, "bCheckCounterAdjusted", bcca)
	if !ok {
		err := fmt.Errorf("processAnalyticAdjustmentAction: bCheckCounterAdjusted property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	return state, nil
}
