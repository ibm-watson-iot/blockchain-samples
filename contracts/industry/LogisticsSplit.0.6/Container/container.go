package main
import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
    "time"
   // "reflect"
 //  "github.com/mcuadros/go-jsonschema-generator"
"github.com/hyperledger/fabric/core/chaincode/shim"
 "github.com/hyperledger/fabric/core/util"
 common "github.com/hyperledger/fabric/examples/chaincode/go/PoD/Common" 
)


// These are common alerts reported by sensor. Example Tetis. 
// http://www.starcomsystems.com/download/Tetis_ENG.pdf 


// This is a logistics contract, written in the context of shipping. It tracks the progress of a Bill of Lading 
// and associated containers, and raises alerts in case of violations in expected conditions

// Assumption 1. Bill of Lading is sacrosanct - Freight forwarders may issue intermediary freight bills, but
// the original B/L is the document we trackend to end. Similarly a 'Corrected B/L' scenario is not considered

// Assumption 2. A Bill of Lading can have multiple containers attached to it. We are, for simplicity, assuming that
// the same transit rules in terms of allowed ranges in temperature, humidity etc. apply across the B/L - i.e. 
// applies to all containers attached to a Bill of Lading

// Assumption 3. A shipment may switch from one container to another in transit, for various reasons. We are assuming,
// for simplicity, that the same containers are used for end to end transit.


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//Container History
type ContainerHistory struct {
	ContHistory []string `json:"conthistory"`
}
var contractState = common.ContContractState{common.MYVERSION, ""}


//************* main *******************
//Create SimpleChaincode instance
//************* main *******************

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}


// ************************************
// invoke callback mode 
// ************************************
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "createContainerLogistics" {
		// create container records
		return t.createContainerLogistics(stub, args)
    } else if function =="updateContainerLogistics" {
        return t.updateContainerLogistics(stub, args)
    } 
	//fmt.Println("Unknown invocation function: ", function)
	return nil, errors.New("Received unknown invocation: " + function)
}

// ************************************
// query callback mode 
// ************************************
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	// Handle different Query functions 
	if function =="readContainerLogisitcsSchemas" {
        return t.readContainerLogisitcsSchemas(stub, args)
    } else if function =="readContainerCurrentStatus" {
        return t.readContainerCurrentStatus(stub, args)
    } else if function =="readContainerHistory" {
            return t.readContainerHistory(stub, args)
    } else if function == "readAssetSchemas" {
		// returns selected sample objects 
		return t.readAssetSchemas(stub, args)
	}
	return nil, errors.New("Received unknown invocation: " + function)
}                   
// ************************************
// deploy functions 
// ************************************

//************* init *******************
//Chaincode initialization
//************* init *******************

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface,  function string, args []string) ([]byte, error) {
	var stateArg common.ContContractState
	var err error

    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version string and chaincode uuid for the compliance")
    }
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    if stateArg.Version != common.MYVERSION {
        return nil, errors.New("Contract version " + common.MYVERSION + " must match version argument: " + stateArg.Version)
    }
    // set the chaincode uuid of the compliance contract 
    // to the global variable
    
    if stateArg.ComplianceCC =="" {
        return nil, errors.New("Compliance chaincode id is mandatory")
    }
    
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    err = stub.PutState(common.CONTSTATEKEY, contractStateJSON)
    if err != nil {
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
}

// ************************************
// invoke functions 
// ************************************

