package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var schemas = `

{
    "API": {
        "createAssetContainer": {
            "description": "Creates a new container (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this container",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
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
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of a container's contents in degrees Celsuis",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "createAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsContainer": {
            "description": "Delete all containers from world state, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAllAssetsContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetContainer": {
            "description": "Delete a container from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistoryContainer": {
            "description": "Delete a container's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteAssetStateHistoryContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetContainer": {
            "description": "Delete one or more properties from a container's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "qprops": {
                                "description": "Qualified property names, e.g. container.barcode",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deletePropertiesFromAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteWorldState": {
            "description": "**** WARNING *** Clears the entire contents of world state, redeploy the contract after using this, in debugging mode, will require a restart",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "deleteWorldState"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "initContract": {
            "description": "Sets contract version and nickname",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "nickname": {
                                "default": "IOT Contract Platform",
                                "description": "The nickname of the current contract instance",
                                "type": "string"
                            },
                            "version": {
                                "description": "The version number of the current contract instance",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "initContract"
                    ],
                    "type": "string"
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssetsContainer": {
            "description": "Returns the state of all containers, supports filters",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllAssetsContainer"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of container states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one container asset, named by its world state ID",
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
        "readAllRoutes": {
            "description": "Returns an array of registered API calls by function (debugging)",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAllRoutes"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "A list of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "assetIDpath": {
                                "description": "Qualified property path to the asset's ID, declared in the contract code",
                                "type": "string"
                            },
                            "class": {
                                "description": "The asset's asset class",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
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
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "An chaincode event emitted by a contract invoke",
                                        "properties": {
                                            "name": {
                                                "description": "The chaincode event's name",
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "The chaincode event's properties",
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
                                "description": "The asset's asset class prefix in world state",
                                "type": "string"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
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
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
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
            "description": "Returns the state a container",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A container's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This container's world state container ID",
                            "type": "string"
                        },
                        "alerts": {
                            "description": "A list of alert names",
                            "items": {
                                "description": "An alert name",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "assetIDpath": {
                            "description": "Qualified property path to the container's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The container's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This container has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetContainer",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
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
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
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
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "container": {
                                    "description": "An chaincode event emitted by a contract invoke",
                                    "properties": {
                                        "name": {
                                            "description": "The chaincode event's name",
                                            "type": "string"
                                        },
                                        "payload": {
                                            "description": "The chaincode event's properties",
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
                            "description": "The container's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this container",
                            "properties": {
                                "container": {
                                    "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                    "properties": {
                                        "barcode": {
                                            "description": "A container's ID",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "description": "The carrier in possession of this container",
                                            "type": "string"
                                        },
                                        "common": {
                                            "description": "Common properties for all assets",
                                            "properties": {
                                                "appdata": {
                                                    "description": "Application managed information as an array of key:value pairs",
                                                    "items": {
                                                        "properties": {
                                                            "K": {
                                                                "type": "string"
                                                            },
                                                            "V": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "minItems": 0,
                                                    "type": "array"
                                                },
                                                "deviceID": {
                                                    "description": "A unique identifier for the device that sent the current event",
                                                    "type": "string"
                                                },
                                                "devicetimestamp": {
                                                    "description": "A timestamp recoded by the device that sent the current event",
                                                    "type": "string"
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
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "temperature": {
                                            "description": "Temperature of a container's contents in degrees Celsuis",
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
                            "description": "Transaction UUID matching the blockchain",
                            "type": "string"
                        },
                        "txnts": {
                            "description": "Transaction timestamp matching the blockchain",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetStateHistoryContainer": {
            "description": "Returns history states for a container",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "daterange": {
                                "description": "if specified, dates must fall in between these values, inclusive",
                                "properties": {
                                    "begin": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    },
                                    "end": {
                                        "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                        "format": "date-time",
                                        "sample": "yyyy-mm-dd hh:mm:ss",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "filter": {
                                "description": "Filter asset states",
                                "properties": {
                                    "match": {
                                        "description": "Defines how to match properties, missing property always fails match",
                                        "enum": [
                                            "n/a",
                                            "all",
                                            "any",
                                            "none"
                                        ],
                                        "type": "string"
                                    },
                                    "select": {
                                        "description": "Qualified property names and values match",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "Qualified property to compare, for example 'asset.assetID'",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "Value to be compared",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "container"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistoryContainer"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of container states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one container asset, named by its world state ID",
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
            "description": "Returns the state of recently updated assets",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "begin": {
                                "description": "zero based beginning of range",
                                "type": "integer"
                            },
                            "end": {
                                "description": "zero based end of range, absence means to end",
                                "type": "integer"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readRecentStates"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of asset states, can mix asset classes",
                    "items": {
                        "description": "A asset's complete state",
                        "properties": {
                            "alerts": {
                                "description": "A list of alert names",
                                "items": {
                                    "description": "An alert name",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assetID": {
                                "description": "This asset's world state asset ID",
                                "type": "string"
                            },
                            "assetIDpath": {
                                "description": "Qualified property path to the asset's ID, declared in the contract code",
                                "type": "string"
                            },
                            "class": {
                                "description": "The asset's asset class",
                                "type": "string"
                            },
                            "compliant": {
                                "description": "This asset has no active alerts",
                                "type": "boolean"
                            },
                            "eventin": {
                                "description": "The contract event that created this state, for example updateAsset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
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
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "eventout": {
                                "description": "The chaincode event emitted on invoke exit, if any",
                                "properties": {
                                    "asset": {
                                        "description": "An chaincode event emitted by a contract invoke",
                                        "properties": {
                                            "name": {
                                                "description": "The chaincode event's name",
                                                "type": "string"
                                            },
                                            "payload": {
                                                "description": "The chaincode event's properties",
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
                                "description": "The asset's asset class prefix in world state",
                                "type": "string"
                            },
                            "state": {
                                "description": "Properties that have been received or calculated for this asset",
                                "properties": {
                                    "asset": {
                                        "description": "The changeable properties for an asset, also considered its 'event' as a partial state",
                                        "properties": {
                                            "assetID": {
                                                "description": "An asset's unique ID, e.g. barcode, VIN, etc.",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "description": "The carrier in possession of this asset",
                                                "type": "string"
                                            },
                                            "common": {
                                                "description": "Common properties for all assets",
                                                "properties": {
                                                    "appdata": {
                                                        "description": "Application managed information as an array of key:value pairs",
                                                        "items": {
                                                            "properties": {
                                                                "K": {
                                                                    "type": "string"
                                                                },
                                                                "V": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "minItems": 0,
                                                        "type": "array"
                                                    },
                                                    "deviceID": {
                                                        "description": "A unique identifier for the device that sent the current event",
                                                        "type": "string"
                                                    },
                                                    "devicetimestamp": {
                                                        "description": "A timestamp recoded by the device that sent the current event",
                                                        "type": "string"
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
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "temperature": {
                                                "description": "Temperature of an asset's contents in degrees Celsuis",
                                                "type": "number"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "txnid": {
                                "description": "Transaction UUID matching the blockchain",
                                "type": "string"
                            },
                            "txnts": {
                                "description": "Transaction timestamp matching the blockchain",
                                "type": "string"
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
        "readWorldState": {
            "description": "Returns the entire contents of world state",
            "properties": {
                "args": {
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readWorldState"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "replaceAssetContainer": {
            "description": "Replaces a container's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this container",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
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
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of a container's contents in degrees Celsuis",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "replaceAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setCreateOnFirstUpdate": {
            "description": "Allow updateAsset to create an asset upon receipt of its first event",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "setCreateOnFirstUpdate": {
                                "description": "Allows updates to create missing assets on first event",
                                "type": "boolean"
                            }
                        },
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "setCreateOnFirstUpdate"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "setLoggingLevel": {
            "description": "Sets the logging level for the contract",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
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
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
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
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "container": {
                                "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                                "properties": {
                                    "barcode": {
                                        "description": "A container's ID",
                                        "type": "string"
                                    },
                                    "carrier": {
                                        "description": "The carrier in possession of this container",
                                        "type": "string"
                                    },
                                    "common": {
                                        "description": "Common properties for all assets",
                                        "properties": {
                                            "appdata": {
                                                "description": "Application managed information as an array of key:value pairs",
                                                "items": {
                                                    "properties": {
                                                        "K": {
                                                            "type": "string"
                                                        },
                                                        "V": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "minItems": 0,
                                                "type": "array"
                                            },
                                            "deviceID": {
                                                "description": "A unique identifier for the device that sent the current event",
                                                "type": "string"
                                            },
                                            "devicetimestamp": {
                                                "description": "A timestamp recoded by the device that sent the current event",
                                                "type": "string"
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
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "temperature": {
                                        "description": "Temperature of a container's contents in degrees Celsuis",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "updateAssetContainer"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "Model": {
        "container": {
            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
            "properties": {
                "barcode": {
                    "description": "A container's ID",
                    "type": "string"
                },
                "carrier": {
                    "description": "The carrier in possession of this container",
                    "type": "string"
                },
                "common": {
                    "description": "Common properties for all assets",
                    "properties": {
                        "appdata": {
                            "description": "Application managed information as an array of key:value pairs",
                            "items": {
                                "properties": {
                                    "K": {
                                        "type": "string"
                                    },
                                    "V": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "deviceID": {
                            "description": "A unique identifier for the device that sent the current event",
                            "type": "string"
                        },
                        "devicetimestamp": {
                            "description": "A timestamp recoded by the device that sent the current event",
                            "type": "string"
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
                        }
                    },
                    "type": "object"
                },
                "temperature": {
                    "description": "Temperature of a container's contents in degrees Celsuis",
                    "type": "number"
                }
            },
            "required": [
                "barcode"
            ],
            "type": "object"
        },
        "containerstate": {
            "description": "A container's complete state",
            "properties": {
                "AssetKey": {
                    "description": "This container's world state container ID",
                    "type": "string"
                },
                "alerts": {
                    "description": "A list of alert names",
                    "items": {
                        "description": "An alert name",
                        "type": "string"
                    },
                    "type": "array"
                },
                "assetIDpath": {
                    "description": "Qualified property path to the container's ID, declared in the contract code",
                    "type": "string"
                },
                "class": {
                    "description": "The container's asset class",
                    "type": "string"
                },
                "compliant": {
                    "description": "This container has no active alerts",
                    "type": "boolean"
                },
                "eventin": {
                    "description": "The contract event that created this state, for example updateAssetContainer",
                    "properties": {
                        "container": {
                            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                            "properties": {
                                "barcode": {
                                    "description": "A container's ID",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "The carrier in possession of this container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "Common properties for all assets",
                                    "properties": {
                                        "appdata": {
                                            "description": "Application managed information as an array of key:value pairs",
                                            "items": {
                                                "properties": {
                                                    "K": {
                                                        "type": "string"
                                                    },
                                                    "V": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "deviceID": {
                                            "description": "A unique identifier for the device that sent the current event",
                                            "type": "string"
                                        },
                                        "devicetimestamp": {
                                            "description": "A timestamp recoded by the device that sent the current event",
                                            "type": "string"
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
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of a container's contents in degrees Celsuis",
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
                    "description": "The chaincode event emitted on invoke exit, if any",
                    "properties": {
                        "container": {
                            "description": "An chaincode event emitted by a contract invoke",
                            "properties": {
                                "name": {
                                    "description": "The chaincode event's name",
                                    "type": "string"
                                },
                                "payload": {
                                    "description": "The chaincode event's properties",
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
                    "description": "The container's asset class prefix in world state",
                    "type": "string"
                },
                "state": {
                    "description": "Properties that have been received or calculated for this container",
                    "properties": {
                        "container": {
                            "description": "The changeable properties for a container, also considered its 'event' as a partial state",
                            "properties": {
                                "barcode": {
                                    "description": "A container's ID",
                                    "type": "string"
                                },
                                "carrier": {
                                    "description": "The carrier in possession of this container",
                                    "type": "string"
                                },
                                "common": {
                                    "description": "Common properties for all assets",
                                    "properties": {
                                        "appdata": {
                                            "description": "Application managed information as an array of key:value pairs",
                                            "items": {
                                                "properties": {
                                                    "K": {
                                                        "type": "string"
                                                    },
                                                    "V": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "minItems": 0,
                                            "type": "array"
                                        },
                                        "deviceID": {
                                            "description": "A unique identifier for the device that sent the current event",
                                            "type": "string"
                                        },
                                        "devicetimestamp": {
                                            "description": "A timestamp recoded by the device that sent the current event",
                                            "type": "string"
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
                                        }
                                    },
                                    "type": "object"
                                },
                                "temperature": {
                                    "description": "Temperature of a container's contents in degrees Celsuis",
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
                    "description": "Transaction UUID matching the blockchain",
                    "type": "string"
                },
                "txnts": {
                    "description": "Transaction timestamp matching the blockchain",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "containerstatearray": {
            "description": "Array of container states, can mix asset classes",
            "items": {
                "patternProperties": {
                    "^CON": {
                        "description": "The external state of one container asset, named by its world state ID",
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
                    "description": "The external state of one container asset, named by its world state ID",
                    "type": "object"
                }
            },
            "type": "object"
        },
        "invokeevent": {
            "description": "An chaincode event emitted by a contract invoke",
            "properties": {
                "name": {
                    "description": "The chaincode event's name",
                    "type": "string"
                },
                "payload": {
                    "description": "The chaincode event's properties",
                    "properties": {},
                    "type": "object"
                }
            },
            "type": "object"
        },
        "ioteventcommon": {
            "description": "Common properties for all assets",
            "properties": {
                "appdata": {
                    "description": "Application managed information as an array of key:value pairs",
                    "items": {
                        "properties": {
                            "K": {
                                "type": "string"
                            },
                            "V": {
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "deviceID": {
                    "description": "A unique identifier for the device that sent the current event",
                    "type": "string"
                },
                "devicetimestamp": {
                    "description": "A timestamp recoded by the device that sent the current event",
                    "type": "string"
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
                }
            },
            "type": "object"
        },
        "stateFilter": {
            "description": "Filter asset states",
            "properties": {
                "match": {
                    "description": "Defines how to match properties, missing property always fails match",
                    "enum": [
                        "n/a",
                        "all",
                        "any",
                        "none"
                    ],
                    "type": "string"
                },
                "select": {
                    "description": "Qualified property names and values match",
                    "items": {
                        "properties": {
                            "qprop": {
                                "description": "Qualified property to compare, for example 'asset.assetID'",
                                "type": "string"
                            },
                            "value": {
                                "description": "Value to be compared",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "type": "array"
                }
            },
            "type": "object"
        }
    }
}`


	var readAssetSchemas iot.ChaincodeFunc = func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
		return []byte(schemas), nil
	}
	func init() {
		iot.AddRoute("readAssetSchemas", "query", iot.SystemClass, readAssetSchemas)
	}
	