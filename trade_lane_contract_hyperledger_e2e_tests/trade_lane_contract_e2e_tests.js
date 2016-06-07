/*******************************************************************************
Copyright (c) 2016 IBM Corporation.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Keerthi Challabotla - Initial Contribution
Carole Corley - Hyperledger update

******************************************************************************/

// Test for IoT Blockchain Tradelane Contract Hyperledger version - 4.0


var q			    = require('q');
var chai 	  	= require('chai');
var should 		= chai.should();
var assert		= require("assert");
var request		= require('request');

var logger		= require('./logger').createLogger();
var internal  = require('./index');
var Config    = require('./config');

describe(" Testing the Tradelane Contract against the obc-peer running on Bluemix ", function() {
	this.timeout(50000);

	var obcHost        = Config.obcHost;
	var obcPort 	     = Config.obcPort;
  var secure_context = Config.secure_context;
  var enroll_secret  = Config.enroll_secret;
  var protocol       = Config.protocol;
	var contract_path	 = Config.contract_path;
	var obcUrl 		     = protocol + "://" + obcHost + ":" + obcPort;
	var deployUrl 	   = obcUrl + "/chaincode";
	var invokeUrl 	   = obcUrl + "/chaincode";
	var queryUrl	     = obcUrl + "/chaincode";
  var registrarUrl   = obcUrl + "/registrar";

	var contractId;
	var messageId;
	var contract_version = "4.0";
	var timeout = 3000;

//-------------------------------------------------------------------------------------------//

    it("Registrar : ", function(complete) {

		var registrarBody = {
			"enrollId": secure_context,
      "enrollSecret": enroll_secret
		};

		var options = {
			url : registrarUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(registrarBody)
		};

		logger.info("Registrar...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to registrar: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("OK");
			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});

//-------------------------------------------------------------------------------------------//

	it(" Should deploy the contract  version : " + contract_version, function(complete) {
		var contractArgs = {
			version : contract_version
		};

		var deployBody = {
			jsonrpc : "2.0",
			method : "deploy",
			params : {
				type : 1,
				chaincodeID : {
          path: contract_path
				},
				ctorMsg : {
					function : "init",
					args : [JSON.stringify(contractArgs)]
				},
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : deployUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(deployBody)
		};

		logger.info("Deploy demo contract...");
		logger.info("path="+contract_path);
		logger.info("url="+deployUrl);
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to deploy demo contract: " + result.error);
			}

			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			contractId = resultBody.message;

			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});

//-------------------------------------------------------------------------------------------//


	it(" Should delete all assets and read all assets", function(complete) {
		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
			  type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "deleteAllAssets",
			    args : [  ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody)
		};

		logger.info("Delete all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	  .then(function (result) {
		  if (result.error) {
			  throw new Error("Failed to invoke the contract: " + result.error);
		  }

		  result.response.statusCode.should.be.equal(200);

		  var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			messageId = resultBody.message;

		  return q.delay(timeout)
	  })
	  .then(function () {

		  var readBody = {
				jsonrpc : "2.0",
				method : "query",
				params : {
				  type : 1,
				  chaincodeID : {
				    name : contractId
				  },
				  ctorMsg : {
				    function : "readAllAssets",
				    args : [  ]
				  },
	        "secureContext" : secure_context
				},
				id : 1234
			};

		var options = {
				url : queryUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(readBody)
			};

		logger.info("Read all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.be.instanceof(Array).and.have.lengthOf(0);
			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
	});


//-------------------------------------------------------------------------------------------//

	it(" Should create asset and read all assets ", function(complete) {

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
			  type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "createAsset",
			    args : [ "{\"assetID\":\"ASSET1\",\"temperature\":0.1}"]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody) //converts a value to JSON notation
		};

		logger.info("Creating Asset...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
		  if (result.error) {
			  throw new Error("Failed to invoke the contract: " + result.error);
		  }

		  result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			messageId = resultBody.message;

		  return q.delay(timeout)
		})
		.then(function () {

		  var readBody = {
				jsonrpc : "2.0",
			  method : "query",
			  params : {
			    type : 1,
			    chaincodeID: {
			      name: contractId
			    },
			    ctorMsg: {
			      function: "readAllAssets",
			      args: [  ]
			    },
          "secureContext" : secure_context
				},
				id : 1234
			};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Read all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}

			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.be.instanceof(Array).and.have.lengthOf(1);

			var arrayBody  = message[0];
			arrayBody.should.have.property("alerts");
			arrayBody['alerts'].should.have.property("active");
			arrayBody['alerts'].should.have.property("cleared");
			arrayBody['alerts'].should.have.property("raised");
			arrayBody.should.have.property("assetID");
			arrayBody.should.have.property("lastEvent");
      arrayBody['lastEvent'].should.have.property("args");
      var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
      eventBody.should.have.property("assetID");
      eventBody['assetID'].should.equal("ASSET1");
      eventBody.should.have.property("temperature");
      eventBody['temperature'].should.equal(0.1);
			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});

//-------------------------------------------------------------------------------------------//

		it(" Should create asset and read all assets ", function(complete) {

			var createBody = {
				jsonrpc : "2.0",
			  method : "invoke",
			  params : {
			    type : 1,
				  chaincodeID : {
				    name : contractId
				  },
				  ctorMsg : {
				    function : "createAsset",
				    args : [ "{\"assetID\":\"ASSET7\",\"temperature\":0}"]
				  },
          "secureContext" : secure_context
				},
				id : 1234
			};

			var options = {
				url : invokeUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(createBody) //converts a value to JSON notation
			};

			logger.info("Creating Asset...");
			var req = request.defaults(options);
			return internal.doRequest(null, "POST", req)
			.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}

			result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			messageId = resultBody.message;

			return q.delay(timeout)
			})

			.then(function () {

			var readBody = {
				jsonrpc : "2.0",
			  method : "query",
			  params : {
			    type : 1,
				  chaincodeID: {
				    name: contractId
				  },
				  ctorMsg: {
				    function: "readAllAssets",
				    args: [ ]
				  },
          "secureContext":secure_context
				},
				id : 1234
			};

			var options = {
				url : queryUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(readBody)
			};

			logger.info("Read all assets...");
			var req = request.defaults(options);
			return internal.doRequest(null, "POST", req)
			})
			.then(function (result) {
				if (result.error) {
					throw new Error("Failed to invoke the contract: " + result.error);
				}

        result.response.statusCode.should.be.equal(200);
				var body = JSON.parse(result.body);
				body.should.have.property("result");
				var resultBody = body.result;
				resultBody.should.have.property("status");
				resultBody.status.should.equal("OK");
				resultBody.should.have.property("message");
				var message = JSON.parse(resultBody.message);
				message.should.be.instanceof(Array).and.have.lengthOf(2);

				var arrayBody  = message[0];
			  arrayBody.should.have.property("alerts");
			  arrayBody['alerts'].should.have.property("active");
			  arrayBody['alerts'].should.have.property("cleared");
				arrayBody['alerts'].should.have.property("raised");
				arrayBody.should.have.property("assetID");
				arrayBody.should.have.property("lastEvent");
        arrayBody['lastEvent'].should.have.property("args");
        var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
        eventBody.should.have.property("assetID");
        eventBody['assetID'].should.equal("ASSET1");
        eventBody.should.have.property("temperature");
        eventBody['temperature'].should.equal(0.1);

        var arrayBody  = message[1];
				arrayBody.should.have.property("assetID");
				arrayBody.should.have.property("lastEvent");
        arrayBody['lastEvent'].should.have.property("args");
      	var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
        eventBody.should.have.property("assetID");
        eventBody['assetID'].should.equal("ASSET7");
        eventBody.should.have.property("temperature");
        eventBody['temperature'].should.equal(0);

				complete();
			})
			.catch(function (error) {
				//assert.fail(null, null, error);
				complete(error);
			})
			.done()
		});

