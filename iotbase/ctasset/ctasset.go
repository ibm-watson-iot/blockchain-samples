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

// v1 KL 09 Aug 2016 Creation of assetUtils as boilerplate for any asset to call for standard
//                   crud like behaviors. Make extensive use of crudUtils.
// v2 KL 02 Nov 2016 new package ctasset

package ctasset

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	cf "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctconfig"
	h "github.com/ibm-watson-iot/blockchain-samples/iotbase/cthistory"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctstate"
	"github.com/op/go-logging"
	"sort"
)

// Asset is a type that holds all information about an asset, including its name,
// its world state prefix, and the qualified property name that is its assetID
type Asset struct {
	Name      string                 `json:"name"`
	Wspref    string                 `json:"wspref"`
	QPAssetID string                 `json:"qpassetID"`
	AssetID   string                 `json:"assetID"`
	State     map[string]interface{} `json:"state"`
}

// NewAsset constructs an asset object with nil state
func NewAsset(n string, p string, q string, a string) Asset {
	return Asset{
		Name:      n,
		Wspref:    p,
		QPAssetID: q,
		AssetID:   a,
		State:     nil,
	}
}

// Logger for the ctstate package
var log = logging.MustGetLogger("asst")

// CreateAsset inializes a new asset and stores it in world state
func CreateAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string, inject []QPropNV) ([]byte, error) {
	var state interface{} = make(map[string]interface{})
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}
	// We have a valid assetID in internal format, so verify whether it already exists.
	assetBytes, err := assetIsActive(stub, assetID)
	if err == nil && len(assetBytes) > 0 {
		err = fmt.Errorf("%s: asset %s already exists", caller, assetID)
		log.Error(err)
		return nil, err
	}
	state = st.DeepMerge(argsMap, make(map[string]interface{}))
	state, err = addTXNTimestampToState(stub, caller, state)
	if err != nil {
		return nil, err
	}
	state = addLastEventToState(stub, caller, argsMap, state, "")
	if state == nil {
		return nil, errors.New("addLastEventToState failed")
	}
	state, err = injectProps(state, inject)
	if err != nil {
		return nil, err
	}
	state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
	if err != nil {
		return nil, err
	}
	err = putMarshalledState(stub, caller, assetName, assetID, state)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateAsset updates an asset and stores it in world state
func UpdateAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string, inject []QPropNV) ([]byte, error) {
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}
	state, err := getUnmarshalledState(stub, caller, assetID)
	if err != nil {
		return nil, err
	}
	state = st.DeepMerge(argsMap, state)
	state, err = addTXNTimestampToState(stub, caller, state)
	if err != nil {
		return nil, err
	}
	state = addLastEventToState(stub, caller, argsMap, state, "")
	if state == nil {
		return nil, errors.New("addLastEventToState failed")
	}
	state, err = injectProps(state, inject)
	if err != nil {
		return nil, err
	}
	state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
	if err != nil {
		return nil, err
	}
	err = putMarshalledState(stub, caller, assetName, assetID, state)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteAsset deletes an asset from world state
func DeleteAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}
	err = removeOneAssetFromWorldState(stub, caller, assetName, assetID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteAllAssets reletes all asstes of a specific asset class from world state
func DeleteAllAssets(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	var err error
	prefix, err := cf.EventNameToAssetPrefix(assetName)
	if err != nil {
		return nil, err
	}
	iter, err := stub.RangeQueryState(prefix, prefix+"}")
	if err != nil {
		err = fmt.Errorf("deleteAllAssets failed to get a range query iterator: %s", err)
		log.Error(err)
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		assetID, _, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("deleteAllAssets iter.Next() failed: %s", err)
			log.Error(err)
			return nil, err
		}
		err = removeOneAssetFromWorldState(stub, caller, assetName, assetID)
		if err != nil {
			err = fmt.Errorf("deleteAllAssets%s failed to remove an asset: %s", assetName, err)
			log.Error(err)
			// continue best efforts?
		}
	}
	return nil, nil
}

// DeletePropertiesFromAsset removes specific properties from an asset in world state
func DeletePropertiesFromAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string, inject []QPropNV) ([]byte, error) {
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}

	var qprops []interface{}
	qpropsBytes, found := st.GetObject(argsMap, "qualPropsToDelete")
	if found {
		qprops, found = qpropsBytes.([]interface{})
		if !found || len(qprops) < 1 {
			err = fmt.Errorf("deletePropertiesFromAsset asset %s qualPropsToDelete is type %T", assetID, qpropsBytes)
			log.Error(err)
			return nil, err
		}
	} else {
		err = fmt.Errorf("deletePropertiesFromAsset asset %s has no qualPropsToDelete argument", assetID)
		log.Error(err)
		return nil, err
	}

	state, err := getUnmarshalledState(stub, caller, assetID)
	if err != nil {
		return nil, err
	}

	log.Debugf("deleteProps: %s state follows *************\n\n%s", assetName, state)

	// now remove properties from state, they are qualified by level
	for p := range qprops {
		prop := qprops[p].(string)
		log.Debugf("deletePropertiesFromAsset deleting property %s from asset %s", prop, assetID)
		state, _ = st.RemoveObject(state, prop)
	}

	state, err = addTXNTimestampToState(stub, caller, state)
	if err != nil {
		return nil, err
	}
	state = addLastEventToState(stub, caller, argsMap, state, "")
	if state == nil {
		return nil, errors.New("addLastEventToState failed")
	}
	state, err = injectProps(state, inject)
	if err != nil {
		return nil, err
	}
	state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
	if err != nil {
		return nil, err
	}
	err = putMarshalledState(stub, caller, assetName, assetID, state)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ReadAsset returns an asset from world state, intended to be returned directly to a client
func ReadAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}
	assetBytes, err := assetIsActive(stub, assetID)
	if err != nil {
		// something went wrong
		err = fmt.Errorf("Asset %s with ID %s not found, err==%s", assetName, assetID, err.Error())
		return nil, err
	}
	return assetBytes, nil
}

// ReadAssetUnmarshalled returns the asset from world state as an object, intended for internal use
func ReadAssetUnmarshalled(stub *shim.ChaincodeStub, assetID string, assetName string, caller string) (interface{}, error) {
	assetBytes, err := assetIsActive(stub, assetID)
	if err != nil || len(assetBytes) == 0 {
		return nil, err
	}
	var state interface{}
	err = json.Unmarshal(assetBytes, &state)
	if err != nil {
		err = fmt.Errorf("readAssetUnmarshalled unmarshal failed: %s", err)
		log.Error(err)
		return nil, err
	}
	return state, nil
}

// ReadAllAssets returns all assets of a specific class from world state as an array
func ReadAllAssets(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	results, err := ReadAllAssetsUnmarshalled(stub, args, assetName, caller)
	if err != nil {
		return nil, err
	}
	resultsBytes, err := json.Marshal(&results)
	if err != nil {
		err = fmt.Errorf("readAllAssets failed to marshal assets structure: %s", err)
		log.Error(err)
		return nil, err
	}
	return resultsBytes, nil
}

// ReadAllAssetsUnmarshalled returns all assets of a specific class from world state as an object, intended for internal use
func ReadAllAssetsUnmarshalled(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]interface{}, error) {
	var assets ByAssetID
	var err error
	var state interface{}
	var filter StateFilter

	prefix, err := cf.EventNameToAssetPrefix(assetName)
	if err != nil {
		err = fmt.Errorf("readAllAssetsUnmarshalled assetName %s has no prefix: %s", assetName, err.Error())
		log.Error(err)
		return nil, err
	}

	filter = getUnmarshalledStateFilter(stub, caller, args)
	//log.Debugf("%s: got filter: %+v from args %+v\n", caller, filter, args)

	iter, err := stub.RangeQueryState(prefix, prefix+"}")
	if err != nil {
		err = fmt.Errorf("readAllAssetsUnmarshalled failed to get a range query iterator: %s", err)
		log.Error(err)
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		assetID, assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("readAllAssetsUnmarshalled iter.Next() failed: %s", err)
			log.Error(err)
			return nil, err
		}
		err = json.Unmarshal(assetBytes, &state)
		if err != nil {
			err = fmt.Errorf("readAllAssetsUnmarshalled unmarshal failed: %s", err)
			log.Error(err)
			return nil, err
		}
		//log.Debugf("%s: about to filter state: %+v using %+v\n", caller, filter)
		if len(filter.Entries) == 0 || filterObject(state, filter) {
			//log.Debugf("%s: filter passed\n", caller)
			assets = append(assets, AssetArr{AssetID: assetID, Asset: state})
		}
	}

	//log.Debugf("%s: Final assets list: %+v\n", caller, assets)
	if len(assets) == 0 {
		return make([]interface{}, 0), nil
	}

	sort.Sort(ByAssetID(assets))

	var results []interface{}
	for _, a := range assets {
		results = append(results, a.Asset)
	}
	return results, nil
}

// ReadAssetHistory returns an asset's history from world state as an array
func ReadAssetHistory(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	argsMap, err := getUnmarshalledArgument(stub, caller, args)
	if err != nil {
		return nil, err
	}
	assetID, err := validateAssetID(caller, assetName, argsMap)
	if err != nil {
		return nil, err
	}
	stateHistory, err := h.ReadStateHistory(stub, assetID)
	if err != nil {
		return nil, err
	}
	// is count present?
	var olen int
	countBytes, found := st.GetObject(argsMap, "count")
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
	assetBytes, err := json.Marshal(hStatesOut)
	if err != nil {
		log.Errorf("readAssetHistory failed to marshal results: %s", err)
		return nil, err
	}

	return []byte(assetBytes), nil
}

//********** implement sort interface for assetID

// AssetArr is a simple sort structure with the assetID called out for sorting
// Used by read all assets
type AssetArr struct {
	AssetID string
	Asset   interface{}
}

func (a AssetArr) String() string {
	return st.PrettyPrint(a.Asset)
}

// ByAssetID implements sort.Interface for []Asset based on
// the AssetID field.
type ByAssetID []AssetArr

func (a ByAssetID) Len() int           { return len(a) }
func (a ByAssetID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAssetID) Less(i, j int) bool { return a[i].AssetID < a[j].AssetID }
