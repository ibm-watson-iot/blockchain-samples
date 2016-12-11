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

// TaggedFilter is a complete filter for a state, inside a "filter" object"
type TaggedFilter struct {
	Filter StateFilter `json:"filter"`
}

var emptyStateFilter = StateFilter{"", make([]QPropNV, 0)}
var emptyTaggedFilter = TaggedFilter{StateFilter{"", make([]QPropNV, 0)}}

// Filter returns true if the filter's conditions are all met
func (a *Asset) Filter(filter StateFilter) bool {
	if len(filter.Select) == 0 {
		return true
	}
	switch filter.Match {
	case "n/a", "":
		return true
	case "all":
		return a.matchAll(filter)
	case "any":
		return a.matchAny(filter)
	case "none":
		return a.matchNone(filter)
	default:
		log.Noticef("filterObject has unknown matchType in filter: %+v", filter)
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
		if strings.TrimSuffix(json, ",omitempty") == strings.TrimSuffix(p, ".") {
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
func getUnmarshalledStateFilter(args []string) (StateFilter, error) {
	var filter StateFilter
	var err error

	if len(args) < 1 {
		// perfectly normal to not have a filter
		return emptyStateFilter, nil
	}

	filter, err = getCanonicalFilterFromEventIn(args)
	if err == nil && filter.Match != "" && filter.Match != "n/a" && len(filter.Select) > 0 {
		return filter, nil
	}
	filter, err = getMapFormatFilterFromEventIn(args)
	if err == nil && filter.Match != "" && filter.Match != "n/a" && len(filter.Select) > 0 {
		return filter, nil
	}
	return emptyStateFilter, nil
}

func getCanonicalFilterFromEventIn(args []string) (StateFilter, error) {
	var filter StateFilter
	var taggedFilter TaggedFilter

	fBytes := []byte(args[0])
	errtagged := json.Unmarshal(fBytes, &taggedFilter)
	if errtagged == nil && taggedFilter.Filter.Match != "" && taggedFilter.Filter.Match != "n/a" {
		return taggedFilter.Filter, nil
	}
	erruntagged := json.Unmarshal(fBytes, &filter)
	if erruntagged == nil && filter.Match != "" && filter.Match != "n/a" {
		return filter, nil
	}
	// log.Debugf("getCanonicalFilterFromEventIn failed to unmarshal %+v\n        tagged error: %s,\n        untagged error: %s\n", args[0], errtagged, erruntagged)
	return emptyStateFilter, nil
}

func getMapFormatFilterFromEventIn(args []string) (StateFilter, error) {
	var filter = StateFilter{"", make([]QPropNV, 0)}
	var f interface{}
	var err error

	fBytes := []byte(args[0])
	err = json.Unmarshal(fBytes, &f)
	if err != nil {
		log.Warningf("getMapFormatFilterFromEventIn failed to unmarshal incoming argument %s, error: %s\n", args[0], err)
		return emptyStateFilter, err
	}

	amap, found := AsMap(f)
	if !found {
		log.Warningf("getMapFormatFilterFromEventIn args: %s is not map shaped\n", args[0])
		return emptyStateFilter, err
	}

	fobj, found := GetObjectAsMap(&amap, "filter")
	if !found {
		// we'll try untagged format then
		fobj = amap
	}

	m, mfound := GetObjectAsString(&fobj, "match")
	sel, selfound := GetObjectAsMap(&fobj, "select")
	if !mfound {
		if selfound {
			log.Warningf("getMapFormatFilterFromEventIn incorrect filter format 'match' found: %t 'select' found %t\n", mfound, selfound)
			return emptyStateFilter, err
		}
	} else {
		if !selfound {
			log.Warningf("getMapFormatFilterFromEventIn incorrect filter format 'match' found: %t 'select' found %t\n", mfound, selfound)
			return emptyStateFilter, err
		}
	}

	filter.Match = m
	qprops := make([]QPropNV, 0)
	for _, e := range sel {
		emap, found := AsMap(e)
		if !found {
			log.Warningf("getMapFormatFilterFromEventIn prop:value not a map shape: %+v\n", e)
			return emptyStateFilter, nil
		}
		k, kfound := GetObjectAsString(&emap, "qprop")
		v, vfound := GetObjectAsString(&emap, "value")
		if !kfound || !vfound {
			log.Warningf("getMapFormatFilterFromEventIn prop or value not found: prop %t value %t\n", kfound, vfound)
			return emptyStateFilter, err
		}
		qprops = append(qprops, QPropNV{k, v})
	}
	filter.Select = qprops

	// log.Debugf("getMapFormatFilterFromEventIn returning filter %+v\n", filter)
	return filter, nil
}
