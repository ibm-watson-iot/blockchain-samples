# Testing the Trade Lane Contract

This JavaScript test script has been used to explore various functions of the Trade Lane contract. It tests all the REST APIs of the contract, which runs alongside an IBM Blockchain peer.

These are the three main RESTful end points supported by an IBM Blockchain peer. 
  1.	/devops/deploy
  2.	/devops/invoke
  3.	/devops/query

There is also a /registrar endpoint for work in secure environments.
Making a REST call to a peer consists of calling one of the above endpoints (for example, /devops/invoke) and passing in a JSON formatted message. 

The following operations have been tested.
1.	Deploying contract to the peer
2.	Create Asset
3.	Read Asset
4.	Update Asset
5.	Delete Asset
6.	Read All Assets
7.	Delete All Assets
8.	Alerts when temperature is above/below the threshold
9.	Read Contract states
10.	Read Asset Schemas
11.	Read Asset Samples
12.	Read Contract Object Model
13.	Read Recent states
14.	Read Asset History
15.	Set Create on Update
16.	Set Logging Level

## Requirements

1. IBM Blockchain peer  
trade_lane_contract_e2e_tests.js was tested with peers running in IBM Bluemix.

2. Node.js, Mocha and additional supporting libraries installed locally on your machine. (Installation steps 1 & 4 below)


## Installation

1. Install Node.js on your local machine.

2. Clone the blockchain-samples project:  
https://github.com/ibm-watson-iot/blockchain-samples   or download the zip of the project. The zip can be downloaded by clicking the Download ZIP button found on the right at the above link.

3.  Navigate to the contract tests directory: trade_lane_contract_e2e_tests/  

4. Run the following command:  
npm install  
  This generates a folder called node_modules in your contract tests directory. In that folder, you should see the following libraries:  
mocha  
log4js@~0.6.21  
chai@~1.9.1  
chai-http@~0.5.0  
request@~2.44.0  
request-promise@~2.0.1  
promise@~6.0.0  
q@~1.0.1  

5. In the contract tests directory, edit config.js.  
HTTP and HTTPS are well known protocols that run on well known ports: HTTP is synonymous with 80 and HTTPS is equally synonymous with 443.  
Replace the following fields with your own credentials which you get from IBM Blockchain:  
    <br>obcHost: "api_host",  
    obcPort: api_port,  
    secure_context: "username",  
    enroll_secret: "secret"

## Execution

In the contract tests directory, run:  
mocha trade_lane_contract_e2e_tests.js

Note: If you see 400 errors after deploying the chaincode, this is likely a timing issue.  You can a) wait and try again or b) extend the timeout value at line 53 in trade_lane_contract_e2e_tests.js.

## Further Reading

trade_lane_contract_e2e_tests.js is written using Mocha and promises. Mocha is the Node.js testing framework. The following are websites relevant to Mocha & promises:

  * [http://www.mochajs.org](http://www.mochajs.org)
  * [https://github.com/mochajs/mocha](https://github.com/mochajs/mocha)
  * [https://github.com/kriskowal/q](https://github.com/kriskowal/q)
