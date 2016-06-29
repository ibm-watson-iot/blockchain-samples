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


// v3.0 HM 25 Feb 2016 Moved the asset state history code into a separate package.
// v3.0.1 HM 03 Mar 2016 Store the state history in descending order.
// v3.0.2 KL 07 Mar 2016 Reduce memory garbage in updateStateHistory 
// v3.0.3 KL             backported from original 3.1/4.0 
// v4.0 KL 17 Mar 2016 Update version number to 4.0 for Hyperledger compatibility.
//                     Clean up lint issues.

package main

import (
	"encoding/json"
    "fmt"
    "sort"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)


// MYVERSION Update for every change, use VX.X.X (Major, Minor, Fix). Suggest that we update
// Major for API break, Minor when adding a feature or behavior, Fix when fixing a bug.
// If the init comes in with the wrong major version, then  we might consider exiting with
// an error.
const MYVERSION string = "4.3"

// DEFAULTNICKNAME is used when a contract is initialized without giving it a nickname
const DEFAULTNICKNAME string = "AVIATION" 

// CONTRACTSTATEKEY is used to store contract state, including version, nickname and activeAssets
const CONTRACTSTATEKEY string = "ContractStateKey"

// ContractState struct defines contract state. Unlike the main contract maps, structs work fine
// for this fixed structure.
type ContractState struct {
	Version      string           `json:"version"`
    Nickname     string           `json:"nickname"`
	ActiveAirlines map[string]bool  `json:"activeAirlines"`
	ActiveAirplanes map[string]bool  `json:"activeAirplanes"`
	ActiveAssemblies map[string]bool  `json:"activeAssemblies"`
	ActiveParts map[string]bool  `json:"activeParts"`
}

// GETContractStateFromLedger retrieves state from ledger and returns to caller
func GETContractStateFromLedger(stub *shim.ChaincodeStub) (ContractState, error) {
    var state = ContractState{ MYVERSION, DEFAULTNICKNAME, 
        make(map[string]bool), make(map[string]bool), 
        make(map[string]bool), make(map[string]bool) }
    var err error
	contractStateBytes, err := stub.GetState(CONTRACTSTATEKEY)
    // minimum string is {"version":""} and version cannot be empty 
	if err == nil && len(contractStateBytes) > 14 {    
		// apparently, this blockchain instance is being reloaded, has the version changed?
		err = json.Unmarshal(contractStateBytes, &state)
		if err != nil {
            err = fmt.Errorf("Unmarshal failed for contract state: %s", err)
            log.Critical(err)
			return ContractState{}, err
		}
        if MYVERSION != state.Version {
            log.Noticef("Contract version has changed from %s to %s", state.Version, MYVERSION)
            state.Version = MYVERSION
        }
	} else {
        // empty state already initialized 
		log.Noticef("Initialized newly deployed contract state version %s", state.Version)
	}
    // this MUST be here
    if state.ActiveAirlines == nil {
        state.ActiveAirlines = make(map[string]bool)
    }
    if state.ActiveAirplanes == nil {
        state.ActiveAirplanes = make(map[string]bool)
    }
    if state.ActiveAssemblies == nil {
        state.ActiveAssemblies = make(map[string]bool)
    }
    if state.ActiveParts == nil {
        state.ActiveParts = make(map[string]bool)
    }
    log.Debug("GETContractState successful")
    return state, nil 
}

// PUTContractStateToLedger writes a contract state into the ledger
func PUTContractStateToLedger(stub *shim.ChaincodeStub, state ContractState) (error) {
    var contractStateJSON []byte
    var err error
    contractStateJSON, err = json.Marshal(state)
    if err != nil {
        err = fmt.Errorf("Failed to marshal contract state: %s", err)
        log.Critical(err)
        return err
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        err = fmt.Errorf("Failed to PUTSTATE contract state: %s", err)
        log.Critical(err)
        return err
    } 
    log.Debugf("PUTContractState: %#v", state)
    return nil 
}

func addAirlineToContractState(stub *shim.ChaincodeStub, airlineID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding airline %s to contract", airlineID)
    state.ActiveAirlines[airlineID] = true
    return PUTContractStateToLedger(stub, state)
}

func addAirplaneToContractState(stub *shim.ChaincodeStub, airplaneID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding airplane %s to contract", airplaneID)
    state.ActiveAirplanes[airplaneID] = true
    return PUTContractStateToLedger(stub, state)
}

