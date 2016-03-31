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

const FUNCTION_PAYLOAD_INDEX = 2

const BlockView = ({isExpanded, blockNumber, timestampString, onBlockClick, urlRestRoot, blockData}) => (
  <Card initiallyExpanded={isExpanded}>
    <CardHeader
      title={"Block #"+blockNumber}
      actAsExpander={true}
      showExpandableButton={true}
      onClick={() => onBlockClick(blockNumber)}
      subtitle={blockData ? moment.unix(blockData.nonHashData.localLedgerCommitTimestamp.seconds).format("M/D/YY LT") : ""}/>
    <CardText expandable={true}>
      <u>{blockData ? blockData.transactions.length + " Transactions" : "" }</u>
      <ol>
      {blockData ? blockData.transactions.map(function(transaction, index){
          return(
             <li key={index}> {transaction.uuid}
               <ul>
                 <li>
                   {moment.unix(transaction.timestamp.seconds).format("M/D/YY LT")}
                 </li>
                 <li>
                   {window.atob(transaction.payload).split('\n')[FUNCTION_PAYLOAD_INDEX]}
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
  timeStampString: momentPropTypes.momentObj,
  transactionsContent: PropTypes.array,
  onBlockClick: PropTypes.func,
  urlRestRoot: PropTypes.string.isRequired
}

export default BlockView
