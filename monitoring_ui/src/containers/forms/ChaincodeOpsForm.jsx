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

import SelectField from 'material-ui/lib/select-field';
import MenuItem from 'material-ui/lib/menus/menu-item';


class ChaincodeOpsForm extends React.Component {

  handleChange(e,i,v) {

    //dispatch an action to change the form model specifically for the form under a particular tab.
    this.props.dispatch(actions.change('chaincodeOpsForm.'+this.props.tab+'.selectedFn',i));
  }

  render() {
    let { chaincodeOpsForm, tab } = this.props;

    return (
      <Form model="chaincodeOpsForm" noValidate
        onSubmit={(chaincodeOpsForm) => this.handleSubmit(chaincodeOpsForm)}>
        <SelectField style={{marginLeft: 10}}value={chaincodeOpsForm[tab] ? chaincodeOpsForm[tab].selectedFn : ""} onChange={(e,i,v) => {this.handleChange(e,i,v)}}>
          {chaincodeOpsForm[tab] && chaincodeOpsForm[tab].fns.map(function(fn, index){
              return(
                <MenuItem value={index} primaryText={fn.name} key={index}/>
              )
          })}
        </SelectField>
      </Form>
    );
  }
}

function mapStateToProps(state) {
  return {
    chaincodeOpsForm: state.chaincodeOpsForm,
    currentOpsTab: state.chaincode.ui.currentTab
  };
}

export default connect(mapStateToProps)(ChaincodeOpsForm);
