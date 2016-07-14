
# The Basic smart contract sample

The Basic contract is a sample hyperledger blockchain contract that is provided by IBM to help you to get started with blockchain development and integration on the IBM Watson IoT Platform. You can use the Basic contract sample to create a blockchain contract that tracks and stores asset data from a device that is connected to your Watson IoT Platform organization.

The following information is provided to help you to get started with the Basic sample:

- [Overview](#overview)
- [Requirements](#requirements)
- [Developing your IBM Blockchain hyperledger by using the Basic contract sample](#instructions)
- [Invoking the contract](#invoke_contract)
- [Generating code](#generate_code)
- [Next steps](next_steps)

## Overview
The Basic blockchain sample, **simple_contract_hyperledger.go**, is the default smart contract for getting started with writing blockchain contracts. The Basic contract includes create, read, update, and delete asset data operations for tracking device data on the IBM Blockchain ledger. Like all other IBM Blockchain contracts, the Basic contract sample is developed in the Go programming language. IBM plan to support contracts that are written in other languages soon.

You can run your smart contracts for Blockchain from either the command line or from a REST interface. For more information, see
[Developing smart contracts for Watson IoT Platform blockchain integration](https://console.ng.bluemix.net/docs/services/IoT/blockchain/dev_blockchain.html).

As outlined in Basic sample, IBM Blockchain contract code includes the following methods:

|Method|Description|
|:---|:---|
|`deploy`|`Used to deploy a smart contract`|  
|`invoke`|`Used to update a smart contract`|
|`query`|`Used to query a smart contract`|

**Note:** In the Basic contract sample, the `deploy` method is called `init`.

When you call any of the methods, you must first pass a JSON string that includes the function name and a set of arguments as key-value pairs to the chain code instance. You can define multiple methods within the same contract.

[Where should this go?]
To create a simple contract to create, read, update, and delete asset data, use the following methods:

|Method|Provides|
|:---|:---|
|'ReadAssetSchema'|The methods and associated properties of the JSON schema contract|  
|'ReadAssetSamples'|An example of the sample JSON data|

For more information about setting up REST and Swagger, see [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md).

## Requirements
{: #requirements}

The Watson IoT Platform includes a data mapping component to route your data to the Blockchain contract. To correctly map the Watson IoT Platform event properties to the corresponding Blockchain contract properties, use the following required functions:

|Function|Used by the data mapping function to....|
|:---|:---|
|`updateAsset`|<ul><li>Create asset records if one doesn't exist on the ledger</li><li>Update asset data for existing records</li><li>Map the asset ID to the Watson IoT Platform serial ID</li><li>Accept IoT device input data in the form of JSON strings</li></ul>|
|`readAssetSchemas`|Expose the function names and properties that are required by the contract so that the data mapper can correctly map the event properties to the contract properties


## Developing your IBM Blockchain hyperledger by using the Basic contract sample
{: #instructions}

The Basic contract is an example recipe that is designed for you to customize and use to experiment with contract development for Blockchain. The current version of the Basic contract is simplistic with similar functions to create and update asset data.

To use the Basic sample **simple_contract_hyperledger.go** sample contract as a foundation to develop your own use cases into deployable chaincode, complete the following procedure:

1. [Download the Basic contract sample](#downloading).
2. [Create the base contract and implement version control](#create_base).
3. [Define the asset data structure](#define_asset_data). The Basic contact sample defines an asset that is in transit and is being tracked as it moves from one location to another. The asset data includes a unique alphanumeric ID that identifies the asset, and other information, for example, location, temperature, and carrier. For more information, see
4. [Initialize the contract](#initialize_contract).
5. [Define the invoke methods](#define_invoke_methods).
6. [Define the query methods for how the contract data is read](#define_query_methods).
7. [Define the callbacks](#define_callbacks).
8. [Develop the contract further](#more_methods)

Detailed information about how to complete each step is hyperlinked to the sections that follow.

### 1. Downloading the Basic samples
{: #downloading}
Download the Basic blockchain contract sample from the [IBM Blockchain contracts](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract_hyperledger) repository on GitHub.  You need the following Basic sample files, which are provided in the repository folder:

- **simple_contract_hyperledger.go**|The Basic contract source file
- **samples.go**|?|
- **schemas.go**|?|  

**Note:** When you install the IBM Blockchain sample environment on Bluemix, the Basic **simple_contract_hyperledger.go** contract is deployed by default.


### 2. Creating the base contract
{: #create_base}

To create the base source file for your Blockchain contract:

1. Create a copy of **simple_contract_hyperledger.go**, which is the main source file of the Basic contract.
2. Using an editor of your choice, open **simple_contract_hyperledger.go**.
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
Note: You can eventually increase the complexity of the `ContractState` code in your contract, for example, you can add more contract state details, introduce asset state alerts, and other items that are outlined in other more advanced examples.

### 3. Defining the asset data structure
{: #define_asset_data}

The Basic contract provides the Blockchain contract code that is required for an asset that is in transit and is also being tracked as it moves from one location to another. In this example, the following asset data is tracked:

- Unique alphanumeric ID that identifies the asset
- Location
- Temperature
- Carrier

The following code provides an example of how to define the data structure for the example scenario of an asset in transit:

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

#### Location

Location has two attributes, latitude and longitude.

#### Pointers

All the fields in the data structure definition are defined as pointers. A pointer is indicated by a preceding asterix character ('\*'), for example, '\*'float64.

When a JSON string is marshaled to a `struct`, if the string does not have a particular field, a default value for the data type is assigned. Because a blank value is valid for a string, and a zero (0) character is also a valid number, incorrect data values might be assigned to fields with missing data types. To work around this problem, use pointers. If the pointer field of a struct isn't represented in the input JSON string, that field's value is not mapped. Another way to work around the problem is to feed the data into a map. Unlike a 'struct', a map doesn't have a specific data structure and accepts the submitted data without requiring and assigning default field values. However, when a specific data structure is expected and types must be validated, it is good to use structs, for example, in the case of a specific IoT device in this example.

Another advantage of marshaling JSON values to or from a struct is that Golang makes every effort to match the fields and look for perfect case-insensitive matches, ignoring the order of the fields. Whereas when you define map elements, you need to handle case mismatches and assert data types explicitly.


### 4. Initializing the contract
{: #initialize_contract}

The `init` function is one of the three required functions of the chaincode and initializes the contract. The other required functions of the chaincode are `invoke` and `query`. The `init` function is called as a 'deploy' function to deploy the chaincode to the fabric. Notice the signature of the function. The `init`, `invoke`, and `query` functions share the following arguments:

|Argument|Description|
|:---|:---|
|Shim |? |
|Function name| |
|Array of strings| |

Other functions that are called from the standard functions typically contain the following arguments:

|Argument|Description|
|:---|:---|
|Pointer to the shim |Enables the function to interact with the shim|
|Function name| |
|Arg attributes|Arguments that are to be consumed by those functions|

```go
// ************************************
// deploy callback mode
// ************************************
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var stateArg ContractState
    var err error
```

The following example shows how you can create an instance of ContractState, which was declared earlier in step 2.

```go
	if len(args) != 1 {
		return nil, errors.New("init expects one argument, a JSON string with tagged version string")
	}
	err = json.Unmarshal([]byte(args[0]), &stateArg)
	if err != nil {
		return nil, errors.New("Version argument unmarshal failed: " + fmt.Sprint(err))
	}
```
**Note:** In this example, the contract version information is held for an instance of the contract. Only one argument is exposed, which is a JSON encoded string that submits the version value.

If the version that is submitted is different from the version in the contract, an error scenario occurs, which is a validation feature. The business logic of a contract can evolve and change over time, and multiple versions of the same contract might be expected to co-exist on the same chain, managing different assets. The version check validation feature ensures that the correct version of the contract is run.

The following example checks to see whether the version submitted matches the version that is defined in the contract. If the version does not match the version that is defined in the contract an error condition is raised and the contract code stops running.

```go
    if stateArg.Version != MYVERSION {
        return nil, errors.New("Contract version " + MYVERSION + " must match version argument: " + stateArg.Version)
    }
```
If the versions match, the contract code continues, and the following code to write the contract state to the ledger is run:
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
## 5. Defining the invoke methods
{: #define_invoke_methods}

Define the `invoke` methods for the create, read, update, and delete operations in your contract, which is where most of the contract action occurs. The `createAsset`, `updateAsset` and `deleteAsset` methods of the create, retrieve, update, and delete interface are explained in this section.

#### Create and update asset data

You can bundle the create and update implementation together in a common function called `createOrupdateAsset`. The `createAsset` and `updateAsset` functions make a call to the `createOrupdateAsset` function as outlined in the following example:
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
If the asset doesn't exist, it is classified as a new asset and is created. If the asset does exists, the new record is classified as an update and the asset is updated with the new values that are retrieved. For example, if, in the new record, only temperature and location values are submitted, only those values of the asset get updated in the shim. If carrier data exists, it is maintained in the state that it was previously received.

- `
The first step of the `createOrupdateAsset` common function is a call to `validateInput`, which is explained later in`‘Validate input data`.

```go
stateIn, err = t.`validateInput`(args)
if err != nil {
return nil, err
}
```
After you run the code, to confirm whether the asset exists, insert a call to the `shim.GetState` method. If the asset does not exist, it is a create scenario and the data is saved as received. If the asset exists, a call is made to the `mergePartialState` method to merge old and new data and the updated record gets written to the shim.

```go
//Partial updates introduced here
// Check if asset record existed in stub
assetBytes, err:= stub.GetState(assetID)
if err != nil || len(assetBytes)==0{
// This implies that this is a 'create' scenario
stateStub = stateIn // The record that goes into the stub is the one that came in
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

## 6. Defining the 'query' methods
{: #define_query_methods}

Use a `query` method to define how the contract data is read. The Basic contract sample uses the following blockchain query implementation methods:

- Read asset data (readAsset) method
- Read asset object model (readAssetObjectModel) method

#### Read asset data (readAsset) method
The`readAsset` method returns the asset data that is stored in the ledger for the asset ID that is input. Method signature is similar to the other methods as is the call to `validateInput`. The only difference is that in this method, the asset ID is used to fetch the asset data from the state and return it, and an error is raised if the asset data does not exist. Note that the same stub GetState call that was used previously in this sample is also used here.
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
         err = errors.New("Unable to unmarshal state data that is obtained from ledger")
        return nil, err
    }
```
#### Read asset object model (readAssetObjectModel) method
The `readAssetObjectModel` method returns the object model of the asset data set that the contract expects. The output of the `readAssetObjectModel` method helps the user interface and mapping component understand the structure of the input JSON string that is expected by the contract. The signature is similar to the others because this is a `query` method.

In the following example, the `readAssetObjectModel` method creates an empty instance of the `AssetState` definition, marshals it to a JSON string, and then returns it so that the caller knows the input data structure.

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
The read asset schema sample returns the schema for the contract in JSON format to the consumer of the contract by reading the contents of the **schemas.go** file.

The **schemas.go** and the **samples.go** files are provided with the Basic contract sample. The  **trade_lane_contract_hyperledger** advanced contract sample demonstrates how you can generate schema files by using the `go generate` command.


## 7. Defining the callbacks
{: #define_callbacks}

In IBM Blockchain contracts, the `init` method handles `init` commands, which are the ones that are called as part of the 'deploy' phase when the code is run. The `invoke` method handles `invoke` method calls and the `query` method handles invocations of type 'query'. In the previous step, you implemented the 'init' comand, so you are now ready to write some simple invoke and query methods to call the following functions:

- `createAsset`
- `updateAsset`
- `deleteAsset`
- `readAsset`
- `readAssetObjectModel`
- `readAssetSchemas`
- `readAssetSamples`

The following example is for the `invoke` implementation.

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
The following example is for the `query` implementation:

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
        err = errors.New("Asset ID is mandatory in the input JSON data")
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

To avoid space issues, implement a procedure for docker log management, for example, you can implement Logrotate, which is supported since Docker V1.8. You can also optionally incorporate contract logging implementations in your contracts.
```
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}f
}
```


### 8. Developing the contract further
{: #more_methods}

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

When you complete all of the required steps in [Developing your IBM Blockchain hyperledger by using the Basic contract sample](#instructions) you will have a simple contract to manage Watson IoT Platform asset data on the blockchain. Your contract will use a simple create, retrieve, update, and delete operations interface. The next step is to use the contract by invoking it. IBM Blockchain contracts expose REST APIs and these can be invoked through the command-line, or you can make REST API calls by using either Swagger or Node.js. No matter which method you use, the actual syntax of the calls that are made to a chaincode does not change. Later on in the more advanced contacts, you can see an example of a sample UI that provides a more user-friendly way of calling the contract.

As mentioned earlier, IBM Blockchain understands `deploy`, `invoke`, and `query` calls, of which must contain a JSON string as input. The input string must also include a Function key and an Args key. The Function key is the name of the actual function you are calling, and the Args are the arguments you want to pass on to it.

There are several IBM Blockchain resources in GitHub that explain in more detail how you can set up the [DevNet](https://github.com/hyperledger/fabric/blob/master/docs/dev-setup/devnet-setup.md) and [Sandbox](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md) environments, and also how to test contracts in the command line. The Sandbox environment is particularly useful for writing new contracts or modifying existing ones as it provides you with an environment to test and debug chaincode without needing to set up a complete IBM Bluemix Hyperledger network. Let's see some examples of the command-line and rest calls to our sample chaincode.

When the network is running, the contract is built, and your instance is registered, you can start calling the contract that you created. If you are using Swagger or another REST interface, the HTTP server and the Swagger or Node.js must be set up as outlined [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md). 	

### Contract calls
**Note:** In the Sandbox, you can specify the contract instance name. In a real peer network, when you register the contract, it will have a lengthy alphanumeric name that is returned by the network and that is the name that you must use. Alternately you can use the GitHub path to point to the contract.

### Deploy

The following code provides an example call to `deploy` your contract:


```
./peer chaincode deploy -n ex01 -c '{"Function":"init", "Args": ["{\"version\":\"1.0\"}"]}'
```

Where:

- Contract name in the Sandbox is ex01
- Peer executable is called with the chaincode deploy command
- Name of the contract is specified as 'ex01'
- Arguments are in JSON format
- Function name is 'init'
- Arguments are listed inside Args: []

Note: Make a note of the signature. Also, note that in IBM Blockchain the content of Args[] can be any string(s), however, since our contract is about Watson IoT Platform data, which is in JSON format, the sample contract requires a JSON string inside Args[].

In command line Sandbox mode, in the tab where you egistered the contract, you can see the response of the Shim, stating that 'init' was received and was successful. If your contract includes print statements, which are useful for debugging, they will also be displayed. Remove print statements when the contract is deployed. IBM Blockchain `deploy` and `invoke` calls are asynchronous.

A REST call using `curl` would look like the following code example:

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
The following snippet provides an example of a Swagger call that uses the deploy section of the IBM Blockchain REST API:
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
The syntax of the function call does not change based on how you call the contract. You need to specify the chaincode method, the function, and the arguments.

You can also explore how to call the other functions that were previously implemented.

### readAssetSchemas
Let's see how our asset's Schema looks like, so that we know what functions to call, and what parameters to pass for the function:
```
./peer chaincode query -l golang -n ex01 -c '{"Function":"readAssetSchemas", "Args":[]}'
```
The following call returns a lengthy JSON object that lists all functions and their properties, with descriptions. The object also indicates clearly which properties are mandatory. The schema can be consumed by applications for easy interaction with the contract. In fact the monitoring_ui example does just this.

###readAssetSamples

The readAssetSamples function provides a sample of what might be passed in.
```
./peer chaincode query -l golang -n ex01 -c '{"Function":"readAssetSamples", "Args":[]}'
```
When you establish what to pass in, you can explore and call the other methods.

### Invoke methods

#### createAsset
```
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"`createAsset`", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":34, \"longitude\":23}, \"temperature\":89, \"carrier\":\"ABCCARRIER\"}"]}'
```
#### readAsset
```
 ./peer chaincode query -l golang -n n ex01 -c '{"Function":"readAsset", "Args":["{\"assetID\":\"CON123\"}"]}'
```
 Which returns:
 ```
 {"assetID":"CON123","location":{"latitude":34,"longitude":23},"temperature":89,"carrier":"ABCCARRIER"}
 ```
#### updateAsset
```
./peer chaincode invoke -l golang -n  ex01  -c '{"Function":"`updateAsset`", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":56, \"longitude\":23}, \"temperature\":78, \"carrier\":\"PQRCARRIER\"}"]}'
```
another`readAsset`with return the updated asset data [??]

#### deleteAsset
```
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"deleteAsset", "Args":["{\"assetID\":\"CON123\"}"]}
```


## Generating code
{: #generate_code}

Depending on the content of your contract, you might need to regenerate the deployable contract schema and sample from the GoLang file. for example, if you include a '//'' go generate directive in your code. To generate contract code, use the following `go generate` command:

```
//go:generate go run scripts/generate_go_schema.go
```

## Next steps
{: #next_steps}

After you create your contract, upload it to IBM Blockchain with Watson IoT Platform or by using the API. The Watson IoT Platform includes a data mapping component to route data to the Blockchain contract. For more information, see [Developing for blockchain](https://console.ng.bluemix.net/docs/services/IoT/blockchain/dev_blockchain.html)

There are more sample contracts available for download in the IBM Blockchain samples folder on [GitHub](https://github.com/ibm-watson-iot/blockchain-samples). Experiment with the available samples, for example, the Trade Lane blockchain contract. Use the samples to enhance your Blockchain smart contracts further.

IBM Blockchain also provides a sandbox environment for you to test contract code before deploying it to your peer network. For more information about how to set up the sandbox environment, see [Writing, building, and running chaincode in a development environment](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md).
