package main
import (
    "encoding/json"
    "errors"
    "fmt"
    "time"
  //  "strings"
  //  "reflect"
 //  "github.com/mcuadros/go-jsonschema-generator"
"github.com/hyperledger/fabric/core/chaincode/shim"
// "github.com/hyperledger/fabric/core/util"
 
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


// DeviceUsageChaincode example simple Chaincode implementation
type DeviceUsageChaincode struct {
}


//************* main *******************
//Create DeviceUsageChaincode instance
//************* main *******************

func main() {
	err := shim.Start(new(DeviceUsageChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}


// ************************************
// invoke callback mode 
// ************************************
func (t *DeviceUsageChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "createDevice" {
		// create container records
		return t.createDevice(stub, args)
    } else if function =="createUsage" {
        return t.createUsage(stub, args)
    }else if function =="extendUsage" {
        return t.extendUsage(stub, args)
    } else if function =="updateDeviceAsAvailable" {
        return t.updateDeviceAsAvailable(stub, args)
    } 
	//fmt.Println("Unknown invocation function: ", function)
	return nil, errors.New("Received unknown invocation: " + function)
}

// ************************************
// query callback mode 
// ************************************
func (t *DeviceUsageChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	// Handle different Query functions 
	if function =="readDevice" {
        return t.readDevice(stub, args)
    } else if function =="readUsage" {
        return t.readUsage(stub, args)
    } else if function =="readUsageHistory" {
            return t.readUsageHistory(stub, args)
    } else if function =="readDeviceList" {
            return t.readDeviceList(stub, args)
    } else if function =="readAssetSchemas" {
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

func (t *DeviceUsageChaincode) Init(stub *shim.ChaincodeStub,  function string, args []string) ([]byte, error) {
	var stateArg ContractState
	var err error

    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with tagged version string")
    }
    //fmt.Println(args[0])
    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
    }
    if stateArg.Version != MYVERSION {
        return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
    }
    // set the chaincode uuid of the compliance contract 
    // to the global variable

    
    contractStateJSON, err := json.Marshal(stateArg)
    if err != nil {
        return nil, errors.New("Marshal failed for contract state" + fmt.Sprint(err))
    }
    err = stub.PutState(CONTSTATEKEY, contractStateJSON)
    if err != nil {
        return nil, errors.New("Contract state failed PUT to ledger: " + fmt.Sprint(err))
    }
    return nil, nil
}

// ************************************
// invoke functions 
// ************************************

func (t *DeviceUsageChaincode) createDevice(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var devicesIn []Device
    //var devicesStub []Device
    var device Device
	var err error
    var sError string = ""
    var sStubUpdate []byte
    var devList DevList
    //var deviceList []string
    var devSlice []string
    var iCount int = 0
    
    var devListData []byte
    if len(args) != 1 {
        return nil, errors.New("Expects one argument, a JSON array with one or more device details")
    }
    err = json.Unmarshal([]byte(args[0]), &devicesIn)
    if err != nil {
        return nil, errors.New("Unable to unmarshal device data " + fmt.Sprint(err))
    }
    // Check if device data already exists in the stub

    for l := range devicesIn {
        //Assuming creteonupdate
        sDeviceId:=*devicesIn[l].DeviceID 
        sDeviceKey:=DEVICESKEY+"_"+sDeviceId
        sDevListKey:=LISTKEY
        stubData, err := stub.GetState(sDeviceKey)
        if err ==nil && len(stubData) >0 {
            // This device is being updated
            //If it is in use, it shouldn't be
            err = json.Unmarshal(stubData, &device)
            if err != nil {
                //fmt.Println(err)
                return nil, err
            }
            if *device.Available {
                //Compare and update
                 if devicesIn[l].MinimumUsageCost !=nil {
                    *device.MinimumUsageCost=*devicesIn[l].MinimumUsageCost
                 }
                if devicesIn[l].MinimumUsageTime !=nil {
                    *device.MinimumUsageTime=*devicesIn[l].MinimumUsageTime
                }
                if devicesIn[l].OvertimeUsageCost !=nil {
                    *device.OvertimeUsageCost=*devicesIn[l].OvertimeUsageCost
                }
                if devicesIn[l].OvertimeUsageTime !=nil {
                    *device.OvertimeUsageTime=*devicesIn[l].OvertimeUsageTime
                }
                sStubUpdate, err = json.Marshal(&devicesIn[l])
                if err !=nil {
                    return nil, err
                }
            } else {
                sError +="Device unavailable for update:  "+sDeviceId
                sStubUpdate = []byte("")

            }
        } else {
             // If device does not exist in stub, this is a new entry
            
             sStubUpdate, err = json.Marshal(&devicesIn[l])

            if err !=nil {
                return nil, err
            }
        }
        if iCount ==0 {
            devSlice = make([]string, 0)
            iCount++
        }
        //Get device list from stub
        
        devSlice = append(devSlice, sDeviceId)
        //fmt.Println("Added to list")
            
        if len(string(sStubUpdate))>0 {
            err = stub.PutState(sDeviceKey,sStubUpdate)
            if err != nil {
                return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
            }
        }
        sListData, err:= stub.GetState(sDevListKey)
        if err!=nil {
            err = json.Unmarshal(sListData, &devList)
            if err != nil {
                return nil, err
            }
            //fmt.Println("Added history to list")
            devSlice = append(devSlice, devList.Devices...)
        }
        devList.Devices = devSlice
        devListData, err = json.Marshal(&devList)
        err = stub.PutState(sDevListKey,devListData)
        //fmt.Println("Data to be put in stub is ", string(devListData))
        if err != nil {
            return nil, errors.New("Usage record failed PUT to ledger: " + fmt.Sprint(err))
        }
       
    }
    if len(sError)>0 {
        return nil, errors.New(sError)
    } 
    return nil, nil
}
/************************ createUsage ********************/
func (t *DeviceUsageChaincode) updateDeviceAsAvailable(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
   
    var device Device
    var deviceIn Device
	var err error
    var sStubUpdate []byte
    
    if len(args) != 1 {
        return nil, errors.New("Expects one argument, a JSON string  with  device details")
    }
    err = json.Unmarshal([]byte(args[0]), &deviceIn)
    if err != nil {
        return nil, errors.New("Unable to unmarshal device data " + fmt.Sprint(err))
    }
    // Check if device data already exists in the stub

    sDeviceId:=*deviceIn.DeviceID 
    sDeviceKey:=DEVICESKEY+"_"+sDeviceId
    stubData, err := stub.GetState(sDeviceKey)
    if err ==nil && len(stubData) >0 {
        // This device is being updated
        // to make it available
        err = json.Unmarshal(stubData, &device)
        if err != nil {
            //fmt.Println(err)
            return nil, err
        }
        *device.Available = true;

        // If device does not exist in stub, this is a new entry
        sStubUpdate, err = json.Marshal(&device)
        if err !=nil {
            return nil, err
        }

        if len(string(sStubUpdate))>0 {
            err = stub.PutState(sDeviceKey,sStubUpdate)
            if err != nil {
                return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
            }
        }
    } else {
        return nil, errors.New("Device record does not exist! " + fmt.Sprint(err))

    }
    return nil, nil
}
 /************************ createUsage ********************/


func (t *DeviceUsageChaincode) createUsage(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var usage Usage
    var device Device
	var err error
    var usageHistory UsageHistory
    var usageHistoryState []byte
    var sAlert AlertLevels
    var sAlertData string
    if len(args) != 1 {
        return nil, errors.New("Expects one argument, a JSON string with device usage details")
    }
     //fmt.Println("Inside createUsage", args[0])
    err = json.Unmarshal([]byte(args[0]), &usage)
    if err != nil {
        return nil, errors.New("Unable to unmarshal device usage data " + fmt.Sprint(err))
    } 
    if usage.Duration==0 {
        return nil, errors.New("Cannot block device for 0 minutes " + fmt.Sprint(err))
    }
    /////////////////////////////////////////////////
    // Create device key for usage stub interactions
    sDeviceId:=usage.DeviceID 
    sDeviceKey:=DEVICESKEY+"_"+sDeviceId
    sUsageKey:=USAGEKEY+"_"+sDeviceId
    stubData, err:= stub.GetState(sDeviceKey)
    if err!=nil {
        return nil, errors.New("Device does not exist in stub!")
    }
    //fmt.Println("Setting avail flag to false")
    /////////////////////////////////////////////////
    //Set the Available flag for the device to false
    err = json.Unmarshal(stubData, &device)
    if err != nil {
        //fmt.Println(err)
        return nil, err
    }
    *device.Available = false
     sStubUpdate, err := json.Marshal(&device)
    if err !=nil {
        return nil, err
    }
    //fmt.Println("PUSHING DEVICE RECORD")
    /////////////////////////////////////////////////
    // Push the updated device data with Available flag set to false, to the stub
    err = stub.PutState(sDeviceKey,sStubUpdate)
    if err != nil {
        return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
    }
    /////////////////////////////////////////////////
    // Assuming start time and duration are provided by the invoke call, calculate end time
    //fmt.Println("calculating time")
    sStartTime := usage.StartTime
    //fmt.Println("The start time is ", sStartTime)

    tEndTime, err := time.Parse("2006-01-02 15:04:05", sStartTime)
    if err != nil {
        //fmt.Println(err)
        return nil, err
    }
    iDuration := time.Duration(usage.Duration)*time.Minute
    dEndTime:= tEndTime.Add(iDuration)
    usage.EndTime = dEndTime.String()
    //Computing usage cost
    usage.UsageCost = *device.MinimumUsageCost* float64(usage.Duration)
    usage.TotalCost = usage.UsageCost
    // Put usage record to state
    sStubUpdate, err = json.Marshal(&usage)
    if err !=nil {
        return nil, err
    }
    //fmt.Println("Usage Key is ", sUsageKey)
    /////////////////////////////////////////////////
    // Put the usage data back into the stub
    err = stub.PutState(sUsageKey,sStubUpdate)
    //fmt.Println(string(sStubUpdate))
    if err != nil {
        return nil, errors.New("Usage record failed PUT to ledger: " + fmt.Sprint(err))
    }
    ///////////////////////////////////////////////
    // Update History
    usageHistKey:=sDeviceId+"_HIST"
    stubHistData, err:= stub.GetState(usageHistKey)
    if err!=nil {
        //fmt.Println("No previous history for device")
        var usageHist = UsageHistory{make([]string, 1)}
        usageHist.History[0] = string(sStubUpdate)
        usageHState, err := json.Marshal(&usageHist)
        if err != nil {
            return nil,err
        }
        usageHistoryState = []byte(usageHState)
        //fmt.Println(string(usageHistoryState))
    } else {
        //fmt.Println("Previous usage history exists for device")
        //fmt.Println("History in stub is ", string(stubHistData))
        var usageHistSlice []string = make([]string, 0)
        usageHistSlice = append(usageHistSlice, string(sStubUpdate))
        if len(stubHistData) > 0 {
            err = json.Unmarshal(stubHistData, &usageHistory)
            if err != nil {
                return nil, err
            }
            if len(usageHistory.History) >= MAXHIST {
                    usageHistory.History = usageHistory.History[:len(usageHistory.History)-1]
            }
            usageHistSlice = append(usageHistSlice, usageHistory.History...)
        }
        usageHistory.History = usageHistSlice
        usageHState, err := json.Marshal(&usageHistory)
        if err != nil {
            return nil,err
        }
        usageHistoryState = []byte(usageHState)
    }
        //fmt.Println("Setting usage history to state: ", string(usageHistoryState))
        err=stub.PutState(usageHistKey, usageHistoryState)
        if err !=nil {
            //fmt.Println("Unable to create usage history ")
            return nil, err
        }
    /////////////////////////////////////////////////
    // Alerts processing
    //Generate Alert Key of device for stub interactions
    //Alert is raised when the timer reaches  the total duration
    // Can be modified later
    // Timer and sleep create new channel. This was checked with Binh over slack
    // But apparently it is messes with the chaincode
    //fmt.Println("Before sleep")
    //time.Sleep(20 * time.Second)
    //fmt.Println("After sleep")
    // Once sleep is done, push the alert into the stub
    
    
    sAlertsKey:=ALERTKEY+"_"+sDeviceId
    
    sAlert = Warning
    sAlertData = "{\""+sDeviceId+"\": \""+string(sAlert)+"\"}"
    fmt.Println(sAlertData)
    err = stub.PutState(sAlertsKey,[]byte(sAlertData))
    if err != nil {
        return nil, errors.New(" Alert failed PUT to ledger: " + fmt.Sprint(err))
    }  
    
    /*
    *device.Available = true
    sStubUpdate, err = json.Marshal(&device)
    if err !=nil {
        return nil, err
    }
    err = stub.PutState(sDeviceKey,sStubUpdate)
    if err != nil {
        return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
    }  */
    //fmt.Println("Create Usage Over")
    return nil, nil
}

/************************ readAssetSchemas ********************/

func (t *DeviceUsageChaincode) readDevice(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var device Device
    if len(args) != 1 {
        return nil, errors.New("readDevice expects one argument, a JSON string with device id")
    }
    //fmt.Println(args[0])
    err := json.Unmarshal([]byte(args[0]), &device)
    if err!=nil {
         return nil, errors.New("Unable to unmarshal input device!")
    } else {
        if device.DeviceID!=nil {
         sDeviceId:=*device.DeviceID 
         sDeviceKey:=DEVICESKEY+"_"+sDeviceId
         stubData, err:= stub.GetState(sDeviceKey)
         if err!=nil {
             return nil, errors.New("Unable to get device from stub !")
         } else {
             //fmt.Println("Device data in stub is :", string(stubData))
            return stubData, nil
         }
         
        } else {
            return nil, errors.New("Device id is mandatory !")
        }
    } 
}
/************************ readUsage ********************/

func (t *DeviceUsageChaincode) readUsage(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var usage Usage
    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with device id")
    }
    //fmt.Println(args[0])
    err := json.Unmarshal([]byte(args[0]), &usage)
    if err!=nil {
         return nil, errors.New("Unable to unmarshal input device!")
    } 
    if len(usage.DeviceID)>0 {
        sDeviceId:=usage.DeviceID 
        sUsageKey:=USAGEKEY+"_"+sDeviceId
        //fmt.Println("Usage Key is ", sUsageKey)
        stubData, err:= stub.GetState(sUsageKey)
        if err!=nil {
            return nil, errors.New("Unable to get device from stub !")
        } else {
            //fmt.Println("Device data in stub is :", string(stubData))
        return stubData, nil
        }
        
    } else {
        return nil, errors.New("Device id is mandatory !")
    }
    return nil, nil
}


/************************ readUsageHistory ********************/

func (t *DeviceUsageChaincode) readUsageHistory(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var usage Usage
    if len(args) != 1 {
        return nil, errors.New("init expects one argument, a JSON string with device id")
    }
    //fmt.Println(args[0])
    err := json.Unmarshal([]byte(args[0]), &usage)
    if err!=nil {
         return nil, errors.New("Unable to unmarshal input device!")
    } 
    if len(usage.DeviceID)>0 {
        sDeviceId:=usage.DeviceID 
         usageHistKey:=sDeviceId+"_HIST"
        stubData, err:= stub.GetState(usageHistKey)
        if err!=nil {
            return nil, errors.New("Unable to get device from stub !")
        } else {
            //fmt.Println("Device data in stub is :", string(stubData))
            return stubData, nil
        }
        
    } else {
        return nil, errors.New("Device id is mandatory !")
    }
    return nil, nil
}

/************************ readUsageHistory ********************/

func (t *DeviceUsageChaincode) readDeviceList(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var devices DevList
    var devSlice []string
    var devList DevList
    devListKey:=LISTKEY
    var devListData []byte
    var iCount int =0
    stubData, err:= stub.GetState(devListKey)
    //fmt.Println("inside device list")
    if err!=nil {
        return nil, errors.New("Unable to get device from stub !")
    } else {
        //fmt.Println("Device data in stub is :", string(stubData))
         err = json.Unmarshal(stubData, &devices)
         fmt.Println()
         sDevicesList:=devices.Devices
         //result := strings.Split(sDevicesList, ",")
         fmt.Println("Devices list is : ", sDevicesList)
          for l := range sDevicesList {
       //Get current status of all devices
            sDeviceId:=sDevicesList[l] 
            fmt.Println(sDeviceId)
            sDeviceKey:=DEVICESKEY+"_"+sDeviceId
            fmt.Println(sDeviceKey)
            stubData, err := stub.GetState(sDeviceKey)
            if err ==nil && len(stubData) >0 {
                if iCount ==0 {
                    devSlice = make([]string, 0)
                    iCount++
                }
                devSlice = append(devSlice, string(stubData))
            }
            iCount++
        }
        devList.Devices = devSlice
        devListData, err = json.Marshal(&devList)
        return devListData, nil
    }
    return nil, nil
}

/************************ extendUsage ********************/


func (t *DeviceUsageChaincode) extendUsage(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    var usage Usage
    var usageStub Usage
    var device Device
	var err error
    
    if len(args) != 1 {
        return nil, errors.New("Expects one argument, a JSON string with device usage details")
    }
     //fmt.Println("Inside createUsage", args[0])
    err = json.Unmarshal([]byte(args[0]), &usage)
    if err != nil {
        return nil, errors.New("Unable to unmarshal device usage data " + fmt.Sprint(err))
    } 
    if usage.Duration==0 {
        return nil, errors.New("Cannot block device for 0 minutes " + fmt.Sprint(err))
    }
    /////////////////////////////////////////////////
    // Create device key for usage stub interactions
    sDeviceId:=usage.DeviceID 
    sDeviceKey:=DEVICESKEY+"_"+sDeviceId
    sUsageKey:=USAGEKEY+"_"+sDeviceId
    stubData, err:= stub.GetState(sDeviceKey)
    if err!=nil {
        return nil, errors.New("Device does not exist in stub!")
    }
    //fmt.Println("Setting avail flag to false")
    /////////////////////////////////////////////////
    //Set the Available flag for the device to false
    err = json.Unmarshal(stubData, &device)
    if err != nil {
        //fmt.Println(err)
        return nil, err
    }
    *device.Available = false
     sStubUpdate, err := json.Marshal(&device)
    if err !=nil {
        return nil, err
    }
    //fmt.Println("PUSHING DEVICE RECORD")
    /////////////////////////////////////////////////
    // Push the updated device data with Available flag set to false, to the stub
    err = stub.PutState(sDeviceKey,sStubUpdate)
    if err != nil {
        return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
    }
    /////////////////////////////////////////////////
    // Get usage data from the stub
    sUsageData, err:= stub.GetState(sUsageKey)
    if err!=nil {
        return nil, errors.New("Device does not have usage data!")
    }
    err = json.Unmarshal(sUsageData, &usageStub)
    if err != nil {
        //fmt.Println(err)
        return nil, err
    }
    /////////////////////////////////////////////////
    // Assuming start time and duration are provided by the invoke call, calculate end time
    fExtension :=*device.OvertimeUsageCost*float64(usage.Duration)
    //fmt.Println("calculating time")
    sOldEndTime := usageStub.EndTime[0:19]
    //fmt.Println("The old end time is ", sOldEndTime)

    tEndTime, err := time.Parse("2006-01-02 15:04:05", sOldEndTime)
    if err != nil {
        //fmt.Println(err)
        return nil, err
    }
    iDuration := time.Duration(usage.Duration)*time.Minute
    dEndTime:= tEndTime.Add(iDuration)
    usageStub.EndTime = dEndTime.String()
    //Computing usage cost
    usageStub.OvertimeCost = fExtension
    usageStub.TotalCost = usageStub.UsageCost+fExtension
    // Put usage record to state
    sStubUpdate, err = json.Marshal(&usageStub)
    if err !=nil {
        return nil, err
    }
    /////////////////////////////////////////////////
    // Put the device data back into the stub
    err = stub.PutState(sUsageKey,sStubUpdate)
    if err != nil {
        return nil, errors.New("Usage record failed PUT to ledger: " + fmt.Sprint(err))
    }
   /*
    *device.Available = true
    sStubUpdate, err = json.Marshal(&device)
    if err !=nil {
        return nil, err
    }
    err = stub.PutState(sDeviceKey,sStubUpdate)
    if err != nil {
        return nil, errors.New("Device record failed PUT to ledger: " + fmt.Sprint(err))
    }  */
    //fmt.Println("Create Usage Over")
    return nil, nil
}
//*************readAssetSchemas*******************/

func (t *DeviceUsageChaincode) readAssetSchemas(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	return []byte(schemas), nil
}
