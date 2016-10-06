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
import Paper from 'material-ui/lib/paper'
import React, { PropTypes } from 'react'

import TextField from 'material-ui/lib/text-field';
import AppBar from 'material-ui/lib/app-bar';
import FlatButton from 'material-ui/lib/flat-button';

import {enableRemoveBtn, disableRemoveBtn, removeResponsePayload, clearResponsePayloads, togglePayloadPolling} from '../actions/ChaincodeActions'
import ResponsePayloadView from '../components/ResponsePayloadView'

import uuid from 'node-uuid'

/**
The ResponsePayloadContainer is responsible for showing the
**/
class ResponsePayloadContainer extends React.Component{

  displayPayload = (responsePayload) => {}

  //objToJsx should be an empty array
  readObjProps = (obj, objToJsx, indents) => {

    var divStyle = {
      whiteSpace: "pre-wrap",
      fontFamily: "Consolas,Monaco,Lucida Console,Liberation Mono,DejaVu Sans Mono,Bitstream Vera Sans Mono,Courier New, monospace"
    }

    objToJsx.push(<p key={uuid.v4() } style={divStyle} > {JSON.stringify(obj, null, 3) }</p>)
    return objToJsx
  }


  render(){
    return(
      <Paper style={{marginBottom:20}}>
        <AppBar
          title={"Response Payloads"}
          showMenuIconButton={false}
          iconElementRight={<FlatButton label="Clear" onTouchTap={this.props.clearResponsePayloads}/>}
        />

      {/*Iterate through all the payloads being monitored and display them on the UI*/}
      {this.props.responsePayloads && this.props.responsePayloads.map( (rPayload, index) => {
          return <ResponsePayloadView key={index} rPayload={rPayload}
            removeFn={this.props.removePayload} displayFn={this.readObjProps}
            enableRemoveBtnFn={this.props.enableRemoveBtnFn} disableRemoveBtnFn={this.props.disableRemoveBtnFn}
            index={index} isRemoveBtnEnabled={rPayload.isRemoveBtnEnabled} togglePayloadPolling={this.props.togglePayloadPolling}/>
        }
      )}
      </Paper>
    );
  }

}

const mapStateToProps = (state) =>{
  return {
    responsePayloads: state.chaincode.ui.responsePayloads
  }
}

const mapDispatchToProps = (dispatch) =>{
  return{
    enableRemoveBtnFn: (index) => {
      dispatch(enableRemoveBtn(index))
    },
    disableRemoveBtnFn: (index) => {
      dispatch(disableRemoveBtn(index))
    },
    removePayload: (index) => {
      dispatch(removeResponsePayload(index))
    },
    clearResponsePayloads: () =>{
      dispatch(clearResponsePayloads())
    },
    togglePayloadPolling: (index) => {
      dispatch(togglePayloadPolling(index))
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(ResponsePayloadContainer)
