# Event Listener Clients

## What is event-listener

The [event listener client application `event_listener.go`](./event-listener.go) will connect to a peer and receive block and custom chaincode events. A block event carries the entire block structure, and thus the event listener can print the result of all transactions in the block, either success or failure. Note that emitted events from the chaincode are also shown.

## To Run in Debug Mode

- go build

- ./event-listener -events-address=< event address >

For the 0.5 developer preview build of Hyperledger, the command line should look something like:

``` sh
vagrant@hyperledger-devenv:v0.0.10-cfc2099:/local-dev/github.com/blockchain-samples/applications/event_listener$ ./event_listener -events-address=0.0.0.0:31315
```

The running event listener will connect to the peer running in debug mode in another terminal window, and the chaincode running in a third terminal window.

## Sample Output When Running Against the PINGPONG Contract

Send a message with assetID "PING":

``` json
POST /chaincode HTTP/1.1
Host: tp-letkemank:5000
Accept: application/json
Content-Type: application/json
Cache-Control: no-cache
Postman-Token: 262df139-9ba2-1260-71c8-a3d04758f7b8

{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
        "type": 1,
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg": {
            "function":"updateAsset",
            "args":["{\"assetID\":\"PING\"}"]
        },
        "secureContext": ""
    },
    "id":1234
}
```

The event catcher prints:

``` txt
Received block
--------------
Success Transaction:
        [uuid:"c85d4006-7525-483e-afad-422f77645474" chaincodeEvent:<chaincodeID:"mycc" txID:"c85d4006-7525-483e-afad-422f77645474" eventName:"EVTPONG" payload:"PONG: 6" > ]

Received chaincode event
------------------------
&{ChaincodeEvent:chaincodeID:"mycc" txID:"c85d4006-7525-483e-afad-422f77645474" eventName:"EVTPONG" payload:"PONG: 6" }
```

> Note the count in the payload. This is the number of PONG events sent.

#Update for IoT Contract Platform

With the advent pof the IoT Contract Platform, there are chaincode events that somewhat duplicate what is available through the block cna rejection events. However, the block event, as the only positive  
indication of a successful transaction is somewhat flawed, in that it is necessary to parse all chaincode events in the block (1 per transaction in our case) and compare Againstthe chaincode id
and the event name in order to make sense of the information.

By contrast, the platform automatically sends success (OK) or failure (ERROR) status with a message if there has been an error. It also includes zero, one or both of alertsRaised and
alertsCleared alert name arrays. Empty arrays are left out of course. An application can declare an interest in this specific event and catch it immediately after each invoke has running
(and the block committed in the case of OK). 

Following is sample output from the event_ listener application that captures all of the relevant events.

``` text
vagrant@hyperledger-devenv:v0.0.11-b111ac5:/local-dev/src/github.com/ibm-watson-iot/blockchain-samples/applications/event_listener$ ./event_listener 
Event Address: 0.0.0.0:7053
Event client appears to have been succesfully created

Received rejection
&{Rejection:tx:<type:CHAINCODE_INVOKE chaincodeID:"\022\004mycc" payload:"\nW\010\001\022\006\022\004mycc\032=\n\013createAsset\n.{\"asset\":{\"assetID\":\"12345\",\"temperature\":-4}}*\014user_type1_0" txid:"558afcd2-e7d0-49d4-8d4e-f8dbfdc393fd" timestamp:<seconds:1480228360 nanos:580007792 > > errorMsg:"Transaction or query returned with failure: Invoke (createAsset) failed with error CreateAsset for class default asset DEF12345 asset already exists" }


Received block
&{Block:stateHash:"\001P\363;\257\220\216\254\321\304\323\261\362f\277{)\177\252\256\2412\267F\311\325\000\352~\347B\204\331/\276\263t\340\300\3041{\276F\034G\205\244\250\310\t\336IW\010\003S\210\017\276\257\205O\035" previousBlockHash:"L\303\232\r\346\316\\\3145P\243\365k\007A^\2600Y\343\307X\244e\377\357\017o\372W\247\250.\340\250u\240\233\263\362j\333\000\3229\332n\266\215.\310\360\305\r\374}\010\210\313p\272\254]I" nonHashData:<localLedgerCommitTimestamp:<seconds:1480228361 nanos:585408919 > chaincodeEvents:<chaincodeID:"mycc" txID:"558afcd2-e7d0-49d4-8d4e-f8dbfdc393fd" eventName:"EVT.IOTCP.INVOKE.RESULT" payload:"{\"message\":\"Invoke (createAsset) failed with error CreateAsset for class default asset DEF12345 asset already exists\",\"status\":\"ERROR\"}" > > }


Received chaincode event
&{ChaincodeEvent:chaincodeID:"mycc" txID:"558afcd2-e7d0-49d4-8d4e-f8dbfdc393fd" eventName:"EVT.IOTCP.INVOKE.RESULT" payload:"{\"message\":\"Invoke (createAsset) failed with error CreateAsset for class default asset DEF12345 asset already exists\",\"status\":\"ERROR\"}" }


Received block
&{Block:transactions:<type:CHAINCODE_INVOKE chaincodeID:"\022\004mycc" payload:"\nW\010\001\022\006\022\004mycc\032=\n\013updateAsset\n.{\"asset\":{\"assetID\":\"12345\",\"temperature\":-4}}*\014user_type1_0" txid:"406325dc-ec30-4a85-8ed1-7fc3317fa243" timestamp:<seconds:1480228474 nanos:971235489 > > stateHash:",U\252|\334\017W\204\"\t\232\206$\007)\210\240\247\200\007\030=T\341\267\374XFX\020\014B\216\251%\031i\201\027\210\240Y\262\200\243\363)Y\031\305\227k\246Y\321\341\025\322mK\266\265\026V" previousBlockHash:"/\271\315\275\366>*e?\336aE\002P$\007\371?\335\177D6\251D\261\261\314\222d\020\013\1772b\310\231\026G\370\346t9\r\002\3238\207\007\300\327c7\362H\010y\024\230\\N\200\312\3644" nonHashData:<localLedgerCommitTimestamp:<seconds:1480228475 nanos:981112352 > chaincodeEvents:<chaincodeID:"mycc" txID:"406325dc-ec30-4a85-8ed1-7fc3317fa243" eventName:"EVT.IOTCP.INVOKE.RESULT" payload:"{\"alertsCleared\":[\"OVERTEMP\"],\"status\":\"OK\"}" > > }


Received chaincode event
&{ChaincodeEvent:chaincodeID:"mycc" txID:"406325dc-ec30-4a85-8ed1-7fc3317fa243" eventName:"EVT.IOTCP.INVOKE.RESULT" payload:"{\"alertsCleared\":[\"OVERTEMP\"],\"status\":\"OK\"}" }

``` 