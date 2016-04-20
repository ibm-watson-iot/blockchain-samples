Blockchain Monitoring UI
=========================
A dynamically generated UI for IBM IoT Blockchain.

Requirements
-----------------
* Node installed locally on your machine.
* An OBC peer with an accessible REST Endpoint.
* A previously deployed contract on the OBC peer network from step 1.

Usage
-----------------
### Downloading
Download the project using either:
```
git clone git@github.com:ibm-watson-iot/blockchain-samples.git
```
in the terminal, or by downloading the zip of the project. The zip can be downloaded by clicking the `Download ZIP` button found on the top right.

### Installing
In the root directory of the project, run the following command:
```
npm install
```

### Running the UI Locally
#### Through the filesystem
Run the following command:
```
npm run build
```
This should generate a file called bundle.js in the public directory. Go to the public directory in the project and open up index.html in a browser.

#### Using the webpack-dev-server. Ideal for development.
This will load the project on a webpack-dev-server. Note this is not suitable for production.
```
npm run dev-server
```

Access the UI by going to:
```
http://localhost:8081/
```

If you run into an issue with the port already being used, set the `PORT` environment variable to the port you'd like to use. Note that hot reload is enabled for the webpack-dev-server, so saving changes to the source will be immediately reflected on the UI. No need to reload manually.

### Bluemix
```
TBD
```

### UI Functionality
The UI is a single page application that is divided into 3 columns. The columns are:

1. Chaincode Operations
2. Response Payloads
3. Blockchain

UI Configuration can be accessed by clicking the `CONFIGURATION` button on the top right of the App Bar. This will display a modal enabling the user to modify the UI configuration.

#### Chaincode Operations
This section of the UI is dynamically generated through a combination of JSON Schema and convention. The tabs are hardcoded in the `ChaincodeReducer.js` file as of now, but will be configurable through a separate configuration file in the near future. Each tab represents a subset of contract functionality. The functions themselves can be selected by clicking the dropdown in the tab's contents and selecting the appropriate function. For instance, assuming that the simple contract is the configured contract, the create tab will have a drop down of one function, `createAsset`. This maps to a function defined in the JSON schema, called `createAsset`. The UI knows to put `createAsset` under the `create` tab because it matches the tab's name as a substring of the function. To see an example of a tab with multiple contract functions, the `read` tab contains three functions, each of which were defined in the JSON schema. Each tab also has a corresponding `type`, which defines whether it should use the OBC invoke or query endpoint.

The arguments form is generated when the user selects a particular function from the dropdown. This creates a form with input fields for every expected argument defined in the JSON schema. For the simple contract, we have the fields `assetID`, `carrier`, `location`, `temperature` and `timestamp`. There are a few things to note:
- `assetID` has an asterisk. This means that this field is required. This is defined in the JSON schema.
- location is a nested object. In other words, it has its own properties, which in turn become fields that we can input data into.
- validation is defined in the JSON schema and is reflected in the form. For instance, if the form is submitted without an `assetID`, the UI will display a message asking for the user to input a value. If a non-numeric value is input into the latitude or longitude fields, the UI will indicate that the value is not of a type(s) number.

If all values are valid and submit is clicked, then the UI will create a valid OBC REST payload with the user input as the arguments and send it along with a request to the configured OBC peer. It then waits for a response from the peer. The response will be output to the section in the second column, which is the Request Payload section.

#### Response Payload
This section of the UI is responsible for displaying the response from the OBC peer. It does so by recursively traversing the payload and outputting the response to the card. Note that it is possible to submit multiple requests from any combination of tabs; the UI will just generate additional cards to display the payload. If a REST request is made with the exact same function and arguments, it will not be shown as an additional card, because it is a duplicate.

When a request payload card initially appears, it will be in the collapsed state. Click the expand button on the right side of the payload's card header to view the contents of the card. You'll see a `Poll for changes` toggle along with a basic formatted string representation of the response payload. When the toggle is on, the UI will actively check for changes to a particular query every time the blockchain height changes. In the case of the simple contract, it is useful for monitoring a particular asset for changes. This toggle is set to off by default.

The user may close an individual card by clicking the `x` button that is on the left of the header. Notice that the button changes colors from gray to orange when hovering over the particular header. This is to facilitate figuring out which card will be closed after clicking the button. If there are many payloads being displayed, the user may also click the clear button on the Request Payload header to remove all payloads from the display.

#### Blockchain
This is the right most column on the UI and shows the current state of the blockchain. The amount of blocks that are displayed is configurable through the configuration modal. To expand a block, click on the expander on the right side of the block header. The contents of the block show what transactions are part of the block and the details of each transaction.

It is important to note that any transactions that occur against the blockchain will appear within blocks on the blockchain. These include invalid transactions and transactions against other contracts. A common error is to have one user use a particular contract to do an update and another user try to read that change. If the other user does not see the change, there is a chance that they have configured the wrong contract.

#### Configuration Modal
This modal can be accessed by clicking the `CONFIGURATION` button on the right side of the main App Bar. This modal will allow you to configure the following settings:

**API host and port**: This should be the in the form of protocol://hostname:port. So for example:
```
http://localhost:3000
```
**Chaincode ID**: The ID of the contract that the UI should use to perform operations. For example:
```
ff89038cb1db8fcddff9f3c786bba06dc1af9afb2616d8bcb851ac50db383be02e25391d979c5eaa499abf2845df270089eb9ac982cf3dec880d24ff70cf95d9
```
These IDs are 128 character hashes that should not trailing nor leading spaces. There should also be no spaces within the ID itself. If copying and pasting this ID from somewhere else, ensure that the Chaincode ID conforms to these parameters.

**Secure Context**: The secure context for doing POST requests to the API endpoint, if necessary.

**Number of blocks to display**: The number of blocks to display from the blockchain. For instance, if this value was `10`, the most recent 10 blocks would show up on the blockchain display.
