var http = require('http');
var fs = require("fs");
/*
REST API  call to the building sensors smart contract
 */

/**
 * Make an HTTP POST call to the REST API
 */
// do a POST request
// create the JSON object

///// Read data from configration file
 var config = fs.readFileSync("configContract.json");
 var configData = JSON.parse(config);
// Get connection and function details from JSON connection file
 var contractId=configData.contractid;
 var hostIP=configData.host;
 var hostPort=configData.port;
 var hostPath=configData.path;
 var method =configData.method;
 var functionName =configData.functionName;
 var securecontext = configData.securecontext;

// If generate data flag is set to true, i.e., value is 'Y'
// generate new device data
// Get device data from devices file
var devices= fs.readFileSync("devices.json");
var jsonContent = JSON.parse(devices);
var items = jsonContent.devices;
var iCounter = 0;
 callRestAPI(contractId, hostIP, hostPort,hostPath, method,functionName, items, iCounter);

 
////////////////////////////callRestAPI///////////////////////////

function callRestAPI(contractId, hostIP, hostPort,hostPath, method, functionName, items, iCounter) {
    //if (iCounter ==items.length)
    if (iCounter ==1)
    {
        process.exit();
    }
    var deviceRaw = items[iCounter];
     deviceStr = JSON.stringify(deviceRaw).trim();
     var deviceData =  deviceStr.replace("\'", "\\\"");
     //console.log ("deviceQuote content is ", JSON.stringify(deviceData))
     var devStrEnd = deviceData.lastIndexOf("}");
     var currentTime = new Date();
     var device = deviceData.substr(0,devStrEnd) + ", \"timestamp\" : \""+ currentTime.getTime() + "\"}"
    // console.log(" Final device string is : \n", device )
    var jsonString1 = '{"jsonrpc": "2.0", "method": "'+method+'", "params": { "type": 1, "chaincodeID":{  "name":"';
    var jsonString2 = '"}, "ctorMsg": { "function":"'+functionName+'", "args":[';
    var jsonString3 = JSON.stringify(device);
    var jsonString4 = '] }, "secureContext": "user_type1_798389e89b" },  "id":1}';
    var jsonString = jsonString1+contractId+jsonString2+jsonString3+jsonString4
    console.log("jsonString is ", jsonString)
    // the post options
    var optionspost = {
        host : hostIP,
        port : hostPort,
        path : hostPath,
        method : 'POST',
        headers: {
            "Content-Type": "application/json",
            "Content-Length": Buffer.byteLength(jsonString)
        }
    };

    // do the POST call
        var reqPost = http.request(optionspost, function(res) {
            console.log("statusCode: ", res.statusCode);
            // uncomment it for header details
        //  console.log("headers: ", res.headers);
        
            res.on('data', function(d) {
                console.info('POST result:\n');
                process.stdout.write(d);
                console.info('\n\nPOST completed');
            });
        });
        
        // write the json data
        reqPost.write(jsonString);
        reqPost.end();
        reqPost.on('error', function(e) {
            console.error(e);
        });
    setTimeout(function() {
    iCounter += 1;
    console.log('Delay for 1000 ms');
    callRestAPI(contractId, hostIP, hostPort,hostPath, method,functionName, items, iCounter);
    }, 1000);
}

