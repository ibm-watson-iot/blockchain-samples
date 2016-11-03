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

//
//  This module is mainly about storage of asset data. In order to properly partition
//  k:v pairs in the database, we need a prefix for the keys based on asset. Prepending
//  the prefix to the assetID when storing and referencing makes it easy to later crawl
//  an iterator over the keyset for queries.
//

// v1 KL 03 Aug 2016 Created to provide configuration data denoting the set of assets
//                   and their database (world state) prefixes, the set of pure events,
//                   and any other useful runtime object that crops up.
// KL 02 Nov 2016 add to package ctconfig

package ctconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctstate"
	"github.com/op/go-logging"
)

// CONFIGSTATEKEY is used to store the dynamic contract config structure.
const CONFIGSTATEKEY = "configStateKey"

var eventMaps = initEventMaps()
var schema = map[string]interface{}(nil)

// DynamicContractConfig is a struct that holds contract configurations that can be
// via application.
type DynamicContractConfig struct {
	ACheckThreshold float64 `json:"aCheckThreshold"`
	BCheckThreshold float64 `json:"bCheckThreshold"`
}

// Logger for the ctconfig package
var log = logging.MustGetLogger("conf")

var defaultDynamicConfig = DynamicContractConfig{2, 4}

// Translation table for event names and prefixes. Includes isAsset property for
// convenience and performance.
// Dictionary:
//    eventPrefix: the prefix used to store the event itself (if needed)
//    assetPrefix: the prefix used to store the asset the event is *about*
//    isAsset:     this event is also an asset (i.e. a partial state)
//    assetName:   the name of the asset this event is *about*
const eventNamePrefixMaps string = `{
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
            },
            "maintenance": {
                "eventPrefix": "MA",
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
            "AA": { "eventName": "analyticAdjustment" },
            "MA": { "eventName": "maintenance" }
        }
    }`

// -----------------------------------------------------------------------------
// static config initialization
// -----------------------------------------------------------------------------

func initEventMaps() map[string]interface{} {
	var myeventMaps map[string]interface{}
	// eventMap is used to move quickly between name and prefix
	err := json.Unmarshal([]byte(eventNamePrefixMaps), &myeventMaps)
	if err != nil {
		// eventMap has syntax error
		log.Criticalf("eventNamePrefixMaps failed to unmarshal: %s\n", err)
		return nil
	}
	//log.Debugf("Conversion maps: %#v\n", eventMaps)
	return myeventMaps
}

// InitSchema takes the passed in schemas string, which was created by running
// go generate to process the payload json schema, and unmarshalling it into a
// usable object (map)
func InitSchema(schemas string) map[string]interface{} {
	var myschema map[string]interface{}
	// schemas is the var created by the "go generate" command in schemas.go
	err := json.Unmarshal([]byte(schemas), &myschema)
	if err != nil {
		// schema has syntax error
		log.Criticalf("The generated schema failed to unmarshal: %s\n", err)
		return nil
	}
	//log.Debugf("GENERATED SCHEMA: %#v\n", schema)
	return myschema
}

// -----------------------------------------------------------------------------
// dynamic config initialization
// -----------------------------------------------------------------------------

// DynamicConfigInit initializes the dynamic (as in mutable) contract configuration
// for thresholds etc.
func DynamicConfigInit(stub *shim.ChaincodeStub) error {

	if err := verifySchema(); err != nil {
		return err
	}

	configbytes, err := stub.GetState(CONFIGSTATEKEY)
	if err == nil && len(configbytes) > 0 {
		// already initialized, this is a reboot etc
		return nil
	}

	config := defaultDynamicConfig
	if err != nil {
		err = fmt.Errorf("configInit cannot get default config: %s", err.Error())
		log.Error(err)
		return err
	}

	err = putDynamicConfigToLedger(stub, config)
	if err != nil {
		err = fmt.Errorf("configInit cannot put config: %s", err.Error())
		log.Error(err)
		return err
	}

	return nil
}

