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
import { combineReducers } from 'redux'
import {blockchain} from './BlockchainReducer.js'
import {configuration} from './ConfigurationReducer.js'
import {asset} from './AssetReducer.js'

import {
  modelReducer,
  createFormReducer
} from 'react-redux-form';

const initialConfigurationState = {
  urlRestRoot: "http://localhost:3000",
  chaincodeId: "7cdd53526ed31f7be5249bfa42c4c73728edebddf91bd29720a289105dafbf1fd8c94306ba128800fc1c2bfbee618ce85717d35f88bb7b481ca3d3ada70d78fd",
  secureContext: "user_context",
  blocksPerPage: "10"
};

/**
Combines all other reducers into one reducer called the root reducer. We will be using the root
reducer when creating the redux store.
**/
const rootReducer = combineReducers({
  blockchain,
  configuration,
  asset,
  //obcConfiguration is the model that deals with any configuration related to obc
  obcConfiguration: modelReducer('obcConfiguration', initialConfigurationState),
})

export default rootReducer
