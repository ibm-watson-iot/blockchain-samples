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

package main

import (
	"encoding/json"
    "fmt"
    "sort"
    "strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func createAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    var state = make(map[string]interface{})
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }
    // We have a valid assetID in internal format, so verify whether it already exists.
    _, err = assetIsActive(stub, assetID)
    if err == nil {
        err = fmt.Errorf("%s: asset %s already exists", caller, assetID)
        log.Error(err)
        return nil, err
    }
    state = deepMerge(argsMap, make(map[string]interface{}))
    state, err = addTXNTimestampToState(stub, caller, state)
    if err != nil { return nil, err }
    state = addLastEventToState(stub, caller, argsMap, state, "")
    state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
    if err != nil { return nil, err }
    err = putMarshalledState(stub, caller, assetName, assetID, state)
    if err != nil { return nil, err }
    return nil, nil
}

func updateAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }
    state, err := getUnmarshalledState(stub, caller, assetID)
    if err != nil { return nil, err }
    state = deepMerge(argsMap, state)
    state, err = addTXNTimestampToState(stub, caller, state)
    if err != nil { return nil, err }
    state = addLastEventToState(stub, caller, argsMap, state, "")
    state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
    if err != nil { return nil, err }
    err = putMarshalledState(stub, caller, assetName, assetID, state)
    if err != nil { return nil, err }
    return nil, nil
}

func deleteAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }
    err = removeOneAssetFromWorldState(stub, caller, assetName, assetID)
    if err != nil { return nil, err }
    return nil, nil
}

func deleteAllAssets(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
	var err error
    prefix, err := eventNameToAssetPrefix(assetName)
    if err != nil { return nil, err }
    iter, err := stub.RangeQueryState(prefix, prefix + "}")
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

func deletePropertiesFromAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }

    var qprops []interface{}
    qpropsBytes, found := getObject(argsMap, "qualPropsToDelete")
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
    if err != nil { return nil, err }

    log.Debugf("deleteProps: %s state follows *************\n\n%s", assetName, state)

    // now remove properties from state, they are qualified by level
    OUTERDELETELOOP:
    for p := range qprops {
        prop := qprops[p].(string)
        log.Debugf("deletePropertiesFromAsset AssetID %s deleting qualified property: %s", assetID, prop)
        // TODO Ugly, isolate in a function at some point
        if (CASESENSITIVEMODE  && strings.HasSuffix(prop, ASSETID)) ||
           (!CASESENSITIVEMODE && strings.HasSuffix(strings.ToLower(prop), strings.ToLower(ASSETID))) {
            log.Warningf("deletePropertiesFromAsset AssetID %s cannot delete protected qualified property: %s", assetID, prop)
        } else {
            levels := strings.Split(prop, ".")
            lm := (map[string]interface{})(state)
            for l := range levels {
                // lev is the name of a level
                lev := levels[l]
                if l == len(levels)-1 {
                    // we're here, delete the actual property name from this level of the map
                    levActual, found := findMatchingKey(lm, lev)
                    if !found {
                        log.Warningf("deletePropertiesFromAsset AssetID %s property match %s not found", assetID, lev)
                        continue OUTERDELETELOOP
                    }
                    log.Debugf("deletePropertiesFromAsset AssetID %s deleting %s", assetID, prop)
                    delete(lm, levActual)
                } else {
                    // navigate to the next level object
                    log.Debugf("deletePropertiesFromAsset AssetID %s navigating to level %s", assetID, lev)
                    lmBytes, found := findObjectByKey(lm, lev)
                    if found {
                        lm, found = lmBytes.(map[string]interface{})
                        if !found {
                            log.Noticef("deletePropertiesFromAsset AssetID %s level %s not found in ledger", assetID, lev)
                            continue OUTERDELETELOOP
                        }
                    } 
                }
            } 
        }
    }

    state, err = addTXNTimestampToState(stub, caller, state)
    if err != nil { return nil, err }
    state = addLastEventToState(stub, caller, argsMap, state, "")
    state, err = handleAlertsAndRules(stub, caller, assetName, assetID, argsMap, state)
    if err != nil { return nil, err }
    err = putMarshalledState(stub, caller, assetName, assetID, state)
    if err != nil { return nil, err }
    return nil, nil
}

func readAsset(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }
    assetBytes, err := assetIsActive(stub, assetID)
    if err != nil { return nil, err }
    return assetBytes, nil
}

func readAllAssets(stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    var assets ByAssetID
	var err error
    var state interface{}

    prefix, err := eventNameToAssetPrefix(assetName)
    if err != nil {
        err = fmt.Errorf("readAllAssets assetName %s has no prefix: %s", assetName, err.Error())
		log.Error(err)
        return nil, err
    }
    iter, err := stub.RangeQueryState(prefix, prefix + "}")
    if err != nil {
        err = fmt.Errorf("readAllAssets failed to get a range query iterator: %s", err)
		log.Error(err)
        return nil, err
    }
    defer iter.Close()
    for iter.HasNext() {
        assetID, assetBytes, err := iter.Next()
        if err != nil {
            err = fmt.Errorf("readAllAssets iter.Next() failed: %s", err)
            log.Error(err)
            return nil, err
        }
        err = json.Unmarshal(assetBytes, &state)
        if err != nil {
            err = fmt.Errorf("readAllAssets unmarshal failed: %s", err)
            log.Error(err)
            return nil, err
        }
        assets = append(assets, AssetArr{AssetID:assetID, Asset:state})
    }

    if len(assets) == 0 {
        return []byte("[]"), nil
    }

	sort.Sort(ByAssetID(assets))

    var results []interface{} 
    for _, a := range assets {
        results = append(results, a.Asset)
    }
    resultsBytes, err := json.Marshal(&results)
    if err != nil {
        err = fmt.Errorf("readAllAssets failed to marshal sorted assets structure: %s", err)
        log.Error(err)
        return nil, err
    }
    return resultsBytes, nil
}

func readAssetHistory (stub *shim.ChaincodeStub, args []string, assetName string, caller string) ([]byte, error) {
    argsMap, err := getUnmarshalledArgument(stub, caller, args)
    if err != nil { return nil, err }
    assetID, err := validateAssetID(caller, assetName, argsMap)
    if err != nil { return nil, err }
    stateHistory, err := readStateHistory(stub, assetID)
    if err != nil { return nil, err }
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
	return prettyPrint(a.Asset)
}

// ByAssetID implements sort.Interface for []Asset based on
// the AssetID field.
type ByAssetID []AssetArr

func (a ByAssetID) Len() int           { return len(a) }
func (a ByAssetID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAssetID) Less(i, j int) bool { return a[i].AssetID < a[j].AssetID }


