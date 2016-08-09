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

// v1 KL 07 Aug 2016 Add event handling in a separate Assembly module

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}
func (t *SimpleChaincode) updateAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}
func (t *SimpleChaincode) deleteAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}
func (t *SimpleChaincode) deleteAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}

func (t *SimpleChaincode) deletePropertiesFromAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}

func (t *SimpleChaincode) readAssetAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}

func (t *SimpleChaincode) readAllAssetsAssembly(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}

func (t *SimpleChaincode) readAssetAssemblyHistory (stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return nil, nil
}




func handleAssemblyEvent(stub *shim.ChaincodeStub, eventName string, assetID string, argsMap ArgsMap, ledgerMap ArgsMap) ArgsMap {
    switch eventName {
        case "flight": 
            // running count of cycles in perpetuity
            cycles, found := getObjectAsInteger(ledgerMap, "aircraft.cycles")
            if found {
                cycles++
            } else {
                cycles = 1
            }
            putObject(ledgerMap, "aircraft.cycles", cycles)
            // counter of cycles between aCheck inspections
            cyclecounter, found := getObjectAsInteger(ledgerMap, "aircraft.aCheckCounter")
            if found {
                cyclecounter++
            } else {
                cyclecounter = 1
            }
            putObject(ledgerMap, "aircraft.aCheckCounter", cyclecounter)
            // counter of consecutive hard landings for bCheck inspections
            gForce, found := getObjectAsNumber(argsMap, "flight.gForce")
            var bcheckcounter int
            if found {
                // TODO perhaps we should store the gForce threshold in the aircraft state
                if gForce > 2 { // temporary threshold for now 
                    bcheckcounter, found = getObjectAsInteger(ledgerMap, "aircraft.bCheckCounter")
                    if found {
                        bcheckcounter++
                    } else {
                        bcheckcounter = 1
                    }
                } else {
                    // reset on soft landing
                    bcheckcounter = 0
                }
            } else {
                // reset on no info
                bcheckcounter = 0
            }
            putObject(ledgerMap, "aircraft.bCheckCounter", bcheckcounter)
            return ledgerMap
        case "inspection":
            // nothing special to do, the rules engine will clear the alerts
        case "analyticAdjustment":
            
    }
    return nil
}
