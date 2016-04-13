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

import Tabs from 'material-ui/lib/tabs/tabs';
import Tab from 'material-ui/lib/tabs/tab';
import Slider from 'material-ui/lib/slider';

import ChaincodeOpsForm from './forms/ChaincodeOpsForm'

import {setCurrentTab} from '../actions/ChaincodeActions'

import JsonSchemaForm from './forms/JsonSchemaForm'

import * as strings from '../resources/strings'

var styles = {
  appBar: {
    flexWrap: 'wrap',
  },
  tabs: {
    width: '100%',
  },
};

/**
The ChaincodeOps container manages the generation of the dynamic forms that
display on the UI.
**/
class ChaincodeOps extends React.Component{

  handleChange = (value) => {
    if(typeof(value) === "string"){
      this.props.dispatch(setCurrentTab(value))
    }
  }

  render(){
    //let {currentTab, possibleTabs} = this.props
    //TODO: Only rerender when going to the particular tab

    return(
      <Paper style={{marginBottom:20}}>
        <AppBar title={strings.APP_BAR_TITLE_CHAINCODE_OPS} style={styles.appBar} showMenuIconButton={false}></AppBar>
          <Tabs value={this.props.currentTab} onChange={this.handleChange} style={styles.tabs}>
            {this.props.possibleTabs.map(function(tab){
              return(
                <Tab label={tab.name} value={tab.name} key={tab.name}>
                  {/*Need to indicate which tab the form belongs to, so we can reference the state.*/}
                  <ChaincodeOpsForm tab={tab.name}/>
                  <div style={{margin: 10}}>
                  <JsonSchemaForm tabRenderedTo={tab.name}/>
                  </div>
                </Tab>
              )
            })}
          </Tabs>
      </Paper>
    )
  }
}

const connectStateToProps = (state) => {
  return{
    possibleTabs : state.chaincode.ui.possibleTabs,
    currentTab : state.chaincode.ui.currentTab
  }
}

export default connect(connectStateToProps, null)(ChaincodeOps)
