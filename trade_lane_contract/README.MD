This is a sample contract that explores a simple Trade Lane scenario, moving *assets* from one place to another. It is made up of several modules in `package main`, which are all small enough to avoid creating separate packages in subfolders in this first version.

This sample has been used to explore features and patterns that may be of interest to smart contract writers in the IoT domain. These are:

- A single contract instance that manages multiple assets
- A CRUD-like API for the assets
- A mechanism for storing asset history (note, this mechanism is early work and will change for better scaling)
- A mechanism for storing the most recent updates to any asset, most recent first. An asset can appear only once in the list and jumps to the top each time it is updated. 
- A simple alerts mechanism that tracks active alerts and marks the threshold events as raised or cleared.
- A simple rules engine that performs the actual threshold tests (e.g. temperature too high) and raises or clears alerts as necessary
- A set of map utilities that enable deep merging of incoming JSON events into the state that is stored in the ledger. This is necessary to implement a pattern where a partial state is used as an event. 
- The map utilities can also operate in case sensitive JSON tag mode (as specified for JSON-RPC) and in case-insensitive mode for the convenience of clients that do not want to be held to JSON-RPC standards (not recommended) 
- A simple logging facility that can be adjusted in real time in order to debug a contract without disrupting it in any way

This contract sample has a JSON Schema 4 compatible schema and a Go script that generates object samples and object schemas in Go form. The contract APIs `getAssetSamples` and `getAssetSchemas` return the generated samples and schemas in JSON object form. This is used by IoT Platform's device mapping so that contracts integrate automatically. Generic GUIs can also be generated using JSON UI plugins for React and Angular. This is demonstrated in another sample in this repository. 