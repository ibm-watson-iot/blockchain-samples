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

class ConfigurationDialog extends React.Component {
  render() {
    const actions = [
      <FlatButton
        label="Cancel"
        secondary={true}
        onTouchTap={this.props.closeDialog}
      />,
      <FlatButton
        label="Submit"
        primary={true}
        disabled={true}
        onTouchTap={this.props.closeDialog}
      />,
    ];

    return (
      <div>
        <FlatButton label="Configuration" onTouchTap={this.props.openDialog} style={{color:"#ffffff", marginTop: 8}}/>
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
    showDialog: state.configuration.showDialog
  }
}

const mapDispatchToProps = (dispatch) => {
  return{
    closeDialog: () => {
      dispatch(setConfigDialogDisplay(false))
    },
    openDialog: () => {
      dispatch(setConfigDialogDisplay(true))
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(ConfigurationDialog)
