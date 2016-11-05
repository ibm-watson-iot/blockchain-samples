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
// KL 14 Mar 2016 backport from 4.0
// KL 27 Mar 2016 support for loose JSON-RPC naming for 3.0.5
// KL 10 Aug 2016 added more flexibility in "GetObjectAsInteger"
// KL 02 Nov 2016 add to package ctstate
// ************************************

package ctstate

import (
	"encoding/json"
	"fmt"
	// "reflect"
	"sort"
	"strings"
)

// ArgsMap is an alias for maps acting as json objects, needed because
// map[string]interface{} cannot be a receiver
type ArgsMap map[string]interface{}

// AsMap does its best to interpret or cast the incoming generic to map[string]interface{}
func AsMap(obj interface{}) (toMap map[string]interface{}, ok bool) {
	var err error
	toMap, found := obj.(map[string]interface{})
	if found {
		return toMap, true
	}
	as, found := obj.(string)
	if found {
		var data interface{}
		err := json.Unmarshal([]byte(as), &data)
		if err == nil {
			return AsMap(interface{}(data))
		}
	}
	err = fmt.Errorf("AsMap: incoming type is %T and is not understood", obj)
	log.Errorf(err.Error())
	return nil, false
}

// AsStringArray does its best to interpret or cast to []string
func AsStringArray(obj interface{}) (toSarr []string, ok bool) {
	var err error
	// 1. array of interface{}, which should of course contain strings
	sa, ok := obj.([]interface{})
	if ok {
		for i, el := range sa {
			sel, ok := el.(string)
			if !ok {
				err = fmt.Errorf("AsStringArray: incoming element %d type is %T from array %#v and is not understood", i, el, obj)
				log.Errorf(err.Error())
				return nil, false
			}
			toSarr = append(toSarr, sel)
		}
		return toSarr, true
	}
	// 2. array of strings, nothing to do
	toSarr, ok = obj.([]string)
	if ok {
		return toSarr, true
	}
	// what about a string argument?
	as, ok := obj.(string)
	if ok {
		if len(as) > 0 && as[0] == '[' {
			// 3. encoded JSON array of strings, unmarshall and call recursively if successful
			var data interface{}
			err := json.Unmarshal([]byte(as), &data)
			if err == nil {
				return AsStringArray(interface{}(data))
			}
			log.Errorf(err.Error())
			return nil, false
		}
		// 4. a non-JSON string, just return that as an array
		return []string{as}, true
	}
	err = fmt.Errorf("AsStringArray: incoming type is %T and is not understood", obj)
	log.Errorf(err.Error())
	return nil, false
}

// GetObject finds an object by its qualified name, which looks like "location.latitude"
// as one example. Returns as interface{} to maintain generic handling
func GetObject(objIn map[string]interface{}, qname string) (interface{}, bool) {
	// return a copy of the selected object
	// handles full qualified name, starting at object's root
	searchObj := objIn
	s := strings.Split(qname, ".")
	// crawl the levels
	for i, v := range s {
		//fmt.Printf("**** FIND level [%d] %s\n", i, v)
		//fmt.Printf("**** FIND level [%d] %s in %+v\n", i, v, searchObj)
		if i+1 < len(s) {
			tmp, found := searchObj[v]
			//fmt.Printf("** tmp is %+v\n", tmp)
			if found {
				searchObj, found = tmp.(map[string]interface{})
				//fmt.Printf("** tmp->searchObj AS MAP is %+v\n", searchObj)
				if !found {
					searchObj, found = tmp.(ArgsMap)
					if !found {
						// log.Debugf("GetObject not map or ArgsMap shape: %s", v)
						//fmt.Printf("** tmp->searchObj AS ARGSMAP is %+v\n", searchObj)
					}
				}
			} else {
				// log.Debugf("GetObject cannot find level: %s in %s", v, qname)
				return nil, false
			}
		} else {
			returnObj, found := searchObj[v]
			if !found {
				// this debug statement is not useful normally as we must be able to
				// handle assetID as part of iot common and as parameter on its own
				// so we get false warnings on read functions, but do enable it if
				// having problems with deep nested structures
				// log.Debugf("GetObject cannot find final level: %s in %s", v, qname)
				return nil, false
			}
			//fmt.Printf("**** Found level [%d] %s\n", i, v)
			return returnObj, true
		}
	}
	return nil, false
}

