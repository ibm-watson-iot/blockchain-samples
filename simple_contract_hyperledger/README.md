
# Using the Basic contract sample for Watson IoT Platform integration

The Basic contract is a sample hyperledger blockchain contract that is provided by IBM to help you to get started with blockchain development and integration on the IBM Watson IoT Platform.

You can use the Basic contract sample to create a blockchain contract that tracks and stores asset data from a device that is connected to your Watson IoT Platform organization. The Basic blockchain sample is the default smart contract for getting started with writing blockchain contracts. The contract is developed in GoLang and includes create, read, update, and delete asset data operations for tracking device data on the IBM Blockchain ledger.   


## Downloading the Basic sample
You can download the Basic blockchain contract sample from the [IBM Blockchain contracts](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract_hyperledger) repository on GitHub. The Basic contract is also deployed by default in the sample IBM Blockchain environment that you can deploy on Bluemix.

Contract code in IBM Blockchain can be written today in the ‘Go’ programming language, and future support is planned for other languages. Contract code written for IBM Blockchain can include three types of methods:

- `deploy`
- `invoke`
- `query`

You can execute contract code from either the command line or from a REST interface. Both methods are possible with the Basic blockchain contract sample, as outlined on this page.

To call a method, you must pass a JSON string to the chaincode instance. This string must include the function name and required arguments passed as key-value pairs. You can have more that one method of a type defined in the contract.   

For more information about setting up REST and Swagger can be found [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md).  


## Creating the base contract

The first step is to create the base source file for your contract, and then you can map your own use cases into deployable chain code.

To create a simple contract to create, read, update and delete asset data, you will need to use the following following methods:

|Method|Provides|
|:---|:---|
|'ReadAssetSchema'|The JSON schema contract's methods and associated properties|
|'ReadAssetSamples'|An example of the sample JSON data|


1. Create the Go (.go) source file for your contract by using an editor of your choice.
2. Define the SimpleChaincode struct:
To comply with the chaincode implementation approach, define the SimpleChaincode struct as outlined in the following code snippet:
```go
type SimpleChaincode struct {
}
```
3. To initialize the contract with a version number, define a constant for the version, as outlined in the following code snippet:
```go
const MYVERSION string = "1.0"
```
In our simple example, the ‘deploy’ method is called ‘init’.
4. Define a contract state, which in this example will keep track of the contract version.

In later examples as we build more complexity into our contract code, we will add more details into contract state and introduce asset states, alerts, etc.
```go
type ContractState struct {
	Version string `json:"version"`
}

var contractState = ContractState{MYVERSION}
```
5. Define the asset data structure.


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

6. Initialize the contract.


### Initializing the contract
Let’s continue with the init function. The 'Init' function is one of the three standard functions expected and mandated in the chaincode, the others being 'Invoke' and 'Query'. The Init is called as a 'deploy' function, to deploy the chaincode to the fabric. Notice the signature of the function. The Init, Invoke and Query functions share the same arguments - the shim, a function name and an array of strings. Other functions, being called from these standard functions will usually have two arguments, the pointer to the shim and the arg s ttributes. Theshim pointer enables the function to interact with the shim and the args are arguments that need to be consumed by those functions.
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
### Invoke methods
Methods on the 'invoke' kind are where the real action happens. We will cover the 'createAsset', 'updateAsset' and 'deleteAsset' methods of the CRUD interface in this section

#### Create and update asset data

The create and update implementation can be together bundled in a common function called ‘createOrUpdateAsset’. So, both createAsset and updateAsset functions simply make a call to createOrUpdateAsset, like so:
```go
/******************** createAsset ********************/
func (t *SimpleChaincode) createAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
_,erval:=t. createOrUpdateAsset(stub, args)
return nil, erval
}

//******************** updateAsset ********************/
func (t *SimpleChaincode) updateAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
_,erval:=t. createOrUpdateAsset(stub, args)
return nil, erval
}
```
The reasoning is that if the asset does not already exist in world space, it is a new asset and gets created. If on the other hand, it already exists the new record is considered an update and the new values that come in are updated. For example, if, in the new record, only temperature and location has been sent in, only those values of the asset get updated in the shim; Carrier data will be maintained as previously received, if any.
createOrUpdateAsset
Now, let us take a look at this common function. The first step is a call to ‘validateInput’, explained later in the section ‘Validate Input data’.
```go
stateIn, err = t.validateInput(args)
if err != nil {
return nil, err
}
```
Once this is done, a quick call to shim.GetState tells us whether the asset already exists. If not, it is a create scenario and the data is saved as received. If the asset exists, a call is made to MergePartialState method to merge old and new data and the updated record gets written to the shim.

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

