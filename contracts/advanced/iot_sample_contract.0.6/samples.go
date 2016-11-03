package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "TRADELANE",
        "version": "The version number of the current contract"
    },
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "extension": {},
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "temperature": 123.456,
        "timestamp": "2016-11-03T06:01:31.011326974Z"
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "alerts": {
            "active": [
                "OVERTTEMP"
            ],
            "cleared": [
                "OVERTTEMP"
            ],
            "raised": [
                "OVERTTEMP"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "carrier": "transport entity currently in possession of asset",
        "compliant": true,
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
        "temperature": 123.456,
        "timestamp": "2016-11-03T06:01:31.011419799Z",
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`