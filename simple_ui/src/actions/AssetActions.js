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
/**
Adds an asset to the asset.assets array in state. This array contains the
assets we are tracking.
**/
export const ADD_ASSET_TO_TRACK = "ADD_ASSET_TO_TRACK"
export const addAssetToTrack = (assetInfo) => {
  /*
  assetInfo: json.OK
  */
  return{
    type: ADD_ASSET_TO_TRACK,
    assetInfo
  }
}

/**
Updates an existing asset. For example, if the temperature is changed and we
are already tracking the asset, we would use this action.
**/
export const UPDATE_ASSET_TO_TRACK = "UPDATE_ASSET_TO_TRACK"
export const updateAssetToTrack = (assetData) => {
  /*
  assetData: json.OK
  */
  return{
    type: UPDATE_ASSET_TO_TRACK,
    assetData
  }
}

export const REMOVE_ASSET_FROM_TRACKING = "REMOVE_ASSET_FROM_TRACKING"
export const removeAssetFromTracking = (assetId) => {

  return{
    type: REMOVE_ASSET_FROM_TRACKING,
    assetId
  }
}

export const SET_ASSET_OBJECT_MODEL = "SET_ASSET_OBJECT_MODEL"
export const setAssetObjectModel = (objectModel) => {
  return{
    type: SET_ASSET_OBJECT_MODEL,
    objectModel
  }
}

export const SET_ASSET_ID_INPUT = "SET_ASSET_INPUT"
export const setAssetIdInput = (assetId) => {
  return{
    type: SET_ASSET_ID_INPUT,
    assetId
  }
}

/**
Toggle the remove button based on the index in the assets array
**/
export const TOGGLE_SHOW_REMOVE_BTN = "TOGGLE_SHOW_REMOVE_BTN"
export const toggleShowRemoveBtn = (index) => {
  return{
    type: TOGGLE_SHOW_REMOVE_BTN,
    index
  }
}

export const HIDE_REMOVE_BTN = "HIDE_REMOVE_BTN"
export const hideRemoveBtn = (index) => {
  return{
    type: HIDE_REMOVE_BTN,
    index
  }
}

export const SHOW_REMOVE_BTN = "SHOW_REMOVE_BTN"
export const showRemoveBtn = (index) => {
  return{
    type: SHOW_REMOVE_BTN,
    index
  }
}

export const CLEAR_TRACKED_ASSETS = "CLEAR_TRACKED_ASSETS"
export const clearTrackedAssets = () => {
  return{
    type: CLEAR_TRACKED_ASSETS
  }
}

export const SET_ASSET_INPUT_ERR_MSG = "SET_ASSET_INPUT_ERR_MSG"
export const setAssetInputErrMsg = (msg) => {
  return{
    type: SET_ASSET_INPUT_ERR_MSG,
    msg
  }
}

/**
Convenience thunk for dispatching the fetchAsset action over all assets
being tracked.
**/
export function fetchAllAssets(){

  return function(dispatch, getState){
    let state = getState();

    //loop through each asset and dispatch
    state.asset.assets.forEach(function(asset){

      //call the singular fetchAsset thunk
      dispatch(fetchAsset(asset.data.assetID))
    })
  }
}

/**
Fetches the asset from the blockchain and adds it to the asset list
assetId: The ID of the asset to fetch

There is no polling function for individual assets. Instead, we will dispatch
fetchAsset from every tracked asset when the blockchain height changes. This
reduces the amount of polling that is required.
**/
export function fetchAsset(assetId){

  return function (dispatch, getState){
    let state = getState();

    let args = {
      assetID: assetId
    }

    let queryRequestPayload = {
      "chaincodeSpec":{
        "type": "GOLANG",
        "chaincodeID":{
          "name":state.configuration.chaincodeId
        },
        "ctorMsg":{
          "function":"readAsset",
          //we need to stringify the object because contract expects a string as args, not an object.
          "args":[JSON.stringify(args)]
        },
        "secureContext":state.configuration.secureContext,
        "confidentialityLevel":"PUBLIC"
      }
    }

    let config = {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(queryRequestPayload)
    }

    return fetch(state.configuration.urlRestRoot + '/devops/query/', config)
    .then(response => response.json())
    .then(json => {
      console.log(json);
      if(json.Error){
        //if there is an error, do not dispatch. Rather, output an error message on the UI.
        dispatch(setAssetInputErrMsg(assetId + " does not exist in the current contract database."));
      }else{

        //ensure that we are not already tracking this asset
        let asset = state.asset.assets.find(function(asset){
          return asset.data.assetID === assetId
        })

        //only add if we could not find a match
        if(!asset){
          dispatch(addAssetToTrack(json.OK))
        }else{
          //otherwise, update
          dispatch(updateAssetToTrack(json.OK))
        }

        //clear the error message
        dispatch(setAssetInputErrMsg(""));
      }
    });
  }
}

/*redux-thunk
Uses fetch to get the object model. The response has the following structure:
{
  "OK": {
    "assetID": "",
    "location": {
      "latitude": 0,
      "longitude": 0
    },
    "temperature": 0,
    "carrier": ""
  }
}
*/
export function fetchAssetObjectModel(){

  return function (dispatch, getState){
    let state = getState();

    //create the payload to communicate with the obc-peer
    let queryRequestPayload = {
      "chaincodeSpec":{
        "type": "GOLANG",
        "chaincodeID":{
          "name":state.configuration.chaincodeId
        },
        "ctorMsg":{
          "function":"readAssetObjectModel",
          //we need to stringify the object because contract expects a string as args, not an object.
          "args":[]
        },
        "secureContext":state.configuration.secureContext,
        "confidentialityLevel":"PUBLIC"
      }
    }

    let config = {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(queryRequestPayload)
    }

    return fetch(state.configuration.urlRestRoot + '/devops/query/', config)
    .then(response => response.json())
    .then(json => {
      //update state to store the object model.
      dispatch(setAssetObjectModel(json.OK))
    })

  }
}
