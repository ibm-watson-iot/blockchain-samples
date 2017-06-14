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
import { Field, Form, actions, getField } from 'react-redux-form';
import TextField from 'material-ui/lib/text-field';
import FlatButton from 'material-ui/lib/flat-button';

import { createFieldClass, controls } from 'react-redux-form';

import { setConfiguration, setConfigDialogDisplay } from '../../actions/ConfigurationActions'

import * as strings from '../../resources/strings'

const MaterialField = createFieldClass({
  'TextField': controls.text
});

class ObcConfigurationForm extends React.Component {

  handleSubmit(obcConfiguration) {
    let { dispatch } = this.props;

    let idWithoutSpaces = obcConfiguration.chaincodeId.replace(/ /g,'');
    //submit if the length is correct
    if(idWithoutSpaces.length > 0 || idWithoutSpaces == "mycc"){
      this.props.dispatch(actions.change('obcConfiguration.chaincodeId',obcConfiguration.chaincodeId.replace(/ /g,'')));

      let config = Object.assign({}, obcConfiguration, {
        chaincodeId: idWithoutSpaces
      })

      //set the properties specific to obc in our configuration store
      dispatch(setConfiguration(config))

      //close the dialog
      dispatch(setConfigDialogDisplay(false))
      this.props.dispatch(actions.setValidity('obcConfiguration.chaincodeId', true));
    }else{
      //set an error message if the length is incorrect
      this.props.dispatch(actions.setValidity('obcConfiguration.chaincodeId', false));
    }


  }
  render() {
    let { obcConfiguration, obcConfigurationForm } = this.props;

    return (
      <Form model="obcConfiguration" noValidate
        onSubmit={(obcConfiguration) => this.handleSubmit(obcConfiguration)}>
        <MaterialField model="obcConfiguration.urlRestRoot">
          <TextField
            hintText={strings.OBC_CONFIG_URL_REST_ROOT_HT}
            floatingLabelText = {strings.OBC_CONFIG_URL_REST_ROOT_FL}
            fullWidth={true}
            />
        </MaterialField>
        <br/>
        <MaterialField model="obcConfiguration.chaincodeId">
          <TextField
            hintText={strings.OBC_CONFIG_CHAINCODE_ID_HT}
            floatingLabelText = {strings.OBC_CONFIG_CHAINCODE_ID_FL}
            fullWidth={true}
            errorText={getField(obcConfigurationForm, 'chaincodeId').valid ? "" : strings.CHAINCODE_LENGTH_ERROR }
            />
        </MaterialField>
        <br/>
        <MaterialField model="obcConfiguration.secureContext">
          <TextField
            hintText={strings.OBC_CONFIG_SECURE_CONTEXT_HT}
            floatingLabelText = {strings.OBC_CONFIG_SECURE_CONTEXT_FL}
            fullWidth={true}
            />
        </MaterialField>
        <MaterialField model="obcConfiguration.blocksPerPage">
          <TextField
            hintText={strings.OBC_CONFIG_BPP_HT}
            floatingLabelText = {strings.OBC_CONFIG_BPP_FL}
            fullWidth={true}
            />
        </MaterialField>
        <br/>
          {/*submit button for form*/}
          <div style={{textAlign: 'right'}}>
            <FlatButton label={strings.OBC_CONFIG_SUBMIT_LABEL} primary={true} type={"submit"}/>
          </div>
      </Form>
    );
  }
}

function mapStateToProps(state) {
  return { obcConfiguration: state.obcConfiguration, obcConfigurationForm: state.obcConfigurationForm };
}

export default connect(mapStateToProps)(ObcConfigurationForm);
