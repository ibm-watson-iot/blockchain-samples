# Blockchain Samples

This project contains sample blockchain smart contracts and generic user interfaces for the [Hyperledger](https://github.com/hyperledger) [fabric](https://github.com/hyperledger/fabric). 

These contracts are intended to:

- provide sample applications and smart contracts to demonstrate how a Hyperledger blockchain and the Watson IoT platform can work together in a solution
- demonstrate advanced smart contract features that can be built on the Hyperledger fabric
- seed new smart contract development with advanced IoT features and integration into the Watson IoT Platform

Notes:

- Development of the Hyperledger fabric and related projects has been moved to [Gerritt](https://gerrit.hyperledger.org/r/#/admin/projects/). However, the [original location](https://github.com/hyperledger/fabric/) remains as a convenient mirror. 
- The Hyperledger Fabric also has a [web site for documentation](http://hyperledger-fabric.readthedocs.io/en/latest/), which is worth exploring, *especially for those who would like to contribute directly to the Hyperledger development effort.*
- Hyperledger on Bluemix is updated periodically, and for compatibility reasons, the blockchain-sample project maintains multiple release levels to line up to bluemix versions. All contracts at head on the master branch are compatible with Hyperledger on Bluemix as of 20 September 2016.

---

## Smart Contracts in this Repository

  - [`simple_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/basic/simple_contract)

  - [`iot_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/advanced/iot_sample_contract)
    - Introduces saving of transaction timestamp and transaction uuid in asset state such that an audit can trace the state of an asset at any point in time to the specific block and transaction on the chain. Cleans up minor issues and bugs from the Trade Lane sample.
  - [`aviation_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/aviation_sample_contract)
    - Introduces complex object models with multiple asset classes managed by a single contract.

- [`cashMachine`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/cashMachine)
    - A sample contract that models an instant teller or other cash machine. Derived from the simple contract for hyperledger.

- [`building`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/contracts/industry/building)
    - A sample contract that models an instant teller or other cash machine. Derived from the simple contract for hyperledger.

- [`docs`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/docs)
    - Documentation 

    ---
    ## OpenBlockChain Supported Components
    The following folders are compatible only with the earliest generation of Hyperledger on Bluemix, which was called Open Blockchain, or OBC:

    - [simple_contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract)
    - [trade_lane_contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract)
    - [trade_lane_contract_e2e_tests](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_e2e_tests)
    - [simple_ui](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_ui)
