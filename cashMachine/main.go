/*******************************************************************************
Copyright (c) 2016 IBM Corporation and other Contributors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Sumabala Nair - Modified SimpleContract for CashMachine use case

******************************************************************************/
//SN: March 2016

// IoT Blockchain Simple Smart Contract v 1.0

// This is a simple contract that creates a CRUD interface to 
// create, read, update and delete an asset

package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
    "time"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

const CONTRACTSTATEKEY string = "ContractStateKey"  
// store contract state - only version in this example
const MYVERSION string = "1.0"

const HISTKEY string = "_HIST"
// ************************************
// asset and contract state 
// ************************************

type ContractState struct {
    Version      string                        `json:"version"`
}

type CashMachineState struct {
    AssetID          string       `json:"assetid,omitempty"`        // all assets must have an ID, primary key of contract
    ActionType       string       `json:"actiontype,omitempty"`       
    Amount           float64      `json:"amount,omitempty"`    
    Balance          float64      `json:"balance,omitempty"`
    Timestamp        string       `json:"timestamp,omitempty"`        
}

type CashMachineHistory struct {
	CashHistory []string `json:"cashhistory"`
}


var contractState = ContractState{MYVERSION}


// ************************************
// deploy callback mode 
// ************************************
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var stateArg ContractState
    var err error
    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version string")
    }
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    if stateArg.Version != MYVERSION {
        return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
    }
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    err = stub.PutState(CONTRACTSTATEKEY, contractStateJSON)
    if err != nil {
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
}

// ************************************
// deploy and invoke callback mode 
// ************************************
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    // Handle different functions
    if function == "createAsset" {
        // create assetID
        return t.createAsset(stub, args)
    } else if function == "updateAsset" {
        // create assetID
        return t.updateAsset(stub, args)
    } else if function == "deleteAsset" {
        // Deletes an asset by ID from the ledger
        return t.deleteAsset(stub, args)
    }
    return nil, errors.New("Received unknown invocation: " + function)
}

// ************************************
// query callback mode 
// ************************************
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    // Handle different functions
    if function == "readAsset" {
        // gets the state for an assetID as a JSON struct
        return t.readAsset(stub, args)
    } else if function =="readAssetHistory" {
        return t.readAssetHistory(stub, args)
    } else if function == "readAssetSamples" {
		// returns selected sample objects 
		return t.readAssetSamples(stub, args)
	} else if function == "readAssetSchemas" {
		// returns selected sample objects 
		return t.readAssetSchemas(stub, args)
	}
    return nil, errors.New("Received unknown invocation: " + function)
}

/**********main implementation *************/

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple Chaincode: %s", err)
    }
}

/*****************ASSET CRUD INTERFACE starts here************/

/****************** 'deploy' methods *****************/

/******************** createCashMachine ********************/

func (t *SimpleChaincode) createAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    _,erval:=t. createOrupdateCashMachine(stub, args)
    return nil, erval
}

//******************** updateCashMachine ********************/

func (t *SimpleChaincode) updateAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
     _,erval:=t. createOrupdateCashMachine(stub, args)
    return nil, erval
}


//******************** deleteCashMachine ********************/

func (t *SimpleChaincode) deleteAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var assetID string // asset ID
    var err error
    var stateIn CashMachineState

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err = t.validateInput(args)
    if err != nil {
        return nil, err
    }
    assetID = stateIn.AssetID
    // Delete the key / asset from the ledger
    err = stub.DelState(assetID)
    if err != nil {
        err = errors.New("Asset record delete failed! : "+ fmt.Sprint(err))
       return nil, err
    }
    cmHistKey:= stateIn.AssetID+HISTKEY
    err = stub.DelState(cmHistKey)
    if err != nil {
        err = errors.New("Asset History delete failed! : "+ fmt.Sprint(err))
       return nil, err
    }
    return nil, nil
}

/******************* Query Methods ***************/

//********************readCashMachine********************/

func (t *SimpleChaincode) readAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var assetID string // asset ID
    var err error
    var state CashMachineState

     // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err:= t.validateInput(args)
    if err != nil {
        return nil, errors.New("Asset does not exist!")
    }
    assetID = stateIn.AssetID
        // Get the state from the ledger
    assetBytes, err:= stub.GetState(assetID)
    if err != nil  || len(assetBytes) ==0{
        err = errors.New("Unable to get asset state from ledger")
        return nil, err
    } 
    err = json.Unmarshal(assetBytes, &state)
    if err != nil {
         err = errors.New("Unable to unmarshal state data obtained from ledger")
        return nil, err
    }
    return assetBytes, nil
}

//*************readCashMachineObjectModel*****************/

func (t *SimpleChaincode) readAssetHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var cmHistKey string // asset ID history key
    var err error
    var state CashMachineState

     // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    stateIn, err:= t.validateInput(args)
    if err != nil {
        return nil, errors.New("Asset does not exist!")
    }
    cmHistKey = stateIn.AssetID+HISTKEY
        // Get the state from the ledger
    assetBytes, err:= stub.GetState(cmHistKey)
    if err != nil  || len(assetBytes) ==0{
        err = errors.New("Unable to get asset history from ledger")
        return nil, err
    } 
    err = json.Unmarshal(assetBytes, &state)
    if err != nil {
         err = errors.New("Unable to unmarshal asset history data obtained from ledger")
        return nil, err
    }
    return assetBytes, nil
}

//*************readCashMachineSamples*******************

func (t *SimpleChaincode) readAssetSamples(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return []byte(samples), nil
}

//*************readCashMachineSchemas*******************

func (t *SimpleChaincode) readAssetSchemas(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return []byte(schemas), nil
}


