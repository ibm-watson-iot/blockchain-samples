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

// v1 KL 07 Aug 2016 Separate crud API for assets

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return createAsset(stub, args, "airline", "createAssetAirline", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) updateAssetAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return updateAsset(stub, args, "airline", "updateAssetAirline", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) deleteAssetAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAsset(stub, args, "airline", "deleteAssetAirline")
}

func (t *SimpleChaincode) deleteAllAssetsAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deleteAllAssets(stub, args, "airline", "deleteAllAssetsAirline")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return deletePropertiesFromAsset(stub, args, "airline", "deletePropertiesFromAssetAirline", []QualifiedPropertyNameValue{})
}

func (t *SimpleChaincode) readAssetAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAsset(stub, args, "airline", "readAssetAirline")
}

func (t *SimpleChaincode) readAllAssetsAirline(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAllAssets(stub, args, "airline", "readAllAssetsAirline")
}

func (t *SimpleChaincode) readAssetAirlineHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return readAssetHistory(stub, args, "airline", "readAssetAirlineHistory")
}