// PutObject inserts an object by its qualified name, which looks like "location.latitude"
// as one example. Creates missing levels.
func PutObject(objIn map[string]interface{}, qname string, value interface{}) (map[string]interface{}, bool) {
	// overwrite the value of the selected object, create if necessary
	// handles full qualified name, starting at object's root
	searchObj, ok := AsMap(objIn)
	if !ok {
		log.Errorf("GetObject passed a non-map / non-ArgsMap: %#v", objIn)
		return objIn, false
	}
	s := strings.Split(qname, ".")
	// crawl the levels
	for i, v := range s {
		//fmt.Printf("**** FIND level [%d] %s\n", i, v)
		//fmt.Printf("**** FIND level [%d] %s in %+v\n", i, v, searchObj)
		if i+1 < len(s) {
			tmp, found := searchObj[v]
			//fmt.Printf("** tmp is %+v\n", tmp)
			if found {
				searchObj, found = tmp.(map[string]interface{})
				//fmt.Printf("** tmp->searchObj AS MAP is %+v\n", searchObj)
				if !found {
					searchObj, found = tmp.(ArgsMap)
					//fmt.Printf("** tmp->searchObj AS ARGSMAP is %+v\n", searchObj)
					if !found {
						// unknown object shape for a non-leaf level
						log.Errorf("PutObject: unknown object shape for a non-leaf level: %+v", tmp)
						return objIn, false
					}
				}
			} else {
				//fmt.Printf("** PutObject level not found in obj %+v, creating %s\n", searchObj, v)
				// level not found, create it and reset searchObj
				searchObj[v] = make(map[string]interface{})
				searchObj = searchObj[v].(map[string]interface{})
			}
		} else {
			//fmt.Printf("** PutObject leaf node to be written into obj %+v, creating %s with value %+v\n", searchObj, v, value)
			// leaf node, assign the value and return
			searchObj[v] = value
			//fmt.Printf("**** Found level [%d] %s\n", i, v)
			return objIn, true
		}
	}
	log.Errorf("PutObject: unknown error -- fell out of loop without returning")
	return objIn, false
}

// RemoveObject removes an object by its qualified name, which looks like
// "location.latitude" as one example.
func RemoveObject(objIn map[string]interface{}, qname string) (map[string]interface{}, bool) {
	searchObj := objIn
	s := strings.Split(qname, ".")
	for i, v := range s {
		if i+1 < len(s) {
			tmp, found := AsMap(searchObj[v])
			if !found {
				return objIn, false
			}
			searchObj = tmp
			continue
		}
		delete(searchObj, v)
		break
	}
	return objIn, true
}

// AddToStringArray merges a specified object (usually in asset state) by qualified name with an incoming
// string or string array. Keeps only unique members (as in a set.)
func AddToStringArray(objIn map[string]interface{}, qname string, value interface{}) (map[string]interface{}, bool) {
	_, ok := AsMap(objIn)
	if !ok {
		log.Errorf("addToStringArray: incoming object not a map type %T[%#v]", objIn, objIn)
		return objIn, false
	}
	arr, found := GetObject(objIn, qname)
	if !found {
		// array object does not exist yet, for convenience we create it
		// log.Debugf("addToStringArray: redirecting to PutObject of %s : %#v", qname, value)
		return PutObject(objIn, qname, value)
	}
	// log.Debugf("addToStringArray: adding %#v to %s : %#v\n", value, qname, arr)
	// we have an array to modify
	sarr, ok := AsStringArray(arr)
	if !ok {
		log.Errorf("addToStringArray: incoming object is not a string array %s : %T[%#v]", qname, arr, arr)
		return objIn, false
	}
	sval, ok := AsStringArray(value)
	if !ok {
		log.Errorf("addToStringArray: incoming value is not a string or string array %s : %T[%#v]", qname, value, value)
		return objIn, false
	}
	for _, v := range sval {
		if Contains(sarr, v) {
			continue
		} else {
			sarr = append(sarr, v)
		}
	}
	sort.Strings(sarr)
	// log.Debugf("addToStringArray: calling PutObject with result %#v\n", sarr)
	return PutObject(objIn, qname, sarr)
}

// RemoveFromStringArray removes from a named object in asset state or other map, an incoming
// string or string array. Assumes unique members (as in a set.)
func RemoveFromStringArray(objIn map[string]interface{}, qname string, value interface{}) (map[string]interface{}, bool) {
	arr, found := GetObject(objIn, qname)
	if !found {
		// log.Debugf("addToStringArray: array %s not found", qname)
		return objIn, false
	}
	// log.Debugf("removeFromStringArray: removing %#v from %s : %#v\n", value, qname, arr)
	sarr, ok := AsStringArray(arr)
	if !ok {
		log.Errorf("removeFromStringArray: incoming object is not a string array %s : %T[%#v]", qname, arr, arr)
		return objIn, false
	}
	sval, ok := AsStringArray(value)
	if !ok {
		log.Errorf("removeFromStringArray: incoming value is not a string or string array %s : %T[%#v]", qname, value, value)
		return objIn, false
	}
	r := sarr[:0]
	for _, v := range sarr {
		// fmt.Printf("r: %#v v: %#v sval: %#v\n", r, v, sval)
		if !Contains(sval, v) {
			r = append(r, v)
		}
	}
	sort.Strings(r)
	// log.Debugf("addToStringArray: calling PutObject with result %#v\n", r)
	return PutObject(objIn, qname, r)
}

// GetObjectAsMap retrieves an object by qualified name and then runs AsMap on it to
// interpret or cast it to map[string]interface{}
func GetObjectAsMap(objIn map[string]interface{}, qname string) (map[string]interface{}, bool) {
	amap, found := GetObject(objIn, qname)
	if found {
		t, found := AsMap(amap)
		if found {
			return t, true
		}
		log.Warningf("GetObjectAsMap object is not a map: %s but rather %T", qname, objIn)
	}
	return nil, false
}

