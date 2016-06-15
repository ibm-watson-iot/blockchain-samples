/*****************************************************************************
Copyright (c) 2016 IBM Corporation and other Contributors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Alex Nguyen - Initial Contribution
*****************************************************************************/
import {
  CONFIG_UPDATE_URL_REST_ROOT, SET_CONFIG_IOT_CONNECTION, SET_CHAIN_HEIGHT_POLLING_INTERVAL_ID, SET_OBC_CONFIGURATION,
  SET_CONFIG_DIALOG_DISPLAY
} from '../actions/ConfigurationActions.js'

//set default configuration. To keep things simple, we don't have any UI elements
//to configure these properties, so they must be populated ahead of time.
export const configuration = (state={
  urlRestRoot: "http://localhost:3000",
   //urlRestRoot: "http://169.44.63.199:37687",
  //chaincodeId: "abf072028033b86aa9d61127c9ac9f0f407a24ce5b464f3afb6e8474169df95e1c1e40d31051553430eca22d10fe8b7083a518c77b80c94d679dba4c6858a90b",
  chaincodeId: "ff89038cb1db8fcddff9f3c786bba06dc1af9afb2616d8bcb851ac50db383be02e25391d979c5eaa499abf2845df270089eb9ac982cf3dec880d24ff70cf95d9",
  //the ID of the chain height polling. This will be populated at runtime.
  chainHeightPollingIntervalId: -1,
  //the chain height polling interval, in milliseconds
  chainHeightPollingInterval: 2000,
  secureContext: "user_context",
  //this is a UI specific toggle. This governs whether or not the configuration modal is showing or not.
  showDialog: false,
  //this is the number of blocks to show on the page at a time.
  blocksPerPage: 10
}, action) => {
  /**
    These state configuration actions are implemented, but we don't use them in the UI. Feel free to wire them to a UI element if needed.
  **/
  switch (action.type){
    case CONFIG_UPDATE_URL_REST_ROOT:
      return Object.assign({}, state, {
        urlRestRoot: action.url
      })
    case SET_CHAIN_HEIGHT_POLLING_INTERVAL_ID: {
      return Object.assign({}, state, {
        chainHeightPollingIntervalId: action.intervalId
      })
    }
    /**
      This action is called from the submit button on the form. It is used to transfer over values from the obcConfiguration model to the configuration used by the ui.
      This allows us to make changes to the form without affecting the queries going on in the background.
    **/
    case SET_OBC_CONFIGURATION:
      return Object.assign({}, state, {
        //set the appropriate properties
        urlRestRoot: action.obcConfigObj.urlRestRoot,
        chaincodeId: action.obcConfigObj.chaincodeId,
        secureContext: action.obcConfigObj.secureContext,
        blocksPerPage: action.obcConfigObj.blocksPerPage
      })

    /**
      Strictly a UI control. This determines whether or not the configuration ui dialog should display or not.
    **/
    case SET_CONFIG_DIALOG_DISPLAY:
      return Object.assign({}, state, {
        showDialog: action.showDialog
      })

    default:
      return state;
  }
}
