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
import {fetchChainHeight, clearBlockchain} from './BlockActions'
import {fetchCcSchema, clearResponsePayloads} from './ChaincodeActions'

export const CONFIG_UPDATE_URL_REST_ROOT = 'CONFIG_UPDATE_URL_REST_ROOT'
export const setUrlRestRoot = (url) =>{
  return{
    type: CONFIG_UPDATE_URL_REST_ROOT,
    url
  }
}

export const SET_CONFIG_IOT_CONNECTION = 'SET_CONFIG_IOT_CONNECTION'
export const setConfigIotConnection = (iotDeviceClient) => {
  return{
    type: SET_CONFIG_IOT_CONNECTION,
    iotDeviceClient
  }
}

export const SET_CHAIN_HEIGHT_POLLING_INTERVAL_ID = 'SET_CHAIN_HEIGHT_POLLING_INTERVAL_ID'
export const setChainHeightPollingIntervalId = (intervalId) => {
  return{
    type: SET_CHAIN_HEIGHT_POLLING_INTERVAL_ID,
    intervalId
  }
}

export const SET_OBC_CONFIGURATION = 'SET_OBC_CONFIGURATION'
export const setObcConfiguration = (obcConfigObj) =>{

  //we first need to clear the interval. Subsequently, we need to start
  //a new interval with the new OBC root and update the interval id.
  return{
    type: SET_OBC_CONFIGURATION,
    obcConfigObj
  }
}

export const SET_CONFIG_DIALOG_DISPLAY = 'SET_CONFIG_DIALOG_DISPLAY'
export const setConfigDialogDisplay = (showDialog) =>{
  return{
    type: SET_CONFIG_DIALOG_DISPLAY,
    showDialog
  }
}

export function setConfiguration(obcConfigObj){

  return function(dispatch, getState){
    //let state = getState();

    //update any OBC related configuration
    dispatch(setObcConfiguration(obcConfigObj))

    //get the current state
    let state = getState();

    //clear the existing interval timer
    let intervalId = state.configuration.chainHeightPollingIntervalId
    clearInterval(intervalId);

    //create a new interval and set it in the state
    intervalId = setInterval(() => {dispatch(fetchChainHeight(state.configuration.urlRestRoot))}, state.configuration.chainHeightPollingInterval);

    //set a new chain
    dispatch(clearBlockchain())

    //clear the response payloads when changing configuration
    dispatch(clearResponsePayloads())

    //set the new timeout interval
    dispatch(setChainHeightPollingIntervalId(intervalId));

    //reload form schema
    dispatch(fetchCcSchema());
  }
}
