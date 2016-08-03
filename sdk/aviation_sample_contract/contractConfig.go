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
    "fmt"
    "time"
    //"github.com/op/go-logging"
    "strings"
)

var assets map[string]map[string]string
var events map[string]map[string]string
var schema map[string]interface{}

// all asset classes should appear here
const assetConfig string =
       `{
            "airline": {
                "id": "AL"
            },
            "aircraft": {
                "ID": "AC"
            },
            "assembly": {
                "ID": "AS"
            }
        }`

// all pure events (non-assets) should appear here
const eventConfig string =
       `{
            "flight": {
                "id": "FL"
            },
            "inspection": {
                "ID": "IN"
            },
            "analyticAdjustment": {
                "ID": "AA"
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

    // assets is used by multi-asset world state etc
    err := json.Unmarshal([]byte(assetConfig), &assets)
    if err != nil {
        // assetConfig has syntax error
        log.Criticalf("assetConfig failed to unmarshal: %s\n", err)
        return err
    }

    //log.Debugf("GENERATED SCHEMA: %+v\n", schema)
    
    // perform assertions to provide a basic error check of schema and config
    err = checkSchema("airline"); if err != nil {return err}
    err = checkSchema("aircraft"); if err != nil {return err}
    err = checkSchema("assembly"); if err != nil {return err}
    err = checkSchema("flight"); if err != nil {return err}
    err = checkSchema("inspection"); if err != nil {return err}
    err = checkSchema("analyticAdjustment"); if err != nil {return err}
    err = checkSchema("testNotThere"); if err != nil {return err}
}

func checkSchema(string checkit) error {
    schemacheck, found := getObject(schema, checkit)
    if !found {
        // schema missing the element
        log.Criticalf("The generated schema does not contain: %s\n", checkit)
        return errors.New("The generated schema does not contain: " + checkit)
    }
}
