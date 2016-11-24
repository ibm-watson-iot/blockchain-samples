package main


	import (
		"github.com/hyperledger/fabric/core/chaincode/shim"
		iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)

var schemas = `

{
    "API": {
        "createAssetSurgicalKit": {
            "description": "Creates a new surgicalkit (e.g. put new)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                                "properties": {
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
                                    "hospital": {
                                        "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                        "properties": {
                                            "address": {
                                                "properties": {
                                                    "city": {
                                                        "type": "string"
                                                    },
                                                    "country": {
                                                        "type": "string"
                                                    },
                                                    "postcode": {
                                                        "type": "string"
                                                    },
                                                    "streetandnumber": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "fence": {
                                                "properties": {
                                                    "center": {
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
                                                    "radius": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sensors": {
                                        "description": "sensor readings for the surgical kit",
                                        "properties": {
                                            "begin": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "currtilt": {
                                                "description": "The current tilt that the kit is experiencing",
                                                "type": "number"
                                            },
                                            "end": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "endlocation": {
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
                                            "maxgforce": {
                                                "description": "The highest (in Gs) force that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "maxtilt": {
                                                "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "startlocation": {
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
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
                                        "type": "string"
                                    },
                                    "status": {
                                        "description": "current kit status as a named entity in possession of the kit",
                                        "enum": [
                                            "",
                                            "oem",
                                            "warehouse",
                                            "dealer",
                                            "retailer",
                                            "hospital",
                                            "scrapped"
                                        ],
                                        "type": "string"
                                    },
                                    "transit": {
                                        "description": "shipping data during transit periods",
                                        "properties": {
                                            "begintransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "type": "string"
                                            },
                                            "endtransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "receiver": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            },
                                            "shipper": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "required": [
                                    "skitID"
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
                        "createAssetSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsSurgicalKit": {
            "description": "Delete all surgicalkits from world state, supports filters",
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
                        "deleteAllAssetsSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetStateHistorySurgicalKit": {
            "description": "Delete a surgicalkit's history from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "properties": {
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
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
                        "deleteAssetStateHistorySurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetSurgicalKit": {
            "description": "Delete a surgicalkit from world state, transactions remain on the blockchain",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "properties": {
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
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
                        "deleteAssetSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetSurgicalKit": {
            "description": "Delete one or more properties from a surgicalkit's state, an example being temperature, which is only relevant for sensitive (as in frozen) shipments",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "qprops": {
                                "description": "Qualified property names, e.g. surgicalkit.skitID",
                                "items": {
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "surgicalkit": {
                                "properties": {
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
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
                        "deletePropertiesFromAssetSurgicalKit"
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
        "readAllAssetsSurgicalKit": {
            "description": "Returns the state of all surgicalkits, supports filters",
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
                        "readAllAssetsSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of surgicalkit states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one surgicalkit asset, named by its world state ID",
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
        "readAssetStateHistorySurgicalKit": {
            "description": "Returns history states for a surgicalkit",
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
                            "surgicalkit"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "enum": [
                        "readAssetStateHistorySurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "Array of surgicalkit states, can mix asset classes",
                    "items": {
                        "patternProperties": {
                            "^CON": {
                                "description": "The external state of one surgicalkit asset, named by its world state ID",
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
        "readAssetSurgicalKit": {
            "description": "Returns the state a surgicalkit",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "properties": {
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
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
                        "readAssetSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A surgicalkit's complete state",
                    "properties": {
                        "AssetKey": {
                            "description": "This surgicalkit's world state surgicalkit ID",
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
                            "description": "Qualified property path to the surgicalkit's ID, declared in the contract code",
                            "type": "string"
                        },
                        "class": {
                            "description": "The surgicalkit's asset class",
                            "type": "string"
                        },
                        "compliant": {
                            "description": "This surgicalkit has no active alerts",
                            "type": "boolean"
                        },
                        "eventin": {
                            "description": "The contract event that created this state, for example updateAssetSurgicalKit",
                            "properties": {
                                "surgicalkit": {
                                    "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                                    "properties": {
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
                                        "hospital": {
                                            "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                            "properties": {
                                                "address": {
                                                    "properties": {
                                                        "city": {
                                                            "type": "string"
                                                        },
                                                        "country": {
                                                            "type": "string"
                                                        },
                                                        "postcode": {
                                                            "type": "string"
                                                        },
                                                        "streetandnumber": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "fence": {
                                                    "properties": {
                                                        "center": {
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
                                                        "radius": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "name": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sensors": {
                                            "description": "sensor readings for the surgical kit",
                                            "properties": {
                                                "begin": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "currtilt": {
                                                    "description": "The current tilt that the kit is experiencing",
                                                    "type": "number"
                                                },
                                                "end": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "endlocation": {
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
                                                "maxgforce": {
                                                    "description": "The highest (in Gs) force that the kit experienced during the sample",
                                                    "type": "number"
                                                },
                                                "maxtilt": {
                                                    "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                                    "type": "number"
                                                },
                                                "startlocation": {
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
                                        "skitID": {
                                            "description": "A surgicalkit's ID",
                                            "type": "string"
                                        },
                                        "status": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        },
                                        "transit": {
                                            "description": "shipping data during transit periods",
                                            "properties": {
                                                "begintransit": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "carrier": {
                                                    "type": "string"
                                                },
                                                "endtransit": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "receiver": {
                                                    "description": "current kit status as a named entity in possession of the kit",
                                                    "enum": [
                                                        "",
                                                        "oem",
                                                        "warehouse",
                                                        "dealer",
                                                        "retailer",
                                                        "hospital",
                                                        "scrapped"
                                                    ],
                                                    "type": "string"
                                                },
                                                "shipper": {
                                                    "description": "current kit status as a named entity in possession of the kit",
                                                    "enum": [
                                                        "",
                                                        "oem",
                                                        "warehouse",
                                                        "dealer",
                                                        "retailer",
                                                        "hospital",
                                                        "scrapped"
                                                    ],
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "required": [
                                        "skitID"
                                    ],
                                    "type": "object"
                                }
                            },
                            "type": "object"
                        },
                        "eventout": {
                            "description": "The chaincode event emitted on invoke exit, if any",
                            "properties": {
                                "surgicalkit": {
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
                            "description": "The surgicalkit's asset class prefix in world state",
                            "type": "string"
                        },
                        "state": {
                            "description": "Properties that have been received or calculated for this surgicalkit",
                            "properties": {
                                "distanceFromCenter": {
                                    "description": "calculated distance from the fence center, can be compared to fence radius",
                                    "type": "number"
                                },
                                "surgicalkit": {
                                    "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                                    "properties": {
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
                                        "hospital": {
                                            "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                            "properties": {
                                                "address": {
                                                    "properties": {
                                                        "city": {
                                                            "type": "string"
                                                        },
                                                        "country": {
                                                            "type": "string"
                                                        },
                                                        "postcode": {
                                                            "type": "string"
                                                        },
                                                        "streetandnumber": {
                                                            "type": "string"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "fence": {
                                                    "properties": {
                                                        "center": {
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
                                                        "radius": {
                                                            "type": "number"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "name": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "sensors": {
                                            "description": "sensor readings for the surgical kit",
                                            "properties": {
                                                "begin": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "currtilt": {
                                                    "description": "The current tilt that the kit is experiencing",
                                                    "type": "number"
                                                },
                                                "end": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "endlocation": {
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
                                                "maxgforce": {
                                                    "description": "The highest (in Gs) force that the kit experienced during the sample",
                                                    "type": "number"
                                                },
                                                "maxtilt": {
                                                    "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                                    "type": "number"
                                                },
                                                "startlocation": {
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
                                        "skitID": {
                                            "description": "A surgicalkit's ID",
                                            "type": "string"
                                        },
                                        "status": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        },
                                        "transit": {
                                            "description": "shipping data during transit periods",
                                            "properties": {
                                                "begintransit": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "carrier": {
                                                    "type": "string"
                                                },
                                                "endtransit": {
                                                    "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                    "format": "date-time",
                                                    "sample": "yyyy-mm-dd hh:mm:ss",
                                                    "type": "string"
                                                },
                                                "receiver": {
                                                    "description": "current kit status as a named entity in possession of the kit",
                                                    "enum": [
                                                        "",
                                                        "oem",
                                                        "warehouse",
                                                        "dealer",
                                                        "retailer",
                                                        "hospital",
                                                        "scrapped"
                                                    ],
                                                    "type": "string"
                                                },
                                                "shipper": {
                                                    "description": "current kit status as a named entity in possession of the kit",
                                                    "enum": [
                                                        "",
                                                        "oem",
                                                        "warehouse",
                                                        "dealer",
                                                        "retailer",
                                                        "hospital",
                                                        "scrapped"
                                                    ],
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "required": [
                                        "skitID"
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
        "replaceAssetSurgicalKit": {
            "description": "Replaces a surgicalkit's state (e.g. put existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                                "properties": {
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
                                    "hospital": {
                                        "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                        "properties": {
                                            "address": {
                                                "properties": {
                                                    "city": {
                                                        "type": "string"
                                                    },
                                                    "country": {
                                                        "type": "string"
                                                    },
                                                    "postcode": {
                                                        "type": "string"
                                                    },
                                                    "streetandnumber": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "fence": {
                                                "properties": {
                                                    "center": {
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
                                                    "radius": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sensors": {
                                        "description": "sensor readings for the surgical kit",
                                        "properties": {
                                            "begin": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "currtilt": {
                                                "description": "The current tilt that the kit is experiencing",
                                                "type": "number"
                                            },
                                            "end": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "endlocation": {
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
                                            "maxgforce": {
                                                "description": "The highest (in Gs) force that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "maxtilt": {
                                                "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "startlocation": {
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
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
                                        "type": "string"
                                    },
                                    "status": {
                                        "description": "current kit status as a named entity in possession of the kit",
                                        "enum": [
                                            "",
                                            "oem",
                                            "warehouse",
                                            "dealer",
                                            "retailer",
                                            "hospital",
                                            "scrapped"
                                        ],
                                        "type": "string"
                                    },
                                    "transit": {
                                        "description": "shipping data during transit periods",
                                        "properties": {
                                            "begintransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "type": "string"
                                            },
                                            "endtransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "receiver": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            },
                                            "shipper": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "required": [
                                    "skitID"
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
                        "replaceAssetSurgicalKit"
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
        "updateAssetSurgicalKit": {
            "description": "Update a contaner's state with one or more property changes (e.g. patch existing)",
            "properties": {
                "args": {
                    "items": {
                        "properties": {
                            "surgicalkit": {
                                "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                                "properties": {
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
                                    "hospital": {
                                        "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                        "properties": {
                                            "address": {
                                                "properties": {
                                                    "city": {
                                                        "type": "string"
                                                    },
                                                    "country": {
                                                        "type": "string"
                                                    },
                                                    "postcode": {
                                                        "type": "string"
                                                    },
                                                    "streetandnumber": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "fence": {
                                                "properties": {
                                                    "center": {
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
                                                    "radius": {
                                                        "type": "number"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "name": {
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "sensors": {
                                        "description": "sensor readings for the surgical kit",
                                        "properties": {
                                            "begin": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "currtilt": {
                                                "description": "The current tilt that the kit is experiencing",
                                                "type": "number"
                                            },
                                            "end": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "endlocation": {
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
                                            "maxgforce": {
                                                "description": "The highest (in Gs) force that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "maxtilt": {
                                                "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                                "type": "number"
                                            },
                                            "startlocation": {
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
                                    "skitID": {
                                        "description": "A surgicalkit's ID",
                                        "type": "string"
                                    },
                                    "status": {
                                        "description": "current kit status as a named entity in possession of the kit",
                                        "enum": [
                                            "",
                                            "oem",
                                            "warehouse",
                                            "dealer",
                                            "retailer",
                                            "hospital",
                                            "scrapped"
                                        ],
                                        "type": "string"
                                    },
                                    "transit": {
                                        "description": "shipping data during transit periods",
                                        "properties": {
                                            "begintransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "carrier": {
                                                "type": "string"
                                            },
                                            "endtransit": {
                                                "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                                "format": "date-time",
                                                "sample": "yyyy-mm-dd hh:mm:ss",
                                                "type": "string"
                                            },
                                            "receiver": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            },
                                            "shipper": {
                                                "description": "current kit status as a named entity in possession of the kit",
                                                "enum": [
                                                    "",
                                                    "oem",
                                                    "warehouse",
                                                    "dealer",
                                                    "retailer",
                                                    "hospital",
                                                    "scrapped"
                                                ],
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "required": [
                                    "skitID"
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
                        "updateAssetSurgicalKit"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "Model": {
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
        },
        "surgicalkit": {
            "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
            "properties": {
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
                "hospital": {
                    "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                    "properties": {
                        "address": {
                            "properties": {
                                "city": {
                                    "type": "string"
                                },
                                "country": {
                                    "type": "string"
                                },
                                "postcode": {
                                    "type": "string"
                                },
                                "streetandnumber": {
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "fence": {
                            "properties": {
                                "center": {
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
                                "radius": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        },
                        "name": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "sensors": {
                    "description": "sensor readings for the surgical kit",
                    "properties": {
                        "begin": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        },
                        "currtilt": {
                            "description": "The current tilt that the kit is experiencing",
                            "type": "number"
                        },
                        "end": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        },
                        "endlocation": {
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
                        "maxgforce": {
                            "description": "The highest (in Gs) force that the kit experienced during the sample",
                            "type": "number"
                        },
                        "maxtilt": {
                            "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                            "type": "number"
                        },
                        "startlocation": {
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
                "skitID": {
                    "description": "A surgicalkit's ID",
                    "type": "string"
                },
                "status": {
                    "description": "current kit status as a named entity in possession of the kit",
                    "enum": [
                        "",
                        "oem",
                        "warehouse",
                        "dealer",
                        "retailer",
                        "hospital",
                        "scrapped"
                    ],
                    "type": "string"
                },
                "transit": {
                    "description": "shipping data during transit periods",
                    "properties": {
                        "begintransit": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        },
                        "carrier": {
                            "type": "string"
                        },
                        "endtransit": {
                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                            "format": "date-time",
                            "sample": "yyyy-mm-dd hh:mm:ss",
                            "type": "string"
                        },
                        "receiver": {
                            "description": "current kit status as a named entity in possession of the kit",
                            "enum": [
                                "",
                                "oem",
                                "warehouse",
                                "dealer",
                                "retailer",
                                "hospital",
                                "scrapped"
                            ],
                            "type": "string"
                        },
                        "shipper": {
                            "description": "current kit status as a named entity in possession of the kit",
                            "enum": [
                                "",
                                "oem",
                                "warehouse",
                                "dealer",
                                "retailer",
                                "hospital",
                                "scrapped"
                            ],
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "required": [
                "skitID"
            ],
            "type": "object"
        },
        "surgicalkitstate": {
            "description": "A surgicalkit's complete state",
            "properties": {
                "AssetKey": {
                    "description": "This surgicalkit's world state surgicalkit ID",
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
                    "description": "Qualified property path to the surgicalkit's ID, declared in the contract code",
                    "type": "string"
                },
                "class": {
                    "description": "The surgicalkit's asset class",
                    "type": "string"
                },
                "compliant": {
                    "description": "This surgicalkit has no active alerts",
                    "type": "boolean"
                },
                "eventin": {
                    "description": "The contract event that created this state, for example updateAssetSurgicalKit",
                    "properties": {
                        "surgicalkit": {
                            "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                            "properties": {
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
                                "hospital": {
                                    "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                    "properties": {
                                        "address": {
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "postcode": {
                                                    "type": "string"
                                                },
                                                "streetandnumber": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "fence": {
                                            "properties": {
                                                "center": {
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
                                                "radius": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "name": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "sensors": {
                                    "description": "sensor readings for the surgical kit",
                                    "properties": {
                                        "begin": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "currtilt": {
                                            "description": "The current tilt that the kit is experiencing",
                                            "type": "number"
                                        },
                                        "end": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "endlocation": {
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
                                        "maxgforce": {
                                            "description": "The highest (in Gs) force that the kit experienced during the sample",
                                            "type": "number"
                                        },
                                        "maxtilt": {
                                            "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                            "type": "number"
                                        },
                                        "startlocation": {
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
                                "skitID": {
                                    "description": "A surgicalkit's ID",
                                    "type": "string"
                                },
                                "status": {
                                    "description": "current kit status as a named entity in possession of the kit",
                                    "enum": [
                                        "",
                                        "oem",
                                        "warehouse",
                                        "dealer",
                                        "retailer",
                                        "hospital",
                                        "scrapped"
                                    ],
                                    "type": "string"
                                },
                                "transit": {
                                    "description": "shipping data during transit periods",
                                    "properties": {
                                        "begintransit": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "type": "string"
                                        },
                                        "endtransit": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "receiver": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        },
                                        "shipper": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "required": [
                                "skitID"
                            ],
                            "type": "object"
                        }
                    },
                    "type": "object"
                },
                "eventout": {
                    "description": "The chaincode event emitted on invoke exit, if any",
                    "properties": {
                        "surgicalkit": {
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
                    "description": "The surgicalkit's asset class prefix in world state",
                    "type": "string"
                },
                "state": {
                    "description": "Properties that have been received or calculated for this surgicalkit",
                    "properties": {
                        "distanceFromCenter": {
                            "description": "calculated distance from the fence center, can be compared to fence radius",
                            "type": "number"
                        },
                        "surgicalkit": {
                            "description": "The changeable properties for a surgicalkit, also considered its 'event' as a partial state",
                            "properties": {
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
                                "hospital": {
                                    "description": "the hospital within which the surgical kit is used, and within which it is geofenced",
                                    "properties": {
                                        "address": {
                                            "properties": {
                                                "city": {
                                                    "type": "string"
                                                },
                                                "country": {
                                                    "type": "string"
                                                },
                                                "postcode": {
                                                    "type": "string"
                                                },
                                                "streetandnumber": {
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "fence": {
                                            "properties": {
                                                "center": {
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
                                                "radius": {
                                                    "type": "number"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "name": {
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "sensors": {
                                    "description": "sensor readings for the surgical kit",
                                    "properties": {
                                        "begin": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "currtilt": {
                                            "description": "The current tilt that the kit is experiencing",
                                            "type": "number"
                                        },
                                        "end": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "endlocation": {
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
                                        "maxgforce": {
                                            "description": "The highest (in Gs) force that the kit experienced during the sample",
                                            "type": "number"
                                        },
                                        "maxtilt": {
                                            "description": "The highest (in degrees from horizontal) tilt that the kit experienced during the sample",
                                            "type": "number"
                                        },
                                        "startlocation": {
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
                                "skitID": {
                                    "description": "A surgicalkit's ID",
                                    "type": "string"
                                },
                                "status": {
                                    "description": "current kit status as a named entity in possession of the kit",
                                    "enum": [
                                        "",
                                        "oem",
                                        "warehouse",
                                        "dealer",
                                        "retailer",
                                        "hospital",
                                        "scrapped"
                                    ],
                                    "type": "string"
                                },
                                "transit": {
                                    "description": "shipping data during transit periods",
                                    "properties": {
                                        "begintransit": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "carrier": {
                                            "type": "string"
                                        },
                                        "endtransit": {
                                            "description": "timestamp formatted yyyy-mm-dd hh:mm:ss",
                                            "format": "date-time",
                                            "sample": "yyyy-mm-dd hh:mm:ss",
                                            "type": "string"
                                        },
                                        "receiver": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        },
                                        "shipper": {
                                            "description": "current kit status as a named entity in possession of the kit",
                                            "enum": [
                                                "",
                                                "oem",
                                                "warehouse",
                                                "dealer",
                                                "retailer",
                                                "hospital",
                                                "scrapped"
                                            ],
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                }
                            },
                            "required": [
                                "skitID"
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
        "surgicalkitstatearray": {
            "description": "Array of surgicalkit states, can mix asset classes",
            "items": {
                "patternProperties": {
                    "^CON": {
                        "description": "The external state of one surgicalkit asset, named by its world state ID",
                        "type": "object"
                    }
                },
                "type": "object"
            },
            "minItems": 0,
            "type": "array"
        },
        "surgicalkitstateexternal": {
            "patternProperties": {
                "^CON": {
                    "description": "The external state of one surgicalkit asset, named by its world state ID",
                    "type": "object"
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
	