// ************************************
// validate input data : common method called by the CRUD functions
// ************************************
func (t *SimpleChaincode) validateInput(args []string) (stateIn CashMachineState, err error) {
    var assetID string // asset ID
    var state  = CashMachineState{} // The calling function is expecting an object of type AssetState

    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a JSON strings with mandatory assetID")
        return state, err
    }
    jsonData:=args[0]
    assetID = ""
    stateJSON := []byte(jsonData)
   // fmt.Println("Input data ",jsonData)
    err = json.Unmarshal(stateJSON, &stateIn)
    if err != nil {
        err = errors.New("Unable to unmarshal input JSON data")
        return state, err
        // state is an empty instance of asset state
    }      
    // was assetID present?
    // The nil check is required because the asset id is a pointer. 
    // If no value comes in from the json input string, the values are set to nil
    
  
    assetID = strings.TrimSpace(stateIn.AssetID)
     //   fmt.Println("assetID ",assetID)
    if assetID==""{
        err = errors.New("AssetID not passed")
        return state, err
    }
    stateIn.AssetID = assetID
   //  fmt.Println("assetID after val ",stateIn.AssetID)
    return stateIn, nil
}
//******************** createOrupdateCashMachine ********************/

func (t *SimpleChaincode) createOrupdateCashMachine(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var assetID string                 // asset ID                    // used when looking in map
    var err error
    var stateIn CashMachineState
    var stateStub CashMachineState
    var cashMachineHistory CashMachineHistory
    var cmHistoryState []byte

    // validate input data for number of args, Unmarshaling to asset state and obtain asset id
    
    stateIn, err = t.validateInput(args)
    if err != nil {
        return nil, err
    }
   // fmt.Println("after validate input ")
    assetID = stateIn.AssetID
    stimeStamp:= strings.TrimSpace(stateIn.Timestamp )
    if stimeStamp==""{
        // Obtain timestamp from the stub
         txnTime, err:= stub.GetTxTimestamp()
        if err !=nil {
            err=errors.New("Unable to get transaction time")
            return nil, err
        }
        txntimestamp := time.Unix(txnTime.Seconds, int64(txnTime.Nanos))
        sTime := txntimestamp.String()
        stateIn.Timestamp = sTime
    } else {
        stateIn.Timestamp= stimeStamp
    }
   // fmt.Println("Time: ", stateIn.Timestamp)
    cmHistKey := stateIn.AssetID+HISTKEY
    // Check if asset record existed in stub
    assetBytes, err:= stub.GetState(assetID)
   // fmt.Println ("error is ", err)
    if err != nil || len(assetBytes)==0{
        // This implies that this is a 'create' scenario
         stateIn.Balance=stateIn.Amount
         stateIn.ActionType="InitialBalance"
         stateStub = stateIn // The record that goes into the stub is the one that comes in, since this is a create scenario
        // fmt.Println("assetBytes ", stateStub.AssetID)
         // If its a create scenario, history cannot exist
        // Do We want to store the raw data in history? Right now it stores processed data
        stateInJSON, err := json.Marshal(stateIn)
        if err != nil {
            return nil, errors.New("Marshal failed for incoming contract state" + fmt.Sprint(err))
        }
        var cmHistory = CashMachineHistory{make([]string, 1)}
        cmHistory.CashHistory[0] = string(stateInJSON)
        cmState, err := json.Marshal(cmHistory)
        if err != nil {
            return nil,err
        }
        cmHistoryState = []byte(cmState)
    } else {
        // This is an update scenario
       // fmt.Println("Update Scenario")
       // fmt.Println("Data in stub was :", string(assetBytes))
        err = json.Unmarshal(assetBytes, &stateStub)
        if err != nil {
            err = errors.New("Unable to unmarshal JSON data from stub")
            return nil, err
            // state is an empty instance of asset state
        }
       // fmt.Println("stateIn.ActionType is ", stateIn.ActionType)
        stateStub.Amount = stateIn.Amount
        if stateIn.ActionType == "Deposit" {
            stateStub.Balance= stateStub.Balance + stateIn.Amount
        } else { 
            stateStub.Balance= stateStub.Balance - stateIn.Amount
            // Assuming stateStub.ActionType is "Withdraw"
        }
        stateIn.Balance = stateStub.Balance // assuming history needs banace too
        stateStub.Timestamp = stateIn.Timestamp // updating the stub record
        stateStub.ActionType = stateIn.ActionType
        // get the histoty of the machine
        // Do We want to store the raw data in history? Right now it stores processed data
        stateInJSON, err := json.Marshal(stateIn)
        if err != nil {
            return nil, errors.New("Marshal failed for incoming contract state" + fmt.Sprint(err))
        }
        cmHistory, err := stub.GetState(cmHistKey)
        if err != nil {
            return nil,err
        }
  
        err = json.Unmarshal(cmHistory, &cashMachineHistory)
        if err != nil {
            return nil, err
        }

        var cmSlice = make([]string, 0)
        cmSlice = append(cmSlice, string(stateInJSON))
        cmSlice = append(cmSlice, cashMachineHistory.CashHistory...)
        cashMachineHistory.CashHistory = cmSlice
        cmState, err := json.Marshal(cashMachineHistory)
        if err != nil {
            return nil,err
        }
        cmHistoryState = []byte(cmState)
    }
    
    
    // Now that the statestub record has the updated data, we can put it in the stub
    stateJSON, err := json.Marshal(stateStub)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
   
    // Write the new state to the ledger
    err = stub.PutState(assetID, stateJSON)
    if err != nil {
        err = errors.New("PUT ledger state failed: "+ fmt.Sprint(err))            
        return nil, err
    } 
     err = stub.PutState(cmHistKey, cmHistoryState)
    if err != nil {
        return nil, errors.New("Cash machine transaction history failed PUT to ledger: " + fmt.Sprint(err))
    }  
    return nil, nil
}
