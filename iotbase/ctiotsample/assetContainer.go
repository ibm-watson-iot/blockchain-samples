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
//                   Tracks containers in trade lane fashion

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
)

// ContainerClass acts as the class of all containers
var ContainerClass = as.AssetClass{"container", "CON", "container.barcode"}

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

func (t *SimpleChaincode) createAssetContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return ContainerClass.CreateAsset(stub, args, "createAssetContainer", []as.QPropNV{})
}

func (t *SimpleChaincode) updateAssetContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return ContainerClass.UpdateAsset(stub, args, "updateAssetContainer", []as.QPropNV{})
}

// func (t *SimpleChaincode) deleteAssetContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	return ContainerClass.DeleteAsset(stub, args)
// }

// func (t *SimpleChaincode) deleteAllAssetsContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	return as.DeleteAllAssets(stub)
// }

// func (t *SimpleChaincode) deletePropertiesFromAssetContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	return ContainerClass.DeletePropertiesFromAsset(stub, args, []as.QPropNV{})
// }

func (t *SimpleChaincode) readAssetContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return ContainerClass.ReadAsset(stub, args)
}

func (t *SimpleChaincode) readAllAssetsContainer(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return ContainerClass.ReadAllAssets(stub, args)
}

// func (t *SimpleChaincode) readAssetIOTHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
// 	return as.ReadAssetHistory(stub, args, "iot", "readAssetIOTHistory")
// }