func (t *SimpleChaincode) createContainerLogistics(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var contInit common.BillOfLadingRegistration
    //var blDefn common.BillOfLadingRegistration
	var err error
    var contState, oldContState common.ContainerLogistics
    //var contHistory ContainerHistory
    
    if len(args) != 1 {
        return nil, errors.New("Expects one argument, a JSON string with bill of lading and container details")
    }
    err = json.Unmarshal([]byte(args[0]), &contInit)
    if err != nil {
        return nil, errors.New("Unable to unmarshal container init " + fmt.Sprint(err))
    }
    // validation again B/L rules can be easily accomplished
    //blDefn = contInit
   // fmt.Println("Max temp: ", contInit.MaxTemperature)
   //  fmt.Println("Min temp: ", contInit.MinTemperature)
   // fmt.Println("Splitting the container list")
    bKey:=contInit.BLNo
    sContainers:=strings.Split(contInit.ContainerNos, ",")
    sTimeStamp:=contInit.Timestamp
    iNos:=len(sContainers)
    mTemp := make(map[string]string, iNos)
    for i := 0; i < len(sContainers); i++ {
       // fmt.Println("Inside container list iteration")
        bCreateContRecord := true // create a new container record
        //This creates a map of [container number] [alerts] with compliance as a true 
        sContKey:=sContainers[i]
	    mTemp[sContKey]= "" // initializing alerts as a blank string
        // Create an initial ContainerLogisitcs record for each container, in the stub
        // Use the container number for the records. If there is an exisitng record,
        //  update it by appending the B/L. 
        sOldContKey:=sContainers[i] // This will be appended with B/L number for old records
       // fmt.Println("Check if old container record exists with the key ", sContKey)
        contData, err := stub.GetState(sContKey)
        if err ==nil && len(contData) >0 {
            //fmt.Println(" This container exists in state, probably used with another B/L")
            // If the container number and B/L number match - unlikely, leave untouched
            err = json.Unmarshal(contData, &oldContState)
            if err != nil {
                //err = errors.New("Unable to unmarshal input JSON data")
                //fmt.Println(err)
                return nil, err
            }
            if oldContState.BLNo != bKey {
                // The container - bill of lading combination does not exist
                // this is the expected case
                sOldContKey = sOldContKey + "_" + oldContState.BLNo 
                // We are going to append the old container record's bill of lading number to the contaienr number
                // This will be the new key for the old record
                 err = stub.PutState(sOldContKey, contData)
                if err != nil {
                    err:=errors.New("re-assigning old container state failed")
                    //fmt.Println(err)
                    return nil, err
                }
               // here bCreateContRecord is true 
            } else {
                // If the bill of lading number is same - shouldnt be - do nothing 
                bCreateContRecord = false   
                // Probably throw an error instead...             
            }
        }
        // If there is no data in the stub for the container, or if there was and we reassigned it,
        // we can now create the container's initial record
        // This is needed, because the container record does not come in with a Bill of Lading.
        // Therefore, we map it here 
        if bCreateContRecord {
           // fmt.Println("Inside 'bCreateContRecord'")
            contState.ContainerNo=sContKey
            contState.BLNo=bKey
            contState.Timestamp = sTimeStamp
            contState.TransitComplete = false
           // fmt.Println("Before constatate marshal")
            contJSON, err := json.Marshal(contState)
           // fmt.Println("After constatate marshal")
            if err != nil {
                err:=errors.New("Marshaling initial container data failed")
                //fmt.Println(err)
                return nil, err
            }
           // fmt.Println("Old container record with different B/L", string(contJSON))
            contHistKey:=contState.ContainerNo+"_HISTORY"
            var contHist = ContainerHistory{make([]string, 1)}
            contHist.ContHistory[0] = string(contJSON)
            contHState, err := json.Marshal(&contHist)
            if err != nil {
                return nil,err
            }
            contHistoryState := []byte(contHState)
            
            err=stub.PutState(sContKey, contJSON)
            if err !=nil {
               // fmt.Println("Unable to create initial container state, ", contJSON)
                return nil, err
            }
            err = stub.PutState(contHistKey, contHistoryState)
            if err != nil {
                return nil, errors.New("container history failed PUT to ledger: " + fmt.Sprint(err))
            } 
           // fmt.Println("New container record generated", string(contJSON))
            // Put Bill Of Lading reg data in the stub to minimize cross-chaincode calls
            err = stub.PutState(bKey, []byte(args[0]))
            if err != nil {
                return nil, errors.New("Bill of Lading data failed PUT to ledger: " + fmt.Sprint(err))
            } 
           // fmt.Println("Bill of Lading rules captured", args[0])
            
        }        
    }
    return nil, nil 
} 
 
 /************************ updateContainerLogistics ********************/

