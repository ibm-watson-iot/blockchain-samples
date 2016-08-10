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

// v1 KL 07 Aug 2016 Add event handling in a separate module

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func handleEvent(stub *shim.ChaincodeStub, eventName string, assetID string, argsMap ArgsMap, ledgerMap ArgsMap) ArgsMap {
    switch eventName {
        case "flight":
            return handleAircraftEvent(stub, eventName, assetID, argsMap, ledgerMap) 
        case "inspection":
            return handleAssemblyEvent(stub, eventName, assetID, argsMap, ledgerMap) 
        case "analyticAdjustment":
            return handleAssemblyEvent(stub, eventName, assetID, argsMap, ledgerMap) 
        default:
            log.Errorf("handleEvent: event %s for assetID %s is unknown\n", eventName, assetID)
            return nil
    }
}
