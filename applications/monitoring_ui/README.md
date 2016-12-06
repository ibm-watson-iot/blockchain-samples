Blockchain Monitoring UI
=========================
The Blockchain Monitoring UI is a dynamically generated user interface for IBM Watson IoT Platform blockchain integration. Use the Monitoring UI to perform actions on the blockchain, see the results of those actions, and monitor the state of your assets in the blockchain ledger.

## Requirements
* Node installed locally on your machine.  
You can install it from https://nodejs.org/.  
* An IBM Blockchain or Hyperledger peer with an accessible REST Endpoint.  
For more information, see: https://console.ng.bluemix.net/docs/services/IoT/blockchain/dev_blockchain.html.
* A deployed contract on the blockchain peer network.

## Downloading

Use `git clone` in the console to clone the following project:  
https://github.com/ibm-watson-iot/blockchain-samples  

The monitoring UI component is in the monitoring_ui folder. You can also download a compressed file of the project by clicking **Download ZIP** from the project page.  

## Installing
To install the Monitoring UI on your local workstation, in the root directory of the project, run the following command:
```
npm install
```

## Running
You can run the monitoring UI directly from the file system or by running the webpack-dev-server.

Method	| Command	|Comment
--- | --- | ---
Filesystem | `npm run build` | The build command generates the bundle.js file in the public directory. </br>To access the Monitoring UI, go to the `monitoring_ui/public` directory and open the *index.html* file in a browser.
webpack-dev-server | `npm run dev-server` | This method is ideal for a development environment but not suitable for a production environment. </br>To access the Monitoring UI, open the following URL in a browser: `http://localhost:8081/` </br>**Note:** If you run into an issue with the port already being used, set the `PORT` environment variable to the port you'd like to use. Note that hot reload is enabled for the webpack-dev-server. Changes that you save to the source are immediately reflected in the Monitoring UI. There is no need to manually reload.

## Configuration
Before you can access the blockchain information with the Monitoring UI, you must point it to a blockchain peer server and provide a contract ID to monitor. Access the configuration by clicking **CONFIGURATION**.  
You can configure the following parameters:

Parameter	|Value	|Comment
--- | --- | ---
API Host and Port	| http://peer_URL:port	| The host and port for the IBM Blockchain REST API prepended with `http://`.
Chaincode ID	| The contract ID that was returned when you registered the contract.	| The contract ID is a 128-character alphanumeric hash that corresponds to the Contract ID entry. </br> **Important:** As you cut-and-paste the contract ID, make sure that no spaces are included in the ID. If the ID is incorrectly entered, the UI will display the blockchain ledger entries, but the asset search function will not work.
Secure Context|Your fabric user	| This is required for connecting to IBM Blockchain instances on Bluemix. </br>**Important:** For secureContext use the user name that was used to configure the fabric.
Number of blocks to display	| A positive integer. Default: 10	| The number of blockchain blocks to display.


## Exploring the Monitoring UI Features
Use the Monitoring UI single page application to perform actions on the blockchain, see the results of those actions, and monitor the state of the blockchain ledger.   

The user interface is divided into three columns.  
1. Chaincode Operations
2. Response Payloads
3. Blockchain

### The Chaincode Operations column
The first section of the Monitoring UI is dynamically generated through a combination of JSON Schema and convention. The tabs each represent a subset of the available contract functions and are hardcoded in the `ChaincodeReducer.js` file.     

The contract functions can be selected from the menu in each tab. The functions and their related input fields are defined in the JSON schema.

For example, if we are connected to the IBM sample conctract on the blockchain fabric the **Create** tab includes just one function: `createAsset`. This function maps to the `createAsset` function that is defined in the JSON schema. The UI knows to put `createAsset` under the `create` tab because it matches the tab's name as a substring of the function. The **Read** tab, in contrast, contains three functions each of which are defined in the JSON schema. Each tab also has a corresponding `type`, which controls the use of Hyperledger invoke or query endpoint.

The *arguments* form is generated when you select a particular function. The form creates input fields for the arguments that are defined in the JSON schema. The basic contract includes the following fields: `assetID`, `carrier`, `location`, `temperature`, and `timestamp`.  

**Note:**
- Fields such as `assetID` that are denoted with asterisk are required. Required fields are defined in the JSON schema.
- Location is a nested object with its own properties that in turn are exposed as data fields.
- Validation is defined in the JSON schema and is reflected in the form.  
For example, if you submit the form without an entry for the required `assetID` field, you are prompted to enter a value. If you enter a non-numeric value in the `latitude` or `longitude` fields, you will also be prompted.

When you submit a form the Monitoring UI creates a valid blockchain REST payload with the field input as arguments. The payload and a request is sent to the configured blockchain peer. The Monitoring UI then waits for a response from the peer. The response is displayed in the Request Payload column.

### The Response Payload column
The second column displays the response from the blockchain peer by recursively traversing the payload and writing the responses to the card. If you submit multiple requests from a combination of tabs the Monitoring UI generates cards as needed to display the payload. **Note:** Duplicate REST request with the exact same function and arguments will not create extra cards.

Request payload cards are displayed in the collapsed state. You must expand the cards to view the contents of the card.

Close individual cards by clicking **x** next to the card header.

Click **Clear** on the Request Payload header to remove all payloads from the display.

**Tip:** Enable the **Poll for changes** toggle to have the Monitoring UI actively check for changes to a particular query every time the blockchain height changes. For the basic contract, use this feature to monitoring a particular asset for changes.


### The Blockchain column
The third column shows the current state of the blockchain.

To expand a block, click the expander. The contents of the block show the transactions in the block and the details for each transaction.

**Important:** Any transactions that occur against a specific blockchain will appear within blocks on the blockchain. These include invalid transactions as well as transactions against other contracts. To see a change on a specific contract, the Monitoring UI must be configured to connect to that contract.
