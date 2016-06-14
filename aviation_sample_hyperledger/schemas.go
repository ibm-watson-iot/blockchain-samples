package main

var schemas = `
{
    "API": {
        "createAsset": {
            "description": "Create an asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of event objects for this contract.",
                        "properties": {
                            "oneOf": {
                                "airline": {
                                    "description": "The set of writable properties for an airline crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "airlineevent": {
                                            "description": "The set of writable properties that define an airline's state.",
                                            "properties": {
                                                "code": {
                                                    "description": "The airline 3 letter code.",
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "description": "The name of the airline.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "name",
                                                "code"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "airplane": {
                                    "description": "The set of writable properties for an airplane crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "airplaneevent": {
                                            "description": "The set of writable properties that define an airplane's state.",
                                            "properties": {
                                                "model": {
                                                    "description": "The airplane model or family name.",
                                                    "type": "string"
                                                },
                                                "variant": {
                                                    "description": "A manufacturer-specific variant identifier for this airplane.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "model"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "assembly": {
                                    "description": "The set of writable properties for an assembly crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "assemblyevent": {
                                            "description": "The set of writable properties that define an assembly's state.",
                                            "properties": {
                                                "ALRS": {
                                                    "description": "Alerting service zone.",
                                                    "type": "string"
                                                },
                                                "ATAcode": {
                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                    "type": "string"
                                                },
                                                "lifeLimitInitial": {
                                                    "description": "Initial assembly life limit.",
                                                    "type": "integer"
                                                },
                                                "lifeLimitUsed": {
                                                    "description": "Assembly life limit that has been used, including adjustments.",
                                                    "type": "integer"
                                                },
                                                "name": {
                                                    "description": "The assembly name.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "ATAcode",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "part": {
                                    "description": "The set of writable properties for a part crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "assetID"
                                            ],
                                            "type": "object"
                                        },
                                        "partevent": {
                                            "description": "The set of writable properties that define a part's state.",
                                            "properties": {
                                                "lifeLimitInitial": {
                                                    "description": "Initial part life limit.",
                                                    "type": "integer"
                                                },
                                                "lifeLimitUsed": {
                                                    "description": "Part life limit that has been used, including adjustments.",
                                                    "type": "integer"
                                                },
                                                "name": {
                                                    "description": "The part name.",
                                                    "type": "string"
                                                },
                                                "partNumber": {
                                                    "description": "Part number.",
                                                    "type": "string"
                                                },
                                                "vendorName": {
                                                    "description": "Vendor name.",
                                                    "type": "string"
                                                },
                                                "vendorNumber": {
                                                    "description": "Vendor part number.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "partNumber",
                                                "name"
                                            ],
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            }
                        },
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
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAsset": {
            "description": "Delete an asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
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
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAsset": {
            "description": "Delete one or more properties from an asset's state. Argument is a JSON encoded string containing an 'assetID' and an array of qualified property names. For example, in an event object containing common and custom properties objects, the argument might look like {'assetID':'A1',['common.location', 'custom.carrier', 'custom.temperature']} and the result of that invoke would be the removal of the location, carrier and temperature properties. The missing temperature would clear a 'OVERTEMP' alert when the rules engine runs.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' with a list of qualified property names.",
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
                },
                "method": "deploy"
            },
            "type": "object"
        },
        "readAllAssets": {
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
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "The set of state objects for this contract.",
                        "properties": {
                            "oneOf": {
                                "airline": {
                                    "description": "The set of all properties for an airline ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured for this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airlinestate": {
                                            "description": "The set of writable properties for an airline crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airlineevent": {
                                                    "description": "The set of writable properties that define an airline's state.",
                                                    "properties": {
                                                        "code": {
                                                            "description": "The airline 3 letter code.",
                                                            "type": "string"
                                                        },
                                                        "name": {
                                                            "description": "The name of the airline.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "name",
                                                        "code"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "airplanes": {
                                            "description": "An array of airplane IDs belonging to this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airplane": {
                                    "description": "The set of all properties for an airplane ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airplanestate": {
                                            "description": "The set of writable properties for an airplane crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airplaneevent": {
                                                    "description": "The set of writable properties that define an airplane's state.",
                                                    "properties": {
                                                        "model": {
                                                            "description": "The airplane model or family name.",
                                                            "type": "string"
                                                        },
                                                        "variant": {
                                                            "description": "A manufacturer-specific variant identifier for this airplane.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "model"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "assemblies": {
                                            "description": "An array of assembly IDs belonging to this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of all properties for an assembly ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "assemblystate": {
                                            "description": "The set of writable properties for an assembly crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "assemblyevent": {
                                                    "description": "The set of writable properties that define an assembly's state.",
                                                    "properties": {
                                                        "ALRS": {
                                                            "description": "Alerting service zone.",
                                                            "type": "string"
                                                        },
                                                        "ATAcode": {
                                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                            "type": "string"
                                                        },
                                                        "lifeLimitInitial": {
                                                            "description": "Initial assembly life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Assembly life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The assembly name.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "ATAcode",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "parts": {
                                            "description": "An array of part IDs belonging to this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        }
                                    },
                                    "type": "object"
                                },
                                "part": {
                                    "description": "The set of all properties for a part ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this part.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "partstate": {
                                            "description": "The set of writable properties for a part crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "assetID"
                                                    ],
                                                    "type": "object"
                                                },
                                                "partevent": {
                                                    "description": "The set of writable properties that define a part's state.",
                                                    "properties": {
                                                        "lifeLimitInitial": {
                                                            "description": "Initial part life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Part life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The part name.",
                                                            "type": "string"
                                                        },
                                                        "partNumber": {
                                                            "description": "Part number.",
                                                            "type": "string"
                                                        },
                                                        "vendorName": {
                                                            "description": "Vendor name.",
                                                            "type": "string"
                                                        },
                                                        "vendorNumber": {
                                                            "description": "Vendor part number.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "partNumber",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
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
            "description": "Returns the state an asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
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
                "method": "query",
                "result": {
                    "description": "The set of state objects for this contract.",
                    "properties": {
                        "oneOf": {
                            "airline": {
                                "description": "The set of all properties for an airline ledger state.",
                                "properties": {
                                    "activities": {
                                        "description": "An array of activity IDs that occured for this airline.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "airlinestate": {
                                        "description": "The set of writable properties for an airline crud event.",
                                        "properties": {
                                            "activityevent": {
                                                "description": "The set of writable properties that define an activity.",
                                                "properties": {
                                                    "activityDate": {
                                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                        "type": "string"
                                                    },
                                                    "activityDetail": {
                                                        "description": "Detailed description of the activity.",
                                                        "type": "string"
                                                    },
                                                    "activityNumber": {
                                                        "description": "Activity number.",
                                                        "type": "string"
                                                    },
                                                    "activityType": {
                                                        "description": "The type of the activity.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "type",
                                                    "name"
                                                ],
                                                "type": "object"
                                            },
                                            "airlineevent": {
                                                "description": "The set of writable properties that define an airline's state.",
                                                "properties": {
                                                    "code": {
                                                        "description": "The airline 3 letter code.",
                                                        "type": "string"
                                                    },
                                                    "name": {
                                                        "description": "The name of the airline.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "name",
                                                    "code"
                                                ],
                                                "type": "object"
                                            },
                                            "commonevent": {
                                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "assetID": {
                                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                    },
                                                    "timestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
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
                                    "airplanes": {
                                        "description": "An array of airplane IDs belonging to this airline.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "commonstate": {
                                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                            "compliant": {
                                                "description": "A contract-specific indication that this asset is compliant.",
                                                "type": "boolean"
                                            },
                                            "txnEvent": {
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
                                            "txntimestamp": {
                                                "description": "Transaction timestamp matching that in the blockchain.",
                                                "type": "string"
                                            },
                                            "txnuuid": {
                                                "description": "Transaction UUID matching that in the blockchain.",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "airplane": {
                                "description": "The set of all properties for an airplane ledger state.",
                                "properties": {
                                    "activities": {
                                        "description": "An array of activity IDs that occured on this airplane.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "airplanestate": {
                                        "description": "The set of writable properties for an airplane crud event.",
                                        "properties": {
                                            "activityevent": {
                                                "description": "The set of writable properties that define an activity.",
                                                "properties": {
                                                    "activityDate": {
                                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                        "type": "string"
                                                    },
                                                    "activityDetail": {
                                                        "description": "Detailed description of the activity.",
                                                        "type": "string"
                                                    },
                                                    "activityNumber": {
                                                        "description": "Activity number.",
                                                        "type": "string"
                                                    },
                                                    "activityType": {
                                                        "description": "The type of the activity.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "type",
                                                    "name"
                                                ],
                                                "type": "object"
                                            },
                                            "airplaneevent": {
                                                "description": "The set of writable properties that define an airplane's state.",
                                                "properties": {
                                                    "model": {
                                                        "description": "The airplane model or family name.",
                                                        "type": "string"
                                                    },
                                                    "variant": {
                                                        "description": "A manufacturer-specific variant identifier for this airplane.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "model"
                                                ],
                                                "type": "object"
                                            },
                                            "commonevent": {
                                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "assetID": {
                                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                    },
                                                    "timestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
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
                                    "assemblies": {
                                        "description": "An array of assembly IDs belonging to this airplane.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "commonstate": {
                                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                            "compliant": {
                                                "description": "A contract-specific indication that this asset is compliant.",
                                                "type": "boolean"
                                            },
                                            "txnEvent": {
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
                                            "txntimestamp": {
                                                "description": "Transaction timestamp matching that in the blockchain.",
                                                "type": "string"
                                            },
                                            "txnuuid": {
                                                "description": "Transaction UUID matching that in the blockchain.",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            },
                            "assembly": {
                                "description": "The set of all properties for an assembly ledger state.",
                                "properties": {
                                    "activities": {
                                        "description": "An array of activity IDs that occured on this assembly.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "assemblystate": {
                                        "description": "The set of writable properties for an assembly crud event.",
                                        "properties": {
                                            "activityevent": {
                                                "description": "The set of writable properties that define an activity.",
                                                "properties": {
                                                    "activityDate": {
                                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                        "type": "string"
                                                    },
                                                    "activityDetail": {
                                                        "description": "Detailed description of the activity.",
                                                        "type": "string"
                                                    },
                                                    "activityNumber": {
                                                        "description": "Activity number.",
                                                        "type": "string"
                                                    },
                                                    "activityType": {
                                                        "description": "The type of the activity.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "type",
                                                    "name"
                                                ],
                                                "type": "object"
                                            },
                                            "assemblyevent": {
                                                "description": "The set of writable properties that define an assembly's state.",
                                                "properties": {
                                                    "ALRS": {
                                                        "description": "Alerting service zone.",
                                                        "type": "string"
                                                    },
                                                    "ATAcode": {
                                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                        "type": "string"
                                                    },
                                                    "lifeLimitInitial": {
                                                        "description": "Initial assembly life limit.",
                                                        "type": "integer"
                                                    },
                                                    "lifeLimitUsed": {
                                                        "description": "Assembly life limit that has been used, including adjustments.",
                                                        "type": "integer"
                                                    },
                                                    "name": {
                                                        "description": "The assembly name.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "ATAcode",
                                                    "name"
                                                ],
                                                "type": "object"
                                            },
                                            "commonevent": {
                                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "assetID": {
                                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                    },
                                                    "timestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
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
                                    "commonstate": {
                                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                            "compliant": {
                                                "description": "A contract-specific indication that this asset is compliant.",
                                                "type": "boolean"
                                            },
                                            "txnEvent": {
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
                                            "txntimestamp": {
                                                "description": "Transaction timestamp matching that in the blockchain.",
                                                "type": "string"
                                            },
                                            "txnuuid": {
                                                "description": "Transaction UUID matching that in the blockchain.",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "parts": {
                                        "description": "An array of part IDs belonging to this assembly.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "part": {
                                "description": "The set of all properties for a part ledger state.",
                                "properties": {
                                    "activities": {
                                        "description": "An array of activity IDs that occured on this part.",
                                        "items": {
                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                            "type": "string"
                                        },
                                        "type": "array"
                                    },
                                    "commonstate": {
                                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                            "compliant": {
                                                "description": "A contract-specific indication that this asset is compliant.",
                                                "type": "boolean"
                                            },
                                            "txnEvent": {
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
                                            "txntimestamp": {
                                                "description": "Transaction timestamp matching that in the blockchain.",
                                                "type": "string"
                                            },
                                            "txnuuid": {
                                                "description": "Transaction UUID matching that in the blockchain.",
                                                "type": "string"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "partstate": {
                                        "description": "The set of writable properties for a part crud event.",
                                        "properties": {
                                            "activityevent": {
                                                "description": "The set of writable properties that define an activity.",
                                                "properties": {
                                                    "activityDate": {
                                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                        "type": "string"
                                                    },
                                                    "activityDetail": {
                                                        "description": "Detailed description of the activity.",
                                                        "type": "string"
                                                    },
                                                    "activityNumber": {
                                                        "description": "Activity number.",
                                                        "type": "string"
                                                    },
                                                    "activityType": {
                                                        "description": "The type of the activity.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "type",
                                                    "name"
                                                ],
                                                "type": "object"
                                            },
                                            "commonevent": {
                                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                "properties": {
                                                    "assetID": {
                                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                    },
                                                    "timestamp": {
                                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "assetID"
                                                ],
                                                "type": "object"
                                            },
                                            "partevent": {
                                                "description": "The set of writable properties that define a part's state.",
                                                "properties": {
                                                    "lifeLimitInitial": {
                                                        "description": "Initial part life limit.",
                                                        "type": "integer"
                                                    },
                                                    "lifeLimitUsed": {
                                                        "description": "Part life limit that has been used, including adjustments.",
                                                        "type": "integer"
                                                    },
                                                    "name": {
                                                        "description": "The part name.",
                                                        "type": "string"
                                                    },
                                                    "partNumber": {
                                                        "description": "Part number.",
                                                        "type": "string"
                                                    },
                                                    "vendorName": {
                                                        "description": "Vendor name.",
                                                        "type": "string"
                                                    },
                                                    "vendorNumber": {
                                                        "description": "Vendor part number.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "partNumber",
                                                    "name"
                                                ],
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            }
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetHistory": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' with item 'count'.",
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
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "The set of state objects for this contract.",
                        "properties": {
                            "oneOf": {
                                "airline": {
                                    "description": "The set of all properties for an airline ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured for this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airlinestate": {
                                            "description": "The set of writable properties for an airline crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airlineevent": {
                                                    "description": "The set of writable properties that define an airline's state.",
                                                    "properties": {
                                                        "code": {
                                                            "description": "The airline 3 letter code.",
                                                            "type": "string"
                                                        },
                                                        "name": {
                                                            "description": "The name of the airline.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "name",
                                                        "code"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "airplanes": {
                                            "description": "An array of airplane IDs belonging to this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airplane": {
                                    "description": "The set of all properties for an airplane ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airplanestate": {
                                            "description": "The set of writable properties for an airplane crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airplaneevent": {
                                                    "description": "The set of writable properties that define an airplane's state.",
                                                    "properties": {
                                                        "model": {
                                                            "description": "The airplane model or family name.",
                                                            "type": "string"
                                                        },
                                                        "variant": {
                                                            "description": "A manufacturer-specific variant identifier for this airplane.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "model"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "assemblies": {
                                            "description": "An array of assembly IDs belonging to this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of all properties for an assembly ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "assemblystate": {
                                            "description": "The set of writable properties for an assembly crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "assemblyevent": {
                                                    "description": "The set of writable properties that define an assembly's state.",
                                                    "properties": {
                                                        "ALRS": {
                                                            "description": "Alerting service zone.",
                                                            "type": "string"
                                                        },
                                                        "ATAcode": {
                                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                            "type": "string"
                                                        },
                                                        "lifeLimitInitial": {
                                                            "description": "Initial assembly life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Assembly life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The assembly name.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "ATAcode",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "parts": {
                                            "description": "An array of part IDs belonging to this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        }
                                    },
                                    "type": "object"
                                },
                                "part": {
                                    "description": "The set of all properties for a part ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this part.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "partstate": {
                                            "description": "The set of writable properties for a part crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "assetID"
                                                    ],
                                                    "type": "object"
                                                },
                                                "partevent": {
                                                    "description": "The set of writable properties that define a part's state.",
                                                    "properties": {
                                                        "lifeLimitInitial": {
                                                            "description": "Initial part life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Part life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The part name.",
                                                            "type": "string"
                                                        },
                                                        "partNumber": {
                                                            "description": "Part number.",
                                                            "type": "string"
                                                        },
                                                        "vendorName": {
                                                            "description": "Vendor name.",
                                                            "type": "string"
                                                        },
                                                        "vendorNumber": {
                                                            "description": "Vendor part number.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "partNumber",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
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
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "The set of state objects for this contract.",
                        "properties": {
                            "oneOf": {
                                "airline": {
                                    "description": "The set of all properties for an airline ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured for this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airlinestate": {
                                            "description": "The set of writable properties for an airline crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airlineevent": {
                                                    "description": "The set of writable properties that define an airline's state.",
                                                    "properties": {
                                                        "code": {
                                                            "description": "The airline 3 letter code.",
                                                            "type": "string"
                                                        },
                                                        "name": {
                                                            "description": "The name of the airline.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "name",
                                                        "code"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "airplanes": {
                                            "description": "An array of airplane IDs belonging to this airline.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airplane": {
                                    "description": "The set of all properties for an airplane ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "airplanestate": {
                                            "description": "The set of writable properties for an airplane crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "airplaneevent": {
                                                    "description": "The set of writable properties that define an airplane's state.",
                                                    "properties": {
                                                        "model": {
                                                            "description": "The airplane model or family name.",
                                                            "type": "string"
                                                        },
                                                        "variant": {
                                                            "description": "A manufacturer-specific variant identifier for this airplane.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "model"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "assemblies": {
                                            "description": "An array of assembly IDs belonging to this airplane.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of all properties for an assembly ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "assemblystate": {
                                            "description": "The set of writable properties for an assembly crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "assemblyevent": {
                                                    "description": "The set of writable properties that define an assembly's state.",
                                                    "properties": {
                                                        "ALRS": {
                                                            "description": "Alerting service zone.",
                                                            "type": "string"
                                                        },
                                                        "ATAcode": {
                                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                            "type": "string"
                                                        },
                                                        "lifeLimitInitial": {
                                                            "description": "Initial assembly life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Assembly life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The assembly name.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "ATAcode",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
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
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "parts": {
                                            "description": "An array of part IDs belonging to this assembly.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        }
                                    },
                                    "type": "object"
                                },
                                "part": {
                                    "description": "The set of all properties for a part ledger state.",
                                    "properties": {
                                        "activities": {
                                            "description": "An array of activity IDs that occured on this part.",
                                            "items": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                                "type": "string"
                                            },
                                            "type": "array"
                                        },
                                        "commonstate": {
                                            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                                "compliant": {
                                                    "description": "A contract-specific indication that this asset is compliant.",
                                                    "type": "boolean"
                                                },
                                                "txnEvent": {
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
                                                "txntimestamp": {
                                                    "description": "Transaction timestamp matching that in the blockchain.",
                                                    "type": "string"
                                                },
                                                "txnuuid": {
                                                    "description": "Transaction UUID matching that in the blockchain.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "partstate": {
                                            "description": "The set of writable properties for a part crud event.",
                                            "properties": {
                                                "activityevent": {
                                                    "description": "The set of writable properties that define an activity.",
                                                    "properties": {
                                                        "activityDate": {
                                                            "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                            "type": "string"
                                                        },
                                                        "activityDetail": {
                                                            "description": "Detailed description of the activity.",
                                                            "type": "string"
                                                        },
                                                        "activityNumber": {
                                                            "description": "Activity number.",
                                                            "type": "string"
                                                        },
                                                        "activityType": {
                                                            "description": "The type of the activity.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "type",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                },
                                                "commonevent": {
                                                    "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                                    "properties": {
                                                        "assetID": {
                                                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                        },
                                                        "timestamp": {
                                                            "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "assetID"
                                                    ],
                                                    "type": "object"
                                                },
                                                "partevent": {
                                                    "description": "The set of writable properties that define a part's state.",
                                                    "properties": {
                                                        "lifeLimitInitial": {
                                                            "description": "Initial part life limit.",
                                                            "type": "integer"
                                                        },
                                                        "lifeLimitUsed": {
                                                            "description": "Part life limit that has been used, including adjustments.",
                                                            "type": "integer"
                                                        },
                                                        "name": {
                                                            "description": "The part name.",
                                                            "type": "string"
                                                        },
                                                        "partNumber": {
                                                            "description": "Part number.",
                                                            "type": "string"
                                                        },
                                                        "vendorName": {
                                                            "description": "Vendor name.",
                                                            "type": "string"
                                                        },
                                                        "vendorNumber": {
                                                            "description": "Vendor part number.",
                                                            "type": "string"
                                                        }
                                                    },
                                                    "required": [
                                                        "partNumber",
                                                        "name"
                                                    ],
                                                    "type": "object"
                                                }
                                            },
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
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
        "updateAsset": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The set of event objects for this contract.",
                        "properties": {
                            "oneOf": {
                                "airline": {
                                    "description": "The set of writable properties for an airline crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "airlineevent": {
                                            "description": "The set of writable properties that define an airline's state.",
                                            "properties": {
                                                "code": {
                                                    "description": "The airline 3 letter code.",
                                                    "type": "string"
                                                },
                                                "name": {
                                                    "description": "The name of the airline.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "name",
                                                "code"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "airplane": {
                                    "description": "The set of writable properties for an airplane crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "airplaneevent": {
                                            "description": "The set of writable properties that define an airplane's state.",
                                            "properties": {
                                                "model": {
                                                    "description": "The airplane model or family name.",
                                                    "type": "string"
                                                },
                                                "variant": {
                                                    "description": "A manufacturer-specific variant identifier for this airplane.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "model"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "assembly": {
                                    "description": "The set of writable properties for an assembly crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "assemblyevent": {
                                            "description": "The set of writable properties that define an assembly's state.",
                                            "properties": {
                                                "ALRS": {
                                                    "description": "Alerting service zone.",
                                                    "type": "string"
                                                },
                                                "ATAcode": {
                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                    "type": "string"
                                                },
                                                "lifeLimitInitial": {
                                                    "description": "Initial assembly life limit.",
                                                    "type": "integer"
                                                },
                                                "lifeLimitUsed": {
                                                    "description": "Assembly life limit that has been used, including adjustments.",
                                                    "type": "integer"
                                                },
                                                "name": {
                                                    "description": "The assembly name.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "ATAcode",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
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
                                "part": {
                                    "description": "The set of writable properties for a part crud event.",
                                    "properties": {
                                        "activityevent": {
                                            "description": "The set of writable properties that define an activity.",
                                            "properties": {
                                                "activityDate": {
                                                    "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                    "type": "string"
                                                },
                                                "activityDetail": {
                                                    "description": "Detailed description of the activity.",
                                                    "type": "string"
                                                },
                                                "activityNumber": {
                                                    "description": "Activity number.",
                                                    "type": "string"
                                                },
                                                "activityType": {
                                                    "description": "The type of the activity.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "type",
                                                "name"
                                            ],
                                            "type": "object"
                                        },
                                        "commonevent": {
                                            "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                            "properties": {
                                                "assetID": {
                                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                                },
                                                "timestamp": {
                                                    "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "assetID"
                                            ],
                                            "type": "object"
                                        },
                                        "partevent": {
                                            "description": "The set of writable properties that define a part's state.",
                                            "properties": {
                                                "lifeLimitInitial": {
                                                    "description": "Initial part life limit.",
                                                    "type": "integer"
                                                },
                                                "lifeLimitUsed": {
                                                    "description": "Part life limit that has been used, including adjustments.",
                                                    "type": "integer"
                                                },
                                                "name": {
                                                    "description": "The part name.",
                                                    "type": "string"
                                                },
                                                "partNumber": {
                                                    "description": "Part number.",
                                                    "type": "string"
                                                },
                                                "vendorName": {
                                                    "description": "Vendor name.",
                                                    "type": "string"
                                                },
                                                "vendorNumber": {
                                                    "description": "Vendor part number.",
                                                    "type": "string"
                                                }
                                            },
                                            "required": [
                                                "partNumber",
                                                "name"
                                            ],
                                            "type": "object"
                                        }
                                    },
                                    "type": "object"
                                }
                            }
                        },
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
        "assetIDKey": {
            "description": "An object containing only an 'assetID' for use as an argument to read or delete.",
            "properties": {
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "assetIDandCount": {
            "description": "Requested 'assetID' with item 'count'.",
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
            "description": "The set of event objects for this contract.",
            "properties": {
                "oneOf": {
                    "airline": {
                        "description": "The set of writable properties for an airline crud event.",
                        "properties": {
                            "activityevent": {
                                "description": "The set of writable properties that define an activity.",
                                "properties": {
                                    "activityDate": {
                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                        "type": "string"
                                    },
                                    "activityDetail": {
                                        "description": "Detailed description of the activity.",
                                        "type": "string"
                                    },
                                    "activityNumber": {
                                        "description": "Activity number.",
                                        "type": "string"
                                    },
                                    "activityType": {
                                        "description": "The type of the activity.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "type",
                                    "name"
                                ],
                                "type": "object"
                            },
                            "airlineevent": {
                                "description": "The set of writable properties that define an airline's state.",
                                "properties": {
                                    "code": {
                                        "description": "The airline 3 letter code.",
                                        "type": "string"
                                    },
                                    "name": {
                                        "description": "The name of the airline.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "name",
                                    "code"
                                ],
                                "type": "object"
                            },
                            "commonevent": {
                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                "properties": {
                                    "assetID": {
                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                    },
                                    "timestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
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
                    "airplane": {
                        "description": "The set of writable properties for an airplane crud event.",
                        "properties": {
                            "activityevent": {
                                "description": "The set of writable properties that define an activity.",
                                "properties": {
                                    "activityDate": {
                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                        "type": "string"
                                    },
                                    "activityDetail": {
                                        "description": "Detailed description of the activity.",
                                        "type": "string"
                                    },
                                    "activityNumber": {
                                        "description": "Activity number.",
                                        "type": "string"
                                    },
                                    "activityType": {
                                        "description": "The type of the activity.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "type",
                                    "name"
                                ],
                                "type": "object"
                            },
                            "airplaneevent": {
                                "description": "The set of writable properties that define an airplane's state.",
                                "properties": {
                                    "model": {
                                        "description": "The airplane model or family name.",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "A manufacturer-specific variant identifier for this airplane.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "model"
                                ],
                                "type": "object"
                            },
                            "commonevent": {
                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                "properties": {
                                    "assetID": {
                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                    },
                                    "timestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
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
                    "assembly": {
                        "description": "The set of writable properties for an assembly crud event.",
                        "properties": {
                            "activityevent": {
                                "description": "The set of writable properties that define an activity.",
                                "properties": {
                                    "activityDate": {
                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                        "type": "string"
                                    },
                                    "activityDetail": {
                                        "description": "Detailed description of the activity.",
                                        "type": "string"
                                    },
                                    "activityNumber": {
                                        "description": "Activity number.",
                                        "type": "string"
                                    },
                                    "activityType": {
                                        "description": "The type of the activity.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "type",
                                    "name"
                                ],
                                "type": "object"
                            },
                            "assemblyevent": {
                                "description": "The set of writable properties that define an assembly's state.",
                                "properties": {
                                    "ALRS": {
                                        "description": "Alerting service zone.",
                                        "type": "string"
                                    },
                                    "ATAcode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
                                        "type": "integer"
                                    },
                                    "lifeLimitUsed": {
                                        "description": "Assembly life limit that has been used, including adjustments.",
                                        "type": "integer"
                                    },
                                    "name": {
                                        "description": "The assembly name.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "ATAcode",
                                    "name"
                                ],
                                "type": "object"
                            },
                            "commonevent": {
                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                "properties": {
                                    "assetID": {
                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                    },
                                    "timestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
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
                    "part": {
                        "description": "The set of writable properties for a part crud event.",
                        "properties": {
                            "activityevent": {
                                "description": "The set of writable properties that define an activity.",
                                "properties": {
                                    "activityDate": {
                                        "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                        "type": "string"
                                    },
                                    "activityDetail": {
                                        "description": "Detailed description of the activity.",
                                        "type": "string"
                                    },
                                    "activityNumber": {
                                        "description": "Activity number.",
                                        "type": "string"
                                    },
                                    "activityType": {
                                        "description": "The type of the activity.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "type",
                                    "name"
                                ],
                                "type": "object"
                            },
                            "commonevent": {
                                "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                "properties": {
                                    "assetID": {
                                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                    },
                                    "timestamp": {
                                        "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "assetID"
                                ],
                                "type": "object"
                            },
                            "partevent": {
                                "description": "The set of writable properties that define a part's state.",
                                "properties": {
                                    "lifeLimitInitial": {
                                        "description": "Initial part life limit.",
                                        "type": "integer"
                                    },
                                    "lifeLimitUsed": {
                                        "description": "Part life limit that has been used, including adjustments.",
                                        "type": "integer"
                                    },
                                    "name": {
                                        "description": "The part name.",
                                        "type": "string"
                                    },
                                    "partNumber": {
                                        "description": "Part number.",
                                        "type": "string"
                                    },
                                    "vendorName": {
                                        "description": "Vendor name.",
                                        "type": "string"
                                    },
                                    "vendorNumber": {
                                        "description": "Vendor part number.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "partNumber",
                                    "name"
                                ],
                                "type": "object"
                            }
                        },
                        "type": "object"
                    }
                }
            },
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
            "description": "The set of state objects for this contract.",
            "properties": {
                "oneOf": {
                    "airline": {
                        "description": "The set of all properties for an airline ledger state.",
                        "properties": {
                            "activities": {
                                "description": "An array of activity IDs that occured for this airline.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "airlinestate": {
                                "description": "The set of writable properties for an airline crud event.",
                                "properties": {
                                    "activityevent": {
                                        "description": "The set of writable properties that define an activity.",
                                        "properties": {
                                            "activityDate": {
                                                "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                "type": "string"
                                            },
                                            "activityDetail": {
                                                "description": "Detailed description of the activity.",
                                                "type": "string"
                                            },
                                            "activityNumber": {
                                                "description": "Activity number.",
                                                "type": "string"
                                            },
                                            "activityType": {
                                                "description": "The type of the activity.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "type",
                                            "name"
                                        ],
                                        "type": "object"
                                    },
                                    "airlineevent": {
                                        "description": "The set of writable properties that define an airline's state.",
                                        "properties": {
                                            "code": {
                                                "description": "The airline 3 letter code.",
                                                "type": "string"
                                            },
                                            "name": {
                                                "description": "The name of the airline.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "name",
                                            "code"
                                        ],
                                        "type": "object"
                                    },
                                    "commonevent": {
                                        "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "assetID": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                            },
                                            "timestamp": {
                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                "type": "string"
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
                            "airplanes": {
                                "description": "An array of airplane IDs belonging to this airline.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "commonstate": {
                                "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "txnEvent": {
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
                                    "txntimestamp": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnuuid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "airplane": {
                        "description": "The set of all properties for an airplane ledger state.",
                        "properties": {
                            "activities": {
                                "description": "An array of activity IDs that occured on this airplane.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "airplanestate": {
                                "description": "The set of writable properties for an airplane crud event.",
                                "properties": {
                                    "activityevent": {
                                        "description": "The set of writable properties that define an activity.",
                                        "properties": {
                                            "activityDate": {
                                                "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                "type": "string"
                                            },
                                            "activityDetail": {
                                                "description": "Detailed description of the activity.",
                                                "type": "string"
                                            },
                                            "activityNumber": {
                                                "description": "Activity number.",
                                                "type": "string"
                                            },
                                            "activityType": {
                                                "description": "The type of the activity.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "type",
                                            "name"
                                        ],
                                        "type": "object"
                                    },
                                    "airplaneevent": {
                                        "description": "The set of writable properties that define an airplane's state.",
                                        "properties": {
                                            "model": {
                                                "description": "The airplane model or family name.",
                                                "type": "string"
                                            },
                                            "variant": {
                                                "description": "A manufacturer-specific variant identifier for this airplane.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "model"
                                        ],
                                        "type": "object"
                                    },
                                    "commonevent": {
                                        "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "assetID": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                            },
                                            "timestamp": {
                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                "type": "string"
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
                            "assemblies": {
                                "description": "An array of assembly IDs belonging to this airplane.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "commonstate": {
                                "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "txnEvent": {
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
                                    "txntimestamp": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnuuid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "assembly": {
                        "description": "The set of all properties for an assembly ledger state.",
                        "properties": {
                            "activities": {
                                "description": "An array of activity IDs that occured on this assembly.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "assemblystate": {
                                "description": "The set of writable properties for an assembly crud event.",
                                "properties": {
                                    "activityevent": {
                                        "description": "The set of writable properties that define an activity.",
                                        "properties": {
                                            "activityDate": {
                                                "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                "type": "string"
                                            },
                                            "activityDetail": {
                                                "description": "Detailed description of the activity.",
                                                "type": "string"
                                            },
                                            "activityNumber": {
                                                "description": "Activity number.",
                                                "type": "string"
                                            },
                                            "activityType": {
                                                "description": "The type of the activity.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "type",
                                            "name"
                                        ],
                                        "type": "object"
                                    },
                                    "assemblyevent": {
                                        "description": "The set of writable properties that define an assembly's state.",
                                        "properties": {
                                            "ALRS": {
                                                "description": "Alerting service zone.",
                                                "type": "string"
                                            },
                                            "ATAcode": {
                                                "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                "type": "string"
                                            },
                                            "lifeLimitInitial": {
                                                "description": "Initial assembly life limit.",
                                                "type": "integer"
                                            },
                                            "lifeLimitUsed": {
                                                "description": "Assembly life limit that has been used, including adjustments.",
                                                "type": "integer"
                                            },
                                            "name": {
                                                "description": "The assembly name.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "ATAcode",
                                            "name"
                                        ],
                                        "type": "object"
                                    },
                                    "commonevent": {
                                        "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "assetID": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                            },
                                            "timestamp": {
                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                "type": "string"
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
                            "commonstate": {
                                "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "txnEvent": {
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
                                    "txntimestamp": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnuuid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "parts": {
                                "description": "An array of part IDs belonging to this assembly.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            }
                        },
                        "type": "object"
                    },
                    "part": {
                        "description": "The set of all properties for a part ledger state.",
                        "properties": {
                            "activities": {
                                "description": "An array of activity IDs that occured on this part.",
                                "items": {
                                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                    "type": "string"
                                },
                                "type": "array"
                            },
                            "commonstate": {
                                "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
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
                                    "compliant": {
                                        "description": "A contract-specific indication that this asset is compliant.",
                                        "type": "boolean"
                                    },
                                    "txnEvent": {
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
                                    "txntimestamp": {
                                        "description": "Transaction timestamp matching that in the blockchain.",
                                        "type": "string"
                                    },
                                    "txnuuid": {
                                        "description": "Transaction UUID matching that in the blockchain.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "partstate": {
                                "description": "The set of writable properties for a part crud event.",
                                "properties": {
                                    "activityevent": {
                                        "description": "The set of writable properties that define an activity.",
                                        "properties": {
                                            "activityDate": {
                                                "description": "Date the activity occured (e.g. date from aechive of activities being copied into this ledger).",
                                                "type": "string"
                                            },
                                            "activityDetail": {
                                                "description": "Detailed description of the activity.",
                                                "type": "string"
                                            },
                                            "activityNumber": {
                                                "description": "Activity number.",
                                                "type": "string"
                                            },
                                            "activityType": {
                                                "description": "The type of the activity.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "type",
                                            "name"
                                        ],
                                        "type": "object"
                                    },
                                    "commonevent": {
                                        "description": "The set of common writable properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event'.",
                                        "properties": {
                                            "assetID": {
                                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                                            },
                                            "timestamp": {
                                                "description": "Optional device timestamp. Note that the contract retains the blockchain-assigned transaction UUID and timestamp, which reflect the time that the event arrived at the Hyperledger fabric. The device timestamp has meaning that is relevant to the device, asset and application context.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "assetID"
                                        ],
                                        "type": "object"
                                    },
                                    "partevent": {
                                        "description": "The set of writable properties that define a part's state.",
                                        "properties": {
                                            "lifeLimitInitial": {
                                                "description": "Initial part life limit.",
                                                "type": "integer"
                                            },
                                            "lifeLimitUsed": {
                                                "description": "Part life limit that has been used, including adjustments.",
                                                "type": "integer"
                                            },
                                            "name": {
                                                "description": "The part name.",
                                                "type": "string"
                                            },
                                            "partNumber": {
                                                "description": "Part number.",
                                                "type": "string"
                                            },
                                            "vendorName": {
                                                "description": "Vendor name.",
                                                "type": "string"
                                            },
                                            "vendorNumber": {
                                                "description": "Vendor part number.",
                                                "type": "string"
                                            }
                                        },
                                        "required": [
                                            "partNumber",
                                            "name"
                                        ],
                                        "type": "object"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    }
                }
            },
            "type": "object"
        }
    }
}`