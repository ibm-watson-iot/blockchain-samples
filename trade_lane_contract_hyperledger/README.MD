This contract is derived from the contract in the folder `trade_lane_contract` to adapt it to the hyperledger project. 

This sample contract implements a simple Trade Lane scenario, moving *assets* from one place to another. It consists of several modules in `package main` that are all small enough to avoid creating separate packages in this first version.

This sample is used to explore features and patterns that are of interest to smart contract writers in the IoT domain. These are:

- A single contract instance that manages multiple assets
- A CRUD-like API for the assets
- An event that is also a partial state
- A deep merge of events to state
- a `JSON Schema 4` compatible schema and a script written in Go that generates object samples and object schemas when `go generate` commands are issued
- A mechanism for storing asset history (note, this mechanism is early work and will change for better scaling)
- A mechanism for storing the most recent updates to any asset, most recent first. An asset can appear only once in the list and jumps to the top each time it is updated. 
- A simple alerts mechanism that tracks active alerts and marks the threshold events as raised or cleared.
- A simple rules engine that performs the actual threshold tests (e.g. temperature too high) and raises or clears alerts as necessary
- A set of map utilities that enable deep merging of incoming JSON events into the state that is stored in the ledger. This is necessary to implement a pattern where a partial state is used as an event. 
- Optional case-insensitivity for JSON tags for the convenience of clients that do not want to be held to the strictness of the JSON-RPC standard (note: insensitivity is not recommended, but can be explored with this sample) 
- A logging facility that can be adjusted in real time in order to debug a deployed contract without disrupting it in any way

The contract APIs `getAssetSamples` and `getAssetSchemas` return generated samples and schemas in JSON object form. This is used by the IoT Platform's device mapping so that contracts integrate automatically.

Generic GUIs can also be generated using JSON UI plugins for React and Angular. This is demonstrated in the UI samples in this repository. 