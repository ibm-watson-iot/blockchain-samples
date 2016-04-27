package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. AssetID is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
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
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "assetID"
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
                }
            },
            "type": "object"
        },
        "deleteAllAssets": {
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
                }
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an assetID.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an assetID for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "deleteAsset function",
                    "enum": [
                        "deleteAsset"
                    ],
                    "type": "string"
                }
            },
            "type": "object"
        },
        "deletePropertiesFromAsset": {
            "description": "Delete one or more properties from an asset. Argument is a JSON encoded string containing an AssetID and an array of qualified property names. An example would be {'assetID':'A1',['event.common.carrier', 'event.customer.temperature']} and the result of that invoke would be the removal of the carrier field and the temperature field with a recalculation of the alert and compliance status.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested assetID with a list or qualified property names.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "qualPropsToDelete": {
                                "items": {
                                    "description": "The qualified name of a property. E.g. 'event.common.carrier', 'event.custom.temperature', etc.",
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "required": [
                            "assetID",
                            "qualPropsToDelete"
                        ],
                        "type": "object"
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
                }
            },
            "type": "object"
        },
        "init": {
            "description": "Initializes the contract when started, either by deployment or by peer restart.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "event sent to init on deployment",
                        "properties": {
                            "nickname": {
                                "default": "TRADELANE",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "version"
                        ],
                        "type": "object"
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
                }
            },
            "type": "object"
        },
        "readAllAssets": {
            "description": "Returns the state of all assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by assetID.",
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
                "result": {
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "A set of fields that constitute the complete asset state.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "inCompliance": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
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
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
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
        "readAsset": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. AssetID is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an assetID for use as an argument to read or delete.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "readAsset function",
                    "enum": [
                        "readAsset"
                    ],
                    "type": "string"
                },
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "OVERTTEMP"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "cleared": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "OVERTTEMP"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "raised": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "OVERTTEMP"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                }
                            },
                            "type": "object"
                        },
                        "assetID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "carrier": {
                            "description": "transport entity currently in possession of asset",
                            "type": "string"
                        },
                        "extension": {
                            "description": "Application-managed state. Opaque to contract.",
                            "properties": {},
                            "type": "object"
                        },
                        "inCompliance": {
                            "description": "A contract-specific indication that this asset is compliant.",
                            "type": "boolean"
                        },
                        "lastEvent": {
                            "description": "function and string parameter that created this state object",
                            "properties": {
                                "args": {
                                    "items": {
                                        "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                        "type": "string"
                                    },
                                    "type": "array"
                                },
                                "function": {
                                    "description": "function that created this state object",
                                    "type": "string"
                                },
                                "redirectedFromFunction": {
                                    "description": "function that originally received the event",
                                    "type": "string"
                                }
                            },
                            "type": "object"
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
                        "temperature": {
                            "description": "Temperature of the asset in CELSIUS.",
                            "type": "number"
                        },
                        "timestamp": {
                            "description": "RFC3339nanos formatted timestamp.",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistory": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. AssetID is required and count is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested assetID with item count.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "count": {
                                "type": "integer"
                            }
                        },
                        "required": [
                            "assetID"
                        ],
                        "type": "object"
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
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of fields that constitute the complete asset state.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "inCompliance": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
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
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
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
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of fields that constitute the complete asset state.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERTTEMP"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "inCompliance": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "args": {
                                        "items": {
                                            "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "function": {
                                        "description": "function that created this state object",
                                        "type": "string"
                                    },
                                    "redirectedFromFunction": {
                                        "description": "function that originally received the event",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
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
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
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
        "setCreateOnUpdate": {
            "description": "Allow updateAsset to redirect to createAsset when assetID does not exist.",
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
                }
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
                }
            },
            "type": "object"
        },
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. AssetID is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
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
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "assetID"
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
                }
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "assetIDKey": {
            "description": "An object containing only an assetID for use as an argument to read or delete.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "assetIDandCount": {
            "description": "Requested assetID with item count.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "count": {
                    "type": "integer"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {},
                    "type": "object"
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
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                },
                "timestamp": {
                    "description": "RFC3339nanos formatted timestamp.",
                    "type": "string"
                }
            },
            "required": [
                "assetID"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "TRADELANE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete asset state.",
            "properties": {
                "alerts": {
                    "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                    "properties": {
                        "active": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "cleared": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "raised": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "OVERTTEMP"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
                    "type": "string"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {},
                    "type": "object"
                },
                "inCompliance": {
                    "description": "A contract-specific indication that this asset is compliant.",
                    "type": "boolean"
                },
                "lastEvent": {
                    "description": "function and string parameter that created this state object",
                    "properties": {
                        "args": {
                            "items": {
                                "description": "parameters to the function, usually args[0] is populated with a JSON encoded event object",
                                "type": "string"
                            },
                            "type": "array"
                        },
                        "function": {
                            "description": "function that created this state object",
                            "type": "string"
                        },
                        "redirectedFromFunction": {
                            "description": "function that originally received the event",
                            "type": "string"
                        }
                    },
                    "type": "object"
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
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                },
                "timestamp": {
                    "description": "RFC3339nanos formatted timestamp.",
                    "type": "string"
                }
            },
            "type": "object"
        }
    }
}`