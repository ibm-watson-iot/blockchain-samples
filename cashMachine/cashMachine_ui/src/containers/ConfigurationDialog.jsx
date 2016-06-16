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
import React from 'react';

import Dialog from 'material-ui/lib/dialog';
import FlatButton from 'material-ui/lib/flat-button';
import Paper from 'material-ui/lib/paper';
import { connect } from 'react-redux'

import ObcConfigurationForm from './forms/ObcConfigurationForm'

import {setConfigDialogDisplay} from '../actions/ConfigurationActions'

import * as strings from '../resources/strings'

import { actions, getField } from 'react-redux-form';

class ConfigurationDialog extends React.Component {
  render() {
    return (
      <div>
        <FlatButton label={strings.OBC_CONFIG_DIALOG_TITLE} onTouchTap={this.props.openDialog} style={{color:"#ffffff", marginTop: 8}}/>
        <Dialog
          title="Configuration"

          modal={false}
          open={this.props.showDialog}
          onRequestClose={this.props.closeDialog}
          autoScrollBodyContent={true}
          autoDetectWindowHeight={false}
          repositionOnUpdate={false}
          >
            <ObcConfigurationForm />
        </Dialog>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return{
    showDialog: state.configuration.showDialog,
    validChaincodeID: state.configuration.chaincodeId,
    obcConfigurationForm: state.obcConfigurationForm
  }
}

const mergeProps = (stateProps, dispatchProps, ownProps) => {
  const {showDialog, validChaincodeID, obcConfigurationForm} = stateProps;
  const {dispatch} = dispatchProps;

  return{
    ...stateProps,
    closeDialog: () => {
      if(getField(obcConfigurationForm, 'chaincodeId').errors){
        dispatch(actions.change('obcConfiguration.chaincodeId',validChaincodeID));
        dispatch(actions.setValidity('obcConfiguration.chaincodeId',true));
      }

      dispatch(setConfigDialogDisplay(false))
    },
    openDialog: () => {
      dispatch(setConfigDialogDisplay(true))
    }
  };
}

export default connect(mapStateToProps, null, mergeProps)(ConfigurationDialog)
