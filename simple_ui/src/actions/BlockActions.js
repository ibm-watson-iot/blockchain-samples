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
import {fetchAllAssets} from './AssetActions'

export const REQUEST_BLOCK_INFO = 'REQUEST_BLOCK_INFO'
function requestBlockInfo(blockNum){
  return{
    type: REQUEST_BLOCK_INFO,
    blockNum
  }
}

export const RECEIVE_BLOCK_INFO = 'RECEIVE_BLOCK_INFO'
function receiveBlockInfo(blockNum, data){
  return{
    type: RECEIVE_BLOCK_INFO,
    blockNum,
    data: data
  }
}

export const ADD_BLOCKS = 'ADD_BLOCKS'
export function addBlocks(chainHeight){
  return{
    type: REQUEST_BLOCK_INFO,
    chainHeight
  }
}

export const CLEAR_BLOCKCHAIN = 'CLEAR_BLOCKCHAIN'
export function clearBlockchain(){
  return{
    type: CLEAR_BLOCKCHAIN
  }
}

export const REC_CHAIN_HEIGHT = 'REC_CHAIN_HEIGHT'
export function recChainHeight(chainHeight, blocksPerPage){
  return{
    type: 'REC_CHAIN_HEIGHT',
    chainHeight,
    blocksPerPage
  }
}

/**
This is a redux-thunk. We basically invoke all other events that are dependent
on a block being added to the blockchain.
**/
export function fetchChainHeight(urlRestRoot){

  return function (dispatch, getState){

    return fetch(urlRestRoot + '/chain')
    .then(response => response.json())
    .then((json) =>{

      //check if the chain height is any different
      let state = getState();

      //call any other functions that need to be invoked on change in blockchain height
      if(state.blockchain.length === 0 || ((state.blockchain[0].blockNumber+1) !== json.height)){

        //when we receive the chain height, we must also use the blocks per page configuration to generate the blockchain correctly
        dispatch(recChainHeight(json.height, state.configuration.blocksPerPage))

        dispatch(fetchAllAssets());
      }
    })

  }
}

/**
Fetches the data for an individual block. This data appears when you click
on a block header on the UI.
**/
export function fetchBlockData(blockNum){

  //dispatch comes from thunk middleware
  return function (dispatch, getState){
    //first update the app state to indicate we are fetching data.
    dispatch(requestBlockInfo(blockNum));

    const state = getState();

    let blockData = (state.blockchain[state.blockchain[0].blockNumber - blockNum]).blockData

    //first make sure that there is no data for this object yet. If there is, we just resolve a promise, because data can't change.
    //The index manipulation is because the blocks are displayed in reverse on the UI, but the blocknums are increasing sequentially.
    if(blockData){
      return Promise.resolve();
    }

    //then return promise. Not necessary but a convienience in case we want to use then against the return.
    return fetch(state.configuration.urlRestRoot + '/chain/blocks/' + blockNum)
    .then(response => response.json())
    .then(json => {
      dispatch(receiveBlockInfo(blockNum, json))
    })
  }
}
