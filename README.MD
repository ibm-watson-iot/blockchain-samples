# Blockchain Samples

This project contains samples of blockchain smart contracts and blockchain applications for the [Hyperledger](https://github.com/hyperledger) [fabric](https://github.com/hyperledger/fabric). 

These are intended to:

- provide sample applications and smart contracts to demonstrate how a Hyperledger blockchain and the Watson IoT platform can work together in a solution
- demonstrate advanced smart contract features that can be built on the Hyperledger fabric
- demonstrate patterns that can make complex IoT contracts easier to develop and maintain
- seed new smart contract development with advanced IoT features and integration into the Watson IoT Platform

Notes:

- Development of the Hyperledger fabric and related projects has been moved to [Gerritt](https://gerrit.hyperledger.org/r/#/admin/projects/). However, the [original location](https://github.com/hyperledger/fabric/) remains as a convenient mirror. 
- The Hyperledger Fabric also has a [web site for documentation](http://hyperledger-fabric.readthedocs.io/en/latest/), which is worth exploring, *especially for those who would like to contribute directly to the Hyperledger development effort.*
- Hyperledger on Bluemix is updated periodically, and for compatibility reasons, the blockchain-samples project has multiple release levels to line up with fabric versions available in IBM Bluemix. All contracts at head on the master branch are compatible with [Hyperledger / IBM Blockchain on Bluemix](http://www.ibm.com/blockchain/bluemix.html) as of 20 September 2016.

---

## Smart Contracts in this Repository

- [`simple_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/basic/simple_contract)
    - A basic IoT sample contract that derives from early work on the trade lane sample. Uses structs to manipulate internal asset state and stores asset state in world state as JSON objects. Adheres to the schema-based protocol for integration with the Watson IoT Platform, but does not include the raw schema or processing from trade lane. Suitable for exploring Hyperledger with a small code base.

- [`iot_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/advanced/iot_sample_contract)
    - An advanced IoT sample contract as the logical successor to the early trade lane work. Exemplifies the _partial state as even pattern_ for IoT contract development. Uses maps instead of structs for internal state with a set of utilities that enable easy manipulation of asset state. Saves transaction timestamp and uuid in asset state for audit purposes. Introduces asset state history, deletion of properties from state, recent state change tracking and query, rules engine, and alerts. Note that the simple trade lane properties used in the original Trade Lane sample contract still exist in this contract and can be used with applications such as the Digital Experience, etc.
    
- [`aviation_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/aviation_sample_contract)
    - Introduces complex object models with multiple asset classes managed by a single contract. Restructures asset manipulation such that new assets delegate default processing. Introduces event handling with asset state manipulation. Models a relatively light weight aviation scenario with tracking of assembly (as in life limited parts) life cycles.

- [`cashMachine`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/cashMachine)
    - A sample contract that models an instant teller or other cash machine. Derived from the simple contract for hyperledger.

- [`building`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/building)
    - A sample contract that models an instant teller or other cash machine. Derived from the simple contract for hyperledger.

---

## Documentation

See the [`docs`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/docs) folder for articles on patterns and customization in these samples. Detailed documentation exists for mode major modules.

---

## Caveats

These samples are developed with minimal resources and so the packaging of features has not yet begun. If you would like to collaborate on exploration of these concepts, please contract [Kim Letkeman](mailto:kletkema@ca.ibm.com).

Enjoy!
