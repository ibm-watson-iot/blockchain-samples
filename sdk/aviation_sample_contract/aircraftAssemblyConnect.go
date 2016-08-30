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

// v1 KL 13 Aug 2016 Add the indexing module for aircraft and assemblies

package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// This module provides relational processing of the aircraft and assembly connections.
// It is a simple module right now, but will grow and eventually be replaced by a real
// database when that arrives in Hyperledger.

// AircraftAssemblyIndexes provides a simple structure around which to build indexing
// operations. These should be internal (world state) IDs with the asset prefixes so that
// the assets can be read directly.
type AircraftAssemblyIndexes struct {
	// golang idiom for map of sets
	AircraftToAssemblies map[string]map[string]struct{} `json:"aircraft2assemblies"`
	// simple 1:1 map
	AssemblyToAircraft map[string]string `json:"assembly2aircraft"`
}

const AIRCRAFTASSEMBLYINDEXESKEY = "AircraftAssemblyIndexes"

func getAircraftAssemblyIndexesFromLedger(stub *shim.ChaincodeStub) (AircraftAssemblyIndexes, error) {
	// indexes must exist as they were created during Init
	var indexes = AircraftAssemblyIndexes{}
	indexBytes, err := stub.GetState(AIRCRAFTASSEMBLYINDEXESKEY)
	if err != nil {
		err := fmt.Errorf("getIndexesFromLedger cannot get the aircraft assembly index structure from the ledger: %s", err.Error())
		log.Error(err)
		return AircraftAssemblyIndexes{}, err
	}
	err = json.Unmarshal(indexBytes, &indexes)
	if err != nil {
		err := fmt.Errorf("getIndexesFromLedger cannot unmarshal the aircraft assembly index structure: %s", err.Error())
		log.Error(err)
		return AircraftAssemblyIndexes{}, err
	}
	return indexes, nil
}

func putAircraftAssemblyIndexesToLedger(stub *shim.ChaincodeStub, indexes AircraftAssemblyIndexes) error {
	indexBytes, err := json.Marshal(&indexes)
	if err != nil {
		err := fmt.Errorf("putIndexesToLedger cannot marshall the aircraft assembly index structure: %s", err.Error())
		log.Error(err)
		return err
	}
	err = stub.PutState(AIRCRAFTASSEMBLYINDEXESKEY, indexBytes)
	if err != nil {
		err := fmt.Errorf("putIndexesToLedger cannot put the aircraft assembly index structure to the ledger: %s", err.Error())
		log.Error(err)
		return err
	}
	return nil
}

func makeAircraftAssemblyIndexes() AircraftAssemblyIndexes {
	return AircraftAssemblyIndexes{
		make(map[string]map[string]struct{}),
		make(map[string]string),
	}
}

func initAircraftAssemblyIndexes(stub *shim.ChaincodeStub) error {
	indexbytes, err := stub.GetState(AIRCRAFTASSEMBLYINDEXESKEY)
	if err == nil && len(indexbytes) > 0 {
		// already initialized, this is a reboot etc
		return nil
	}
	indexes := makeAircraftAssemblyIndexes()
	err = putAircraftAssemblyIndexesToLedger(stub, indexes)
	if err != nil {
		err = fmt.Errorf("initAircraftAssemblyIndexes cannot put the index structure to world state: %s", err.Error())
		log.Error(err)
		return err
	}
	return nil
}

func (indexes AircraftAssemblyIndexes) isAssemblyOnThisAircraft(assemblyID string, aircraftID string) bool {
	aircraft, found := indexes.AssemblyToAircraft[assemblyID]
	return found && aircraft == aircraftID
}

func (indexes AircraftAssemblyIndexes) isAssemblyOnAnyAircraft(assemblyID string) (string, bool) {
	aircraft, found := indexes.AssemblyToAircraft[assemblyID]
	if found {
		return aircraft, true
	}
	return "", false
}

func (indexes AircraftAssemblyIndexes) getAircraftAssemblies(aircraftID string) (map[string]struct{}, bool) {
	assemblies, found := indexes.AircraftToAssemblies[aircraftID]
	return assemblies, found
}

func (indexes AircraftAssemblyIndexes) addAssemblyToAircraft(assemblyID string, aircraftID string) error {
	currAircraft, found := indexes.isAssemblyOnAnyAircraft(assemblyID)
	if !found {
		aircraftAssemblies := indexes.AircraftToAssemblies[aircraftID]
		if aircraftAssemblies == nil {
			indexes.AircraftToAssemblies[aircraftID] = make(map[string]struct{})
		}
		indexes.AircraftToAssemblies[aircraftID][assemblyID] = struct{}{}
		indexes.AssemblyToAircraft[assemblyID] = aircraftID
		return nil
	}
	err := fmt.Errorf("addAssemblyToAircraft: cannot add assembly %s to aircraft %s as assembly is already on aircraft %s", assemblyID, aircraftID, currAircraft)
	log.Error(err)
	return err
}

func (indexes AircraftAssemblyIndexes) removeAssemblyFromAircraft(assemblyID string, aircraftID string) error {
	if indexes.isAssemblyOnThisAircraft(assemblyID, aircraftID) {
		delete(indexes.AircraftToAssemblies[aircraftID], assemblyID)
		delete(indexes.AssemblyToAircraft, assemblyID)
		return nil
	}
	err := fmt.Errorf("removeAssemblyFromAircraft: cannot remove assembly %s from aircraft %s as assembly is not on this aircraft", assemblyID, aircraftID)
	log.Error(err)
	return err
}
