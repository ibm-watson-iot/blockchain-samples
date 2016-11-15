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
	iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

// ContainerClass acts as the class of all containers
var ContainerClass = iot.AssetClass{
	Name:        "container",
	Prefix:      "CON",
	AssetIDPath: "container.barcode",
}

func newContainer() iot.Asset {
	return iot.Asset{
		Class:        ContainerClass,
		AssetKey:     "",
		State:        nil,
		EventIn:      nil,
		TXNID:        "",
		TXNTS:        nil,
		EventOut:     nil,
		AlertsActive: nil,
		Compliant:    true,
	}
}

var createAssetContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.CreateAsset(stub, args, "createAssetContainer", []iot.QPropNV{})
}

var updateAssetContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.UpdateAsset(stub, args, "updateAssetContainer", []iot.QPropNV{})
}

var deleteAssetContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistoryContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetContainer", []iot.QPropNV{})
}

var readAssetContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.ReadAsset(stub, args)
}

var readAllAssetsContainer iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.ReadAllAssets(stub, args)
}

var readAssetStateHistoryContainer = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return ContainerClass.ReadAssetStateHistory(stub, args)
}

var overtempAlert iot.AlertName = "OVERTEMP"
var overtempRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, container *iot.Asset) error {
	temp, found := iot.GetObjectAsNumber(container.State, "container.temperature")
	if found {
		if temp > 0 {
			iot.RaiseAlert(container, overtempAlert)
		} else {
			iot.ClearAlert(container, overtempAlert)
		}
	}
	return nil
}

func init() {
	iot.AddRule("Over Temperature Alert", ContainerClass, []iot.AlertName{overtempAlert}, overtempRule)
	iot.AddRoute("createAssetContainer", "invoke", ContainerClass, createAssetContainer)
	iot.AddRoute("updateAssetContainer", "invoke", ContainerClass, updateAssetContainer)
	iot.AddRoute("deleteAssetContainer", "invoke", ContainerClass, deleteAssetContainer)
	iot.AddRoute("deleteAssetStateHistoryContainer", "invoke", ContainerClass, deleteAssetStateHistoryContainer)
	iot.AddRoute("deleteAllAssetsContainer", "invoke", ContainerClass, deleteAllAssetsContainer)
	iot.AddRoute("deletePropertiesFromAssetContainer", "invoke", ContainerClass, deletePropertiesFromAssetContainer)
	iot.AddRoute("readAssetContainer", "query", ContainerClass, readAssetContainer)
	iot.AddRoute("readAssetStateHistoryContainer", "query", ContainerClass, readAssetStateHistoryContainer)
	iot.AddRoute("readAllAssetsContainer", "query", ContainerClass, readAllAssetsContainer)
}
