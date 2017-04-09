/**
 * Copyright 2016 IBM All Rights Reserved.
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

// This is an SDK for the IoT Contract Platform

'use strict';

var fs = require('fs');
var path = require('path');
var utils = require('fabric-client/lib/utils.js');
var log4js = require('log4js');
var logger = log4js.getLogger('CHANNEL');
logger.setLevel('DEBUG');

var hfc = require('fabric-client');
var Peer = require('fabric-client/lib/Peer.js');
var Orderer = require('fabric-client/lib/Orderer.js');
var EventHub = require('fabric-client/lib/EventHub.js');

var creator = require('./creator.js');
var proposal = require('./proposal.js');
var sdk = require('./cpsdk');

module.exports.chanHeight = (appc) => {
    logger.debug("Entered chanheight for", appc.myChannelID);
    return appc.myChain.queryInfo();
}

module.exports.blocks = (appc, from, to) => {
    logger.debug("Entered blocks for %s from block %d to block %d", appc.myChannelID, from, to);
    return appc.myChain.queryBlock(Number(from));
}

module.exports.newchan = (appc, chanName) => {
    var client = new hfc();
    if (chanName === null || chanName === "") {
        chanName = appc.myChannelID;
    }
    var chain = appc.myChain;
    // add orderer to chain object
    var caRootsPath = appc.myOrdererCACertsPath;
    let data = fs.readFileSync(path.join(__dirname, caRootsPath));
    let caroots = Buffer.from(data).toString();
    logger.debug("calling addOrderer");
    appc.myChain.addOrderer(
        new Orderer(
            appc.myOrdererURI, {
                'pem': caroots,
                'ssl-target-name-override': appc.myOrdererName
            }
        )
    );
	// Acting as a client in org1 when creating the channel
	var org = appc.myORGName;

	utils.setConfigSetting('key-value-store', 'fabric-client/lib/impl/FileKeyValueStore.js');
	return hfc.newDefaultKeyValueStore({
		path: appc.myKeyValueStore
	}).then((store) => {
		client.setStateStore(store);
		return creator.getCreator(client, appc);
	})
	.then((admin) => {
		logger.debug('Successfully enrolled user \'admin\'');
		the_user = admin;

		// readin the envelope to send to the orderer
		data = fs.readFileSync('../../test/fixtures/channel/mychannel.tx');
		var request = {
			envelope : data,
			name : chanName,
			orderer : orderer
		};
		// send to orderer
		return client.createChannel(request);
	}, (err) => {
		logger.debug('Failed to enroll user \'admin\'. ' + err);
		throw('Failed to enroll user \'admin\'. ' + err);
	})
	.then((chain) => {
		logger.debug(' response ::%j',chain);

		if (chain) {
			var test_orderers = chain.getOrderers();
			if(test_orderers) {
				var test_orderer = test_orderers[0];
				if(test_orderer === orderer) {
					logger.debug('Successfully created the channel.');
				}
				else {
					logger.debug('Chain did not have the orderer.');
                    throw('Chain did not have the orderer.');
				}
			}
			return sleep(5000);
		} else {
			logger.debug('Failed to create the channel. ');
			throw('Failed to create the channel. ');
		}
	}, (err) => {
		logger.debug('Failed to initialize the channel: ' + err.stack ? err.stack : err);
		throw('Failed to initialize the channel: ' + err.stack ? err.stack : err);
	})
	.then((nothing) => {
		logger.debug('Successfully waited to make sure new channel was created.');
	}, (err) => {
		logger.debug('Failed to sleep due to error: ' + err.stack ? err.stack : err);
		throw('Failed to sleep due to error: ' + err.stack ? err.stack : err);
	});
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}