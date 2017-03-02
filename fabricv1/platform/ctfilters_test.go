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
	"testing"
)

var f = "{\"filter\":{\"match\":\"all\", \"select\":[{\"qprop\":\"b\", \"value\":\"c\"},{\"qprop\":\"c\",\"value\":\"d\"}]}}"
var f2 = "{\"match\":\"all\", \"select\":[{\"qprop\":\"b\", \"value\":\"c\"},{\"qprop\":\"c\",\"value\":\"d\"}]}"
var f3 = "{\"filter\":{\"match\":\"all\",\"select\":{\"0\":{\"qprop\":\"a\",\"value\":\"b\"},\"1\":{\"qprop\":\"c\",\"value\":\"d\"}}}}"
var f4 = "{\"match\":\"all\",\"select\":{\"0\":{\"qprop\":\"a\",\"value\":\"b\"},\"1\":{\"qprop\":\"c\",\"value\":\"d\"}}}"

func TestJSONUnmarshal(t *testing.T) {
	var tagout TaggedFilter
	err := json.Unmarshal([]byte(f), &tagout)
	if err != nil || tagout.Filter.Match != "all" || len(tagout.Filter.Select) == 0 {
		t.Fail()
		fmt.Printf("tagged to tagged: %+v || %+v\n", tagout, err)
	}

	var out2 StateFilter
	err = json.Unmarshal([]byte(f2), &out2)
	if err != nil || out2.Match != "all" || len(out2.Select) == 0 {
		t.Fail()
		fmt.Printf("untagged to untagged: %+v || %+v\n", out2, err)
	}
}

func TestCanonicalTaggedToTagged(t *testing.T) {
	var filter StateFilter
	filter, err := getCanonicalFilterFromEventIn([]string{f})
	if err != nil || filter.Match != "all" || len(filter.Select) == 0 {
		t.Fail()
		fmt.Printf("*** canonical tagged to tagged: [%+v]==>[%+v] : err [%+v]\n", f, filter, err)
	}
}

func TestCanonicalUntaggedToUntagged(t *testing.T) {
	var filter StateFilter
	filter, err := getCanonicalFilterFromEventIn([]string{f2})
	if err != nil || filter.Match != "all" || len(filter.Select) == 0 {
		t.Fail()
		fmt.Printf("*** canonical untagged to untagged: [%+v]==>[%+v] : err [%+v]\n", f2, filter, err)
	}
}

func TestMapFormatTaggedToTagged(t *testing.T) {
	var filter StateFilter
	filter, err := getMapFormatFilterFromEventIn([]string{f3})
	if err != nil || filter.Match != "all" || len(filter.Select) == 0 {
		t.Fail()
		fmt.Printf("*** mapformat tagged to tagged: [%+v]==>[%+v] : err [%+v]\n", f3, filter, err)
	}
}

func TestMapFormatUntaggedToUntagged(t *testing.T) {
	var filter StateFilter
	filter, err := getMapFormatFilterFromEventIn([]string{f4})
	if err != nil || filter.Match != "all" || len(filter.Select) == 0 {
		t.Fail()
		fmt.Printf("*** mapformat untagged to untagged: [%+v]==>[%+v] : err [%+v]\n", f4, filter, err)
	}
}

func TestGetUnmarshalledStateFilterTagged(t *testing.T) {
	var filter StateFilter
	var err error
	filter, err = getUnmarshalledStateFilter([]string{f})
	if err != nil || filter.Match != "all" || len(filter.Select) == 0 {
		t.Fail()
		fmt.Printf("*** getUnmarshalledStateFilter tagged: [%+v]==>[%+v] : err [%+v]\n", f, filter, err)
	}
}

func TestGetUnmarshalledStateFilterUntagged(t *testing.T) {
	var filter2 StateFilter
	filter2, err := getUnmarshalledStateFilter([]string{f2})
	if err != nil || filter2.Match != "all" || len(filter2.Select) == 0 {
		t.Fail()
		fmt.Printf("*** getUnmarshalledStateFilter untagged: [%+v]==>[%+v] : err [%+v]\n", f2, filter2, err)
	}
}

func TestGetUnmarshalledStateFilterTaggedObject(t *testing.T) {
	var filter3 StateFilter
	filter3, err := getUnmarshalledStateFilter([]string{f3})
	if err != nil || filter3.Match != "all" || len(filter3.Select) == 0 {
		t.Fail()
		fmt.Printf("*** getUnmarshalledStateFilter tagged object: [%+v]==>[%+v] : err [%+v]\n", f3, filter3, err)
	}
}

func TestGetUnmarshalledStateFilterUntaggedObject(t *testing.T) {
	var filter4 StateFilter
	filter4, err := getUnmarshalledStateFilter([]string{f4})
	if err != nil || filter4.Match != "all" || len(filter4.Select) == 0 {
		t.Fail()
		fmt.Printf("*** getUnmarshalledStateFilter untagged object: [%+v]==>[%+v] : err [%+v]\n", f4, filter4, err)
	}
}
