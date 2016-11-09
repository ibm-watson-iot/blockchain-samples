package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
	)

var schemas = `

{
    "API": {
        "createAssetContainer": {
            "description": "Create an asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "barcode": {
                                "description": "The ID of a container.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of the container",
                                "type": "string"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                "properties": {
                                    "devicetimestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
                                    },
                                    "extension": {
                                        "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                        "items": {
                                            "properties": {},
                                            "type": "object"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "location": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "references": {
                                        "description": "An array of external references relevant to this asset.",
                                        "items": {
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "temperature": {
                                "description": "Temperature of the inside of the container in CELSIUS.",
                                "type": "number"
                            }
                        },
                        "required": [
                            "barcode"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAsset function",
                    "enum": [
                        "createAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsContainer": {
            "description": "Delete the state of all assets. No arguments are accepted. For each managed asset, the state and history are erased, and the asset is removed if necessary from recent states.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAllAssets function",
                    "enum": [
                        "deleteAllAssets"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetContainer": {
            "description": "Delete an asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The ID of a container.",
                        "type": "string"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetContainer": {
            "description": "Delete one or more properties from an asset's state. Argument is a JSON encoded string containing an 'assetID' and an array of qualified property names. For example, in an event object containing common and custom properties objects, the argument might look like {'assetID':'A1',['common.location', 'custom.carrier', 'custom.temperature']} and the result of that invoke would be the removal of the location, carrier and temperature properties. The missing temperature would clear a 'OVERTEMP' alert when the rules engine runs.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "barcode": {
                            "description": "The ID of a container.",
                            "type": "string"
                        },
                        "qprops": {
                            "items": {
                                "description": "The qualified name of a property. E.g. 'event.common.carrier', 'event.custom.temperature', etc.",
                                "type": "string"
                            },
                            "type": "array"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deletePropertiesFromAsset function",
                    "enum": [
                        "deletePropertiesFromAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "nickname": {
                            "default": "CTIORSAMPLE",
                            "description": "The nickname of the current contract",
                            "type": "string"
                        },
                        "version": {
                            "description": "The version number of the current contract",
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "init function",
                    "enum": [
                        "init"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssetsContainer": {
            "description": "Returns the state of all assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by 'assetID'.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAllAssets function",
                    "enum": [
                        "readAllAssets"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of container states, often for different assets",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one container asset.",
                                "properties": {
                                    "AssetKey": {
                                        "description": "The World State asset ID. Used to read and write state.",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "List of alert names that are currently active.",
                                        "items": {
                                            "description": "Container alerts. Triggered or cleared by contract rules.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the asset's ID.",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "Asset class.",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The event that created this state.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "A chaincode event to be emitted on invoke exit.",
                                        "properties": {
                                            "container": {
                                                "description": "An event that is emitted at the end of an invoke, if it is present.",
                                                "properties": {
                                                    "name": {
                                                        "description": "Event name.",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A json object containing the event's properties.",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "Asset class prefix in World State.",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "The state of the one asset.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetContainer": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "barcode": {
                            "description": "The ID of a container.",
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "The external state of one container asset.",
                    "properties": {
                        "AssetKey": {
                            "description": "The World State asset ID. Used to read and write state.",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "List of alert names that are currently active.",
                            "items": {
                                "description": "Container alerts. Triggered or cleared by contract rules.",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the asset's ID.",
                            "type": "string"
                        },
                        "class": {
                            "description": "Asset class.",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "A contract-specific indication that this asset is compliant.",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The event that created this state.",
                            "properties": {
                                "container": {
                                    "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                    "properties": {
                                        "barcode": {
                                            "description": "The ID of a container.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "transport entity currently in possession of the container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                            "properties": {
                                                "devicetimestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                },
                                                "extension": {
                                                    "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                    "items": {
                                                        "properties": {},
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "references": {
                                                    "description": "An array of external references relevant to this asset.",
                                                    "items": {
                                                        "type": "string"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of the inside of the container in CELSIUS.",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "A chaincode event to be emitted on invoke exit.",
                            "properties": {
                                "container": {
                                    "description": "An event that is emitted at the end of an invoke, if it is present.",
                                    "properties": {
                                        "name": {
                                            "description": "Event name.",
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A json object containing the event's properties.",
                                            "properties": {},
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "Asset class prefix in World State.",
                            "type": "string"
                        },
                        "state": {
                            "description": "The state of the one asset.",
                            "properties": {
                                "container": {
                                    "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                    "properties": {
                                        "barcode": {
                                            "description": "The ID of a container.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "transport entity currently in possession of the container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                            "properties": {
                                                "devicetimestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                },
                                                "extension": {
                                                    "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                    "items": {
                                                        "properties": {},
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "references": {
                                                    "description": "An array of external references relevant to this asset.",
                                                    "items": {
                                                        "type": "string"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of the inside of the container in CELSIUS.",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching that in the blockchain.",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching that in the blockchain.",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistoryContainer": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "barcode": {
                            "description": "The ID of a container.",
                            "type": "string"
                        },
                        "end": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        },
                        "start": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetHistory function",
                    "enum": [
                        "readAssetHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of container states, often for different assets",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one container asset.",
                                "properties": {
                                    "AssetKey": {
                                        "description": "The World State asset ID. Used to read and write state.",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "List of alert names that are currently active.",
                                        "items": {
                                            "description": "Container alerts. Triggered or cleared by contract rules.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the asset's ID.",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "Asset class.",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The event that created this state.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "A chaincode event to be emitted on invoke exit.",
                                        "properties": {
                                            "container": {
                                                "description": "An event that is emitted at the end of an invoke, if it is present.",
                                                "properties": {
                                                    "name": {
                                                        "description": "Event name.",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A json object containing the event's properties.",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "Asset class prefix in World State.",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "The state of the one asset.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readRecentStates": {
            "description": "Returns the state of recently updated assets as an array of objects sorted with the most recently updated asset first. Each asset appears exactly once up to a maxmum of 20 in this version of the contract.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readRecentStates function",
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of container states, often for different assets",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one container asset.",
                                "properties": {
                                    "AssetKey": {
                                        "description": "The World State asset ID. Used to read and write state.",
                                        "type": "string"
                                    },
                                    "alerts": {
                                        "description": "List of alert names that are currently active.",
                                        "items": {
                                            "description": "Container alerts. Triggered or cleared by contract rules.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assetIDpath": {
                                        "description": "Qualified property path to the asset's ID.",
                                        "type": "string"
                                    },
                                    "class": {
                                        "description": "Asset class.",
                                        "type": "string"
                                    },
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "eventin": {
                                        "description": "The event that created this state.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "eventout": {
                                        "description": "A chaincode event to be emitted on invoke exit.",
                                        "properties": {
                                            "container": {
                                                "description": "An event that is emitted at the end of an invoke, if it is present.",
                                                "properties": {
                                                    "name": {
                                                        "description": "Event name.",
                                                        "type": "string"
                                                    },
                                                    "payload": {
                                                        "description": "A json object containing the event's properties.",
                                                        "properties": {},
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "prefix": {
                                        "description": "Asset class prefix in World State.",
                                        "type": "string"
                                    },
                                    "state": {
                                        "description": "The state of the one asset.",
                                        "properties": {
                                            "container": {
                                                "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "barcode": {
                                                        "description": "The ID of a container.",
                                                        "type": "string"
                                                    },
                                                    "carrier": {
                                                        "description": "transport entity currently in possession of the container",
                                                        "type": "string"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                        "properties": {
                                                            "devicetimestamp": {
                                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                                "type": "string"
                                                            },
                                                            "extension": {
                                                                "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                                "items": {
                                                                    "properties": {},
                                                                    "type": "object"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            },
                                                            "location": {
                                                                "description": "A geographical coordinate",
                                                                "properties": {
                                                                    "latitude": {
                                                                        "type": "number"
                                                                    },
                                                                    "longitude": {
                                                                        "type": "number"
                                                                    }
                                                                },
                                                                "type": "object"
                                                            },
                                                            "references": {
                                                                "description": "An array of external references relevant to this asset.",
                                                                "items": {
                                                                    "type": "string"
                                                                },
                                                                "minItems": 0,
                                                                "type": "array"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "temperature": {
                                                        "description": "Temperature of the inside of the container in CELSIUS.",
                                                        "type": "number"
                                                    }
                                                },
                                                "required": [
                                                    "barcode"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "txnid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnts": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "setCreateOnUpdate": {
            "description": "Allow updateAsset to redirect to createAsset when 'assetID' does not exist.",
            "properties": {
                "args": {
                    "description": "True for redirect allowed, false for error on asset does not exist.",
                    "items": {
                        "setCreateOnUpdate": {
                            "type": "boolean"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setCreateOnUpdate function",
                    "enum": [
                        "setCreateOnUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level in the contract.",
            "properties": {
                "args": {
                    "description": "logging levels indicate what you see",
                    "items": {
                        "logLevel": {
                            "enum": [
                                "CRITICAL",
                                "ERROR",
                                "WARNING",
                                "NOTICE",
                                "INFO",
                                "DEBUG"
                            ],
                            "type": "string"
                        }
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "setLoggingLevel function",
                    "enum": [
                        "setLoggingLevel"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetContainer": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "barcode": {
                                "description": "The ID of a container.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of the container",
                                "type": "string"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                "properties": {
                                    "devicetimestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
                                    },
                                    "extension": {
                                        "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                        "items": {
                                            "properties": {},
                                            "type": "object"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "location": {
                                        "description": "A geographical coordinate",
                                        "properties": {
                                            "latitude": {
                                                "type": "number"
                                            },
                                            "longitude": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "references": {
                                        "description": "An array of external references relevant to this asset.",
                                        "items": {
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "temperature": {
                                "description": "Temperature of the inside of the container in CELSIUS.",
                                "type": "number"
                            }
                        },
                        "required": [
                            "barcode"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAsset function",
                    "enum": [
                        "updateAsset"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "container": {
            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
            "properties": {
                "barcode": {
                    "description": "The ID of a container.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of the container",
                    "type": "string"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                    "properties": {
                        "devicetimestamp": {
                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                            "type": "string"
                        },
                        "extension": {
                            "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                            "items": {
                                "properties": {},
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "location": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "latitude": {
                                    "type": "number"
                                },
                                "longitude": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
                        "references": {
                            "description": "An array of external references relevant to this asset.",
                            "items": {
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of the inside of the container in CELSIUS.",
                    "type": "number"
                }
            },
            "required": [
                "barcode"
            ],
            "type": "object"
        },
        "containerstate": {
            "description": "The external state of one container asset.",
            "properties": {
                "AssetKey": {
                    "description": "The World State asset ID. Used to read and write state.",
                    "type": "string"
                },
                "alerts": {
                    "description": "List of alert names that are currently active.",
                    "items": {
                        "description": "Container alerts. Triggered or cleared by contract rules.",
                        "enum": [
                            "OVERTTEMP"
                        ],
                        "type": "string"
                    },
                    "type": "array"
                },
                "assetIDpath": {
                    "description": "Qualified property path to the asset's ID.",
                    "type": "string"
                },
                "class": {
                    "description": "Asset class.",
                    "type": "string"
                },
                "compliant": {
                    "description": "A contract-specific indication that this asset is compliant.",
                    "type": "boolean"
                },
                "eventin": {
                    "description": "The event that created this state.",
                    "properties": {
                        "container": {
                            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                            "properties": {
                                "barcode": {
                                    "description": "The ID of a container.",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "transport entity currently in possession of the container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                    "properties": {
                                        "devicetimestamp": {
                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                            "type": "string"
                                        },
                                        "extension": {
                                            "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                            "items": {
                                                "properties": {},
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "location": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "references": {
                                            "description": "An array of external references relevant to this asset.",
                                            "items": {
                                                "type": "string"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of the inside of the container in CELSIUS.",
                                    "type": "number"
                                }
                            },
                            "required": [
                                "barcode"
                            ],
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "eventout": {
                    "description": "A chaincode event to be emitted on invoke exit.",
                    "properties": {
                        "container": {
                            "description": "An event that is emitted at the end of an invoke, if it is present.",
                            "properties": {
                                "name": {
                                    "description": "Event name.",
                                    "type": "string"
                                },
                                "payload": {
                                    "description": "A json object containing the event's properties.",
                                    "properties": {},
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "prefix": {
                    "description": "Asset class prefix in World State.",
                    "type": "string"
                },
                "state": {
                    "description": "The state of the one asset.",
                    "properties": {
                        "container": {
                            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                            "properties": {
                                "barcode": {
                                    "description": "The ID of a container.",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "transport entity currently in possession of the container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                    "properties": {
                                        "devicetimestamp": {
                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                            "type": "string"
                                        },
                                        "extension": {
                                            "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                            "items": {
                                                "properties": {},
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "location": {
                                            "description": "A geographical coordinate",
                                            "properties": {
                                                "latitude": {
                                                    "type": "number"
                                                },
                                                "longitude": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "references": {
                                            "description": "An array of external references relevant to this asset.",
                                            "items": {
                                                "type": "string"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of the inside of the container in CELSIUS.",
                                    "type": "number"
                                }
                            },
                            "required": [
                                "barcode"
                            ],
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "txnid": {
                    "description": "Transaction UUID matching that in the blockchain.",
                    "type": "string"
                },
                "txnts": {
                    "description": "Transaction timestamp matching that in the blockchain.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "containerstatearray": {
            "description": "an array of container states, often for different assets",
            "items": {
                "patternProperties": {
                    "^CON": {
                        "description": "The external state of one container asset.",
                        "properties": {
                            "AssetKey": {
                                "description": "The World State asset ID. Used to read and write state.",
                                "type": "string"
                            },
                            "alerts": {
                                "description": "List of alert names that are currently active.",
                                "items": {
                                    "description": "Container alerts. Triggered or cleared by contract rules.",
                                    "enum": [
                                        "OVERTTEMP"
                                    ],
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetIDpath": {
                                "description": "Qualified property path to the asset's ID.",
                                "type": "string"
                            },
                            "class": {
                                "description": "Asset class.",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The event that created this state.",
                                "properties": {
                                    "container": {
                                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "barcode": {
                                                "description": "The ID of a container.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "transport entity currently in possession of the container",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                "properties": {
                                                    "devicetimestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
                                                    },
                                                    "extension": {
                                                        "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                        "items": {
                                                            "properties": {},
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "references": {
                                                        "description": "An array of external references relevant to this asset.",
                                                        "items": {
                                                            "type": "string"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of the inside of the container in CELSIUS.",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "barcode"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "A chaincode event to be emitted on invoke exit.",
                                "properties": {
                                    "container": {
                                        "description": "An event that is emitted at the end of an invoke, if it is present.",
                                        "properties": {
                                            "name": {
                                                "description": "Event name.",
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "A json object containing the event's properties.",
                                                "properties": {},
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "prefix": {
                                "description": "Asset class prefix in World State.",
                                "type": "string"
                            },
                            "state": {
                                "description": "The state of the one asset.",
                                "properties": {
                                    "container": {
                                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "barcode": {
                                                "description": "The ID of a container.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "transport entity currently in possession of the container",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                                "properties": {
                                                    "devicetimestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
                                                    },
                                                    "extension": {
                                                        "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                        "items": {
                                                            "properties": {},
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "location": {
                                                        "description": "A geographical coordinate",
                                                        "properties": {
                                                            "latitude": {
                                                                "type": "number"
                                                            },
                                                            "longitude": {
                                                                "type": "number"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "references": {
                                                        "description": "An array of external references relevant to this asset.",
                                                        "items": {
                                                            "type": "string"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of the inside of the container in CELSIUS.",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "barcode"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    }
                },
                "type": "object"
            },
            "minItems": 0,
            "type": "array"
        },
        "containerstateexternal": {
            "patternProperties": {
                "^CON": {
                    "description": "The external state of one container asset.",
                    "properties": {
                        "AssetKey": {
                            "description": "The World State asset ID. Used to read and write state.",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "List of alert names that are currently active.",
                            "items": {
                                "description": "Container alerts. Triggered or cleared by contract rules.",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the asset's ID.",
                            "type": "string"
                        },
                        "class": {
                            "description": "Asset class.",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "A contract-specific indication that this asset is compliant.",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The event that created this state.",
                            "properties": {
                                "container": {
                                    "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                    "properties": {
                                        "barcode": {
                                            "description": "The ID of a container.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "transport entity currently in possession of the container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                            "properties": {
                                                "devicetimestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                },
                                                "extension": {
                                                    "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                    "items": {
                                                        "properties": {},
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "references": {
                                                    "description": "An array of external references relevant to this asset.",
                                                    "items": {
                                                        "type": "string"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of the inside of the container in CELSIUS.",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "A chaincode event to be emitted on invoke exit.",
                            "properties": {
                                "container": {
                                    "description": "An event that is emitted at the end of an invoke, if it is present.",
                                    "properties": {
                                        "name": {
                                            "description": "Event name.",
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "A json object containing the event's properties.",
                                            "properties": {},
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "prefix": {
                            "description": "Asset class prefix in World State.",
                            "type": "string"
                        },
                        "state": {
                            "description": "The state of the one asset.",
                            "properties": {
                                "container": {
                                    "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                                    "properties": {
                                        "barcode": {
                                            "description": "The ID of a container.",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "transport entity currently in possession of the container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
                                            "properties": {
                                                "devicetimestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                },
                                                "extension": {
                                                    "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                                                    "items": {
                                                        "properties": {},
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "location": {
                                                    "description": "A geographical coordinate",
                                                    "properties": {
                                                        "latitude": {
                                                            "type": "number"
                                                        },
                                                        "longitude": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "references": {
                                                    "description": "An array of external references relevant to this asset.",
                                                    "items": {
                                                        "type": "string"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of the inside of the container in CELSIUS.",
                                            "type": "number"
                                        }
                                    },
                                    "required": [
                                        "barcode"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "txnid": {
                            "description": "Transaction UUID matching that in the blockchain.",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching that in the blockchain.",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "invokeevent": {
            "description": "An event that is emitted at the end of an invoke, if it is present.",
            "properties": {
                "name": {
                    "description": "Event name.",
                    "type": "string"
                },
                "payload": {
                    "description": "A json object containing the event's properties.",
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "ioteventcommon": {
            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
            "properties": {
                "devicetimestamp": {
                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                    "type": "string"
                },
                "extension": {
                    "description": "Application managed array of extension properties. Opaque to contract. To be used in emergencies or for sidecar information that is not relevant to contract rule processing.",
                    "items": {
                        "properties": {},
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "location": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "latitude": {
                            "type": "number"
                        },
                        "longitude": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                },
                "references": {
                    "description": "An array of external references relevant to this asset.",
                    "items": {
                        "type": "string"
                    },
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        }
    }
}`


	var readAssetSchemas as.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		as.AddRoute("readAssetSchemas", "query", as.SystemClass, readAssetSchemas)
	}
	