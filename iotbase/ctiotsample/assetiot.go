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
// V2 KL 03 Nov 2016 Adapted from aviation contract as new iot sample

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
)

func (t *SimpleChaincode) createAssetIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.CreateAsset(stub, args, "iot", "createAssetIOT", []as.QPropNV{})
}

func (t *SimpleChaincode) updateAssetIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.UpdateAsset(stub, args, "iot", "updateAssetIOT", []as.QPropNV{})
}

func (t *SimpleChaincode) deleteAssetIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.DeleteAsset(stub, args, "iot", "deleteAssetIOT")
}

func (t *SimpleChaincode) deleteAllAssetsIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.DeleteAllAssets(stub, args, "iot", "deleteAllAssetsIOT")
}

func (t *SimpleChaincode) deletePropertiesFromAssetIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.DeletePropertiesFromAsset(stub, args, "iot", "deletePropertiesFromAssetIOT", []as.QPropNV{})
}

func (t *SimpleChaincode) readAssetIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.ReadAsset(stub, args, "iot", "readAssetIOT")
}

func (t *SimpleChaincode) readAllAssetsIOT(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.ReadAllAssets(stub, args, "iot", "readAllAssetsIOT")
}

func (t *SimpleChaincode) readAssetIOTHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return as.ReadAssetHistory(stub, args, "iot", "readAssetIOTHistory")
}
