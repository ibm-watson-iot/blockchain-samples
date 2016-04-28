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
import equal from 'deep-equal';
import {openSnackbar, hideSnackbar, setSnackbarMsg} from '../actions/AppActions';
import * as strings from '../resources/strings'

export const SET_CC_SCHEMA = "SET_CC_SCHEMA"
export const setCcSchema = (schema) => {
  return{
    type: SET_CC_SCHEMA,
    schema
  }
}

export const SET_CC_SCHEMA_OBJ = "SET_CC_SCHEMA_OBJ"
export const setCcSchemaObj = (obj) => {
  return{
    type: SET_CC_SCHEMA_OBJ,
    schema
  }
}

export const SET_CC_OPS = "SET_CC_OPS"
export const setCcOps = (ops) => {
  return{
    type: SET_CC_OPS,
    ops
  }
}

export const SET_CURRENT_TAB = "SET_CURRENT_TAB"
export const setCurrentTab = (tab) => {
  return{
    type: SET_CURRENT_TAB,
    tab
  }
}

//this enables the remove button on the UI. This happens when the user
//hovers over the header for a particular payload.
export const ENABLE_REMOVE_BTN = "ENABLE_REMOVE_BTN"
export const enableRemoveBtn = (index) => {
  return{
    type: ENABLE_REMOVE_BTN,
    index
  }
}

//this disables the remove button on the UI.
export const DISABLE_REMOVE_BTN = "DISABLE_REMOVE_BTN"
export const disableRemoveBtn = (index) => {
  return{
    type: DISABLE_REMOVE_BTN,
    index
  }
}

export const REMOVE_RESPONSE_PAYLOAD = "REMOVE_RESPONSE_PAYLOAD"
export const removeResponsePayload = (index) => {
  return{
    type: REMOVE_RESPONSE_PAYLOAD,
    index
  }
}

//Add a response payload to the payload
export const ADD_RESPONSE_PAYLOAD = "ADD_RESPONSE_PAYLOAD"
export const addResponsePayload = (args, fn, opType, rPayload, isPolling, isRemoveBtnEnabled) => {
  return{
    type: ADD_RESPONSE_PAYLOAD,
    args,
    fn,
    opType,
    rPayload,
    isPolling,
    isRemoveBtnEnabled
  }
}

export const CLEAR_RESPONSE_PAYLOADS = "CLEAR_RESPONSE_PAYLOADS"
export const clearResponsePayloads = () => {
  return{
    type: CLEAR_RESPONSE_PAYLOADS
  }
}

export const ENABLE_PAYLOAD_POLLING = "ENABLE_PAYLOAD_POLLING"
export const enablePayloadPolling = (index) => {
  return{
    type: ENABLE_PAYLOAD_POLLING,
    index
  }
}

export const DISABLE_PAYLOAD_POLLING = "DISABLE_PAYLOAD_POLLING"
export const disablePayloadPolling = (index) => {
  return{
    type: DISABLE_PAYLOAD_POLLING,
    index
  }
}

export const UPDATE_RESPONSE_PAYLOAD = "UPDATE_RESPONSE_PAYLOAD"
export const updateResponsePayload = (index, payload) => {
  return{
    type: UPDATE_RESPONSE_PAYLOAD,
    index,
    payload
  }
}

export const TAB_CREATE = "CREATE";
export const TAB_READ = "READ";
export const TAB_UPDATE = "UPDATE";
export const TAB_DELETE = "DELETE";

export const INVOKE = "INVOKE";
export const QUERY = "QUERY";

import {
  actions
} from 'react-redux-form';

//create an object that stores all functions. The possibleTabs is the UI representation
//of all the tabs that are possible. This model is specifically for the form.
const createChaincodeOpsModel = (schema, possibleTabs) => {

  let obj = {}

  possibleTabs.forEach(function(tab){
    obj[tab.name] = {
      fns: [],
      selectedFn: ""
    }
  })

  let api = schema.API;

  //we assume a one to one correlation between the possible tabs and the first word of every function
  //loop through the api object and pick up any functions
  for (var fn in api) {
    if (api.hasOwnProperty(fn)) {
      let lowerFn = fn.toLowerCase();

      for(let i = 0; i < possibleTabs.length; i++){
        //look through the lowercased function to figure out what tab it belongs to
        if(lowerFn.indexOf(possibleTabs[i].name.toLowerCase()) === 0){
          //push this to the list of functions
          obj[possibleTabs[i].name].fns.push({
            name: fn,
            desc: api[fn].description,
          });

          //if this is the first entry, make it the default selected function
          if(obj[possibleTabs[i].name].fns.length === 1){
            obj[possibleTabs[i].name].selectedFn = 0
          }
        }
      }
    }
  }

  return obj;
}

