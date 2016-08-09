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


// v1 KL 03 Aug 2016 Created to provide configuration data denoting the set of assets
//                   and their database (world state) prefixes, the set of pure events, 
//                   and any other useful runtime object that crops up. 

package main

import (
	"encoding/json"
    "errors"
)

var eventMaps map[string]interface{}
var schema map[string]interface{}

// Translation table for event names and prefixes. Includes isAsset property for
// convenience and performance.
// Dictionary:
//    eventPrefix: the prefix used to store the event itself (if needed)
//    assetPrefix: the prefix used to store the asset the event is *about*
//    isAsset:     this event is also an asset (i.e. a partial state)
//    assetName:   the name of the asset this event is *about*
const eventNamePrefixMaps string =
    `{
        "eventNameConfig": {
            "airline": {
                "eventPrefix": "AL",
                "assetPrefix": "AL",
                "isAsset": true,
                "assetName": "airline"
            },
            "aircraft": {
                "eventPrefix": "AC",
                "assetPrefix": "AC",
                "isAsset": true,
                "assetName": "aircraft"
            },
            "assembly": {
                "eventPrefix": "AS",
                "assetPrefix": "AS",
                "isAsset": true,
                "assetName": "assembly"
            },
            "flight": {
                "eventPrefix": "FL",
                "assetPrefix": "AC",
                "isAsset": false,
                "assetName": "aircraft"
            },
            "inspection": {
                "eventPrefix": "IN",
                "assetPrefix": "AS",
                "isAsset": false,
                "assetName": "assembly"
            },
            "analyticAdjustment": {
                "eventPrefix": "AA",
                "assetPrefix": "AS",
                "isAsset": false,
                "assetName": "assembly"
            }
        },
        "prefixToEventName": {
            "AL": { "eventName": "airline" },
            "AC": { "eventName": "aircraft" },
            "AS": { "eventName": "assembly" },
            "FL": { "eventName": "flight" },
            "IN": { "eventName": "inspection" },
            "AA": { "eventName": "analyticAdjustment" }
        }
    }`

func configInit() error {
    // schemas is the var created by the "go generate" command in schemas.go
    err := json.Unmarshal([]byte(schemas), &schema)
    if err != nil {
        // schema has syntax error
        log.Criticalf("The generated schema failed to unmarshal: %s\n", err)
        return err
    }
    //log.Debugf("GENERATED SCHEMA: %#v\n", schema)

    // eventMap is used to move quickly between name and prefix
    err = json.Unmarshal([]byte(eventNamePrefixMaps), &eventMaps)
    if err != nil {
        // eventMap has syntax error
        log.Criticalf("eventNamePrefixMaps failed to unmarshal: %s\n", err)
        return err
    }
    //log.Debugf("Conversion maps: %#v\n", eventMaps)
    
    // perform assertions to provide a basic alignment check of schema and config
    err = checkSchema("objectModelSchemas.airlineEvent"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.aircraftEvent"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.assemblyEvent"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.airlineState"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.aircraftState"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.assemblyState"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.flightEvent"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.inspectionEvent"); if err != nil {return err}
    err = checkSchema("objectModelSchemas.analyticAdjustmentEvent"); if err != nil {return err}

    // uncomment the following line if you are curious how alignement errors look in the logs
    //err = checkSchema("objectModelSchemas.THIS_SHOULD_FAIL"); if err != nil {return err}

    // now print the list of assets and events into the logs
    evIF, found := getObject(eventMaps, "eventNameConfig")
    if found {
        evMap, found := evIF.(map[string]interface{})
        if found {
            for k := range evMap {
                if isAssetName(k) {
                    log.Infof("%s is asset\n", k)
                } else {
                    log.Infof("%s is event\n", k)
                }
            }
        } else {
            log.Critical("eventNameConfig map is not map of interface{}")
        }
    } else {
        log.Critical("eventNameConfig map is corrupted")
    }
    return nil
}

// a bit of self-checking code
func checkSchema (checkit string) error {
    _, found := getObject(schema, checkit)
    if !found {
        // schema missing the element
        log.Criticalf("The generated schema does not contain: %s\n", checkit)
        return errors.New("The generated schema does not contain: " + checkit)
    }
    return nil
}

func isAssetName (checkit string) bool {
    b, found := getObjectAsBoolean(eventMaps, "eventNameConfig." + checkit + ".isAsset")
    if found {return b}
    return false 
}

func isAssetPrefix (checkit string) bool {
    n, found := getObjectAsString(eventMaps, "prefixToEventName." + checkit + ".eventName")
    if found {
        b, found := getObjectAsBoolean(eventMaps, "eventNameConfig." + n + ".isAsset")
        if found {
            return b
        }
    }
    return false 
}

func isEventName (checkit string) bool {
    _, found := getObject(eventMaps, "eventNameConfig." + checkit)
    return found 
}

func isEventPrefix (checkit string) bool {
    _, found := getObject(eventMaps, "prefixToEventName." + checkit)
    return found 
}

// An event can be a partial state, in which case the event name is also the 
// asset name. Or an event can be a pure event that is *about* an asset, and thus
// we have to go to the contract config to get the assetName. 
func eventNameToAssetName (eventName string) (string, error) {
    p, found := getObjectAsString(eventMaps, "eventNameConfig." + eventName + ".assetName")
    if found {
        return p, nil
    }
    err := errors.New("eventNameToAssetName: mapping event name to asset name failed for eventName = " + eventName)
    log.Error(err)
    return "", err 
}

func eventNameToEventPrefix (eventName string) (string, error) {
    p, found := getObjectAsString(eventMaps, "eventNameConfig." + eventName + ".eventPrefix")
    if found {
        return p, nil
    }
    err := errors.New("eventNameToEventPrefix: mapping event name to event prefix failed for eventName = " + eventName)
    log.Error(err)
    return "", err 
}

func eventNameToAssetPrefix (eventName string) (string, error) {
    p, found := getObjectAsString(eventMaps, "eventNameConfig." + eventName + ".assetPrefix")
    if found {
        return p, nil
    }
    err := errors.New("eventNameToAssetPrefix: mapping event name to asset prefix failed for eventName = " + eventName)
    log.Error(err)
    return "", err 
}

func eventPrefixToEventName (eventPrefix string) (string, error) {
    n, found := getObjectAsString(eventMaps, "prefixToEventName." + eventPrefix + ".eventName")
    if found { 
        return n, nil 
    }
    err := errors.New("eventPrefixToEventName: mapping event prefix to event name failed for eventPrefix = " + eventPrefix)
    log.Error(err)
    return "", err 
}

func assetIDToInternal (eventName string, assetID string) (string, error) {
    p, found := getObjectAsString(eventMaps, "nameToPrefix." + eventName + ".prefix")
    if found {
        return p + assetID, nil
    }
    err := errors.New("assetIDToInternal: mapping asset name to prefix failed for eventName = " + eventName + " and assetID = " + assetID)
    log.Error(err)
    return "", err
}

func assetIDToExternal (assetID string) (string, error) {
    _, found := getObject(eventMaps, "prefixToName." + assetID[0:2])
    if found {
        return assetID[2:], nil
    }
    err := errors.New("assetIDToEnternal: asset prefix not present in prefixToName: " + assetID)
    log.Error(err)
    return "", err
}

