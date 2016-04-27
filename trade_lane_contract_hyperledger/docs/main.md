# IoT Sample Contract   
[`main.go`](../main.go "main file for the IoT sample contract")  

## Brief Introduction to Hyperledger, IoT Sample Contract, and Watson IoT Platform
The reader have some background in 
[Hyperledger fabric](https://github.com/hyperledger/fabric) concepts before reading this document and its siblings. Pre-reading suggestions include:

- the contents of the subfolders under the Htperledger [`/docs`](https://github.com/hyperledger/fabric/tree/master/docs) folder, with particular attention paid to chaincode topics
- a comprehensive explanation of Hyperledger concepts and implementation in the [Protocol Specification](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md). 
- the [Effective Go](https://golang.org/doc/effective_go.html) online ebook

###Contract
The sample smart contract is written 
in the `Go` language. It runs in a _**[Hyperledger](https://github.com/hyperledger "open source blockchain fabric incubating in the Linux Foundation")** blockchain fabric_ and is designed for integration with the _**[Watson IoT Platform](http://www.ibm.com/internet-of-things/) (IoTP)**_. 

Smart contracts implement business logic specifically oriented towards a blockchain's distributed, transparent ledger. Smart contracts service two classes of message:
- transactions, which change ledger state, and
- queries, which read ledger state

Transactions can include, but are not limited to:
- deploy a *contract instance* into the fabric
- upgrade a *contract instance* in the fabric with a new version (future feature)
- create an asset
- modify an asset's state
- delete an asset

####Contract Instance
The term *contract instance* defines one deployment of a specific body of contract code. For example, say that you have a contract code that you think of as version 1. You can deploy that contract several times for several purposes, assuming that it is designed for multiple deployments. A contract in its physical form is referred to as *chaincode*, and these two terms are often used interchangeably.

A chaincode's ID as returned as the synchronous result of a deployment transaction defines the contract instance's target ID for messages. The UUID is calculated over the source code, the path from which it is deployed, the constructor function name and the initialization parameters. See [Protocol Specification - 3.1.2.1 Transaction Data Structure](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#3121-transaction-data-structure)

Changing the constructor function or arguments will change the identity of the deployed contract instance, each instance being separately addressable by applications. An instance has a separate partition in world state but shares the same blockchain with other instances. When upgrading a chaincode to a new version, each contract instance is upgraded separately.

####Behavior and State
A contract chaincode encapsulates behavior and state. This sample's behavior and state are not specific to any given scenario or domain, but the architecture and design flexibly support IoT device events, for example as sent by the IoTP. In other words, *behavior* is event-oriented and *state* is asset-oriented.

This sample's event and state data model has a few *trade lane* customizations to simulate tracking of assets such as packages and containers as they move between locations in the hands of one or more carriers. 

Extending this sample contract to handle more complex scenarios or other domains is a straightforward exercise as discussed in [`CustomizingTheSampleContract.md`](./CustomizingTheSampleContract.md). 

###Hyperledger Fabric
![Diagram 1: Hyperledger Fabric](images/hyperledger_fabric.jpg)
Diagram 1: Hyperledger Fabric

In diagram 1, the fabric is shown as a set of interconnected validating peers (VPs). These are tasked with validating transactions and negotiating consensus. They execute transactions by delegating message payloads to the contract. The more sophisticated consensus algorithms can vote on both inputs and outputs, but with the Bluemix demo algorithm consensus is automatic.  

>*Note: Any change to the ledger world state requires transaction __consensus__.*

Bluemix demo environments follow [Protocol Specification - 2.2.1 Single Validating Peer](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#221-single-validating-peer) and use only VPs. 

The diagram also shows optional non-validating peers in the role of application gateways. NVPs are responsible for verifying inputs and forwarding transactions to connected VPs. NVP's do not process messages locally as at this time they do not have a copy of either the blockchain or the world state. NVP architecture is evolving and this may change in the future. 

Applications can communicate with the fabric through an NVP or a VP.

See [Protocol Specification - 2.2.2 Multiple Validating Peers](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#222-multiple-validating-peers) for more information.

###Hyperledger Contract Connection and Flow
A contract instance in the Hyperledger fabric communicates directly with its associated peer through a `shim`. A shim is similar to a *service provider interface* in other architectures, allowing many contracts with varied logic and purpose to communicate with a peer.

![Diagram 2: Peer to contract connection](images/peer_contract_connection.jpg)  
Diagram 2: Peer to contract connection

As shown in diagram 2, shim logic is split between the peer (proxy side) and the contract (stub side), with the stub side defining a contract API for routing contract payload messages. The stub is passed into the contract as the parameter `stub`, which enables the contract to communicate with its peer, the world state database, and of course applications and devices. 

See:
- proxy side implementation in [`chaincode_support.go`](https://github.com/hyperledger/fabric/blob/master/core/chaincode/chaincode_support.go)
- stub side implementation in [`shim`](https://github.com/hyperledger/fabric/tree/master/core/chaincode/shim) folder
- [Protocol-Specification 3.3 Chaincode and 3.3.1 Virtual Machine Instantiation](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#33-chaincode)

### Transaction Walkthrough
A transaction starts with an *invoke* or *deploy* message and runs in the ledger's *write mode*,seeing and modify both uncommitted state and committed state. Invokes can thus build upon one another before a block is committed with all of its transactions, after consensus is achieved of course.

The process looks like this:
- an application or message broker (e.g. Watson IoT Platform) sends a transaction message
- a non-validating peer (NVP) or validating peer (VP) receives the message
- the peer (NVP or VP) that received the message
    - verifies the inputs and creates a transaction if inputs verify successfully
    - the peer broadcasts the transaction to its connected VPs to initiate consensus and execute the specified contract function
- each VP then 
    - executes the transaction by delegating it to the specified function in the contract
    - this generally constitutes validation of the transaction
- the peer that received the initial message then
    - receives or generates the response
    - forwards the response to the calling application
- if a majority of VPs vote __*yes*__ for consensus
    - **all** VPs commit the transaction to the blockchain and the resulting state to world state
        - those who voted *no* are overridden and forced to go along with consensus
- else (a majority of peers vote __*no*__ for consensus)
    - the transaction is rolled back along with transactions built upon its changed state, including chaincode to chaincode invokes 

> Note that the default consensus plugin is called NOOPs (NOOP == no operation) that defaults consensus to "yes" in all cases.

### Query Walkthrough
A *Query* runs in *read only mode* and can see only committed world state with no ability to change it. Thus, queries show state only after the whole process of validation, consensus and execution has completed. This means that, if the query can see the state change, it is permanent and a hash of the delta has been recorded in the blockchain next to the transaction that created the new state.

The process: 
- an application sends a query message
- an NVP or a VP receives the message and executes the query
  - the peer passes the message payload (function name and arguments) through the shim to the contract
  - the query executes by reading committed state from world state
  - the contract returns the result to the peer
- the peer forwards the result to the calling application synchronously

>Note that transactions execute asynchronously, returning the transaction's UUID synchronously and a later event that documents the success or failure of the transaction. Transactions can see both committed and uncommitted state. Queries execute synchronously and can see only committed state. Thus, queries will not see the result of a transaction immediately after the UUID is received, or even for several seconds or longer. It is of **critical** importance to design applications with a tolerance for the inherent hysteresis.

###Messaging Between Application and Contract
See [Protocol-Specification - 6.2.1.3 Chaincode API](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#6213-chaincode-api)

See the fabric's [RESTful API](https://github.com/hyperledger/fabric/blob/master/core/rest/rest_api.json) in the 
[Hyperledger project at github.com](https://github.com/hyperledger/fabric). 

Applications send messages to a peer, targeting a RESTful API endpoint. There are endpoints for blocks, transactions, chaincode, and so on. The RESTful API target that is relevant to smart contract interaction is at the endpoint `/chaincode`. Messages targeted at `/chaincode` contain a contract API payload with a specified function and array of string arguments. 

A contract *invoke* message with function `updateAsset` and a JSON encoded object argument in `args[0]` is illustrated in the following diagram:

![Diagram 3: Message protocol between applications and peers](images/message_payloads.jpg) 
Diagram 3: Message protocol between applications and peers

In diagram 3, four message layers are shown:

1. `HTTP POST` message addressed to the NVP's (or less often a VP's or potentially a cluster's) internet address with headers as shown in the diagram. The target path for all messages is `/chaincode`
2. POST body carrying a [JSON RPC](http://json-rpc.org/ "http://json-rpc.org/") compatible envelope that specifies:
  - `method` : selects the operational mode from`deploy`, `invoke`, and `query`
  - `id` : coordinates message and response
  - `params` : an inner envelope 
3. `params` envelope to define and address the chaincode
  - `chaincodeID` : the targeted contract instance identifier as returned from a deploy message
  - `type` : 1 == golang
  - `secureContext` : credentials for a secure peer, ignored by peers running in unsecured mode
  - `ctorMsg` : contract payload
4. `ctorMsg` : contract API invocation with function selector and arguments
    - `function` : the name of the function in the contract that is to process the event
    - `args` : an array of strings as arguments to the function, in this contract sample there is one argument that is a JSON encoded object 

The rest of this document is concerned with the contract's API as implemented in the `main.go` file.

---

##Contract API and Operation
User chaincode runs in a virtual environment such as a Docker container and communicates with its associated peer via the shim messaging protocol as shown in Diagram 2. The contract executable is built after extracting it from a github repository either locally as a github clone inside the peer build, or remotely at [github](http://github.com). 

This sample's physical structure is flat, using one `main.go` Go file and several related Go files in the top level folder. Together these embody `package main`. The main file implements the shim interface, the API delegation logic, and the contract's state and business logic.

Related go files add features that include (but are not limited to) asset history, contract state, recent states, alerts, and rules, also as a part of `package main`. See [README.md](./README.md). 

This file is in the `/docs` subfolder, along with other documents describing the related go files. The `/scripts` subfolder contains a Go script that is executed when the `go generate` command is typed at the command line while in the main folder context. The execution is controlled by the following comment line near the top of `main.go`:

``` go
//go:generate go run scripts/generate_go_schema.go
```

###Contract Payload and Delegation
The *JSON RPC* message envelope targets chaincode modes, which translates to the following functions and arguments in the shim's API:

- `Init` : called for `deploy` mode, initializes the contract, may delegate by switching on `function` but this sample does not
- `Invoke` : called for `invoke` mode, delegates by switching on `function`
- `Query` : called for `query` mode, delegates by switching on `function`

These must be implemented in the contract to the same specific interface:

``` go
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
```
A SimpleChaincode struct is used as the receiver for shim functions and their delegates. The shim itself is passed into the parameter `stub`.

The contract payload is carried in the next two parameters: 

- `function` : a string that specifies an internal function in the contract API
- `args` : an array of strings that contain the arguments to the *function*

The following example shows the basic mechanism for `invoke` mode to delegate to `createAsset`:

``` go
type SimpleChaincode struct {}

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "createAsset" {
		return t.createAsset(stub, args)
    } else ...
    return nil, fmt.Errorf("unknown function: %s", function)
}

// ... and later in the file ...

// contract API delegated by Invoke
func (t *SimpleChaincode) createAsset(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
    code body
    if err = bad thing happens
        return nil, err
    code body
    return nil, nil
}
```
`Invoke` and `Query` functions delegate internally for all but the most trivial of APIs. In deploy mode, a simple contract may implement initialization code directly in the `Init` delegator function, while more complex contracts might add, say, an `upgrade` function.

###Ledger State
Transactions generally exist to change ledger state as the result of an incoming event. The state might pertain to the contract or all assets (e.g. recent state changes), or to an individual asset. Ledger state manipulation is performed with the stub API. 

Ledger state is stored as key:value (K:V) pairs with string keys and []byte values. The byte stream is often a cast of a string value. Hyperledger examples show string names as keys and converted integers (using iToA) as values while this contract chooses to store state as a JSON encoded object. 

The shim supports direct communication with other smart contracts, which is known as *chaincode to chaincode* invokes. This is useful for complex systems with multiple interacting smart contracts, each with a different asset focus. Examples of assets that could interact on separate subchains are devices, *things* (a.k.a. assets), legal contracts, warranties, financial transactions, and so on. The separation can also be implemented for security or business rule segmentation.  

> When an invoke transaction fails at some point in its execution, the entire transaction and all state changes in all participating contracts are rolled back. This is equivalent to a database rollback in traditional applications.

Code snippets for the ledger state APIs used in this sample contract:

- `stub.PutState` : writes or overwrites the state at a specific key
  ``` go
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
  ```
  The JSON marshal function takes the address of the map that stores the output state and generates a string representing the object. Arguments to the `PutState` function are the assetID string as key and the stateJSON string cast to `[]byte` as value.

- `stub.GetState` : reads an existing state at a specific key, returning a JSON encoded object string as `[]byte`
``` go
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
 ```
 The byte array is then unmarshalled into a generic variable of type interface{}, which creates a nested structure of `map[string]interface{}` at each level in the object. In Go, the maps can be directly processed by asserting the type to a new variable:
 ``` go
     // assert the existing state as a map
    ledgerMap, found = ledgerBytes.(map[string]interface{})
    if !found {
        log.Errorf("updateAsset assetID %s LEDGER state is not a map shape", assetID)
        return nil, err
    }
```
- `stub.DelState` : removes the state value at a specified key, but does not remove transactions from the blockchain as those are indelible

###State Keys
See [Protocol-Specification - 3.2 Ledger](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#32-ledger)

It is important to define key patterns that can be managed in a straightforward manner without performing indirect key lookups. It is more efficient and clear to use asset IDs directly where possible. This contract stores values as JSON encoded object strings. 

This sample contract uses the following keys and patterns to store state:

- `assetID`: state for one asset
- `assetID + ".StateHistory"`: array of asset states in order of arrival, most recent first
- `ContractStateKey` : contract state including version, nickname, and array of managed assetIDs
- `RecentStatesKey` : array of state changes in order by time, most recent first, assets appear only once and jump to the front when they get a new state change

##Contract Patterns 
###CRUD API Pattern
The contract payload definition (i.e. Contract API) is generic, so there are no restrictions on the names that can be used for the functions and arguments. On the other hand, following naming patterns offers advantages where automation is concerned, as demonstrated by the `monitoring_ui` that configures itself based on the CRUD pattern this sample contract uses for API function names.

The CRUD pattern is a familiar way to model database access and while it is not without its issues, it offers a clear and straighforward resource model for state manipulation and avoids API proliferation that is common as devices get smarter with larger subsets of state data sent in each device event.

The contract's primary key is `assetID`. As mentioned in earlier sections, an asset can be from a shipping container bar code to a package tracking number, a vehicle identifier (VIN) or anything else. What matters is that the asset is the focus of the contract. This contract sample implements a light weight trade lane scenario and so assets are packages or containers.

The state data is polymorphic, in that assets do not have to be of the same type. They just need to respond to a common partial subset of the writable state properties, for example location and temperature. 

It would be a mistake, though, to combine several disparate API subsets by asset classification. In that situation, collaboration contracts properly isolates state and behavior while allowing asset classes to post status or event to other asset classes managed by a different contract. This avoids rampant *specialization by guard code*.

CRUD standards for `CREATE`, `READ`, `UPDATE`, and `DELETE` operations. This contract models all APIs in one of those four classifications and used their lowercase equivalents as prefixes for all function names. This again offers automation in the form on the generic UIs in this project.

|CREATE|READ|UPDATE|DELETE|SET|
|---|---|---|---|---|
|createAsset|readAsset|updateAsset|deleteAsset|setLoggingLevel|
||readAllAssets||deleteAllAssets|setCreateOnUpdate|
||readAssetHistory||deletePropertiesFromAsset||
||readRecentStates||||
||readAssetSamples||||
||readAssetSchemas||||
||readContractState||||

>The appearance of two functions starting with `set` is meant to separate them from asset and contract state. They are actually configuration parameters for the control flow of the contract execution.

In addition to the generic UIs in this project, the consistency of this pattern is also leveraged in the mapping features in the IoT Platform. 

###Partial State as Event Pattern
The term state is generally used as a short form for *World State*, which is the total of all K:V pairs in the ledger. However, state is also used when referring to the current values of all of an asset's properties, which are stored in a single object called `state`. The writable properties are grouped into a subset object called `event`, which models the incoming events sent by devices and applications. 

This contract is focused on assets and so are the `event` and `state` objects:
- `event` : defines all of the writable properties in the state.
- `state` : defines all asset properties, as in the `event` properties plus read-only properties such as `alerts` and `inCompliance`

####Event
Having previously established that the function names in this contract API rigidly follow the CRUD pettern, the Partial State as Event pattern defines how arguments are shaped when passing through the `args` string array. This contract sample chooses to use only `Args[0]` and to expect that to be a valid JSON encoded `event`, as defined in the [contract payoad schema]((../payloadSchema.json). 

This pattern's utility lies in using *partial states* as events rather than defining a large number of rigid event objects. With JSON's flexible tagging property tags, JSON objects need not enforce completeness. It is perfectly valid to be missing all but the properties tagged as `required`. This contract's `event` object has only the `assetID` defined as required.

One major advantage of this approach to IoT device events processing is that the contract need not be respun when a new composite device comes online. Previously unseen combinations of writable properties are always valid. For example, a new device that can send a location, a temperature and a gForce in obe event does not affect the contract's processing at all. Yet were this a rigidly defined API, a new function would be required and the contract would need to be upgraded across the entire fabric. 

####State
Asset state is built up as events are received. For example, an event with only the `assetID` could create the state, and all that would be stored is the asset's ID. Another event could arrive with a location, and then another with the temperature. As each event is processed, a new state is calculated and written into the ledger.

Asset is therefore an object that is the sum of all events and calculations. It is encoded into a JSON string and stored into the ledger with `assetID` as key. Incoming events are deep merged into the state ad infinitem, each creating the next asset state in the ledger. 

####Ad Hoc Extension
The partial state as event pattern helps with object model extension in a running contract. Because JSON is tagged and flexible, a property has been introduced into `event` and `state` that is meant to offer extra flexibility for applications to store data alongside the pre-defined properties. 

This property is called `extension` and can be treated as a nested object to which application-managed state is added as necessary. This enables an application to add "sidecar" data to asset state without upgrading a contract. The deep merge algorithm handles extension data just as it handles schema defined properties. 

Treat this feature with care though, as the contract knows nothing of this data and thus cannot apply rules to it. But it can save a lot of time and cost when related data can be tracked alongside and is thus available to all monitors and applications. 

###Maps over Structs
The partial state as event pattern requires the use of another pattern, that being the use of maps over structs for internal object management. 

In Go, JSON will populate a struct from an incoming event using reflection. However, properties that are not present in the struct are lost, and all properties not present in the event are set to their zero value, which cannot be used to determine presence or absence of the property since `0` is a valid value with some fields such as temperature.

It is also difficult to generically merge an incoming event into a state struct since again there is no way to know which properties were sent. Structs necessitate a complex and rigid API based on getters and setters to be effective for event processing and state management. And in such a case all flexibility to adapt to new devices is lost.

Maps, on the other hand, reflect exactly what was present in the event or state and offer a very flexible way to deal with incoming data. It is even possible to validate incoming events against the contract payload schema, although this sample does not yet implement that check. 

Deep merging of sparsely populated events nested as deeply as necessary to render the asset object model into the ledger state avoids all of the primary issues with structs. Note that, after deep merging the current state is fed into the rule engine to calculate alerts and compliance, so the merging in of any subset of properties is seen by all rules and thus lowers the likelihood of bugs in the common processing as the flows are all generic.

##Contract Object Model
###Asset
A smart contract derived from this sample must have a primary area of focus. It is the `thing` on which the contract is focused. The thing could be a container or package, as in this sample contract, or it could be a vehicle, a refrigerator, an automobile part, and so on. 

This contract generically calls the *things* about which it cares **assets** and is therefore built around the identifier for those assets, the `assetID`.

> A reminder that this sample contract is written for IoT scenarios, where assets are identified and are the subject of periodic readings from **devices**, including bar code scanners, GPS locators, accelerometers, and so on. Composite devices may send readings containing several such properties. This was covered earlier in the *Partial State as Event* pattern discussion. 

###Device
This contract does not have device identifiers in it, but it could. For example, every event is generated by a device or an application (which can and should be thought of as a *logical device* or a *device gateway* if it is sending readings or composite readings). A device will have a unique identifier if its message is sent through the IoT Platform. 

The key point about devices is that they send data about *things*, which we generically call *assets*. Thus, devices play a key role in the lifecycle of assets. 

> *The deviceID from which an event was generated would be saved automatically inside the `lastEvent` portion of the asset state so that lightweight analytics could be performed on devices using asset history.*

###Event
As mentioned earlier, the `event` object consists of the writable asset properties and its JSON encoding supports the partial state as event pattern. The following properties define an event:

####`event` - Schema (output from getAssetSamples)
``` json
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "extension": {},
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "temperature": 123.456,
        "timestamp": "2016-04-15T10:53:30.3964036-04:00"
    },
```
Properties used by contract sample:

* Mandatory for the contract to function at all
    * `assetID` - the ID of the asset that is the central focus of the contract, a good example being a container or package
* Useful in all IoT contracts 
    * `location` - a `Geolocation` object containing the properties `latitude` and `longitude` specifying the last known location of the object
    * `timestamp` - date and time that the event occured or state was created. 
    * `extension` - application-managed state as discussed elsewhere
* Specific to this sample contract's lightweight trade lane scenario
    * `temperature` - an asset-specific measurement of the temperature in Celsuis. This example contract uses zero as the threshold value, inclusive. That is, any temperature above zero is too high. The example is meant to simulate shipment of frozen goods. This is used to demonstrate how alerts work, so a temperature above zero raises an `OVERTEMP` alert, which would cause inspection of the contents on arrival and possible penalties to the responsible carrier.
    * `carrier` - the entity currently transporting the asset. Generally, a shipping company that is moving the asset between any combination of origin, way points and destination.

###State
The `event` is a subset of the asset's `state`. Thus state becomes event plus calculated (as in read-only) properties. These include `alerts`, `incompliance`, and `lastEvent`.

####`State` Schema (output from getAssetSamples)
The state object looks like:
``` json
    "state": {
        "alerts": {
            "active": [
                "OVERTTEMP"
            ],
            "cleared": [
                "OVERTTEMP"
            ],
            "raised": [
                "OVERTTEMP"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "extension": {},
        "inCompliance": true,
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "temperature": 123.456,
        "timestamp": "2016-04-26T23:51:16.8214639-04:00"
    }
```

Read-only state properties additional to previously documented event properties:

* `alerts` - three arrays of alert name string from the `alertName` enum in the schema, of which only `OVERTEMP` is implemented in this sample
    * `alerts.active` - alert names that are currently active
    * `alerts.raised` - alert names that were raised by this event, defining the exact moment of transition from clear to active
    * `alerts.cleared` - alert namess that were cleared by this event, defining the exact moment of transition from active to clear
* `incompliance` - a boolean property specifying whether there are any active alerts for this asset
   * the sample contract defines compliant as no alerts active, but derived contracts can have much more complex algorithms
* `lastEvent` - the event that created this state
   * `function` - the function that handled the incoming event
   * `args` - the string args composing the incoming event, always one in this contract pattern
   * `redirectedFromFunction` - the function that redirected the event to the function that created this event

##Contract API
The contract's CRUD API that generically handles creation and update of assets without getting drawn into the minutia of many getters and setters. These generic flows allow the processing to separate into clearly identifiable chunks, mitigating complexity. That said, there are subtle differences between create and update that should be studied carefully before making changes. 

The documentation for the most complex functions will use annotated pseudo-code to separate the intent from the noise of the language and error handling. Consider following this document and the code in side by side windows in your favourite editor. Examples of excellent editors include VSCode, Sublime, Brackets and many others.

###Invokes (State Changing)
####createAsset
Creates an asset that does not already exist, using `assetID` as the key. Attempting to create an existing asset will generate an error and fail. 

>*Note that there is one feature that can send a second parameter to `createAsset`. Redirected invokes to update non-existent assets are redirected from the `updateAsset` function by default. The name of the redirecting function is stuffed into `Args[1]` and saved in the `lastEvent` calculated property in the state.*

Following is an annotated pseudo-code description of `createAsset`:

``` go
// 1 arg, arg[0] is JSON partial state as event, optional arg[1] is redirecting function
fail if args < 1 or > 2
// next two lines create map that supports partial state as event
// note, since asset does not exist argsMap is also initial state
unmarshal args[0] as var event of type interface{}
assert event as type map[string]interface{} into var argsMap of type ArgsMap 
fail if assetID does not exist in args
fail if assetID is already in ledger
if timestamp does not exist
    insert one into argsMap using transaction time
// alerts removed from state if nothing is active, raised or cleared
// incompliance removed if false 
run rules engine against argsMap to generate alerts
if any alert is active, raised or cleared
    add alerts to argsMap
    delete incompliance from argsMap
else 
    delete alerts from argsMap
    add incompliance=true to argsMap
copy argsMap to stateOut
// lastEvent is the event that specifically created this state
create lastEvent with function, event, and redirect if present
marshal stateOut to JSON
PUT state to ledger
add asset to activeAssets in contract state
create history bucket for asset and add state
push state change into recent states
```

####updateAsset
Update assets that already exist. An assetID for a nonexistant asset will redirect to `createAsset` by default, or generate an error if the redirect feature is disabled using the `setCreateOnUpdate` API. 

The primary difference from a `createAsset` invocation is that `updateAsset` gets the existing asset state from the ledger and merges the incoming event with it to generate the new state. This is the cornerstone of the *partial state as event* pattern, since events carry data that are a subset of the writable properties of the state. Calculated fields are created after the deep merge has completed so that the entire state can be pushed into the rules engine and any other post processing that may exist in a derivation of this sample.

Following is an annotated pseudo-code description of `updateAsset`:

``` go
// 1 arg, arg[0] is JSON partial state as event 
fail if number of args != 1
unmarshal args[0] as var event of type interface{}
assert event as type map[string]interface{} into var argsMap of type ArgsMap 
fail if assetID does not exist in args or is blank
if assetID is not in ledger
    redirect to createAsset if enabled
    else fail
if timestamp does not exist
    insert one into argsMap using transaction time
//args ready, process existing state from ledger    
GET asset state from ledger as assetBytes of type []byte
unmarshal assetBytes as ledgerBytes of type interface{} // bad name
assert ledgerBytes as type map[string]interface{} into var ledgerMap of type ArgsMap 
// deepMerge is a function in mapUtils.go
deep merge argsMap into ledgerMap and assign to stateOut
// alerts removed from state if nothing is active, raised or cleared
// incompliance removed if false 
run rules engine against argsMap to generate alerts
if any alert is active, raised or cleared
    add alerts to argsMap
    delete incompliance from argsMap
else 
    delete alerts from argsMap
    add incompliance=true to argsMap
// lastEvent is the event that specifically created this state
create lastEvent with function, event, and redirect if present
marshal stateOut to JSON
PUT state to ledger
add state to history bucket for asset
push state change into recent states
```

####deleteAsset
Delete an asset that already exists. Deleting a nonexistent returns an error. Asset state and history are permanently removed from the database. Asset is removed from recent states and from activeAssets in contract state. It ceases to exist except as transactions on the blockchain.

This function makes irreversible changes to World State. It should not be used in situations where historical data is used unless that data is backed up outside the blockchain. An alternative to deletion would be an asset status property that designates the asset as `idle` or `retired`. 

For example, a container sitting in storage after refurbishment might be allocated a new bar code, in which case the previous assetID could be deleted (permanent wipe) if the info is not needed for audit purposes or retired (logical deletion retaining history). 

``` go
// 1 arg, arg[0] is JSON partial state as event 
fail if number of args != 1
unmarshal args[0] as var event of type interface{}
assert event as type map[string]interface{} into var argsMap of type ArgsMap 
fail if assetID does not exist in args or is blank
fail if assetID is not in ledger
DELETE state from ledger
remove history bucket for asset
remove entry for asset if any from recent states
```
####deletePropertiesFromAsset
>Unlike `deleteAsset`, this function does not change history, other than to add the new state to history. Removing properties is a semantic operation, and only to be performed when the properties are semantically incorrect with the asset's current status.

Removes specified properties from an asset. Processing in general is similar to updateAsset in that asset state (i.e. the ledger) is modified. But rather than deep merging during update, this function performs deep deletion.

Removing a property from a map is quite different from setting it to `zero` or `false` or `empty` in a struct. The absence of a property is semantically useful in that it indicates that the contract should no longer be concerned with it. Rules are designed to ignore missing properties, unless that is specifically what the rule is about. 

The `OVERTEMP` rule, for example, does not perform a calculation if the temperature is not present. And since the rules are run after state update, deleting temperature from the state has the same effect as receiving a compliant reading - any active alerts will clear. 

> Why might one remove the temperature property? 

> A container may have frozen goods removed at a way point before proceeding along its journey, causing the supervising application to realize that the container no longer contains frozen goods. There is therefore no longer a need to track the temperature, which effectively disables the `OVERTEMP` rule. Of course, if temperature readings are expected to continue with refrigeration is turned off, then new API would be needed to explicitly disable the rule or change the `OVERTEMP` threshold.

> Additionally, a more sophisticated contract might have a status property, where the transition from status delivered to status idle would remove all transient properties like temperature as they may or may not be relevant to the next shipment. The removal should only be performed after status has transitioned to `IDLE` because the compliance status is important in status `DELIVERED` in that inspections must have the opportunity to be performed if the contract has ever been out of compliance during the shipment.

``` go
// 1 arg, arg[0] is JSON partial state as event 
fail if number of args != 1
unmarshal args[0] as var event of type interface{}
assert event as type map[string]interface{} into var argsMap of type ArgsMap 
fail if assetID does not exist in args or is blank
fail if assetID is not in ledger
fail if qualified property names array is missing
    
//args are ready, now process existing state from ledger    
GET asset state from ledger as assetBytes of type []byte
unmarshal assetBytes as ledgerBytes of type interface{} // bad name
assert ledgerBytes as type map[string]interface{} into var ledgerMap of type ArgsMap 
// removal means splitting the qualified properties by level
// and removing the property from its submap. 
for all qualified properties to delete
    continue on next if deleting assetID or timeStamp 
    split qualified property by '.' into levels of type []string
    for all levels
        if this is the actual property (last level, leaf node)
            delete if propertyname exists
        else
            set context to new level if it exists

if timestamp does not exist
    insert one into argsMap using transaction time
// alerts removed from state if nothing is active, raised or cleared
// incompliance removed if false 
run rules engine against argsMap to generate alerts
if any alert is active, raised or cleared
    add alerts to argsMap
    delete incompliance from argsMap
else 
    delete alerts from argsMap
    add incompliance=true to argsMap
// lastEvent is the event that specifically created this state
create lastEvent with function, event, and redirect if present
marshal stateOut to JSON
PUT state to ledger
add state to history bucket for asset
push state change into recent states
```
####deleteAllAssets
Perform deleteAsset on every asset in the contract state's `activeAssets` array. The end result is that the ledger's state for this contract is empty.

###Queries (These do not Change State)
####readAsset
Remember that state is stored as an encoded JSON object, i.e. a string. So if the assetID is present and non-blank and if the asset exists, then get the state from the ledger and return the string as is, cast to `[]byte` of course.

####readAllAssets
Perform readAsset on every asset in the contract state's `activeAssets` array after sorting the `assetID` list. This makes it easy to search the list at the other end. Add each asset state to an array after unmarshaling. Once the array is completed, marshal the entire array and return to the client.

>*Warning: This API may not scale in applications with many active assets. Be aware of the costs and use at your own risk.*

####readAssetHistory
Events received on behalf of an asset generate new asset states. Each new state is prepended to the asset's state history array, which is stored in reverse order with the most recent first. 

This function is intended to query the asset's history for purposes such as lightweight analytics, an example being plotting the temperature by date, time and location in order to assess penalties for spoilt goods.

If the assetID is present and non-blank and if the asset exists, then get the history state from the ledger and return the string as is cast to `[]byte`. Asset history uses a composite key made up of the `assetID` + `STATEHISTORYKEY`, which is defined as:

``` go
const STATEHISTORYKEY string = ".StateHistory"
```
Composite keys such as this are useful for storing additional data *about* assets.

> *Note that a future update to the state history module will introduce partitioning of the history data for better scaling. The partitioning will follow a key naming scheme that allows range iterators to select the appropriate buckets to process for any given timestamp range.*

####readRecentStates
Every time an asset's state is changed, its new state is added to the contract's recent states array. It is placed at the beginning of the array so that the array is maintained in order by event arrival time, most recent first. This should be thought of as the *most recently created asset states across the entire contract*.

> *If an asset is already present in the recent states array, it is removed and the entries before its original position are shifted to close the gap. Its new state is then inserted into the array at the beginning as always. If it did not already exist, then the last one is dropped as the entire array is shifted to open a space for the new event.*

The returned recent states array is currently limited to 20 states. 

####readAssetSamples
The schema processing script generates two Go files that can be read from the contract by applications through this and the following APIs. Samples are full JSON structures with every property populated with a value of the right type. These can be used to gain familiarity or to populate templates etc.
 
Returns the generated string `samples` that contains the samples that were configured to be generated when the contract was built. See the JSON configuration file in the `/scripts` folder.

####readAssetSchemas
The payload schema contains the complete API and object model. The JSON configuration for the script generator creates a Go files with a string variable called `schemas` that can be parsed by applications such as the Watson IoT Platform and the generic UIs in this samples project. 

> **All IoT contracts _must_ implement `readAssetSchemas` so that these applications can automatically integrate.**

Returns the generated string `schemas` that contains the schemas that were configured to be generated when the contract was built.

####readContractState
Returns an object containing the contract's version number, its nickname, and the array of managed assets.

###Contract-Specific API
In addition to the asset's CRUD API, this sample contract has two functions to manage the contract's behavior. They use the prefix `set` instead of a CRUD prefix so that they stand apart from asset-oriented API.

####setLoggingLevel
Change the logging level of a specific peer. The values are case-independent strings and legal values are `critical`, `error`, `warning`, `notice`, `info` and `debug`, matching exactly the levels in the go-logging package. 

The logging level changes immediately and stays at the set level until it is changed or the peer is restarted, in which case it comes up again at the default level. I.e. the new logging level is not persisted in World State.

> This setter is implemented as a transaction and so will affect all peers in a fabric. It will be reimplemented shortly as a query so that only one peer will see the setter. Since it does not attempt to persist in World State, a query is an acceptable mode for the transaction.

####setCreateOnUpdate
Change the contract's redirect permission to allow or deny sending an update of an unknown asset to the `createAsset` function. The default is allow. 

