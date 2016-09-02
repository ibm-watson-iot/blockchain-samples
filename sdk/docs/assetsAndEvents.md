# Assets and Events

## Introduction

### **Asset**

The focus of an IoT smart contract is often referred to as an asset. Some tamgible thing that  For example, in a contract designed to track warranty status on cars, it could be the warranty itself, or the car. 

For simple contracts, the primary asset is selected and the API is designed to manipulate that asset class. As discussed in the [Introduction to Hyperledger Smart Contracts for IoT, Best Practices and Patterns](HyperledgerContractsIntroBestPracticesPatterns.html) document, these samples use a pattern based on a simplified CRUD (create, read update, delete) API for manipulating assets. The IoT sample, for example, had functions called `createAsset`, `readAsset`, `updateAsset`, and `deleteAsset` among others.

Generic names reduce the need for rework when the sample is extended or adapted into other domains. However, the generic names do not reinforce the asset class in use, and so the advanced multi-asset aviation contract sample adopts more explicit and accurate naming conventions. For example, there are three asset classes in the  

BLAH BLAH XX