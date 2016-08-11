

#Purpose

The node.js scripts shared here are **purely meant for contract testing purposes**, to generate sample data and test it against the sample contract. The actual implementation should be routed through the IoTP. Once the proxy is configured, the data will flow through the proxy to the smart contract, making sendingIoT data to the blockchain a hassle-free, seamless flow.

###deviceGenerator.js

This node.js code reads the _configGenerator.json_ to understand the requirements of data generation (date ranges for dates, data generation flag - could be eliminated and number of new devices to be added). It then creates new device data and appends it to the _devices.json_ file. This could be changed to create an entirely new set of data. An example of an updated file can be seen in here : *example devices-after generator run.json*


###deviceContractCall.js
This node.js code reads the _configContract.json_ to understand the contract details (contract id, host, port etc.) and the nature of the call to be made (create/ read / update etc). It then reads device device data from the _devices.json_ file and makes calls to the contract, one record at a time. _configContract.json_ has an example for a local environment and *configContract_bluemix example.json* is an example of how it would look like if calling the contract registered in Bluemix.
For example, to send new device data to the contract, set "method" as "invoke" and "functionName" as "createAsset". In order to test an update scenario, just update the data in devices.json (except for assetid and name) and change function name as "updateAsset". To read, change "method" to "query" and "functionName" to "readAsset" and so on....
Note: readAsset function in our sample contract only cares about being able to find the assetid and name in the input string, so even though the script sends the entire data, it shouldn't impact the read function.


Note: This is sample wip code. Please use and modify your copy as per your discretion.