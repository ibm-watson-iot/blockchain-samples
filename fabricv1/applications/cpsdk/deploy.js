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

var hfc = require('fabric-client');
var utils = require('fabric-client/lib/utils.js');
var Peer = require('fabric-client/lib/Peer.js');
var Orderer = require('fabric-client/lib/Orderer.js');
var EventHub = require('fabric-client/lib/EventHub.js');

var creator = require('./creator.js');
var proposal = require('./proposal.js');

var sdk = require('./cpsdk');

logger.setLevel('DEBUG');

var client = new hfc();
var chain;
var eventhub;
var tx_id = null;
var clientContext = {};

if (!process.env.GOPATH) {
    process.env.GOPATH = fabric.goPath;
}

module.exports.deploy = function(appc, func, arg) {
    var mychan = appc.myChannelID;
    logger.debug("Entered deploy -- about to call newChain with ", mychan);
    var chain = client.newChain(mychan, clientContext);
    logger.debug("calling addOrderer");
    chain.addOrderer(new Orderer(appc.myOrdererURI));
    logger.debug("Setting up event hub");
    var eventhub = new EventHub();
    eventhub.setPeerAddr(appc.myEndorserURIEvents);
    eventhub.connect();
    // for (var endorser in fabric.organizations[me.myORG].endorsers) {
    //     if (fabric.organizations[me.myORG].endorsers.hasOwnProperty(endorser)) {
    //         chain.addPeer(new Peer(endorser.requests));
    //     }
    // }
    logger.debug("adding endorser peer");
    chain.addPeer(new Peer(appc.myEndorserURIRequests));
    var kvstore = appc.myKeyValueStore
    logger.debug("setup KV store: " + kvstore);
    hfc.newDefaultKeyValueStore({
        path: kvstore,
    }).then(function(store) {
        client.setStateStore(store);
        logger.debug("calling getCreator");
        return creator.getCreator(client, appc);
    }).then(
        function(admin) {
            logger.info('Successfully obtained enrolled user to deploy the chaincode');
            var tx_id = proposal.getTxId();
            // send proposal to endorser
            var request = {
                chaincodePath: appc.myChaincodePath,
                chaincodeId: appc.myChaincodeID,
                fcn: func,
                args: [arg],
                chainId: appc.myChannelID,
                txId: tx_id,
                nonce: utils.getNonce(),
            };
            logger.debug('Sending instantiate proposal to endorser, request = \n' + JSON.stringify(request, null, 4));
            return chain.sendInstantiateProposal(request);
        }
    ).then(
        function(results) {
            logger.info('Successfully obtained proposal responses from endorsers');
            return proposal.processProposalResults(tx_id, eventhub, chain, results, 'deploy');
        }
    ).then(
        function(response) {
            if (response.status === 'SUCCESS') {
                logger.info('Successfully sent deployment transaction to the orderer.');
                return response
            } else {
                logger.error('Failed to order the deployment endorsement. Error code: ' + response.status);
                throw response
            }
        }
    ).catch(
        function(err) {
            eventhub.disconnect();
            logger.error(err.stack ? err.stack : err);
            return err
        }
    );
}