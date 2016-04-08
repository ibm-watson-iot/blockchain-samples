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
import { createStore, applyMiddleware, compose } from 'redux'
import thunkMiddleware from 'redux-thunk'
import createLogger from 'redux-logger'
import rootReducer from '../reducers/RootReducer'
import persistState from 'redux-localstorage'

const loggerMiddleware = createLogger()

/**
Configures the redux store and applies middleware.
**/
export default function configureStore(initialState) {
  return createStore(rootReducer, initialState, compose(
      applyMiddleware(thunkMiddleware, loggerMiddleware),
      persistState(["configuration","obcConfiguration"]),
      window.devToolsExtension ? window.devToolsExtension() : f => f
    ));

}
