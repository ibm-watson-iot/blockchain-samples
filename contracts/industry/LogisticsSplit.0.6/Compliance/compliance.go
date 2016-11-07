package main
import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
    "reflect"
 //  "github.com/mcuadros/go-jsonschema-generator"
"github.com/hyperledger/fabric/core/chaincode/shim"
common "github.com/hyperledger/fabric/examples/chaincode/go/PoD/Common" 
)
// This contract is meant for complianace management
// Originally designed to track compliance for shipping contract's container data
// An Event listener could listen to this contract and raise a violation
// on the UI everytime a compliance update comes in.

const STATE string = "STATE"

const CONTRACTVERSIONKEY string = "COMPLIANCE" 
const SEP string = "_"
const DEFTYPE string = "SHIPPING"

type ContractVersion struct {
    Version      string                        `json:"version"`
    Type         string                        `json:"type"`
} 

type ComplianceHistory struct {
	CompHistory []string `json:"comphistory"`
}


//Features: Originally written in the context of a Shipping B/L
// 1. Create initial Compliance state (originally Bill of Lading compliance state)
// 2. Update  state with latest non-compliance
// 3. Store history of compliance alerts with timestamp
// 4. Provide the option to reset alerts - can be extended in future to include authentication
// 5. Flag state as 'Transit Complete' - 'Active' key gets set to inactive
// 6. Provide option to delete compliance history - this is primarily intended to support development

// SimpleChaincode  implementation
type SimpleChaincode struct {
}


//************* main *******************
//Create SimpleChaincode instance
//************* main *******************

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}
//***********************************************************
// First the chaincode functions - Init, Invoke and Query
//***********************************************************

/************************ Init function ********************/

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// The  compliance data is categorized based on complaince type, passed in as function 
    //Interpretation: This stands for the business function being addressed
	// For example, the compliance type could be 'Shipping' in the shipping case
	// Assuming one contract instance can manage multiple asset (eg. container), BLNo(eg. B/L) scenarios
	// A new intance need NOT be created for every business instance eg. every bill of lading
	
	var cVersion ContractVersion
	var err error

    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version and type")
    }
    err = json.Unmarshal([]byte(args[0]), &cVersion)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    // This needs expansion: Once versioning implementation is sorted out on the hyperledger side
   // fmt.Println ("Inside init")
    ver:=strings.TrimSpace(cVersion.Version)
    if ver != common.MYVERSION {
      //  fmt.Println ("Contract version " + common.MYVERSION + " must match version argument: " + cVersion.Version)
        return nil, errors.New("Contract version " + common.MYVERSION + " must match version argument: " + cVersion.Version)
    }
    bizType:=strings.TrimSpace(cVersion.Type)
     if bizType=="" {
         bizType=DEFTYPE // Type is  added to the key used for 'put' to the state
    } else {
        bizType=strings.ToUpper(function)
    }
    versionJSON, err := json.Marshal(cVersion)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    contractKey := CONTRACTVERSIONKEY+SEP+bizType 
    // This could be modified to implement as a composite key - similar to implementation @ ledger/statemgmt/commons.go 
  //  fmt.Println ("Contract key " , contractKey)
    err = stub.PutState(contractKey, versionJSON)
    if err != nil {
        // fmt.Println ("Contract state failed PUT to ledger")
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
	return nil, nil
}

/************************ Invoke function ********************/

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
     invFunc :=strings.TrimSpace(function)
    switch invFunc {
        case "createUpdateComplianceRecord" :
            return t.createUpdateComplianceRecord(stub, args)
        case "archiveComplianceRecord" :
            return t.archiveComplianceRecord(stub, args)
        case "resetComplianceHistory" :
            return t.resetComplianceHistory(stub, args)
        default:
            return nil, errors.New("Unknown function call to compliance : invoke")
    }
}

