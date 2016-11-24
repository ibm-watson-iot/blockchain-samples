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

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// ChaincodeRoute stores a route for an asset class or event
type ChaincodeRoute struct {
	FunctionName string
	Method       string
	Class        AssetClass
	Function     func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)
}

// SimpleChaincode is the receiver for all shim API
type SimpleChaincode struct{}

// ChaincodeFunc is the signature for all functions that eminate from deploy, invoke or query messages to the contract
type ChaincodeFunc func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)

var router = make(map[string]ChaincodeRoute, 0)

// AddRoute allows a class definition to register its payload API, one route at a time
// functionName is the function that will appear in a rest or gRPC message
// method is one of deploy, invoke or query
// class is the asset class that created the route
// function is the actual function to be executed when the router is triggered
func AddRoute(functionName string, method string, class AssetClass, function ChaincodeFunc) error {
	if r, found := router[functionName]; found {
		err := fmt.Errorf("AddRoute: function name %s attempt to register against class %s as method %s but is already registered against class %s as method %s", class.Name, method, r.FunctionName, r.Class.Name, r.Method)
		log.Error(err)
		return err
	}
	r := ChaincodeRoute{
		FunctionName: functionName,
		Method:       method,
		Class:        class,
		Function:     function,
	}
	router[functionName] = r
	log.Debugf("Class %s added route with function name %s as method %s", r.Class.Name, r.FunctionName, r.Method)
	return nil
}

func getDeployFunctions() []ChaincodeFunc {
	var results = make([]ChaincodeFunc, 0)
	for _, r := range router {
		if r.Method == "deploy" {
			results = append(results, r.Function)
		}
	}
	return results
}

// EVTCCERR is a chaincode event ID that is emitted by the platform when an error
// is encountered during deploy or invoke
const EVTCCERR string = "EVTCCERR"

// EVTCCOK is a chaincode event ID that is emitted by the platform when a transaction
// completes successfully, invoke or deploy only
const EVTCCOK string = "EVTCCOK"

func setStubEvent(stub shim.ChaincodeStubInterface, event string, function string, method string, err string) {
	_ = stub.SetEvent(event, []byte(fmt.Sprintf(`{"event": %s,"function": %s,"method": %s,"result": %s}`, event, function, method, err)))
}

// Init is called by deploy messages
func Init(stub shim.ChaincodeStubInterface, function string, args []string, ContractVersion string) ([]byte, error) {
	var iargs = make([]string, 2)
	if len(args) == 0 {
		err := fmt.Errorf("Init received no args, expecting a json object in args[0]")
		log.Error(err)
		setStubEvent(stub, EVTCCERR, function, "deploy", err.Error())
		return nil, err
	}
	iargs[0] = args[0]
	iargs[1] = ContractVersion
	fs := getDeployFunctions()
	if len(fs) == 0 {
		err := fmt.Errorf("Init found no registered functions '%s'", function)
		log.Error(err)
		setStubEvent(stub, EVTCCERR, function, "deploy", err.Error())
		return nil, err
	}
	for _, f := range fs {
		_, err := f(stub, iargs)
		if err != nil {
			err := fmt.Errorf("Init (%s) failed with error %s", function, err)
			log.Error(err)
			setStubEvent(stub, EVTCCERR, function, "deploy", err.Error())
			return nil, err
		}
	}
	setStubEvent(stub, EVTCCOK, function, "deploy", "ok")
	return nil, nil
}

// Invoke is called when an invoke message is received
func Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var r ChaincodeRoute
	r, found := router[function]
	if !found {
		err := fmt.Errorf("Invoke did not find registered invoke function %s", function)
		log.Error(err)
		setStubEvent(stub, EVTCCERR, function, "invoke", err.Error())
		return nil, err
	}
	_, err := r.Function(stub, args)
	if err != nil {
		err := fmt.Errorf("Invoke (%s) failed with error %s", function, err)
		log.Error(err)
		setStubEvent(stub, EVTCCERR, function, "invoke", err.Error())
		return nil, err
	}
	setStubEvent(stub, EVTCCOK, function, "invoke", "ok")
	return nil, nil
}

// Query is called when a query message is received
func Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var r ChaincodeRoute
	r, found := router[function]
	if !found {
		err := fmt.Errorf("Query did not find registered query function %s", function)
		log.Error(err)
		return nil, err
	}
	result, err := r.Function(stub, args)
	if err != nil {
		err := fmt.Errorf("Query (%s) failed with error %s", function, err)
		log.Error(err)
		return nil, err
	}
	return result, nil
}

// readAllRoutes shows all registered routes
var readAllRoutes = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	type RoutesOut struct {
		FunctionName string     `json:"functionname"`
		Method       string     `json:"method"`
		Class        AssetClass `json:"class"`
	}
	var r = make([]RoutesOut, 0, len(router))
	for _, route := range router {
		ro := RoutesOut{
			route.FunctionName,
			route.Method,
			route.Class,
		}
		r = append(r, ro)
	}
	return json.Marshal(r)
}

func init() {
	AddRoute("readAllRoutes", "query", SystemClass, readAllRoutes)
}
