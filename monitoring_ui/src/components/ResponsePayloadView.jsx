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
import Toggle from 'material-ui/lib/toggle';

const ResponsePayloadHeaderView = ({args, removeFn, isRemoveBtnEnabled, index}) => (
  <div>
    {isRemoveBtnEnabled ? <FlatButton label="X" primary={true} onClick={() => {removeFn(index)}}/> :
    <FlatButton label="X" disabled={true}/>
    }

    {/*the label for the response payloads are the args used to retrieve it.*/}
    {JSON.stringify(args)}
  </div>
)

/**
onMouseEnter and onMouseLeave enable UI alterations for whether or not the user
is hovering over the card header

removeFn is responsible for removing the asset from the tracking list
showRemoveFn is responsible for showing the remove asset button that calls removeFn
hideRemoveFn is responsible for hiding the remove asset button that calls removeFn
**/
const ResponsePayloadView = ({rPayload, removeFn, displayFn, enableRemoveBtnFn, disableRemoveBtnFn, index, isRemoveBtnEnabled, togglePayloadPolling}) => (
  <Card initiallyExpanded={false}>
    <CardHeader
      title={<ResponsePayloadHeaderView args={rPayload.args} removeFn={removeFn} isRemoveBtnEnabled={isRemoveBtnEnabled} index={index}/>}
      actAsExpander={false}
      showExpandableButton={true}
      onMouseEnter={() => {enableRemoveBtnFn(index)}}
      onMouseLeave={() => {disableRemoveBtnFn(index)}}
    />
    <CardText expandable={true}>
      {/*displayObj is a function that describes how to display the data.*/}
      <Toggle
        onToggle={()=>{togglePayloadPolling(index)}}
        label="Poll for changes"
      />
      {displayFn(rPayload.responsePayload, [], 0).map(function(item){
        return item;
      })}
    </CardText>
  </Card>
)

export default ResponsePayloadView
