package main
import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
    "time"
 
    "github.com/op/go-logging"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/hyperledger/fabric/core/util"
    common "github.com/hyperledger/fabric/examples/chaincode/go/PoD/Common" 

)

var myLogger = logging.MustGetLogger("logistics")
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var contractState = common.BLContractState{common.MYVERSION, "",""}

const blContFunctionCall string = "createContainerLogistics"
const blCompFunctionCall string = "createUpdateComplianceRecord"

//************* main *******************
//Create SimpleChaincode instance
//************* main *******************

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}

//************* init *******************
//Chaincode initialization
//************* init *******************
/*
The Init function initializes the contract ids for container and compliance contracts
*/
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var stateArg common.BLContractState
	var err error
    myLogger.Info("[LogisticsChaincode-BillOfLading] Init")
    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version string, container and compliance chaincode uuids")
    }
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    if stateArg.Version != common.MYVERSION {
        return nil, errors.New("Contract version " + common.MYVERSION + " must match version argument: " + stateArg.Version)
    }
    
    //fmt.Println("chaincodes assigned")
    if stateArg.ContainerCC=="" || stateArg.ComplianceCC =="" {
        return nil, errors.New("Container and compliance chaincode ids are mandatory")
    }
    //fmt.Println("complianceChainCode ", complianceChainCode)
    //fmt.Println("containerChainCode ", containerChainCode)
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    
    err = stub.PutState(common.BLSTATEKEY, contractStateJSON)
    if err != nil {
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
}

// ************************************
// invoke callback mode 
// ************************************
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different functions
    switch function {
        case "registerBillOfLading" :
            return t.registerBillOfLading(stub, args)
        case "deregisterBillOfLading" :
            return t.deregisterBillOfLading(stub, args)
        default:
            return nil, errors.New("Unknown function call to compliance : invoke")
    }
}

// ************************************
// query callback mode 
// ************************************
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different Query functions 
    switch function {
        case "getBillOfLadingRegistration" :
            return t.getBillOfLadingRegistration(stub, args)
        case "getBillOfLadingRegistrationSchema" :
            return t.getBillOfLadingRegistrationSchema(stub, args)
        default:
            return nil, errors.New("Unknown function call to compliance : invoke")
    }
}                   

// ***********************registerBillOfLading************************

