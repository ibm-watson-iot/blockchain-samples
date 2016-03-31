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

import {fetchAssetObjectModel} from '../actions/AssetActions'
import AssetContainer from './AssetContainer'

import ConfigurationDialog from './ConfigurationDialog'

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
    this.props.fetchAssetObjectModel()
  }

  render(){
    return(
      <div>

        <AppBar style={{position: 'fixed', zIndex: 9999, background: Colors.grey800}}
          title="IoT Blockchain Monitor"
          showMenuIconButton={false}
          zDepth={4}
          iconElementRight={<ConfigurationDialog/>}
        />

      <Grid style={{paddingTop: 100, marginBottom:30}}>
          <Row>
            <Col xs={12} md={6}>
              <AssetContainer />
            </Col>
            <Col xs={12} md={6}>
              <Blockchain />
            </Col>
          </Row>
        </Grid>


    </div>
    )
  }
}

const mapDispatchToProps = (dispatch) => {
  return{
    fetchAssetObjectModel: () => {
      dispatch(fetchAssetObjectModel())
    }
  }
}

export default connect(
  null,
  mapDispatchToProps
)(App)
