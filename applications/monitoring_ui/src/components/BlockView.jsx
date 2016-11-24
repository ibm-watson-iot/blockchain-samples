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
import React, { PropTypes } from 'react'
import Card from 'material-ui/lib/card/card'
import CardActions from 'material-ui/lib/card/card-actions'
import CardHeader from 'material-ui/lib/card/card-header'
import FlatButton from 'material-ui/lib/flat-button'
import CardText from 'material-ui/lib/card/card-text'
import moment from 'moment'
import momentPropTypes from 'react-moment-proptypes'
import * as strings from '../resources/strings'

const FUNCTION_PAYLOAD_INDEX = 2

const BlockView = ({isExpanded, blockNumber, timestampString, onBlockClick, urlRestRoot, blockData}) => (
  <Card initiallyExpanded={isExpanded}>
    <CardHeader
      title={strings.BLOCK_CARD_HEADER_TEXT + " #" + blockNumber}
      actAsExpander={true}
      showExpandableButton={true}
      subtitle={(typeof blockData === "object") ? moment.unix(blockData.nonHashData.localLedgerCommitTimestamp.seconds).format("M/D/YY LT") : ""} />
    <CardText expandable={true}>
      <h5>{(typeof blockData === "object" && typeof blockData.transactions === "object") ? blockData.transactions.length + " " + strings.BLOCK_CARD_CONTENTS_TRANSACTION_TEXT : "No Transactions"}</h5>
      <ol>
        {(typeof blockData === "object" && typeof blockData.transactions === "object") ? blockData.transactions.map(function (transaction, index) {
          return (
            <li key={index}> {transaction.txid}
              <ul>
                <li>
                  {moment.unix(transaction.timestamp.seconds).format("M/D/YY LT")}
                </li>
                <li>
                  {window.atob(transaction.payload).split('\n')[FUNCTION_PAYLOAD_INDEX]}
                </li>
                <li>
                  {window.atob(transaction.payload).split('\n')[3].substr(1)}
                </li>
              </ul>
            </li>
          );
        }) : null}
      </ol>
    </CardText>
  </Card>
)

// When working with the new platform, you can insert this block of code (uncommented of course) just above the </CardText>.
// For some reason, the addition of this code causes display of BlockView when there are no chaincodeEvents.
// None of this makes much sense, and we'll have to restructure the UI with redux at some point to clean it up.
// Meanwhile, please do not commit it back to blockchain-samples with this in place, as older contracts do not work with it this way.
      // <h5>{(typeof blockData === "object" && typeof blockData.nonHashData === "object" && typeof blockData.nonHashData.chaincodeEvents === "object") ? blockData.nonHashData.chaincodeEvents.length + " " + "Chaincode Events" : "No Chaincode Events"}</h5>
      // <ol>
      //   {(typeof blockData === "object" && typeof blockData.nonHashData === "object" && typeof blockData.nonHashData.chaincodeEvents === "object") ? blockData.nonHashData.chaincodeEvents.map(function (event, index) {
      //     return (
      //       <li key={index}> {event.txID}
      //         <ul>
      //           <li>
      //             {event.eventName}
      //           </li>
      //           <li>
      //             {window.atob(event.payload)}
      //           </li>
      //         </ul>
      //       </li>
      //     );
      //   }) : null}
      // </ol>


//define the properties that a BlockView is expecting.
BlockView.propTypes = {
  isExpanded: PropTypes.bool.isRequired,
  blockNumber: PropTypes.number.isRequired,
  timeStampString: momentPropTypes.momentObj,
  transactionsContent: PropTypes.array,
  onBlockClick: PropTypes.func,
  urlRestRoot: PropTypes.string.isRequired
}

export default BlockView
