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
Howard McKinney- Initial Contribution
Sumabala Nair - modified the contract for Building Sensors use case.
*/

// Adapted from the  IoT Blockchain Demo Smart Contract at
// https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_hyperledger
// v 1 : Building sensor contract  - adaptation  -June - July 2016

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"strings"
	"time"
)

// For now, We will not use go generate here : go:generate go run scripts/generate_go_schema.go
// This is because there are different assets of various structures coming in. The only
// certainity is that the assets should have an asset id and asset type. These will be extracted
// and the asset data stored and manipulated as such
// Alerts have been temporarily written in a manner that they will run only for applicable types -
// essentially based on whether the observed values are in the incoming data stream
// This needs to be modified as well.
// Major revamp required to the approach based on the use case.

//***************************************************
//***************************************************
//* CONTRACT initialization and runtime engine
//***************************************************
//***************************************************

// ************************************
// definitions
// ************************************

// SimpleChaincode is the receiver for all shim API
type SimpleChaincode struct {
}

// ASSETID is the JSON tag for the assetID
const ASSETID string = "assetID"

// ASSETTYPE is the JSON tag for the asset type
const ASSETTYPE string = "assettype"

// ASSETNAME Asset description from which type is inferred
const ASSETNAME string = "name"

// TIMESTAMP is the JSON tag for timestamps, devices must use this tag to be compatible!
const TIMESTAMP string = "timestamp"

// ArgsMap is a generic map[string]interface{} to be used as a receiver
type ArgsMap map[string]interface{}

var log = NewContractLogger(DEFAULTNICKNAME, DEFAULTLOGGINGLEVEL)

// ************************************
// start the message pumps
// ************************************
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		log.Infof("ERROR starting Simple Chaincode: %s", err)
	}
}

// Init is called in deploy mode when contract is initialized
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var stateArg ContractState
	var err error

	log.Info("Entering INIT")

	if len(args) != 1 {
		err = errors.New("init expects one argument, a JSON string with  mandatory version and optional nickname")
		log.Critical(err)
		return nil, err
	}

	err = json.Unmarshal([]byte(args[0]), &stateArg)
	if err != nil {
		err = fmt.Errorf("Version argument unmarshal failed: %s", err)
		log.Critical(err)
		return nil, err
	}

	if stateArg.Nickname == "" {
		stateArg.Nickname = DEFAULTNICKNAME
	}

	(*log).setModule(stateArg.Nickname)

	err = initializeContractState(stub, stateArg.Version, stateArg.Nickname)
	if err != nil {
		return nil, err
	}

	log.Info("Contract initialized")
	return nil, nil
}

// Invoke is called in invoke mode to delegate state changing function messages
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "createAsset" {
		return t.createAsset(stub, args)
	} else if function == "updateAsset" {
		return t.updateAsset(stub, args)
	} else if function == "deleteAsset" {
		return t.deleteAsset(stub, args)
	} else if function == "deleteAllAssets" {
		return t.deleteAllAssets(stub, args)
	} else if function == "deletePropertiesFromAsset" {
		return t.deletePropertiesFromAsset(stub, args)
	} else if function == "setLoggingLevel" {
		return nil, t.setLoggingLevel(stub, args)
	} else if function == "setCreateOnUpdate" {
		return nil, t.setCreateOnUpdate(stub, args)
	}
	err := fmt.Errorf("Invoke received unknown invocation: %s", function)
	log.Warning(err)
	return nil, err
}

// Query is called in query mode to delegate non-state-changing queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "readAsset" {
		return t.readAsset(stub, args)
	} else if function == "readAllAssets" {
		return t.readAllAssets(stub, args)
	} else if function == "readRecentStates" {
		return readRecentStates(stub)
	} else if function == "readAssetHistory" {
		return t.readAssetHistory(stub, args)
	} else if function == "readAssetSamples" {
		return t.readAssetSamples(stub, args)
	} else if function == "readAssetSchemas" {
		return t.readAssetSchemas(stub, args)
	} else if function == "readContractObjectModel" {
		return t.readContractObjectModel(stub, args)
	} else if function == "readContractState" {
		return t.readContractState(stub, args)
	}
	// To be added
	/*   else if function == "readAllAssetsOfType" {
	return t.readAllAssetsOfType(stub, args)*/
	err := fmt.Errorf("Query received unknown invocation: %s", function)
	log.Warning(err)
	return nil, err
}

