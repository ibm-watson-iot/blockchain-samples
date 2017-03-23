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

var hfc = require('fabric-client');
//var hfcca = require('fabric-ca');
var sdk = require("../cpsdk/cpsdk.js");
var utils = require('fabric-client/lib/utils.js');
var log4js = require('log4js');

var client = new hfc();
var logger = log4js.getLogger('CLI');
logger.setLevel('DEBUG');

var command = process.argv[2]
var func = process.argv[3]
var payload = process.argv[4]

logger.debug("%s %s(%s)", command.toUpperCase(), func, payload)

var user = require("./wallet.json")
var appc = sdk.newAppContext("IOT Contract Platform CLI", user);

switch (command) {
    case "deploy":
        result = appc.deploy(func, payload);
        logger.debug("Deploy result: ", result);
        break;
    case "test":
        testcli();
        break;
}

function testcli() {
    logger.debug("%s", JSON.stringify(appc, null, 4));
}