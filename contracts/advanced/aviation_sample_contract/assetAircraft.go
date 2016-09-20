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

// v1 KL 07 Aug 2016 Add event handling in a separate Aircraft module

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	_, err := createAsset(stub, args, "aircraft", "createAssetAircraft")
	// an opportunity to augment the state
	return nil, err
}

func (t *SimpleChaincode) updateAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	_, err := updateAsset(stub, args, "aircraft", "updateAssetAircraft")
	// an opportunity to augment the state
	return nil, err
}

func (t *SimpleChaincode) deleteAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAsset(stub, args, "aircraft", "deleteAssetAircraft")
}

func (t *SimpleChaincode) deleteAllAssetsAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAllAssets(stub, args, "aircraft", "deleteAllAssetsAircraft")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	_, err := deletePropertiesFromAsset(stub, args, "aircraft", "deletePropertiesFromAssetAircraft")
	// an opportunity to augment the state
	return nil, err
}

func injectAssemblies(stub *shim.ChaincodeStub, aircraft interface{}) (interface{}, error) {
	var err error
	var assemblies []string
	ac := asMap(aircraft)
	if len(ac) == 0 {
		err = fmt.Errorf("injectAssemblies failed to convert aircraft %+v to map: %s", aircraft, err)
		log.Error(err)
		return nil, err
	}
	aircraftID, ok := getObjectAsString(ac, "common.assetID")
	if !ok {
		err = fmt.Errorf("injectAssemblies failed to find assetID in aircraft %+v", aircraft)
		log.Error(err)
		return nil, err
	}
	aircraftID, err = assetIDToInternal("aircraft", aircraftID)
	if err != nil {
		err = errors.New("injectAssemblies failed to convert aircraftID to internal format: " + err.Error())
		log.Error(err)
		return nil, err
	}
	indexes, err := getAircraftAssemblyIndexesFromLedger(stub)
	if err != nil {
		err = errors.New("injectAssemblies failed to get indexes from ledger: " + err.Error())
		log.Error(err)
		return nil, err
	}
	assembliesmap, ok := indexes.getAircraftAssemblies(aircraftID)
	if !ok {
		// aircraft has no assemblies, this is ok, assign empty assemblies list
		assembliesmap = make(map[string]struct{}, 0)
	}
	for assemblyID := range assembliesmap {
		assemblyID, err = assetIDToExternal(assemblyID)
		if err != nil {
			err = fmt.Errorf("injectAssemblies failed to convert assemblyID to external %s: %s", assemblyID, err)
			log.Error(err)
			return nil, err
		}
		assemblies = append(assemblies, assemblyID)
	}
	acbytes, ok := putObject(ac, "assemblies", assemblies)
	if !ok {
		err = fmt.Errorf("injectAssemblies failed to put assemblies into aircraft %s: %s", aircraftID, err)
		log.Error(err)
		return nil, err
	}
	return acbytes, nil
}

func (t *SimpleChaincode) readAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var aircraft interface{}
	aircraftstring, err := readAsset(stub, args, "aircraft", "readAssetAircraft")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(aircraftstring), &aircraft)
	if err != nil {
		return nil, err
	}
	aircraft, err = injectAssemblies(stub, aircraft)
	if err != nil {
		return nil, err
	}
	aircraftbytes, err := json.Marshal(aircraft)
	if err != nil {
		err = errors.New("readAssetAircraft failed to marshall aircraft: " + err.Error())
		log.Error(err)
		return nil, err
	}
	return aircraftbytes, nil
}

func (t *SimpleChaincode) readAllAssetsAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	aircrafts, err := readAllAssetsUnmarshalled(stub, args, "aircraft", "readAllAssetsAircraft")
	if err != nil {
		return nil, err
	}
	for _, aircraft := range aircrafts {
		aircraft, err = injectAssemblies(stub, aircraft)
		if err != nil {
			return nil, err
		}
	}
	aircraftsbytes, err := json.Marshal(&aircrafts)
	if err != nil {
		err = errors.New("readAllAssetsAircraft failed to marshall aircraftsbytes: " + err.Error())
		log.Error(err)
		return nil, err
	}
	return aircraftsbytes, nil
}

func (t *SimpleChaincode) readAssetAircraftHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "aircraft", "readAssetAircraftHistory")
}
