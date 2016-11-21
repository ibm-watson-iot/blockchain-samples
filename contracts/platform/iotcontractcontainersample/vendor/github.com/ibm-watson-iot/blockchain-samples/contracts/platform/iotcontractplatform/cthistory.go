/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Howard McKinney- Initial Contribution
Kim Letkeman - Initial Contribution
*/

// v0.1 KL -- new iot chaincode platform
// v0.2 KL -- complete rewrite, history will be stored one state at a time so that
//            it poses minimal additional burden on state writes, read will use an iterator,
//            read will allow start and end time range, all, or last <n>

package iotcontractplatform

import (
	"encoding/json"
	"fmt"

	"time"

	"sort"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Changing to prepend history key so that an asset's history is separated from it's
// current state.

// STATEHISTORYKEY is used to separate history from current asset state and is prepended
// to the assetID
const STATEHISTORYKEY string = "IOTCP.HIST." // + assetKey + '.' + txnts

// AssetStateHistory is used to hold the output array of strings
type AssetStateHistory struct {
	AssetHistory []Asset `json:"assetHistory"`
}

// EmptyDateRange is used to determine that a range does not exist, which works in this
// case because the zero values of time are nonsense
var EmptyDateRange = DateRange{}

// DateRange allows a function to return states between a begin and end time, inclusive
type DateRange struct {
	DateRange struct {
		Begin string `json:"begin"`
		End   string `json:"end"`
	} `json:"daterange"`
}

// PUTAssetStateHistory write an Asset state with history key
func (a *Asset) PUTAssetStateHistory(stub shim.ChaincodeStubInterface) error {
	historyKey := STATEHISTORYKEY + a.AssetKey + "." + a.TXNTS.Format(time.RFC3339Nano)
	assetBytes, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("Failed to marshal Asset for history: %s", err)
		log.Error(err)
		return err
	}
	err = stub.PutState(historyKey, assetBytes)
	if err != nil {
		err = fmt.Errorf("Failed to PUT Asset history: %s", err)
		log.Error(err)
		return err
	}
	return nil
}

// DeleteAssetStateHistory deletes all history for an asset
func (c *AssetClass) DeleteAssetStateHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var arg = c.NewAsset()

	if err = arg.unmarshallEventIn(stub, args); err != nil {
		err := fmt.Errorf("DeleteAssetStateHistory for class %s could not unmarshall, err is %s", c.Name, err)
		log.Error(err)
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("DeleteAssetStateHistory for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Error(err)
		return nil, err
	}

	var historyKey = STATEHISTORYKEY + assetKey + "."

	iter, err := stub.RangeQueryState(historyKey, historyKey+"}")
	if err != nil {
		err = fmt.Errorf("DeleteAssetStateHistory failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		key, _, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("DeleteAssetStateHistory iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		err = stub.DelState(key)
		if err != nil {
			err = fmt.Errorf("DeleteAssetStateHistory DelState for asset %s failed: %s ", key, err)
			log.Errorf(err.Error())
			return nil, err
		}
	}

	return nil, nil
}

// ReadAssetStateHistory gets the state history for an asset.
func (c *AssetClass) ReadAssetStateHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var assets = make(AssetArray, 0)
	var err error
	var filter StateFilter
	var dr DateRange
	var arg = c.NewAsset()
	var begin string
	var end string

	if err = arg.unmarshallEventIn(stub, args); err != nil {
		err := fmt.Errorf("ReadAssetStateHistory for class %s could not unmarshall, err is %s", c.Name, err)
		log.Error(err)
		return nil, err
	}
	assetKey, err := arg.getAssetKey()
	if err != nil {
		err = fmt.Errorf("ReadAssetStateHistory for class %s could not find id at %s, err is %s", c.Name, c.AssetIDPath, err)
		log.Error(err)
		return nil, err
	}

	filter, err = getUnmarshalledStateFilter(args)
	if err != nil {
		err = fmt.Errorf("ReadAssetStateHistory failed while getting filter for %s %s, err is %s", c.Name, assetKey, err)
		log.Error(err)
		return nil, err
	}

	dr, err = getUnmarshalledDateRange(stub, args)
	if err != nil {
		err = fmt.Errorf("ReadAssetStateHistory failed while getting daterange for %s %s, err is %s", c.Name, assetKey, err)
		log.Error(err)
		return nil, err
	}

	if dr == EmptyDateRange {
		begin = ""
		end = "}"
	} else {
		begin = dr.DateRange.Begin
		end = dr.DateRange.End + "}"
	}

	var historyKey = STATEHISTORYKEY + assetKey + "."

	iter, err := stub.RangeQueryState(historyKey+begin, historyKey+end)
	if err != nil {
		err = fmt.Errorf("ReadAssetStateHistory failed to get a range query iterator: %s", err)
		log.Errorf(err.Error())
		return nil, err
	}
	defer iter.Close()
	for iter.HasNext() {
		key, assetBytes, err := iter.Next()
		if err != nil {
			err = fmt.Errorf("ReadAssetStateHistory iter.Next() failed: %s", err)
			log.Errorf(err.Error())
			return nil, err
		}
		var state = new(Asset)
		err = json.Unmarshal(assetBytes, state)
		if err != nil {
			err = fmt.Errorf("ReadAssetStateHistory unmarshal %s failed: %s", key, err)
			log.Errorf(err.Error())
			return nil, err
		}
		if state.Filter(filter) {
			assets = append(assets, *state)
		}
	}

	// return history, newest first
	sort.Sort(sort.Reverse(ByTimestamp(assets)))

	return json.Marshal(assets)
}

// Returns a date range found in the json object in args[0]
func getUnmarshalledDateRange(stub shim.ChaincodeStubInterface, args []string) (DateRange, error) {
	var dr DateRange
	var err error

	if len(args) == 0 {
		// perfectly normal to not have no date range
		return DateRange{}, nil
	}

	fBytes := []byte(args[0])
	err = json.Unmarshal(fBytes, &dr)
	if err != nil {
		err = fmt.Errorf("getUnmarshalledDateRange failed to unmarshal %s as a date range, error: %s", args[0], err)
		log.Error(err)
		return EmptyDateRange, err
	}

	return dr, nil
}
