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
        "code": "The airline 3 letter code.",
        "name": "The name of the airline."
    },
    "airplane": {
        "model": "The airplane model or family name.",
        "variant": "A manufacturer-specific variant identifier for this airplane."
    },
    "assembly": {
        "ALRS": "Alerting service zone.",
        "ATAcode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
        "lifeLimitInitial": 789,
        "lifeLimitUsed": 789,
        "name": "The assembly name."
    },
    "event": {
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
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "part": {
        "lifeLimitInitial": 789,
        "lifeLimitUsed": 789,
        "name": "The part name.",
        "partNumber": "Part number.",
        "vendorName": "Vendor name.",
        "vendorNumber": "Vendor part number."
    },
    "state": {
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
    }
}`