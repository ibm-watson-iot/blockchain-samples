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
	"strconv"

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

// QPropNV is a single search entry to be matched
// - the qualifiedProperty field denotes a path to a leaf node in the object
// - the value property denotes the value to match against
type QPropNV struct {
	QProp string `json:"qprop"`
	Value string `json:"value"`
}

// FilterGroup is a matchmode with a list of K:V pairs
type FilterGroup struct {
	Match  string    `json:"match"`
	Select []QPropNV `json:"select"`
}

// StateFilter is a complete filter for a state
type StateFilter struct {
	Filter FilterGroup `json:"filter"`
}

var emptyFilter = StateFilter{
	FilterGroup{
		"n/a",
		make([]QPropNV, 0),
	},
}

// Filter returns true if the filter's conditions are all met
func (a *Asset) Filter(filter StateFilter) bool {
	if len(filter.Filter.Select) == 0 {
		return true
	}
	switch filter.Filter.Match {
	case "n/a":
		return true
	case "all":
		return matchAll(a, filter)
	case "any":
		return matchAny(a, filter)
	case "none":
		return matchNone(a, filter)
	default:
		err := fmt.Errorf("filterObject has unknown matchType in filter: %+v", filter)
		log.Errorf(err.Error())
		return false
	}
}

func matchAll(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Filter.Select {
		if !performOneMatch(a.State, f) {
			// must match all
			return false
		}
	}
	// success
	return true
}

func matchAny(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Filter.Select {
		if performOneMatch(a.State, f) {
			// must match at least one
			return true
		}
	}
	// fail
	return false
}

func matchNone(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Filter.Select {
		if performOneMatch(a.State, f) {
			// must not match any
			return false
		}
	}
	// success, none matched
	return true
}

func performOneMatch(obj *map[string]interface{}, prop QPropNV) bool {
	o, found := GetObject(obj, prop.QProp)
	if found {
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
			err = fmt.Errorf("Cannot convert %s to float64 in filter when comparing to object %s %+v", prop.Value, prop.QProp, obj)
			log.Errorf(err.Error())
			return false
		case int:
			i, err := strconv.Atoi(prop.Value)
			if err == nil {
				return o.(int) == i
			}
			err = fmt.Errorf("Cannot convert %s to int in filter when comparing to object %s %+v", prop.Value, prop.QProp, obj)
			log.Errorf(err.Error())
			return false
		case bool:
			if b, err := strconv.ParseBool(prop.Value); err == nil {
				return b == o.(bool)
			}
			err := fmt.Errorf("Cannot convert %s to bool in filter when comparing to object %s %+v", prop.Value, prop.QProp, obj)
			log.Errorf(err.Error())
			return false
		default:
			err := fmt.Errorf("Unexpected property to compare type: %T %s", prop.Value, t)
			log.Errorf(err.Error())
			return false
		}
	}
	return false
}

// Returns a filter found in the json object in args[0]
func getUnmarshalledStateFilter(stub shim.ChaincodeStubInterface, args []string) (StateFilter, error) {
	var filter StateFilter
	var err error

	if len(args) != 1 {
		// perfectly normal to not have a filter
		return emptyFilter, nil
	}

	fBytes := []byte(args[0])
	err = json.Unmarshal(fBytes, &filter)
	if err != nil {
		err = fmt.Errorf("getUnmarshalledStateFilter failed to unmarshal %s as filter, error: %s", args[0], err)
		log.Error(err)
		return emptyFilter, err
	}

	return filter, nil
}
