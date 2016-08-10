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

// v1 KL 07 Aug 2016 Add event handling in a separate Aircraft module

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (t *SimpleChaincode) createAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return createAsset(stub, args, "aircraft", "createAssetAircraft")
}

func (t *SimpleChaincode) updateAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return updateAsset(stub, args, "aircraft", "updateAssetAircraft")
}

func (t *SimpleChaincode) deleteAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return deleteAsset(stub, args, "aircraft", "deleteAssetAircraft")
}

func (t *SimpleChaincode) deleteAllAssetsAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return deleteAllAssets(stub, args, "aircraft", "deleteAllAssetsAircraft")
}

func (t *SimpleChaincode) deletePropertiesFromAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return deletePropertiesFromAsset(stub, args, "aircraft", "deletePropertiesFromAssetAircraft")
}

func (t *SimpleChaincode) readAssetAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return readAsset(stub, args, "aircraft", "readAssetAircraft")
}

func (t *SimpleChaincode) readAllAssetsAircraft(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return readAllAssets(stub, args, "aircraft", "readAllAssetsAircraft")
}

func (t *SimpleChaincode) readAssetAircraftHistory (stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    return readAssetHistory(stub, args, "aircraft", "readAssetAircraftHistory")
}





func handleAircraftEvent(stub *shim.ChaincodeStub, assetName string, assetID string, argsMap ArgsMap, ledgerMap ArgsMap) ArgsMap {
    switch assetName {
        case "flight": 
            // running count of cycles in perpetuity
            cycles, found := getObjectAsInteger(ledgerMap, "aircraft.cycles")
            if found {
                cycles++
            } else {
                cycles = 1
            }
            putObject(ledgerMap, "aircraft.cycles", cycles)
            propagateEventToAssemblies(stub, assetName, assetID, argsMap, ledgerMap)
            return ledgerMap
        default:
            // NOTE: If your schema and contractConfig are correctly set up, this
            // CANNOT happen.
            err := fmt.Errorf("handleAircraftEvent: received incorrect event %s for asset %s", assetName, assetID)
            log.Error(err)
            return nil
    }
}

func propagateEventToAssemblies(stub *shim.ChaincodeStub, assetName string, assetID string, argsMap ArgsMap, ledgerMap ArgsMap) (error) {
    var me = "propagateEventToAssemblies"
    assemBytes, found := getObject(ledgerMap, "aircraft.assemblies")
    if found {
        assemblies, found := assemBytes.([]string)
        if found {
            for assem := range assemblies {
                // each of these is an *internal* prefixed assetID because the assembly
                // state has the original assetID and so we can use physical database
                // keys for connection
                // NOTE: This is the moral equivalent of "updateAsset"
                assembly := assemblies[assem]
                state, err := getUnmarshalledState(stub, me, assembly)
                if err != nil { return err }  // already logged
                state = handleAssemblyEvent(stub, assetName, assembly, argsMap, state)
                state, err = addTXNTimestampToState(stub, me, state)
                if err != nil { return err }  // already logged
                state = addLastEventToState(stub, me, argsMap, state, "")
                state, err = handleAlertsAndRules(stub, me, assetName, assembly, argsMap, state)
                if err != nil { return err }  // already logged
                err = putMarshalledState(stub, me, assetName, assembly, state)
                if err != nil { return err }  // already logged
            }
        }
    }
    return nil
}


