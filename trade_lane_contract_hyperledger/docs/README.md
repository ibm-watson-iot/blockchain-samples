The document 
[`main.md`](main.md "main go file for trade lane sample contract")
discusses the main file of the trade lane sample contract for 
[Hyperledger](https://github.com/hyperledger), which is an open source blockchain fabric under the Linux Foundation. Your understanding of this contract's structure and implementation intent starts with this module, and then proceeds through the supporting modules listed below, in order by your areas of interest. 

> Note that the main document also contains a tutorial on the Hyperledger fabric and smart contracts.

* [`schema.md`](schema.md)
  - discusses the [JSON SCHEMA v4](http://json-schema.org/documentation.html) compatible schema in the file [`payloadSchema.json`](../payloadSchema.json), which defines the API and data model for the contract
* [`alerts.md`](alerts.md)
  - discusses the alerts module [`alerts.go`](../alerts.go), which supports contract-specific alert conditions
* [`rules.md`](rules.md)
  - discusses the rules engine in module [`rules.go`](../rules.go), which is called after every state change in order to perform property manipulation (including calculating new properties) and to test thresholds and raise alerts and then calculate contract compliance for the asset 
* [`mapUtils.md`](mapUtils.md)
  - discusses the module [`mapUtils.go`](../mapUtils.go) that supports property access in JSON objects plus deep merging of events into state in support of the *Partial State as Event* contract pattern
* [`recentStates.md`](recentStates.md)
  - discusses the module [`recentStates.go`](../recentStates.go), which manages a list of state changes across all assets sorted by time most recent first and containing only one event per asset
* [`contractState.md`](contractState.md)
  - discusses the module[`contractState.go`](../contractState.go), which manages contract state including version, nickname, and activeAssets list
* [`assetHistory.md`](contractState.md)
  - discusses the module [assethistory.go](`../assethistory.go`), which manages asset state history sorted by time with the most recent state first; each asset's history is stored in a separate bucket
