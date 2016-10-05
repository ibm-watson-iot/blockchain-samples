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

// v1 KL 07 Aug 2016 Add event handling in a separate Assembly module

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return createAsset(stub, args, "assembly", "createAssetAssembly", []QualifiedPropertyNameValue{{"status", "new"}})
}

func (t *SimpleChaincode) updateAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return updateAsset(stub, args, "assembly", "updateAssetAssembly", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) deleteAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAsset(stub, args, "assembly", "deleteAssetAssembly")
}

func (t *SimpleChaincode) deleteAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAllAssets(stub, args, "assembly", "deleteAllAssetsAssembly")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deletePropertiesFromAsset(stub, args, "assembly", "deletePropertiesFromAssetAssembly", []QualifiedPropertyNameValue{})
}

func injectAircraft(stub *shim.ChaincodeStub, assembly interface{}) (interface{}, error) {
	var err error
	assem, ok := asMap(assembly)
	if !ok {
		err = fmt.Errorf("injectAircraft failed to convert assembly %+v to map: %s", assembly, err)
		log.Error(err)
		return nil, err
	}
	assemblyID, ok := getObjectAsString(assem, "common.assetID")
	if !ok {
		err = fmt.Errorf("injectAircraft failed to find assetID in assembly %+v", assembly)
		log.Error(err)
		return nil, err
	}
	assemblyID, err = assetIDToInternal("assembly", assemblyID)
	if err != nil {
		err = errors.New("injectAircraft failed to convert assemblyID to internal format: " + err.Error())
		log.Error(err)
		return nil, err
	}
	indexes, err := getAircraftAssemblyIndexesFromLedger(stub)
	if err != nil {
		err = errors.New("injectAircraft failed to get indexes from ledger: " + err.Error())
		log.Error(err)
		return nil, err
	}
	aircraftID, ok := indexes.AssemblyToAircraft[assemblyID]
	if !ok {
		// assembly has no associated aircraft, this is ok, return original assembly
		return assembly, nil
	}
	aircraftID, err = assetIDToExternal(aircraftID)
	if err != nil {
		err = fmt.Errorf("injectAircraft failed to convert aircraftID to external %s: %s", aircraftID, err)
		log.Error(err)
		return nil, err
	}
	assembytes, ok := putObject(assem, "aircraft", aircraftID)
	if !ok {
		err = fmt.Errorf("injectAircraft failed to put aircraft into assembly %s: %s", aircraftID, err)
		log.Error(err)
		return nil, err
	}
	return assembytes, nil
}

func (t *SimpleChaincode) readAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var assembly interface{}
	aircraftstring, err := readAsset(stub, args, "assembly", "readAssetAssembly")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(aircraftstring), &assembly)
	if err != nil {
		return nil, err
	}
	assembly, err = injectAircraft(stub, assembly)
	if err != nil {
		return nil, err
	}
	assembytes, err := json.Marshal(assembly)
	if err != nil {
		err = errors.New("readAssetAssembly failed to marshall assembly: " + err.Error())
		log.Error(err)
		return nil, err
	}
	return assembytes, nil
}

func (t *SimpleChaincode) readAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	assemblies, err := readAllAssetsUnmarshalled(stub, args, "assembly", "readAllAssetsAssembly")
	if err != nil {
		return nil, err
	}
	for _, assembly := range assemblies {
		assembly, err = injectAircraft(stub, assembly)
		if err != nil {
			return nil, err
		}
	}
	assembliesbytes, err := json.Marshal(&assemblies)
	if err != nil {
		err = errors.New("readAllAssetsAssembly failed to marshall assembliesbytes: " + err.Error())
		log.Error(err)
		return nil, err
	}
	return assembliesbytes, nil
}

func (t *SimpleChaincode) readAssetAssemblyHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "assembly", "readAssetAssemblyHistory")
}
