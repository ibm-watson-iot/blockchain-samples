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

// v1 KL 10 Aug 2016 Add flight event

// Flight Event
// This event makes changes to both aircraft assets and assembly assets.
// For the aircraft to which the flight applies, "cycles" is incremented by one.
// For the assemblies attached to said aircraft:
//    "cycles" and "adjustedCycles" are incremented
//    "aCheckCounter" and "aCheckCounterAdjusted" are incremented
//    "bCheckCounter" and "bCheckCounterAdjusted" are incremented

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func eventFlight(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	event, err := getUnmarshalledArgument(stub, "eventFlight", args)
	if err != nil {
		return nil, err
	}

	aircraftState, err := handleAircraftFlightEvent(stub, event)
	if err != nil {
		return nil, err
	}

	aircraftID, found := getObjectAsString(aircraftState, "common.assetID")
	if !found {
		err := errors.New("eventFlight: can't find assetID in returned aircraft state")
		log.Error(err)
		return nil, err
	}

	aircraftID, err = assetIDToInternal("flight", aircraftID)
	if err != nil {
		err := fmt.Errorf("eventFlight: failed to convert assetID to internal: %s", aircraftID)
		log.Error(err)
		return nil, err
	}

	indexes, err := getAircraftAssemblyIndexesFromLedger(stub)
	if err != nil {
		return nil, err
	}

	// propagate to assemblies
	assemblies, found := indexes.getAircraftAssemblies(aircraftID)
	log.Debugf("eventFlight: propagating aircraft %s to %d assemblies", aircraftID, len(assemblies))
	if found {
		for assetID := range assemblies {
			_, err := handleAssemblyFlightEvent(stub, event, assetID)
			if err != nil {
				return nil, err
			}
		}
	}

	// no need to put indexes to ledger since we used them in read only mode

	return nil, nil
}

func handleAircraftFlightEvent(stub shim.ChaincodeStubInterface, event interface{}) (interface{}, error) {
	log.Debugf("handleAircraftFlightEvent: %+v", event)
	// get the aircraft ledger state
	assetID, err := getEventAssetID("handleAircraftFlightEvent", "flight", "flight.aircraft", event)
	if err != nil {
		return nil, err
	}

	// the assetID should be an internal aircraft ID
	state, err := getUnmarshalledState(stub, "handleAircraftFlightEvent", assetID)
	if err != nil {
		return nil, err
	}

	// count cycles for the life of the aircraft
	cycles, found := getObjectAsNumber(state, "cycles")
	if found {
		cycles++
	} else {
		cycles = 1
	}

	state, ok := putObject(state, "cycles", cycles)
	if !ok {
		err := errors.New("handleAircraftFlightEvent: putObject failed for cycles")
		log.Error(err)
		return nil, err
	}

	// count adjusted cycles for the life of the aircraft
	adjustedCycles, found := getObjectAsNumber(state, "adjustedCycles")
	if found {
		adjustedCycles++
	} else {
		adjustedCycles = 1
	}

	state, ok = putObject(state, "adjustedCycles", adjustedCycles)
	if !ok {
		err := errors.New("handleAircraftFlightEvent: putObject failed for adjustedCycles")
		log.Error(err)
		return nil, err
	}

	state, err = addTXNTimestampToState(stub, "handleAircraftFlightEvent", state)
	if err != nil {
		return nil, err
	}

	state = addLastEventToState(stub, "handleAircraftFlightEvent", event, state, "")

	state, err = handleAlertsAndRules(stub, "handleAircraftFlightEvent", "flight", assetID, event, state)
	if err != nil {
		return nil, err
	}

	err = putMarshalledState(stub, "handleAircraftFlightEvent", "flight", assetID, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func handleAssemblyFlightEvent(stub shim.ChaincodeStubInterface, event interface{}, assetID string) (interface{}, error) {
	var ok bool

	log.Debugf("handleAssemblyFlightEvent: %+v", event)
	// the assetID is an internal prefixed assembly ID
	state, err := getUnmarshalledState(stub, "handleAssemblyFlightEvent", assetID)
	if err != nil {
		return nil, err
	}

	// count cycles for the life of the assembly
	cycles, found := getObjectAsNumber(state, "cycles")
	if found {
		cycles++
	} else {
		cycles = 1
	}
	state, ok = putObject(state, "cycles", cycles)
	if !ok {
		err := errors.New("handleAssemblyFlightEvent: putObject failed for cycles")
		log.Error(err)
		return nil, err
	}

	adjustedCycles, found := getObjectAsNumber(state, "adjustedCycles")
	if found {
		adjustedCycles++
	} else {
		adjustedCycles = 1
	}

	state, ok = putObject(state, "adjustedCycles", adjustedCycles)
	if !ok {
		err := errors.New("handleAssemblyFlightEvent: putObject failed for adjustedCycles")
		log.Error(err)
		return nil, err
	}

	acc, found := getObjectAsNumber(state, "aCheckCounter")
	if found {
		acc++
	} else {
		acc = 1
	}

	state, ok = putObject(state, "aCheckCounter", acc)
	if !ok {
		err := fmt.Errorf("handleAssemblyFlightEvent: aCheckCounter property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	acca, found := getObjectAsNumber(state, "aCheckCounterAdjusted")
	if found {
		acca++
	} else {
		acca = 1
	}

	state, ok = putObject(state, "aCheckCounterAdjusted", acca)
	if !ok {
		err := fmt.Errorf("handleAssemblyFlightEvent: aCheckCounterAdjusted property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	bcc, found := getObjectAsNumber(state, "bCheckCounter")
	if found {
		bcc++
	} else {
		bcc = 1
	}

	state, ok = putObject(state, "bCheckCounter", bcc)
	if !ok {
		err := fmt.Errorf("handleAssemblyFlightEvent: bCheckCounter property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	bcca, found := getObjectAsNumber(state, "bCheckCounterAdjusted")
	if found {
		bcca++
	} else {
		bcca = 1
	}

	state, ok = putObject(state, "bCheckCounterAdjusted", bcca)
	if !ok {
		err := fmt.Errorf("handleAssemblyFlightEvent: bCheckCounterAdjusted property could not be written into state: %+v", state)
		log.Error(err)
		return nil, err
	}

	state, err = addTXNTimestampToState(stub, "handleAssemblyFlightEvent", state)
	if err != nil {
		return nil, err
	}
	state = addLastEventToState(stub, "handleAssemblyFlightEvent", event, state, "")

	state, err = handleAlertsAndRules(stub, "handleAssemblyFlightEvent", "flight", assetID, event, state)
	if err != nil {
		return nil, err
	}

	err = putMarshalledState(stub, "handleAssemblyFlightEvent", "flight", assetID, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