func (t *SimpleChaincode) updateContainerLogistics(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
    var contState, contIn   common.ContainerLogistics
    var compState common.ComplianceState
    var containerHistory ContainerHistory
    var contractState common.ContContractState
    //var oldAlert, newAlert  Alerts
    
    
    // This invoke function is the heart of the logistics contract.
    // The container details come in one by one and are tagged to the bill of lading.
    // They are validated to ensure no violation have happened. 
    //If yes, invoke the shipping contract for notiifcation - on hold now - only the state gets updated today
     if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a single JSON string with mandatory Container Number and optional details")
		//fmt.Println(err)
		return nil, err
	}
	jsonData:=args[0]
    conJSON := []byte(jsonData)
    //fmt.Println("Input Container data arg: ", jsonData)
    
    // Unmarshal imput data into ContainerLogistics struct   
    err = json.Unmarshal(conJSON, &contIn)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
   // fmt.Println(" contIn after unmarshaling [", contIn, "]")  
    
     
     // Container can't be an empty string
     contIn.ContainerNo = strings.TrimSpace(contIn.ContainerNo)
     
     if contIn.ContainerNo=="" {
        err = errors.New("Container number cannot be blank")
        //fmt.Println(err)
        return nil, err
    }
    //fmt.Println("After container number check")
    // Check if an initial definiton has been created for this container number
    // If not, an update shouldn't be allowed since we don't know which B/L its associated with.
    // and what are the parameters under which the shipment is supposed to operate
    // Fetch the record in the stub with the container number
    sContKey := contIn.ContainerNo
    contData, err := stub.GetState(sContKey)
    if err!=nil {
         err = errors.New("Container record not created during registration")
          //  fmt.Println(err)
            return nil, err
    }
   // fmt.Println("After cont state check")
    // This container record has been created in the registration phase
    // or this is not the first container record coming in. 
  //  fmt.Println("contData ", string(contData))
    err = json.Unmarshal(contData, &contState)
    if err != nil {
        err = errors.New("Unable to unmarshal JSON data from stub")
       // fmt.Println(err)
        return nil, err
    }
    
    // Update stub record with new data from the record just read
    // This will be maintained as 'Current State'. This is done because a record may come in 
    // with a partial update. The 'current state' record should be as complete as possible
    // If a property is null in the new record, ignore it, else update the 'state' record
    // using reflection to parse the records and update
    //contState, err =t.mergePartialState(contState,contIn)
    contIn.BLNo = contState.BLNo
    txnTime, err:= stub.GetTxTimestamp()
    if err !=nil {
        err=errors.New("Unable to get transction time")
        return nil, err
    }
    txntimestamp := time.Unix(txnTime.Seconds, int64(txnTime.Nanos))
    sTime := txntimestamp.String()
    
    if strings.TrimSpace(contIn.Timestamp)=="" {
        contIn.Timestamp = sTime
    }
    
   // fmt.Println(" Timestamp is ", sTime)
    
    contState = contIn

  
    
    // Container state never comes in with Bill of Lading number. This gets added here
    // The state record already has it, now we add it to the incoming record
    // This is used in the alertsCheck call
  //  fmt.Println("B/L number in container state is ", contState.BLNo)
    blKey :=contState.BLNo
    contIn.BLNo = blKey
  //  fmt.Println("B/L number in container in is ", contIn.BLNo)

    
  //  fmt.Println("Perform a compliance check on the new record")
    newAlerts, err:= t.alertsCheck(stub, contIn)
    sAlerts := string(newAlerts)
  //  fmt.Println("Alerts data is : ", string(newAlerts))
    if len(sAlerts)>0 {
        // This implies a compliance violation. 
     //   fmt.Println("Update container record with new alert status pertaining to that record")
        
        contState.AlertRecord = sAlerts
        
        // call the compliance contract to maintain the state
        //*************************************************************************
    // Invoke the compliance contract to create and maintain the B/L compliance
        //get compliance contract uuid from the stub
        contractStateJSON, err := stub.GetState(common.CONTSTATEKEY)
        if err != nil {
            return nil,errors.New("Unable to fetch container and compliance contract keys")
        }
        err = json.Unmarshal(contractStateJSON, &contractState)
        if err != nil {
            return nil, err
        }
        complianceChainCode:=contractState.ComplianceCC
        f := "createUpdateComplianceRecord"
        compState.BLNo = blKey
        compState.Type = "SHIPPING"
        sTimeStamp:=contState.Timestamp 
        sCompAlerts:= sAlerts
        mAlerts:= make(map[string]string)
        mAlerts[contState.ContainerNo]=sCompAlerts
        compState.AssetAlerts=mAlerts // every alert even triggers a ui action. cumulation not needed. history maintained
        compState.Timestamp = sTimeStamp
        compState.Compliance =false
        compState.Active=true
        compJSON, err := json.Marshal(compState)
        if err != nil {
            err:=errors.New("Marshaling compliance initialization failed")
            //fmt.Println(err)
            return nil, err
        }
        
        //////////////////////
        invokeArgs := string(compJSON)
        var callArgs = make([]string, 0)
        callArgs = append(callArgs, f)
        callArgs = append(callArgs, invokeArgs)
      
        _, err = stub.InvokeChaincode(complianceChainCode, util.ToChaincodeArgs(callArgs...))
        if err != nil {
            errStr := fmt.Sprintf("Failed to invoke chaincode. Got error: %s", err.Error())
           // fmt.Printf(errStr)
            return nil, errors.New(errStr)
        }
     //*************************************************************************
       
    }
    
    updContJSON, err := json.Marshal(contState)
    if err != nil {
        err:=errors.New("Marshaling container data failed")
        //fmt.Println(err)
        return nil, err
    }
    
    // Now updated container state data
    //fmt.Printf("Putting updated container state data %s to ledger\n", string(updContJSON))
    err = stub.PutState(sContKey, updContJSON)
    if err != nil {
        err:=errors.New("Writing updated container state data to the ledger failed")
       // fmt.Println(err)
        return nil, err
    }
    contHistKey:=contIn.ContainerNo+"_HISTORY"
    var contSlice = make([]string, 0)
    contSlice = append(contSlice, jsonData)
    contSlice = append(contSlice, containerHistory.ContHistory...)
    containerHistory.ContHistory = contSlice
    
    contHState, err := json.Marshal(&containerHistory)
    if err != nil {
        return nil,err
    }
    contHistoryState:= []byte(contHState)
    err = stub.PutState(contHistKey, contHistoryState)
    if err != nil {
        return nil, errors.New("Container history updatefailed PUT to ledger: " + fmt.Sprint(err))
    } 
    
    // Check if lat-long is in notification range
    // Not implementing notification till clarity from IRL Side
    // later, implement it to call the shipping contract
   // fmt.Printf("Container %s state successfully written to ledger : %s\n", sContKey, string(updContJSON))
    return nil, nil
 
}


