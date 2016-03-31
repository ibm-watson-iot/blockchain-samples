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
const path = require('path');
const webpack = require('webpack');

//target for using npm run [target] in the terminal
//const TARGET = process.env.npm_lifecycle_event;

//process.env.BABEL_ENV = TARGET;

const PATHS = {
  app: path.join(__dirname, 'src'),
  build: path.join(__dirname, 'public'),
  lib: path.join(__dirname, 'lib')
}

module.exports = {
  entry:{
    app: PATHS.app
  },
  resolve: {
    extensions: ['','.js','.jsx']
  },
  output: {
    path: PATHS.build,
    filename: 'bundle.js'
  },
  module:{
    loaders:[
      {
        //use regex to test for js and jsx
        test: /\.jsx?$/,
        loaders: ['babel?cacheDirectory'],
        //only include files in the PATHS.app path
        include: PATHS.app
      },
      {
        test: /\.css$/,
        loaders: ['style','css'],
        include: [PATHS.app,PATHS.lib]
      }
    ]
  },
  devTool: 'eval-source-map',
  devServer:{
    contentBase: PATHS.build,
    historyApiFallback: true,
    hot: true,
    inline: true,
    progress: true,
    stats: 'errors-only',
    host: process.env.HOST,
    port: process.env.PORT || 8081
  },
  plugins:[
    new webpack.HotModuleReplacementPlugin()
  ]
}
