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
import { connect } from 'react-redux';
import { Field, Form, actions } from 'react-redux-form';
import TextField from 'material-ui/lib/text-field';
import FlatButton from 'material-ui/lib/flat-button';

import { createFieldClass, controls } from 'react-redux-form';

import { setConfiguration, setConfigDialogDisplay } from '../../actions/ConfigurationActions'

const MaterialField = createFieldClass({
  'TextField': controls.text
});

class ObcConfigurationForm extends React.Component {

  handleSubmit(obcConfiguration) {
    let { dispatch } = this.props;

    console.log(obcConfiguration);

    //set the properties specific to obc in our configuration store
    dispatch(setConfiguration(obcConfiguration))

    //close the dialog
    dispatch(setConfigDialogDisplay(false))
  }
  render() {
    let { obcConfiguration } = this.props;

    return (
      <Form model="obcConfiguration" noValidate
        onSubmit={(obcConfiguration) => this.handleSubmit(obcConfiguration)}>
        <MaterialField model="obcConfiguration.urlRestRoot">
          <TextField
            hintText="http://localhost:3000"
            floatingLabelText = "API Host and Port"
            fullWidth={true}
            />
        </MaterialField>
        <br/>
        <MaterialField model="obcConfiguration.chaincodeId">
          <TextField
            hintText="7cdd53526ed31f7be5249bfa42..."
            floatingLabelText = "Chaincode ID"
            fullWidth={true}
            />
        </MaterialField>
        <br/>
        <MaterialField model="obcConfiguration.secureContext">
          <TextField
            hintText="user_context"
            floatingLabelText = "Secure Context"
            fullWidth={true}
            />
        </MaterialField>
        <MaterialField model="obcConfiguration.blocksPerPage">
          <TextField
            hintText="ex. 100 or empty to display all"
            floatingLabelText = "Number of blocks to display"
            fullWidth={true}
            />
        </MaterialField>
        <br/>
          {/*submit button for form*/}
          <div style={{textAlign: 'right'}}>
            <FlatButton label="Submit" primary={true} type={"submit"}/>
          </div>
      </Form>
    );
  }
}

function mapStateToProps(state) {
  return { obcConfiguration: state.obcConfiguration };
}

export default connect(mapStateToProps)(ObcConfigurationForm);
