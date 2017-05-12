# Blockchain Samples

This project contains samples of blockchain smart contracts and blockchain applications for the [Hyperledger](https://github.com/hyperledger) [fabric](https://github.com/hyperledger/fabric). 

Purpose:

- provide sample applications and smart contracts to demonstrate how a Hyperledger blockchain and the Watson IoT platform can work together in a solution
- demonstrate advanced smart contract features that can be built on the Hyperledger fabric
- demonstrate patterns that can make complex IoT contracts easier to develop and maintain
- seed new smart contract development with advanced IoT features and integration into the Watson IoT Platform
 
Documentation:

- [Introduction to Best Practices and Patterns for IoT Contract Development on Hyperledger](https://github.com/ibm-watson-iot/blockchain-samples/blob/master/docs/HyperledgerContractsIntroBestPracticesPatterns.md)
- [Customizing or Extending these Sample Contracts](https://github.com/ibm-watson-iot/blockchain-samples/blob/master/docs/CustomizingTheSampleContract.md)
- [Index to the Docs](https://github.com/ibm-watson-iot/blockchain-samples/blob/master/docs/README.md)

Notes:

- Development of the Hyperledger fabric and related projects has been moved to [Gerritt](https://gerrit.hyperledger.org/r/#/admin/projects/). However, the [original location](https://github.com/hyperledger/fabric/) remains as a convenient mirror. 
- The Hyperledger Fabric also has a [web site for documentation](http://hyperledger-fabric.readthedocs.io/en/latest/), which is worth exploring, *especially for those who would like to contribute directly to the Hyperledger development effort.*
- Hyperledger on Bluemix is updated periodically, and for compatibility reasons the blockchain-samples project has multiple release levels to line up with fabric versions available in IBM Bluemix. 
  - All contracts that have `.0.6` appended to their folder names are compatible with [Hyperledger v0.6](https://github.com/hyperledger/fabric/tree/v0.6) and new fabrics created on [IBM Bluemix](https://console.ng.bluemix.net/) as of 10 November 2016.
  - All contracts without this suffix are compatible with [Hyperledger v0.5 Developer Preview](https://github.com/hyperledger-archives/fabric/tree/v0.5-developer-preview) and preexisting IBM Bluemix fabrics.

---

## Smart Contracts in this Repository

- [`simple_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/basic/simple_contract)

    The Basic contract is a sample hyperledger blockchain contract that is provided by IBM to help you to get started with blockchain development and integration on the IBM Watson IoT Platform. You can use the Basic contract sample to create a blockchain contract that tracks and stores asset data from a device that is connected to your Watson IoT Platform organization.

- [`iot_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/advanced/iot_sample_contract)
    
    This sample contract implements a simple Trade Lane scenario, moving *assets* from one place to another. It consists of several modules in `package main` that are all small enough to avoid creating separate packages in this first version. These add features that can be used without changing the code. This sample is used to explore features and patterns that are of interest to smart contract writers in the IoT domain. These are:

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

- [PingPong Contract]

    This contract has been derived from the IoT Sample contract. Please read the introductory articles in the [docs folder](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/docs) if you are not familiar with how these contracts function.

    This contract adds the ability to send outgoing events from an advanced IoT contract. Three events are registered, as clipped from the code:

    ``` go
    // EVTINVOKEERR is sent out whenever there is an error in the chaincode
    const EVTINVOKEERR = "EVTINVOKEERR"

    // EVTPONG is sent out whenever a PING assetID is received
    const EVTPONG = "EVTPONG"

    // EVTPING is sent out whenever a PONG assetID is received
    const EVTPING = "EVTPING"
    ```

    See the [event listener application](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/applications/event_listener) README and code to understand how the client registers interest in specific events and then catches them in a gRPB stream (modeled in Go as channels).

- [`aviation_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/aviation_sample_contract)

    The aviation scenario is as follows:

        - Three primary asset classes:

            - Airline -- owns zero to many aircraft
            - Aircraft -- encompasses zero ro many assemblies
            - Assembly -- landing gear, wing, etc. These are life limited parts.

        - The CRUD features for each asset are abstracted in the module [assetCommon.go](./assetCommon.go), which makes heavy use of abstracted services in [crudCommon.go](./crudCommon.go)  
        - each asset follows the *partial state as event* pattern, where the asset's writable properties make up its primary `event` to be passed to create and update
        - additional events exist for:

            - flights that record a takeoff and landing sequence, which is known as a cycle
            - inspections against an assembly
            - analyticAdjustment events record calculated changes in wear and tear based on such analytics as weather patterns and runway conditions
            - maintenance events record the mounting and unmounting of assemblies with a full state machine

        - rules exist to track cycles and hard landings

            - ACHECK rule compares adjusted (by analytics) cycle counters to a dynamically configurable threshold, raising an alert as necessary, typically short time
            - BCHECK rule compares adjusted (by analytics) cycle counters to a dynamically configurable threshold, raising an alert as necessary, typically long time
            - HARDLANDING rule raises an alert when a hard landing is known to have occured
            
        - inspection events clear these alerts, note that bcheck clears both acheck and bcheck alerts

        > Note that the usual common properties such as geolocation, extension, etc. are available in the `common` subsection of asset event and state.

        Physical changes from the Generic IoT Contract include:

        - Three assets tracked with full CRUD APIs: airline, aircraft, assembly
        - Four event types handled with event APIs: flight, maintenance, inspection, analyticAdjustment
        - 2-way inverted index maintains 1 to many relationship from aircraft to assemblies
        - filters allow sophisticated queries and offer lightweight relationships between asset classes

            - a filter is a match type (all, any, none) and an array of k:v pairs with qualified property names and values

        - contractConfig module supports static and dynamic configuration of contract
        - new common layer for quick addition of a new asset class
        - new common layer to support crud operations
        - rules for acheck and bcheck (short and long term inspection cycles) and hard landing alerts 

- [`cashMachine`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/cashMachine)

- [`building`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/building)

- [`Parking Meter Demo`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/parkingmeter)

    The Parking meter is FRDM-K64 and Cordio Boards combined. The FRDM board acts as the meter and the cordio acts as the BLE Beacon. The beacon transmits short urls in Eddystone standard and this points to the bluemix page https://mbed-parkingmeter.mybluemix.net. This page supports pakring meters PKM-001 through PKM-010. The code for the UI is inside mbedParkingMeterUI.

    The UI makes a call to the IBM commerce system to initiate payment and the chaincode (code availalbe in mbedParkingMeter / mbedParkingMeter.0.6) for recording parking meter usage data.

    It makes http calls to the node-red flow - NodeFlow.json -for setting the Parking meter to free (beacon stops emitting, UI says 'Free Parking') or paid (beacon starts emitting, UI says 'Paid parking') and also initiating the countdown once payment is made.

- [`Carbon Trading Demo`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/carbon_trading)

- [`Logistics Split`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/LogisticsSplit.0.6)
