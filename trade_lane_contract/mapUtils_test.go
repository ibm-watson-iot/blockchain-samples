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


// ************************************
// KL 27 Mar 2016 add testing for mapUtils as strict RPC is coming in
// ************************************

package main

import (
    "fmt"
    "testing"
    "strings"
	"encoding/json"
    "reflect"
)

var samplesStartLine int = 36

var testsamples = `
{
    "event1": {
        "assetID": "ASSET001",
        "carrier": "UPS",
        "extension": {
            "arr": ["s1", "s2", "s3"]
        },
        "location": {
        },
        "temperature": 123.456,
        "timestamp": "2016-03-17T01:51:23.51620144Z"
    },
    "event2": {
        "AssetID": "ASSET001",
        "carrier": "UPS",
        "extension": {
            "arrint": [1, 2]
        },
        "location": {
            "latitude": 123.456
        },
        "Temperature": 123.456,
        "timestamp": "2016-03-17T01:51:23.51620144Z"
    },
    "event3": {
        "assetid": "ASSET001",
        "carrier": "UPS",
        "extension": {
            "arr": []
        },
        "location": {
            "longitude": 123.456
        },
        "tEmperature": 123.456,
        "timestamp": "2016-03-17T01:51:23.51620144Z"
    }
}`

var testparm1 = `
{
    "assetID": "ASSET001",
    "carrier": "UPS",
    "temperature": 2
}`

func printUnmarshalError(js string, err interface{}) {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
        fmt.Println("*********** ERR trying to get syntax error location **************\n", err)
		return
	}
	
	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}
	
	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset) - start -1
	// note, the offset here is the line number in this file
    // of the test samples json string definition (it happens to work out)
	fmt.Printf("Error in line %d: %s \n", line + samplesStartLine, err)
	fmt.Printf("%s\n%s^\n\n", js[start:end], strings.Repeat(" ", pos))
}

func getTestObjects (t *testing.T) (map[string]interface{}) {
    var o interface{}
    err := json.Unmarshal([]byte(testsamples), &o)
    if err != nil { 
        printUnmarshalError(testsamples, err)
        t.Fatalf("unmarshal test samples failed: %s", err)
    } else {
        omap, found := o.(map[string]interface{})
        if found { 
            return omap
        } else {
            t.Fatalf("test samples not map shape, is: %s", reflect.TypeOf(o))
        }
    }
    return make(map[string]interface{})
}

func getTestParms (t *testing.T) (interface{}) {
    var o interface{}
    err := json.Unmarshal([]byte(testparm1), &o)
    if err != nil { 
        printUnmarshalError(testsamples, err)
        t.Fatalf("unmarshal test samples failed: %s", err)
    }
    return o
}

func TestContains(t *testing.T) {
    t.Log("Enter TestContains")
    o := getTestObjects(t)
    ev1, found := getObject(o, "event1.extension.arr")
    if !found {
        t.Fatal("event1.extension.arr not found")
    }
    if !contains(ev1, "s2") {
        t.Fatal("event1.extension.arr should contain s2")
    }
    if contains(ev1, "s6") {
        t.Fatal("event1.extension.arr should not contain s6")
    }
    ev2, found := getObject(o, "event2.extension.arrint")
    if !found {
        t.Fatal("event2.extension.arrint not found")
    }
    // for the next 2, remember that JSON unmarshals numbers as float64
    if !contains(ev2, float64(2)) {
        t.Fatal("event2.extension.arr should contain 2")
    }
    if contains(ev2, float64(3)) {
        t.Fatal("event2.extension.arr should not contain 3")
    }
    ev3, found := getObject(o, "event3.extension.arr")
    if !found {
        t.Fatal("event3.extension.arr not found")
    }
    if contains(ev3, "s2") {
        t.Fatal("event2.extension.arr should not contain s2")
    }
}

func TestDeepMerge(t *testing.T)  {
    t.Log("Enter TestDeepMerge")
    o := getTestObjects(t)
    ev1, found := getObject(o, "event1")
    if !found {
        t.Fatal("event1 not found")
    }
    ev2, found := getObject(o, "event2")
    if !found {
        t.Fatal("event2 not found")
    }
    ev3, found := getObject(o, "event3")
    if !found {
        t.Fatal("event3 not found")
    }
    state1 := ev1
    //fmt.Printf("*** State1: %s\n", prettyPrint(state1))
    _, found = getObject(state1, "location.latitude") 
    if  found {
        t.Fatal("state1.location should not contain latitude")
    }
    _, found = getObject(state1, "location.longitude")    
    if found {
        t.Fatal("state1.location should not contain longitude")
    }
    state2 := deepMerge(ev2, state1)
    //fmt.Printf("*** State2: %s\n", prettyPrint(state2))
    _, found = getObject(state1, "location.latitude") 
    if !found {
        t.Fatal("state2.location should contain latitude")
    }
    _, found = getObject(state1, "location.longitude")    
    if found {
        t.Fatal("state2.location should not contain longitude")
    }
    state3 := deepMerge(ev3, state2)
    //fmt.Printf("*** State3: %s\n", prettyPrint(state3))
    _, found = getObject(state3, "location.latitude") 
    if !found {
        t.Fatal("state2.location should contain latitude")
    }
    _, found = getObject(state3, "location.longitude")    
    if !found {
        t.Fatal("state3.location should contain longitude")
    }
}

func TestParms(t *testing.T)  {
    fmt.Println("Enter TestContains")
    o := getTestParms(t)
    aid1, found := getObject(o, "assetID")
    if !found {
        t.Fatal("assetID not found")
    }
    fmt.Println("AssetID: ", aid1, " TypeOF parms: ", reflect.TypeOf(o))
}

func TestArgsMap(t *testing.T)  {
    fmt.Println("Enter TestArgsMap")
    o := getTestParms(t)
    var a ArgsMap = ArgsMap(o.(map[string]interface{})) 
    aid1, found := getObject(a, "assetID")
    if !found {
        t.Fatal("assetID not found")
    }
    fmt.Println("AssetID: ", aid1, " TypeOF parms: ", reflect.TypeOf(a))
}