#### Delete Asset Data
The deleteAsset method enables asset data to be removed from the state, by asset ID. The blockchain is a decentralized transaction ledger. Every single validated transaction will be retained by the blockchain. So if an asset's record was created and updated in the blockchain, all that information will still reside in the ledger. The 'delete' we are performing here is just a removal of asset data from the current state. In future, if there was a need to retrieve and inspect data pertaining to the deleted asset, it can be retrieved from the blockchain. This is an important distinction. We are not eliminating the asset information from the ledger, we are only removing it from the current state as pertains to the contract.

**Important:** Some notes about deleting assets  
An important point to note is that a blockchain retains every single validated transaction in a decentralized, peer-to-peer fashion. This means that peers can keep a copy of the complete ledger and that every piece of data can be potentially retrieved and inspected at any point in time. So why are we talking about deleting an asset?
IBM Blockchain allows us to maintain a state database in the context of a contract (somewhat analogous to Ethereum's state database or Bitcoin's UTXO set). The delete feature is there to enable us to remove the data as required from the state. The advantage is that in contracts where history is maintained, typically you will query this state database rather than finding all transactions for a given contract and simulating the state history yourself.
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
### 'Query' methods
Query methods, as the name implies, simply allow us to read data pertaining to the contract. We have two implementations of query in this example - the readAsset and the readAssetObjectModel

#### Read Asset Data
This is a method of type 'query' as understood by IBM Blockchain. The asset id is expected as input and the method returns the asset data, as stored in the ledger. Method signature is quite similar to the other methods and so is the call to validateInput. The part where it differs is that in this method, we use the asset id to fetch the asset data from the state and return it, raising an error if the asset data does not exist. You will notice that we are using the same stub.GetState call here that we've previously come across.
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
#### Read the Asset Object Model
This method returns the Object model of the Asset dataset the contract expects. This helps the UI and mapping component understand the structure of the input JSON string expected by the contract. We will see examples of how this the data it returns looks like, shortly.
Since this is a 'query' method, the signature is similar to the others. What we do here is create an empty instance of the AssetState definition, marshal it to a JSON string and return it so that the caller knows the input data structure.
```go
func (t *SimpleChaincode) readAssetObjectModel(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var state AssetState = AssetState{}

	// Marshal and return
	stateJSON, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}
	return stateJSON, nil
}
```
#### Read Asset Sample
The Read Asset Sample returns a 'sample' of the schema to the consumer of the contract by reading the contents of the samples.go file in the project folder.  

#### Read Asset Schema
The Read Asset Schema returns the contract's schema in JSON format to the consumer of the contract by reading the contents of the schemas.go file. schemas.go and samples.go are made available as-is in this contract. In  'trade_lane_contract_hyperledger', we will see how these files can be generated using the 'go generate' command.

