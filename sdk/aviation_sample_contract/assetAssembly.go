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
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	state, err := createAsset(stub, args, "assembly", "createAssetAssembly")
	if err != nil {
		return nil, err
	}
	// an opportunity to augment the state
	state, ok := putObject(state, "status", "new")
	if !ok {
		err = errors.New("createAssetAssembly failed to put object status into state")
		log.Error(err)
		return nil, err
	}
	assetID, ok := getObjectAsString(state, "common.assetID")
	if !ok {
		err = errors.New("createAssetAssembly failed to get assetID state")
		log.Error(err)
		return nil, err
	}
	assetID, err = assetIDToInternal("assembly", assetID)
	if err != nil {
		return nil, err
	}
	err = putMarshalledState(stub, "createAssetAssembly", "assembly", assetID, state)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *SimpleChaincode) updateAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	_, err := updateAsset(stub, args, "assembly", "updateAssetAssembly")
	// an opportunity to augment the state
	return nil, err
}

func (t *SimpleChaincode) deleteAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAsset(stub, args, "assembly", "deleteAssetAssembly")
}

func (t *SimpleChaincode) deleteAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAllAssets(stub, args, "assembly", "deleteAllAssetsAssembly")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	_, err := deletePropertiesFromAsset(stub, args, "assembly", "deletePropertiesFromAssetAssembly")
	// an opportunity to augment the state
	return nil, err
}

func (t *SimpleChaincode) readAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAsset(stub, args, "assembly", "readAssetAssembly")
}

func (t *SimpleChaincode) readAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAllAssets(stub, args, "assembly", "readAllAssetsAssembly")
}

func (t *SimpleChaincode) readAssetAssemblyHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "assembly", "readAssetAssemblyHistory")
}
