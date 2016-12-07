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
import { connect } from 'react-redux'
import { fetchBlockData } from '../actions/BlockActions'
import React from 'react'
import BlockView from '../components/BlockView.jsx'

class Block extends React.Component{

  constructor(props){
    super(props)

  }

  //when the block is instantiated, we load block information from the server.
  //block information can't change, so we just load once.
  componentDidMount(){
    this.props.fetchBlockData(this.props.blockNumber)
  }

  isEmpty(obj) {
    for(var prop in obj) {
        if(obj.hasOwnProperty(prop))
            return false;
    }

    return JSON.stringify(obj) === JSON.stringify({});
  };
  
  render(){
    // combine transactions and chaincodeEvents, either can be missing or incomplete!
    let blockMap = {}
    const bd = this.props.blockData
    if (bd) {
      if (bd.transactions) {
        for (let i=0; i<bd.transactions.length; i++) {
          let t = bd.transactions[i]
          let p = window.atob(t.payload).split(new RegExp('[^\u0020-\u007e]+', 'g'))
          let f = "n/a"
          let a = "n/a"
          if (p.length === 5) {
            // invoke without arguments
            f = p[3].substr(0,p[3].length-1)
          } else if (p.length === 7) {
            // invoke with arguments
            f = p[4]
            a = p[5].substr(1,p[5].length-2)
          } else if (p.length === 6) {
            // deploy functions have 6 elements
            f = p[4]
            a = p[5].substr(1,p[5].length-1)
          } else {
            // not known to happen in v0.6 fabric
            f = "ERR: ARRAY LENGTH: " + p.length + " [" + JSON.stringify(p) + "]"
          }
          blockMap[t.txid] = {
            timestamp: t.timestamp,
            function: f,
            args: a,
            chaincodeID: window.atob(t.chaincodeID)
          }
        }
      }
      if (bd.nonHashData) {
        if (bd.nonHashData.chaincodeEvents) {
          for (let i=0; i<bd.nonHashData.chaincodeEvents.length; i++) {
            // INCREDIBLY: v0.6 Hyperledger will include one empty opject when the array should be empty
            let e = bd.nonHashData.chaincodeEvents[i]
            if (!this.isEmpty(e)) {
              if (blockMap[e.txID]) {
                blockMap[e.txID].eventName = e.eventName
                blockMap[e.txID].event = window.atob(e.payload)
                blockMap[e.txID].chaincodeID = e.chaincodeID
              } else {
                blockMap[e.txID] = {
                  eventName: e.eventName,
                  event: window.atob(e.payload),
                  chaincodeID: e.chaincodeID
                }
              }
            } 
          }
        }
      }
    }
    let blockArr = []
    for (var p in blockMap) {
      if (blockMap.hasOwnProperty(p)) {
          let a = blockMap[p]
          a.txid = p
          blockArr.push(a)
      }
    }
    console.log("BLOCKARR ----  #", this.props.blockNumber, blockArr)
    return(
      <BlockView isExpanded={this.props.isExpanded} blockNumber={this.props.blockNumber} blockData={this.props.blockData} blockArr={blockArr} />
    )
  }

}

const mapStateToProps = (state, ownProps) =>{
  //calculate the inverse
  let adjustedIndex = state.blockchain[0].blockNumber - ownProps.blockNumber;
  let currBlock = state.blockchain[adjustedIndex];
  currBlock.urlRestRoot = state.configuration.urlRestRoot;
  return currBlock;
}

const mapDispatchToProps = (dispatch) =>{
  return{
    fetchBlockData: (blockNumber) => {
      dispatch(fetchBlockData(blockNumber))
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Block)
