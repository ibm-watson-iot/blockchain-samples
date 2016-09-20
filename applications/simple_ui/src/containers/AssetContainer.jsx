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

import {setAssetIdInput, fetchAsset, removeAssetFromTracking, toggleShowRemoveBtn, showRemoveBtn, hideRemoveBtn, clearTrackedAssets, setAssetInputErrMsg} from '../actions/AssetActions'
import AssetView from '../components/AssetView'

/**
The AssetContainer class contains a list of assets being tracked by the user.
Clicking the expand button on the individual asset will show the entire payload
contained within the response from the asset query.
**/
class AssetContainer extends React.Component{

  //objToJsx should be an empty array
  readObjProps = (obj, objToJsx, indents) => {

    //If the object itself is a primitive, we just return that as a string.
    //for example {"OK":100}
    if(obj && Object.prototype.toString.call(obj) !== '[object Object]'){
      return (<p>{obj.toString()}</p>);
    }

    for(var propertyName in obj){
      if (obj.hasOwnProperty(propertyName)) {
        if(Object.prototype.toString.call(obj[propertyName]) === '[object Object]'){
          //print the correct number of indentations for parent level object
          objToJsx.push(<p key={indents+propertyName} style={{textIndent: (indents * 20)}}> {propertyName + ": "} </p>);
          //to prettify the output, we should indent the nested objects
          indents ++;
          this.readObjProps(obj[propertyName], objToJsx, indents)
          //we finished going into the nested object, so remove one level of indents
          indents --;
          //objToJsx.push(<p key={indents+propertyName+"end"} style={{textIndent: (indents * 20)}}>}</p>);
        }else{
          objToJsx.push(<p key={indents+propertyName} style={{textIndent: (indents * 20)}}> {propertyName + ": " + obj[propertyName]}</p>);
        }
      }
    }

    return objToJsx;
  }

  handleOnChange = (e,i,v) => {
    this.props.setAssetIdInput(e.target.value);
  }

  handleAddAsset = (e,i,v) => {
    //verify that the asset doesn't already exist

    let assetIdInput = this.props.assetIdInput;

    let asset = this.props.assets.find(function(asset){
      return asset.data.assetID === assetIdInput
    })

    //if we found an asset, it means we are already tracking it
    if(asset){
      this.props.setAssetInputErrMsg(this.props.assetIdInput + " is already being tracked.")
    }else{
      //otherwise, add this asset to the tracking list.
      this.props.fetchAsset(this.props.assetIdInput);
      this.props.setAssetIdInput("");
    }
  }

  render(){
    return(
      <Paper style={{marginBottom:20}}>
        <AppBar
          title={<div>Assets for Contract: <span title={this.props.chaincodeId}>{this.props.chaincodeId.substring(0,7)}...</span></div>}
          showMenuIconButton={false}
          iconElementRight={<FlatButton label="Reset" onTouchTap={this.props.clearTrackedAssets}/>}
        />

      <div style = {{paddingLeft: 8, paddingRight: 8}}>
        <TextField
          hintText="Asset-123"
          floatingLabelText = "Asset ID"
          type="text"
          onChange={this.handleOnChange}
          value={this.props.assetIdInput}
          style={{width:'calc(100% - 100px)'}}
          errorText={this.props.errorMsg}
          />

          {/*submit button for form*/}
          <FlatButton label="Submit" primary={true} onClick={this.handleAddAsset}/>
      </div>

      {/*Iterate through all the assets being tracked and display them on the UI*/}
      {this.props.assets.map( (asset, index) => {
        //TODO: Remove this, it adds an extra object to test nesting
          //asset.data.location.location2 = {some: 1, other: 2}
          return <AssetView key={index} assetData={asset.data} removeFn={this.props.removeAsset} displayObj={this.readObjProps} showRemoveFn={this.props.showRemoveBtn} hideRemoveFn={this.props.hideRemoveBtn} index={index} showRemoveBtn={asset.showRemove}/>
        }
      )}
      </Paper>
    )
  }

}

const mapStateToProps = (state) =>{
  return {
    assetIdInput: state.asset.assetIdInput,
    assets: state.asset.assets,
    errorMsg: state.asset.assetInputErrMsg,
    chaincodeId: state.configuration.chaincodeId
  }
}

const mapDispatchToProps = (dispatch) =>{
  return{
    setAssetIdInput: (assetId) => {
      dispatch(setAssetIdInput(assetId))
    },
    fetchAsset: (assetId) => {
      dispatch(fetchAsset(assetId))
    },
    removeAsset: (assetId) => {
      dispatch(removeAssetFromTracking(assetId))
    },
    showRemoveBtn: (index) => {
      dispatch(showRemoveBtn(index))
    },
    hideRemoveBtn: (index) => {
      dispatch(hideRemoveBtn(index))
    },
    clearTrackedAssets: () => {
      dispatch(clearTrackedAssets())
    },
    setAssetInputErrMsg: (msg) => {
      dispatch(setAssetInputErrMsg(msg))
    }
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(AssetContainer)
