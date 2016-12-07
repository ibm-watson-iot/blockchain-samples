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
import React, {PropTypes} from 'react'
import Card from 'material-ui/lib/card/card'
import CardActions from 'material-ui/lib/card/card-actions'
import CardHeader from 'material-ui/lib/card/card-header'
import FlatButton from 'material-ui/lib/flat-button'
import CardText from 'material-ui/lib/card/card-text'
import moment from 'moment'
import momentPropTypes from 'react-moment-proptypes'
import * as strings from '../resources/strings'

const FUNCTION_PAYLOAD_INDEX = 2

const BlockView = ({isExpanded, blockNumber, blockData, blockArr}) => (
  <Card initiallyExpanded={isExpanded}>
    <CardHeader
      title={strings.BLOCK_CARD_HEADER_TEXT + " #"+blockNumber}
      actAsExpander={true}
      showExpandableButton={true}
      subtitle={blockData ? moment.unix(blockData.nonHashData.localLedgerCommitTimestamp.seconds).format("M/D/YY LT") : ""}/>
    <CardText expandable={true}>
      <u>{blockArr ? blockArr.length + " " + strings.BLOCK_CARD_CONTENTS_TRANSACTION_TEXT : "ERROR Block Array Missing" }</u>
      <ol>
      {blockArr ? blockArr.map(function(transaction, index){
          return(
             <li key={index}> {transaction.txid}
               <ul>
                 <li>
                   Chaincode ID: {transaction.chaincodeID}
                 </li>
                 <li>
                   Timestamp: {transaction.timestamp ? moment.unix(transaction.timestamp.seconds).format("M/D/YY HH:mm:ss.") + Math.floor(transaction.timestamp.nanos/1000000) : "n/a"}
                 </li>
                 <li>
                   Function: {transaction.function ? transaction.function : "n/a"}
                 </li>
                 <li>
                   Args: {transaction.args ? transaction.args : "n/a"}
                 </li>
                 <li>
                   Event Emitted: {transaction.eventName}
                 </li>
                 <li>
                   Event Payload: {transaction.event}
                 </li>
               </ul>
             </li>
             );
      }) : null}
    </ol>
    </CardText>
  </Card>
)

//define the properties that a BlockView is expecting.
BlockView.propTypes ={
  isExpanded: PropTypes.bool.isRequired,
  blockNumber: PropTypes.number.isRequired,
  blockData: PropTypes.object.isRequired,
  blockArr: PropTypes.array.isRequired
}

export default BlockView
