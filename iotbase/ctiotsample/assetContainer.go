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

// v0.1 KL -- new IOT sample with Trade Lane properties and behaviors

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
)

// ContainerClass acts as the class of all containers
var ContainerClass = as.AssetClass{
	Name:        "container",
	Prefix:      "CON",
	AssetIDPath: "container.barcode",
}

func newContainer() as.Asset {
	return as.Asset{
		Class:     ContainerClass,
		AssetKey:  "",
		State:     nil,
		EventIn:   nil,
		TXNID:     "",
		TXNTS:     nil,
		EventOut:  nil,
		Alerts:    nil,
		Compliant: true,
	}
}

var createAssetContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.CreateAsset(stub, args, "createAssetContainer", []as.QPropNV{})
}

var updateAssetContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.UpdateAsset(stub, args, "updateAssetContainer", []as.QPropNV{})
}

var deleteAssetContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeleteAsset(stub, args)
}

var deleteAllAssetsContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetContainer", []as.QPropNV{})
}

var readAssetContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.ReadAsset(stub, args)
}

var readAllAssetsContainer as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.ReadAllAssets(stub, args)
}

// func (t *SimpleChaincode) readAssetIOTHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
// 	return as.ReadAssetHistory(stub, args, "iot", "readAssetIOTHistory")
// }

func init() {
	as.AddRoute("createAssetContainer", "invoke", ContainerClass, createAssetContainer)
	as.AddRoute("updateAssetContainer", "invoke", ContainerClass, updateAssetContainer)
	as.AddRoute("deleteAssetContainer", "invoke", ContainerClass, deleteAssetContainer)
	as.AddRoute("deleteAllAssetsContainer", "invoke", ContainerClass, deleteAllAssetsContainer)
	as.AddRoute("deletePropertiesFromAssetContainer", "invoke", ContainerClass, deletePropertiesFromAssetContainer)
	as.AddRoute("readAssetContainer", "query", ContainerClass, readAssetContainer)
	as.AddRoute("readAllAssetsContainer", "query", ContainerClass, readAllAssetsContainer)
}
