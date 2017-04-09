/**
 * Copyright 2017 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

var util = require('util');
var hfc = require('fabric-client');
var utils = require('fabric-client/lib/utils.js');
var Peer = require('fabric-client/lib/Peer.js');
var Orderer = require('fabric-client/lib/Orderer.js');
var EventHub = require('fabric-client/lib/EventHub.js');
var log4js = require('log4js');
var fs = require('fs');
var path = require('path');

// for TLS
process.env.GRPC_SSL_CIPHER_SUITES = process.env.GRPC_SSL_CIPHER_SUITES ?
    process.env.GRPC_SSL_CIPHER_SUITES :
    'ECDHE-RSA-AES128-GCM-SHA256:' +
    'ECDHE-RSA-AES128-SHA256:' +
    'ECDHE-RSA-AES256-SHA384:' +
    'ECDHE-RSA-AES256-GCM-SHA384:' +
    'ECDHE-ECDSA-AES128-GCM-SHA256:' +
    'ECDHE-ECDSA-AES128-SHA256:' +
    'ECDHE-ECDSA-AES256-SHA384:' +
    'ECDHE-ECDSA-AES256-GCM-SHA384';

var newChan = require('./newchan.js');
//var joinChan = require('./joinchan.js');
var channel = require('./channel.js');
//var installCC = require('./installcc.js');
//var instantiateCC = require('./instantiatess.js');
//var invoke = require('./invoke.js');
//var query = require('./query.js');
//var deploy = require('./deploy.js');
var client = new hfc();
var logger = log4js.getLogger('IOTCPSDK');
logger.setLevel('DEBUG');


/**
 * @class
 */
class AppContext {
    /**
     * @param {string} name -- used to identify this app's instance
     */
    constructor(name, user) {
        var err
        if (typeof name === 'undefined' || !name) {
            logger.error(err = 'Failed to create AppContext. Missing required "name" parameter.');
            throw new Error(err);
        }
        this.name = name;
        if (typeof user === 'undefined' || !user) {
            logger.error(err = 'Failed to create AppContext. Missing required "user" parameter.');
            throw new Error(err);
        }
        this.user = user;

        this._myChain = null;

        logger.debug("IOTCP SDK Instantiated as [%s] with wallet [%s]", this.name, JSON.stringify(this.user).substr(0, 30) + "...");
    }

    // user based info
    get myUserName() { return this.user["name"] }
    get myUserSecret() { return this.user["secret"] }
    get myKeyValueStore() { return this.user["myKeyValueStore"] }
    get myORGName() { return this.user["myOrganization"]["name"] }

    // targeted fabric info
    get myOrdererName() { return this.user["myOrderer"]["name"] }
    get myOrdererURI() { return this.user["myOrderer"]["uri"] }
    get myOrdererCACertsPath() { return this.user["myOrderer"]["tls_cacerts"] }
    get myEndorserName() { return this.user["myOrganization"]["endorsers"][0]["name"] }
    get myEndorserURIRequests() { return this.user["myOrganization"]["endorsers"][0]["requests"] }
    get myEndorserURIEvents() { return this.user["myOrganization"]["endorsers"][0]["events"] }
    get myEndorserCACertsPath() { return this.user["myOrganization"]["endorsers"][0]["tls_cacerts"] }
    get myCAURI() { return this.user["myOrganization"]["ca"] }
    get myMSPID() { return this.user["myOrganization"]["mspid"] }

    // targeted channel info
    get myChannelID() { return this.user["myChannel"]["channelID"] }
    get myContractID() { return this.user["myContract"]["chaincodeID"] }
    get myContractPath() { return this.user["myContract"]["chaincodePath"] }
    get myContractID() { return this.user["myContract"]["chaincodeID"] }

    get myChain() { 
        if ( this._myChain === null) {
            this._myChain = client.newChain(this.myChannelID);
        }
        return this._myChain;
    }

    // methods for channel manipulation and monitoring
    chanHeight() { return channel.chanHeight(this) }
    blocks(from, to) { return channel.blocks(this, from, to) }
    newChan(channame) {
        return channel.newchan(this, channame);
    }

}

module.exports.newAppContext = function(name, user) {
    return new AppContext(name, user);
}

module.exports.getQProp = function(obj, str) {
    return str.split(".").reduce(function(o, x) { return o[x] }, obj);
}