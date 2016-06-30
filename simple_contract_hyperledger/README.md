
# The Basic smart contract sample

The Basic contract is a sample hyperledger blockchain contract that is provided by IBM to help you to get started with blockchain development and integration on the IBM Watson IoT Platform. You can use the Basic contract sample to create a blockchain contract that tracks and stores asset data from a device that is connected to your Watson IoT Platform organization.

The following information is provided to help you to get started with the Basic sample:

- [Overview](#overview)
- [Requirements](#requirements)
- [Developing your IBM Blockchain hyperledger by using the Basic contract sample](#downloading)
- [More contract code functions](#more_functions)
- [Next steps](next_steps)


## Overview
{: #Overview}
The Basic blockchain sample, **simple_contract_hyperledger.go**, is the default smart contract for getting started with writing blockchain contracts. The Basic contract includes create, read, update, and delete asset data operations for tracking device data on the IBM Blockchain ledger. Blockchain contracts including the Sample contract are currently developed in the ‘Go’ programming language but other languages will be supported by IBM in the future.

You can run your smart contracts for Blockchain from either the command line or from a REST interface. For information, see
[Developing smart contracts for Watson IoT Platform blockchain integration](https://console.ng.bluemix.net/docs/services/IoT/blockchain/dev_blockchain.html).

As outlined in Basic sample, IBM Blockchain contract code includes the following methods:

|Method|Description|
|:---|:---|
|`deploy`|?|  
|`invoke`|?|
|`query`|?|

**Note:** In the Basics sample, the `deploy` method is called `init`.

When you call any of the methods, you must first pass a JSON string that includes the function name and a set of arguments as key-value pairs to the chain code instance. You can define multiple methods in within the same contract.

[Where should this go?]
To create a simple contract to create, read, update and delete asset data, you will also need to use the following following methods:

|Method|Provides|
|:---|:---|
|'ReadAssetSchema'|The methods and associated properties of the JSON schema contract.|  
|'ReadAssetSamples'|An example of the sample JSON data|

For more information about setting up REST and Swagger, see [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md).

## Requirements
{: #Requirements}

The Watson IoT Platform includes a data mapping component to route data to the Blockchain contract. In order for the Watson IoT Platform event properties to be correctly mapped to the contract properties, the following contract features are required for data mapping:

1. The ``updateAsset`` function.  
This simple contract is a recipe, an example intended to be tweaked as people experiment with contracts. Today's simple contract, though it implements CRUD, is very accommodating - it has the same features in create and update functions. The Data mapping component expects an `updateAsset` function that can do the following:
 * create the asset record if it doesn't exist on the ledger.
 * update asset data if the asset already exists.
 * asset id the key. The asset id maps to the serial id the comes in from the Watson IoT Platform.
 * Accepts input as a JSON string. This makes sense as IoT device data generally comes in as JSON strings.

2. The `readAssetSchemas` function.
The readAssetSchemas function exposes the function names and properties expected by the contract. This function is called by the mapper in order to know what the contract properties are, thus enabling the mapping of the event properties to the contract properties.

## Developing your IBM Blockchain hyperledger by using the Basic contract sample
{: #instructions}

To use the Basic sample **simple_contract_hyperledger.go** sample contract as a foundation to develop your own use cases into deployable chaincode, complete the following procedure:

1. [Download the Basic contract sample](#downloading).
2. [Create the base contract and implement version control](#create_base).
3. [Define the asset data structure](#define_asset_data). The Basic contact sample defines an asset that is in transit and is being tracked as it moves from one location to another. The asset data includes a unique alphanumeric ID that identifies the asset, and other information, for example, location, temperature, and carrier. For more information, see
4. [Initialize the contract](#initialize_contract).
5. [Define the 'methods' for invoking the contract](#define_invoke_methods).
6. [Define the query methods for how the contract data is read](#define_query_methods).
7. [Define the callbacks](#callbacks).

Detailed information about how to complete each step is hyperlinked to the sections that follow.

### 1. Downloading the Basic samples
{: #downloading}
Download the Basic blockchain contract sample from the [IBM Blockchain contracts](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract_hyperledger) repository on GitHub.  You will need the following Basic sample files, which are provided in the repository folder:

|File name|Description|
|:---|:---|
|**simple_contract_hyperledger.go**|The Basic contract source file|
|**samples.go**|?|
|**schemas.go**|?|  

**Note:** When you install the IBM Blockchain sample environment on Bluemix, the Basic **simple_contract_hyperledger.go** contract is deployed by default.

### 2. Creating the base contract
{: #create_base}

To create the base source file for your Blockchain contract:

1. Create a copy of the Basic contract file,  **simple_contract_hyperledger.go**.
2. Open **simple_contract_hyperledger.go** by using an editor of your choice:
     1. Define the SimpleChaincode struct, as outlined in the following code snippet:
     ```go
     type SimpleChaincode struct {
     }
     ```
     2. Initialize the contract with a version number, as outlined in the following code snippet:
     ```go
     const MYVERSION string = "1.0"
     ```
     3. Define a contract state to keep track of the contract version.
```go
type ContractState struct {
	Version string `json:"version"`
}

var contractState = ContractState{MYVERSION}
```
Note: You can eventually increase the complexity of the ContractState code in your contract, for example, you can add more details into contract state, and you can also introduce asset states alerts and other items, as outlined in later examples.  [Which examples: On this page? or in other samples?]

### 3. Defining the asset data structure
{: #define_asset_data}

For our simple contract, let’s take the example of an asset that is in transit, being tracked as it moves from one location to another.  The asset data would include a unique alphanumeric id that identifies the asset, and other relevant information like location, temperature and carrier.

So lets define the asset data structure. As you can see, Location has two attributes – Latitude and Longitude. You will notice that all the fields are pointers (They have a '\*' in front of the data type). When a JSON string is Marshaled to a struct, if the string does not have a particular field, a default value, based on the datatype gets assigned. Since a blank value is valid for a string in our use case, and a zero a valid number, this means that missing data types would get incorrect data values assigned to this. An easy way to work around this is to use pointers. If a struct's pointer field isn't represented in the input JSON string, that field's value is not mapped.   Another option would be to use maps and marshal the data into the same. Since a map doesn't have a specific data structure, it will only take the data that is sent in, and will not assign default values to fields, simply because there are not fields defined for the map. However, when a specific data strcuture is expected, like from a specific IoT device in our example,  and types need to be validated, it is good to use structs. Another advantage is that when we marshal or unmarshal json to or from a struct, golang makes every effort to match the fields - it looks for prefect matches, case-insensitive match, and doesnt care about any mismatch in the order of the fields. On the other hand, with a map, you will need to handle case mismatches and assert datatypes explicitly for elements of a map.So lets define the asset data structure. As you can see, Location has two attributes – Latitude and Longitude. You will notice that all the fields are pointers (They have a '\*' in front of the data type). When a JSON string is Marshaled to a struct, if the string does not have a particular field, a default value, based on the datatype gets assigned. Since a blank value is valid for a string in our use case, and a zero a valid number, this means that missing data types would get incorrect data values assigned to this. An easy way to work around this is to use pointers. If a struct's pointer field isn't represented in the input JSON string, that field's value is not mapped.   Another option would be to use maps and marshal the data into the same. Since a map doesn't have a specific data structure, it will only take the data that is sent in, and will not assign default values to fields, simply because there are not fields defined for the map. However, when a specific data strcuture is expected, like from a specific IoT device in our example,  and types need to be validated, it is good to use structs. Another advantage is that when we marshal or unmarshal json to or from a struct, golang makes every effort to match the fields - it looks for prefect matches, case-insensitive match, and doesnt care about any mismatch in the order of the fields. On the other hand, with a map, you will need to handle case mismatches and assert datatypes explicitly for elements of a map.

```go
type Geolocation struct {
    Latitude    *float64 `json:"latitude,omitempty"`
    Longitude   *float64 `json:"longitude,omitempty"`
}

type AssetState struct {
    AssetID        *string       `json:"assetID,omitempty"`        // all assets must have an ID, primary key of contract
    Location       *Geolocation  `json:"location,omitempty"`       // current asset location
    Temperature    *float64      `json:"temperature,omitempty"`    // asset temp
    Carrier        *string       `json:"carrier,omitempty"`        // the name of the carrier
}
```


### 4. Initializing the contract
{: #initialize_contract}

Let’s continue with the init function. The 'Init' function is one of the three standard functions expected and mandated in the chaincode, the others being `invoke` and 'Query'. The Init is called as a 'deploy' function, to deploy the chaincode to the fabric. Notice the signature of the function. The Init, Invoke and Query functions share the same arguments - the shim, a function name and an array of strings. Other functions, being called from these standard functions will usually have two arguments, the pointer to the shim and the arg s ttributes. Theshim pointer enables the function to interact with the shim and the args are arguments that need to be consumed by those functions.
```go
// ************************************
// deploy callback mode
// ************************************
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var stateArg ContractState
    var err error
```
We are creating an instance of ContractState that we declared earlier. For this simple example, it will hold the contract version information for an instance of the contract. We are expecting only one argument – a JSON encoded string that sends in the version.
```go
	if len(args) != 1 {
		return nil, errors.New("init expects one argument, a JSON string with tagged version string")
	}
	err = json.Unmarshal([]byte(args[0]), &stateArg)
	if err != nil {
		return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
	}
```
If the version sent in is different from the one in the contract, we have a problem. This feature is included as a safeguard. The business logic of a contract may evolve over time, and multiple versions may co-exist on the same chain, managing different assets. This version check ensures that the right version of the contract is being executed. In this simple example, we will just check if the version sent in matches the one defined in the contract. If not, we simply raise an error and exit gracefully.
```go
    if stateArg.Version != MYVERSION {
        return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
    }
```
If versions match, simply write the contract state to the ledger.
```go
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
```
## Defining invoke methods
{: #define_invoke_methods}

Methods on the `invoke` kind are where the real action happens. We will cover the `createAsset`, `updateAsset` and `deleteAsset` methods of the CRUD interface in this section

#### Create and update asset data

The create and update implementation can be together bundled in a common function called `createOrupdateAsset`. The `createAsset` and `updateAsset` functions make a call to the `createOrupdateAsset` function as outlined in the following example:
```go
/******************** `createAsset` ********************/
func (t *SimpleChaincode) `createAsset`(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
_,erval:=t. createOr`updateAsset`(stub, args)
return nil, erval
}

//******************** `updateAsset` ********************/
func (t *SimpleChaincode) `updateAsset`(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
_,erval:=t. createOr`updateAsset`(stub, args)
return nil, erval
}
```
The reasoning is that if the asset does not already exist in world space, it is a new asset and gets created. If on the other hand, it already exists the new record is considered an update and the new values that come in are updated. For example, if, in the new record, only temperature and location has been sent in, only those values of the asset get updated in the shim; Carrier data will be maintained as previously received, if any.
createOr`updateAsset`
Now, let us take a look at this common function. The first step is a call to ‘`validateInput`’, explained later in the section ‘Validate Input data’.
```go
stateIn, err = t.`validateInput`(args)
if err != nil {
return nil, err
}
```
Once this is done, a quick call to `shim.GetState` tells us whether the asset already exists. If not, it is a create scenario and the data is saved as received. If the asset exists, a call is made to `mergePartialState` method to merge old and new data and the updated record gets written to the shim.

```go
//Partial updates introduced here
// Check if asset record existed in stub
assetBytes, err:= stub.GetState(assetID)
if err != nil || len(assetBytes)==0{
// This implies that this is a 'create' scenario
stateStub = stateIn // The record that goes into the stub is the one that cme in
} else {
// This is an update scenario
err = json.Unmarshal(assetBytes, &stateStub)
if err != nil {
err = errors.New("Unable to unmarshal JSON data from stub")
return nil, err
// state is an empty instance of asset state
}
// Merge partial state updates
stateStub, err =t.mergePartialState(stateStub,stateIn)
if err != nil {
err = errors.New("Unable to merge state")
return nil,err

}
}
```

## 5. Defining the 'query' methods
{: #define_query_methods}

Use a `query` method to read data pertaining to the blockchain contract. The Basic contract sample uses the following query implementation methods:

- Read asset data (readAsset) method
- Read asset object model (readAssetObjectModel) method

#### Read asset data (readAsset) method
The`readAsset` method returns the asset data that is stored in the ledger for the asset ID that is input. Method signature is quite similar to the other methods and so is the call to `validateInput`. The part where it differs is that in this method, we use the asset id to fetch the asset data from the state and return it, raising an error if the asset data does not exist. You will notice that we are using the same stub GetState call here that we've previously come across.
```go
    assetID = *stateIn.AssetID
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
```
#### Read asset object model (readAssetObjectModel) method
The `readAssetObjectModel` method returns the object model of the Asset dataset that the contract expects. This helps the UI and mapping component understand the structure of the input JSON string expected by the contract. We will see examples of how this the data it returns looks like, shortly.
The signature is similar to the others because this is a 'query' method. What we do here is create an empty instance of the `AssetState` definition, marshal it to a JSON string and return it so that the caller knows the input data structure.
```go
func (t *SimpleChaincode) readAssetObjectModel(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var state `AssetState` = AssetState{}

	// Marshal and return
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}
	return stateJSON, nil
}
```
#### Read asset sample
The read asset sample returns a sample of the schema to the consumer of the contract by reading the contents of the **samples.go** file that is in the project folder.  

#### Read asset schema
The read asset schema sample returns the schema for the contract in JSON format to the consumer of the contract by reading the contents of the **schemas.go** file. The **schemas.go** and the **samples.go** files are provided with the Basic contract sample, whereas the **trade_lane_contract_hyperledger** advanced sample demonstrates how you can generate these files by using the `go generate` command.


#### Deleting asset data
Use the `deleteAsset` method to remove asset data from the current blockchain state by asset ID. Blockchain is a decentralized transaction ledger. Every single validated transaction is retained in the blockchain. When the record for a specific asset is created and updated in the blockchain, all activity is recorded in the ledger. The `deleteAsset` method removes all data for the specified asset from the current state but does not remove it from the blockchain. If you need to retrieve or view any data pertaining to an asset that was deleted from the current state, you can still retrieve it from the blockchain ledger if you need to.

The following code snippet provides an example of how you can delete assets from the current state:
```go
    ........
    ........
   assetID = *stateIn.AssetID
    // Delete the key / asset from the ledger
    err = stub.DelState(assetID)
    if err != nil {
        err = errors.New("DELSTATE failed! : "+ fmt.Sprint(err))
       return nil, err
    }
    .......
    .......
```
**Important:** IBM Blockchain retains every single validated transaction in a decentralized peer-to-peer fashion. Peers can keep a copy of the complete ledger, meaning that every piece of data can be potentially retrieved and inspected at any point in time. IBM Blockchain maintains a state database in the context of a contract, which is somewhat analogous to Ethereum's state database or Bitcoin's UTXO set. Use the asset delete feature to remove your asset data from the current state as and when you need to. An advantage of the peer-to-peer feature is that in contracts where history is maintained, typically you query the state database rather than finding all transactions for a given contract and simulating the state history yourself.

## Invoking the contract
{: #invoking_the_contract}

We have created a simple contract which implements a CRUD interface and enables us to manage Watson IoT Platform asset data on the blockchain. Great! Now how do we use it? IBM Blockchain contracts expose REST APIs and these can be invoked through the command-line, or we could make REST API calls using, say, Swagger or Node.js. Later on, we will see an example of a sample UI that provides us with a more user-friendly way of calling the contract.
No matter which method you use, the actual syntax of the calls made to a chaincode will not change. Like we saw earlier, IBM Blockchain understands three types of chaincode calls - 'deploy', `invoke` and 'query'. Every call expects a JSON string as input. This string will include a Function key and an Args key. The Function key is the name of the actual function you are calling and the Args are the arguments you want to pass on to it.

IBM Blockchain github pages clearly explain how to set up the [DevNet](https://github.com/hyperledger/fabric/blob/master/docs/dev-setup/devnet-setup.md) and [Sandbox](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md) environments and how to test contracts in the command-line. The Sandbox environment is particularly useful when writing new contracts or modifying existing ones - it provides us with an environment to test and debug chaincode without having to set up a complete IBM Bluemix Hyperledger network. Let's see some examples of the command-line and rest calls to our sample chaincode.
Once the network is running, contract is built and instance registered, (all covered in the links above) we can start calling our contract.If we are using a Swagger or other REST interface, the HTTP server and the Swagger  / Node.js setup needs to be in place as described [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md). 	

### Contract calls
**Note:** In Sandbox, we get to set the contract instance name. In a real peer network, when we register the contract, it will have a lengthy alphanumeric name that is returned by the network and that is the name that needs to be used. Alternately we can use the github path to point to the contract.

### Deploy
This is an example call to 'deploy'. The contract name in the Sandbox is ex01. Notice the signature. We are calling the peer executable with the chaincode deploy command. We specify the name of the contract as ex01 and the arguments, in a JSON format are the Function name 'init' and the arguments it expects inside Args: []. IBM blockchain allows the content of Args[] to be any string(s), however, since our contract is about IOT data, which comes in as json, we have written our contract to expect a json string inside Args[].
In command line Sandbox mode, in the tab where you have registered the contract, you can see the response of the Shim, stating that init was received and was successful. If you have print statements in the contract, those can also be seen here. They are useful for debugging, but are best removed once the contract is deployed. IBM Blockchain 'deploy' and `invoke`s are asynchronous.
```
./peer chaincode deploy -n ex01 -c '{"Function":"init", "Args": ["{\"version\":\"1.0\"}"]}'
```
A REST call using curl would look like this:

```
curl -X POST --header 'Content-Type: application/json' --header 'Accept:
application/json' -d '{
  "type": "GOLANG",
  "chaincodeID": {
    "name": "ex01"
  },
  "ctorMsg": {
    "function": "init",
    "args": [
      "{\"version\":\"1.0\"}"
    ]
  }
}' 'http://127.0.0.1:3000/devops/deploy'
```
And a Swagger call, using the IBM Blockchain REST API's deploy section would look like:
```
{
  "type": "GOLANG",
  "chaincodeID": {
    "name": "ex01"
  },
  "ctorMsg": {
    "function": "init",
    "args": [
      "{\"version\":\"1.0\"}"
    ]
  },
  "secureContext": "string",
  "confidentialityLevel": "PUBLIC"
}
```
As you can see,  the syntax of the actual function call does not change based on how you call the contract. You need to specify the chaincode method, the function and arguments.

Now let us look at how to call the other functions we have implemented.

### readAssetSchemas
Let's see how our asset's Schema looks like, so that we know what functions to call, and what parameters to pass for the function
```
./peer chaincode query -l golang -n ex01 -c '{"Function":"readAssetSchemas", "Args":[]}'
```
This call returns a lengthy JSON object that lists all the functions and properties of the functions, with descriptions. It also indicates clearly which properties are mandatory. The schema can be consumed by applications for easy interaction with the contract. In fact the monitoring_ui example does just this.

readAssetSamples
This function gives us an example  - a sample of what could be passed in.
```
./peer chaincode query -l golang -n ex01 -c '{"Function":"readAssetSamples", "Args":[]}'
```
Now we know what to send in, lets call the other methods.

### Invoke methods

#### createAsset
```
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"`createAsset`", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":34, \"longitude\":23}, \"temperature\":89, \"carrier\":\"ABCCARRIER\"}"]}'
```
#### readAsset
```
 ./peer chaincode query -l golang -n n ex01 -c '{"Function":"readAsset", "Args":["{\"assetID\":\"CON123\"}"]}'
```
 This returns
 ```
 {"assetID":"CON123","location":{"latitude":34,"longitude":23},"temperature":89,"carrier":"ABCCARRIER"}
 ```
#### updateAsset
```
./peer chaincode invoke -l golang -n  ex01  -c '{"Function":"`updateAsset`", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":56, \"longitude\":23}, \"temperature\":78, \"carrier\":\"PQRCARRIER\"}"]}'
```
another`readAsset`with return the updated asset data

#### deleteAsset
```
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"deleteAsset", "Args":["{\"assetID\":\"CON123\"}"]}
```

### Callbacks
In IBM Blockchain contracts, the Init method handles init commands, the ones called as part of 'deploy' during execution. The Invoke method handles `invoke` method calls and the 'Query' method handles invocations of type 'query'. We've already implemented 'Init'. Now lets write simple Invoke and Query methods that will call our functions -`createAsset`, `updateAsset`, `deleteAsset`, readAsset, `readAssetObjectModel`, readAssetSchemas and readAssetSamples methods, based on the function name that is passed to the contract invocation. The following example is for the `invoke` implementation
```go
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    // Handle different functions
    if function == "`createAsset`" {
        // create assetID
        return t.`createAsset`(stub, args)
    } else if function == "`updateAsset`" {
        // create assetID
        return t.`updateAsset`(stub, args)
    } else if function == "deleteAsset" {
        // Deletes an asset by ID from the ledger
        return t.deleteAsset(stub, args)
    }
    return nil, errors.New("Received unknown invocation: " + function)
}
```
The following example is for the 'Query' implementation
```go
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    // Handle different functions
    if function == "readAsset" {
        // gets the state for an assetID as a JSON struct
        return t.readAsset(stub, args)
    } else if function =="readAssetObjectModel" {
        return t.readAssetObjectModel(stub, args)
    }  else if function == "readAssetSamples" {
  // returns selected sample objects
  return t.readAssetSamples(stub, args)
 } else if function == "readAssetSchemas" {
  // returns selected sample objects
  return t.readAssetSchemas(stub, args)
 }
    return nil, errors.New("Received unknown invocation: " + function)
}
```
#### Validate your input data (validateInput)

The ``validateInput`` function is not an implementation of an IBM Blockchain deploy, invoke, or query methods, but is a common function that is called by all of the `invoke` type functions (`createAsset`, `updateAsset` and `deleteAsset`) and the`readAsset`'query' function that does the following:
1. Validates the input data for the right number of arguments,
2. Unmarshals the input data to the asset state and
3. Checks for the asset id.
```go
func (t *SimpleChaincode) `validateInput`(args []string) (stateIn AssetState, err error) {
    var assetID string // asset ID
    var state AssetState = AssetState{} // The calling function is expecting an object of type AssetState

    if len(args) !=1 {
        err = errors.New("Incorrect number of arguments. Expecting a JSON strings with mandatory assetID")
        return state, err
    }
    jsonData:=args[0]
    assetID = ""
    stateJSON := []byte(jsonData)
    err = json.Unmarshal(stateJSON, &stateIn)
    if err != nil {
        err = errors.New("Unable to unmarshal input JSON data")
        return state, err
        // state is an empty instance of asset state
    }      
    // was assetID present?
    // The nil check is required because the asset id is a pointer.
    // If no value comes in from the json input string, the values are set to nil

    if stateIn.AssetID !=nil {
        assetID = strings.TrimSpace(*stateIn.AssetID)
        if assetID==""{
            err = errors.New("AssetID not passed")
            return state, err
        }
    } else {
        err = errors.New("Asset id is mandatory in the input JSON data")
        return state, err
    }
    stateIn.AssetID = &assetID
    return stateIn, nil
}
```



### mergePartialState
The Basic contract supports partial updates of data. The `mergePartialState` function iteratively scans old and new records, replacing the old values with the new values when it identifies a field with updates. The merged record is then used by the `updatepdateAsset` function.

 The following snippet provides an example of how the `mergePartialState` function is defined:
```go
func (t *SimpleChaincode) mergePartialState(oldState AssetState, newState AssetState) (AssetState,  error) {

    old := reflect.ValueOf(&oldState).Elem()
    new := reflect.ValueOf(&newState).Elem()
    for i := 0; i < old.NumField(); i++ {
        oldOne:=old.Field(i)
        newOne:=new.Field(i)
        if ! reflect.ValueOf(newOne.Interface()).IsNil() {
            oldOne.Set(reflect.Value(newOne))
        }
    }
    return oldState, nil
 }
 ```
### Main function
The 'main' function creates a 'shim' instance.

Note that we have a fmt.Printf command here. Printf and Println commands can easily be used to debug the chaincode as we write / modify and test chaincode. However, since the chaincode runs inside a docker container, these will get written to the log of the docker instance and space gets filled up very quickly, especially if the environment is running in debug mode. It is prudent to limit or avoid Print statements in actual deployed chaincode as the chaincode is essentially supposed to run in the background.

Best practices around docker log management, like Logrotate (supported since Docker 1.8) could be adopted to avoid space issues. Contract logging implementations could also be incorporated in future contracts.
```
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}f
}
```

## Next steps
{: #next_steps}

After you create your contract, upload it to IBM Blockchain with Watson IoT Platform or by using the API. The Watson IoT Platform includes a data mapping component to route data to the Blockchain contract. For more information, see [Developing for blockchain](https://console.ng.bluemix.net/docs/services/IoT/blockchain/dev_blockchain.html)

There are more sample contracts available for download in the IBM Blockchain samples folder on [GitHub](https://github.com/ibm-watson-iot/blockchain-samples). Experiment with the available samples, for example, the Trade Lane blockchain contract. Use the samples to enhance your Blockchain smart contracts further.

IBM Blockchain also provides a sandbox environment for you to test contract code before deploying it to your peer network. For more information about how to set up the sandbox environment, see [Writing, building, and running chaincode in a development environment](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md).
