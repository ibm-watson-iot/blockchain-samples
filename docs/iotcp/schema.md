# Schema

[`payloadSchema.json`](../payloadSchema.json "the contract schema for API and object model")

## Introduction

As referenced in other documents in this folder, the REST messaging protocol for communicating with the block chain and with smart contracts is documented in a file called 
[`rest_api.json`](https://github.com/hyperledger/fabric/blob/master/core/rest/rest_api.json "on hyperledger project in github.com"). This file documents all service end points to which a client application can GET, PUT and POST messages.

This document describes the schema file - 
[`payloadschema.json`](../payloadschema.json "in this folder hierarchy"), documenting the payloads carried by RESTful messages specifically for execution by a smart contract. 

This payload schema describes both the contract's API and its data model. The schema is 
[JSON Schema 4](http://json-schema.org/ "JSON Schema 4 specifications on the web") compatible
and works with compatible tools. (See the generic UI referenced below.) 

The contract API follows a CRUD-like pattern that, combined with other important patterns, have proven useful in contract development with a strong IoT component. Read the 
[tutorial on contract patterns](TutorialSmartContractPatterns.md "explains the key patterns on which the trade lane sample contract is based") for more information on how the schema APIs, data model, and IoT Platform offer an end to end environment for IoT applications.

>*Note that we maintain a semblance of Swagger compatibility in order that the [Swagger editor](http://editor.swagger.io/#/ "on the web") can be used to explore the schema, but Swagger supports only a small subset of JSON Schema 4 and so is not as useful as it could otherwise be when describing payloads. It is rigidly oriented towards resource-oriented RESTful APIs and the JSON-RPC flavour of the hyperledger chaincode (a.k.a. smart contract) API is not a perfect fit. Swagger's editor is not needed, though, for error checking of the schema. This is covered in the next section*

## Schema Error Checking

Schema error checking is accomplished with a script implemented in Go. It checks the schema for errors when you run the `go generate` command in the main folder of any contract derived from this sample. This line near the top of the main contract is responsible for connecting the contract to the script and schema:

`//go:generate go run scripts/generate_go_schema.go`

The script resides in the `/scripts` folder and assumes the presence of a schema in the main contract folder. It reads from that file and immediately unmarshals the JSON in order to check for errors. The output of a contrived error looks like:

``` text

vagrant@hyperledger-devenv:v0.0.9-4ba4f96:/wip/trade_lane_contract$ go generate
JSON CONFIG FILEPATH:
   /wip/trade_lane_contract/scripts/generate.json
*********** UNMARSHAL ERR **************
 invalid character 't' looking for beginning of object key string
Error in line 49: invalid character 't' looking for beginning of object key string
"definitions": { this is an error
                 ^

```

Remove the offending text *this is an error* and rerun the `go generate` command to get this output instead:

``` bash

vagrant@hyperledger-devenv:v0.0.9-4ba4f96:/wip/trade_lane_contract$ go generate
JSON CONFIG FILEPATH:
   /wip/trade_lane_contract/scripts/generate.json
Generate Go SCHEMA file schemas.go for:
   [init createAsset updateAsset deleteAsset deletePropertiesFromAsset deleteAllAssets readAsset readAllAssets readAssetHistory readRecentStates setLoggingLevel setCreateOnUpdate] and:
   [ChaincodeOpPayload ChaincodeOpSuccess ChaincodeOpFailure assetIDandCount assetIDKey initEvent event state]
Generate Go SAMPLE file samples.go for:
   [ChaincodeOpPayload ChaincodeOpSuccess ChaincodeOpFailure initEvent event state contractState]

```

The output specifies what the script generated based on the commands in the JSON configuration file.

## Schema File Generation

When the unmarshal test is successful, the script generates contract-consumable Go files named 
[`schemas.go`](schemas.go "in the main contract folder") and
[`samples.go`](samples.go "in the main contract folder"). These automatically become part of the contract build.

The exported data structures are then exposed to clients through a pair of APIs, which by convention are named `getAssetSchemas` and `getAssetSamples` (as implemented in this sample contract).

Clients can make these calls and use them as desired. Examples include displaying samples of specific APIs or data inputs or outputs for debugging purposes (a useful way to explore a running contract without accessing its source code), mapping device inputs dynamically (which is how the IoT Platform Proxy can offer mapping services from any device to any contract), and validating parameters (although there are also good reasons for flexibility with properties accepted in state changing API). 

To see a generic GUI in action, explore the [generic UI](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/generic_ui "on github") and watch as the UI self-initializes from the schema. 

The selection of APIs and objects that appear in the returned schemas or samples are configured in the script's companion [JSON configuration file](scripts/generate.json "in the scripts folder"). 
When extending the sample contract for a specific domain, client applications can rely on these APIs being accurate if the schema file is maintained and a few conventions are followed.

Summary of conventions that enable integration with the IoT Platform and generic tooling:
+ implement getAssetSchemas
+ run `go generate` for every schema change and commit those files
+ ideally, use the prefixes create, read, update and delete for invokes and queries
+ retain the schema name `payloadSchema.json` and leave it in the main contract folder

## Customizing the Schema

There is more to customizing the contract than meets the eye, so this information is gathered in the [tutorial on customizing the sample contract](docs/CustomizingTheSampleContract.md "explains how to extend the sample contract for your own domain").

## Additional Information: Wiring the Contract into the Fabric

The API for a contract has been updated for the hyperledger project to a 
[JSON RPC](http://json-rpc.org/ "http://json-rpc.org/") compatible REST messaging protocol. The endpoint for these messages is:

- /chaincode

This service endpoint expects POST messages at a peer in the fabric. The payload for the message
is defined in the schema element `ChaincodeOpPayload`, which contains a JSON RPC `method` that can contain only three service endpoint names. These can also be thought of as *modes* of operation:

- `deploy` : chaincode is deployed and initialized
- `invoke` : transaction modifies uncommitted state asynchronously
- `query` : query committed state and return results synchronously

In addition to the `method`, an integer `id` (JSON RPC specifies string, integer or null and this will be addressed in the future) is sent by the client with every transaction and is then returned in the asynchronous response, allowing the client to track transaction progress. (Asynchronous responses under development at the time of writing.)

This outer JSON RPC payload also has `params`, which contain the inner contract payload, defined in the element `ChaincodeSpec`. The salient elements in ChaincodeSpec are:

- `function` : a string that specifies an internal contract function
- `args` : an array of strings as arguments to the *function*

These functions are defined in the schema in the section called API. This looks like:

``` json

        "API": {
            "type": "object",
            "description": "The API for the tradelane sample contract consisting of the init function, the crud functions to change state, and a set of query functions for asset state, asset history, recent states, and so on.",
            "properties": {
                "init": {
                    "type": "object",
                    "description": "Initializes the contract when started, either by deployment or by peer restart.",
                    "properties": {
                        "method": "deploy",
                        "function": {
                            "type": "string",
                            "enum": [
                                "init"
                            ],
                            "description": "init function"
                        },
                        "args": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/initEvent"
                            },
                            "minItems": 1,
                            "maxItems": 1,
                            "description": "args are JSON encoded strings"
                        }
                    }
                },
                "createAsset": {
                    "type": "object",
                    "description": "Create an asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
                    "properties": {
                        "method": "invoke",
                        "function": {
                            "type": "string",
                            "enum": [
                                "createAsset"
                            ],
                            "description": "createAsset function"
                        },
                        "args": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/event"
                            },
                            "minItems": 1,
                            "maxItems": 1,
                            "description": "args are JSON encoded strings"
                        }
                    }
                },

```

__NEW__ The schema now documents the JSON RPC `method` property, which can be `deploy`, `invoke` or `query`, so that a mapping or other dynamic user interface can determine which functions are intended as invokes and which are queries.

Contract API is represented as objects inside the API object with names to match the contract's functions. The schema shows the delegation to these functions as a tree from the main chaincode input through layers that are named for deploy, invoke, and query. The actual contract function schemas are shaped to match the expected payload with `function` and `args` being specified. With queries, results are also specified.

Note that this does not intend to replace the JSON RPC schema for the outer JSON RPC protocol, but rather the embedded contract API.

The schema can be explored in the Swagger editor, but Swagger errors on choice verbs like `oneIf` must be ignored. It can show the API options and allow you to drill down into the parameters etc. This is sometimes useful for interactive learning, but in fact the schema can be generated in its entirety and explored on the screen. 

The contract API approximates JSON-RPC compatibility withint the limits of the payload definitions to which we are restricted by hyperledger. There are plugins for `react` and `angular` that can interpret the schema and generate forms etc. These can be seen in the UI folders in this project.
