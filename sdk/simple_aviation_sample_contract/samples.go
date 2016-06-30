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
        "airplane": {
            "acmodel": "Aircraft model",
            "acnumber": "Aircraft number",
            "airline": "Airline name",
            "lifelimitinitial": 123.456
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "flight": {
            "flightnumber": "A flight number",
            "hardlanding": true
        },
        "inspection": "ACHECK",
        "iotcommon": {
            "extension": {},
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "timestamp": "2016-06-28T15:44:07.447759544Z"
        },
        "lifelimitdeduct": 123.456
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "airplane": {
            "acmodel": "Aircraft model",
            "acnumber": "Aircraft number",
            "airline": "Airline name",
            "cyclecounter": 123.456,
            "hardlanding": true,
            "initiallifelimit": 123.456,
            "lifelimitused": 123.456
        },
        "alerts": {
            "active": [
                "NONE",
                "ACHECK",
                "BCHECK"
            ],
            "cleared": [
                "NONE",
                "ACHECK",
                "BCHECK"
            ],
            "raised": [
                "NONE",
                "ACHECK",
                "BCHECK"
            ]
        },
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "compliant": true,
        "flight": {
            "flightnumber": "A flight number",
            "hardlanding": true
        },
        "inspection": "ACHECK",
        "iotcommon": {
            "extension": {},
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "timestamp": "2016-06-28T15:44:07.447871569Z"
        },
        "lastEvent": {
            "arg": {
                "airplane": {
                    "acmodel": "Aircraft model",
                    "acnumber": "Aircraft number",
                    "airline": "Airline name",
                    "lifelimitinitial": 123.456
                },
                "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
                "flight": {
                    "flightnumber": "A flight number",
                    "hardlanding": true
                },
                "inspection": "ACHECK",
                "iotcommon": {
                    "extension": {},
                    "location": {
                        "latitude": 123.456,
                        "longitude": 123.456
                    },
                    "timestamp": "2016-06-28T15:44:07.447855645Z"
                },
                "lifelimitdeduct": 123.456
            },
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "lifelimitdeduct": 123.456,
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`