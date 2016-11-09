package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
	)

var samples = `

{
    "container": {
        "barcode": "The ID of a container.",
        "carrier": "transport entity currently in possession of the container",
        "common": {
            "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "references": [
                "carpe noctem"
            ]
        },
        "temperature": 123.456
    },
    "containerstate": {
        "AssetKey": "The World State asset ID. Used to read and write state.",
        "alerts": [
            "OVERTTEMP"
        ],
        "assetIDpath": "Qualified property path to the asset's ID.",
        "class": "Asset class.",
        "compliant": true,
        "eventin": {
            "container": {
                "barcode": "The ID of a container.",
                "carrier": "transport entity currently in possession of the container",
                "common": {
                    "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                    "extension": [
                        {}
                    ],
                    "location": {
                        "latitude": 123.456,
                        "longitude": 123.456
                    },
                    "references": [
                        "carpe noctem"
                    ]
                },
                "temperature": 123.456
            }
        },
        "eventout": {
            "container": {
                "name": "Event name.",
                "payload": {}
            }
        },
        "prefix": "Asset class prefix in World State.",
        "state": {
            "container": {
                "barcode": "The ID of a container.",
                "carrier": "transport entity currently in possession of the container",
                "common": {
                    "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                    "extension": [
                        {}
                    ],
                    "location": {
                        "latitude": 123.456,
                        "longitude": 123.456
                    },
                    "references": [
                        "carpe noctem"
                    ]
                },
                "temperature": 123.456
            }
        },
        "txnid": "Transaction UUID matching that in the blockchain.",
        "txnts": "Transaction timestamp matching that in the blockchain."
    },
    "containerstatearray": [
        {
            "^CON": {
                "AssetKey": "The World State asset ID. Used to read and write state.",
                "alerts": [
                    "OVERTTEMP"
                ],
                "assetIDpath": "Qualified property path to the asset's ID.",
                "class": "Asset class.",
                "compliant": true,
                "eventin": {
                    "container": {
                        "barcode": "The ID of a container.",
                        "carrier": "transport entity currently in possession of the container",
                        "common": {
                            "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                            "extension": [
                                {}
                            ],
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            },
                            "references": [
                                "carpe noctem"
                            ]
                        },
                        "temperature": 123.456
                    }
                },
                "eventout": {
                    "container": {
                        "name": "Event name.",
                        "payload": {}
                    }
                },
                "prefix": "Asset class prefix in World State.",
                "state": {
                    "container": {
                        "barcode": "The ID of a container.",
                        "carrier": "transport entity currently in possession of the container",
                        "common": {
                            "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                            "extension": [
                                {}
                            ],
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            },
                            "references": [
                                "carpe noctem"
                            ]
                        },
                        "temperature": 123.456
                    }
                },
                "txnid": "Transaction UUID matching that in the blockchain.",
                "txnts": "Transaction timestamp matching that in the blockchain."
            }
        }
    ],
    "containerstateexternal": {
        "^CON": {
            "AssetKey": "The World State asset ID. Used to read and write state.",
            "alerts": [
                "OVERTTEMP"
            ],
            "assetIDpath": "Qualified property path to the asset's ID.",
            "class": "Asset class.",
            "compliant": true,
            "eventin": {
                "container": {
                    "barcode": "The ID of a container.",
                    "carrier": "transport entity currently in possession of the container",
                    "common": {
                        "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                        "extension": [
                            {}
                        ],
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "references": [
                            "carpe noctem"
                        ]
                    },
                    "temperature": 123.456
                }
            },
            "eventout": {
                "container": {
                    "name": "Event name.",
                    "payload": {}
                }
            },
            "prefix": "Asset class prefix in World State.",
            "state": {
                "container": {
                    "barcode": "The ID of a container.",
                    "carrier": "transport entity currently in possession of the container",
                    "common": {
                        "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                        "extension": [
                            {}
                        ],
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        },
                        "references": [
                            "carpe noctem"
                        ]
                    },
                    "temperature": 123.456
                }
            },
            "txnid": "Transaction UUID matching that in the blockchain.",
            "txnts": "Transaction timestamp matching that in the blockchain."
        }
    },
    "invokeevent": {
        "name": "Event name.",
        "payload": {}
    },
    "ioteventcommon": {
        "devicetimestamp": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
        "extension": [
            {}
        ],
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "references": [
            "carpe noctem"
        ]
    }
}`


	var readAssetSamples as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(samples), nil
	}

	func init() {
		as.AddRoute("readAssetSamples", "query", as.SystemClass, readAssetSamples)
	}
	