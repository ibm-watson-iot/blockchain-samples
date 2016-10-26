package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "Ping Pong",
        "version": "The version number of the current contract"
    },
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "extension": {},
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "timestamp": "2016-10-25T15:09:03.78690912Z"
    },
    "initEvent": {
        "nickname": "Ping Pong",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "alerts": {
            "active": [
                "N/A"
            ],
            "cleared": [
                "N/A"
            ],
            "raised": [
                "N/A"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "compliant": true,
        "errorcount": 789,
        "extension": {},
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "pingcount": 789,
        "timestamp": "2016-10-25T15:09:03.787045505Z",
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`