// ************************************
// query functions 
// ************************************


// ************************************
// getContainerLogisitcsSchema
// ************************************
// This is a 'convenience' function, to provide the consumer of a contract an example of 
// the Bill of Lading definition dataset.
func (t *SimpleChaincode) readContainerLogisitcsSamples(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    cont := []byte(`{ "ContainerNo":"MSKU000000","Location" :{"Latitude":10, "Longitude":15}, "Temperature":2, 
        "Carrier":"Carrier", "Timestamp":"2016-03-03T20:27:23.969676659Z", "Humidity":15, "Light":5,
    "DoorClosed":true, "Acceleration":0}`)
    return cont, nil
}

// ************************************
// getContainerCurrentStatus
// ************************************
// This returns the container data
func (t *SimpleChaincode) readContainerCurrentStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var err error
    var  contIn   common.ContainerLogistics
    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a single JSON string with mandatory Container Number")
	//	fmt.Println(err)
		return nil, err
	}
	jsonData:=args[0]
    conJSON := []byte(jsonData)
   // fmt.Println("Input Container data arg: ", jsonData)
    
    // Unmarshal imput data into ContainerLogistics struct   
    err = json.Unmarshal(conJSON, &contIn)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
   // fmt.Println(" contIn after unmarshaling [", contIn, "]")        
     
     // Container can't be an empty string
     contIn.ContainerNo = strings.TrimSpace(contIn.ContainerNo)
     
     if contIn.ContainerNo=="" {
        err = errors.New("Container number cannot be blank")
       // fmt.Println(err)
        return nil, err
    }
    
    contData, err := stub.GetState(contIn.ContainerNo)
    if err!=nil {
         err = errors.New("Container record not available")
          // fmt.Println(err)
            return nil, err
    }
    return contData, nil
}
func (t *SimpleChaincode) readContainerLogisitcsSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(schemas), nil
}
/*********************************  resetContainerHistory ****************************/
 func (t *SimpleChaincode) readContainerHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var  contIn   common.ContainerLogistics
    //var contHist  ContainerHistory
    var err error
    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a single JSON string with mandatory Container Number")
		//fmt.Println(err)
		return nil, err
	}
	jsonData:=args[0]
    conJSON := []byte(jsonData)
   // fmt.Println("Input Container data arg: ", jsonData)
    
    // Unmarshal imput data into ContainerLogistics struct   
    err = json.Unmarshal(conJSON, &contIn)
    if err != nil {
        //err = errors.New("Unable to unmarshal input JSON data")
		//fmt.Println(err)
		return nil, err
    }
   // fmt.Println(" contIn after unmarshaling [", contIn, "]")        
     if contIn.ContainerNo==""{
       //  fmt.Println(" Container number is blank")
        err = errors.New("Container number is mandatory")
       // fmt.Println(err)
		return nil, err
     }
     
    contHistKey := contIn.ContainerNo+"_HISTORY"
    conthistory, err := stub.GetState(contHistKey   )
    if err != nil {
        return nil, err
    }
     return conthistory, nil
 }

