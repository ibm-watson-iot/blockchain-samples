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

// v0.1 KL -- new iot chaincode platform

package iotcontractplatform

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// InvokeEvent carries the event that is to be set upon exit from the chaincode
type InvokeEvent struct {
	Name    string                 `json:"name"`
	Payload map[string]interface{} `json:"payload"`
}

// AssetClass defines a receiver for rules and other class-specific execution
type AssetClass struct {
	Name        string `json:"name"`        // asset class name
	Prefix      string `json:"prefix"`      // asset class prefix for key in world state
	AssetIDPath string `json:"assetIDpath"` // property that is unique key for this class
}

// NewAsset create an instance of an asset class
func (c AssetClass) NewAsset() Asset {
	var a = Asset{
		c, "", nil, nil, "", "", nil, nil, AlertNameArray(make([]AlertName, 0)), true,
	}
	return a
}

// AllAssetClass is the class of all assets
var AllAssetClass = AssetClass{"All", "", ""}

func (c AssetClass) String() string {
	return fmt.Sprintf("CLS=%s | PRF=%s | ID=%s", c.Name, c.Prefix, c.AssetIDPath)
}

// Asset is a type that holds all information about an asset, including its name,
// its world state prefix, and the qualified property name that is its assetID
type Asset struct {
	Class        AssetClass              `json:"assetclass"`         // asset's classifier with metadata
	AssetKey     string                  `json:"assetkey"`           // asset's world state key
	State        *map[string]interface{} `json:"assetstate"`         // asset's current state
	EventIn      *map[string]interface{} `json:"eventpayload"`       // most recent event body
	FunctionIn   string                  `json:"eventfunction"`      // most recent event function
	TXNID        string                  `json:"txnid"`              // transaction UUID matching blockchain
	TXNTS        *time.Time              `json:"txnts,omitempty"`    // transaction timestamp matching blockchain
	EventOut     *InvokeEvent            `json:"eventout,omitempty"` // event (if any) emitted upon exit from an invoke
	AlertsActive AlertNameArray          `json:"alerts,omitempty"`   // array of active alerts
	Compliant    bool                    `json:"compliant"`          // true if the asset complies with the contract terms
}

// AssetArray is an array of assets, used by read all, recent states, history, etc.
type AssetArray []Asset

func (a Asset) String() string {
	return PrettyPrint(a)
}

func (aa AssetArray) String() string {
	return PrettyPrint(aa)
}

