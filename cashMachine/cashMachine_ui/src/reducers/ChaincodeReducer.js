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
import update from 'react-addons-update'
import {SET_CC_SCHEMA, SET_CURRENT_TAB,TAB_CREATE, TAB_READ, TAB_UPDATE, TAB_DELETE, SET_CC_OPS, INVOKE, QUERY,
ENABLE_REMOVE_BTN, DISABLE_REMOVE_BTN, REMOVE_RESPONSE_PAYLOAD, ADD_RESPONSE_PAYLOAD, CLEAR_RESPONSE_PAYLOADS,
ENABLE_PAYLOAD_POLLING, UPDATE_RESPONSE_PAYLOAD, DISABLE_PAYLOAD_POLLING, OPEN_SNACKBAR, HIDE_SNACKBAR} from '../actions/ChaincodeActions'


//the chaincode reducer default state is an empty object
export const chaincode = (state={
  ui:{
    currentTab:TAB_CREATE,
    possibleTabs: [
      {name: TAB_CREATE, type: INVOKE},
      {name:TAB_READ, type: QUERY},
      {name: TAB_UPDATE, type: INVOKE},
      {name: TAB_DELETE, type: INVOKE}
    ],
    /*this is the list of payloads that we need to display on the ui
      {
        args: {}
        fn: string
        opType: string -> what type of operation is this. QUERY or INVOKE.
        isPolling: true/false -> this determines if we should continually update the payload display
        responsePayload: {} -> the actual contents of the response payload
        isRemoveEnabled: true/false -> this determines whether or not the remove button is in the enabled or disabled state.
      }
    */
    responsePayloads:[

    ],
    snackbar:{
      open: false,
    }
  }
}, action) =>{
  switch (action.type){
    case SET_CC_SCHEMA:
      return Object.assign({}, state, {
        schema: action.schema
      })
    case SET_CURRENT_TAB:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          currentTab: action.tab
        }
      })
    //set the chaincode operations object
    case SET_CC_OPS:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          ops: action.ops
        }
      })
    case ENABLE_REMOVE_BTN:
      return update(state, {
        ui:{
          responsePayloads:{
            [action.index]:{
              isRemoveBtnEnabled:{
                 $set:true
              }
            }
          }
        }
      })
    case DISABLE_REMOVE_BTN:
      return update(state, {
        ui:{
          responsePayloads:{
            [action.index]:{
              isRemoveBtnEnabled:{
                $set:false
              }
            }
          }
        }
      })
    case REMOVE_RESPONSE_PAYLOAD:
    return Object.assign({}, state, {
      ui:{
        ...state.ui,
        responsePayloads:
          [
            ...state.ui.responsePayloads.slice(0, action.index),
            ...state.ui.responsePayloads.slice(action.index+1)
          ]
      }
    })
    case ADD_RESPONSE_PAYLOAD:
      return update(state, {
        ui:{
          responsePayloads:{
            $push:
              [
                {
                  args: action.args,
                  fn: action.fn,
                  opType: action.opType,
                  isPolling: action.isPolling,
                  responsePayload: action.rPayload,
                  isRemoveBtnEnabled: action.isRemoveBtnEnabled,
                }
              ]
          }
        }
      })
    case CLEAR_RESPONSE_PAYLOADS:
      return Object.assign({}, state, {
        ui:{
          ...state.ui,
          responsePayloads:
            []
        }
      })
    case ENABLE_PAYLOAD_POLLING:
      return update(state,{
        ui:{
          responsePayloads:{
            [action.index]:{
              isPolling:{
                $set: true
              }
            }
          }
        }
      })
    case DISABLE_PAYLOAD_POLLING:
      return update(state,{
        ui:{
          responsePayloads:{
            [action.index]:{
              isPolling:{
                $set: false
              }
            }
          }
        }
      })
    case UPDATE_RESPONSE_PAYLOAD:
      return update(state,{
        ui:{
          responsePayloads:{
            [action.index]:{
              responsePayload:{
                $set: action.payload
              }
            }
          }
        }
      })
    default:
      return state
  }
}
