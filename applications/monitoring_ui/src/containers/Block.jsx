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
    //this.loadBlockInfoFromServer(this.props.url);
    this.props.fetchBlockData(this.props.blockNumber)
  }

  render(){
    // combine transactions and chaincodeEvents, either can be missing or incomplete!
    let blockMap = {}
    const bd = this.props.blockData
    if (bd) {
      if (bd.transactions) {
        for (let i=0; i<bd.transactions.length; i++) {
          let t = bd.transactions[i]
          blockMap[t.txid] = {
            timestamp: t.timestamp,
            function: window.atob(t.payload).split("\n")[2],
            args: window.atob(t.payload).split("\n")[3].substr(1),
            chaincodeID: window.atob(t.chaincodeID)
          }
        }
      }
      if (bd.nonHashData) {
        if (bd.nonHashData.chaincodeEvents) {
          for (let i=0; i<bd.nonHashData.chaincodeEvents.length; i++) {
            let e = bd.nonHashData.chaincodeEvents[i]
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
    let blockArr = []
    for (var p in blockMap) {
      if (blockMap.hasOwnProperty(p)) {
          let a = blockMap[p]
          a.txid = p
          blockArr.push(a)
      }
    }
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