//***************************************************
//***************************************************
//* ASSET CRUD INTERFACE
//***************************************************
//***************************************************

// ************************************
// createAsset
// ************************************
func (t *SimpleChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var event interface{}
	var found bool
	var err error
	//var timeIn time.Time

	log.Info("Entering createAsset")

	// allowing 2 args because updateAsset is allowed to redirect when
	// asset does not exist
	if len(args) < 1 || len(args) > 2 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	assetID = ""
	assetType = ""
	assetName = ""
	eventBytes := []byte(args[0])
	log.Debugf("createAsset arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("createAsset failed to unmarshal arg: %s", err)
		return nil, err
	}

	if event == nil {
		err = errors.New("createAsset unmarshal arg created nil event")
		log.Error(err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("createAsset arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("createAsset arg does not include assetID ")
			log.Error(err)
			return nil, err
		}
	}
	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}

	log.Info(assetType)
	sAssetKey := assetID + "_" + assetType
	found = assetIsActive(stub, sAssetKey)
	if found {
		err := fmt.Errorf("createAsset arg asset %s of type %s already exists", assetID, assetType)
		log.Error(err)
		return nil, err
	}

	// For now, timestamp is being sent in from the invocation to the contract
	// Once the BlueMix instance supports GetTxnTimestamp, we will incorporate the
	// changes to the contract

	// run the rules and raise or clear alerts
	alerts := newAlertStatus()
	if argsMap.executeRules(&alerts) {
		// NOT compliant!
		log.Noticef("createAsset assetID %s of type %s is noncompliant", assetID, assetType)
		argsMap["alerts"] = alerts
		delete(argsMap, "incompliance")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(argsMap, "alerts")
		} else {
			argsMap["alerts"] = alerts
		}
		argsMap["incompliance"] = true
	}

	// copy incoming event to outgoing state
	// this contract respects the fact that createAsset can accept a partial state
	// as the moral equivalent of one or more discrete events
	// further: this contract understands that its schema has two discrete objects
	// that are meant to be used to send events: common, and custom
	stateOut := argsMap

	// save the original event
	stateOut["lastEvent"] = make(map[string]interface{})
	stateOut["lastEvent"].(map[string]interface{})["function"] = "createAsset"
	stateOut["lastEvent"].(map[string]interface{})["args"] = args[0]
	if len(args) == 2 {
		// in-band protocol for redirect
		stateOut["lastEvent"].(map[string]interface{})["redirectedFromFunction"] = args[1]
	}

	// marshal to JSON and write
	stateJSON, err := json.Marshal(&stateOut)
	if err != nil {
		err := fmt.Errorf("createAsset state for assetID %s failed to marshal", assetID)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	log.Infof("Putting new asset state %s to ledger", string(stateJSON))
	// The key i 'assetid'_'type'

	err = stub.PutState(sAssetKey, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("createAsset AssetID %s of Type %s PUTSTATE failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}
	log.Infof("createAsset AssetID %s of type %s state successfully written to ledger: %s", assetID, assetType, string(stateJSON))

	// add asset to contract state
	err = addAssetToContractState(stub, sAssetKey)
	if err != nil {
		err := fmt.Errorf("createAsset asset %s of type %s failed to write asset state: %s", assetID, assetType, err)
		log.Critical(err)
		return nil, err
	}

	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("createAsset AssetID %s of type %s push to recentstates failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// save state history
	err = createStateHistory(stub, sAssetKey, string(stateJSON))
	if err != nil {
		err := fmt.Errorf("createAsset asset %s of type %s state history save failed: %s", assetID, sAssetKey, err)
		log.Critical(err)
		return nil, err
	}
	return nil, nil
}

// ************************************
// updateAsset
// ************************************
func (t *SimpleChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var event interface{}
	var ledgerMap ArgsMap
	var ledgerBytes interface{}
	var found bool
	var err error
	//var timeIn time.Time

	log.Info("Entering updateAsset")

	if len(args) != 1 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	assetID = ""
	assetType = ""
	assetName = ""
	eventBytes := []byte(args[0])
	log.Debugf("updateAsset arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("updateAsset failed to unmarshal arg: %s", err)
		return nil, err
	}

	if event == nil {
		err = errors.New("createAsset unmarshal arg created nil event")
		log.Error(err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("updateAsset arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("updateAsset arg does not include assetID")
			log.Error(err)
			return nil, err
		}
	}
	log.Noticef("updateAsset found assetID %s", assetID)

	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}
	log.Noticef("updateAsset found assetID %s of type %s ", assetID, assetType)

	sAssetKey := assetID + "_" + assetType
	found = assetIsActive(stub, sAssetKey)
	if !found {
		// redirect to createAsset with same parameter list
		if canCreateOnUpdate(stub) {
			log.Noticef("updateAsset redirecting asset %s of type %s to createAsset", assetID, assetType)
			var newArgs = []string{args[0], "updateAsset"}
			return t.createAsset(stub, newArgs)
		}
		err = fmt.Errorf("updateAsset asset %s of type %s does not exist", assetID, assetType)
		log.Error(err)
		return nil, err
	}
	// For now, timestamp is being sent in from the invocation to the contract
	// Once the BlueMix instance supports GetTxnTimestamp, we will incorporate the
	// changes to the contract

	// **********************************
	// find the asset state in the ledger
	// **********************************
	log.Infof("updateAsset: retrieving asset %s state from ledger", sAssetKey)
	assetBytes, err := stub.GetState(sAssetKey)
	if err != nil {
		log.Errorf("updateAsset assetID %s of type %s GETSTATE failed: %s", assetID, assetType, err)
		return nil, err
	}

	// unmarshal the existing state from the ledger to theinterface
	err = json.Unmarshal(assetBytes, &ledgerBytes)
	if err != nil {
		log.Errorf("updateAsset assetID %s of type %s unmarshal failed: %s", assetID, assetType, err)
		return nil, err
	}

	// assert the existing state as a map
	ledgerMap, found = ledgerBytes.(map[string]interface{})
	if !found {
		log.Errorf("updateAsset assetID %s of type %s LEDGER state is not a map shape", assetID, assetType)
		return nil, err
	}

	// now add incoming map values to existing state to merge them
	// this contract respects the fact that updateAsset can accept a partial state
	// as the moral equivalent of one or more discrete events
	// further: this contract understands that its schema has two discrete objects
	// that are meant to be used to send events: common, and custom
	// ledger has to have common section
	stateOut := deepMerge(map[string]interface{}(argsMap),
		map[string]interface{}(ledgerMap))
	log.Debugf("updateAsset assetID %s merged state: %s of type %s", assetID, assetType, stateOut)

	// handle compliance section
	alerts := newAlertStatus()
	a, found := stateOut["alerts"] // is there an existing alert state?
	if found {
		// convert to an AlertStatus, which does not work by type assertion
		log.Debugf("updateAsset Found existing alerts state: %s", a)
		// complex types are all untyped interfaces, so require conversion to
		// the structure that is used, but not in the other direction as the
		// type is properly specified
		alerts.alertStatusFromMap(a.(map[string]interface{}))
	}
	// important: rules need access to the entire calculated state
	if ledgerMap.executeRules(&alerts) {
		// true means noncompliant
		log.Noticef("updateAsset assetID %s of type %s is noncompliant", assetID, assetType)
		// update ledger with new state, if all clear then delete
		stateOut["alerts"] = alerts
		delete(stateOut, "incompliance")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(stateOut, "alerts")
		} else {
			stateOut["alerts"] = alerts
		}
		stateOut["incompliance"] = true
	}

	// save the original event
	stateOut["lastEvent"] = make(map[string]interface{})
	stateOut["lastEvent"].(map[string]interface{})["function"] = "updateAsset"
	stateOut["lastEvent"].(map[string]interface{})["args"] = args[0]

	// Write the new state to the ledger
	stateJSON, err := json.Marshal(ledgerMap)
	if err != nil {
		err = fmt.Errorf("updateAsset AssetID %s of type %s marshal failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	err = stub.PutState(sAssetKey, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAsset AssetID %s of type %s PUTSTATE failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}
	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAsset AssetID %s push to recentstates failed: %s", assetID, err)
		log.Error(err)
		return nil, err
	}

	// add history state
	err = updateStateHistory(stub, sAssetKey, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("updateAsset AssetID %s of type %s push to history failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// NOTE: Contract state is not updated by updateAsset

	return nil, nil
}

// ************************************
// deleteAsset
// ************************************
func (t *SimpleChaincode) deleteAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var event interface{}
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("Expecting one JSON state object with an assetID")
		log.Error(err)
		return nil, err
	}

	assetID = ""
	assetType = ""
	assetName = ""
	eventBytes := []byte(args[0])
	log.Debugf("deleteAsset arg: %s", args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Errorf("deleteAsset failed to unmarshal arg: %s", err)
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("deleteAsset arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("deleteAsset arg does not include assetID")
			log.Error(err)
			return nil, err
		}
	}

	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}
	sAssetKey := assetID + "_" + assetType
	found = assetIsActive(stub, sAssetKey)
	if !found {
		err = fmt.Errorf("deleteAsset assetID %s of type  %s does not exist", assetID, assetType)
		log.Error(err)
		return nil, err
	}

	// Delete the key / asset from the ledger
	err = stub.DelState(sAssetKey)
	if err != nil {
		log.Errorf("deleteAsset assetID %s of type %s failed DELSTATE", assetID, assetType)
		return nil, err
	}
	// remove asset from contract state
	err = removeAssetFromContractState(stub, sAssetKey)
	if err != nil {
		err := fmt.Errorf("deleteAsset asset %s of type %s failed to remove asset from contract state: %s", assetID, assetType, err)
		log.Critical(err)
		return nil, err
	}
	// save state history
	err = deleteStateHistory(stub, sAssetKey)
	if err != nil {
		err := fmt.Errorf("deleteAsset asset %s of type %s state history delete failed: %s", assetID, assetType, err)
		log.Critical(err)
		return nil, err
	}
	// push the recent state
	err = removeAssetFromRecentState(stub, sAssetKey)
	if err != nil {
		err := fmt.Errorf("deleteAsset asset %s recent state removal failed: %s", assetID, assetType, err)
		log.Critical(err)
		return nil, err
	}

	return nil, nil
}

// ************************************
// deletePropertiesFromAsset
// ************************************
func (t *SimpleChaincode) deletePropertiesFromAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var event interface{}
	var ledgerMap ArgsMap
	var ledgerBytes interface{}
	var found bool
	var err error
	var alerts AlertStatus

	if len(args) < 1 {
		err = errors.New("Not enough arguments. Expecting one JSON object with mandatory AssetID and property name array")
		log.Error(err)
		return nil, err
	}
	eventBytes := []byte(args[0])

	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		log.Error("deletePropertiesFromAsset failed to unmarshal arg")
		return nil, err
	}

	argsMap, found = event.(map[string]interface{})
	if !found {
		err := errors.New("updateAsset arg is not a map shape")
		log.Error(err)
		return nil, err
	}
	log.Debugf("deletePropertiesFromAsset arg: %+v", argsMap)

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("deletePropertiesFromAsset arg does not include assetID")
			log.Error(err)
			return nil, err
		}
	}
	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}
	sAssetKey := assetID + "_" + assetType

	found = assetIsActive(stub, sAssetKey)
	if !found {
		err = fmt.Errorf("deletePropertiesFromAsset assetID %s of type %s does not exist", assetID, assetType)
		log.Error(err)
		return nil, err
	}

	// is there a list of property names?
	var qprops []interface{}
	qpropsBytes, found := getObject(argsMap, "qualPropsToDelete")
	if found {
		qprops, found = qpropsBytes.([]interface{})
		log.Debugf("deletePropertiesFromAsset qProps: %+v, Found: %+v, Type: %+v", qprops, found, reflect.TypeOf(qprops))
		if !found || len(qprops) < 1 {
			log.Errorf("deletePropertiesFromAsset asset %s of type %s qualPropsToDelete is not an array or is empty", assetID, assetType)
			return nil, err
		}
	} else {
		log.Errorf("deletePropertiesFromAsset asset %s of type %s has no qualPropsToDelete argument", assetID, assetType)
		return nil, err
	}

	// **********************************
	// find the asset state in the ledger
	// **********************************
	log.Infof("deletePropertiesFromAsset: retrieving asset %s of type %s state from ledger", assetID, assetType)
	assetBytes, err := stub.GetState(sAssetKey)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s GETSTATE failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// unmarshal the existing state from the ledger to the interface
	err = json.Unmarshal(assetBytes, &ledgerBytes)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s unmarshal failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// assert the existing state as a map
	ledgerMap, found = ledgerBytes.(map[string]interface{})
	if !found {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s LEDGER state is not a map shape", assetID, assetType)
		log.Error(err)
		return nil, err
	}

	// now remove properties from state, they are qualified by level
