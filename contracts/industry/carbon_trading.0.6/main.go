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
Howard McKinney- Initial Contribution
Risham Chokshi - Initial Contribution
*/

// IoT Blockchain Demo Smart Contract for Carbon Emission
// v4   KL 23 June 2016 Initial contract supporting multiple assets with carbon reading, threshold and carbon
// v5   RC 5 July 2016 updated credit sold attributes
// v5.1 RC 6 July 2016 added some basic trading attributes
// v5.2 RC 12 July 2016 added attributes to request buying credits
// v6   RC 14 July 2016 added trade history attributes
// v6.1 RC 22 July 2016 added "trade" asset to store all trade history
// v7 RC 26 July 2016 added weather data, notification alerts and contact information attributes
package main

import (
    "encoding/json"
    "fmt"
    "strings"

    "strconv"
    "errors"
    "time"
    
    "reflect"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    
)

//go:generate go run scripts/generate_go_schema.go


//***************************************************
//***************************************************
//* CONTRACT initialization and runtime engine
//***************************************************
//***************************************************

// ************************************
// definitions 
// ************************************

// SimpleChaincode is the receiver for all shim API
type SimpleChaincode struct {
}

// ASSETID is the JSON tag for the assetID
const ASSETID string = "assetID"
// TIMESTAMP is the JSON tag for timestamps, devices must use this tag to be compatible! 
const TIMESTAMP string = "timestamp"
// ArgsMap is a generic map[string]interface{} to be used as a receiver 
type ArgsMap map[string]interface{} 

var log = NewContractLogger(DEFAULTNICKNAME, DEFAULTLOGGINGLEVEL)

// ************************************
// start the message pumps 
// ************************************
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        log.Infof("ERROR starting Simple Chaincode: %s", err)
    }
}

// Init is called in deploy mode when contract is initialized
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    var stateArg ContractState
    var err error

    log.Info("Entering INIT")
    
    if len(args) != 1 {
        err = errors.New("init expects one argument, a JSON string with  mandatory version and optional nickname") 
        log.Critical(err)
        return nil, err 
    }

    err = json.Unmarshal([]byte(args[0]), &stateArg)
    if err != nil {
        err = fmt.Errorf("Version argument unmarshal failed: %s", err)
        log.Critical(err)
        return nil, err 
    }
    
    if stateArg.Nickname == "" {
        stateArg.Nickname = DEFAULTNICKNAME
    } 

    (*log).setModule(stateArg.Nickname)
    
    err = initializeContractState(stub, stateArg.Version, stateArg.Nickname)
    if err != nil {
        return nil, err
    }
    
    log.Info("Contract initialized")
    return nil, nil
}

// Invoke is called in invoke mode to delegate state changing function messages 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "createAsset" {
        return t.createAsset(stub, args)
    } else if function == "updateAsset" {
        return t.updateAsset(stub, args)
    } else if function == "deleteAsset" {
        return t.deleteAsset(stub, args)
    } else if function == "deleteAllAssets" {
        return t.deleteAllAssets(stub, args)
    } else if function == "deletePropertiesFromAsset" {
        return t.deletePropertiesFromAsset(stub, args)
    } else if function == "setLoggingLevel" {
        return nil, t.setLoggingLevel(stub, args)
    } else if function == "setCreateOnUpdate" {
        return nil, t.setCreateOnUpdate(stub, args)
    }
    err := fmt.Errorf("Invoke received unknown invocation: %s", function)
    log.Warning(err)
    return nil, err
}

// Query is called in query mode to delegate non-state-changing queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if function == "readAsset" {
        return t.readAsset(stub, args)
    } else if function == "readAllAssets" {
        return t.readAllAssets(stub, args)
    } else if function == "readRecentStates" {
        return readRecentStates(stub)
    } else if function == "readAssetHistory" {
        return t.readAssetHistory(stub, args)
    } else if function == "readAssetSamples" {
        return t.readAssetSamples(stub, args)
    } else if function == "readAssetSchemas" {
        return t.readAssetSchemas(stub, args)
    } else if function == "readContractObjectModel" {
        return t.readContractObjectModel(stub, args)
    } else if function == "readContractState" {
        return t.readContractState(stub, args)
    }
    err := fmt.Errorf("Query received unknown invocation: %s", function)
    log.Warning(err)
    return nil, err
}


//***************************************************
//***************************************************
//* ASSET CRUD INTERFACE
//***************************************************
//***************************************************