/************************ Query function ********************/

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    qFunc :=strings.TrimSpace(function)
    switch qFunc {
        case "readComplianceHistory" :
            return t.readComplianceHistory(stub, args)
        case "readCurrentComplianceState":
            return t.readCurrentComplianceState(stub, args)
        default:
            return nil, errors.New("Unknown function call to compliance : query")
    }
}

/************************ Invoke functions ********************/

func (t *SimpleChaincode) createUpdateComplianceRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var compHistoryState []byte
    var compHistKey     string
    var compStubState   common.ComplianceState
    var complianceHistory ComplianceHistory
  
    compInState, compData, err:= t.validateAndGet(stub, args)
    jsonData := args[0]
    // Type is just an indication of the business function now. It could be used for additional validation later.
    compHistKey = compInState.BLNo+SEP+STATE
    //If the record exists in the stub, this is an update, else a new record creation
    if compData =="" {
        // This is a create scenario
        if compInState.Type =="" {
            compInState.Type=DEFTYPE
        }
        compStubState = compInState
        // If its a create scenario, history cannot exist
        // We want to store the raw data in history
        var compHist = ComplianceHistory{make([]string, 1)}
        compHist.CompHistory[0] = jsonData
        compState, err := json.Marshal(&compHist)
        if err != nil {
            return nil,err
        }
        compHistoryState = []byte(compState)
    } else {
        // This is an update scenario
        err = json.Unmarshal([]byte(compData), &compStubState)
        if err != nil {
            err = errors.New("Unable to unmarshal JSON data from stub")
          //  fmt.Println(err)
            return nil, err
        }
        //compStubState, err =t.mergePartialState(compStubState,compInState)
	    compStubState = compInState
        comphistory, err := stub.GetState(compHistKey)
        if err != nil {
            return nil,err
        }
        
        err = json.Unmarshal(comphistory, &complianceHistory)
        if err != nil {
            return nil, err
        }

        var compSlice []string = make([]string, 0)
        compSlice = append(compSlice, jsonData)
        compSlice = append(compSlice, complianceHistory.CompHistory...)
        complianceHistory.CompHistory = compSlice
        compState, err := json.Marshal(&complianceHistory)
        if err != nil {
            return nil,err
        }
        compHistoryState = []byte(compState)
    }
    // Store compliance and history to the state
    compStubJSON, err := json.Marshal(compStubState)
    if err != nil {
        err:=errors.New("Marshaling compliance data failed")
        //fmt.Println(err)
        return nil, err
    }
     
    err = stub.PutState(compInState.BLNo, compStubJSON)
    if err != nil {
        return nil, errors.New("compliance state failed PUT to ledger: " + fmt.Sprint(err))
    }
    err = stub.PutState(compHistKey, compHistoryState)
    if err != nil {
        return nil, errors.New("compliance history failed PUT to ledger: " + fmt.Sprint(err))
    }  
    return nil, nil
}
    

 /*********************************  archiveComplianceRecord ****************************/
    
 func (t *SimpleChaincode) archiveComplianceRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
   var compStubState   common.ComplianceState
    compInState, compData, err:= t.validateAndGet(stub, args)
    err = json.Unmarshal([]byte(compData), &compStubState)
    if err != nil {
        err = errors.New("Unable to unmarshal JSON data from stub")
       // fmt.Println(err)
        return nil, err
    }
    compStubState.Active = false
    // Store compliance and history to the state
    compStubJSON, err := json.Marshal(compStubState)
    if err != nil {
        err:=errors.New("Marshaling compliance data failed")
        //fmt.Println(err)
        return nil, err
    }
     err = stub.PutState(compInState.BLNo, compStubJSON)
    if err != nil {
        return nil, errors.New("compliance state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
 }
     
  /*********************************  resetComplianceHistory ****************************/
      
 func (t *SimpleChaincode) resetComplianceHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var compStubState   common.ComplianceState
    mAlerts:= make(map[string]string)
    mAlerts[" "]=" "
    compInState, compData, err:= t.validateAndGet(stub, args)
    err = json.Unmarshal([]byte(compData), &compStubState)
    if err != nil {
        err = errors.New("Unable to unmarshal JSON data from stub")
      //  fmt.Println(err)
        return nil, err
    }
    
    compStubState.AssetAlerts=mAlerts
    compStubState.Compliance = true
    compStubState.Active = true
    
    compHistKey := compInState.BLNo+SEP+STATE
    err = stub.DelState(compHistKey)
    if err!=nil {
        err = errors.New("Compliance data history delete failed")
	//	fmt.Println(err)
		return nil, err
    }
    compStubJSON, err := json.Marshal(compStubState)
    if err != nil {
        err:=errors.New("Marshaling compliance data failed")
        //fmt.Println(err)
        return nil, err
    }
    err = stub.PutState(compInState.BLNo, compStubJSON)
    if err != nil {
        return nil, errors.New("compliance state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
 } 
  /*********************************  resetComplianceHistory ****************************/
 func (t *SimpleChaincode) readCurrentComplianceState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

     _, compData, err:= t.validateAndGet(stub, args)
     if err !=nil {
         return nil, nil
     }
     return []byte(compData), nil
 }
 /*********************************  resetComplianceHistory ****************************/
 func (t *SimpleChaincode) readComplianceHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

    compInState, _, err:= t.validateAndGet(stub, args)
    compHistKey := compInState.BLNo+SEP+STATE
     comphistory, err := stub.GetState(compHistKey)
    if err != nil {
        return nil, err
    }
     return comphistory, nil
 }
  
 /*********************************  internal: mergePartialState ****************************/	
 func (t *SimpleChaincode) mergePartialState(oldState common.ComplianceState, newState common.ComplianceState) (common.ComplianceState,  error) {
     
    old := reflect.ValueOf(&oldState).Elem()
    new := reflect.ValueOf(&newState).Elem()
    for i := 0; i < old.NumField(); i++ {
        oldOne:=old.Field(i)
        newOne:=new.Field(i)
        if ! reflect.ValueOf(newOne.Interface()).IsNil() {
         //   fmt.Println("New is", newOne.Interface())
         //   fmt.Println("Old is ",oldOne.Interface())
            oldOne.Set(reflect.Value(newOne))
         //   fmt.Println("Updated Old is ",oldOne.Interface())
        } else {
            fmt.Println("Old is ",oldOne.Interface())
        }
    }
    return oldState, nil
 }
 /*********************************  internal: validateAndGet ****************************/	
 func (t *SimpleChaincode) validateAndGet(stub shim.ChaincodeStubInterface, args []string) (common.ComplianceState,  string, error){
     //Assumption: One record at a time.
    var compInState     common.ComplianceState
    //var compStubState   ComplianceState
    var sErr string = ""
    var compstr string
    var err error
    
    if len(args) !=1 {
        err = errors.New("Expecting a single JSON string with mandatory Key, compliance flag, compliance data, active flag and timestamp ")
	//	fmt.Println(err)
		return compInState , sErr, err
        // Though written for a shipping use case, this can easily be explanded to listen to multiple compliance scenarios
	}
    jsonData:=args[0]
    // Marshaling the input to the ComplianceState structure. If a value is not
    // sent in, it is defaulted as null. This allows partial updates.
    compJSON := []byte(jsonData)
  //  fmt.Println("Input asset arg: ", jsonData)
    
    err = json.Unmarshal(compJSON, &compInState)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
	//	fmt.Println(err)
		return compInState , sErr, err
    }
    if strings.TrimSpace(compInState.BLNo) =="" {
        err = errors.New("The Primary Key is mandatory. For example Bill of Lading number for Shipping")
	//	fmt.Println(err)
		return compInState , sErr, err
    }
    
    compData, err := stub.GetState(compInState.BLNo)
    if err !=nil {
        //This implies that this is a new record
        compstr = ""
    } else {
        compstr = string(compData)
    }
    
    return compInState, compstr, nil
 }