OUTERDELETELOOP:
	for p := range qprops {
		prop := qprops[p].(string)
		log.Debugf("deletePropertiesFromAsset AssetID %s of type %s deleting qualified property: %s", assetID, assetType, prop)
		// TODO Ugly, isolate in a function at some point
		if (CASESENSITIVEMODE && strings.HasSuffix(prop, ASSETID)) ||
			(!CASESENSITIVEMODE && strings.HasSuffix(strings.ToLower(prop), strings.ToLower(ASSETID)) ||
				CASESENSITIVEMODE && strings.HasSuffix(prop, ASSETTYPE)) ||
			(!CASESENSITIVEMODE && strings.HasSuffix(strings.ToLower(prop), strings.ToLower(ASSETTYPE))) {
			log.Warningf("deletePropertiesFromAsset AssetID %s of type %s cannot delete protected qualified property: %s or type %s", assetID, assetType, prop)
		} else {
			levels := strings.Split(prop, ".")
			lm := (map[string]interface{})(ledgerMap)
			for l := range levels {
				// lev is the name of a level
				lev := levels[l]
				if l == len(levels)-1 {
					// we're here, delete the actual property name from this level of the map
					levActual, found := findMatchingKey(lm, lev)
					if !found {
						log.Warningf("deletePropertiesFromAsset AssetID %s of type %s property match %s not found", assetID, assetType, lev)
						continue OUTERDELETELOOP
					}
					log.Debugf("deletePropertiesFromAsset AssetID %s of type %s deleting %s", assetID, assetType, prop)
					delete(lm, levActual)
				} else {
					// navigate to the next level object
					log.Debugf("deletePropertiesFromAsset AssetID %s of type %s navigating to level %s", assetID, assetType, lev)
					lmBytes, found := findObjectByKey(lm, lev)
					if found {
						lm, found = lmBytes.(map[string]interface{})
						if !found {
							log.Noticef("deletePropertiesFromAsset AssetID %s of type %s level %s not found in ledger", assetID, assetType, lev)
							continue OUTERDELETELOOP
						}
					}
				}
			}
		}
	}
	log.Debugf("updateAsset AssetID %s final state: %s of type %s ", assetID, assetType, ledgerMap)

	// set timestamp
	// TODO timestamp from the stub - GetTxnTimestamp
	ledgerMap[TIMESTAMP] = time.Now()

	// handle compliance section
	alerts = newAlertStatus()
	a, found := argsMap["alerts"] // is there an existing alert state?
	if found {
		// convert to an AlertStatus, which does not work by type assertion
		log.Debugf("deletePropertiesFromAsset Found existing alerts state: %s", a)
		// complex types are all untyped interfaces, so require conversion to
		// the structure that is used, but not in the other direction as the
		// type is properly specified
		alerts.alertStatusFromMap(a.(map[string]interface{}))
	}
	// important: rules need access to the entire calculated state
	if ledgerMap.executeRules(&alerts) {
		// true means noncompliant
		log.Noticef("deletePropertiesFromAsset assetID %s of type %s is noncompliant", assetID, assetType)
		// update ledger with new state, if all clear then delete
		ledgerMap["alerts"] = alerts
		delete(ledgerMap, "incompliance")
	} else {
		if alerts.AllClear() {
			// all false, no need to appear
			delete(ledgerMap, "alerts")
		} else {
			ledgerMap["alerts"] = alerts
		}
		ledgerMap["incompliance"] = true
	}

	// save the original event
	ledgerMap["lastEvent"] = make(map[string]interface{})
	ledgerMap["lastEvent"].(map[string]interface{})["function"] = "deletePropertiesFromAsset"
	ledgerMap["lastEvent"].(map[string]interface{})["args"] = args[0]

	// Write the new state to the ledger
	stateJSON, err := json.Marshal(ledgerMap)
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s marshal failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// finally, put the new state
	err = stub.PutState(sAssetKey, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s PUTSTATE failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}
	err = pushRecentState(stub, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s push to recentstates failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// add history state
	err = updateStateHistory(stub, sAssetKey, string(stateJSON))
	if err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset AssetID %s of type %s push to history failed: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	return nil, nil
}