/**
This function does an OBC request for every payload that has polling enabled.
This allows the UI to update without the user having to manually run another query.
**/
export function sendObcPollingRequests(){
  return function(dispatch, getState){
    let state = getState();

    state.chaincode.ui.responsePayloads.forEach(function(payload){
      if(payload.isPolling){
        dispatch(sendObcRequest(payload.args, payload.fn, payload.opType))
      }
    })

  }
}

/**
Send an http request to the OBC Peer. The requestType is either query or invoke.
The args is the form data.
**/
export function sendObcRequest(args, fn, requestType){

  return function(dispatch, getState){
    let state = getState();
    //iterate through the args and delete any empty strings

    for(var propertyName in args){
      if (args.hasOwnProperty(propertyName)) {
        if(args[propertyName] === ""){
          delete args[propertyName]
        }
      }
    }

    let requestPayload = {
      "jsonrpc": "2.0",
      "method": requestType.toLowerCase(),
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":state.configuration.chaincodeId
          },
          "ctorMsg":{
            "function":fn,
            //we need to stringify the object because contract expects a string as args, not an object.
            "args": args ? [JSON.stringify(args)] : []
          },
          "secureContext":state.configuration.secureContext
      },
      "id": 5
    }

    let config = {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestPayload)
    }

    return fetch(state.configuration.urlRestRoot + '/chaincode/', config)
    .then(response => response.json())
    .then(json => {

      if(json.error){
        dispatch(setSnackbarMsg(json.error.data));
        dispatch(openSnackbar());
      }else{
        let alreadyRequested = false;
        let indexOfMatch = -1;

        //If this is a query type, then we should display the response payload on the UI.
        if(requestType === QUERY){
          //first we check if the response payload already exists. If it does, then we update. Otherwise, we add.
          for(let i=0; i < state.chaincode.ui.responsePayloads.length; i++){
            //we compare 3 properties to verify equality: args, fn and type.
            let payload = state.chaincode.ui.responsePayloads[i];

            if(equal(payload.args, args) && payload.fn === fn && payload.opType === requestType){
              alreadyRequested = true;
              indexOfMatch = i;
              break;
            }
          }

          //we found a match, which means we should be updating, not appending.
          if(alreadyRequested){
            dispatch(updateResponsePayload(indexOfMatch, JSON.parse(json.result.message)))
          }else{
            dispatch(addResponsePayload(args, fn, QUERY, JSON.parse(json.result.message), false, false))
          }
        }
      }
    })
  }
}

/**
This is a redux-thunk. We request the chaincode schema.
**/
export function fetchCcSchema(){

  return function (dispatch, getState){
    let state = getState();

    let queryRequestPayload = {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":state.configuration.chaincodeId
          },
          "ctorMsg":{
            "function":"readAssetSchemas",
            //we need to stringify the object because contract expects a string as args, not an object.
            "args": []
          },
          "secureContext":state.configuration.secureContext
      },
      "id": 5
    }

    let config = {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(queryRequestPayload)
    }

    return fetch(state.configuration.urlRestRoot + '/chaincode/', config)
    .then(response => response.json())
    .then(json => {

      console.log(json);

      //if there is an error, display it
      if(json.error){
        dispatch(setSnackbarMsg(json.error.data));
        dispatch(openSnackbar())

        //update state to store the object model.
        dispatch(setCcSchema({}))

      }else{
        //update state to store the object model.
        dispatch(setCcSchema(JSON.parse(json.result.message)))

        //then parse through the cc schema and create an object
        let chaincodeOpsModel = createChaincodeOpsModel(JSON.parse(json.result.message), state.chaincode.ui.possibleTabs)

        //set the chaincode ops
        //this is tied directly to the form model, so we use the react-redux-form actions.change function
        dispatch(actions.change('chaincodeOpsForm', chaincodeOpsModel))
      }
    })

  }
}

export function togglePayloadPolling(index){

  //we do the necessary operations to enable or disable polling
  return function(dispatch,getState){
    let state = getState();

    //this is the ui container object. The raw payload data is a property called
    //response payloads.
    let payloadUi = state.chaincode.ui.responsePayloads[index];

    //we are currently polling, perform operations to disable polling
    if(payloadUi.isPolling){
      dispatch(disablePayloadPolling(index))
    }else{
      //we are currently not polling, perform operations to enable polling
      dispatch(enablePayloadPolling(index))

      //and do an immediate update, in case the asset has changed when polling
      //was disabled.
      dispatch(sendObcRequest(payloadUi.args, payloadUi.fn, payloadUi.opType))

    }
  }
}
