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
        "timestamp": "2016-04-27T13:03:48.9253142-04:00"
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
        "extension": {},
        "inCompliance": true,
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
        "timestamp": "2016-04-27T13:03:48.9263143-04:00"
    }
}`