// ************************************
// deletaAllAssets
// ************************************
func (t *SimpleChaincode) deleteAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var sAssetKey string
	var err error

	if len(args) > 0 {
		err = errors.New("Too many arguments. Expecting none.")
		log.Error(err)
		return nil, err
	}

	aa, err := getActiveAssets(stub)
	if err != nil {
		err = fmt.Errorf("deleteAllAssets failed to get the active assets: %s", err)
		log.Error(err)
		return nil, err
	}
	for i := range aa {
		sAssetKey = aa[i]

		// Delete the key / asset from the ledger
		err = stub.DelState(sAssetKey)
		if err != nil {
			err = fmt.Errorf("deleteAllAssets arg %d AssetKey %s failed DELSTATE", i, sAssetKey)
			log.Error(err)
			return nil, err
		}
		// remove asset from contract state
		err = removeAssetFromContractState(stub, sAssetKey)
		if err != nil {
			err = fmt.Errorf("deleteAllAssets asset %s failed to remove asset from contract state: %s", sAssetKey, err)
			log.Critical(err)
			return nil, err
		}
		// save state history
		err = deleteStateHistory(stub, sAssetKey)
		if err != nil {
			err := fmt.Errorf("deleteAllAssets asset %s state history delete failed: %s", sAssetKey, err)
			log.Critical(err)
			return nil, err
		}
	}
	err = clearRecentStates(stub)
	if err != nil {
		err = fmt.Errorf("deleteAllAssets clearRecentStates failed: %s", err)
		log.Error(err)
		return nil, err
	}
	return nil, nil
}