### Callbacks
In IBM Blockchain contracts, the Init method handles init commands, the ones called as part of 'deploy' during execution. The Invoke method handles  'invoke' method calls and the 'Query' method handles invocations of type 'query'. We've already implemented 'Init'. Now lets write simple Invoke and Query methods that will call our functions -createAsset, updateAsset, deleteAsset, readAsset, readAssetObjectModel, readAssetSchemas and readAssetSamples methods, based on the funtion name passed to the contract invocation. The following example is for the 'Invoke' implementation
```go
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
#### Validate Input data
This is not an implementation of an IBM Blockchain deploy, invoke, or query method, but a common function called by all the 'invoke' type functions (createAsset, updateAsset and deleteAsset) and the readAsset 'query' function that does the following:
1. Validates the input data for the right number of arguments,
2. Unmarshals the input data to the asset state and
3. Checks for the asset id.
```go
func (t *SimpleChaincode) validateInput(args []string) (stateIn AssetState, err error) {
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
This contract allows partial updates of data. The mergePartialState funtion iterates through the old and rew record. For those fields that have updates, the function replaces the old values with the new value. The merged record is then used to updatepdateAsset function
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
### main
And last but not least, the 'main' function  which creates a shim instance.

Note that we have a fmt.Printf command here. Printf and Println commands can easily be used to debug the chaincode as we write / modify and test chaincode. However, since the chaincode runs inside a docker container, these will get written to the log of the docker instance and space gets filled up very quickly, especially if the environment is running in debug mode. It is prudent to limit or avoid Print statements in actual deployed chaincode as the chaincode is essentially supposed to run in the background.

Best practices around docker log management, like Logrotate (supported since Docker 1.8) could be adopted to avoid space issues. Contract logging implementations could also be incorporated in future contracts.
```
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}
```


## Invoking the contract
{: #invoking_the_contract}

We have created a simple contract which implements a CRUD interface and enables us to manage Watson IoT Platform asset data on the blockchain. Great! Now how do we use it? IBM Blockchain contracts expose REST APIs and these can be invoked through the command-line, or we could make REST API calls using, say, Swagger or Node.js. Later on, we will see an example of a sample UI that provides us with a more user-friendly way of calling the contract.
No matter which method you use, the actual syntax of the calls made to a chaincode will not change. Like we saw earlier, IBM Blockchain understands three types of chaincode calls - 'deploy', 'invoke' and 'query'. Every call expects a JSON string as input. This string will include a Function key and an Args key. The Function key is the name of the actual function you are calling and the Args are the arguments you want to pass on to it.

IBM Blockchain github pages clearly explain how to set up the [DevNet](https://github.com/hyperledger/fabric/blob/master/docs/dev-setup/devnet-setup.md) and [Sandbox](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md) environments and how to test contracts in the command-line. The Sandbox environment is particularly useful when writing new contracts or modifying existing ones - it provides us with an environment to test and debug chaincode without having to set up a complete IBM Bluemix Hyperledger network. Let's see some examples of the command-line and rest calls to our sample chaincode.
Once the network is running, contract is built and instance registered, (all covered in the links above) we can start calling our contract.If we are using a Swagger or other REST interface, the HTTP server and the Swagger  / Node.js setup needs to be in place as described [here](https://github.com/hyperledger/fabric/blob/master/docs/API/CoreAPI.md). 	

### Contract calls
**Note:** In Sandbox, we get to set the contract instance name. In a real peer network, when we register the contract, it will have a lengthy alphanumeric name that is returned by the network and that is the name that needs to be used. Alternately we can use the github path to point to the contract.

### Deploy
This is an example call to 'deploy'. The contract name in the Sandbox is ex01. Notice the signature. We are calling the peer executable with the chaincode deploy command. We specify the name of the contract as ex01 and the arguments, in a JSON format are the Function name 'init' and the arguments it expects inside Args: []. IBM blockchain allows the content of Args[] to be any string(s), however, since our contract is about IOT data, which comes in as json, we have written our contract to expect a json string inside Args[].
In command line Sandbox mode, in the tab where you have registered the contract, you can see the response of the Shim, stating that init was received and was successful. If you have print statements in the contract, those can also be seen here. They are useful for debugging, but are best removed once the contract is deployed. IBM Blockchain 'deploy' and 'invoke's are asynchronous.
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
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"createAsset", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":34, \"longitude\":23}, \"temperature\":89, \"carrier\":\"ABCCARRIER\"}"]}'
```
### readAsset
```
 ./peer chaincode query -l golang -n n ex01 -c '{"Function":"readAsset", "Args":["{\"assetID\":\"CON123\"}"]}'
```
 This returns
 ```
 {"assetID":"CON123","location":{"latitude":34,"longitude":23},"temperature":89,"carrier":"ABCCARRIER"}
 ```
### updateAsset
```
./peer chaincode invoke -l golang -n  ex01  -c '{"Function":"updateAsset", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":56, \"longitude\":23}, \"temperature\":78, \"carrier\":\"PQRCARRIER\"}"]}'
```
another readAsset with return the updated asset data

### deleteAsset
```
./peer chaincode invoke -l golang -n  ex01 -c '{"Function":"deleteAsset", "Args":["{\"assetID\":\"CON123\"}"]}
```

## Using the Simple Contract with the Watson IoT Platform
{: #using_contract_iotp}

 So far we have looked at how to write a simple contract in IBM Blockchain for use in a Tradelane scenario, where device / asset data comes in and gets captured on the blockchain. It implements a CRUD pattern, one familiar to most users where data storage and management is involved. There could be other patterns, say, for example getter-setters. We have also seen examples of how to deploy and test the contract. However, how do we use this contract in conjunction with the Watson IoT Platform? For a beta customer, what are the bare minimum steps required to ensure that the contract and Watson IoT Platform work together?

### Data Mapping
The Watson IoT Platform Data Mapping component is what causes the Watson IoT Platform data to be routed to the Blockchain contract, thereby enabling them to work together. This is where incoming IoT Event properties are mapped to corresponding contract properties.

#### Contract features required for Data Mapping
1. The `updateAsset` function.  
This simple contract is a recipe, an example intended to be tweaked as people experiment with contracts. Today's simple contract, though it implements CRUD, is very accomodating - it has the same features in create and update functions. The Data mapping component expects an updateAsset function that can do the following:
 * create the asset record if it doesn't exist on the ledger.
 * update asset data if the asset already exists.
 * asset id the key. The asset id maps to the serial id the comes in from the Watson IoT Platform.
 * Accepts input as a JSON string. This makes sense as IoT device data generally comes in as JSON strings.

2. The `readAssetSchemas` function.
The readAssetSchemas function exposes the function names and properties expected by the contract. This function is called by the mapper in order to know what the contract properties are thus enabling mapping of the Event properties to the contract properties.

## Next steps

When your contract is created, you can upload the contract to IBM Blockchain with Watson IoT Platform, or by using the API. For more information, see [Developing for blockchain](https://TBC)

IBM Blockchain also provides us with a sandbox environment to test contract code before we deploy it to the peer network. For more information about how to set up the sandbox, see [Writing, building, and running chaincode in a development environment](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md).
