package main

var samples = `
{
    "activity": {
        "activityDate": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
        "activityDetail": "Detailed description of the activity.",
        "activityNumber": "Activity number.",
        "activityType": "The type of the activity."
    },
    "airline": {
        "activities": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "airplanes": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "code": "The airline 3 letter code.",
        "name": "The name of the airline."
    },
    "airplane": {
        "activities": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "assemblies": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "model": "The airplane model or family name.",
        "variant": "A manufacturer-specific variant identifier for this airplane."
    },
    "assembly": {
        "ALRS": "Alerting service zone.",
        "ATAcode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
        "activities": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "lifeLimitInitial": 789,
        "lifeLimitUsed": 789,
        "name": "The assembly name.",
        "parts": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ]
    },
    "event": {
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "extension": [
            {}
        ],
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "oneOf": [
            {
                "airline": {
                    "activities": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "airplanes": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "code": "The airline 3 letter code.",
                    "name": "The name of the airline."
                }
            },
            {
                "airplane": {
                    "activities": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "assemblies": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "model": "The airplane model or family name.",
                    "variant": "A manufacturer-specific variant identifier for this airplane."
                }
            },
            {
                "assembly": {
                    "ALRS": "Alerting service zone.",
                    "ATAcode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                    "activities": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "lifeLimitInitial": 789,
                    "lifeLimitUsed": 789,
                    "name": "The assembly name.",
                    "parts": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ]
                }
            },
            {
                "part": {
                    "activities": [
                        "The ID of a managed asset. The resource focal point for a smart contract."
                    ],
                    "lifeLimitInitial": 789,
                    "lifeLimitUsed": 789,
                    "name": "The part name.",
                    "partNumber": "Part number.",
                    "vendorName": "Vendor name.",
                    "vendorNumber": "Vendor part number."
                }
            },
            {
                "activity": {
                    "activityDate": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                    "activityDetail": "Detailed description of the activity.",
                    "activityNumber": "Activity number.",
                    "activityType": "The type of the activity."
                }
            }
        ],
        "references": [
            "carpe noctem"
        ],
        "timestamp": "2016-06-12T15:19:48.45276341Z"
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "part": {
        "activities": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "lifeLimitInitial": 789,
        "lifeLimitUsed": 789,
        "name": "The part name.",
        "partNumber": "Part number.",
        "vendorName": "Vendor name.",
        "vendorNumber": "Vendor part number."
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
        "compliant": true,
        "event": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "oneOf": [
                {
                    "airline": {
                        "activities": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "airplanes": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "code": "The airline 3 letter code.",
                        "name": "The name of the airline."
                    }
                },
                {
                    "airplane": {
                        "activities": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "assemblies": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "model": "The airplane model or family name.",
                        "variant": "A manufacturer-specific variant identifier for this airplane."
                    }
                },
                {
                    "assembly": {
                        "ALRS": "Alerting service zone.",
                        "ATAcode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                        "activities": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "lifeLimitInitial": 789,
                        "lifeLimitUsed": 789,
                        "name": "The assembly name.",
                        "parts": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ]
                    }
                },
                {
                    "part": {
                        "activities": [
                            "The ID of a managed asset. The resource focal point for a smart contract."
                        ],
                        "lifeLimitInitial": 789,
                        "lifeLimitUsed": 789,
                        "name": "The part name.",
                        "partNumber": "Part number.",
                        "vendorName": "Vendor name.",
                        "vendorNumber": "Vendor part number."
                    }
                },
                {
                    "activity": {
                        "activityDate": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                        "activityDetail": "Detailed description of the activity.",
                        "activityNumber": "Activity number.",
                        "activityType": "The type of the activity."
                    }
                }
            ],
            "references": [
                "carpe noctem"
            ],
            "timestamp": "2016-06-12T15:19:48.452829951Z"
        },
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`