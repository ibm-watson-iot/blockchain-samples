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

// v1 KL 12 Aug 2016 Implement maintenance event

// Maintenance Event
// This event targets assembly assets.
// For the assembly, the following actions result in the status
//  Action                 status
//  commission             inventory
//  Installation           aircraft
//  Uninstallation         inventory
//  MaintenanceStart       maintenance
//  MaintenanceComplete    inventory
//  Scrap                  scrapped
//
// An assembly:
//   - can only be installed on an aircraft from inventory
//   - can only go to maintenance from inventory
//   - providing an aircraft serial number for maintenanceStarted is an error
//
// NOTE: This is not an enforced finite state machine. It is possible to go from
// any status to any other status by executing the appropriate action. Some combinations
// make no sense and will be errored.

package main

import (
	"fmt"
	//"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func eventMaintenance(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	event, err := getUnmarshalledArgument(stub, "eventMaintenance", args)
	if err != nil {
		return nil, err
	}
	_, err = handleAssemblyMaintenanceEvent(stub, event)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func handleAssemblyMaintenanceEvent(stub *shim.ChaincodeStub, event interface{}) (interface{}, error) {
	assetID, err := getEventAssetID("handleAssemblyMaintenanceEvent", "maintenance", "maintenance.assembly", event)
	if err != nil {
		return nil, err
	}

	state, err := getUnmarshalledState(stub, "handleAssemblyMaintenanceEvent", assetID)
	if err != nil {
		return nil, err
	}

	state, err = processMaintenanceAction(stub, state, event, assetID)
	if err != nil {
		return nil, err
	}

	// state will be of type interface{} for use with crudUtils
	state, err = addTXNTimestampToState(stub, "handleAssemblyMaintenanceEvent", state)
	if err != nil {
		return nil, err
	}

	state = addLastEventToState(stub, "handleAssemblyMaintenanceEvent", event, state, "")

	// no rules at this time, so commenting this out
	//state, err = handleAlertsAndRules(stub, "handleAssemblyMaintenanceEvent", "maintenance", assetID, event, state)
	//if err != nil { return nil, err }

	err = putMarshalledState(stub, "handleAssemblyMaintenanceEvent", "maintenance", assetID, state)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func processMaintenanceAction(stub *shim.ChaincodeStub, state interface{}, event interface{}, assetID string) (interface{}, error) {
	var ok bool

	indexes, err := getAircraftAssemblyIndexesFromLedger(stub)
	if err != nil {
		return nil, err
	}

	action, found := getObjectAsString(event, "maintenance.action")
	if !found {
		err := fmt.Errorf("processMaintenanceAction: action property missing from event: %+v", event)
		log.Error(err)
		return nil, err
	}

	eventAssemblyID, found := getObjectAsString(event, "maintenance.assembly")
	if !found {
		// no assembly in event
		err = fmt.Errorf("processMaintenanceAction: assembly id missing from maintenance event but is required")
		log.Error(err)
		return nil, err
	}
	eventAssemblyID, err = assetIDToInternal("maintenance", eventAssemblyID)
	if err != nil {
		return nil, err
	}

	eventAircraftID, found := getObjectAsString(event, "maintenance.aircraft")
	if !found {
		// no aircraft in event, check to see if it was mandatory
		if contains([]string{"install", "uninstall"}, action) {
			err = fmt.Errorf("processMaintenanceAction: aircraft id missing from maintenance event but is required for action %s", action)
			log.Error(err)
			return nil, err
		}
	} else {
		eventAircraftID, err = assetIDToInternal("aircraft", eventAircraftID)
		if err != nil {
			return nil, err
		}
	}

	log.Info(fmt.Sprintf("\n\nProcess Maintenance Action: \n\nSTATE: %+v \n\nEVENT: %+v\n\n EVENT ASSEM: %s   EVENT AIRCRAFT: %s   ACTION: %s\n\n", state, event, eventAssemblyID, eventAircraftID, action))

	switch action {
	case "commission":
		// commission always goes into inventory, error if not "new"
		err := validateStatus(state, []string{"new"})
		if err != nil {
			return nil, err
		}
		state, ok = putObject(state, "status", "inventory")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to commmission as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	case "install":
		err := validateStatus(state, []string{"inventory"})
		if err != nil {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to install to aircraft %s: %s", eventAssemblyID, eventAircraftID, err.Error())
			log.Error(err)
			return nil, err
		}
		currAircraft, found := indexes.isAssemblyOnAnyAircraft(eventAssemblyID)
		if found {
			err := fmt.Errorf("processMaintenanceAction: assembly %s cannot be installed on aircraft %s as it is already on aircraft %s", eventAssemblyID, eventAircraftID, currAircraft)
			log.Error(err)
			return nil, err
		}
		// good to install
		err = indexes.addAssemblyToAircraft(eventAssemblyID, eventAircraftID)
		if err != nil {
			return nil, err
		}
		state, ok = putObject(state, "status", "aircraft")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to install as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	case "uninstall":
		err := validateStatus(state, []string{"aircraft"})
		if err != nil {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to uninstall from aircraft %s: %s", eventAssemblyID, eventAircraftID, err.Error())
			log.Error(err)
			return nil, err
		}
		if !indexes.isAssemblyOnThisAircraft(eventAssemblyID, eventAircraftID) {
			err := fmt.Errorf("processMaintenanceAction: assembly %s cannot be uninstalled as it is not on aircraft: %s", eventAssemblyID, eventAircraftID)
			log.Error(err)
			return nil, err
		}
		// good to uninstall
		err = indexes.removeAssemblyFromAircraft(eventAssemblyID, eventAircraftID)
		if err != nil {
			return nil, err
		}
		state, ok = putObject(state, "status", "inventory")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to uninstall as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	case "startMaintenance":
		err := validateStatus(state, []string{"inventory"})
		if err != nil {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to start maintenence: %s", eventAssemblyID, err.Error())
			log.Error(err)
			return nil, err
		}
		state, ok = putObject(state, "status", "maintenance")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to startMaintenance as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	case "endMaintenance":
		err := validateStatus(state, []string{"maintenance"})
		if err != nil {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to end maintenence: %s", eventAssemblyID, err.Error())
			log.Error(err)
			return nil, err
		}
		state, ok = putObject(state, "status", "inventory")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to endMaintenance as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	case "scrap":
		// note that "" is included so that an assembly can be scrapped before it is commissioned
		err := validateStatus(state, []string{"inventory", "maintenance", "new"})
		if err != nil {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to start maintenence: %s", eventAssemblyID, err.Error())
			log.Error(err)
			return nil, err
		}
		state, ok = putObject(state, "status", "scrapped")
		if !ok {
			err := fmt.Errorf("processMaintenanceAction: assembly %s failed to scrap as putObject failed", eventAssemblyID)
			log.Error(err)
			return nil, err
		}
	default:
		err := fmt.Errorf("processMaintenanceAction: action property unknown: %s", action)
		log.Error(err)
		return nil, err
	}

	err = putAircraftAssemblyIndexesToLedger(stub, indexes)
	if err != nil {
		return nil, err
	}

	return state, nil
}

func validateStatus(state interface{}, statuses []string) error {
	status, ok := getObjectAsString(state, "status")
	if !ok {
		// no status means new, should be checking an empty array
		if len(statuses) == 0 {
			return nil
		}
		err := fmt.Errorf("validateStatus: expecting one of %+v statuses but found no status property", statuses)
		log.Error(err)
		return err
	}
	if contains(statuses, status) {
		return nil
	}
	err := fmt.Errorf("validateStatus: status %s not in expected statuses %+v", status, statuses)
	log.Error(err)
	return err
}
