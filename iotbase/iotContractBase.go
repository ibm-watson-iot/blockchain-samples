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

// v0.1 KL -- new IOT sample with Trade Lane properties and behaviors

package iotbase

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
	ru "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctcompliance"
	cf "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctconfig"
	hs "github.com/ibm-watson-iot/blockchain-samples/iotbase/cthistory"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/cthistory"
)

var log = shim.NewLogger("base")

// Init is used to load all sub packaged and allow their init() registrations to run
func Init() {
	as.Enabled = true
	cf.Enabled = true
	ru.Enabled = true
	hs.Enabled = true
	st.Enabled = true
}