// CreateAsset inializes a new asset and stores it in world state
func (c *AssetClass) CreateAsset(stub shim.ChaincodeStubInterface, args []string, caller string, inject []QPropNV) ([]byte, error) {

	var a = c.NewAsset()

	if err := a.unmarshallEventIn(stub, args); err != nil {
		err = fmt.Errorf("CreateAsset for class %s could not unmarshall, err is %s", c.Name, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetKey, err := a.getAssetKey()
	if err != nil {
		err = fmt.Errorf("CreateAsset for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	_, exists, err := c.getAssetFromWorldState(stub, assetKey)
	if err != nil {
		err := fmt.Errorf("CreateAsset for class %s asset %s read from world state returned error %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if exists {
		err := fmt.Errorf("CreateAsset for class %s asset %s asset already exists", c.Name, a.AssetKey)
		log.Errorf(err.Error())
		return nil, err
	}

	// copy the event into a new state
	astate := DeepMergeMap(*a.EventIn, make(map[string]interface{}))
	a.State = &astate
	if err := a.addTXNTimestampToState(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed to add txn timestamp for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}

	// save original asset function
	a.FunctionIn = caller

	if len(inject) > 0 {
		err := a.injectProps(inject)
		if err != nil {
			err = fmt.Errorf("CreateAsset for class %s failed to inject properties %+v for %s, err is %s", c.Name, inject, a.AssetKey, err)
			log.Errorf(err.Error())
			return nil, err
		}
	}
	if err := a.ExecuteRules(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed in rules engine for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if err := a.putMarshalledState(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed to marshall for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	return nil, nil
}

// UpdateAsset updates an asset and stores it in world state
func (c *AssetClass) UpdateAsset(stub shim.ChaincodeStubInterface, args []string, caller string, inject []QPropNV) ([]byte, error) {

	var arg = c.NewAsset()
	var a = c.NewAsset()

	if err := arg.unmarshallEventIn(stub, args); err != nil {
		err = fmt.Errorf("UpdateAsset for class %s could not unmarshall, err is %s", c.Name, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("UpdateAsset for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetBytes, exists, err := c.getAssetFromWorldState(stub, assetKey)
	if err != nil {
		err := fmt.Errorf("UpdateAsset for class %s asset %s read from world state returned error %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if !exists {
		err := fmt.Errorf("UpdateAsset for class %s asset %s asset does not exist", c.Name, a.AssetKey)
		log.Errorf(err.Error())
		return nil, err
	}
	err = json.Unmarshal(assetBytes, &a)
	if err != nil {
		err := fmt.Errorf("UpdateAsset for class %s asset %s Unmarshal failed with err %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	// save the incoming EventIn
	a.EventIn = arg.EventIn
	a.FunctionIn = arg.FunctionIn

	// merge the event into the state
	astate := DeepMergeMap(*a.EventIn, *a.State)
	a.State = &astate

	if err := a.addTXNTimestampToState(stub); err != nil {
		err = fmt.Errorf("UpdateAsset for class %s failed to add txn timestamp for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}

	// save original asset function
	a.FunctionIn = caller

	if len(inject) > 0 {
		err := a.injectProps(inject)
		if err != nil {
			err = fmt.Errorf("UpdateAsset for class %s failed to inject properties %+v for %s, err is %s", c.Name, inject, a.AssetKey, err)
			log.Errorf(err.Error())
			return nil, err
		}
	}
	if err := a.ExecuteRules(stub); err != nil {
		err = fmt.Errorf("UpdateAsset for class %s failed in rules engine for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if err := a.putMarshalledState(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed to to marshall for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}

	return nil, nil
}

// DeleteAsset deletes an asset from world state
func (c *AssetClass) DeleteAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var arg = c.NewAsset()

	if err := arg.unmarshallEventIn(stub, args); err != nil {
		err := fmt.Errorf("DeleteAsset for class %s could not unmarshall, err is %s", c.Name, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("DeleteAsset for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	err = arg.removeOneAssetFromWorldState(stub)
	if err != nil {
		err := fmt.Errorf("DeleteAsset: removeOneAssetFromWorldState class %s, asset %s, returned error: %s", c.Name, assetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	return nil, nil
}

// DeleteAllAssets reletes all asstes of a specific asset class from world state
func (c *AssetClass) DeleteAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var filter StateFilter

	filter, err := getUnmarshalledStateFilter(stub, args)
	if err != nil {
		err = fmt.Errorf("DeleteAllAssets failed to get the filter: %s", err)
		log.Error(err)
		return nil, err
	}
	iter, err := stub.RangeQueryState(c.Prefix, c.Prefix+"}")
	if err != nil {
		err = fmt.Errorf("DeleteAllAssets failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		key, stateBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("DeleteAllAssets iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		var state Asset
		err = json.Unmarshal(stateBytes, &state)
		if err != nil {
			err = fmt.Errorf("DeleteAllAssets state unmarshal failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		if state.Filter(filter) {
			err = state.removeOneAssetFromWorldState(stub)
			if err != nil {
				err = fmt.Errorf("DeleteAllAssets removeOneAssetFromWorldState for asset %s failed: %s", key, err)
				log.Errorf(err.Error())
				return nil, err
			}
		}
	}

	return nil, nil
}

// DeletePropertiesFromAsset removes specific properties from an asset in world state
func (c *AssetClass) DeletePropertiesFromAsset(stub shim.ChaincodeStubInterface, args []string, caller string, inject []QPropNV) ([]byte, error) {

	var arg = c.NewAsset()
	var a = c.NewAsset()

	if err := arg.unmarshallEventIn(stub, args); err != nil {
		err = fmt.Errorf("DeletePropertiesFromAsset for class %s could not unmarshall, err is %s", c.Name, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("DeletePropertiesFromAsset for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetBytes, exists, err := c.getAssetFromWorldState(stub, assetKey)
	if err != nil {
		err := fmt.Errorf("DeletePropertiesFromAsset for class %s asset %s read from world state returned error %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if !exists {
		err := fmt.Errorf("DeletePropertiesFromAsset for class %s asset %s asset does not exist", c.Name, a.AssetKey)
		log.Errorf(err.Error())
		return nil, err
	}
	err = json.Unmarshal(assetBytes, &a)
	if err != nil {
		err := fmt.Errorf("DeletePropertiesFromAsset for class %s asset %s Unmarshal failed with err %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	// save the incoming EventIn
	a.EventIn = arg.EventIn
	a.FunctionIn = arg.FunctionIn

	var qprops []string
	qprops, found := GetObjectAsStringArray(arg.EventIn, "qprops")
	if !found {
		err = fmt.Errorf("deletePropertiesFromAsset asset %s has no qprops argument or qprops not a string array", assetKey)
		log.Errorf(err.Error())
		return nil, err
	}

	// remove qualified properties from state
	for _, p := range qprops {
		_ = RemoveObject(a.State, p)
	}

	if err := a.addTXNTimestampToState(stub); err != nil {
		err = fmt.Errorf("deletePropertiesFromAsset for class %s failed to add txn timestamp for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}

	// save original asset function
	a.FunctionIn = caller

	if len(inject) > 0 {
		err := a.injectProps(inject)
		if err != nil {
			err = fmt.Errorf("deletePropertiesFromAsset for class %s failed to inject properties %+v for %s, err is %s", c.Name, inject, a.AssetKey, err)
			log.Errorf(err.Error())
			return nil, err
		}
	}
	if err := a.ExecuteRules(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed in rules engine for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if err := a.putMarshalledState(stub); err != nil {
		err = fmt.Errorf("CreateAsset for class %s failed to marshall for %s, err is %s", c.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}

	return nil, nil
}

// ReadAsset returns an asset from world state, intended to be returned directly to a client
func (c *AssetClass) ReadAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var arg = c.NewAsset()

	if err := arg.unmarshallEventIn(stub, args); err != nil {
		err := fmt.Errorf("ReadAsset for class %s could not unmarshall, err is %s", c.Name, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("ReadAsset for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Errorf(err.Error())
		return nil, err
	}
	assetBytes, exists, err := c.getAssetFromWorldState(stub, assetKey)
	if err != nil {
		err := fmt.Errorf("ReadAsset for class %s, asset %s returned error: %s", c.Name, assetKey, err)
		log.Errorf(err.Error())
		return nil, err
	}
	if !exists {
		err := fmt.Errorf("ReadAsset for class %s, asset %s does not exist", c.Name, assetKey)
		log.Errorf(err.Error())
		return nil, err
	}
	return assetBytes, nil
}

// ReadAllAssets returns all assets of a specific class from world state as an array
func (c AssetClass) ReadAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	results, err := c.ReadAllAssetsUnmarshalled(stub, args)
	if err != nil {
		return nil, err
	}
	resultsBytes, err := json.Marshal(&results)
	if err != nil {
		err = fmt.Errorf("readAllAssets failed to marshal assets structure: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	return resultsBytes, nil
}

// ReadAllAssetsUnmarshalled returns all assets of a specific class from world state as an object, intended for internal use
func (c AssetClass) ReadAllAssetsUnmarshalled(stub shim.ChaincodeStubInterface, args []string) (AssetArray, error) {
	var assets AssetArray
	var err error
	var filter StateFilter

	filter, err = getUnmarshalledStateFilter(stub, args)
	if err != nil {
		err = fmt.Errorf("readAllAssetsUnmarshalled failed to get a filter: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}

	iter, err := stub.RangeQueryState(c.Prefix, c.Prefix+"}")
	if err != nil {
		err = fmt.Errorf("readAllAssetsUnmarshalled failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		key, assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("readAllAssetsUnmarshalled iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		var state = new(Asset)
		err = json.Unmarshal(assetBytes, state)
		if err != nil {
			err = fmt.Errorf("readAllAssetsUnmarshalled unmarshal %s failed: %s", key, err)
			log.Errorf(err.Error())
			return nil, err
		}
		if state.Filter(filter) {
			assets = append(assets, *state)
		}
	}

	if len(assets) == 0 {
		return make(AssetArray, 0), nil
	}

	sort.Sort(assets)

	return assets, nil
}

//********** sort interface for AssetArray

func (aa AssetArray) Len() int           { return len(aa) }
func (aa AssetArray) Swap(i, j int)      { aa[i], aa[j] = aa[j], aa[i] }
func (aa AssetArray) Less(i, j int) bool { return aa[i].AssetKey < aa[j].AssetKey }

// ByTimestamp alias for sorting by timestamp
type ByTimestamp AssetArray

func (aa ByTimestamp) Len() int           { return len(aa) }
func (aa ByTimestamp) Swap(i, j int)      { aa[i], aa[j] = aa[j], aa[i] }
func (aa ByTimestamp) Less(i, j int) bool { return (*aa[i].TXNTS).Before(*aa[j].TXNTS) }
