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

import {
  ADD_ASSET_TO_TRACK, REMOVE_ASSET_FROM_TRACKING, SET_ASSET_OBJECT_MODEL, SET_ASSET_ID_INPUT, UPDATE_ASSET_TO_TRACK,
  TOGGLE_SHOW_REMOVE_BTN, HIDE_REMOVE_BTN, SHOW_REMOVE_BTN, CLEAR_TRACKED_ASSETS, SET_ASSET_INPUT_ERR_MSG
} from '../actions/AssetActions.js'

//set default configuration
export const asset = (state={
  /*assets will end up looking like:
  assets:[
    {
      showRemove: false
      data: {
        //json.OK response data
      }
    }
  ]
  */
  assets: [],
  //stores the raw asset object model
  objectModel: {},
  //UI field for entering the asset ID to track
  assetIdInput: "",
  assetInputErrMsg: "",
}, action) => {
  switch (action.type){
    case ADD_ASSET_TO_TRACK:
      return update(state,{
        assets:{
          $push: [
            {
              showRemove: false,
              data: action.assetInfo,
            }
          ]
        }
      })

    case REMOVE_ASSET_FROM_TRACKING:
      //find the index and then remove
      let index = 0;
      for(index; index < state.assets.length; index++){
        if(state.assets[index].data.assetID === action.assetId){
          break;
        }
      }

      //remove the asset by slicing around it and then assigning the new assets array.
      return Object.assign({}, state, {
        assets:[
          ...state.assets.slice(0, index),
          ...state.assets.slice(index+1)
        ]
      })

    case CLEAR_TRACKED_ASSETS:
      return Object.assign({}, state, {
        assets:[]
      })

    //set the asset object model.
    case SET_ASSET_OBJECT_MODEL:
      return Object.assign({}, state, {
        objectModel: action.objectModel
      })

    case SET_ASSET_ID_INPUT:
      return Object.assign({}, state, {
        assetIdInput: action.assetId
      })

    case UPDATE_ASSET_TO_TRACK:

      let i = 0;
      for(i; i < state.assets.length; i++){
        if(state.assets[i].data.assetID === action.assetData.assetID){
          break;
        }
      }

      return update(state,{
        assets:{
          [i]:{
            data:{
              $set: action.assetData
            }
          }
        }
      })

    /**
    Strictly a UI action. Governs whether or not the remove asset button is shown
    for a particular asset.
    **/

    case HIDE_REMOVE_BTN:
      return update(state,{
        assets:{
          [action.index]:{
            showRemove:{
              $set: false
            }
          }
        }
      })

      case SHOW_REMOVE_BTN:
        return update(state,{
          assets:{
            [action.index]:{
              showRemove:{
                $set: true
              }
            }
          }
        })

    case TOGGLE_SHOW_REMOVE_BTN:
      //invert the showRemove property to get what the status should be after toggling
      let endStatus = !state.assets[action.index].showRemove;

      return update(state,{
        assets:{
          [action.index]:{
            showRemove:{
              $set: endStatus
            }
          }
        }
      })

    case SET_ASSET_INPUT_ERR_MSG:
      return Object.assign({},state,{
        assetInputErrMsg: action.msg
      })
    default:
      return state;
  }
}