// ************************************
// readAsset
// ************************************
func (t *SimpleChaincode) readAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var request interface{}
	var assetBytes []byte
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("Expecting one JSON event object")
		log.Error(err)
		return nil, err
	}

	requestBytes := []byte(args[0])
	log.Debugf("readAsset arg: %s", args[0])

	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		log.Errorf("readAsset failed to unmarshal arg: %s", err)
		return nil, err
	}

	argsMap, found = request.(map[string]interface{})
	if !found {
		err := errors.New("readAsset arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("readAsset arg does not include assetID")
			log.Error(err)
			return nil, err
		}
	}
	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	sMsg := "Inside readAsset assetName: " + assetName
	log.Info(sMsg)
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}
	sMsgTyoe := "Inside readAsset assetType: " + assetType
	log.Info(sMsgTyoe)
	sAssetKey := assetID + "_" + assetType
	found = assetIsActive(stub, sAssetKey)
	if !found {
		err := fmt.Errorf("readAsset arg asset %s of type %s does not exist", assetID, assetType)
		log.Error(err)
		return nil, err
	}

	// Get the state from the ledger
	assetBytes, err = stub.GetState(sAssetKey)
	if err != nil {
		log.Errorf("readAsset assetID %s of type %s failed GETSTATE", assetID, assetType)
		return nil, err
	}

	return assetBytes, nil
}