func (t *SimpleChaincode) registerBillOfLading(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
	var err error
    var blReg common.BillOfLadingRegistration
    var compState common.ComplianceState
    var contractState common.BLContractState

    // This is where the initial regestration of the Bill of Lading takes place. It would correspond to a 
    // Bill of lading being generated in shipping and contains list of associated containers
     
    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a single JSON string with mandatory BillofLading, Container Numbers and Hazmat")
		//fmt.Println(err)
		return nil, err
	}
    jsonData:=args[0]
    
    // Marshaling the input to the BillOfLadingRegistration structure. If a value is not
    // sent in, it is defaulted as null
    initJSON := []byte(jsonData)
    //fmt.Println("Input asset arg: ", jsonData)
    
    err = json.Unmarshal(initJSON, &blReg)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
    //fmt.Println(" blDef after unmarshaling [", blDef, "]") 
   
    // fmt.Println(" Check if Bill of Lading and container numbers have been sent in correctly")
     blReg.BLNo = strings.TrimSpace(blReg.BLNo)
     blReg.ContainerNos = strings.TrimSpace(blReg.ContainerNos)
     if blReg.BLNo=="" || blReg.ContainerNos =="" {
        err = errors.New("Bill of Lading  / Container numbers cannot be blank")
        fmt.Println(err)
        return nil, err
    }
     //fmt.Println(" After checking blank")
     // Implementing the transaction timestamp feature
    //Transaction id can also be obtained as  stub.UUID
    // Could be leveraged in custom Event listener being planned
    txnTime, err:= stub.GetTxTimestamp()
    if err !=nil {
        err=errors.New("Unable to get transction time")
        return nil, err
    }
    txntimestamp := time.Unix(txnTime.Seconds, int64(txnTime.Nanos))
    sTime := txntimestamp.String()
    
    if strings.TrimSpace(blReg.Timestamp)=="" {
        blReg.Timestamp = sTime
    }
    
     //fmt.Println(" Timestamp is ", sTime)
    bKey:=blReg.BLNo                       // Bill of lading record - registration key
   
   // Prepare the record to put Bill of Lading Registration information in the stub
    //fmt.Println("First check if it exists in state")
    blStub, err := stub.GetState(bKey)
    if err == nil && (len(blStub)>0) {
        //fmt.Println("blStub ", string(blStub))
        err = errors.New("You cannot create an existing Bill of Lading record: "+ bKey)          
        //fmt.Println(err)
		return nil, err
    }

    //If the mandatory fields of Bill of Lading Number, Container Number and Hazmat are satisfactory,
    // and the B/L doesn't already exist in the stub we can put the B/L details into the stub.
    // Set TransitComplete to false
    
    blReg.TransitComplete = false
    // Marshal back to a JSON string which will be stored in the stub
    regJSON, err := json.Marshal(blReg)
    if err != nil {
        err:=errors.New("Marshaling bill of lading data in registration failed")
        //fmt.Println(err)
        return nil, err
    }
    // Before we put the registration data on the stub, let's create an instance of the state record
    // We won't do an unnecessary stub read for state, since the bill of lading registration doesn't exist
    // If for some unexpected reason it does exist, we don't care -it is invalid and will get overwritten
    
    //fmt.Println("before calling container logistics")
    //*************************************************************************
    // Invoke the container logistics contract to create the container's initial state
    
    f := blContFunctionCall
    // Fetch container and compliance contract ids from the ledger
    contractStateJson, err := stub.GetState(common.BLSTATEKEY)
    if err != nil {
        return nil,errors.New("Unable to fetch container and compliance contract keys")
    }
    err = json.Unmarshal(contractStateJson, &contractState)
    if err != nil {
        return nil, err
    }
    containerChainCode:=contractState.ContainerCC
    complianceChainCode:=contractState.ComplianceCC
    //fmt.Println("containerChainCode ", containerChainCode)
    //fmt.Println("complianceChainCode ", complianceChainCode)
	var invokeArgs = make([]string, 0) 
    invokeArgs = append(invokeArgs, f)
    invokeArgs = append(invokeArgs, string(regJSON))
	_, err = stub.InvokeChaincode(containerChainCode, util.ToChaincodeArgs(invokeArgs...))
	if err != nil {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
		fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
    //fmt.Println("before calling compliance")
     //*************************************************************************
    // Invoke the compliance contract to create and maintain the B/L compliance
    f = blCompFunctionCall
    compState.BLNo = bKey
    //fmt.Println("compState.BLNo", compState.BLNo)
    compState.Type = "SHIPPING"
    //fmt.Println("compState.Type", compState.Type)
    compState.Timestamp = blReg.Timestamp 
    //fmt.Println("compState.Timestamp", compState.Timestamp)
    compState.Compliance = true
    //fmt.Println("compState.Compliance", compState.Compliance)
    compState.Active=true
   // fmt.Println("compState.Active", compState.Active)
    compJSON, err := json.Marshal(compState)
    if err != nil {
        err:=errors.New("Marshaling compliance initialization failed")
        //fmt.Println(err)
        return nil, err
    }
	//fmt.Println("marshal success")
    invokeArgs = make([]string, 0) 
    invokeArgs = append(invokeArgs, f)
    invokeArgs = append(invokeArgs, string(compJSON))
    _, err = stub.InvokeChaincode(complianceChainCode, util.ToChaincodeArgs(invokeArgs...))
	if err != nil {
		errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
		//fmt.Printf(errStr)
		return nil, errors.New(errStr)
	}
     //*************************************************************************
    //fmt.Println("After chaincode calls")   
    // Now, lets put the Registration information in the stub

   // fmt.Printf("Putting new bill of lading registration data %s to ledger\n", string(regJSON))
    err = stub.PutState(bKey, regJSON)
    if err != nil {
        err:=errors.New("Writing bill of lading registration data to the ledger failed")
        //fmt.Println(err)
        return nil, err
    }
	return nil, nil
}