// GetDynamicConfigFromLedger retrieves the current dynamic config from world state
func GetDynamicConfigFromLedger(stub *shim.ChaincodeStub) (DynamicContractConfig, error) {
	// config must exist as it was created during Init
	config := defaultDynamicConfig
	configbytes, err := stub.GetState(CONFIGSTATEKEY)
	if err != nil || len(configbytes) == 0 {
		err := fmt.Errorf("getDynamicConfigFromLedger cannot get dynamic config from the ledger: %s", err.Error())
		log.Error(err)
		return config, err
	}
	err = json.Unmarshal(configbytes, &config)
	if err != nil {
		err := fmt.Errorf("getDynamicConfigFromLedger cannot unmarshal dynamic config: %s", err.Error())
		log.Error(err)
		return config, err
	}
	return config, nil
}

func putDynamicConfigToLedger(stub *shim.ChaincodeStub, config DynamicContractConfig) error {
	configBytes, err := json.Marshal(&config)
	if err != nil {
		err := fmt.Errorf("putDynamicConfigToLedger cannot marshall the dynamic config: %s", err.Error())
		log.Error(err)
		return err
	}
	err = stub.PutState(CONFIGSTATEKEY, configBytes)
	if err != nil {
		err := fmt.Errorf("putDynamicConfigToLedger cannot put the dynamic config to the ledger: %s", err.Error())
		log.Error(err)
		return err
	}
	return nil
}

// UpdateContractConfig updates world state for the contract configuration
func UpdateContractConfig(stub *shim.ChaincodeStub, args []string) error {
	var config DynamicContractConfig
	var err error
	if len(args) != 1 {
		err = errors.New("updateDynamicConfig: Incorrect number of arguments. Expecting a JSON encoded dynamic config")
		log.Error(err)
		return err
	}

	config, err = GetDynamicConfigFromLedger(stub)
	if err != nil {
		err = fmt.Errorf("updateDynamicConfig: failed to get dynamic config from ledger: %s", err)
		log.Error(err)
		return err
	}

	err = json.Unmarshal([]byte(args[0]), &config)
	if err != nil {
		err = fmt.Errorf("updateDynamicConfig failed to unmarshal arg: %s", err)
		log.Error(err)
		return err
	}

	err = putDynamicConfigToLedger(stub, config)
	if err != nil {
		err = fmt.Errorf("updateDynamicConfig failed to put config ro ledger: %s", err)
		log.Error(err)
		return err
	}

	return nil
}

// ReadContractConfig returns the latest dynamic contract configuration
func ReadContractConfig(stub *shim.ChaincodeStub) ([]byte, error) {
	configbytes, err := stub.GetState(CONFIGSTATEKEY)
	if err != nil || len(configbytes) == 0 {
		err := fmt.Errorf("readDynamicConfig cannot get dynamic config from the ledger: %s", err.Error())
		log.Error(err)
		return nil, err
	}
	return configbytes, nil
}

// -----------------------------------------------------------------------------
// static config accessors
// -----------------------------------------------------------------------------

// IsAssetName returns true if the incoming name is an asset
func IsAssetName(checkit string) bool {
	b, found := st.GetObjectAsBoolean(eventMaps, "eventNameConfig."+checkit+".isAsset")
	if found {
		return b
	}
	return false
}

// IsAssetPrefix returns true if the incoming string is an asset prefix
func IsAssetPrefix(checkit string) bool {
	n, found := st.GetObjectAsString(eventMaps, "prefixToEventName."+checkit+".eventName")
	if found {
		b, found := st.GetObjectAsBoolean(eventMaps, "eventNameConfig."+n+".isAsset")
		if found {
			return b
		}
	}
	return false
}

// IsEventPrefix returns true if the incoming string is an event prefix
func isEventPrefix(checkit string) bool {
	_, found := st.GetObject(eventMaps, "eventNameConfig."+checkit)
	return found
}

// IsEventName returns true if the incoming string is an event name
func IsEventName(checkit string) bool {
	_, found := st.GetObject(eventMaps, "prefixToEventName."+checkit)
	return found
}

