#A Hyperledger Sample Contract with Trade Lane Schema

> **This smart contract is described in a set of documents in the `/docs` folder. The [main Go file documentation](./docs/main.md) includes a tutorial on the Hyperledger fabric and smart contract development. See the [README](./docs/README.md) in the `/docs` folder.**

This contract is derived from the openblockchain contract in the folder `trade_lane_contract` to adapt it to the hyperledger project.  

This sample contract implements a simple Trade Lane scenario, moving *assets* from one place to another. It consists of several modules in `package main` that are all small enough to avoid creating separate packages in this first version.

This sample is used to explore features and patterns that are of interest to smart contract writers in the IoT domain. These are:

- A single contract instance that manages multiple assets
- A CRUD-like API for the assets
- An event that is also a partial state
- A deep merge of events to state
- a `JSON Schema 4` compatible schema and a script written in Go that generates object samples and object schemas when `go generate` commands are issued
- A mechanism for storing asset history (note, this mechanism is early work and will change for better scaling)
- A mechanism for storing the most recent updates to any asset, most recent first. An asset can appear only once in the list and jumps to the top each time it is updated. 
- An alerts mechanism that tracks active alerts and marks the threshold events as raised or cleared.
- A rules engine that performs threshold tests (e.g. temperature too high) and raises or clears alerts as necessary (and note that the rules need not be limited to alerts testing etc, they can in fact generate read-only properties directly if need be)
- A set of map utilities that enable deep merging of incoming JSON events into the state that is stored in the ledger. This is necessary to implement a pattern where a partial state is used as an event. 
- Optional case-insensitivity for JSON tags for the convenience of clients that do not want to be held to the strictness of the JSON-RPC standard (note: insensitivity is not recommended, but can be explored with this sample) 
- A logging facility that can be adjusted in real time in order to debug a deployed contract without disrupting it in any way

Generic UIs exist in other folders in this project, driven from the schema in this and other contracts. Plugins for React are used to generate forms from the schema, and of course processing the schema directly would enable a host of other schema-driven features. 

To enable application access to the schema, contract APIs `getAssetSamples` and `getAssetSchemas` return generated samples and schemas in JSON object form. This is used by the generic UIs and by the Watson IoT Platform for automated integration.
