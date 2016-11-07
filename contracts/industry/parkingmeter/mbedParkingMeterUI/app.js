/*eslint-env node*/

//------------------------------------------------------------------------------
// node.js starter application for Bluemix
//------------------------------------------------------------------------------

// This application uses express as its web server
// for more info, see: http://expressjs.com
var express = require('express');

// cfenv provides access to your Cloud Foundry environment
// for more info, see: https://www.npmjs.com/package/cfenv
var cfenv = require('cfenv');

// create a new express server
var app = express();

// serve the files out of ./public as our main files
app.use(express.static(__dirname + '/public'));

// get the app environment from Cloud Foundry
var appEnv = cfenv.getAppEnv();

var fs = require('fs');
// https call to IBM payment service
var https = require('https');
// Parse xml response 
//var parseString = require('xml2js').parseString;


///////////////////////////////////////////
//IBM Payment integration : Get session link
// start the session
app.post('/session', function(req, res) {
	 invokePaymentInstance(function(err, result){
    if(err){
      res.send(500, { error: 'something blew up' });
    } else {
      res.send(result);
    }
  });
});
var invokePaymentInstance = function(callback){

	var xmlString ='<TRX> <SVC>ProfileStartSession</SVC> <PRJ>WATSONIOT</PRJ> <CTY>US</CTY> <PRF>MyWallet-12345</PRF> ' +
	' <COM>Internet</COM> <CUR>USD</CUR> <SES>  <SUU>http://www.example.com</SUU> <FAU>http://www.example.com </FAU> ' +
	'<CAU>http://www.example.com</CAU> <LAN>en</LAN> <MBV>true</MBV> <ACT>AutoPay</ACT> <VOD>V22</VOD> </SES> </TRX>';
  	console.log(xmlString);
	var fileContents;
	try {
		  fileKeyContents = fs.readFileSync('./watsoniot.key.pem');
		  fileCertContents = fs.readFileSync('./watsoniot.cert.pem');
	  } catch (err) {
	  	console.log("Error encountered"+ err);
	}
	var options = {
	  hostname: 'ips-preprod.ihost.com',
	  port: 50443,
	  path: '/trx',
	  key: fileKeyContents,
	  cert: fileCertContents,
	  method: 'POST',
	  connection: 'close',
	  headers: {
	    'Content-Length': Buffer.byteLength(xmlString)
	  }
	};  
	var request = https.request(options, function(response) {
	  console.log('Inside request');
	  console.log('Status: ' + response.statusCode);
	  console.log('Headers: ' + JSON.stringify(response.headers));
	  response.on('data', function (body) {
	    console.log('Body: ' + body);
	    var sessionData = body.toString('utf8');
	    var startPosn = sessionData.lastIndexOf ("<SEU>");
	    var endPosn = sessionData.lastIndexOf ("</SEU>");
	    console.log('start posn: '+startPosn);
	    console.log('end posn: '+endPosn);
	    sessionPath= sessionData.substr(startPosn+5, (endPosn-startPosn-5));
	    console.log('returning sessionPath: ' + sessionPath);
	    callback(null, {sessionPath : sessionPath});
	    //parseString(body, function (err, result) {
	    //	console.log("Session path is : " +JSON.stringify(result.TRX.SES.SEU));
     	//console.dir(JSON.stringify(result));
     	//res.send(data); 
	  });
	 //);
	});
	request.write(xmlString);
	request.end();
	request.on('error', function(e) {
	  console.log('returning : problem with request: ' + e.message);
	  callback(e);
	});

};
/////////////////////////////////////////////
//IBM Payment call AuthCap
 
 app.post('/pymt', function(req, res) {
	 invokeAndProcessPayment(function(err, result){
    if(err){
      res.send(500, { error: 'something blew up' });
    } else {
      res.send(result);
    }
  });
});
var invokeAndProcessPayment = function(callback){

	var xmlString ='<TRX> <SVC>AuthCap</SVC> <PRJ>WATSONIOT</PRJ> <CTY>US</CTY> <ORD>12345</ORD> <COM>Internet</COM> '+
  ' <CUR>USD</CUR> <NET>10.00</NET> <TAX>1.00</TAX>  <GRS>11.00</GRS> <PRF>MyWallet-12345</PRF> </TRX> ';
  	console.log(xmlString);
	var fileContents;
	try {
		  fileKeyContents = fs.readFileSync('./watsoniot.key.pem');
		  fileCertContents = fs.readFileSync('./watsoniot.cert.pem');
	  } catch (err) {
	  	console.log("Error encountered"+ err);
	}
	var options = {
	  hostname: 'ips-preprod.ihost.com',
	  port: 50443,
	  path: '/trx',
	  key: fileKeyContents,
	  cert: fileCertContents,
	  method: 'POST',
	  connection: 'close',
	  headers: {
	    'Content-Length': Buffer.byteLength(xmlString)
	  }
	};  
	var request = https.request(options, function(response) {
	  console.log('Inside request');
	  console.log('Status: ' + response.statusCode);
	  console.log('Headers: ' + JSON.stringify(response.headers));
	  response.on('data', function (body) {
	    console.log('Body: ' + body);
	    var responseMessage = body.toString('utf8');
	    var startPosn = responseMessage.lastIndexOf ("<MSG>");
	    var endPosn = responseMessage.lastIndexOf ("</MSG>");
	    console.log('start posn: '+startPosn);
	    console.log('end posn: '+endPosn);
	    successMessage= responseMessage.substr(startPosn+5, (endPosn-startPosn-5));
	    console.log('returning sessionPath: ' + successMessage);
	    callback(null, {successMessage : successMessage});
	    //parseString(body, function (err, result) {
	    //	console.log("Session path is : " +JSON.stringify(result.TRX.SES.SEU));
     	//console.dir(JSON.stringify(result));
     	//res.send(data); 
	  });
	 //);
	});
	request.write(xmlString);
	request.end();
	request.on('error', function(e) {
	  console.log('returning : problem with request: ' + e.message);
	  callback(e);
	});

};
/////////////////////////////////////////////
// start server on the specified port and binding host
app.listen(appEnv.port, '0.0.0.0', function() {
  // print a message when the server starts listening
  console.log("server starting on " + appEnv.url);
});
