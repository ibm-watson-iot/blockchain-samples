# Watson IoT Blockchain Sample Documentation

## Introduction

Watson IoT Contract Platform samples currently fall into three broad categories:

- simple single-asset hello world sample
- simple single-asset sample with numerous additional features
- advanced multi-asset sample with events and additional features

All of these use a schema to drive the related monitoring UI, which can be used to initiate any transaction or query that is exported for external use via the JSON configuration file.

>This repository is evolving, and there will be changes in how we partition the samples. One useful change will the packaging of the additional features. At this time, the features are implemented in *sidecar* files in package main. This tends to clutter the folder in which the smart contract resides and so will be addressed in time.

## Start Here

Read the [Introduction to Hyperledger Smart Contracts for IoT, Best Practices and Patterns](HyperledgerContractsIntroBestPracticesPatterns.md) and its specified pre-reading __before__ reading the detailed contract documents. 

Read the tutorial [Customizing the Sample Contract](CustomizingTheSampleContract.md) to get a feel for the simplicity and flexibility inherent in the *partial state as event* pattern that is used heavily in the trade lane sample contract.

When you have completed the pre-reading, start with [`main.md`](main.md) for detailed documentation on the IoT Trade Lane sample contract's main body and processing. Then move on to the detailed module documents shown below in your order of interest.

## Topics

* [`FAQ.md`](FAQ.md)
  - answers *frequently asked questions*, or perhaps more accurately *fervently anticipated questions*
* [`assetsAndEvents.md`](assetsAndEvents.md)
  - discusses how assets and events are defined, modelled and used in these samples
* [`alerts.md`](alerts.md)
  - discusses the alerts module [`alerts.go`](../alerts.go), which supports contract-specific alert conditions
* [`generator.md`](generator.md)
  - discusses the schema processor script [`generate_go_schema.go`](../scripts/generate_go_schema.go), which is executed when the `go generate` command it typed at the main folder level; the schema processor checks syntax and validates as JSON Schema 4 compatible, generating object model and API samples and schema in Go language for the use of the contract and applications 
* [`main.md`](main.md "main go file for trade lane sample contract")
  - discusses the main module [`main.go`](../main.go) of the trade lane sample contract for [Hyperledger](https://github.com/hyperledger), which is an open source blockchain fabric under the Linux Foundation. Your understanding of this contract's structure and implementation intent starts with this module, and then proceeds through the supporting modules listed below, in order by your areas of interest. 
* [`mapping.md`](mapping.md)
  - a general discussion of device schema to contract schema mapping as an illustration of how each views its own data model and how these will intersect in an IoT application
* [`mapUtils.md`](mapUtils.md)
  - discusses the module [`mapUtils.go`](../mapUtils.go) that supports property access in JSON objects plus deep merging of events into state in support of the *Partial State as Event* contract pattern
* [`rules.md`](rules.md)
  - discusses the rules engine in module [`rules.go`](../rules.go), which is called after every state change in order to perform property manipulation (including calculating new properties) and to test thresholds and raise alerts and then calculate contract compliance for the asset 
* [`schema.md`](schema.md)
  - discusses the [JSON SCHEMA v4](http://json-schema.org/documentation.html) compatible schema in the file [`payloadSchema.json`](../payloadSchema.json), which defines the API and data model for the contract

## Under Construction

* [`assetHistory.md`](contractState.md)
  - discusses the module [assethistory.go](`../assethistory.go`), which manages asset state history sorted by time with the most recent state first; each asset's history is stored in a separate bucket
* [`contractState.md`](contractState.md)
  - discusses the module[`contractState.go`](../contractState.go), which manages contract state including version, nickname, and activeAssets list
* [`recentStates.md`](recentStates.md)
  - discusses the module [`recentStates.go`](../recentStates.go), which manages a list of state changes across all assets sorted by time most recent first and containing only one event per asset
