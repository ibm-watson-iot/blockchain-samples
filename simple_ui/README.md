Blockchain Simple UI
=========================
Simple UI example for IBM IoT Blockchain. Assets can be tracked by typing the asset ID into the field in the Assets section and clicking submit.

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

#### Asset Tracking
- Track an asset by typing in the Asset ID in the `Asset ID` field on the left panel and click submit.
- If you typed in an asset that currently exists in the current contract database, you will see a header appear. Click the expand arrow on the right to view the details of this asset.
- If anyone makes changes to this asset, you will see live updates appear in the details section. There is no need to refresh.
- It is possible to track multiple assets by typing the name of another Asset ID in the `Asset ID` field. This will make another header appear which functions the same as the one before it.
- If you would like to stop monitoring an asset, click on the `X` button located to the left of the header.

#### Blockchain Viewer
- The Blockchain display on the right half of the UI shows all transactions on the Blockchain. Click on each header to view the details of the transactions for each block.
- These transactions may have been done against a different contract than yours. In which case, not all assets that appear on the Blockchain will be accessible for Asset Tracking. In order to see those assets, you must switch the contract that you are querying against. This can be found in the configuration modal.

#### Configuration Modal
- The configuration can be accessed by clicking on the `CONFIGURATION` button located near the top right of the screen. This will bring you to the configuration form, which allows you to configure the UI.

The initial configuration is defined in the configuration reducer:
```javascript
//src/reducers/ConfigurationReducer.js
  state={
    urlRestRoot: "http://localhost:3000",
    //This is the root address of the OBC REST API
    chaincodeId: "7cdd53526ed31f7be5249bfa42c4c73728edebddf91bd29720a289105dafbf1fd8c94306ba128800fc1c2bfbee618ce85717d35f88bb7b481ca3d3ada70d78fd",
    //the ID of the chain height polling. This will be populated at runtime.
    chainHeightPollingIntervalId: -1,
    //the chain height polling interval, in milliseconds
    chainHeightPollingInterval: 2000,
    secureContext: "user_context",
    showDialog: false
  }
}
```
These defaults are for the public Bluemix environment and most are customizable from within the UI. The settings you may configure from the UI are:

API Host and Port: The host and port for the OBC REST API. This string is used directly as the `urlRestRoot` property.
Chaincode ID: The ID for the chaincode we are working with. This string is used directly as the `chaincodeId` property.
Secure Context: This is required for connecting to OBC instances on Bluemix. This string is used directly as the `secureContext` property.

Development Guide
-------------------
The UI is written using react and redux using ES2015. All required babel presets and webpack loaders are already included in this project. If you are not familiar with react, redux or webpack, I recommend spending some time getting familiar with the technologies:
```
React: https://facebook.github.io/react/
Redux: https://github.com/reactjs/redux
Webpack: https://webpack.github.io/
```

In order to inspect assets through the UI, you will need to run invoke commands against the OBC peer through the CLI or RESTful interface. The easiest way to do this is through a utility that can send POST and GET requests, such as postman. Once that is done, you can inspect asset state by inputting the asset ID.

### Program Flow
There are two major containers for this application, the blockchain container and the asset container. The blockchain container can be found here:
```
src/containers/Blockchain.jsx
```
The asset container can be found here:
```
src/containers/AssetContainer.jsx
```

### Blockchain Container
This is where the blockchain is rendered and also where the polling logic takes place. Since OBC does not support event broadcasting, we are required to poll  the blockchain if we want to detect if changes have occurred. We do this by setting an interval:
```javascript
//src/containers/Blockchain.jsx

componentDidMount(){
  this.props.fetchChainHeight(this.props.urlRestRoot);

  let intervalId = setInterval(() => {this.props.fetchChainHeight(this.props.urlRestRoot)}, 2000);
  this.props.setChainHeightPollingIntervalId(intervalId);
}
```

fetchChainHeight will pull the chain height from the OBC REST API, which is then diffed with the current chain height. We know that a new transaction has occurred on the blockchain when the fetched chain height is different than the stored one:

```javascript
//src/actions/BlockActions.js
export function fetchChainHeight(urlRestRoot){

  return function (dispatch, getState){

    return fetch(urlRestRoot + '/chain')
    .then(response => response.json())
    .then((json) =>{
      //console.log(json);

      //check if the chain height is any different

      let state = getState();

      //call any other functions that need to be invoked on change in blockchain height
      if(state.blockchain.length !== json.height){
        dispatch(recChainHeight(json.height))

        dispatch(fetchAllAssets());
      }
    })
  }
}
```

Once we are sure the height is different, we can proceed with dispatching any actions that depend on new transactions. The main benefit with this approach is that it reduces the number of polling routines that we require to track an asset. So rather than having each asset poll the API for it's own state, we only do the check if the blockchain height has changed. At this time, the payload received from the response is basically a byte stream. It isn't very standardized, which means we don't really have an idea of what the change was, but we do know a change happened. Thus, we have to run all dispatches that are dependent on new transactions coming in.

### Asset Container
The asset container has one text field that a user can use to input asset IDs. When as asset ID is entered, a material-ui card will be appended to the asset list in the container. Since we are fetching all assets on blockchain height change, the information contained in these asset cards will be updated at a regular interval. You can try this out by opening one of the asset cards and then updating an asset. You will see that the information in the card changes soon after.
