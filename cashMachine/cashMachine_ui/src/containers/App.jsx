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
import React from 'react'
import Blockchain from './Blockchain'

import Grid from 'react-bootstrap/lib/Grid'
import Row from 'react-bootstrap/lib/Row'
import Col from 'react-bootstrap/lib/Col'

import { connect } from 'react-redux'

import Tabs from 'material-ui/lib/tabs/tabs';
import Tab from 'material-ui/lib/tabs/tab';
import AppBar from 'material-ui/lib/app-bar';

import getMuiTheme from 'material-ui/lib/styles/getMuiTheme';
import IotTheme from '../theme/theme.js';

import "../styles/main.css";
//import "../../node_modules/hint.css/hint.css"

import * as Colors from 'material-ui/lib/styles/colors';

import {fetchCcSchema} from '../actions/ChaincodeActions'
import ChaincodeOps from './ChaincodeOps'

import ConfigurationDialog from './ConfigurationDialog'

import ResponsePayloadContainer from './ResponsePayloadContainer'

import * as strings from '../resources/strings'

import {openSnackbar, hideSnackbar, setSnackbarMsg} from '../actions/AppActions';
import Snackbar from 'material-ui/lib/snackbar';

class App extends React.Component{

  static childContextTypes = {
    muiTheme : React.PropTypes.object,
  };

  //the key passed through context must be called "muiTheme"
  getChildContext() {
    return { muiTheme: getMuiTheme(IotTheme) }
  }

  //get the asset object model as soon as this component mounts
  componentDidMount(){
    //this.props.fetchAssetObjectModel()
    this.props.fetchCcSchema();
  }

  render(){
    return(
      <div>
        <AppBar style={{position: 'fixed', zIndex: 9999, background: Colors.grey800}}
          title={strings.APP_BAR_MAIN_TITLE_TEXT}
          showMenuIconButton={false}
          zDepth={4}
          iconElementRight={<ConfigurationDialog/>}
          fluid={true}
        />

        <Grid style={{paddingTop: 100, marginBottom:30}} fluid={true}>
            <Row>
              <Col xs={12} md={4}>
                  <ChaincodeOps />
              </Col>
              <Col xs={12} md={4}>
                  <ResponsePayloadContainer />
              </Col>
              <Col xs={12} md={4}>
                  <Blockchain />
              </Col>
            </Row>
          </Grid>

          <Snackbar
            open={this.props.snackbarIsOpen}
            message={this.props.snackbarMsg}
            autoHideDuration={4000}
            onRequestClose={()=>{this.props.hideSnackbar()}}
          />
      </div>
    )
  }
}

const mapStateToProps = (state) => {
  return {
    snackbarIsOpen: state.app.ui.snackbar.open,
    snackbarMsg: state.app.ui.snackbar.msg
  }
}

const mapDispatchToProps = (dispatch) => {
  return{
    fetchAssetObjectModel: () => {
      dispatch(fetchAssetObjectModel())
    },
    fetchCcSchema: () => {
      dispatch(fetchCcSchema())
    },
    hideSnackbar: () => {
      dispatch(hideSnackbar())
    }
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