// EventNameToAssetName returns the asset name of the primary targeted asset
func EventNameToAssetName(eventName string) (string, error) {
	p, found := st.GetObjectAsString(eventMaps, "eventNameConfig."+eventName+".assetName")
	if found {
		return p, nil
	}
	err := errors.New("eventNameToAssetName: mapping event name to asset name failed for eventName = " + eventName)
	log.Error(err)
	return "", err
}

// EventNameToEventPrefix returns the prefix associated with the event name
func EventNameToEventPrefix(eventName string) (string, error) {
	p, found := st.GetObjectAsString(eventMaps, "eventNameConfig."+eventName+".eventPrefix")
	if found {
		return p, nil
	}
	err := errors.New("eventNameToEventPrefix: mapping event name to event prefix failed for eventName = " + eventName)
	log.Error(err)
	return "", err
}

// EventNameToAssetPrefix returns the prefix associated with the primary targeted asset
func EventNameToAssetPrefix(eventName string) (string, error) {
	p, found := st.GetObjectAsString(eventMaps, "eventNameConfig."+eventName+".assetPrefix")
	if found {
		return p, nil
	}
	err := errors.New("eventNameToAssetPrefix: mapping event name to asset prefix failed for eventName = " + eventName)
	log.Error(err)
	return "", err
}

// EventPrefixToEventName returns the name associated with the event prefix
func EventPrefixToEventName(eventPrefix string) (string, error) {
	n, found := st.GetObjectAsString(eventMaps, "prefixToEventName."+eventPrefix+".eventName")
	if found {
		return n, nil
	}
	err := errors.New("eventPrefixToEventName: mapping event prefix to event name failed for eventPrefix = " + eventPrefix)
	log.Error(err)
	return "", err
}

// AssetIDToInternal returns the "prefixed" assetID based on its name
func AssetIDToInternal(eventName string, assetID string) (string, error) {
	p, found := st.GetObjectAsString(eventMaps, "eventNameConfig."+eventName+".assetPrefix")
	if found {
		return p + assetID, nil
	}
	err := errors.New("assetIDToInternal: mapping asset name to prefix failed for eventName = " + eventName + " and assetID = " + assetID)
	log.Error(err)
	return "", err
}

// AssetIDToExternal returns the "non-prefixed" assetID based on its name
func AssetIDToExternal(assetID string) (string, error) {
	log.Debugf("assetIDToExternal ASSETID: %s", assetID)
	_, found := st.GetObject(eventMaps, "prefixToEventName."+assetID[0:2])
	if found {
		return assetID[2:], nil
	}
	err := errors.New("assetIDToEnternal: asset prefix not present in prefixToName: " + assetID)
	log.Error(err)
	return "", err
}

// -----------------------------------------------------------------------------
// schema verification
// -----------------------------------------------------------------------------

func verifySchema() error {
	// perform assertions to provide a basic alignment check of schema and config
	err := checkSchema("objectModelSchemas.airlineEvent")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.aircraftEvent")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.assemblyEvent")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.airlineState")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.aircraftState")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.assemblyState")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.flightEvent")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.inspectionEvent")
	if err != nil {
		return err
	}
	err = checkSchema("objectModelSchemas.analyticAdjustmentEvent")
	if err != nil {
		return err
	}

	// uncomment the following line if you are curious how alignement errors look in the logs
	//err = checkSchema("objectModelSchemas.THIS_SHOULD_FAIL"); if err != nil {return err}

	// now print the list of assets and events into the logs
	evIF, found := st.GetObject(eventMaps, "eventNameConfig")
	if found {
		evMap, found := evIF.(map[string]interface{})
		if found {
			for k := range evMap {
				if IsAssetName(k) {
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
func checkSchema(checkit string) error {
	_, found := st.GetObject(schema, checkit)
	if !found {
		// schema missing the element
		log.Criticalf("The generated schema does not contain: %s\n", checkit)
		return errors.New("The generated schema does not contain: " + checkit)
	}
	return nil
}
