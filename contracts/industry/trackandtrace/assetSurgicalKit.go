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

// SurgicalKitClass acts as the class of all SurgicalKits
var SurgicalKitClass = iot.AssetClass{
	Name:        "SurgicalKit",
	Prefix:      "SKT",
	AssetIDPath: "surgicalkit.skitID",
}

func newSurgicalKit() iot.Asset {
	return iot.Asset{
		Class:        SurgicalKitClass,
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

var createAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.CreateAsset(stub, args, "createAssetSurgicalKit", []iot.QPropNV{})
}

var replaceAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.ReplaceAsset(stub, args, "replaceAssetSurgicalKit", []iot.QPropNV{})
}

var updateAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.UpdateAsset(stub, args, "updateAssetSurgicalKit", []iot.QPropNV{})
}

var deleteAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.DeleteAsset(stub, args)
}

var deleteAssetStateHistorySurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.DeleteAssetStateHistory(stub, args)
}

var deleteAllAssetsSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.DeleteAllAssets(stub, args)
}

var deletePropertiesFromAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.DeletePropertiesFromAsset(stub, args, "deletePropertiesFromAssetSurgicalKit", []iot.QPropNV{})
}

var readAssetSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.ReadAsset(stub, args)
}

var readAllAssetsSurgicalKit iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.ReadAllAssets(stub, args)
}

var readAssetStateHistorySurgicalKit = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return SurgicalKitClass.ReadAssetStateHistory(stub, args)
}

var excessForceAlert iot.AlertName = "EXCESSFORCE"
var excessForceRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, SurgicalKit *iot.Asset) error {
	force, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.sensors.maxgforce")
	if found {
		if force > 2 {
			iot.RaiseAlert(SurgicalKit, excessForceAlert)
		} else {
			iot.ClearAlert(SurgicalKit, excessForceAlert)
		}
	}
	return nil
}

var excessTiltAlert iot.AlertName = "EXCESSTILT"
var excessTiltRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, SurgicalKit *iot.Asset) error {
	tilt, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.sensors.maxtilt")
	if found {
		if tilt > 90 || tilt < -90 {
			iot.RaiseAlert(SurgicalKit, excessTiltAlert)
		} else {
			iot.ClearAlert(SurgicalKit, excessTiltAlert)
		}
	}
	return nil
}

var outOfAreaAlert iot.AlertName = "OUTOFAREA"
var outOfAreaRule iot.RuleFunc = func(stub shim.ChaincodeStubInterface, SurgicalKit *iot.Asset) error {
	status, found := iot.GetObjectAsString(SurgicalKit.State, "surgicalkit.status")
	if !found || status != "hospital" {
		return nil
	}
	lat, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.sensors.endlocation.latitude")
	if !found {
		return nil
	}
	long, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.sensors.endlocation.longitude")
	if !found {
		return nil
	}
	flat, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.hospital.fence.center.latitude")
	if !found {
		return nil
	}
	flong, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.hospital.fence.center.longitude")
	if !found {
		return nil
	}
	radius, found := iot.GetObjectAsNumber(SurgicalKit.State, "surgicalkit.hospital.fence.radius")
	if !found {
		return nil
	}
	distance := iot.Distance(lat, long, flat, flong)
	if distance > radius {
		iot.RaiseAlert(SurgicalKit, outOfAreaAlert)
	} else {
		iot.ClearAlert(SurgicalKit, outOfAreaAlert)
	}
	_ = iot.PutObject(SurgicalKit.State, "distanceFromFenceCenter", distance)
	return nil
}

func init() {
	iot.AddRule("Excess Force Alert", SurgicalKitClass, []iot.AlertName{excessForceAlert}, excessForceRule)
	iot.AddRule("Excess Tilt Alert", SurgicalKitClass, []iot.AlertName{excessTiltAlert}, excessTiltRule)
	iot.AddRule("Out Of Area Alert", SurgicalKitClass, []iot.AlertName{outOfAreaAlert}, outOfAreaRule)

	iot.AddRoute("createAssetSurgicalKit", "invoke", SurgicalKitClass, createAssetSurgicalKit)
	iot.AddRoute("replaceAssetSurgicalKit", "invoke", SurgicalKitClass, replaceAssetSurgicalKit)
	iot.AddRoute("updateAssetSurgicalKit", "invoke", SurgicalKitClass, updateAssetSurgicalKit)
	iot.AddRoute("deleteAssetSurgicalKit", "invoke", SurgicalKitClass, deleteAssetSurgicalKit)
	iot.AddRoute("deleteAssetStateHistorySurgicalKit", "invoke", SurgicalKitClass, deleteAssetStateHistorySurgicalKit)
	iot.AddRoute("deleteAllAssetsSurgicalKit", "invoke", SurgicalKitClass, deleteAllAssetsSurgicalKit)
	iot.AddRoute("deletePropertiesFromAssetSurgicalKit", "invoke", SurgicalKitClass, deletePropertiesFromAssetSurgicalKit)
	iot.AddRoute("readAssetSurgicalKit", "query", SurgicalKitClass, readAssetSurgicalKit)
	iot.AddRoute("readAssetStateHistorySurgicalKit", "query", SurgicalKitClass, readAssetStateHistorySurgicalKit)
	iot.AddRoute("readAllAssetsSurgicalKit", "query", SurgicalKitClass, readAllAssetsSurgicalKit)
}