// ************************************
// readAllAssets
// ************************************
func (t *SimpleChaincode) readAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var sAssetKey string
	var err error
	var results []interface{}
	var state interface{}

	if len(args) > 0 {
		err = errors.New("readAllAssets expects no arguments")
		log.Error(err)
		return nil, err
	}

	aa, err := getActiveAssets(stub)
	if err != nil {
		err = fmt.Errorf("readAllAssets failed to get the active assets: %s", err)
		log.Error(err)
		return nil, err
	}
	results = make([]interface{}, 0, len(aa))
	for i := range aa {
		sAssetKey = aa[i]
		// Get the state from the ledger
		assetBytes, err := stub.GetState(sAssetKey)
		if err != nil {
			// best efforts, return what we can
			log.Errorf("readAllAssets assetID %s failed GETSTATE", sAssetKey)
			continue
		} else {
			err = json.Unmarshal(assetBytes, &state)
			if err != nil {
				// best efforts, return what we can
				log.Errorf("readAllAssets assetID %s failed to unmarshal", sAssetKey)
				continue
			}
			results = append(results, state)
		}
	}

	resultsStr, err := json.Marshal(results)
	if err != nil {
		err = fmt.Errorf("readallAssets failed to marshal results: %s", err)
		log.Error(err)
		return nil, err
	}

	return []byte(resultsStr), nil
}