// GetObjectAsString retrieves an object by qualified name and interprets or casts it to string
func GetObjectAsString(objIn map[string]interface{}, qname string) (string, bool) {
	tbytes, found := GetObject(objIn, qname)
	if found {
		t, found := tbytes.(string)
		if found {
			return t, true
		}
		log.Warningf("GetObjectAsString object is not a string: %s", qname)
	}
	return "", false
}

// GetObjectAsStringArray retrieves an object by qualified name and interprets or casts it to []string
func GetObjectAsStringArray(objIn map[string]interface{}, qname string) ([]string, bool) {
	tbytes, found := GetObject(objIn, qname)
	if found {
		return AsStringArray(tbytes)
	}
	return []string{}, false
}

// GetObjectAsBoolean retrieves an object by qualified name and interprets or casts it to bool
func GetObjectAsBoolean(objIn map[string]interface{}, qname string) (bool, bool) {
	tbytes, found := GetObject(objIn, qname)
	if found {
		t, found := tbytes.(bool)
		if found {
			return t, true
		}
		log.Warningf("GetObjectAsBoolean object is not a boolean: %s", qname)
	}
	return false, false
}

// GetObjectAsNumber retrieves an object by qualified name and interprets or casts it to float64
func GetObjectAsNumber(objIn map[string]interface{}, qname string) (float64, bool) {
	tbytes, found := GetObject(objIn, qname)
	if found {
		t, found := tbytes.(float64)
		if found {
			return t, true
		}
		log.Warningf("GetObjectAsNumber object is not a number (float64): %s", qname)
	}
	return 0, false
}

// GetObjectAsInteger retrieves an object by qualified name and interprets or casts it to integer
// NOTE: will truncate in incoming JSON Number (float64)
func GetObjectAsInteger(objIn map[string]interface{}, qname string) (int, bool) {
	tbytes, found := GetObject(objIn, qname)
	if found {
		// try as int first
		i, found := tbytes.(int)
		if found {
			return i, true
		}
		// try as JSON number and then cast
		f, found := tbytes.(float64)
		if found {
			return int(f), true
		}
		log.Warningf("GetObjectAsInteger object is not an integer: %s", qname)
	}
	return 0, false
}

// Contains does its best to assert an array type on the incoming array
// and the matching type on the incoming val, and then searches for val
// in arr.
func Contains(arr interface{}, val interface{}) bool {
	switch t := arr.(type) {
	case []string:
		arr2 := arr.([]string)
		for _, v := range arr2 {
			if v == val {
				return true
			}
		}
	case []int:
		arr2 := arr.([]int)
		for _, v := range arr2 {
			if v == val {
				return true
			}
		}
	case []float64:
		arr2 := arr.([]float64)
		for _, v := range arr2 {
			if v == val {
				return true
			}
		}
	case []interface{}:
		//todo: try cast instead of assertion
		//todo: use schema to determine if we even call this function or just add the value
		arr2 := arr.([]interface{})
		for _, v := range arr2 {
			switch tt := val.(type) {
			case string:
				if v.(string) == val.(string) {
					return true
				}
			case int:
				if v.(int) == val.(int) {
					return true
				}
			case float64:
				if v.(float64) == val.(float64) {
					return true
				}
			case interface{}:
				if v.(interface{}) == val.(interface{}) {
					return true
				}
			default:
				log.Errorf("Contains passed array containing unknown type: %+v\n", tt)
				return false
			}
		}
	default:
		log.Errorf("Contains passed array of unknown type: %+v\n", t)
		return false
	}
	return false
}

// DeepMergeMap all levels of a src map into a dst map and return dst
func DeepMergeMap(srcIn map[string]interface{}, dstIn map[string]interface{}) map[string]interface{} {
	for k, v := range srcIn {
		switch v.(type) {
		case map[string]interface{}:
			dstv, found := GetObject(dstIn, k)
			if found {
				// recursive DeepMerge into existing key
				dstIn[k] = DeepMergeMap(v.(map[string]interface{}), dstv.(map[string]interface{}))
			} else {
				// copy src to dst at same key
				dstIn[k] = v
			}
		case []interface{}:
			dstv, found := GetObject(dstIn, k)
			if found {
				if _, isString := dstv.([]string); isString {
					dstm, ok := AddToStringArray(dstv.(map[string]interface{}), k, v)
					if ok {
						dstIn[k] = dstm
					} else {
						// failed, must copy
						dstIn[k] = v
					}
				} else {
					// not strings, copy for now
					dstIn[k] = v
				}
			} else {
				// copy
				dstIn[k] = v
			}
		default:
			// copy discrete type
			dstIn[k] = v
		}
	}
	return dstIn
}

// PrettyPrint returns a string that is a nicely indented representation
// of js object (map); if json fails for some reason, returns the %#v representation
func PrettyPrint(m interface{}) string {
	bytes, err := json.MarshalIndent(m, "", "  ")
	if err == nil {
		return string(bytes)
	}
	return fmt.Sprintf("%#v", m)
}
