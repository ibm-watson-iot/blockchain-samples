#OBSOLETE

This contract has been retired and will be expunged soon. Please instead use the [generic iot sample contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/advanced/iot_sample_contract).

#A Hyperledger sample contract with the Trade Lane schema

> **This smart contract is described in a set of documents in the `/docs` folder. The [main Go file documentation](./docs/main.md) includes a tutorial on the Hyperledger fabric and smart contract development. See the [README](./docs/README.md) in the `/docs` folder.**

The Trade Lane blockchain contract is derived from the 'openblockchain' contract in the folder `trade_lane_contract` and is adapted for the Hyperledger project.  

More advanced than the [Basic contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract_hyperledger), the Trade Lane contract sample provides an example of how you can implement a simple Trade Lane scenario of moving *assets* from one place to another. The contract consists of several modules in `package main`, and in the initial version, the modules small enough to avoid creating separate packages.

Use the Trade Lane Blockchain contract sample to track and manage your Watson IoT Platform asset data. By using this sample, you can develop a contract that includes any of the following features and patterns:

- A single contract instance that manages multiple assets.
- A create, retrieve, update, and delete operations API for assets.
- An event that is also a partial state.
- A deep merge of events to state.
- A `JSON Schema 4` compatible schema and a script written in the Go programming language that generates object samples and object schemas when `go generate` commands are issued.
- A mechanism for storing asset history. Note, that this mechanism is in the early stages and will require further changes to improve scalability.
- A mechanism for storing the most recent updates to any asset, storing the most recent first. An asset can appear only once in the list and jumps to the top each time that it is updated.
- An alerts mechanism that tracks active alerts and marks the threshold events as raised or cleared.
- A rules engine that performs threshold tests, for example, detects when the temperature too high, and raises or clears alerts as necessary. Note that the rules need not be limited to alerts testing and can also generate read-only properties directly as required.
- A set of map utilities that enable deep merging of incoming JSON events into the state that is stored in the ledger. This is necessary to implement a pattern where a partial state is used as an event.
- Optional case-insensitivity for JSON tags for the convenience of clients that do not want to be held to the strictness of the JSON-RPC standard. Note: insensitivity is not recommended, but can be explored with this sample.
- A logging facility that can be adjusted in real time in order to debug a deployed contract without disrupting it in any way.

Generic user interfaces exist in the other folders in this project and are driven from the schema in this and other contracts. Plugins for React are used to generate forms from the schema. Processing the schema directly also enables a host of other schema-driven features.

To enable application access to the schema, contract APIs `getAssetSamples` and `getAssetSchemas` return generated samples and schemas in JSON object form. These are used by the generic UIs and by the Watson IoT Platform for automated integration.