// ************************************
// createAsset 
// ************************************
func (t *SimpleChaincode) createAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var argsMap ArgsMap
    var event interface{}
    var found bool
    var err error
    var timeIn time.Time

    log.Info("Entering createAsset")

    // allowing 2 args because updateAsset is allowed to redirect when
    // asset does not exist
    if len(args) < 1 || len(args) > 2 {
        err = errors.New("Expecting one JSON event object")
        log.Error(err)
        return nil, err
    }
    
    assetID = ""
    eventBytes := []byte(args[0])
    log.Debugf("createAsset arg: %s", args[0])

    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        log.Errorf("createAsset failed to unmarshal arg: %s", err)
        return nil, err
    } 
    
    if event == nil {
        err = errors.New("createAsset unmarshal arg created nil event")
        log.Error(err)
        return nil, err
    }

    argsMap, found = event.(map[string]interface{})
    if !found {
        err := errors.New("createAsset arg is not a map shape")
        log.Error(err)
        return nil, err
    }

    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("createAsset arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }
    //checking if allottedCredits is present or not? 
    //check if asset created is not trade, if it is, then all the field do not need to be present
    if assetID != "trade" {
        if argsMap["allottedCredits"] == nil{
            err := errors.New("createAsset arg does not include allottedCredits credits specified")
            log.Error(err)
            return nil, err
        }
        //check if allottedCredits value is negative, throw error if it is
        if argsMap["allottedCredits"].(float64) < 0.0 {
            err := errors.New("createAsset arg allottedCredits credits cannot be negative")
            log.Error(err)
            return nil, err
        }
        //check for email field specified
        if argsMap["email"] == nil{
            err := errors.New("createAsset arg does not have Email specified")
            log.Error(err)
            return nil, err
        }
        //create a contactInformation attribute
        contactInfo := make(map[string]interface{})
        contactInfo["email"] = argsMap["email"].(string)
        if argsMap["phoneNum"] != nil{
            contactInfo["phoneNum"] = argsMap["phoneNum"].(string)
        }
        argsMap["contactInformation"] = contactInfo
        //initializing value of reading to equal 0
        if argsMap["reading"] == nil{
            argsMap["reading"] = 0.0
        }
        //initialize credits sold for the trade
        if argsMap["soldCredits"] == nil{
            argsMap["soldCredits"] = 0.0
        }
        //initializing boughtCredits attribute to 0
        if argsMap["boughtCredits"] == nil{
            argsMap["boughtCredits"] = 0.0
        }
        //check if value of pricePerCredit given to update is not negative
        if argsMap["pricePerCredit"]!=nil && argsMap["pricePerCredit"].(float64) < 0.0 {
            err := errors.New("updateAsset arg pricePerCredit needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //check if value of creditsForSale given to update is not negative
        if argsMap["creditsForSale"]!=nil && argsMap["creditsForSale"].(float64) < 0.0 {
            err := errors.New("updateAsset arg creditsForSale needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //check if value of priceRequestBuy given to update is not negative
        if argsMap["priceRequestBuy"]!=nil && argsMap["priceRequestBuy"].(float64) < 0.0 {
            err := errors.New("updateAsset arg priceRequestBuy needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //check if value of creditsRequestBuy given to update is not negative
        if argsMap["creditsRequestBuy"]!=nil && argsMap["creditsRequestBuy"].(float64) < 0.0 {
            err := errors.New("updateAsset arg creditsRequestBuy needs be a positive value")
            log.Error(err)
            return nil, err
        }
    }
    found = assetIsActive(stub, assetID)
    if found {
        err := fmt.Errorf("createAsset arg asset %s already exists", assetID)
        log.Error(err)
        return nil, err
    }

    // test and set timestamp
    // TODO get time from the shim as soon as they support it, we cannot
    // get consensus now because the timestamp is different on all peers.
    var timeOut = time.Now() // temp initialization of time variable - not really needed.. keeping old line
    timeInBytes, found := getObject(argsMap, TIMESTAMP)
    
    if found {
        timeIn, found = timeInBytes.(time.Time)
        if found && !timeIn.IsZero() {
            timeOut = timeIn
        }
    }
    txnunixtime, err := stub.GetTxTimestamp()
    if err != nil {
        err = fmt.Errorf("Error getting transaction timestamp: %s", err)
        log.Error(err)
        return nil, err
    }
    txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
    timeOut = txntimestamp
    //*************************************************//
    argsMap[TIMESTAMP] = timeOut
    // run the rules and raise or clear alerts, if assetID is not 'trade'
    if assetID == "trade" {
        alerts := newAlertStatus()
        if argsMap.executeRules(&alerts) {
            // NOT compliant!
            log.Noticef("createAsset assetID %s is noncompliant", assetID)
            argsMap["alerts"] = alerts
            delete(argsMap, "incompliance")
        } else {
            if alerts.AllClear() {
                // all false, no need to appear
                delete(argsMap, "alerts")
            } else {
                argsMap["alerts"] = alerts
            }
            argsMap["incompliance"] = true
        }
    }
    // copy incoming event to outgoing state
    // this contract respects the fact that createAsset can accept a partial state
    // as the moral equivalent of one or more discrete events
    // further: this contract understands that its schema has two discrete objects
    // that are meant to be used to send events: common, and custom
    stateOut := argsMap
    
    // save the original event
    stateOut["lastEvent"] = make(map[string]interface{})
    stateOut["lastEvent"].(map[string]interface{})["function"] = "createAsset"
    stateOut["lastEvent"].(map[string]interface{})["args"] = args[0]
    if len(args) == 2 {
        // in-band protocol for redirect
        stateOut["lastEvent"].(map[string]interface{})["redirectedFromFunction"] = args[1]
    }

    // marshal to JSON and write
    stateJSON, err := json.Marshal(&stateOut)
    if err != nil {
        err := fmt.Errorf("createAsset state for assetID %s failed to marshal", assetID)
        log.Error(err)
        return nil, err
    }

    // finally, put the new state
    log.Infof("Putting new asset state %s to ledger", string(stateJSON))
    err = stub.PutState(assetID, []byte(stateJSON))
    if err != nil {
        err = fmt.Errorf("createAsset AssetID %s PUTSTATE failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    log.Infof("createAsset AssetID %s state successfully written to ledger: %s", assetID, string(stateJSON))

    // add asset to contract state
    err = addAssetToContractState(stub, assetID)
    if err != nil {
        err := fmt.Errorf("createAsset asset %s failed to write asset state: %s", assetID, err)
        log.Critical(err)
        return nil, err 
    }

    err = pushRecentState(stub, string(stateJSON))
    if err != nil {
        err = fmt.Errorf("createAsset AssetID %s push to recentstates failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }

    // save state history
    err = createStateHistory(stub, assetID, string(stateJSON))
    if err != nil {
        err := fmt.Errorf("createAsset asset %s state history save failed: %s", assetID, err)
        log.Critical(err)
        return nil, err 
    }
    
    return nil, nil
}

// ************************************
// updateAsset 
// ************************************
func (t *SimpleChaincode) updateAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var argsMap ArgsMap
    var event interface{}
    var ledgerMap ArgsMap
    var ledgerBytes interface{}
    var found bool
    var err error
    var timeIn time.Time
    
    log.Info("Entering updateAsset")

    if len(args) != 1 {
        err = errors.New("Expecting one JSON event object")
        log.Error(err)
        return nil, err
    }
    
    assetID = ""
    eventBytes := []byte(args[0])
    log.Debugf("updateAsset arg: %s", args[0])
    
    
    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        log.Errorf("updateAsset failed to unmarshal arg: %s", err)
        return nil, err
    }

    if event == nil {
        err = errors.New("createAsset unmarshal arg created nil event")
        log.Error(err)
        return nil, err
    }

    argsMap, found = event.(map[string]interface{})
    if !found {
        err := errors.New("updateAsset arg is not a map shape")
        log.Error(err)
        return nil, err
    }
    
    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("updateAsset arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }
    log.Noticef("updateAsset found assetID %s", assetID)

    found = assetIsActive(stub, assetID)
    if !found {
        // redirect to createAsset with same parameter list
        if canCreateOnUpdate(stub) {
            log.Noticef("updateAsset redirecting asset %s to createAsset", assetID)
            var newArgs = []string{args[0], "updateAsset"}
            return t.createAsset(stub, newArgs)
        }
        err = fmt.Errorf("updateAsset asset %s does not exist", assetID)
        log.Error(err)
        return nil, err
    }

    // test and set timestamp
    // TODO get time from the shim as soon as they support it, we cannot
    // get consensus now because the timestamp is different on all peers.
    //*************************************************//
    // Suma quick fix for timestamp  - Aug 1
    var timeOut = time.Now() // temp initialization of time variable - not really needed.. keeping old line
    timeInBytes, found := getObject(argsMap, TIMESTAMP)
    
    if found {
        timeIn, found = timeInBytes.(time.Time)
        if found && !timeIn.IsZero() {
            timeOut = timeIn
        }
    }
    txnunixtime, err := stub.GetTxTimestamp()
    if err != nil {
        err = fmt.Errorf("Error getting transaction timestamp: %s", err)
        log.Error(err)
        return nil, err
    }
    txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
    timeOut = txntimestamp
    //*************************************************//
    argsMap[TIMESTAMP] = timeOut
    //else, it will be timestamped with data taken in from sensor
    // **********************************
    // find the asset state in the ledger
    // **********************************
    log.Infof("updateAsset: retrieving asset %s state from ledger", assetID)
    assetBytes, err := stub.GetState(assetID)
    if err != nil {
        log.Errorf("updateAsset assetID %s GETSTATE failed: %s", assetID, err)
        return nil, err
    }
    // unmarshal the existing state from the ledger to theinterface
    err = json.Unmarshal(assetBytes, &ledgerBytes)
    if err != nil {
        log.Errorf("updateAsset assetID %s unmarshal failed: %s", assetID, err)
        return nil, err
    } 
    // assert the existing state as a map
    ledgerMap, found = ledgerBytes.(map[string]interface{})
    if !found {
        log.Errorf("updateAsset assetID %s LEDGER state is not a map shape", assetID)
        return nil, err
    }
    //For upadating reading value for an asset
    //Adding what was previous value of ledgerMap reading with what we are getting 
    if assetID != "trade" {
        /****************************** DELETE OR NOT? TO DOOOO *****************************/
        if argsMap["reading"] != nil && argsMap["temperatureCelsius"] != nil && argsMap["temperatureFahrenheit"] != nil && argsMap["windSpeed"] != nil && argsMap["windGustSpeed"] != nil && argsMap["precipitation"] != nil && argsMap["windDegrees"] != nil && argsMap["iconUrl"] != nil{
            //float attributes will be equal to argsMap value and that value will be appended to sensorWeatherHistory attribute
            var readingVal string = checkandConvertType(argsMap["reading"])
            var tempCel string = checkandConvertType(argsMap["temperatureCelsius"])
            var tempFah string = checkandConvertType(argsMap["temperatureFahrenheit"])
            var windSp string = checkandConvertType(argsMap["windSpeed"])
            var windGustSp string = checkandConvertType(argsMap["windGustSpeed"])
            var windDeg string = checkandConvertType(argsMap["windDegrees"])
            var prep string = checkandConvertType(argsMap["precipitation"])
            //setting value to be reading to 0 if value of current used carbon emission exerted needs to be 0
            if argsMap["reading"].(float64) < 0.0 {
                readingVal = "0.0" 
                argsMap["reading"] = 0.0
            } else {
                argsMap["reading"] = argsMap["reading"].(float64) + ledgerMap["reading"].(float64)
            }
            ledgerMap["sensorWeatherHistory"] = ledgerMap.updateSensorWeatherBlock(readingVal, timeOut.String(), tempCel, tempFah, windSp, windGustSp, prep, windDeg, argsMap["iconUrl"].(string))
        } else if argsMap["reading"] != nil || argsMap["temperatureCelsius"] != nil || argsMap["temperatureFahrenheit"] != nil || argsMap["windSpeed"] != nil || argsMap["windGustSpeed"] != nil || argsMap["precipitation"] != nil || argsMap["windDegrees"] != nil || argsMap["iconUrl"] != nil{
            //all the fields should be present, throw an error
            err := errors.New("updateAsset some of the attributes are missing for sensorWeather history: reading, temperatureFahrenheit, temperatureCelsius, windSpeed, windGustSpeed, precipitation, wind Degrees, icon URL - should be all present")
            log.Error(err)
            return nil, err
        }
        //check if allottedCredits value is negative, throw error if it is
        if argsMap["allottedCredits"] != nil{
            if argsMap["allottedCredits"].(float64) < 0.0 {
                err := errors.New("updateAsset arg allottedCredits credits cannot be negative")
                log.Error(err)
                return nil, err
            }
        }
        //update a contactInformation attribute
        if argsMap["email"] != nil || argsMap["phoneNum"] != nil{
            contactInfo := make(map[string]interface{})
            //update email if argument is given
            if argsMap["email"] != nil{
                contactInfo["email"] = argsMap["email"].(string)
            } else {
                contactInfo["email"] = ledgerMap["contactInformation"].(map[string]interface{})["email"]
            }
            //update phone number if argument is given
            if argsMap["phoneNum"] != nil{
                contactInfo["phoneNum"] = argsMap["phoneNum"].(string)
            } else {
                contactInfo["phoneNum"] = ledgerMap["contactInformation"].(map[string]interface{})["phoneNum"]
            }
            argsMap["contactInformation"] = contactInfo
        }
        //argument to change number of credits which was sold
        if argsMap["soldCredits"] != nil{
            argsMap["soldCredits"] = argsMap["soldCredits"].(float64) + ledgerMap["soldCredits"].(float64)
        }

        //argument to change number of credits which was bought
        if argsMap["boughtCredits"] != nil{
            argsMap["boughtCredits"] = argsMap["boughtCredits"].(float64) + ledgerMap["boughtCredits"].(float64)
        }
        //check if pricePerCredit or creditsForSale values was updated
        //checking error cases with pricePerCredit and creditsForSale
        //throw an error if pricePerCredit was updated without putting in value for creditsForSale
        if argsMap["pricePerCredit"] != nil && argsMap["creditsForSale"] == nil && ledgerMap["creditsForSale"] == nil {
            err := errors.New("updateAsset arg creditsForSale need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err

        }
        //throw an error if creditsForSale was updated without putting in value for pricePerCredit
        if argsMap["pricePerCredit"] == nil && ledgerMap["pricePerCredit"] == nil && argsMap["creditsForSale"] != nil {
            err := errors.New("updateAsset arg pricePerCredit need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        //check if value of pricePerCredit given to update is not negative
        if argsMap["pricePerCredit"]!=nil && argsMap["pricePerCredit"].(float64) < 0.0 {
            err := errors.New("updateAsset arg pricePerCredit needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //check if value of creditsForSale given to update is not negative
        if argsMap["creditsForSale"]!=nil && argsMap["creditsForSale"].(float64) < 0.0 {
            err := errors.New("updateAsset arg creditsForSale needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //creditsForSale needs to be less than or equal to remaining credits  
        //credits put for sale should be less than remaining credits
        var totalCred float64 = 0
        if argsMap["creditsForSale"] != nil && ledgerMap["creditsSellList"] != nil {
            //calculate all the credits put on sale
            totalCred = addValuesInArray(ledgerMap["creditsSellList"].([]interface{}))
        }
        if argsMap["creditsForSale"] != nil && (argsMap["creditsForSale"].(float64) > (ledgerMap["allottedCredits"].(float64) - ledgerMap["reading"].(float64) - ledgerMap["soldCredits"].(float64) + ledgerMap["boughtCredits"].(float64) - totalCred)){
            err := errors.New("updateAsset arg creditsForSale needs be less and or equal to remaining credits" + strconv.FormatFloat(ledgerMap["allottedCredits"].(float64) - ledgerMap["reading"].(float64), 'f', -1, 64))
            log.Error(err)
            return nil, err
        }
        //it passed all the error checks and credits and price for sell should be added to the list
        if argsMap["creditsForSale"] != nil {
            //first request on sell, converting float to string 
            var creditSellVal string = strconv.FormatFloat(argsMap["creditsForSale"].(float64), 'f', -1, 64)
            var priceSellVal string = strconv.FormatFloat(argsMap["pricePerCredit"].(float64), 'f', -1, 64)
            //if the list is null, add to the list
            if ledgerMap["creditsSellList"] == nil {
                ledgerMap["creditsSellList"] = []string{creditSellVal}
                ledgerMap["priceSellList"] = []string{priceSellVal}
            } else {
                //if the list isn't null, then append to the list
                ledgerMap["creditsSellList"] = append(ledgerMap["creditsSellList"].([]interface{}), creditSellVal)
                ledgerMap["priceSellList"] = append(ledgerMap["priceSellList"].([]interface{}), priceSellVal)
            }
        }    
        //credits requested to buy and price it was requested
        if argsMap["priceRequestBuy"] != nil && argsMap["creditsRequestBuy"] == nil && ledgerMap["creditsRequestBuy"] == nil {
            err := errors.New("updateAsset arg creditsRequestBuy need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        //throw an error if creditsForSale was updated without putting in value for pricePerCredit
        if argsMap["priceRequestBuy"] == nil && ledgerMap["priceRequestBuy"] == nil && argsMap["creditsRequestBuy"] != nil {
            err := errors.New("updateAsset arg priceRequestBuy need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        //check if value of pricePerCredit given to update is not negative
        if argsMap["priceRequestBuy"]!=nil && argsMap["priceRequestBuy"].(float64) < 0.0 {
            err := errors.New("updateAsset arg priceRequestBuy needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //check if value of creditsForSale given to update is not negative
        if argsMap["creditsRequestBuy"]!=nil && argsMap["creditsRequestBuy"].(float64) < 0.0 {
            err := errors.New("updateAsset arg creditsRequestBuy needs be a positive value")
            log.Error(err)
            return nil, err
        }
        //add it to the list for requesting different number of credits
        //it passed all the error checks and credits and price for sell should be added to the list
        if argsMap["creditsRequestBuy"] != nil {
            //first request on sell, converting float to string 
            var creditSellVal string = strconv.FormatFloat(argsMap["creditsRequestBuy"].(float64), 'f', -1, 64)
            var priceSellVal string = strconv.FormatFloat(argsMap["priceRequestBuy"].(float64), 'f', -1, 64)
            //if the list is null, add to the list
            if ledgerMap["creditsBuyList"] == nil {
                ledgerMap["creditsBuyList"] = []string{creditSellVal}
                ledgerMap["priceBuyList"] = []string{priceSellVal}
            } else {
                //if the list isn't null, then append to the list
                ledgerMap["creditsBuyList"] = append(ledgerMap["creditsBuyList"].([]interface{}), creditSellVal)
                ledgerMap["priceBuyList"] = append(ledgerMap["priceBuyList"].([]interface{}), priceSellVal)
            }
        }
        //check if any of the arguments are about updating sell list
        if argsMap["updateSellIndex"] != nil && argsMap["updateSellCredits"] == nil {
            err := errors.New("updateAsset arg updateSellCredits need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        if argsMap["updateSellIndex"] == nil && argsMap["updateSellCredits"] != nil {
            err := errors.New("updateAsset arg updateSellIndex need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        //update sell list credit attribute depending on index and credit values
        if argsMap["updateSellIndex"] != nil && argsMap["updateSellCredits"] != nil {
            var index int = int(argsMap["updateSellIndex"].(float64))
            var credits float64 = argsMap["updateSellCredits"].(float64)
            if index < len(ledgerMap["creditsSellList"].([]interface{})) && index >= 0 {
                if credits > 0 {
                    ledgerMap["creditsSellList"].([]interface{})[index] = strconv.FormatFloat(credits, 'f', -1, 64)
                } else {
                    ledgerMap["creditsSellList"] = append(ledgerMap["creditsSellList"].([]interface{})[:index], ledgerMap["creditsSellList"].([]interface{})[index+1:]...)
                    ledgerMap["priceSellList"] = append(ledgerMap["priceSellList"].([]interface{})[:index], ledgerMap["priceSellList"].([]interface{})[index+1:]...)
                }
            } else{
                //the index provided is out of bound, throwing an error
                err := errors.New("updateAsset arg updateSellIndex cannot be out of bound for SellList")
                log.Error(err)
                return nil, err
            }
        }
        //check if any of update arguments are about buy list
        if argsMap["updateBuyIndex"] != nil && argsMap["updateBuyCredits"] == nil {
            err := errors.New("updateAsset arg updateBuyCredits need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        if argsMap["updateBuyIndex"] == nil && argsMap["updateBuyCredits"] != nil {
            err := errors.New("updateAsset arg updateBuyIndex need to be present, that cannot be empty. Current value is nil")
            log.Error(err)
            return nil, err
        }
        //update sell list credit attribute depending on index and credit values
        if argsMap["updateBuyIndex"] != nil && argsMap["updateBuyCredits"] != nil {
            var index int = int(argsMap["updateBuyIndex"].(float64))
            var credits float64 = argsMap["updateBuyCredits"].(float64)
            if index < len(ledgerMap["creditsBuyList"].([]interface{}))  && index >=0 {
                if credits > 0 {
                    ledgerMap["creditsBuyList"].([]interface{})[index] = strconv.FormatFloat(credits, 'f', -1, 64)
                } else {
                    ledgerMap["creditsBuyList"] = append(ledgerMap["creditsBuyList"].([]interface{})[:index], ledgerMap["creditsBuyList"].([]interface{})[index+1:]...)
                    ledgerMap["priceBuyList"] = append(ledgerMap["priceBuyList"].([]interface{})[:index], ledgerMap["priceBuyList"].([]interface{})[index+1:]...)
                }
            } else{
                //the index provided is out of bound, throwing an error
                err := errors.New("updateAsset arg updateBuyIndex cannot be out of bound for BuyList")
                log.Error(err)
                return nil, err
            }
        }
    }
    //check if the udpate is for trade 
    //********************TRADE HISTORY ASPECT********************
    //checking for all the attributes
     //checking for all the attributes
    if assetID == "trade" && argsMap["tradeCredits"] != nil && argsMap["tradePrice"] != nil && argsMap["tradeTimestamp"] != nil {
        ledgerMap["tradeHistory"] = ledgerMap.updateTradeBlock(false, argsMap["tradeCredits"].(string), argsMap["tradePrice"].(string), argsMap["tradeTimestamp"].(string), "", "")
    } else if assetID == "trade" && (argsMap["tradeCredits"] != nil || argsMap["tradePrice"] != nil || argsMap["tradeTimestamp"] != nil) {
        //all the fields should be present, throw an error
        err := errors.New("updateAsset some of the attributes are missing for 'trade' asset in trade history: tradeCredits, tradePrice, tradeTimestamp - should be all present")
        log.Error(err)
        return nil, err
    } else if argsMap["tradeCredits"] != nil && argsMap["tradePrice"] != nil && argsMap["tradeTimestamp"] != nil && argsMap["tradeCompany"] != nil && argsMap["tradeBuySell"] != nil {
        ledgerMap["tradeHistory"] = ledgerMap.updateTradeBlock(true, argsMap["tradeCredits"].(string), argsMap["tradePrice"].(string), argsMap["tradeTimestamp"].(string), argsMap["tradeCompany"].(string), argsMap["tradeBuySell"].(string))
    } else if argsMap["tradeCredits"] != nil || argsMap["tradePrice"] != nil || argsMap["tradeTimestamp"] != nil || argsMap["tradeCompany"] != nil || argsMap["tradeBuySell"] != nil {
        //all the fields should be present, throw an error
        err := errors.New("updateAsset some of the attributes are missing for trade history: tradeCredits, tradeCompany, tradePrice, tradeTimestamp, tradeBuySell - should be all present")
        log.Error(err)
        return nil, err
    }
    // now add incoming map values to existing state to merge them
    // this contract respects the fact that updateAsset can accept a partial state
    // as the moral equivalent of one or more discrete events
    // further: this contract understands that its schema has two discrete objects
    // that are meant to be used to send events: common, and custom
    // ledger has to have common section
    stateOut := deepMerge(map[string]interface{}(argsMap), 
                          map[string]interface{}(ledgerMap))
    log.Debugf("updateAsset assetID %s merged state: %s", assetID, stateOut)

    // handle compliance section
    if assetID != "trade" {
        alerts := newAlertStatus()
        a, found := stateOut["alerts"] // is there an existing alert state?
        if found {
            // convert to an AlertStatus, which does not work by type assertion
            log.Debugf("updateAsset Found existing alerts state: %s", a)
            // complex types are all untyped interfaces, so require conversion to
            // the structure that is used, but not in the other direction as the
            // type is properly specified
            alerts.alertStatusFromMap(a.(map[string]interface{}))
        }
        // important: rules need access to the entire calculated state 
        if ledgerMap.executeRules(&alerts) {
            // true means noncompliant
            log.Noticef("updateAsset assetID %s is noncompliant", assetID)
            // update ledger with new state, if all clear then delete
            stateOut["alerts"] = alerts
            delete(stateOut, "incompliance")
        } else {
            if alerts.AllClear() {
                // all false, no need to appear
                delete(stateOut, "alerts")
            } else {
                stateOut["alerts"] = alerts
            }
            stateOut["incompliance"] = true
        }
    }
    
    // save the original event
    stateOut["lastEvent"] = make(map[string]interface{})
    stateOut["lastEvent"].(map[string]interface{})["function"] = "updateAsset"
    stateOut["lastEvent"].(map[string]interface{})["args"] = args[0]

    // Write the new state to the ledger
    stateJSON, err := json.Marshal(ledgerMap)
    if err != nil {
        err = fmt.Errorf("updateAsset AssetID %s marshal failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    // finally, put the new state
    err = stub.PutState(assetID, []byte(stateJSON))
    if err != nil {
        err = fmt.Errorf("updateAsset AssetID %s PUTSTATE failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    err = pushRecentState(stub, string(stateJSON))
    if err != nil {
        err = fmt.Errorf("updateAsset AssetID %s push to recentstates failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    // add history state
    err = updateStateHistory(stub, assetID, string(stateJSON))
    if err != nil {
        err = fmt.Errorf("updateAsset AssetID %s push to history failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    // NOTE: Contract state is not updated by updateAsset 
    return nil, nil
}

// ************************************
// deleteAsset 
// ************************************
func (t *SimpleChaincode) deleteAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var argsMap ArgsMap
    var event interface{}
    var found bool
    var err error

    if len(args) != 1 {
        err = errors.New("Expecting one JSON state object with an assetID")
        log.Error(err)
        return nil, err
    }
    
    assetID = ""
    eventBytes := []byte(args[0])
    log.Debugf("deleteAsset arg: %s", args[0])

    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        log.Errorf("deleteAsset failed to unmarshal arg: %s", err)
        return nil, err
    }

    argsMap, found = event.(map[string]interface{})
    if !found {
        err := errors.New("deleteAsset arg is not a map shape")
        log.Error(err)
        return nil, err
    }
    
    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("deleteAsset arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }

    found = assetIsActive(stub, assetID)
    if !found {
        err = fmt.Errorf("deleteAsset assetID %s does not exist", assetID)
        log.Error(err)
        return nil, err
    }

    // Delete the key / asset from the ledger
    err = stub.DelState(assetID)
    if err != nil {
        log.Errorf("deleteAsset assetID %s failed DELSTATE", assetID)
        return nil, err
    }
    // remove asset from contract state
    err = removeAssetFromContractState(stub, assetID)
    if err != nil {
        err := fmt.Errorf("deleteAsset asset %s failed to remove asset from contract state: %s", assetID, err)
        log.Critical(err)
        return nil, err 
    }
    // save state history
    err = deleteStateHistory(stub, assetID)
    if err != nil {
        err := fmt.Errorf("deleteAsset asset %s state history delete failed: %s", assetID, err)
        log.Critical(err)
        return nil, err 
    }
    // push the recent state
    err = removeAssetFromRecentState(stub, assetID)
    if err != nil {
        err := fmt.Errorf("deleteAsset asset %s recent state removal failed: %s", assetID, err)
        log.Critical(err)
        return nil, err 
    }
    
    return nil, nil
}

// ************************************
// deletePropertiesFromAsset 
// ************************************
func (t *SimpleChaincode) deletePropertiesFromAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var argsMap ArgsMap
    var event interface{}
    var ledgerMap ArgsMap
    var ledgerBytes interface{}
    var found bool
    var err error
    var alerts AlertStatus

    if len(args) < 1 {
        err = errors.New("Not enough arguments. Expecting one JSON object with mandatory AssetID and property name array")
        log.Error(err)
        return nil, err
    }
    eventBytes := []byte(args[0])

    err = json.Unmarshal(eventBytes, &event)
    if err != nil {
        log.Error("deletePropertiesFromAsset failed to unmarshal arg")
        return nil, err
    }
    
    argsMap, found = event.(map[string]interface{})
    if !found {
        err := errors.New("updateAsset arg is not a map shape")
        log.Error(err)
        return nil, err
    }
    log.Debugf("deletePropertiesFromAsset arg: %+v", argsMap)
    
    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("deletePropertiesFromAsset arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }

    found = assetIsActive(stub, assetID)
    if !found {
        err = fmt.Errorf("deletePropertiesFromAsset assetID %s does not exist", assetID)
        log.Error(err)
        return nil, err
    }

    // is there a list of property names?
    var qprops []interface{}
    qpropsBytes, found := getObject(argsMap, "qualPropsToDelete")
    if found {
        qprops, found = qpropsBytes.([]interface{})
        log.Debugf("deletePropertiesFromAsset qProps: %+v, Found: %+v, Type: %+v", qprops, found, reflect.TypeOf(qprops))
        if !found || len(qprops) < 1 {
            log.Errorf("deletePropertiesFromAsset asset %s qualPropsToDelete is not an array or is empty", assetID)
            return nil, err
        }
    } else {
        log.Errorf("deletePropertiesFromAsset asset %s has no qualPropsToDelete argument", assetID)
        return nil, err
    }

    // **********************************
    // find the asset state in the ledger
    // **********************************
    log.Infof("deletePropertiesFromAsset: retrieving asset %s state from ledger", assetID)
    assetBytes, err := stub.GetState(assetID)
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s GETSTATE failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }

    // unmarshal the existing state from the ledger to the interface
    err = json.Unmarshal(assetBytes, &ledgerBytes)
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s unmarshal failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    
    // assert the existing state as a map
    ledgerMap, found = ledgerBytes.(map[string]interface{})
    if !found {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s LEDGER state is not a map shape", assetID)
        log.Error(err)
        return nil, err
    }

    // now remove properties from state, they are qualified by level
    OUTERDELETELOOP:
    for p := range qprops {
        prop := qprops[p].(string)
        log.Debugf("deletePropertiesFromAsset AssetID %s deleting qualified property: %s", assetID, prop)
        // TODO Ugly, isolate in a function at some point
        if (CASESENSITIVEMODE  && strings.HasSuffix(prop, ASSETID)) ||
           (!CASESENSITIVEMODE && strings.HasSuffix(strings.ToLower(prop), strings.ToLower(ASSETID))) {
            log.Warningf("deletePropertiesFromAsset AssetID %s cannot delete protected qualified property: %s", assetID, prop)
        } else {
            levels := strings.Split(prop, ".")
            lm := (map[string]interface{})(ledgerMap)
            for l := range levels {
                // lev is the name of a level
                lev := levels[l]
                if l == len(levels)-1 {
                    // we're here, delete the actual property name from this level of the map
                    levActual, found := findMatchingKey(lm, lev)
                    if !found {
                        log.Warningf("deletePropertiesFromAsset AssetID %s property match %s not found", assetID, lev)
                        continue OUTERDELETELOOP
                    }
                    log.Debugf("deletePropertiesFromAsset AssetID %s deleting %s", assetID, prop)
                    delete(lm, levActual)
                } else {
                    // navigate to the next level object
                    log.Debugf("deletePropertiesFromAsset AssetID %s navigating to level %s", assetID, lev)
                    lmBytes, found := findObjectByKey(lm, lev)
                    if found {
                        lm, found = lmBytes.(map[string]interface{})
                        if !found {
                            log.Noticef("deletePropertiesFromAsset AssetID %s level %s not found in ledger", assetID, lev)
                            continue OUTERDELETELOOP
                        }
                    } 
                }
            } 
        }
    }
    log.Debugf("updateAsset AssetID %s final state: %s", assetID, ledgerMap)

    // set timestamp
    //*************************************************//
    // Suma quick fix for timestamp  - Aug 1
     txnunixtime, err := stub.GetTxTimestamp()
    if err != nil {
        err = fmt.Errorf("Error getting transaction timestamp: %s", err)
        log.Error(err)
        return nil, err
    }
    txntimestamp := time.Unix(txnunixtime.Seconds, int64(txnunixtime.Nanos))
    ledgerMap[TIMESTAMP] = txntimestamp
    //*************************************************//

    // handle compliance section
    alerts = newAlertStatus()
    a, found := argsMap["alerts"] // is there an existing alert state?
    if found {
        // convert to an AlertStatus, which does not work by type assertion
        log.Debugf("deletePropertiesFromAsset Found existing alerts state: %s", a)
        // complex types are all untyped interfaces, so require conversion to
        // the structure that is used, but not in the other direction as the
        // type is properly specified
        alerts.alertStatusFromMap(a.(map[string]interface{}))
    }
    // important: rules need access to the entire calculated state 
    if ledgerMap.executeRules(&alerts) {
        // true means noncompliant
        log.Noticef("deletePropertiesFromAsset assetID %s is noncompliant", assetID)
        // update ledger with new state, if all clear then delete
        ledgerMap["alerts"] = alerts
        delete(ledgerMap, "incompliance")
    } else {
        if alerts.AllClear() {
            // all false, no need to appear
            delete(ledgerMap, "alerts")
        } else {
            ledgerMap["alerts"] = alerts
        }
        ledgerMap["incompliance"] = true
    }
    
    // save the original event
    ledgerMap["lastEvent"] = make(map[string]interface{})
    ledgerMap["lastEvent"].(map[string]interface{})["function"] = "deletePropertiesFromAsset"
    ledgerMap["lastEvent"].(map[string]interface{})["args"] = args[0]
    
    // Write the new state to the ledger
    stateJSON, err := json.Marshal(ledgerMap)
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s marshal failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }

    // finally, put the new state
    err = stub.PutState(assetID, []byte(stateJSON))
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s PUTSTATE failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    err = pushRecentState(stub, string(stateJSON))
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s push to recentstates failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }

    // add history state
    err = updateStateHistory(stub, assetID, string(stateJSON))
    if err != nil {
        err = fmt.Errorf("deletePropertiesFromAsset AssetID %s push to history failed: %s", assetID, err)
        log.Error(err)
        return nil, err
    }

    return nil, nil
}

// ************************************
// deletaAllAssets 
// ************************************
func (t *SimpleChaincode) deleteAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var err error

    if len(args) > 0 {
        err = errors.New("Too many arguments. Expecting none.")
        log.Error(err)
        return nil, err
    }
    
    aa, err := getActiveAssets(stub)
    if err != nil {
        err = fmt.Errorf("deleteAllAssets failed to get the active assets: %s", err)
        log.Error(err)
        return nil, err
    }
    for i := range aa {
        assetID = aa[i]
        
        // Delete the key / asset from the ledger
        err = stub.DelState(assetID)
        if err != nil {
            err = fmt.Errorf("deleteAllAssets arg %d assetID %s failed DELSTATE", i, assetID)
            log.Error(err)
            return nil, err
        }
        // remove asset from contract state
        err = removeAssetFromContractState(stub, assetID)
        if err != nil {
            err = fmt.Errorf("deleteAllAssets asset %s failed to remove asset from contract state: %s", assetID, err)
            log.Critical(err)
            return nil, err 
        }
        // save state history
        err = deleteStateHistory(stub, assetID)
        if err != nil {
            err := fmt.Errorf("deleteAllAssets asset %s state history delete failed: %s", assetID, err)
            log.Critical(err)
            return nil, err 
        }
    }
    err = clearRecentStates(stub)
    if err != nil {
        err = fmt.Errorf("deleteAllAssets clearRecentStates failed: %s", err)
        log.Error(err)
        return nil, err
    }
    return nil, nil
}

// ************************************
// readAsset 
// ************************************
func (t *SimpleChaincode) readAsset(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var argsMap ArgsMap
    var request interface{}
    var assetBytes []byte
    var found bool
    var err error
    
    if len(args) != 1 {
        err = errors.New("Expecting one JSON event object")
        log.Error(err)
        return nil, err
    }
    
    requestBytes := []byte(args[0])
    log.Debugf("readAsset arg: %s", args[0])
    
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        log.Errorf("readAsset failed to unmarshal arg: %s", err)
        return nil, err
    }

    argsMap, found = request.(map[string]interface{})
    if !found {
        err := errors.New("readAsset arg is not a map shape")
        log.Error(err)
        return nil, err
    }
    
    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("readAsset arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }
    
    found = assetIsActive(stub, assetID)
    if !found {
        err := fmt.Errorf("readAsset arg asset %s does not exist", assetID)
        log.Error(err)
        return nil, err
    }

    // Get the state from the ledger
    assetBytes, err = stub.GetState(assetID)
    if err != nil {
        log.Errorf("readAsset assetID %s failed GETSTATE", assetID)
        return nil, err
    } 

    return assetBytes, nil
}

// ************************************
// readAllAssets 
// ************************************
func (t *SimpleChaincode) readAllAssets(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetID string
    var err error
    var results []interface{}
    var state interface{}

    if len(args) > 0 {
        err = errors.New("readAllAssets expects no arguments")
        log.Error(err)
        return nil, err
    }
    
    aa, err := getActiveAssets(stub)
    if err != nil {
        err = fmt.Errorf("readAllAssets failed to get the active assets: %s", err)
        log.Error(err)
        return nil, err
    }
    results = make([]interface{}, 0, len(aa))
    for i := range aa {
        assetID = aa[i]
        // Get the state from the ledger
        assetBytes, err := stub.GetState(assetID)
        if err != nil {
            // best efforts, return what we can
            log.Errorf("readAllAssets assetID %s failed GETSTATE", assetID)
            continue
        } else {
            err = json.Unmarshal(assetBytes, &state)
            if err != nil {
                // best efforts, return what we can
                log.Errorf("readAllAssets assetID %s failed to unmarshal", assetID)
                continue
            }
            results = append(results, state)
        }
    }
    
    resultsStr, err := json.Marshal(results)
    if err != nil {
        err = fmt.Errorf("readallAssets failed to marshal results: %s", err)
        log.Error(err)
        return nil, err
    }

    return []byte(resultsStr), nil
}

// ************************************
// readAssetHistory 
// ************************************
func (t *SimpleChaincode) readAssetHistory(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var assetBytes []byte
    var assetID string
    var argsMap ArgsMap
    var request interface{}
    var found bool
    var err error

    if len(args) != 1 {
        err = errors.New("readAssetHistory expects a JSON encoded object with assetID and count")
        log.Error(err)
        return nil, err
    }
    
    requestBytes := []byte(args[0])
    log.Debugf("readAssetHistory arg: %s", args[0])
    
    err = json.Unmarshal(requestBytes, &request)
    if err != nil {
        err = fmt.Errorf("readAssetHistory failed to unmarshal arg: %s", err)
        log.Error(err)
        return nil, err
    }
    
    argsMap, found = request.(map[string]interface{})
    if !found {
        err := errors.New("readAssetHistory arg is not a map shape")
        log.Error(err)
        return nil, err
    }
    
    // is assetID present or blank?
    assetIDBytes, found := getObject(argsMap, ASSETID)
    if found {
        assetID, found = assetIDBytes.(string) 
        if !found || assetID == "" {
            err := errors.New("readAssetHistory arg does not include assetID")
            log.Error(err)
            return nil, err
        }
    }
    
    found = assetIsActive(stub, assetID)
    if !found {
        err := fmt.Errorf("readAssetHistory arg asset %s does not exist", assetID)
        log.Error(err)
        return nil, err
    }

    // Get the history from the ledger
    stateHistory, err := readStateHistory(stub, assetID)
    if err != nil {
        err = fmt.Errorf("readAssetHistory assetID %s failed readStateHistory: %s", assetID, err)
        log.Error(err)
        return nil, err
    }
    
    // is count present?
    var olen int
    countBytes, found := getObject(argsMap, "count")
    if found {
        olen = int(countBytes.(float64))
    }
    if olen <= 0 || olen > len(stateHistory.AssetHistory) { 
        olen = len(stateHistory.AssetHistory) 
    }
    var hStatesOut = make([]interface{}, 0, olen) 
    for i := 0; i < olen; i++ {
        var obj interface{}
        err = json.Unmarshal([]byte(stateHistory.AssetHistory[i]), &obj)
        if err != nil {
            log.Errorf("readAssetHistory JSON unmarshal of entry %d failed [%#v]", i, stateHistory.AssetHistory[i])
            return nil, err
        }
        hStatesOut = append(hStatesOut, obj)
    }
    assetBytes, err = json.Marshal(hStatesOut)
    if err != nil {
        log.Errorf("readAssetHistory failed to marshal results: %s", err)
        return nil, err
    }
    
    return []byte(assetBytes), nil
}


//***************************************************
//***************************************************
//* CONTRACT STATE 
//***************************************************
//***************************************************

func (t *SimpleChaincode) readContractState(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var err error

    if len(args) != 0 {
        err = errors.New("Too many arguments. Expecting none.")
        log.Error(err)
        return nil, err
    }

    // Get the state from the ledger
    chaincodeBytes, err := stub.GetState(CONTRACTSTATEKEY)
    if err != nil {
        err = fmt.Errorf("readContractState failed GETSTATE: %s", err)
        log.Error(err)
        return nil, err
    }

    return chaincodeBytes, nil
}

//***************************************************
//***************************************************
//* CONTRACT METADATA / SCHEMA INTERFACE
//***************************************************
//***************************************************

// ************************************
// readAssetSamples 
// ************************************
func (t *SimpleChaincode) readAssetSamples(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    return []byte(samples), nil
}

// ************************************
// readAssetSchemas 
// ************************************
func (t *SimpleChaincode) readAssetSchemas(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    return []byte(schemas), nil
}

// ************************************
// readContractObjectModel 
// ************************************
func (t *SimpleChaincode) readContractObjectModel(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var state = ContractState{MYVERSION, DEFAULTNICKNAME, make(map[string]bool)}

    stateJSON, err := json.Marshal(state)
    if err != nil {
        err := fmt.Errorf("JSON Marshal failed for get contract object model empty state: %+v with error [%s]", state, err)
        log.Error(err)
        return nil, err
    }
    return stateJSON, nil
}

// ************************************
// setLoggingLevel 
// ************************************
func (t *SimpleChaincode) setLoggingLevel(stub shim.ChaincodeStubInterface, args []string) (error) {
    type LogLevelArg struct {
        Level string `json:"logLevel"`
    }
    var level LogLevelArg
    var err error
    if len(args) != 1 {
        err = errors.New("Incorrect number of arguments. Expecting a JSON encoded LogLevel.")
        log.Error(err)
        return err
    }
    err = json.Unmarshal([]byte(args[0]), &level)
    if err != nil {
        err = fmt.Errorf("setLoggingLevel failed to unmarshal arg: %s", err)
        log.Error(err)
        return err
    }
    for i, lev := range logLevelNames {
        if strings.ToUpper(level.Level) == lev {
            (*log).SetLoggingLevel(LogLevel(i))
            return nil
        } 
    }
    err = fmt.Errorf("Unknown Logging level: %s", level.Level)
    log.Error(err)
    return err
}

// CreateOnUpdate is a shared parameter structure for the use of 
// the createonupdate feature
type CreateOnUpdate struct {
    CreateOnUpdate bool `json:"createOnUpdate"`
}

// ************************************
// setCreateOnUpdate 
// ************************************
func (t *SimpleChaincode) setCreateOnUpdate(stub shim.ChaincodeStubInterface, args []string) (error) {
    var createOnUpdate CreateOnUpdate
    var err error
    if len(args) != 1 {
        err = errors.New("setCreateOnUpdate expects a single parameter")
        log.Error(err)
        return err
    }
    err = json.Unmarshal([]byte(args[0]), &createOnUpdate)
    if err != nil {
        err = fmt.Errorf("setCreateOnUpdate failed to unmarshal arg: %s", err)
        log.Error(err)
        return err
    }
    err = PUTcreateOnUpdate(stub, createOnUpdate)
    if err != nil {
        err = fmt.Errorf("setCreateOnUpdate failed to PUT setting: %s", err)
        log.Error(err)
        return err
    }
    return nil
}

// PUTcreateOnUpdate marshals the new setting and writes it to the ledger
func PUTcreateOnUpdate(stub shim.ChaincodeStubInterface, createOnUpdate CreateOnUpdate) (err error) {
    createOnUpdateBytes, err := json.Marshal(createOnUpdate)
    if err != nil {
        err = errors.New("PUTcreateOnUpdate failed to marshal")
        log.Error(err)
        return err
    }
    err = stub.PutState("CreateOnUpdate", createOnUpdateBytes)
    if err != nil {
        err = fmt.Errorf("PUTSTATE createOnUpdate failed: %s", err)
        log.Error(err)
        return err
    }
    return nil
}

// canCreateOnUpdate retrieves the setting from the ledger and returns it to the calling function
func canCreateOnUpdate(stub shim.ChaincodeStubInterface) (bool) {
    var createOnUpdate CreateOnUpdate
    createOnUpdateBytes, err := stub.GetState("CreateOnUpdate")
    if err != nil {
        err = fmt.Errorf("GETSTATE for canCreateOnUpdate failed: %s", err)
        log.Error(err)
        return true  // true is the default
    }
    err = json.Unmarshal(createOnUpdateBytes, &createOnUpdate)
    if err != nil {
        err = fmt.Errorf("canCreateOnUpdate failed to marshal: %s", err)
        log.Error(err)
        return true  // true is the default
    }
    return createOnUpdate.CreateOnUpdate
}

//calculates total number of credits given array of string
func addValuesInArray(args []interface {})(float64){
    if args != nil {
        var totalVal float64 = 0
        for i := range args {
            val, err := strconv.ParseFloat(args[i].(string), 64)
            if err == nil {
                totalVal = totalVal + val
            }
        }
        return totalVal
    }
    return 0.0
}

//check type for the input and conver it to string if needed
func checkandConvertType(args interface{})(string){
    switch v := args.(type) {
    case float64:
        //convert it to a string
        return strconv.FormatFloat(args.(float64), 'f', -1, 64)
    case string:
        return strings.Trim(args.(string), " ")
    default:
        fmt.Printf("Other:%v\n", v)
        return "Type mismatched"
    }
}
