# Introduction to Hyperledger Smart Contracts for IoT<br/>Best Practices and Patterns

This will introduce you to smart contracts for the Hyperledger Fabric and a few best practices in the form of patterns that the author has found to be useful in the context of the Internet of Things. 

## A Brief Introduction to Smart Contracts for Hyperledger

The reader should have some background in 
[Hyperledger fabric](https://github.com/hyperledger/fabric) concepts before reading this document and its siblings. Pre-reading suggestions include:

- the contents of the subfolders under the Htperledger [`/docs`](https://github.com/hyperledger/fabric/tree/master/docs) folder, with particular attention paid to chaincode topics
- a comprehensive explanation of Hyperledger concepts and implementation in the [Protocol Specification](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md). 
- the [Effective Go](https://golang.org/doc/effective_go.html) online ebook

### Smart Contract

The [IoT Trade Lane Sample for Hyperledger](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_hyperledger) 
smart contract is written in the `Go` language. It runs in a _**[Hyperledger](https://github.com/hyperledger "open source blockchain fabric incubating in the Linux Foundation")** blockchain fabric_ and is designed for integration with the _**[Watson IoT Platform](http://www.ibm.com/internet-of-things/) (IoTP)**_. 

> From this point on, the *IoT Trade Lane Sample for Hyperledger* is implied as having implemented what is discussed herein. When referenced directly, it is simply called *trade lane*.

Smart contracts implement business logic specifically oriented towards a blockchain's distributed, transparent ledger. Smart contracts service two classes of message:

- transactions, which change ledger state, and
- queries, which read ledger state

Transactions can include, but are not limited to:

- deploy a *contract instance* into the fabric
- upgrade a *contract instance* in the fabric with a new version (future feature)
- create an asset
- modify an asset's state
- delete an asset

#### Contract Instance

The term *contract instance* defines one deployment of a specific body of contract code. For example, say that you have a contract code that you think of as version 1. You can deploy that contract several times for several purposes, assuming that it is designed for multiple deployments. A contract in its physical form is referred to as *chaincode*, and these two terms are often used interchangeably.

A chaincode's ID as returned as the synchronous result of a deployment transaction defines the contract instance's target ID for messages. The UUID is calculated over the source code, the path from which it is deployed, the constructor function name and the initialization parameters. See [Protocol Specification - 3.1.2.1 Transaction Data Structure](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#3121-transaction-data-structure)

Changing the constructor function or arguments will change the identity of the deployed contract instance, each instance being separately addressable by applications. An instance has a separate partition in world state but shares the same blockchain with other instances. 

Now say that you have added some features and want to deploy what you consider to be version 1.1 of the smart contract. The key point with respect to upgrading smart contract *instances* to a new version of the *chaincode* is that each contract instance must be upgraded separately in the fabric. This because the fabric sees three completely different entities, despite their sharing a common code base. 

> Note that contracts are difficult to upgrade at this time, and Hyperledger is evolving rapidly.

#### Behavior and State

A contract's chaincode encapsulates behavior and state. Trade lane's behavior and state are not specific to any given scenario or domain, but the architecture and design flexibly support IoT device events such as those sent by the IoTP. Note that *behavior* is event-oriented and *state* is asset-oriented.

Trade lane's event and state data model has a few customizations to simulate tracking of assets such as packages and containers as they move between locations in the hands of one or more carriers. 

Extending trade lane to handle more complex scenarios or changing it for other domains is a straightforward exercise as will be discussed in [`CustomizingTheSampleContract.md`](./CustomizingTheSampleContract.md). 

### Hyperledger Fabric

![Diagram 1: Hyperledger Fabric](images/hyperledger_fabric.jpg)
Diagram 1: Hyperledger Fabric

In diagram 1, the fabric is shown as a set of interconnected validating peers (VPs). These are tasked with validating transactions and negotiating consensus. They execute transactions by delegating message payloads to the contract. The more sophisticated consensus algorithms can vote on both inputs and outputs, but with the Bluemix demo algorithm consensus is automatic.  

>*Note: Any change to the ledger world state requires transaction __consensus__.*

Bluemix demo environments follow [Protocol Specification - 2.2.1 Single Validating Peer](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#221-single-validating-peer) and use only VPs. 

The diagram also shows optional non-validating peers in the role of application gateways. NVPs are responsible for verifying inputs and forwarding transactions to connected VPs. NVP's do not process messages locally as at this time they do not have a copy of either the blockchain or the world state. NVP architecture is evolving and this may change in the future. 

Applications can communicate with the fabric through an NVP or a VP.

See [Protocol Specification - 2.2.2 Multiple Validating Peers](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#222-multiple-validating-peers) for more information.

### Hyperledger Contract Connection and Flow

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

### Messaging Between Application and Contract

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
    - `args` : an array of strings as arguments to the function, in trade lane there is one argument that is a JSON encoded object 

The rest of this document is concerned with the contract's API as implemented in the `main.go` file.

---

## Contract API and Operation

User chaincode runs in a virtual environment such as a Docker container and communicates with its associated peer via the shim messaging protocol as shown in Diagram 2. The contract executable is built after extracting it from a github repository either locally as a github clone inside the peer build, or remotely at [github](http://github.com). 

Trade lane's physical structure is flat, using one `main.go` Go file and several related Go files in the top level folder. Together these embody `package main`. The main file implements the shim interface, the API delegation logic, and the contract's state and business logic.

Related go files add features that include (but are not limited to) asset history, contract state, recent states, alerts, and rules, also as a part of `package main`. See [README.md](./README.md). 

This file is in the `/docs` subfolder, along with other documents describing the related go files. The `/scripts` subfolder contains a Go script that is executed when the `go generate` command is typed at the command line while in the main folder context. The execution is controlled by the following comment line near the top of `main.go`:

``` go
//go:generate go run scripts/generate_go_schema.go
```

### Contract Payload and Delegation

The *JSON RPC* message envelope targets chaincode modes, which translates to the following functions and arguments in the shim's API:

- `Init` : called for `deploy` mode, initializes the contract, may delegate by switching on `function` but trade lane does not
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

### Ledger State

Transactions generally exist to change ledger state as the result of an incoming event. The state might pertain to the contract or all assets (e.g. recent state changes), or to an individual asset. Ledger state manipulation is performed with the stub API. 

Ledger state is stored as key:value (K:V) pairs with string keys and []byte values. The byte stream is often a cast of a string value. Hyperledger examples show string names as keys and converted integers (using iToA) as values while this contract chooses to store state as a JSON encoded object. 

The shim supports direct communication with other smart contracts, which is known as *chaincode to chaincode* invokes. This is useful for complex systems with multiple interacting smart contracts, each with a different asset focus. Examples of assets that could interact on separate subchains are devices, *things* (a.k.a. assets), legal contracts, warranties, financial transactions, and so on. The separation can also be implemented for security or business rule segmentation.  

> When an invoke transaction fails at some point in its execution, the entire transaction and all state changes in all participating contracts are rolled back. This is equivalent to a database rollback in traditional applications.

Code snippets for the ledger state APIs used in trade lane:

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

### State Keys

See [Protocol-Specification - 3.2 Ledger](https://github.com/hyperledger/fabric/blob/master/docs/protocol-spec.md#32-ledger)

It is important to define key patterns that can be managed in a straightforward manner without performing indirect key lookups. It is more efficient and clear to use asset IDs directly where possible. This contract stores values as JSON encoded object strings. 

Trade lane uses the following keys and patterns to store state:

- `assetID`: state for one asset
- `assetID + ".StateHistory"`: array of asset states in order of arrival, most recent first
- `ContractStateKey` : contract state including version, nickname, and array of managed assetIDs
- `RecentStatesKey` : array of state changes in order by time, most recent first, assets appear only once and jump to the front when they get a new state change

## Contract Patterns

### CRUD API Pattern

The contract payload definition (i.e. Contract API) is generic, so there are no restrictions on the names that can be used for the functions and arguments. On the other hand, following naming patterns offers advantages where automation is concerned, as demonstrated by the `monitoring_ui` that configures itself based on the CRUD pattern that trade lane uses for API function names.

The CRUD pattern is a familiar way to model database access and while it is not without its issues, it offers a clear and straighforward resource model for state manipulation and avoids API proliferation that is common as devices get smarter with larger subsets of state data sent in each device event.

The contract's primary key is `assetID`. As mentioned in earlier sections, an asset can be from a shipping container bar code to a package tracking number, a vehicle identifier (VIN) or anything else. What matters is that the asset is the focus of the contract. Trade lane implements a light weight shipping scenario and so assets are packages or containers.

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

### Partial State as Event Pattern

The term state is generally used as a short form for *World State*, which is the total of all K:V pairs in the ledger. However, state is also used when referring to the current values of all of an asset's properties, which are stored in a single object called `state`. The writable properties are grouped into a subset object called `event`, which models the incoming events sent by devices and applications. 

This contract is focused on assets and so are the `event` and `state` objects:
- `event` : defines all of the writable properties in the state.
- `state` : defines all asset properties, as in the `event` properties plus read-only properties such as `alerts` and `compliant`

#### Event

Having previously established that the function names in this contract API rigidly follow the CRUD pattern, the *partial state as event* pattern defines how arguments are shaped when passed through the `args` string array. Trade lane uses only `Args[0]` and expects a valid JSON encoded `event`, as defined in the [contract payload schema](../payloadSchema.json). 

This pattern's utility lies in using partial states as events rather than defining a large number of rigid event objects. With flexible property tags, JSON objects need not enforce completeness. It is perfectly valid to be missing all but the properties tagged as `required`. This contract's `event` object has only the `assetID` defined as required.

An important advantage of this approach for IoT device event processing is that the contract need not be updated and redeployed (i.e. upgraded) when a new composite device appears in the system. Previously unseen combinations of writable properties are always valid. For example, a new device that can send a location, a temperature and a gForce in a single event does not affect the contract's processing at all. But were this a rigidly defined API, a new event API would be required and the contract would need to be upgraded across the entire fabric. 

#### State

Asset state builds up as events are received. For example, an event with only the `assetID` would create an empty state, effectively registering the asset on the blockchain. A second event could then arrive with a location, and a third with a temperature. As each event is processed, a new state is calculated as the union of current state and incoming event, and then written into the ledger. 

Asset state is therefore the sum of all events with their calculations (e.g. alert status), encoded as a JSON string and stored into the ledger with the `assetID` as key. And just in case I have not repeated it often enough, incoming events are [deep merged](mapUtils.md) into asset state ad infinitem, each event creating a new asset state in the ledger. 

#### Ad Hoc Extension

The partial state as event pattern helps with object model extension in a running contract. Because JSON is tagged and flexible, a property has been introduced into `event` and `state` that is meant to offer extra flexibility for applications to store data alongside the pre-defined properties. 

This property is called `extension` and can be treated as a nested object to which application-managed state is added as necessary. This enables an application to add "sidecar" data to asset state without upgrading a contract. The deep merge algorithm handles extension data just as it handles schema defined properties. 

Treat this feature with care though, as the contract knows nothing of this data and thus cannot apply rules to it. But it can save a lot of time and cost when related data can be tracked alongside and is thus available to all monitors and applications. 

### Maps over Structs

The partial state as event pattern requires the use of another pattern, that being the use of maps over structs for internal object management. 

In Go, JSON will populate a struct from an incoming event using reflection. However, properties that are not present in the struct are lost, and all properties not present in the event are set to their zero value, which cannot be used to determine presence or absence of the property since `0` is a valid value with some fields such as temperature.

It is also difficult to generically merge an incoming event into a state struct since again there is no way to know which properties were sent. Structs necessitate a complex and rigid API based on getters and setters to be effective for event processing and state management. And in such a case all flexibility to adapt to new devices is lost.

> There is an alternate mechanism for managing the data model with structs where every member of every struct is defined as a pointer to a type, and where the presence or absence can be checked explicitly by testing for nil. Events and states based on this mechanism would create sparse nested structs. Working with these would create a fair amount of code to pass pointers back and forth and a `structUtils` module similar to the `mapUtils` module used in this contract would need to exist. This may be explored in a future derived sample.

Maps, on the other hand, naturally capture what is present in the event or state and offer some flexibility in dealing with incoming data. While it is also possible to validate incoming events against the contract payload schema, trade lane does not yet implement that check and favours instead a relaxed approach to extra properties.

Deep merging of sparsely populated nested event objects provides the accumulation and maintenance of the asset's state. After deep merging, the newly merged state is fed into the rule engine to calculate alerts and compliance. Thus, the merging in of any subset of properties is seen by all rules and thus lowers the likelihood of bugs in the common processing as the flows are all generic.