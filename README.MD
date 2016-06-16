# Blockchain Samples

This project contains sample blockchain smart contracts and generic user interfaces for the [Hyperledger](https://github.com/hyperledger) [fabric](https://github.com/hyperledger/fabric).

These contracts are intended for exploration of the IBM Blockchain Bluemix environment and the Hyperledger fabric. They can be used to seed smart contract development with advanced features.

> **Note:** The release strategies for Hyperledger and for this project are evolving, so contract development is currently separated into folders that remain compatible with specific generations of the Hyperledger project on Bluemix.

---
## Hyperledger Supported Components

The Hyperledger Bluemix environment will evolve with the intent that releases will be tagged in the Hyperledger [fabric project](https://github.com/hyperledger/fabric) to indicate Bluemix compatibility levels. At the time of writing -- 16 June 2016 -- the Hyperledger fabric on Bluemix is compatible with the following folders in this project:

- [simple_contract_hyperledger](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract_hyperledger)
- [trade_lane_contract_hyperledger](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_hyperledger)
- [trade_lane_contract_hyperledger_e2e_tests](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_hyperledger_e2e_tests)
- [monitoring_ui](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/monitoring_ui)

Release 1.3 of this project was meant to guarantee compatibility when developing against Bluemix Hyperledger in the May-June 2016 time frame:

https://github.com/ibm-watson-iot/blockchain-samples/releases/tag/1.3

However, the folder strategy was adopted after that and so the folders named above will remain compatible with Bluemix for the forseeable future. As the strategy evolves, this document will capture changes.

---
## Upcoming Hyperledger Supported Components

The following contracts are compatible with the latest Hyperledger code base, expected on Bluemix in early summer:

- [`sdk`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/sdk)  

    - The sdk is under construction and its evolution is intended to include future packaging of boiler plate code for asset CRUD interfaces and related features such as asset state history, recent state activity, alerts and rules, and so on. This is intended to shrink the size of derived contracts while improving reliability and productivity.

  - [`iot_sample_contract`]((https://github.com/ibm-watson-iot/blockchain-samples/tree/master/sdk/iot_sample_contract)
    - Introduces saving of transaction timestamp and transaction uuid in asset state such that an audit can trace the state of an asset at any point in time to the specific block and transaction on the chain. Cleans up minor issues and bugs from the Trade Lane sample.
  - [`aviation_sample_contract`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/sdk/aviation_sample_contract)
    - Introduces complex object models with multiple asset classes managed by a single contract.
  - [`docs`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/sdk/docs)
    - A complete set of documentation that is derived from the document set in the Trade Lane sample but is intended to evolve over time to include new features and more articles.

- [`cashMachine`](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/cashMachine)
    - A sample contract that models an instant teller or other cash machine. Derived from the simple contract for hyperledger.

    ---
    ## OpenBlockChain Supported Components
    The following folders are compatible only with the earliest generation of Hyperledger on Bluemix, which was called Open Blockchain, or OBC:

    - [simple_contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_contract)
    - [trade_lane_contract](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract)
    - [trade_lane_contract_e2e_tests](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_e2e_tests)
    - [simple_ui](https://github.com/ibm-watson-iot/blockchain-samples/tree/master/simple_ui)
