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

// v1 KL 10 Aug 2016 Add filters to queries to enable powerful searching
// v2 KL 02 Nov 2016 new package ctasset

package ctasset

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctstate"
)

// MatchType denotes how a filter should operate.
type MatchType int32

const (
	// MatchAll requires that every property in the filter be present and have
	// the same value
	MatchAll MatchType = 0
	// MatchAny requires that at least one property in the filter be present and have
	// the same value
	MatchAny MatchType = 1
	// MatchNone requires that every property in the filter either be present and have
	// a different value. or not be present
	MatchNone MatchType = 2
)

// MatchName is a map of ID to name
var MatchName = map[int]string{
	0: "matchall",
	1: "matchany",
	2: "matchnone",
}

// MatchValue is a map of name to ID
var MatchValue = map[string]int32{
	"matchall":  0,
	"matchany":  1,
	"matchnone": 2,
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

// StateFilter is an array of QPropNV
type StateFilter struct {
	MatchMode string    `json:"matchmode"`
	Entries   []QPropNV `json:"entries"`
}

var emptyFilter = StateFilter{"matchall", make([]QPropNV, 0)}

// Filter returns true if the filter's conditions are all met
func (a *Asset) Filter(filter StateFilter) bool {
	switch filter.MatchMode {
	case "ALL":
		return matchAll(a, filter)
	case "ANY":
		return matchAny(a, filter)
	case "NONE":
		return matchNone(a, filter)
	default:
		err := fmt.Errorf("filterObject has unknown matchType in filter: %+v", filter)
		log.Errorf(err.Error())
		return false
	}
}

func matchAll(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Entries {
		if !performOneMatch(a.State, f) {
			// must match all
			return false
		}
	}
	// success
	return true
}

func matchAny(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Entries {
		if performOneMatch(a.State, f) {
			// must match at least one
			return true
		}
	}
	// fail
	return false
}

func matchNone(a *Asset, filter StateFilter) bool {
	for _, f := range filter.Entries {
		if performOneMatch(a.State, f) {
			// must not match any
			return false
		}
	}
	// success, none matched
	return true
}

func performOneMatch(obj *map[string]interface{}, prop QPropNV) bool {
	o, found := st.GetObject(obj, prop.QProp)
	if found {
		switch t := o.(type) {
		case []interface{}:
			return st.Contains(o, prop.Value)
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

// Returns a map containing the JSON object represented by args[0]
func getUnmarshalledStateFilter(stub shim.ChaincodeStubInterface, caller string, args []string) StateFilter {
	var filter StateFilter
	var err error

	if len(args) != 1 {
		// perfectly normal to not have a filter
		return emptyFilter
	}

	fBytes := []byte(args[0])
	err = json.Unmarshal(fBytes, &filter)
	if err != nil {
		err = fmt.Errorf("%s failed to unmarshal filter: %s error: %s", caller, args[0], err)
		log.Errorf(err.Error())
		return emptyFilter
	}

	return filter
}
