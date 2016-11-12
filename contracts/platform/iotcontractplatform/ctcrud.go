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
	"errors"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// **************************************************
// CRUD utility functions
// **************************************************

// SystemClass acts as the class of all containers
var SystemClass = AssetClass{
	Name:        "system",
	Prefix:      "SYS",
	AssetIDPath: "n/a",
}

// Incoming asset CRUD events must have an assetID, which must be where the asset
// definition says it is. This function creates the world state representation by
// prepending the Prefix to it.
func (a *Asset) getAssetKey() (string, error) {
	assetID, found := GetObjectAsString(a.EventIn, a.Class.AssetIDPath)
	if !found {
		err := fmt.Errorf("getAssetID: %s not found", a.Class.AssetIDPath)
		log.Errorf(err.Error())
		return "", err
	}

	if assetID == "" {
		err := fmt.Errorf("getAssetID: %s is blank", a.Class.AssetIDPath)
		log.Errorf(err.Error())
		return "", err
	}

	a.AssetKey = a.Class.Prefix + assetID
	return a.AssetKey, nil
}

// The assetId is in the database and has > 0 bytes of info
func (c AssetClass) getAssetFromWorldState(stub shim.ChaincodeStubInterface, assetKey string) (stateout []byte, exists bool, err error) {
	stateout, err = stub.GetState(assetKey)
	if err != nil {
		err := fmt.Errorf("getAssetFromWorldState: GetState of %s returned error %s", assetKey, err)
		log.Errorf(err.Error())
		return nil, false, err
	}

	// some keys exist with no data in them
	if len(stateout) == 0 {
		err := fmt.Errorf("getAssetFromWorldState: state for asset %s is zero length", assetKey)
		log.Warningf(err.Error())
		// log, but suppress this error as the key can be reused (i.e. it is not really
		// there and does not show up when iterating)
		return nil, false, nil
	}

	// We do not Unmarshal here because the result is ignored more often than it is used

	return stateout, true, nil
}

// Decodes args[0], which must be a map containing a JSON object representing
// a partial state containing one or more direct readings for specific state
// properties (e.g. gForce, temperature, location, etc.)
func (a *Asset) unmarshallEventIn(stub shim.ChaincodeStubInterface, args []string) error {
	var event interface{}
	var err error

	if len(args) != 1 && len(args) != 2 {
		err = errors.New("Expecting a JSON event [and optional redirect function name]")
		log.Errorf(err.Error())
		return err
	}

	eventBytes := []byte(args[0])
	err = json.Unmarshal(eventBytes, &event)
	if err != nil {
		err = fmt.Errorf("%s failed to unmarshal arg: %s", a.Class, err)
		log.Errorf(err.Error())
		return err
	}

	if event == nil {
		err = fmt.Errorf("%s unmarshal arg created nil event", a.Class)
		log.Errorf(err.Error())
		return err
	}

	amap, found := AsMap(event)
	if !found {
		err := fmt.Errorf("%s arg is not a map shape", a.Class)
		log.Errorf(err.Error())
		return err
	}
	a.EventIn = &amap

	return nil
}

// // Returns the world state represented by prefix + assetID unmarshalled.
// func (c AssetClass) getUnmarshalledState(stub shim.ChaincodeStubInterface, assetID string) (*Asset, error) {
//     var stateBytes []byte
//     var err error

//     stateBytes, exists, err := c.getAssetFromWorldState(stub, assetID)
//     if err != nil {
//         err := fmt.Errorf("getUnmarshalledState for class %s asset %s read from world state returned error %s", c.Name, assetID, err)
//         log.Errorf(err.Error())
//         return nil, err
//     }
//     if !exists {
//         err := fmt.Errorf("getUnmarshalledState for class %s asset %s asset does not exist", c.Name, assetID)
//         log.Errorf(err.Error())
//         return nil, err
//     }

//     var a Asset
//     // unmarshal the existing state from the ledger to theinterface
//     err = json.Unmarshal(stateBytes, &a)
//     if err != nil {
//         log.Errorf("%s: assetID %s unmarshal failed: %s", c.Name, assetID, err)
//         return nil, err
//     }

//     return &a, nil
// }

// Pushes state to the ledger using assetID, which is expected to be prefixed.
func (a *Asset) putMarshalledState(stub shim.ChaincodeStubInterface) error {
	// Write the new state to the ledger
	stateJSON, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("putMarshalledState: assetID %s marshal failed: %s", a.AssetKey, err)
		log.Errorf(err.Error())
		return err
	}

	err = stub.PutState(a.AssetKey, []byte(stateJSON))
	if err != nil {
		err = fmt.Errorf("putMarshalledState: PUTSTATE for assetID %s failed: %s", a.AssetKey, err)
		log.Errorf(err.Error())
		return err
	}

	err = a.PushRecentState(stub)
	if err != nil {
		err = fmt.Errorf("%s: assetID %s push recent states failed: %s", a.Class.Name, a.AssetKey, err)
		log.Errorf(err.Error())
		return err
	}

	// // add history state
	// err = h.UpdateStateHistory(stub, assetID, string(stateJSON))
	// if err != nil {
	//     err = fmt.Errorf("%s: event %s assetID %s push history failed: %s", caller, eventName, assetID, err)
	//     log.Errorf(err.Error())
	//     return err
	// }
	return nil
}

// RemoveOneAssetFromWorldState remove the asset from world state
func removeOneAssetFromWorldState(stub shim.ChaincodeStubInterface, assetKey string) error {
	err := stub.DelState(assetKey)
	if err != nil {
		err = fmt.Errorf("removeOneAssetFromWorldState: asset %s failed", assetKey)
		log.Error(err)
		return err
	}
	err = RemoveAssetFromRecentStates(stub, assetKey)
	if err != nil {
		err = fmt.Errorf("removeOneAssetFromWorldState: asset %s could not be removed from recent states: %s", assetKey, err)
		log.Error(err)
		return err
	}
	// err = h.DeleteStateHistory(stub, assetID)
	// if err != nil {
	// 	err = fmt.Errorf("%s: %s assetID %s history deletion failed", caller, assetName, assetID)
	// 	log.Error(err)
	// 	return err
	// }
	return nil
}

// Adds the current transaction timestamp into world state, replacing it if it was already there.
func (a *Asset) addTXNTimestampToState(stub shim.ChaincodeStubInterface) error {
	// add transaction uuid and timestamp
	a.TXNID = stub.GetTxID()
	txnunixtime, err := stub.GetTxTimestamp()
	if err != nil {
		err = fmt.Errorf("error getting transaction timestamp, err is %s", err)
		log.Errorf(err.Error())
		return err
	}
	txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
	a.TXNTS = &txntimestamp
	return nil
}

// ********** property injection implementation
func (a *Asset) injectProps(qprops []QPropNV) error {
	var ok bool
	for _, qp := range qprops {
		ok = PutObject(a.State, qp.QProp, qp.Value)
		if !ok {
			err := fmt.Errorf("injectProps->putObject failed to put %s:%s to state %#v", qp.QProp, qp.Value, a)
			log.Errorf(err.Error())
			return err
		}
	}
	return nil
}
