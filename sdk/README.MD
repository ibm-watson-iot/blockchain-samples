#SDK for IoT Smart Contracts for Hyperledger 

#UNDER CONSTRUCTION

##iot_sample_contract

This is a generic IoT sample contract that follows a *recorder* style in that sensor readings or related data are part of an event object that is included in a state object for the target asset that would be the *Thing* in IoT. The sensor readings or equivalent come from *devices* or *applications*.

See the article  
[Introduction to Hyperledger Smart Contracts for IoT<br/>Best Practices and Patterns](docs/HyperledgerContractsIntroBestPracticesPatterns.md) 

##simple_aviation_sample_contract

A simple extension of the IoT sample contract that uses a *composite event* consisting of several separate objects and a common set of properties. The rules are set up to show how one rule can fire an alert (e.g. too many hard landings) and another can cancel it (e.g. an inspection has been performed.)

##aviation_sample_contract

**Under Construction**

This contract derives from the IoT sample, introducing a complex data model with 3 asset classes and several event functions. There are abstracted layers that can be used to quickly implement new asset classes. Docs for this are TBD. 

##docs

A set of articles discussing patterns, practices, customization of these samples, and the individual files that make up a contract and its features.

##common packages

The intent is to evolve to a set of shared packages that can be used to power any Hyperledger Smart Contract using IoT contract patterns like recorder, partial state as event, multi-asset, and crud API. 
