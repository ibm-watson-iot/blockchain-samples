Blockchain Generic UI
=========================
A generic UI for IBM IoT Blockchain.

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

#### JSON Schema UI Generation
- Contract function names need to start with create, read, update or destroy. This allows the UI to map function calls to their appropriate OBC REST endpoints. For instance, `createAsset` is a valid function name, but `makeAsset` would not work with the UI.
- Each tab has its own form. The state of each form is managed in the `chaincodeOpsForm` store. Data is persisted even when changing the selected function from the select dropdown. The form data is stored under the `args` property of each function as an array.
- When a form is submitted for a read function, the response payload will be displayed in a section below the form. Otherwise, a snackbar will be displayed indicated a request was sent. This is because OBC read endpoints are synchronous while invoke endpoints are asynchronous.
- Click on the switch to enable polling for a particular request.
- Click on the x button to hide a response panel
- Use local storage to store configuration
- Two forms, one for the function dropdown and another for the args.

#### Blockchain Viewer
- The Blockchain display on the right half of the UI shows all validated transactions on the Blockchain. Click on each header to view the details of the transactions for each block.
- These transactions may have been done against a different contract than yours. In which case, not all assets that appear on the Blockchain will be accessible for Asset Tracking. In order to see those assets, you must switch the contract that you are querying against. This can be found in the configuration modal.

#### Configuration Modal
- The configuration can be accessed by clicking on the `CONFIGURATION` button located near the top right of the screen. This will bring you to the configuration form, which allows you to configure the UI. The fields that can be modified are:

  * API Host and Port: The host URL and port number written as a web URL. For instance: http://localhost:3000
  * Chaincode ID: The ID of the contract that the UI should use to perform operations.
  * Secure Context: Required if the OBC Peer has security enabled.
  * Number of blocks to display: The number of blocks to display on the right side.