//-------------------------------------------------------------------------------------------//

		it(" Should create asset and read all assets ", function(complete) {

			var createBody = {
				jsonrpc : "2.0",
			  method : "invoke",
			  params : {
			    type : 1,
				  chaincodeID : {
				    name : contractId
				  },
				  ctorMsg : {
				    function : "createAsset",
				    args : [ "{\"assetID\":\"ASSET8\",\"temperature\":-4}"]
				  },
          "secureContext" : secure_context
				},
				id : 1234
			};

			var options = {
				url : invokeUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(createBody) //converts a value to JSON notation
			};

			logger.info("Creating Asset...");
			var req = request.defaults(options);
			return internal.doRequest(null, "POST", req)
			.then(function (result) {
				if (result.error) {
					throw new Error("Failed to invoke the contract: " + result.error);
				}

				result.response.statusCode.should.be.equal(200);

				var body = JSON.parse(result.body);
				body.should.have.property("result");
				var resultBody = body.result;
				resultBody.should.have.property("status");
				resultBody.status.should.equal("OK");
				resultBody.should.have.property("message");
				messageId = resultBody.message;

				return q.delay(timeout)
			})
			.then(function () {

				var readBody = {
					jsonrpc : "2.0",
			  	method : "query",
			  	params : {
			    	type : 1,
				    chaincodeID: {
				      name: contractId
				    },
				    ctorMsg: {
				      function: "readAllAssets",
				      args: [  ]
				    },
            "secureContext" : secure_context
					},
					id : 1234
				};

				var options = {
					url : queryUrl,
					headers : {'Content-Type':'application/json'},
					body : JSON.stringify(readBody)
				};

				logger.info("Read all assets...");
				var req = request.defaults(options);
				return internal.doRequest(null, "POST", req)
			})
			.then(function (result) {
				if (result.error) {
					throw new Error("Failed to invoke the contract: " + result.error);
				}

				result.response.statusCode.should.be.equal(200);
				var body = JSON.parse(result.body);
				body.should.have.property("result");
				var resultBody = body.result;
				resultBody.should.have.property("status");
				resultBody.status.should.equal("OK");
				resultBody.should.have.property("message");
				var message = JSON.parse(resultBody.message);
				message.should.be.instanceof(Array).and.have.lengthOf(3);

				var arrayBody  = message[0];
				arrayBody.should.have.property("alerts");
				arrayBody['alerts'].should.have.property("active");
				arrayBody['alerts'].should.have.property("cleared");
				arrayBody['alerts'].should.have.property("raised");
				arrayBody.should.have.property("assetID");
				arrayBody.should.have.property("lastEvent");
        arrayBody['lastEvent'].should.have.property("args");
        var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
        eventBody.should.have.property("assetID");
        eventBody['assetID'].should.equal("ASSET1");
        eventBody.should.have.property("temperature");
        eventBody['temperature'].should.equal(0.1);

        var arrayBody  = message[1];
				arrayBody.should.have.property("assetID");
				arrayBody.should.have.property("lastEvent");
        arrayBody['lastEvent'].should.have.property("args");
        var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
        eventBody.should.have.property("assetID");
        eventBody['assetID'].should.equal("ASSET7");
        eventBody.should.have.property("temperature");
        eventBody['temperature'].should.equal(0);

        var arrayBody  = message[2];
				arrayBody.should.have.property("assetID");
				arrayBody.should.have.property("lastEvent");
        arrayBody['lastEvent'].should.have.property("args");
        var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
        eventBody.should.have.property("assetID");
        eventBody['assetID'].should.equal("ASSET8");
        eventBody.should.have.property("temperature");
        eventBody['temperature'].should.equal(-4);

				complete();
			})
			.catch(function (error) {
				//assert.fail(null, null, error);
				complete(error);
			})
			.done()
		});