// ************************************
// readAssetHistory
// ************************************
func (t *SimpleChaincode) readAssetHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assetBytes []byte
	var assetID string
	var assetType string
	var assetName string
	var argsMap ArgsMap
	var request interface{}
	var found bool
	var err error

	if len(args) != 1 {
		err = errors.New("readAssetHistory expects a JSON encoded object with assetID and count")
		log.Error(err)
		return nil, err
	}

	requestBytes := []byte(args[0])
	log.Debugf("readAssetHistory arg: %s", args[0])

	err = json.Unmarshal(requestBytes, &request)
	if err != nil {
		err = fmt.Errorf("readAssetHistory failed to unmarshal arg: %s", err)
		log.Error(err)
		return nil, err
	}

	argsMap, found = request.(map[string]interface{})
	if !found {
		err := errors.New("readAssetHistory arg is not a map shape")
		log.Error(err)
		return nil, err
	}

	// is assetID present or blank?
	assetIDBytes, found := getObject(argsMap, ASSETID)
	if found {
		assetID, found = assetIDBytes.(string)
		if !found || assetID == "" {
			err := errors.New("readAssetHistory arg does not include assetID")
			log.Error(err)
			return nil, err
		}
	}
	// Is asset name present?
	assetTypeBytes, found := getObject(argsMap, ASSETNAME)
	if found {
		assetName, found = assetTypeBytes.(string)
		if !found || assetName == "" {
			err := errors.New("createAsset arg does not include assetName ")
			log.Error(err)
			return nil, err
		}
	}
	if strings.Contains(assetName, "Plug") {
		assetType = "smartplug"
	} else {
		assetType = "motor"
	}
	sAssetKey := assetID + "_" + assetType
	found = assetIsActive(stub, sAssetKey)
	if !found {
		err := fmt.Errorf("readAssetHistory arg asset %s does not exist", assetID)
		log.Error(err)
		return nil, err
	}

	// Get the history from the ledger
	stateHistory, err := readStateHistory(stub, sAssetKey)
	if err != nil {
		err = fmt.Errorf("readAssetHistory assetID %s of type %s failed readStateHistory: %s", assetID, assetType, err)
		log.Error(err)
		return nil, err
	}

	// is count present?
	var olen int
	countBytes, found := getObject(argsMap, "count")
	if found {
		olen = int(countBytes.(float64))
	}
	if olen <= 0 || olen > len(stateHistory.AssetHistory) {
		olen = len(stateHistory.AssetHistory)
	}
	var hStatesOut = make([]interface{}, 0, olen)
	for i := 0; i < olen; i++ {
		var obj interface{}
		err = json.Unmarshal([]byte(stateHistory.AssetHistory[i]), &obj)
		if err != nil {
			log.Errorf("readAssetHistory JSON unmarshal of entry %d failed [%#v]", i, stateHistory.AssetHistory[i])
			return nil, err
		}
		hStatesOut = append(hStatesOut, obj)
	}
	assetBytes, err = json.Marshal(hStatesOut)
	if err != nil {
		log.Errorf("readAssetHistory failed to marshal results: %s", err)
		return nil, err
	}

	return []byte(assetBytes), nil
}

//***************************************************
//***************************************************
//* CONTRACT STATE
//***************************************************
//***************************************************

func (t *SimpleChaincode) readContractState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error

	if len(args) != 0 {
		err = errors.New("Too many arguments. Expecting none.")
		log.Error(err)
		return nil, err
	}

	// Get the state from the ledger
	chaincodeBytes, err := stub.GetState(CONTRACTSTATEKEY)
	if err != nil {
		err = fmt.Errorf("readContractState failed GETSTATE: %s", err)
		log.Error(err)
		return nil, err
	}

	return chaincodeBytes, nil
}