func (t *SimpleChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	return []byte(schemas), nil
}


// ************************************
// alertsCheck
// ************************************
// This is an 'internal' function, to check for alert
func (t *SimpleChaincode) alertsCheck(stub shim.ChaincodeStubInterface, contIn common.ContainerLogistics) ([]byte,  error) {
    // I will rework thisd - possibly with reflection
    var blDefn common.BillOfLadingRegistration
    var blReg common.BillOfLadingRegistration 
    //alerts will not get raised: blReg doesn't exist today in state
    // Expensive?

    var alert =new(common.Alerts)
    var contAlert []byte
    var val common.Variation
    
    bKey:= contIn.BLNo
   // fmt.Println("B/L number  inside alertscheck is ",bKey)
    complianceAlert := false
    //  use value in global variable to check alerts compliance.
    blData, err := stub.GetState(bKey)
    if err!=nil {
        err:=errors.New("Unable to retrieve Bill of Lading data from the stub")
       // fmt.Println(err)
        return nil, err
    }
     err = json.Unmarshal(blData, &blDefn)
    if err != nil {
        return nil, errors.New("Bill of Lading record unmarshal failed: " + fmt.Sprint(err))
    }
    blReg = blDefn
    //Temperature check
    var actVal = contIn.Temperature

    val, _ = t.inRange(blReg.MinTemperature, blReg.MaxTemperature, actVal)
    if (val !=common.Normal) {
        alert.TempAlert = val
        complianceAlert = true
    }
     cAlert, err := json.Marshal(&alert)
       // fmt.Println(string(cAlert))
    //Humidity Check
    actVal = contIn.Humidity

    val,_ = t.inRange(blReg.MinHumidity, blReg.MaxHumidity, actVal)
    if (val !=common.Normal) {
        alert.HumAlert = val
        complianceAlert = true
    }

     cAlert, err = json.Marshal(&alert)
       // fmt.Println(string(cAlert))
    //Light Check
    actVal = contIn.Light

    val,_ = t.inRange(blReg.MinLight, blReg.MaxLight, actVal)
    if (val !=common.Normal) {
        alert.LightAlert = val
        complianceAlert = true
    }
    
    // Acceleration check
    actVal = contIn.Acceleration
    
    val,_ = t.inRange(blReg.MinAcceleration, blReg.MaxAcceleration, actVal)
    if (val !=common.Normal) {
        alert.AccAlert = val
        complianceAlert = true
     }
    

    if contIn.DoorClosed == false {
        alert.DoorAlert =true
        complianceAlert = true
    }
        

    if complianceAlert {
        cAlert, err = json.Marshal(&alert)
       // fmt.Println(string(cAlert))
        if err !=nil {
            err = errors.New("Unable to marshal alert data")
		   // fmt.Println(err)
		    return nil, err
        }
        contAlert = cAlert
    } 
      return contAlert, nil
}

/*********************************  internal: inRange ****************************/	
 func (t *SimpleChaincode) inRange(minVal float64, maxVal float64, actVal float64) (common.Variation,  error) {
    var val common.Variation 

    if actVal > maxVal {
      //  fmt.Println(actVal)
        val =common.Above
    } else if actVal < minVal {
        val =common.Below
    } else {
    val = common.Normal
    }
    return val, nil
 }