//-------------------------------------------------------------------------------------------//

    it(" Should read particular asset ", function(complete) {

			var readBody = {
				jsonrpc : "2.0",
				method : "query",
				params : {
					type : 1,
					chaincodeID: {
				  	name: contractId
				  },
				  ctorMsg: {
				    function: "readAsset",
				    args: [ "{\"assetID\":\"ASSET7\"}" ]
				  },
            "secureContext" : secure_context
					},
					id : 1234
				};

			var options = {
				url : queryUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(readBody)
			};

			logger.info("Reading asset...")
			var req = request.defaults(options);
			return internal.doRequest(null, "POST", req)
			.then(function (result) {
				if (result.error) {
					throw new Error("Failed to invoke the contract: " + result.error);
				}
				result.response.statusCode.should.be.equal(200);
				var body = JSON.parse(result.body);
				body.should.have.property("result");
				var resultBody = body.result;
				resultBody.should.have.property("status");
				resultBody.status.should.equal("OK");
				resultBody.should.have.property("message");
				var message = JSON.parse(resultBody.message);

			  message.should.have.property("assetID");
        message.assetID.should.equal("ASSET7");
				message.should.have.property("lastEvent");
				var eventBody = message.lastEvent;
				eventBody.should.have.property("args");
        eventBody.should.have.property("function");
				var assetBody = JSON.parse(eventBody['args']);
				assetBody.should.have.property("assetID");
				assetBody['assetID'].should.equal("ASSET7");
        assetBody.should.have.property("temperature");
				assetBody['temperature'].should.equal(0);
				eventBody['function'].should.equal("createAsset");

				complete();
			})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});