//***************************************************
//***************************************************
//* CONTRACT METADATA / SCHEMA INTERFACE
//***************************************************
//***************************************************

// ************************************
// readAssetSamples
// ************************************
func (t *SimpleChaincode) readAssetSamples(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(samples), nil
}

// ************************************
// readAssetSchemas
// ************************************
func (t *SimpleChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(schemas), nil
}

// ************************************
// readContractObjectModel
// ************************************
func (t *SimpleChaincode) readContractObjectModel(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var state = ContractState{MYVERSION, DEFAULTNICKNAME, make(map[string]bool)}

	stateJSON, err := json.Marshal(state)
	if err != nil {
		err := fmt.Errorf("JSON Marshal failed for get contract object model empty state: %+v with error [%s]", state, err)
		log.Error(err)
		return nil, err
	}
	return stateJSON, nil
}

// ************************************
// setLoggingLevel
// ************************************
func (t *SimpleChaincode) setLoggingLevel(stub shim.ChaincodeStubInterface, args []string) error {
	type LogLevelArg struct {
		Level string `json:"logLevel"`
	}
	var level LogLevelArg
	var err error
	if len(args) != 1 {
		err = errors.New("Incorrect number of arguments. Expecting a JSON encoded LogLevel.")
		log.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(args[0]), &level)
	if err != nil {
		err = fmt.Errorf("setLoggingLevel failed to unmarshal arg: %s", err)
		log.Error(err)
		return err
	}
	for i, lev := range logLevelNames {
		if strings.ToUpper(level.Level) == lev {
			(*log).SetLoggingLevel(LogLevel(i))
			return nil
		}
	}
	err = fmt.Errorf("Unknown Logging level: %s", level.Level)
	log.Error(err)
	return err
}

// CreateOnUpdate is a shared parameter structure for the use of
// the createonupdate feature
type CreateOnUpdate struct {
	CreateOnUpdate bool `json:"createOnUpdate"`
}

// ************************************
// setCreateOnUpdate
// ************************************
func (t *SimpleChaincode) setCreateOnUpdate(stub shim.ChaincodeStubInterface, args []string) error {
	var createOnUpdate CreateOnUpdate
	var err error
	if len(args) != 1 {
		err = errors.New("setCreateOnUpdate expects a single parameter")
		log.Error(err)
		return err
	}
	err = json.Unmarshal([]byte(args[0]), &createOnUpdate)
	if err != nil {
		err = fmt.Errorf("setCreateOnUpdate failed to unmarshal arg: %s", err)
		log.Error(err)
		return err
	}
	err = PUTcreateOnUpdate(stub, createOnUpdate)
	if err != nil {
		err = fmt.Errorf("setCreateOnUpdate failed to PUT setting: %s", err)
		log.Error(err)
		return err
	}
	return nil
}

// PUTcreateOnUpdate marshals the new setting and writes it to the ledger
func PUTcreateOnUpdate(stub shim.ChaincodeStubInterface, createOnUpdate CreateOnUpdate) (err error) {
	createOnUpdateBytes, err := json.Marshal(createOnUpdate)
	if err != nil {
		err = errors.New("PUTcreateOnUpdate failed to marshal")
		log.Error(err)
		return err
	}
	err = stub.PutState("CreateOnUpdate", createOnUpdateBytes)
	if err != nil {
		err = fmt.Errorf("PUTSTATE createOnUpdate failed: %s", err)
		log.Error(err)
		return err
	}
	return nil
}

// canCreateOnUpdate retrieves the setting from the ledger and returns it to the calling function
func canCreateOnUpdate(stub shim.ChaincodeStubInterface) bool {
	var createOnUpdate CreateOnUpdate
	createOnUpdateBytes, err := stub.GetState("CreateOnUpdate")
	if err != nil {
		err = fmt.Errorf("GETSTATE for canCreateOnUpdate failed: %s", err)
		log.Error(err)
		return true // true is the default
	}
	err = json.Unmarshal(createOnUpdateBytes, &createOnUpdate)
	if err != nil {
		err = fmt.Errorf("canCreateOnUpdate failed to marshal: %s", err)
		log.Error(err)
		return true // true is the default
	}
	return createOnUpdate.CreateOnUpdate
}
