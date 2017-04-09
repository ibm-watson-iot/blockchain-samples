/**
 * Copyright 2016 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the 'License');
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an 'AS IS' BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */
'use strict';

var log4js = require('log4js');
var logger = log4js.getLogger('Creator');

var path = require('path');
var util = require('util');

var User = require('fabric-client/lib/User.js');
var utils = require('fabric-client/lib/utils.js');
var copService = require('fabric-ca-client/lib/FabricCAClientImpl.js');

logger.setLevel('DEBUG');

module.exports.getCreator = function(client, appc) {
    var username = appc.myUserName;
    var password = appc.myUserSecret;
    var member;
    return client.getUserContext(username)
        .then((user) => {
            if (user && user.isEnrolled()) {
                logger.info('Successfully loaded member from persistence');
                return user;
            } else {
				// Need to enroll it with CA server
				var tlsOptions = {
					trustedRoots: [appc.certificate],
					verify: false
				};
                var ca_client = new copService(appc.myCAURI);
                // need to enroll it with CA server
                return ca_client.enroll({
                    enrollmentID: username,
                    enrollmentSecret: password
                }).then((enrollment) => {
                    logger.info('Successfully enrolled user \'' + username + '\'');

                    member = new User(username, client);
                    return member.setEnrollment(enrollment.key, enrollment.certificate);
                }).then(() => {
                    return client.setUserContext(member);
                }).then(() => {
                    return member;
                }).catch((err) => {
                    logger.error('Failed to enroll and persist user. Error: ' + err.stack ? err.stack : err);
                    throw new Error('Failed to obtain an enrolled user');
                });
            }
        });
};