//-------------------------------------------------------------------------------------------//

it(" Should read contract state ", function(complete) {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
				chaincodeID: {
				  name: contractId
				},
				ctorMsg: {
				  function: "readContractState",
				  args: [  ]
				},
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Reading Contract State...")
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.have.property("version");
			message.version.should.equal(contract_version);

			message.should.have.property("nickname");
			message.nickname.should.equal("TRADELANE");

			message.should.have.property("activeAssets");
			message.activeAssets.should.have.property("ASSET1");
			should.equal(message.activeAssets.ASSET1,true);
			message.activeAssets.should.have.property("ASSET7");
			should.equal(message.activeAssets.ASSET7,true);
			message.activeAssets.should.have.property("ASSET8");
			should.equal(message.activeAssets.ASSET8,true);

			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});

//-------------------------------------------------------------------------------------------//


it(" Should read Asset Schemas ", function(complete) {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
				chaincodeID: {
			  	name: contractId
				},
				ctorMsg: {
			  	function: "readAssetSchemas",
			  	args: [  ]
				},
      	"secureContext" : secure_context
			},
			id : 1234
	  };

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Reading Asset Schemas...")
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.have.property("API");
			message.should.have.property("objectModelSchemas");

			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});


//-------------------------------------------------------------------------------------------//


it(" Should read Asset Samples ", function(complete) {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			  chaincodeID: {
			    name: contractId
			  },
			  ctorMsg: {
			    function: "readAssetSamples",
			    args: [  ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Reading Asset Samples...")
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.have.property("contractState");
			message.should.have.property("event");
			message.should.have.property("initEvent");
			message.should.have.property("state");

			complete();
		})
		.catch(function (error) {
			//assert.fail(null, null, error);
			complete(error);
		})
		.done()
	});


//-------------------------------------------------------------------------------------------//

