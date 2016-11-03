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
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return createAsset(stub, args, "aircraft", "createAssetAircraft", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) updateAssetAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return updateAsset(stub, args, "aircraft", "updateAssetAircraft", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) deleteAssetAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return deleteAsset(stub, args, "aircraft", "deleteAssetAircraft")
}

func (t *SimpleChaincode) deleteAllAssetsAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return deleteAllAssets(stub, args, "aircraft", "deleteAllAssetsAircraft")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return deletePropertiesFromAsset(stub, args, "aircraft", "deletePropertiesFromAssetAircraft", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) readAssetAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return readAsset(stub, args, "aircraft", "readAssetAircraft")
}

func (t *SimpleChaincode) readAllAssetsAircraft(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return readAllAssets(stub, args, "aircraft", "readAllAssetsAircraft")
}

func (t *SimpleChaincode) readAssetAircraftComplete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var result = struct {
		Aircraft   interface{}   `json:"aircraft"`
		Assemblies []interface{} `json:"assemblies"`
	}{}
	argsMap, err := getUnmarshalledArgument(stub, "readAssetAircraftComplete", args)
	if err != nil {
		return nil, err
	}
	assetIDInternal, err := validateAssetID("readAssetAircraftComplete", "aircraft", argsMap)
	if err != nil {
		return nil, err
	}
	result.Aircraft, err = readAssetUnmarshalled(stub, assetIDInternal, "aircraft", "readAssetAircraftComplete:aircraft")
	if err != nil {
		return nil, err
	}
	assemblies, ok := getObjectAsStringArray(result.Aircraft, "assemblies")
	if !ok {
		// ignore it, since it could be that there simply aren't any assemblies
		assemblies = []string{}
	}
	for _, asm := range assemblies {
		asmInternal, err := assetIDToInternal("assembly", asm)
		if err != nil {
			return nil, err
		}
		assembly, err := readAssetUnmarshalled(stub, asmInternal, "assembly", "readAssetAircraftComplete:asssembly")
		if err != nil {
			return nil, err
		}
		result.Assemblies = append(result.Assemblies, assembly)
	}
	aircraftsbytes, err := json.Marshal(&result)
	if err != nil {
		err = errors.New("readAssetAircraftComplete failed to marshall aircraftsbytes: " + err.Error())
		log.Error(err)
		return nil, err
	}
	return aircraftsbytes, nil
}

func (t *SimpleChaincode) readAssetAircraftHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "aircraft", "readAssetAircraftHistory")
}
