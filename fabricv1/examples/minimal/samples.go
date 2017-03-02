package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/fabricv1/platform"
)

var samples = `

{
    "API": {
        "createAsset": {
            "args": [
                {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                }
            ],
            "function": "createAsset"
        }
    },
    "Model": {
        "asset": {
            "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
            "carrier": "The carrier in possession of this asset",
            "common": {
                "appdata": [
                    {
                        "K": "carpe noctem",
                        "V": "carpe noctem"
                    }
                ],
                "deviceID": "A unique identifier for the device that sent the current event",
                "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                "location": {
                    "latitude": 123.456,
                    "longitude": 123.456
                }
            },
            "temperature": 123.456
        },
        "assetstate": {
            "alerts": [
                "An alert name"
            ],
            "assetID": "This asset's world state asset ID",
            "class": {},
            "compliant": true,
            "eventin": {
                "asset": {
                    "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                    "carrier": "The carrier in possession of this asset",
                    "common": {
                        "appdata": [
                            {
                                "K": "carpe noctem",
                                "V": "carpe noctem"
                            }
                        ],
                        "deviceID": "A unique identifier for the device that sent the current event",
                        "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        }
                    },
                    "temperature": 123.456
                }
            },
            "eventout": {
                "asset": {
                    "name": "EVT.IOTCP.INVOKE.RESULT",
                    "payload": {
                        "properties": "NO TYPE PROPERTY"
                    }
                }
            },
            "state": {
                "asset": {
                    "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                    "carrier": "The carrier in possession of this asset",
                    "common": {
                        "appdata": [
                            {
                                "K": "carpe noctem",
                                "V": "carpe noctem"
                            }
                        ],
                        "deviceID": "A unique identifier for the device that sent the current event",
                        "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                        "location": {
                            "latitude": 123.456,
                            "longitude": 123.456
                        }
                    },
                    "temperature": 123.456
                }
            },
            "txnid": "Transaction UUID matching the blockchain",
            "txnts": "Transaction timestamp matching the blockchain"
        },
        "assetstatearray": [
            {
                "alerts": [
                    "An alert name"
                ],
                "assetID": "This asset's world state asset ID",
                "class": {},
                "compliant": true,
                "eventin": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "eventout": {
                    "asset": {
                        "name": "EVT.IOTCP.INVOKE.RESULT",
                        "payload": {
                            "properties": "NO TYPE PROPERTY"
                        }
                    }
                },
                "state": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "txnid": "Transaction UUID matching the blockchain",
                "txnts": "Transaction timestamp matching the blockchain"
            }
        ],
        "assetstateexternal": {
            "^DEF": {
                "alerts": [
                    "An alert name"
                ],
                "assetID": "This asset's world state asset ID",
                "class": {},
                "compliant": true,
                "eventin": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "eventout": {
                    "asset": {
                        "name": "EVT.IOTCP.INVOKE.RESULT",
                        "payload": {
                            "properties": "NO TYPE PROPERTY"
                        }
                    }
                },
                "state": {
                    "asset": {
                        "assetID": "An asset's unique ID, e.g. barcode, VIN, etc.",
                        "carrier": "The carrier in possession of this asset",
                        "common": {
                            "appdata": [
                                {
                                    "K": "carpe noctem",
                                    "V": "carpe noctem"
                                }
                            ],
                            "deviceID": "A unique identifier for the device that sent the current event",
                            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
                            "location": {
                                "latitude": 123.456,
                                "longitude": 123.456
                            }
                        },
                        "temperature": 123.456
                    }
                },
                "txnid": "Transaction UUID matching the blockchain",
                "txnts": "Transaction timestamp matching the blockchain"
            }
        },
        "eventIOTContractPlatformInvokeResult": {
            "name": "EVT.IOTCP.INVOKE.RESULT",
            "payload": {
                "properties": "NO TYPE PROPERTY"
            }
        },
        "eventIOTContractPlatformStatus": {
            "activeAlerts": [
                "An alert name"
            ],
            "alertsCleared": [
                "An alert name"
            ],
            "alertsRaised": [
                "An alert name"
            ],
            "invokeresult": {
                "message": "carpe noctem",
                "status": "ERROR"
            }
        },
        "ioteventcommon": {
            "appdata": [
                {
                    "K": "carpe noctem",
                    "V": "carpe noctem"
                }
            ],
            "deviceID": "A unique identifier for the device that sent the current event",
            "devicetimestamp": "A timestamp recoded by the device that sent the current event",
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            }
        },
        "stateFilter": {
            "match": "all",
            "select": [
                {
                    "qprop": "Qualified property to compare, for example 'asset.assetID'",
                    "value": "Value to be compared"
                }
            ]
        }
    }
}`


	var readAssetSamples iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(samples), nil
	}

	func init() {
		iot.AddRoute("readAssetSamples", "query", iot.SystemClass, readAssetSamples)
	}
	