func addAssemblyToContractState(stub *shim.ChaincodeStub, assemblyID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding assemblie %s to contract", assemblyID)
    state.ActiveAssemblies[assemblyID] = true
    return PUTContractStateToLedger(stub, state)
}

func addPartToContractState(stub *shim.ChaincodeStub, partID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding part %s to contract", partID)
    state.ActiveParts[partID] = true
    return PUTContractStateToLedger(stub, state)
}

func removeAirlineFromContractState(stub *shim.ChaincodeStub, airlineID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding airline %s to contract", airlineID)
    delete(state.ActiveAirlines, airlineID)
    return PUTContractStateToLedger(stub, state)
}

func removeAirplaneFromContractState(stub *shim.ChaincodeStub, airplaneID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding airplane %s to contract", airplaneID)
    delete(state.ActiveAirplanes, airplaneID)
    return PUTContractStateToLedger(stub, state)
}

func removeAssemblyFromContractState(stub *shim.ChaincodeStub, assemblyID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding assemblie %s to contract", assemblyID)
    delete(state.ActiveAssemblies, assemblyID)
    return PUTContractStateToLedger(stub, state)
}

func removePartFromContractState(stub *shim.ChaincodeStub, partID string) (error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return err
    }
    log.Debugf("Adding part %s to contract", partID)
    delete(state.ActiveParts, partID)
    return PUTContractStateToLedger(stub, state)
}

func getActiveAirlines(stub *shim.ChaincodeStub) ([]string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return []string{}, err
    }
    var a = make([]string, len(state.ActiveAirlines))
    i := 0
    for id := range state.ActiveAirlines {
        a[i] = id
        i++ 
    }
    sort.Strings(a)
    return a, nil
}

func getActiveAirplanes(stub *shim.ChaincodeStub) ([]string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return []string{}, err
    }
    var a = make([]string, len(state.ActiveAirplanes))
    i := 0
    for id := range state.ActiveAirplanes {
        a[i] = id
        i++ 
    }
    sort.Strings(a)
    return a, nil
}

func getActiveAssemblies(stub *shim.ChaincodeStub) ([]string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return []string{}, err
    }
    var a = make([]string, len(state.ActiveAssemblies))
    i := 0
    for id := range state.ActiveAssemblies {
        a[i] = id
        i++ 
    }
    sort.Strings(a)
    return a, nil
}

func getActiveParts(stub *shim.ChaincodeStub) ([]string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return []string{}, err
    }
    var a = make([]string, len(state.ActiveParts))
    i := 0
    for id := range state.ActiveParts {
        a[i] = id
        i++ 
    }
    sort.Strings(a)
    return a, nil
}

func initializeContractState(stub *shim.ChaincodeStub, version string, nickname string) (error) {
    var state ContractState
    var err error
    if version != MYVERSION {
        err = fmt.Errorf("Contract version: %s does not match version argument: %s", MYVERSION, version)
        log.Critical(err)
        return err
    }
    state, err = GETContractStateFromLedger(stub)
    if err != nil {
        return err
    }  
    if version != state.Version {
        log.Noticef("Contract version has changed from %s to %s", version, MYVERSION)
        // keep going, this is an update of version -- later this will
        // be handled by pulling state from the superseded contract version
    }
    state.Version = MYVERSION
    state.Nickname = nickname
    return PUTContractStateToLedger(stub, state)
}

func getLedgerContractVersion(stub *shim.ChaincodeStub) (string, error) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)  
    if err != nil {
        return "", err
    }
    return state.Version, nil   
}

func airlineIsActive(stub *shim.ChaincodeStub, airlineID string) (bool) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)
    if err != nil { return false}
    found, _ := state.ActiveAirlines[airlineID]
    return found
}                      

func airplaneIsActive(stub *shim.ChaincodeStub, airplaneID string) (bool) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)
    if err != nil { return false}
    found, _ := state.ActiveAirlines[airplaneID]
    return found
}                      

func assemblyIsActive(stub *shim.ChaincodeStub, assemblyID string) (bool) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)
    if err != nil { return false}
    found, _ := state.ActiveAirlines[assemblyID]
    return found
}                      

func partIsActive(stub *shim.ChaincodeStub, partID string) (bool) {
    var state ContractState
    var err error
    state, err = GETContractStateFromLedger(stub)
    if err != nil { return false}
    found, _ := state.ActiveAirlines[partID]
    return found
}                      