it(" Should Read the contract object model ", function(complete) {

	var readBody = {
		jsonrpc : "2.0",
		method : "query",
		params : {
			type : 1,
			chaincodeID : {
				name : contractId
			},
			ctorMsg : {
				function : "readContractObjectModel",
				args : []
			},
      "secureContext" : secure_context
		},
		id : 1234
	};

	var options = {
		url : queryUrl,
		headers : {'Content-Type':'application/json'},
		body : JSON.stringify(readBody)
	};
	logger.info("Reading Contract object model...");
	var req = request.defaults(options);
	return internal.doRequest(null, "POST", req)
	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to Read demo contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);
		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		var message = JSON.parse(resultBody.message);
		message.should.have.property("version");
		message.version.should.equal(contract_version);
		message.should.have.property("nickname");
		message.nickname.should.equal("TRADELANE");
		message.should.have.property("activeAssets");
		//body['OK']['activeAssets'].should.equal('{}');

		complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//

it(" Should Read the recent states ", function(complete) {

	var readBody = {
		jsonrpc : "2.0",
		method : "query",
		params : {
			type : 1,
			chaincodeID : {
				name : contractId
			},
			ctorMsg : {
				function : "readRecentStates",
				args : []
			},
    	"secureContext" : secure_context
		},
		id : 1234
	};

	var options = {
		url : queryUrl,
		headers : {'Content-Type':'application/json'},
		body : JSON.stringify(readBody)
	};

	logger.info("Reading recent states...");
	var req = request.defaults(options);
	return internal.doRequest(null, "POST", req)
	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to Read demo contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);
		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		var message = JSON.parse(resultBody.message);
		message.should.be.instanceof(Array).and.have.lengthOf(3);

		var arrayBody  = message[0];
		arrayBody.should.have.property("assetID");
		arrayBody.should.have.property("lastEvent");
    arrayBody['lastEvent'].should.have.property("args");
    var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
    eventBody.should.have.property("assetID");
    eventBody['assetID'].should.equal("ASSET8");
    eventBody.should.have.property("temperature");
    eventBody['temperature'].should.equal(-4);

    var arrayBody  = message[1];
		arrayBody.should.have.property("assetID");
		arrayBody.should.have.property("lastEvent");
    arrayBody['lastEvent'].should.have.property("args");
    var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
    eventBody.should.have.property("assetID");
    eventBody['assetID'].should.equal("ASSET7");
    eventBody.should.have.property("temperature");
    eventBody['temperature'].should.equal(0);

  	var arrayBody  = message[2];
		arrayBody.should.have.property("alerts");
		arrayBody['alerts'].should.have.property("active");
		arrayBody['alerts'].should.have.property("cleared");
		arrayBody['alerts'].should.have.property("raised");
		arrayBody.should.have.property("assetID");
		arrayBody.should.have.property("lastEvent");
    arrayBody['lastEvent'].should.have.property("args");
    var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
    eventBody.should.have.property("assetID");
    eventBody['assetID'].should.equal("ASSET1");
    eventBody.should.have.property("temperature");
  	eventBody['temperature'].should.equal(0.1);

		complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//

it(" Should Read the Asset History ", function(complete) {

	var readBody = {
		jsonrpc : "2.0",
		method : "query",
		params : {
			type : 1,
			chaincodeID : {
				name : contractId
			},
			ctorMsg : {
				function : "readAssetHistory",
				args : [ "{\"assetID\":\"ASSET8\"}" ]
			},
      "secureContext" : secure_context
		},
		id : 1234
	};

	var options = {
		url : queryUrl,
		headers : {'Content-Type':'application/json'},
		body : JSON.stringify(readBody)
	};

	logger.info("Reading the asset history...");
	var req = request.defaults(options);
	return internal.doRequest(null, "POST", req)
	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to Read demo contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);
		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		var message = JSON.parse(resultBody.message);
		message.should.be.instanceof(Array).and.have.lengthOf(1);

		var arrayBody  = message[0];
		arrayBody.should.have.property("assetID");
		arrayBody.should.have.property("lastEvent");
    arrayBody['lastEvent'].should.have.property("args");
    var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
    eventBody.should.have.property("assetID");
    eventBody['assetID'].should.equal("ASSET8");
    eventBody.should.have.property("temperature");
    eventBody['temperature'].should.equal(-4);

		complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//

it(" Should update an asset and read asset - Bad temp scenario", function(complete) {
	var timestamp = new Date();

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "updateAsset",
			    args : [ "{\"assetID\":\"ASSET8\",\"temperature\":3}" ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
				url : invokeUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(createBody)
			};

		logger.info("Updating asset...");

		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)

	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}
		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

		return q.delay(timeout)
	})

	.then(function () {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			    chaincodeID: {
			      name: contractId
			    },
			    ctorMsg: {
			      function: "readAsset",
			      args: [ "{\"assetID\":\"ASSET8\"}" ]
			    },
          "secureContext" : secure_context
				},
				id : 1234
			};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Bad temperature - Reading an asset...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);

      message.should.have.property("alerts");
      message.alerts.should.have.property("active");

      message.alerts.active.should.be.instanceof(Array).and.have.lengthOf(1);
      should.equal(message.alerts.active[0], "OVERTEMP");

      message.alerts.should.have.property("raised");
      message.alerts.raised.should.be.instanceof(Array).and.have.lengthOf(1);
      should.equal(message.alerts.raised[0], "OVERTEMP");

      message.alerts.should.have.property("cleared");
      message.alerts.cleared.should.be.instanceof(Array).and.have.lengthOf(0);

			message.should.have.property("assetID");
      message.assetID.should.equal("ASSET8");
			message.should.have.property("lastEvent");
			var eventBody = message.lastEvent;
			eventBody.should.have.property("args");
			var assetBody = JSON.parse(eventBody['args']);
			assetBody.should.have.property("assetID");
			assetBody['assetID'].should.equal("ASSET8");
			assetBody.should.have.property("temperature");
			assetBody['temperature'].should.equal(3);

			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//

