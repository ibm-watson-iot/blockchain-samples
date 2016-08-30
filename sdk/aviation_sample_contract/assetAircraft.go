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

func (t *SimpleChaincode) readAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAsset(stub, args, "aircraft", "readAssetAircraft")
}

func (t *SimpleChaincode) readAllAssetsAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAllAssets(stub, args, "aircraft", "readAllAssetsAircraft")
}

func (t *SimpleChaincode) readAssetAircraftHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "aircraft", "readAssetAircraftHistory")
}
