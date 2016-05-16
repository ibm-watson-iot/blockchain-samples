#Customize the Trade Lane Sample Contract for Hyperledger
##Prerequisite
Please read the document [Hyperledger Contracts Introduction to Best Practices and Patterns](./HyperledgerContractsIntroBestPracticesPatterns.md) before reading this document as it helps to have an understanding of the flow that is intended for the sample contract before adding new features.
## Introduction
This generic IoT sample smart contract is intended as a template in the spirit of a well-featured *hello world*. It contains customizations that simulate a simple trade lane scenario: shipping a consignment in an identifiable container or package between locations. There can be many such simultaneous shipments.

This contract tracks some of the key parameters in a simple fulfillment scenario: 

![Simple Fullfilment per UBL 2.1.10](http://docs.oasis-open.org/ubl/os-UBL-2.1/art/UBL-2.1-Fulfilment-1simple.png "Universal Business Language 2.1 Specification")

> See [UBL 2.1.10](http://docs.oasis-open.org/ubl/os-UBL-2.1/UBL-2.1.html#S-SHIPMENT-CONSIGNMENT) for a description of the many possible relationships between shipments and consignments.   

###Scenario
For simplicity, the consignment, shipper, receiver, consignor and consignee are not identified. The container / package *is* tracked and identified by `assetID` as per the singleton contract instance and multiple managed asset patterns used in this contract. The container's location and temperature are tracked and the carrier that bears responsibility is followed through events that set a new carrier whenever it changes. 

The event stream is fluid and there is no formal state machine with a managed set of statuses. For example, the contract does not track statuses like `idle`, `packing`, `loading`, `shipped`, `arrived`, and so on. 

The container is considered to be moving when location events are received. It is considered contain frozen goods when temperature events are received. This because the contract has an explicit rule that throws the `OVERTEMP` alert when the temperature exceeds 0 degrees Celsuis exclusive. Carrier events are sent when a new carrier assumes reponsibility, which implies that the container has been moved. The state machine is can be thought of as *implied* in this sample. 

###Asset State Management
This sample makes use of the *partial state as event* pattern as introduced in the document [*introduction to hyperledger best practices and patterns*](./HyperledgerContractsIntroBestPracticesPatterns.md). Events arrive ad hoc with the asset state remembering each property as last set by an event. 

In other words, an asset's state builds up as events are received. This pattern allows  creation or update events like `createAsset` or `updateAsset` to carry any combination of the writable properties of the state. 

> Events received by the `updateAsset` function will automatically redirect to the `createAsset` function by default if the asset is not found. This behavior can be changed (as in toggled off or on) with this message:

>``` go
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"setCreateOnUpdate",
            "args":["{\"createOnUpdate\":false}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}
```

The asset state is essentially the same when a carrier, a temperature and a geo location all arrive as discrete events, or when a single event arrives carrying all three properties. 

###Rules and Alerts
Differences that can be observed between these two scenarios are related to the calculated alert status in this contract:

- the alert state will be different at the end of the triplet, depending on which contains the temperature, with the raised status being only set once for any given transition gtom cleared to raised, so raised will be set of the temperature is the last event to be received and it will be off if temperature is the first or second to be received
  - in the second scenario, the temperature is included in the only message, and so raised is definitely on after the event is processed
  - such behaviors are expected and make perfect sense when the thresholds are interpreted correctly, as in they mean specifically that *the event that generated this asset state raised this alert*, and such a threshold can only happen once in a row
- in cases where there is a dependency between properties for a complex alert, the order in which they are received may matter, although it would be wise to guard against such a non-deterministic behavior

This flexibility in order of arrival and complexity of event makes such a contract relatively immune to API expansion issues where composite devices with different combinations of sensors might force new API to be developed and deployed in a less flexible design.

And so this contract specifically defines these behaviors:

- any combination of properties can arrive in any combination of *discrete* or *composite* events
- temperature events are relevant to frozen goods and should not be sent when the consignment does not contain frozen goods
  - the `OVERTEMP` rule uses an inclusive threshold of zero degrees Celsius
  - corollary: temperature events not sent indicates that the container does not contain frozen goods
- an asset can be created by any event if the default `createOnUpdate` behavior is not changed
- asset state builds up as properties are received in events, discrete or otherwise
- applications do not need to manage state externally
  - they can ask the contract for the current world state of any asset
- the `OVERTEMP` rule is sensitive to the absence of a temperature property and will always clear the alert when no temperature is present
  - this is a critical behavior for the alerts module

Of course, should the lack of a temperature property be an alertable condition, then the rule should raise an alert (e.g. `MISSINGTEMP`) when a temperature is missing for a specified period of time.

In a more elaborate contract, a consignment would be modelled explicitly as a member of the container's contents manifest. Each consignment could be typed in some way with defined thresholds, making alerts sensitive to the consignment type. 
- Defining types known to the rule with associated thresholds has many ways of becoming complex and brittle. Built-in constants would be the obvious approach, but that is inflexible - especially when new goods types and associate thresholds appear. Injecting a new data model from outside is an option, but is beyond the scope of this tutorial.

Or the consignment could have an explicit temperature threshold set from the application such that the rule is sensitive to the threshold's presence or absence and uses the threshold when present to calculate compliance for that consignment. 
- Allowing the application to set the threshold makes sense in several ways: the user that is managing the loading has the best idea of what the consignment contains and can apply some level of external checking to ensure the most appropriate alert threshold; and the contract has no limitations as to types of goods or thresholds to be applied. The contract is entirely flexible. 

##System Engineering Considerations
A blockchain provides a distributed, resilient, indelible, auditable, transparent, and shared ledger that can build trust in trustless environments. An example is a system built for collaboration between competitors and regulatory agencies. Blockchains come with performance-related constraints owing to the inherent delays and state hysteresis in systems where transactions are committed to blocks and world state *only after they achieve consensus*. 

There is an architectural balance to be struck in the system's data model. Data stored in world state is subject to consensus delays (that manifests as hysteresis between transaction execution and state availability to queries) while data stored in back-office servers is typically not. 

This behavior affects both the speed and volume at which data can be added to the blockchain, and the responsiveness of applications that interact with either the back-office or the blockchain is therefore different and will feel at times variable. 

Further, a *fire-hose* of real-time data from an IoT device will overwhelm a blockchain, but streams of processed data with redundancy removed and at an appropriate cadence will work. But do note that, in cases where immediate real-time access to posted data is a requirement, a blockchain will not be likely to meet the goal.

The single most important detail when designing a large system combining a real-time data handling component and a blockchain component is to ensure an appropriate separation between real-time processing with a need for immediacy in data access, and the slower data stream that is intended for the blockchain. And note again that *slower* does not imply *slow*, but rather *not immediate*.

This sample contract therefore assumes a simplified scenario such that:
- the blockchain is tracking properties as measured by devices at regular, but not real-time, intervals
- device event data is sent to the *Watson IoT Platform* and mapped onto contract API
- one or more monitor applications are polling asset world state and acting on changes
  - by polling world state, the inherent hysteresis is no longer a concern for most functions

##A Simple Customization
The simplest possible customizations involve changes mainly to the data model and related rules. This will be common when a system or application uses a blockchain to store a record of device events and asset states without performing complex business operations. 

This sample contract is one such data-centric implementation in that the primary application-specific logic is the handling of alerts for temperatures exceeding a threshold. 

Te *partial state as event* pattern is therefore the primary event and state processing algorithm used in this contract sample, and is the default processing that is inherited by derived contracts that use this sample as a template. 

So as assets move from place to place, the contract receives location updates, temperature updates and occasionally transfers of custodianship. Events therefore contain any combination of temperature, geo location, carrier and RFC3339nano timestamp properties. 

As discussed above, the temperature rule's behavior can be improved by the addition of a threshold that is specific to the container contents. To keep the tutorial simple, a threshold property is added to asset state for the container only. The `OVERTEMP` rule is adjusted to become sensitive to the presence or absence of the new threshold property. 

> Adding the threshold property as an event requires no specialized logic as the deep merge of event into state handles the new property. The combined asset state is presented to the rules engine before committing to the ledger so that the adjusted `OVERTEMP` rule can now see the temperature and the new threshold.

###Schema Change
Two changes are required in the [schema](../payloadSchema.go): 
 - add the `threshold` property to the event object
 - copy it to the state object

>Explicit duplication of properties between event and state as opposed to import of a common set of properties is a matter of history and convenience at this time. To be addressed in a future version of the schema.

Before:

``` json
        "event": {
            "type": "object",
            "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "assetID": {
                    "$ref": "#/definitions/assetID"
                },
                "timestamp": {
                    "type": "string",
                    "description": "RFC3339nanos formatted timestamp."
                },
                "location": {
                    "$ref": "#/definitions/geo"
                },
                "extension": {
                    "type": "object",
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {}
                },
                "temperature": {
                    "type": "number",
                    "description": "Temperature of the asset in CELSIUS."
                },
                "carrier": {
                    "type": "string",
                    "description": "transport entity currently in possession of asset"
                }
            },
            "required": [
                "assetID"
            ]
        },
```

After:

``` json
        "event": {
            "type": "object",
            "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "assetID": {
                    "$ref": "#/definitions/assetID"
                },
                "timestamp": {
                    "type": "string",
                    "description": "RFC3339nanos formatted timestamp."
                },
                "location": {
                    "$ref": "#/definitions/geo"
                },
                "extension": {
                    "type": "object",
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {}
                },
                "temperature": {
                    "type": "number",
                    "description": "Temperature of the asset in CELSIUS."
                },
                "threshold": {
                    "type": "number",
                    "description": "Temperature threshold inclusive in CELSUIS."
                },
                "carrier": {
                    "type": "string",
                    "description": "transport entity currently in possession of asset"
                }
            },
            "required": [
                "assetID"
            ]
        },
```

### Rule Change
Once the threshold is added to the data model, the `OVERTEMP` rule can become sensitive to it. 

> The goal in this style of contract is to build up asset state from incoming events, storing the calculated state in *world state*. Just before the state calculation is completed, the rules engine is run against the new asset state, which must raise or clear its specific alert for the target asset.

The `OVERTEMP` rule is defined in [rules.go](../rules.go).

Before:

``` go
func (alerts *AlertStatusInternal) overTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
                return
            }
        }
    }
    alerts.clearAlert(AlertsOVERTEMP)
}
```

After:

This is a larger change and yet is extremely straighforward.

``` go
func (alerts *AlertStatusInternal) overTempRule (a *ArgsMap) {
    var temperatureThreshold  float64 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            tbytes, found = getObject(*a, "threshold")
            if found {
                temperatureThreshold, found = tbytes.(float64)
                if found {
                    if t > temperatureThreshold {
                        alerts.raiseAlert(AlertsOVERTEMP)
                        return
                    }
                }
            }
        }
    }
    alerts.clearAlert(AlertsOVERTEMP)
}
```

The style here is obvious and for most multi-property-dependent rules should be considered mandatory. To clarify the technique: only when every property has been found and correctly asserted to be of the right type should the calculation proceed, raising or clearing the alert as the rule demands. In **all** other cases, the code should fall through to clear the alert.

>The assignments for tbytes and t use the Go ":=" operator, which creates a new variable without a previous declaration and with an inferred type. For clarity with the newly-dynamic threshold, I changed the declaration of the threshold from a constant to a variable, assigning it from the JSON property, if it exists. Since the variable is predeclared, the ":=" is changed to "=" in the assignment. The compiler actually found the error where I left in the ":=" originally because it is pretty subtle difference. This is a good use of the Go compiler since it is blazingly fast.

###Generate and Build
Once the schema changes have been made, the `go generate` command can be executed at the root level of the contract. It relies on a comment at the top of the [main.go](../main.go) file:

``` go
//go:generate go run scripts/generate_go_schema.go
```
The named go script is executed in the scripts folder and that script generates the `schemas.go` and `samples.go` files that are used to define the key schema elements to be sent to any application when asked for. The Watson IoT Platform uses the `getAssetSchemas` call to return the contents of `schemas.go` and processes the API and Data Model elements returned for its device event to contract event mapping feature.

The execution of `go generate` looks like:

``` sh
vagrant@hyperledger-devenv:v0.0.9-5cd67fd:/local-dev/github.ibm.com/blockchain-samples-wip/trade_lane_contract_hyperledger$ go generate
JSON CONFIG FILEPATH:
   /local-dev/github.ibm.com/blockchain-samples-wip/trade_lane_contract_hyperledger/scripts/generate.json
Generate Go SCHEMA file schemas.go for:
   [init createAsset updateAsset deleteAsset deletePropertiesFromAsset deleteAllAssets readAsset readAllAssets readAssetHistory readRecentStates setLoggingLevel setCreateOnUpdate] and:
   [assetIDandCount assetIDKey initEvent event state]
Generate Go SAMPLE file samples.go for:
   [initEvent event state contractState]
vagrant@hyperledger-devenv:v0.0.9-5cd67fd:/local-dev/github.ibm.com/blockchain-samples-wip/trade_lane_contract_hyperledger$
```
The output documents what the script was told to generate based on its [configuration file](../scripts/generate.json). The convention is to always generate the entire API so that applications can generate forms or messages to the contract.

The data model can be very handy, since it isolates certain structures by name, rather than the API inputs, which imply an incoming event without specifying it as `event`.

The build command looks like this:
``` sh
vagrant@hyperledger-devenv:v0.0.9-5cd67fd:/local-dev/github.ibm.com/blockchain-samples-wip/trade_lane_contract_hyperledger$ go build
vagrant@hyperledger-devenv:v0.0.9-5cd67fd:/local-dev/github.ibm.com/blockchain-samples-wip/trade_lane_contract_hyperledger$
```
There is no output when it works.

##Testing the New Contract
The author assumes that the reader has been through the [sandbox document](https://github.com/hyperledger/fabric/blob/master/docs/API/SandboxSetup.md) and has been able to build and run the peer and the contract in two separate terminal windows.

> Tips:

> - The author runs everything from build to execution in the Vagrant environment for its consistency and clarity
- The author chooses to use git shell windows running bash on a Windows system and the output in this document is all from such a window
- Follow the sandbox instructions *exactly* as they have changed substantially recently

Once the new contract and the peer are running as per the sandbox instructions, it is time to test the contract's new feature.

Start by deploying the contract in debug mode, which means naming it "mycc".

The command in POSTMAN is:

``` json
{
    "jsonrpc": "2.0",
    "method": "deploy",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"init",
            "args":["{\"version\":\"4.0\",\"nickname\":\"THRESHOLD\"}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}
```

The response is OK in this case, since both are running and communicating:

``` json
{
  "jsonrpc": "2.0",
  "result": {
    "status": "OK",
    "message": "mycc"
  },
  "id": 1234
}
```

The logs will show that the contract executed, the nickname *THRESHOLD* standing out in the logs:

``` sh
2016/05/16 06:50:24 [THRESHOLD-4.0] DEBU PUTContractState: []interface {}{main.ContractState{Version:"4.0", Nickname:"THRESHOLD", ActiveAssets:map[string]bool{}}}
2016/05/16 06:50:24 [THRESHOLD-4.0] INFO Contract initialized
```

To test the threshold, we will need an asset with a threshold in it. We will use `ASSET1` and `100` as the values. No other data will be sent in the initial event. 

``` json`
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"createAsset",
            "args":["{\"assetID\":\"ASSET1\",\"threshold\":100}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}
```
The return message specifies the transaction ID in the `message` property and the `status` is `OK`.

``` json
{
  "jsonrpc": "2.0",
  "result": {
    "status": "OK",
    "message": "bf80f7cb-46da-4dd1-92b6-02c7c402b0ea"
  },
  "id": 1234
}
```

To verify the asset exists, we can perform a `readAsset` on `asset1`:

``` json
{
    "jsonrpc": "2.0",
    "method": "query",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"readAsset",
            "args":["{\"assetID\":\"ASSET1\"}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}
```

Queries return the entire response immediately:

``` json
{
  "jsonrpc": "2.0",
  "result": {
    "status": "OK",
    "message": "{\"assetID\":\"ASSET1\",\"incompliance\":true,\"lastEvent\":{\"args\":\"{\"assetID\":\"ASSET1\",\"threshold\":100}\",\"function\":\"createAsset\"},\"threshold\":100,\"timestamp\":\"2016-05-16T06:56:51.055683317Z\"}"
  },
  "id": 1234
}
```

> The JSON RPC 2.0 envelope for contract payloads stringifies inputs and outputs, which shows up as escaped strings. This means that the examples from POSTMAN cannot display objects in *pretty* format. 

The `threshold` tag is clearly visible and its value is 100 degrees Calcuis. The contract shows as being in compliance, since there is no temperature event so far. With nothing to calculate, the rule must clear the alert.

Now we will send a temperature of 99 and the contract will show as being in compliance again. 

>From here on, `updateAsset` and `readAsset` calls are used but the response from the invoke and the command for the query are left out as unnecessary detail.

``` json
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"updateAsset",
            "args":["{\"assetID\":\"ASSET1\",\"temperature\":99}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}

{
  "jsonrpc": "2.0",
  "result": {
    "status": "OK",
    "message": "{\"assetID\":\"ASSET1\",\"incompliance\":true,\"lastEvent\":{\"args\":\"{\"assetID\":\"ASSET1\",\"temperature\":99}\",\"function\":\"updateAsset\"},\"temperature\":99,\"threshold\":100,\"timestamp\":\"2016-05-16T07:07:35.557031032Z\"}"
  },
  "id": 1234
}
```

The temperature has now shown up and is below the threshold, so the contract remains in compliance.

And finally, we send the event with temperature as 101 and the contract will go out of compliance. 

``` json
{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"updateAsset",
            "args":["{\"assetID\":\"ASSET1\",\"temperature\":101}"]
        },
        "secureContext": "user_type1_400022b418"
    },
    "id":1234
}

{
  "jsonrpc": "2.0",
  "result": {
    "status": "OK",
    "message": "{\"alerts\":{\"active\":[\"OVERTEMP\"],\"raised\":[\"OVERTEMP\"],\"cleared\":[]},\"assetID\":\"ASSET1\",\"lastEvent\":{\"args\":\"{\"assetID\":\"ASSET1\",\"temperature\":101}\",\"function\":\"updateAsset\"},\"temperature\":101,\"threshold\":100,\"timestamp\":\"2016-05-16T07:12:02.802425472Z\"}"
  },
  "id": 1234
}
```

- The temperature now shows 101, which is above the threshold of 100. Thus, we now see the `OVERTEMP` alert as both raised and active. 
  - Raised means that this specific event raised the alert by changing it from inactive state to active state, and active says that the temperature for this asset is too high at this point in time.
- The `incompliance` property is now missing, which means that the asset is no longer compliant with the terms in the contract.

##Conclusion
This was a basic introduction to customization. The desceptively simple changes in this article add a significant feature to the contract in that it can now deal with temperatures that are specific to an asset's cargo.

The flexibility and simplicity of the *partial state as event* pattern makes data-driven contracts easy to construct and maintain and should be the default design until it fails to meet the goals of the contract.