func (t *SimpleChaincode) deregisterBillOfLading(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
    //call fetchBlData and update the flags to false.
    // Ideally update container and compliance contracts. Not required today because container contract
    // will automatically update
    var blReg common.BillOfLadingRegistration
    blRegData, err:=t.fetchBLData( stub, args)
    if err !=nil{
        return nil, err     
    }
    err = json.Unmarshal(blRegData, &blReg)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
   //Updating the timestamp 

    txnTime, err:= stub.GetTxTimestamp()
    if err !=nil {
        err=errors.New("Unable to get transction time")
        return nil, err
    }
    txntimestamp := time.Unix(txnTime.Seconds, int64(txnTime.Nanos))
    sTime := txntimestamp.String()
    if strings.TrimSpace(blReg.Timestamp)=="" {
        blReg.Timestamp = sTime
    }

    blReg.TransitComplete=true;
    regJSON, err := json.Marshal(blReg)
    if err != nil {
        err:=errors.New("Marshaling bill of lading data in deregistration failed")
        //fmt.Println(err)
        return nil, err
    }
     bKey:=blReg.BLNo
    err = stub.PutState(bKey, regJSON)
    if err != nil {
        err:=errors.New("Updating bill of lading registration data in the ledger failed")
        //fmt.Println(err)
        return nil, err
    }
    return nil, nil
    
}

// ************************************
// query functions 
// ************************************

// ***********************getBillOfLadingRegistration************************
    
// ************************************
// getBillOfLadingRegistrationSchema
// ************************************
// This is a 'convenience' function, to provide the consumer of a contract an example of 
// the Bill of Lading definition dataset.
func (t *SimpleChaincode) getBillOfLadingRegistrationSchema(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
    // Temporarily hardcoded. This is not at present intended to be JSON compatible
    // Can be explanded later to use a combination of ast and default definitions
    bl := []byte (`{ "BLNo": "0000000000", "ContainerNos" : "MSKU000000, MRSK000000",  "Hazmat"  : false,
     "MinTemperature" : -20.00,  "MaxTemperature" : 0.00,   "MinHumidity" : 20.00,  "MaxHumidity" : 50.00,  
     "MinLight" : 0.00,   "MaxLight" : 100.00, "MinAcceleration" : 0.001,  "MaxAcceleration" : 1.9  }`)
      // Notify-aspects are not represented at present. Will be added later
      // Will be replaced by the schema implementation later for consumption by the UI
	return bl, nil
}



// ************************************
// getBillOfLadingRegistration
// ************************************
// This returns the actual Bill of Lading registration dataset.

func (t *SimpleChaincode) getBillOfLadingRegistration(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var err error
    blRegData, err:=t.fetchBLData( stub, args)
    if err !=nil{
        return nil, err     
    }
    return blRegData, nil
}
// ************************************
// getBillOfLadingState
// ************************************
// This returns the actual Bill of Lading state dataset.


// ************************************
// pullBLData 
// ************************************
// internal utiltiy function

func (t *SimpleChaincode) fetchBLData(stub shim.ChaincodeStubInterface,  args []string) ([]byte, error) {
    var qKey string
    var err error
    var blReg common.BillOfLadingRegistration
    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a single JSON string with mandatory BillofLading")
		//fmt.Println(err)
		return nil, err
	}
    jsonData:=args[0]
    
    // Marshaling the input to the BillOfLadingRegistration structure. If a value is not
    // sent in, it is defaulted as null
    initJSON := []byte(jsonData)
    //fmt.Println("Input asset arg: ", jsonData)
    
    err = json.Unmarshal(initJSON, &blReg)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
   // fmt.Println(" blDef after unmarshaling [", blDef, "]") 
   
   //  Nil check for mandatory fields. Since these are defined as pointers in the
   // struct definition, they will be unmarshalled as null (json) / nil (golang)      
  
   //  fmt.Println(" Trimming blanks out")
   
   // Check if Bill of Lading and container numbers have been sent in correctly
     blReg.BLNo = strings.TrimSpace(blReg.BLNo)
     
     if blReg.BLNo=="" {
        err = errors.New("Bill of Lading cannot be blank")
        //fmt.Println(err)
        return nil, err
    }
    
   
    qKey=blReg.BLNo                       // Bill of lading record - registration 
    
    blData, err := stub.GetState(qKey)
    if err!=nil {
        err:=errors.New("Unable to retrieve Bill of Lading data from the stub")
        //fmt.Println(err)
        return nil, err
    }
    // if it was a reg call, it will be reg info
    // if state call, state info is returned
    //fmt.Println(string(blData)) 
    return blData, nil
}
