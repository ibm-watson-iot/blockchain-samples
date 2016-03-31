import {fetchChainHeight, clearBlockchain} from './BlockActions'
import {clearTrackedAssets, setAssetInputErrMsg} from './AssetActions'

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
    console.log('Interval to clear: ' + intervalId)
    clearInterval(intervalId);

    //create a new interval and set it in the state
    intervalId = setInterval(() => {dispatch(fetchChainHeight(state.configuration.urlRestRoot))}, state.configuration.chainHeightPollingInterval);

    //clear assets being tracked
    dispatch(clearTrackedAssets())

    //clear error message
    dispatch(setAssetInputErrMsg(""))

    //set a new chain
    dispatch(clearBlockchain())

    //set the new timeout interval
    dispatch(setChainHeightPollingIntervalId(intervalId));
  }
}
