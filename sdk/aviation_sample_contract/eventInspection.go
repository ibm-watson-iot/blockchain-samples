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

// v1 KL 12 Aug 2016 Implement inspection event

// Inspection Event
// This event targets assembly assets.
// For the assembly:
//    "aCheckCounter" is set to zero and rules are called to clear ACHECK alert
//    "bCheckCounter" is set to zero and rules are called to clear BCHECK alert

package main

import (
	//"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func eventInspection(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	event, err := getUnmarshalledArgument(stub, "eventInspection", args)
	if err != nil {
		return nil, err
	}
	_, err = handleAssemblyInspectionEvent(stub, event)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func handleAssemblyInspectionEvent(stub *shim.ChaincodeStub, event interface{}) (interface{}, error) {
	assetID, err := getEventAssetID("handleAssemblyInspectionEvent", "inspection", "inspection.assembly", event)
	if err != nil {
		return nil, err
	}

	state, err := getUnmarshalledState(stub, "handleAssemblyInspectionEvent", assetID)
	if err != nil {
		return nil, err
	}

	state, err = addTXNTimestampToState(stub, "handleAssemblyInspectionEvent", state)
	if err != nil {
		return nil, err
	}

	state = addLastEventToState(stub, "handleAssemblyInspectionEvent", event, state, "")

	// the rules will clear the appropriate alerts
	state, err = handleAlertsAndRules(stub, "handleAssemblyInspectionEvent", "inspection", assetID, event, state)
	if err != nil {
		return nil, err
	}

	err = putMarshalledState(stub, "handleAssemblyInspectionEvent", "inspection", assetID, state)
	if err != nil {
		return nil, err
	}

	return state.(ArgsMap), nil
}
