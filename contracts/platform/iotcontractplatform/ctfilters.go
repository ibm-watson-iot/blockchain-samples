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
	"reflect"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// MatchType denotes how a filter should operate.
type MatchType int32

const (
	// MatchDisabled causes the filter to not execute
	MatchDisabled = 0
	// MatchAll requires that every property in the filter be present and have
	// the same value
	MatchAll MatchType = 1
	// MatchAny requires that at least one property in the filter be present and have
	// the same value
	MatchAny MatchType = 2
	// MatchNone requires that every property in the filter either be present and have
	// a different value. or not be present
	MatchNone MatchType = 3
)

// MatchName is a map of ID to name
var MatchName = map[int]string{
	0: "n/a",
	1: "all",
	2: "any",
	3: "none",
}

// MatchValue is a map of name to ID
var MatchValue = map[string]int32{
	"n/a":  0,
	"all":  1,
	"any":  2,
	"none": 3,
}

func (x MatchType) String() string {
	return MatchName[int(x)]
}

// QPropNV is a name : value pair to be matched
type QPropNV struct {
	QProp string `json:"qprop"`
	Value string `json:"value"`
}

// StateFilter is a complete filter for a state
type StateFilter struct {
	Match  string    `json:"match"`
	Select []QPropNV `json:"select"`
}

// Filter returns true if the filter's conditions are all met
func (a *Asset) Filter(filter StateFilter) bool {
	switch filter.Match {
	case "n/a":
		return true
	case "all":
		return a.matchAll(filter)
	case "any":
		return a.matchAny(filter)
	case "none":
		return a.matchNone(filter)
	default:
		err := fmt.Errorf("filterObject has unknown matchType in filter: %+v", filter)
		log.Notice(err)
		return true
	}
}

func (a *Asset) matchAll(filter StateFilter) bool {
	for _, f := range filter.Select {
		if !a.performOneMatch(f) {
			// must match all
			return false
		}
	}
	// success
	return true
}

func (a *Asset) matchAny(filter StateFilter) bool {
	for _, f := range filter.Select {
		if a.performOneMatch(f) {
			// must match at least one
			return true
		}
	}
	// fail
	return false
}

func (a *Asset) matchNone(filter StateFilter) bool {
	for _, f := range filter.Select {
		if a.performOneMatch(f) {
			// must not match any
			return false
		}
	}
	// success, none matched
	return true
}

func findJSONPropInStruct(p string, v reflect.Value) (reflect.Value, interface{}, reflect.Kind, bool) {
	dump("findJSONPropInStruct", p, nil, nil, nil)
	if v.Kind() != reflect.Struct {
		return reflect.Value{}, nil, reflect.Invalid, false
	}
	typear := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		json := typear.Field(i).Tag.Get("json")
		dump("findJSONPropInStruct loop", f, f.Kind(), p, json)
		if strings.TrimSuffix(json, ",omitempty") == p {
			return f, f.Interface(), f.Kind(), true
		}
	}
	return reflect.Value{}, nil, reflect.Invalid, false
}

func dump(s string, i interface{}, j interface{}, k interface{}, l interface{}) {
	fmt.Printf("%s: first: %+v || second: %+v || third: %+v || fourth: %+v\n", s, i, j, k, l)
}

func (a *Asset) performOneMatch(prop QPropNV) bool {
	dump("performOneMatch", a, prop, nil, nil)
	var levels []string
	var found = false
	var kind reflect.Kind
	var v reflect.Value
	var o interface{}
	if prop.QProp == "" {
		return false
	}
	levels = strings.SplitAfterN(prop.QProp, ".", 2)
	ar := reflect.ValueOf(a).Elem()
	v, o, kind, found = findJSONPropInStruct(strings.TrimSuffix(levels[0], ","), ar)
	dump("JSON prop in struct returned", v, o, kind, found)

	if found {
		if len(levels) == 2 {
			omap, found := o.(*map[string]interface{})
			if found {
				o, found = GetObject(omap, levels[1])
				if !found {
					return false
				}
			} else if kind == reflect.Struct {
				v, o, kind, found = findJSONPropInStruct(levels[1], v)
				if !found {
					return false
				}
			} else {
				return false
			}
		} else if kind == reflect.Slice {
			return Contains(o, prop.Value)
		}
		// ** at this point, we have a leaf node interface{} value "o" to compare
		switch t := o.(type) {
		case []interface{}:
			return Contains(o, prop.Value)
		case string:
			return o.(string) == prop.Value
		case float64:
			f, err := strconv.ParseFloat(prop.Value, 64)
			if err == nil {
				return o.(float64) == f
			}
			err = fmt.Errorf("Cannot convert %s to float64 in filter when comparing to object %s %+v", prop.Value, prop.QProp, a)
			log.Error(err)
			return false
		case int:
			i, err := strconv.Atoi(prop.Value)
			if err == nil {
				return o.(int) == i
			}
			err = fmt.Errorf("Cannot convert %s to int in filter when comparing to object %s %+v", prop.Value, prop.QProp, a)
			log.Error(err)
			return false
		case bool:
			if b, err := strconv.ParseBool(prop.Value); err == nil {
				return b == o.(bool)
			}
			err := fmt.Errorf("Cannot convert %s to bool in filter when comparing to object %s %+v", prop.Value, prop.QProp, a)
			log.Error(err)
			return false
		default:
			err := fmt.Errorf("Unexpected property to compare type: %T %s", prop.Value, t)
			log.Error(err)
			return false
		}
	}
	return false
}

// Returns a filter found in the json object in args[0]
func getUnmarshalledStateFilter(stub shim.ChaincodeStubInterface, args []string) (StateFilter, error) {
	var filter = StateFilter{"", make([]QPropNV, 0)}
	var f interface{}
	var err error

	dump("getUnmarshalledStateFilter args", args, nil, nil, nil)

	if len(args) != 1 {
		// perfectly normal to not have a filter
		return filter, nil
	}

	fBytes := []byte(args[0])
	err = json.Unmarshal(fBytes, &f)
	if err != nil {
		err = fmt.Errorf("getUnmarshalledStateFilter failed to unmarshal %s, error: %s", args[0], err)
		log.Error(err)
		return filter, err
	}

	amap, found := AsMap(f)
	if !found {
		err = fmt.Errorf("getUnmarshalledStateFilter %s is not map shaped", args[0])
		log.Error(err)
		return filter, err
	}

	fobj, found := GetObjectAsMap(&amap, "filter")
	if !found {
		return filter, nil
	}

	m, mfound := GetObjectAsString(&fobj, "match")
	sel, selfound := GetObjectAsMap(&fobj, "select")
	if !mfound || !selfound {
		err = fmt.Errorf("getUnmarshalledStateFilter matchfound: %t selectfound %t", mfound, selfound)
		log.Error(err)
		return filter, err
	}

	filter.Match = m
	qprops := make([]QPropNV, 0)
	for _, e := range sel {
		emap, found := AsMap(e)
		if !found {
			return filter, nil
		}
		k, kfound := GetObjectAsString(&emap, "qprop")
		v, vfound := GetObjectAsString(&emap, "value")
		if !kfound || !vfound {
			err = fmt.Errorf("getUnmarshalledStateFilter matchfound: %t selectfound %t", kfound, vfound)
			log.Error(err)
			return filter, err
		}
		qprops = append(qprops, QPropNV{k, v})
	}
	filter.Select = qprops

	return filter, nil
}