it(" Should update an asset and read asset - Good temp scenario", function(complete) {
	var timestamp = new Date();

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "updateAsset",
			    args : [ "{\"assetID\":\"ASSET8\",\"temperature\":-3}" ]
			  },
      	"secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody)
		};

		logger.info("Updating asset...");

		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)

	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}
		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

		return q.delay(timeout)
	})

	.then(function () {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			  chaincodeID: {
			    name: contractId
			  },
			  ctorMsg: {
			    function: "readAsset",
			    args: [ "{\"assetID\":\"ASSET8\"}" ]
			  },
        "secureContext":secure_context
			},
			id : 1234
		};

		var options = {
				url : queryUrl,
				headers : {'Content-Type':'application/json'},
				body : JSON.stringify(readBody)
			};

		logger.info("Good temperature - Reading an asset...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);
			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);

      message.should.have.property("alerts");
      message.alerts.should.have.property("active");
			message.should.have.property("assetID");
      message.assetID.should.equal("ASSET8");
			message.should.have.property("lastEvent");
			var eventBody = message.lastEvent;
			eventBody.should.have.property("args");
			var assetBody = JSON.parse(eventBody['args']);
			assetBody.should.have.property("assetID");
			assetBody['assetID'].should.equal("ASSET8");
			assetBody.should.have.property("temperature");
			assetBody['temperature'].should.equal(-3);

			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//


it(" Should update an asset and read asset", function(complete) {
		var timestamp = new Date();

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "updateAsset",
			    args : [ "{\"assetID\":\"ASSET8\",\"location\":{\"latitude\":49,\"longitude\":-97},\"carrier\":\"UPS\"}" ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody)
		};

		logger.info("Updating asset...");

		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)

	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}
		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

		return q.delay(timeout)
	})

	.then(function () {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			  chaincodeID: {
			    name: contractId
			  },
			  ctorMsg: {
			    function: "readAsset",
			    args: [ "{\"assetID\":\"ASSET8\"}" ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Reading an asset...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);

			message.should.have.property("assetID");
      message.should.have.property("carrier");
			message.should.have.property("lastEvent");
			var eventBody = message.lastEvent;
			eventBody.should.have.property("args");
			var assetBody = JSON.parse(eventBody['args']);
			assetBody.should.have.property("assetID");
			assetBody['assetID'].should.equal("ASSET8");
			assetBody.should.have.property("location");
			assetBody['location'].should.have.property("latitude");
			assetBody['location'].should.have.property("longitude");
			assetBody['location']['latitude'].should.equal(49);
			assetBody['location']['longitude'].should.equal(-97);
			assetBody.should.have.property("carrier");
			assetBody['carrier'].should.equal("UPS");

			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
});

//-------------------------------------------------------------------------------------------//


	it(" Should delete an asset and read all assets", function(complete) {
		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "deleteAsset",
			    args : [ "{\"assetID\":\"ASSET8\"}" ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody)
		};

		logger.info("Delete all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

		return q.delay(timeout)
	})

	.then(function () {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			  chaincodeID: {
			    name: contractId
			  },
			  ctorMsg: {
			    function: "readAllAssets",
			    args: [  ]
			  },
        "secureContext":secure_context
			},
			id : 1234
		};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Read all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.be.instanceof(Array).and.have.lengthOf(2);

			var arrayBody  = message[0];
			arrayBody.should.have.property("alerts");
			arrayBody['alerts'].should.have.property("active");
			arrayBody['alerts'].should.have.property("cleared");
			arrayBody['alerts'].should.have.property("raised");
			arrayBody.should.have.property("assetID");
			arrayBody.should.have.property("lastEvent");
      arrayBody['lastEvent'].should.have.property("args");
      var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
      eventBody.should.have.property("assetID");
      eventBody['assetID'].should.equal("ASSET1");
      eventBody.should.have.property("temperature");
      eventBody['temperature'].should.equal(0.1);

      var arrayBody  = message[1];
			arrayBody.should.have.property("assetID");
			arrayBody.should.have.property("lastEvent");
      arrayBody['lastEvent'].should.have.property("args");
      var eventBody = JSON.parse(arrayBody['lastEvent']['args']);
      eventBody.should.have.property("assetID");
      eventBody['assetID'].should.equal("ASSET7");
      eventBody.should.have.property("temperature");
      eventBody['temperature'].should.equal(0);

			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
	});

