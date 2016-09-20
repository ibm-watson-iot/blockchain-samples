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
)

var ac1AssetID = "5678"
var as1AssetID = "acb"
var as2AssetID = "def"
var as3AssetID = "ghi"

var ac2AssetID = "1234"
var as4AssetID = "jkl"
var as5AssetID = "mno"

func TestIndexes(t *testing.T) {
	fmt.Println("Starting the index testing...")
	t.Log("Enter TestIndexes")

	indexes := makeAircraftAssemblyIndexes()
	fmt.Printf("Indexes %+v\n", indexes)

	_, found := indexes.isAssemblyOnAnyAircraft(as5AssetID)
	if found {
		t.Fatalf("aircraft assembly %s is on aircraft when there aren't any", as5AssetID)
	}

	err := indexes.addAssemblyToAircraft(as5AssetID, ac2AssetID)
	if err != nil {
		t.Fatalf("aircraft assembly %s could not be added to aircraft %s, error is %+v", as5AssetID, ac2AssetID, err)
	}
	err = indexes.addAssemblyToAircraft(as4AssetID, ac2AssetID)
	if err != nil {
		t.Fatalf("aircraft assembly %s could not be added to aircraft %s error is %+v", as4AssetID, ac2AssetID, err)
	}
	err = indexes.addAssemblyToAircraft(as5AssetID, ac2AssetID)
	if err == nil {
		t.Fatalf("aircraft assembly %s could not be added to aircraft %s error is %+v", as5AssetID, ac2AssetID, err)
	}
	if !indexes.isAssemblyOnThisAircraft(as4AssetID, ac2AssetID) {
		t.Fatalf("aircraft assembly %s not found on aircraft %s", as4AssetID, ac2AssetID)
	}
	if indexes.isAssemblyOnThisAircraft("shouldnotbefound", ac2AssetID) {
		t.Fatalf("aircraft assembly %s not found on aircraft %s", as4AssetID, ac2AssetID)
	}
	if ac, found := indexes.isAssemblyOnAnyAircraft("shouldnotbefound"); found {
		t.Fatalf("aircraft assembly %s found on aircraft %s", "shouldnotbefound", ac)
	}
	if err = indexes.removeAssemblyFromAircraft(as4AssetID, ac2AssetID); err != nil {
		t.Fatalf("aircraft assembly %s could not be removed from aircraft %s", as4AssetID, ac2AssetID)
	}
	if indexes.isAssemblyOnThisAircraft(as4AssetID, ac2AssetID) {
		t.Fatalf("aircraft assembly %s should not be found on aircraft %s", as4AssetID, ac2AssetID)
	}

	fmt.Printf("Indexes after %+v\n", indexes)

	assemblies, found := indexes.getAircraftAssemblies(ac2AssetID)
	if !found {
		t.Fatalf("aircraft %s failed to return assembles", ac2AssetID)
	}
	fmt.Printf("Aircraft %s has assemblies %+v\n", ac2AssetID, assemblies)

}
