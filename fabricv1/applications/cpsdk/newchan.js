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

var log4js = require('log4js');
var logger = log4js.getLogger('DEPLOY');
logger.setLevel('DEBUG');

var hfc = require('fabric-client');
var utils = require('fabric-client/lib/utils.js');
var Peer = require('fabric-client/lib/Peer.js');
var Orderer = require('fabric-client/lib/Orderer.js');
var EventHub = require('fabric-client/lib/EventHub.js');

var creator = require('./creator.js');
var proposal = require('./proposal.js');
var sdk = require('./cpsdk');

var client = new hfc();

module.exports.newchan = function(appc) {
    // get new chain object
    var mychan = appc.myChannelID;
    logger.debug("Entered deploy -- about to call newChain with ", mychan);
    var chain = client.newChain(mychan);
    // add orderer to chain object
    var caRootsPath = appc.myOrdererCACertsPath;
    let data = fs.readFileSync(path.join(__dirname, caRootsPath));
    let caroots = Buffer.from(data).toString();
    logger.debug("calling addOrderer");
    chain.addOrderer(
        new Orderer(
            appc.myOrdererURI, {
                'pem': caroots,
                'ssl-target-name-override': appc.myOrdererName
            }
        )
    );
    // add keystore to chain object
    var kvstore = appc.myKeyValueStore
    logger.debug("setup KV store: " + kvstore);
    hfc.newDefaultKeyValueStore({
            path: kvstore,
        }).then((store) => {
            client.setStateStore(store);
            logger.debug("calling getCreator");
            return creator.getCreator(client, appc);
        }).then(
            (admin) => {
                logger.info('Successfully enrolled user');
                the_user = admin;

                // readin the envelope to send to the orderer
                data = fs.readFileSync('./test/fixtures/channel/mychannel.tx');
                var request = {
                    envelope: data
                };
                // send to orderer
                return chain.createChannel(request);
            }, (err) => {
                t.fail('Failed to enroll user \'admin\'. ' + err);
                t.end();
            })
        .then((response) => {
            logger.debug(' response ::%j', response);

            if (response && response.status === 'SUCCESS') {
                t.pass('Successfully created the channel.');
                return sleep(5000);
            } else {
                t.fail('Failed to create the channel. ');
                t.end();
            }
        }, (err) => {
            t.fail('Failed to initialize the channel: ' + err.stack ? err.stack : err);
            t.end();
        })
        .then((nothing) => {
            t.pass('Successfully waited to make sure new channel was created.');
            t.end();
        }, (err) => {
            t.fail('Failed to sleep due to error: ' + err.stack ? err.stack : err);
            t.end();
        });
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}