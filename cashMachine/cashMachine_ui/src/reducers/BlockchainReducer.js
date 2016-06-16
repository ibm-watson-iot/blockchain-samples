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
import _ from 'underscore'
import moment from 'moment'
import update from 'react-addons-update'

import{
  CLEAR_BLOCKCHAIN, REC_CHAIN_HEIGHT, RECEIVE_BLOCK_INFO
} from '../actions/BlockActions.js'


//the blockchain reducer default state is an empty array
export const blockchain = (state=[], action) =>{
  switch (action.type){
    case REC_CHAIN_HEIGHT:
    //Need to calculate the diff based on last block number, not state.length. State.length
    //is not configurable and may not necessarily reflect the entire length of the blockchain
      //get the difference between the action chain height and the existing chain height

      let diff = 0;

      //if there are actually elements in the blockchain, we can calculate the diff
      if(state.length > 0){
        diff = action.chainHeight - (state[0].blockNumber + 1); //the first blocknumber + 1 should be the height of what we have stored.
      }else{
        diff = action.chainHeight
      }
      //only do this if a parameter for blocks per page is passed in. Otherwise, display all blocks.
      if (action.blocksPerPage && (diff > action.blocksPerPage)){
        diff = action.blocksPerPage;
      }

      //if there is no difference, we return the state back and do nothing.
      if (diff <= 0){
        return state;
      }

      //if we hit the limit, we have to slice
      if (action.chainHeight > action.blocksPerPage){
        return[        //append appropriate titles to each block in state
          ...(_.range(diff).map(function(item,index){
            let adjustedIndex = action.chainHeight - index - 1;
            return {
              blockNumber: adjustedIndex,
              isExpanded: false,
            }
          })),
        ...state.slice(0, state.length - diff)
        ]
      }else{
        //return an array that prepends the new block range to the existing blocks and slices off the new blocks.
        return[        //append appropriate titles to each block in state
          ...(_.range(diff).map(function(item,index){
            let adjustedIndex = action.chainHeight - index - 1;
            return {
              blockNumber: adjustedIndex,
              isExpanded: false,
            }
          })),
        ...state
        ]
      }
    case RECEIVE_BLOCK_INFO:
      //use a computed property to update the correct block with data. The index
      //manipulation is because the blocks are in reverse order on the UI.
      return update(state, {
        [state[0].blockNumber - action.blockNum]:{blockData: {$set: action.data}}
      });
    case CLEAR_BLOCKCHAIN:
      return []
    default:
      return state
  }
}
