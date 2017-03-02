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
// KL 02 Nov 2016 add to package ctstate
// ************************************

package iotcontractplatform

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

var samplesStartLine = 36

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
    "temperature": 2.2,
    "integer": 2,
    "bool": true,
	"sarr": ["a","b"],
	"aa" : {
		"bb" : {
			"cc" : "d"
		}
	}
}`

func printUnmarshalError(js string, err interface{}) {
	syntax, ok := err.(*json.SyntaxError)
	if !ok {
		// fmt.Println("*********** ERR trying to get syntax error location **************\n", err)
		return
	}

	start, end := strings.LastIndex(js[:syntax.Offset], "\n")+1, len(js)
	if idx := strings.Index(js[start:], "\n"); idx >= 0 {
		end = start + idx
	}

	line, pos := strings.Count(js[:start], "\n"), int(syntax.Offset)-start-1
	// note, the offset here is the line number in this file
	// of the test samples json string definition (it happens to work out)
	fmt.Printf("Error in line %d: %s \n", line+samplesStartLine, err)
	fmt.Printf("%s\n%s^\n\n", js[start:end], strings.Repeat(" ", pos))
}

func getTestObjects(t *testing.T) map[string]interface{} {
	var o interface{}
	err := json.Unmarshal([]byte(testsamples), &o)
	if err != nil {
		printUnmarshalError(testsamples, err)
		t.Fatalf("unmarshal test samples failed: %s", err)
	} else {
		omap, found := o.(map[string]interface{})
		if found {
			return omap
		}
		t.Fatalf("test samples not map shape, is: %s", reflect.TypeOf(o))
	}
	return make(map[string]interface{})
}

func getTestParms(t *testing.T) map[string]interface{} {
	var o interface{}
	err := json.Unmarshal([]byte(testparm1), &o)
	if err != nil {
		printUnmarshalError(testsamples, err)
		t.Fatalf("unmarshal test samples failed: %s", err)
	}
	omap, _ := AsMap(o)
	return omap
}

func TestContains(t *testing.T) {
	t.Log("Enter TestContains")
	o := getTestObjects(t)
	ev1, found := GetObject(&o, "event1.extension.arr")
	if !found {
		t.Fatal("event1.extension.arr not found")
	}
	if !Contains(ev1, "s2") {
		t.Fatal("event1.extension.arr should contain s2")
	}
	if Contains(ev1, "s6") {
		t.Fatal("event1.extension.arr should not contain s6")
	}
	ev2, found := GetObject(&o, "event2.extension.arrint")
	if !found {
		t.Fatal("event2.extension.arrint not found")
	}
	// for the next 2, remember that JSON unmarshals numbers as float64
	if !Contains(ev2, float64(2)) {
		t.Fatal("event2.extension.arr should contain 2")
	}
	if Contains(ev2, float64(3)) {
		t.Fatal("event2.extension.arr should not contain 3")
	}
	ev3, found := GetObject(&o, "event3.extension.arr")
	if !found {
		t.Fatal("event3.extension.arr not found")
	}
	if Contains(ev3, "s2") {
		t.Fatal("event2.extension.arr should not contain s2")
	}
}

func TestDeepMerge(t *testing.T) {
	t.Log("Enter TestDeepMerge")
	o := getTestObjects(t)
	ev1, found := GetObjectAsMap(&o, "event1")
	// fmt.Printf("*** Event1: %s\n", PrettyPrint(ev1))
	if !found {
		t.Fatal("event1 not found")
	}
	ev2, found := GetObjectAsMap(&o, "event2")
	if !found {
		t.Fatal("event2 not found")
	}
	// fmt.Printf("*** Event2: %s\n", PrettyPrint(ev2))
	ev3, found := GetObjectAsMap(&o, "event3")
	if !found {
		t.Fatal("event3 not found")
	}
	// fmt.Printf("*** Event3: %s\n", PrettyPrint(ev3))
	_, found = GetObject(&ev1, "location.latitude")
	if found {
		t.Fatal("state1.location should not contain latitude")
	}
	_, found = GetObject(&ev1, "location.longitude")
	if found {
		t.Fatal("state1.location should not contain longitude")
	}
	state2 := DeepMergeMap(ev2, ev1)
	//fmt.Printf("*** State2 = ev2 + ev1: %s\n", PrettyPrint(state2))
	_, found = GetObject(&state2, "location.latitude")
	if !found {
		t.Fatal("state2.location should contain latitude")
	}
	_, found = GetObject(&state2, "location.longitude")
	if found {
		t.Fatal("state2.location should not contain longitude")
	}
	//fmt.Printf("*** ev3: %s\n", PrettyPrint(ev3))
	state3 := DeepMergeMap(ev3, state2)
	fmt.Printf("*** State3 = ev3 + state2: %s\n", PrettyPrint(state3))
	_, found = GetObject(&state3, "location.latitude")
	if !found {
		t.Fatal("state3.location should contain latitude")
	}
	_, found = GetObject(&state3, "location.longitude")
	if !found {
		t.Fatal("state3.location should contain longitude")
	}
}

func TestParms(t *testing.T) {
	// fmt.Println("Enter TestContains")
	o := getTestParms(t)
	_, found := GetObject(&o, "assetID")
	if !found {
		t.Fatal("assetID not found")
	}
}

func TestGetByType(t *testing.T) {
	// fmt.Println("Enter TestByType")
	o := getTestParms(t)
	_, found := GetObjectAsString(&o, "assetID")
	if !found {
		t.Fatal("typeof assetID should be string")
	}

	_, found = GetObjectAsStringArray(&o, "sarr")
	if !found {
		t.Fatalf("typeof sarr should be []string")
	}

	_, found = GetObjectAsNumber(&o, "temperature")
	if !found {
		t.Fatal("typeof temperature should be number")
	}

	_, found = GetObjectAsInteger(&o, "temperature")
	if !found {
		t.Fatal("type of temperature should be integer")
	}

	_, found = GetObjectAsInteger(&o, "integer")
	if !found {
		t.Fatal("typeof integer should be integer")
	}

	_, found = GetObjectAsMap(&o, "aa")
	if !found {
		t.Fatal("typeof aa should be map")
	}
}

func TestPutObject(t *testing.T) {
	// fmt.Println("Enter TestPutObject")
	o := getTestParms(t)

	// fmt.Printf("Object before: %+v\n\n", o)

	ok := PutObject(&o, "time", time.Now())
	if !ok {
		t.Fatal("could not put time")
	}
	tm, ok := GetObject(&o, "time")
	if !ok {
		t.Fatal("put time failed")
	}
	fmt.Printf("Time after retrieved: %+v\n\n", tm)

	ok = PutObject(&o, "anInt", 1)
	if !ok {
		t.Fatal("could not put anInt")
	}
	_, ok = GetObject(&o, "anInt")
	if !ok {
		t.Fatal("put anInt failed")
	}

	ok = PutObject(&o, "aFloat", 1.567)
	if !ok {
		t.Fatal("could not put aFloat")
	}
	_, ok = GetObject(&o, "aFloat")
	if !ok {
		t.Fatal("put aFloat failed")
	}

	ok = PutObject(&o, "maintenance.status", "inventory")
	if !ok {
		t.Fatal("could not put maintenance.status")
	}
	_, ok = GetObjectAsString(&o, "maintenance.status")
	if !ok {
		t.Fatal("put maintenance.status failed")
	}

	ok = PutObject(&o, "a.b.c.d.lastmaplevel.status", "installed")
	if !ok {
		t.Fatal("could not put a.b.c.d.lastmaplevel.status")
	}
	_, ok = GetObject(&o, "anInt")
	if !ok {
		t.Fatal("put a.b.c.d.lastmaplevel.status failed")
	}

	// fmt.Printf("Object after: %+v\n\n", o)

}

func TestRemoveObject(t *testing.T) {
	// fmt.Println("Enter TestRemoveObject")
	o := getTestParms(t)

	fmt.Printf("Object before: %+v\n\n", o)

	ok := RemoveObject(&o, "assetID")
	if !ok {
		t.Fatal("could not remove assetID")
	}

	ok = RemoveObject(&o, "carrier")
	if !ok {
		t.Fatal("could not remove carrier")
	}

	ok = RemoveObject(&o, "aa.bb.cc")
	if !ok {
		t.Fatal("could not remove aa.bb.cc")
	}

	// fmt.Printf("Object after removal of aa.bb.cc: %+v\n\n", o)

	ok = RemoveObject(&o, "aa")
	if !ok {
		t.Fatal("could not remove aa")
	}

	fmt.Printf("Object after: %+v\n\n", o)
}

func TestAsStringArray(t *testing.T) {
	// fmt.Println("Enter TestAsStringArray")

	_, ok := AsStringArray([]string{"a"})
	if !ok {
		t.Fatal("could not convert []string{'a'} to string array")
	}
	// fmt.Printf("TestAsStringArray: conversion of []string{'a'} created %#v\n", s)

	_, ok = AsStringArray([]int{2, 3, 4})
	if ok {
		t.Fatal("converted []int{2, 3, 4} to string array, how?")
	}

	_, ok = AsStringArray("astring")
	if !ok {
		t.Fatal("failed to convert 'astring' to string array")
	}
	// fmt.Printf("TestAsStringArray: conversion of 'astring' created %#v\n", s)

	_, ok = AsStringArray(`["a", "b", "c"]`)
	if !ok {
		t.Fatal("failed to convert JSON a,b,c to string array")
	}
	// fmt.Printf("TestAsStringArray: conversion of JSON array ['a', 'b', 'c'] created %#v\n", s)
}

func TestAddToStringArray(t *testing.T) {
	// fmt.Println("Enter TestAddToStringArray")
	o := getTestParms(t)

	// fmt.Printf("Object before: %+v\n\n", o)

	sarr, ok := GetObjectAsStringArray(&o, "sarr")
	if !ok {
		t.Fatal("could not get sarr as string array")
	}
	AddToStringArray([]string{"d", "b", "c"}, &sarr)
	// fmt.Printf("Addtostringarray TEST for equality: %+v === %+v\n\n", sarr, []string{"a", "b", "c", "d"})
	if !reflect.DeepEqual(sarr, []string{"a", "b", "c", "d"}) {
		t.Fatal("merge of [d,b,c] into [a,b] failed")
	}
	// fmt.Printf("Object d,b,c added: %+v\n\n", o)

	sarr, ok = GetObjectAsStringArray(&o, "sarr")
	if !ok {
		t.Fatal("could not get sarr as string array")
	}
	AddToStringArray([]string{"astring"}, &sarr)
	if !ok {
		t.Fatal("could not merge 'astring' into sarr")
	}
	// fmt.Printf("Object after: %+v\n\n", o)

	// fmt.Printf("Object after destruction by not checking for nil: %+v\n\n", o)

}

func TestRemoveFromStringArray(t *testing.T) {
	// fmt.Println("Enter TestRemoveFromStringArray")
	o := getTestParms(t)
	arr, ok := o["sarr"]
	if !ok {
		t.Fatal("sarr missing from test parms")
	}
	sarr, ok := AsStringArray(arr)
	if !ok {
		t.Fatal("sarr not an array of strings")
	}

	fmt.Printf("TestRemoveFromStringArray before: %+v\n\n", sarr)

	RemoveFromStringArray([]string{"d", "b", "c"}, &sarr)
	fmt.Printf("Addtostringarray TEST for equality: %+v === %+v\n\n", sarr, []string{"a"})
	if !reflect.DeepEqual(sarr, []string{"a"}) {
		t.Fatal("could not remove [d,b,c] from sarr")
	}

	RemoveFromStringArray([]string{"a"}, &sarr)
	if len(sarr) > 0 {
		t.Fatal("could not remove 'a' from sarr")
	}

	// fmt.Printf("Object after: %+v\n\n", o)

}
