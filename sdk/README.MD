#SDK for IoT Smart Contracts for Hyperledger 

#UNDER CONSTRUCTION

##iot_sample_contract

This is a generic IoT sample contract that follows a *recorder* style in that senro readings or related data are part of an event object that is included in a state object for the target asset that would be the *Thing* in IoT. The sensor readings or equivalent come from *devices* or *applications*.

See the article  
[Introduction to Hyperledger Smart Contracts for IoT<br/>Best Practices and Patterns](docs/HyperledgerContractsIntroBestPracticesPatterns.md) 

##aviation_sample_contract

**Under Construction**

This contract derives from the IoT sample, introducing a complex data model with 5 asset classes and 4 separate CRUD APIs.

##docs

A set of articles discussing patterns, practices, customization of these samples, and the individual files that make up a contract and its features.

##common packages

The intent is to evolve to a set of shared packages that can be used to power any Hyperledger Smart Contract using IoT contract patterns like recorder, partial state as event, multi-asset, and crud API. 