//-------------------------------------------------------------------------------------------//

	it(" Should delete all asset and read all assets", function(complete) {
		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "deleteAllAssets",
			    args : [  ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody)
		};

		logger.info("Delete all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

		return q.delay(timeout)
	})

	.then(function () {

		var readBody = {
			jsonrpc : "2.0",
			method : "query",
			params : {
				type : 1,
			  chaincodeID: {
			    name: contractId
			  },
			  ctorMsg: {
			    function: "readAllAssets",
			    args: [  ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : queryUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(readBody)
		};

		logger.info("Read all assets...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
	})
		.then(function (result) {
			if (result.error) {
				throw new Error("Failed to invoke the contract: " + result.error);
			}
			result.response.statusCode.should.be.equal(200);

			var body = JSON.parse(result.body);
			body.should.have.property("result");
			var resultBody = body.result;
			resultBody.should.have.property("status");
			resultBody.status.should.equal("OK");
			resultBody.should.have.property("message");
			var message = JSON.parse(resultBody.message);
			message.should.be.instanceof(Array).and.have.lengthOf(0);

			complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
	});

//-------------------------------------------------------------------------------------------//

	it(" Invoking setCreateOnUpdate ", function(complete) {

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "setCreateOnUpdate",
			    args : [ "{\"createOnUpdate\":false}"]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody) //converts a value to JSON notation
		};

		logger.info("Setting CreateOnUpdate to false...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}

		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

    complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
	});

//-------------------------------------------------------------------------------------------//

	it(" Invoking setLoggingLevel ", function(complete) {

		var createBody = {
			jsonrpc : "2.0",
			method : "invoke",
			params : {
				type : 1,
			  chaincodeID : {
			    name : contractId
			  },
			  ctorMsg : {
			    function : "setLoggingLevel",
			    args : [ "{\"logLevel\":\"DEBUG\"}" ]
			  },
        "secureContext" : secure_context
			},
			id : 1234
		};

		var options = {
			url : invokeUrl,
			headers : {'Content-Type':'application/json'},
			body : JSON.stringify(createBody) //converts a value to JSON notation
		};

		logger.info("Setting LoggingLevel to DEBUG...");
		var req = request.defaults(options);
		return internal.doRequest(null, "POST", req)
		.then(function (result) {
		if (result.error) {
			throw new Error("Failed to invoke the contract: " + result.error);
		}
		result.response.statusCode.should.be.equal(200);

		var body = JSON.parse(result.body);
		body.should.have.property("result");
		var resultBody = body.result;
		resultBody.should.have.property("status");
		resultBody.status.should.equal("OK");
		resultBody.should.have.property("message");
		messageId = resultBody.message;

    complete();
	})
	.catch(function (error) {
		//assert.fail(null, null, error);
		complete(error);
	})
	.done()
	});

});
