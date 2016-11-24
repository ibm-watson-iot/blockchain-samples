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

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

// Update the path to match your configuration
//go:generate go run /local-dev/src/github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform/scripts/processSchema.go

// SimpleChaincode is the receiver for all shim API
type SimpleChaincode struct {
}

// CONTRACTVERSION is mandatory to use the platform **
const CONTRACTVERSION = "0.1"

// Logger for the cthistory package
var log = shim.NewLogger("main")

func main() {
	iot.SetContractLogger(shim.NewLogger("skit.track.trace"))
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		log.Infof("ERROR starting Simple Chaincode: %s", err)
	}
}

// Init is called in deploy mode and calls the router's Init function
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return iot.Init(stub, function, args, CONTRACTVERSION)
}

// Invoke is called in invoke mode and calls the router's Invoke function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return iot.Invoke(stub, function, args)
}

// Query is called in query mode and calls the router's Query function
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return iot.Query(stub, function, args)
}
