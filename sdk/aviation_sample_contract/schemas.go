package main

var schemas = `
{
    "API": {
        "createAssetAircraft": {
            "description": "Create an aircraft asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The aircraft CRUD event",
                        "properties": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAssetAircraft function",
                    "enum": [
                        "createAssetAircraft"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetAirline": {
            "description": "Create an airline asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The airline CRUD event",
                        "properties": {
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAssetAirline function",
                    "enum": [
                        "createAssetAirline"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createAssetAssembly": {
            "description": "Create an assembly asset. One argument, a JSON encoded event. The 'assetID' property is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The assembly event. Note that assetID is the assembly serial number",
                        "properties": {
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createAssetAssembly function",
                    "enum": [
                        "createAssetAssembly"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsAircraft": {
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
                    "description": "deleteAllAssetsAircraft function",
                    "enum": [
                        "deleteAllAssetsAircraft"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsAirline": {
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
                    "description": "deleteAllAssetsAirline function",
                    "enum": [
                        "deleteAllAssetsAirline"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAllAssetsAssembly": {
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
                    "description": "deleteAllAssetsAssembly function",
                    "enum": [
                        "deleteAllAssetsAssembly"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetAircraft": {
            "description": "Delete an aircraft asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "deleteAssetAircraft function",
                    "enum": [
                        "deleteAssetAircraft"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetAirline": {
            "description": "Delete an airline asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "deleteAssetAirline function",
                    "enum": [
                        "deleteAssetAirline"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deleteAssetAssembly": {
            "description": "Delete an assembly asset, its history, and any recent state activity. Argument is a JSON encoded string containing only an 'assetID'.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "deleteAssetAssembly function",
                    "enum": [
                        "deleteAssetAssembly"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetAircraft": {
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
                    "description": "deletePropertiesFromAssetAircraft function",
                    "enum": [
                        "deletePropertiesFromAssetAircraft"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetAirline": {
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
                    "description": "deletePropertiesFromAssetAirline function",
                    "enum": [
                        "deletePropertiesFromAssetAirline"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "deletePropertiesFromAssetAssembly": {
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
                    "description": "deletePropertiesFromAssetAssembly function",
                    "enum": [
                        "deletePropertiesFromAssetAssembly"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "eventAnalyticAdjustment": {
            "description": "Post an analytic adjustment event to the contract. One argument, a JSON encoded event. The 'assetID' property is the target for the event, which is an assembly in the case of the analytic adjustment event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                        "properties": {
                            "analyticAdjustment": {
                                "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                "properties": {
                                    "action": {
                                        "enum": [
                                            "adjustLifeLimit"
                                        ],
                                        "type": "string"
                                    },
                                    "amount": {
                                        "type": "number"
                                    },
                                    "assembly": {
                                        "description": "Assembly serial number",
                                        "type": "string"
                                    },
                                    "reason": {
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
                    "description": "eventAnalyticAdjustment function",
                    "enum": [
                        "eventAnalyticAdjustment"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "eventFlight": {
            "description": "Post a flight event to the contract. One argument, a JSON encoded event. The 'assetID' property is the target for the event, which is an aircraft in the case of the flight event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "flight event, assetID defines airplane against which the event occurred",
                        "properties": {
                            "flight": {
                                "description": "A takeoiff and a landing",
                                "properties": {
                                    "aircraft": {
                                        "description": "Aircraft tail or serial number (tbd)",
                                        "type": "string"
                                    },
                                    "analyticHardlanding": {
                                        "description": "landing considered hard by analytics",
                                        "type": "boolean"
                                    },
                                    "atd": {
                                        "description": "actual time departure",
                                        "type": "string"
                                    },
                                    "flightnumber": {
                                        "description": "A flight number",
                                        "type": "string"
                                    },
                                    "from": {
                                        "description": "3 letter code of originating airport",
                                        "type": "string"
                                    },
                                    "gForce": {
                                        "description": "force incurred on landing",
                                        "type": "number"
                                    },
                                    "hardlanding": {
                                        "description": "landing considered hard by pilot or aircraft sensor",
                                        "type": "boolean"
                                    },
                                    "landingType": {
                                        "description": "code defining landing quality??",
                                        "type": "string"
                                    },
                                    "sta": {
                                        "description": "standard time arrival",
                                        "type": "string"
                                    },
                                    "std": {
                                        "description": "standard time departure",
                                        "type": "string"
                                    },
                                    "to": {
                                        "description": "3 letter code of terminating airport",
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
                    "description": "eventFlight function",
                    "enum": [
                        "eventFlight"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "eventInspection": {
            "description": "Post an inspection event to the contract. One argument, a JSON encoded event. The 'assetID' property is the target for the event, which is an assembly in the case of the inspection event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                        "properties": {
                            "inspection": {
                                "description": "indicates that an inspection has occured for this assembly",
                                "properties": {
                                    "action": {
                                        "description": "inspection that has been performed",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "aircraft": {
                                        "description": "Aircraft serial number",
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
                    "description": "eventInspection function",
                    "enum": [
                        "eventInspection"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "eventMaintenance": {
            "description": "Post a maintenance event to the contract. One argument, a JSON encoded event. The 'assetID' property is the target for the event, which is an assembly in the case of the maintenance event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "maintenance event",
                        "properties": {
                            "maintenance": {
                                "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                "properties": {
                                    "action": {
                                        "enum": [
                                            "commission",
                                            "install",
                                            "uninstall",
                                            "startMaintenance",
                                            "endMaintenance",
                                            "scrap"
                                        ],
                                        "type": "string"
                                    },
                                    "aircraft": {
                                        "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                        "type": "string"
                                    },
                                    "assembly": {
                                        "description": "This assembly's serial number",
                                        "type": "string"
                                    },
                                    "note": {
                                        "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "assembly",
                                    "action"
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
                    "description": "eventMaintenance function",
                    "enum": [
                        "eventMaintenance"
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
        "readAllAssetsAircraft": {
            "description": "Returns the state of all aircraft assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by 'assetID'.",
            "properties": {
                "args": {
                    "description": "accepts one optional argument, a filter",
                    "items": {
                        "description": "A filter that can be used to restrict a read operation to a subset of objects that meet a criteria.",
                        "properties": {
                            "filter": {
                                "properties": {
                                    "entries": {
                                        "description": "Array of qualified propery name and value pairs.",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "A qualified property as dot separated levels terminated by a leaf node. An example would be 'common.assetID'.",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "The value to be compared.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    },
                                    "matchhow": {
                                        "enum": [
                                            "matchall",
                                            "matchany",
                                            "matchnone"
                                        ],
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "entries",
                                    "matchhow"
                                ],
                                "type": "object"
                            }
                        },
                        "required": [
                            "filter"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAllAssetsAircraft function",
                    "enum": [
                        "readAllAssetsAircraft"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsAirline": {
            "description": "Returns the state of all airline assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by 'assetID'.",
            "properties": {
                "args": {
                    "description": "accepts one optional argument, a filter",
                    "items": {
                        "description": "A filter that can be used to restrict a read operation to a subset of objects that meet a criteria.",
                        "properties": {
                            "filter": {
                                "properties": {
                                    "entries": {
                                        "description": "Array of qualified propery name and value pairs.",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "A qualified property as dot separated levels terminated by a leaf node. An example would be 'common.assetID'.",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "The value to be compared.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    },
                                    "matchhow": {
                                        "enum": [
                                            "matchall",
                                            "matchany",
                                            "matchnone"
                                        ],
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "entries",
                                    "matchhow"
                                ],
                                "type": "object"
                            }
                        },
                        "required": [
                            "filter"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAllAssetsAirline function",
                    "enum": [
                        "readAllAssetsAirline"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAllAssetsAssembly": {
            "description": "Returns the state of all assembly assets as an array of JSON encoded strings. Accepts no arguments. For each managed asset, the state is read from the ledger and added to the returned array. Array is sorted by 'assetID'.",
            "properties": {
                "args": {
                    "description": "accepts one optional argument, a filter",
                    "items": {
                        "description": "A filter that can be used to restrict a read operation to a subset of objects that meet a criteria.",
                        "properties": {
                            "filter": {
                                "properties": {
                                    "entries": {
                                        "description": "Array of qualified propery name and value pairs.",
                                        "items": {
                                            "properties": {
                                                "qprop": {
                                                    "description": "A qualified property as dot separated levels terminated by a leaf node. An example would be 'common.assetID'.",
                                                    "type": "string"
                                                },
                                                "value": {
                                                    "description": "The value to be compared.",
                                                    "type": "string"
                                                }
                                            },
                                            "type": "object"
                                        },
                                        "type": "array"
                                    },
                                    "matchhow": {
                                        "enum": [
                                            "matchall",
                                            "matchany",
                                            "matchnone"
                                        ],
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "entries",
                                    "matchhow"
                                ],
                                "type": "object"
                            }
                        },
                        "required": [
                            "filter"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAllAssetsAssembly function",
                    "enum": [
                        "readAllAssetsAssembly"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states, often for different assets",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetAircraft": {
            "description": "Returns the state of an aircraft asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "readAssetAircraft function",
                    "enum": [
                        "readAssetAircraft"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "cleared": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "raised": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
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
                        "iotCommon": {
                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                        "lastEvent": {
                            "description": "function and string parameter that created this state object",
                            "properties": {
                                "arg": {
                                    "description": "A set of properties that constitute a message to the contract.",
                                    "properties": {
                                        "iotCommon": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                        "oneOf": {
                                            "aircraft": {
                                                "description": "The aircraft CRUD event",
                                                "properties": {
                                                    "aircraft": {
                                                        "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                        "properties": {
                                                            "airline": {
                                                                "description": "AssetID of airline that owns this airplane",
                                                                "type": "string"
                                                            },
                                                            "code": {
                                                                "description": "Aircraft code -- e.g. WN / SWA",
                                                                "type": "string"
                                                            },
                                                            "dateOfBuild": {
                                                                "description": "Aircraft build completed / in service date",
                                                                "type": "string"
                                                            },
                                                            "mode-s": {
                                                                "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                "type": "string"
                                                            },
                                                            "model": {
                                                                "description": "Aircraft model -- e.g. 737-5H4",
                                                                "type": "string"
                                                            },
                                                            "operator": {
                                                                "description": "AssetID of operator that flies this airplane",
                                                                "type": "string"
                                                            },
                                                            "serialNumber": {
                                                                "description": "Aircraft serial number (manufacturer assigned)",
                                                                "type": "string"
                                                            },
                                                            "tailNumber": {
                                                                "description": "Aircraft tail number (airline assigned)",
                                                                "type": "string"
                                                            },
                                                            "variant": {
                                                                "description": "Aircraft model variant -- e.g. B735",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "airline": {
                                                "description": "The airline CRUD event",
                                                "properties": {
                                                    "airline": {
                                                        "description": "The writable properties for an airline",
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
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "analyticAdjustment": {
                                                "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                "properties": {
                                                    "analyticAdjustment": {
                                                        "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "adjustLifeLimit"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "amount": {
                                                                "type": "number"
                                                            },
                                                            "assembly": {
                                                                "description": "Assembly serial number",
                                                                "type": "string"
                                                            },
                                                            "reason": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "assembly": {
                                                "description": "The assembly event. Note that assetID is the assembly serial number",
                                                "properties": {
                                                    "assembly": {
                                                        "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                        "properties": {
                                                            "arlsZone": {
                                                                "description": "tbd",
                                                                "type": "string"
                                                            },
                                                            "assemblyNumber": {
                                                                "description": "Assembly type identifier",
                                                                "type": "string"
                                                            },
                                                            "ataCode": {
                                                                "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                "type": "string"
                                                            },
                                                            "lifeLimitInitial": {
                                                                "description": "Initial assembly life limit.",
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
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "flight": {
                                                "description": "flight event, assetID defines airplane against which the event occurred",
                                                "properties": {
                                                    "flight": {
                                                        "description": "A takeoiff and a landing",
                                                        "properties": {
                                                            "aircraft": {
                                                                "description": "Aircraft tail or serial number (tbd)",
                                                                "type": "string"
                                                            },
                                                            "analyticHardlanding": {
                                                                "description": "landing considered hard by analytics",
                                                                "type": "boolean"
                                                            },
                                                            "atd": {
                                                                "description": "actual time departure",
                                                                "type": "string"
                                                            },
                                                            "flightnumber": {
                                                                "description": "A flight number",
                                                                "type": "string"
                                                            },
                                                            "from": {
                                                                "description": "3 letter code of originating airport",
                                                                "type": "string"
                                                            },
                                                            "gForce": {
                                                                "description": "force incurred on landing",
                                                                "type": "number"
                                                            },
                                                            "hardlanding": {
                                                                "description": "landing considered hard by pilot or aircraft sensor",
                                                                "type": "boolean"
                                                            },
                                                            "landingType": {
                                                                "description": "code defining landing quality??",
                                                                "type": "string"
                                                            },
                                                            "sta": {
                                                                "description": "standard time arrival",
                                                                "type": "string"
                                                            },
                                                            "std": {
                                                                "description": "standard time departure",
                                                                "type": "string"
                                                            },
                                                            "to": {
                                                                "description": "3 letter code of terminating airport",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "inspection": {
                                                "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                "properties": {
                                                    "inspection": {
                                                        "description": "indicates that an inspection has occured for this assembly",
                                                        "properties": {
                                                            "action": {
                                                                "description": "inspection that has been performed",
                                                                "enum": [
                                                                    "ACHECK",
                                                                    "BCHECK",
                                                                    "HARDLANDING"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "Aircraft serial number",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "maintenance": {
                                                "description": "maintenance event",
                                                "properties": {
                                                    "maintenance": {
                                                        "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "commission",
                                                                    "install",
                                                                    "uninstall",
                                                                    "startMaintenance",
                                                                    "endMaintenance",
                                                                    "scrap"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                "type": "string"
                                                            },
                                                            "assembly": {
                                                                "description": "This assembly's serial number",
                                                                "type": "string"
                                                            },
                                                            "note": {
                                                                "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "required": [
                                                            "assembly",
                                                            "action"
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
                        "oneOf": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            }
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
        "readAssetAircraftHistory": {
            "description": "Requests a specified number of history states for aircraft assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
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
                    "description": "readAssetAircraftHistory function",
                    "enum": [
                        "readAssetAircraftHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetAirline": {
            "description": "Returns the state of an airline asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "readAssetAirline function",
                    "enum": [
                        "readAssetAirline"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "cleared": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "raised": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
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
                        "iotCommon": {
                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                        "lastEvent": {
                            "description": "function and string parameter that created this state object",
                            "properties": {
                                "arg": {
                                    "description": "A set of properties that constitute a message to the contract.",
                                    "properties": {
                                        "iotCommon": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                        "oneOf": {
                                            "aircraft": {
                                                "description": "The aircraft CRUD event",
                                                "properties": {
                                                    "aircraft": {
                                                        "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                        "properties": {
                                                            "airline": {
                                                                "description": "AssetID of airline that owns this airplane",
                                                                "type": "string"
                                                            },
                                                            "code": {
                                                                "description": "Aircraft code -- e.g. WN / SWA",
                                                                "type": "string"
                                                            },
                                                            "dateOfBuild": {
                                                                "description": "Aircraft build completed / in service date",
                                                                "type": "string"
                                                            },
                                                            "mode-s": {
                                                                "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                "type": "string"
                                                            },
                                                            "model": {
                                                                "description": "Aircraft model -- e.g. 737-5H4",
                                                                "type": "string"
                                                            },
                                                            "operator": {
                                                                "description": "AssetID of operator that flies this airplane",
                                                                "type": "string"
                                                            },
                                                            "serialNumber": {
                                                                "description": "Aircraft serial number (manufacturer assigned)",
                                                                "type": "string"
                                                            },
                                                            "tailNumber": {
                                                                "description": "Aircraft tail number (airline assigned)",
                                                                "type": "string"
                                                            },
                                                            "variant": {
                                                                "description": "Aircraft model variant -- e.g. B735",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "airline": {
                                                "description": "The airline CRUD event",
                                                "properties": {
                                                    "airline": {
                                                        "description": "The writable properties for an airline",
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
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "analyticAdjustment": {
                                                "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                "properties": {
                                                    "analyticAdjustment": {
                                                        "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "adjustLifeLimit"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "amount": {
                                                                "type": "number"
                                                            },
                                                            "assembly": {
                                                                "description": "Assembly serial number",
                                                                "type": "string"
                                                            },
                                                            "reason": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "assembly": {
                                                "description": "The assembly event. Note that assetID is the assembly serial number",
                                                "properties": {
                                                    "assembly": {
                                                        "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                        "properties": {
                                                            "arlsZone": {
                                                                "description": "tbd",
                                                                "type": "string"
                                                            },
                                                            "assemblyNumber": {
                                                                "description": "Assembly type identifier",
                                                                "type": "string"
                                                            },
                                                            "ataCode": {
                                                                "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                "type": "string"
                                                            },
                                                            "lifeLimitInitial": {
                                                                "description": "Initial assembly life limit.",
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
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "flight": {
                                                "description": "flight event, assetID defines airplane against which the event occurred",
                                                "properties": {
                                                    "flight": {
                                                        "description": "A takeoiff and a landing",
                                                        "properties": {
                                                            "aircraft": {
                                                                "description": "Aircraft tail or serial number (tbd)",
                                                                "type": "string"
                                                            },
                                                            "analyticHardlanding": {
                                                                "description": "landing considered hard by analytics",
                                                                "type": "boolean"
                                                            },
                                                            "atd": {
                                                                "description": "actual time departure",
                                                                "type": "string"
                                                            },
                                                            "flightnumber": {
                                                                "description": "A flight number",
                                                                "type": "string"
                                                            },
                                                            "from": {
                                                                "description": "3 letter code of originating airport",
                                                                "type": "string"
                                                            },
                                                            "gForce": {
                                                                "description": "force incurred on landing",
                                                                "type": "number"
                                                            },
                                                            "hardlanding": {
                                                                "description": "landing considered hard by pilot or aircraft sensor",
                                                                "type": "boolean"
                                                            },
                                                            "landingType": {
                                                                "description": "code defining landing quality??",
                                                                "type": "string"
                                                            },
                                                            "sta": {
                                                                "description": "standard time arrival",
                                                                "type": "string"
                                                            },
                                                            "std": {
                                                                "description": "standard time departure",
                                                                "type": "string"
                                                            },
                                                            "to": {
                                                                "description": "3 letter code of terminating airport",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "inspection": {
                                                "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                "properties": {
                                                    "inspection": {
                                                        "description": "indicates that an inspection has occured for this assembly",
                                                        "properties": {
                                                            "action": {
                                                                "description": "inspection that has been performed",
                                                                "enum": [
                                                                    "ACHECK",
                                                                    "BCHECK",
                                                                    "HARDLANDING"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "Aircraft serial number",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "maintenance": {
                                                "description": "maintenance event",
                                                "properties": {
                                                    "maintenance": {
                                                        "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "commission",
                                                                    "install",
                                                                    "uninstall",
                                                                    "startMaintenance",
                                                                    "endMaintenance",
                                                                    "scrap"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                "type": "string"
                                                            },
                                                            "assembly": {
                                                                "description": "This assembly's serial number",
                                                                "type": "string"
                                                            },
                                                            "note": {
                                                                "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "required": [
                                                            "assembly",
                                                            "action"
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
                        "oneOf": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            }
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
        "readAssetAirlineHistory": {
            "description": "Requests a specified number of history states for airline assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
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
                    "description": "readAssetAirlineHistory function",
                    "enum": [
                        "readAssetAirlineHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readAssetAssembly": {
            "description": "Returns the state of an assembly asset. Argument is a JSON encoded string. The arg is an 'assetID' property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested 'assetID' in an object.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
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
                    "description": "readAssetAssembly function",
                    "enum": [
                        "readAssetAssembly"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "cleared": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "raised": {
                                    "items": {
                                        "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
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
                        "iotCommon": {
                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                        "lastEvent": {
                            "description": "function and string parameter that created this state object",
                            "properties": {
                                "arg": {
                                    "description": "A set of properties that constitute a message to the contract.",
                                    "properties": {
                                        "iotCommon": {
                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                        "oneOf": {
                                            "aircraft": {
                                                "description": "The aircraft CRUD event",
                                                "properties": {
                                                    "aircraft": {
                                                        "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                        "properties": {
                                                            "airline": {
                                                                "description": "AssetID of airline that owns this airplane",
                                                                "type": "string"
                                                            },
                                                            "code": {
                                                                "description": "Aircraft code -- e.g. WN / SWA",
                                                                "type": "string"
                                                            },
                                                            "dateOfBuild": {
                                                                "description": "Aircraft build completed / in service date",
                                                                "type": "string"
                                                            },
                                                            "mode-s": {
                                                                "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                "type": "string"
                                                            },
                                                            "model": {
                                                                "description": "Aircraft model -- e.g. 737-5H4",
                                                                "type": "string"
                                                            },
                                                            "operator": {
                                                                "description": "AssetID of operator that flies this airplane",
                                                                "type": "string"
                                                            },
                                                            "serialNumber": {
                                                                "description": "Aircraft serial number (manufacturer assigned)",
                                                                "type": "string"
                                                            },
                                                            "tailNumber": {
                                                                "description": "Aircraft tail number (airline assigned)",
                                                                "type": "string"
                                                            },
                                                            "variant": {
                                                                "description": "Aircraft model variant -- e.g. B735",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "airline": {
                                                "description": "The airline CRUD event",
                                                "properties": {
                                                    "airline": {
                                                        "description": "The writable properties for an airline",
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
                                                        "type": "object"
                                                    },
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "analyticAdjustment": {
                                                "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                "properties": {
                                                    "analyticAdjustment": {
                                                        "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "adjustLifeLimit"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "amount": {
                                                                "type": "number"
                                                            },
                                                            "assembly": {
                                                                "description": "Assembly serial number",
                                                                "type": "string"
                                                            },
                                                            "reason": {
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "assembly": {
                                                "description": "The assembly event. Note that assetID is the assembly serial number",
                                                "properties": {
                                                    "assembly": {
                                                        "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                        "properties": {
                                                            "arlsZone": {
                                                                "description": "tbd",
                                                                "type": "string"
                                                            },
                                                            "assemblyNumber": {
                                                                "description": "Assembly type identifier",
                                                                "type": "string"
                                                            },
                                                            "ataCode": {
                                                                "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                "type": "string"
                                                            },
                                                            "lifeLimitInitial": {
                                                                "description": "Initial assembly life limit.",
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
                                                    "common": {
                                                        "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "flight": {
                                                "description": "flight event, assetID defines airplane against which the event occurred",
                                                "properties": {
                                                    "flight": {
                                                        "description": "A takeoiff and a landing",
                                                        "properties": {
                                                            "aircraft": {
                                                                "description": "Aircraft tail or serial number (tbd)",
                                                                "type": "string"
                                                            },
                                                            "analyticHardlanding": {
                                                                "description": "landing considered hard by analytics",
                                                                "type": "boolean"
                                                            },
                                                            "atd": {
                                                                "description": "actual time departure",
                                                                "type": "string"
                                                            },
                                                            "flightnumber": {
                                                                "description": "A flight number",
                                                                "type": "string"
                                                            },
                                                            "from": {
                                                                "description": "3 letter code of originating airport",
                                                                "type": "string"
                                                            },
                                                            "gForce": {
                                                                "description": "force incurred on landing",
                                                                "type": "number"
                                                            },
                                                            "hardlanding": {
                                                                "description": "landing considered hard by pilot or aircraft sensor",
                                                                "type": "boolean"
                                                            },
                                                            "landingType": {
                                                                "description": "code defining landing quality??",
                                                                "type": "string"
                                                            },
                                                            "sta": {
                                                                "description": "standard time arrival",
                                                                "type": "string"
                                                            },
                                                            "std": {
                                                                "description": "standard time departure",
                                                                "type": "string"
                                                            },
                                                            "to": {
                                                                "description": "3 letter code of terminating airport",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "inspection": {
                                                "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                "properties": {
                                                    "inspection": {
                                                        "description": "indicates that an inspection has occured for this assembly",
                                                        "properties": {
                                                            "action": {
                                                                "description": "inspection that has been performed",
                                                                "enum": [
                                                                    "ACHECK",
                                                                    "BCHECK",
                                                                    "HARDLANDING"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "Aircraft serial number",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "type": "object"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "maintenance": {
                                                "description": "maintenance event",
                                                "properties": {
                                                    "maintenance": {
                                                        "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                        "properties": {
                                                            "action": {
                                                                "enum": [
                                                                    "commission",
                                                                    "install",
                                                                    "uninstall",
                                                                    "startMaintenance",
                                                                    "endMaintenance",
                                                                    "scrap"
                                                                ],
                                                                "type": "string"
                                                            },
                                                            "aircraft": {
                                                                "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                "type": "string"
                                                            },
                                                            "assembly": {
                                                                "description": "This assembly's serial number",
                                                                "type": "string"
                                                            },
                                                            "note": {
                                                                "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                "type": "string"
                                                            }
                                                        },
                                                        "required": [
                                                            "assembly",
                                                            "action"
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
                        "oneOf": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            }
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
        "readAssetAssemblyHistory": {
            "description": "Requests a specified number of history states for assembly assets. Returns an array of states sorted with the most recent first. The 'assetID' property is required and the count property is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
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
                    "description": "readAssetAssemblyHistory function",
                    "enum": [
                        "readAssetAssemblyHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "minItems": 0,
                    "type": "array"
                }
            },
            "type": "object"
        },
        "readContractConfig": {
            "description": "Returns the contract's dynamic configuration -- those thresholds and values that can be set by external applications.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readContractConfig function",
                    "enum": [
                        "readContractConfig"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "properties": {
                        "aCheckThreshold": {
                            "description": "Cycles threshold for the aCheck inspection alert.",
                            "type": "number"
                        },
                        "bCheckThreshold": {
                            "description": "Cycles threshold for the bCheck inspection alert.",
                            "type": "number"
                        }
                    },
                    "type": "object"
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
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "cleared": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "raised": {
                                        "items": {
                                            "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                            "enum": [
                                                "ACHECK",
                                                "BCHECK",
                                                "HARDLANDING"
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
                            "iotCommon": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                            "lastEvent": {
                                "description": "function and string parameter that created this state object",
                                "properties": {
                                    "arg": {
                                        "description": "A set of properties that constitute a message to the contract.",
                                        "properties": {
                                            "iotCommon": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                            "oneOf": {
                                                "aircraft": {
                                                    "description": "The aircraft CRUD event",
                                                    "properties": {
                                                        "aircraft": {
                                                            "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                            "properties": {
                                                                "airline": {
                                                                    "description": "AssetID of airline that owns this airplane",
                                                                    "type": "string"
                                                                },
                                                                "code": {
                                                                    "description": "Aircraft code -- e.g. WN / SWA",
                                                                    "type": "string"
                                                                },
                                                                "dateOfBuild": {
                                                                    "description": "Aircraft build completed / in service date",
                                                                    "type": "string"
                                                                },
                                                                "mode-s": {
                                                                    "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                                    "type": "string"
                                                                },
                                                                "model": {
                                                                    "description": "Aircraft model -- e.g. 737-5H4",
                                                                    "type": "string"
                                                                },
                                                                "operator": {
                                                                    "description": "AssetID of operator that flies this airplane",
                                                                    "type": "string"
                                                                },
                                                                "serialNumber": {
                                                                    "description": "Aircraft serial number (manufacturer assigned)",
                                                                    "type": "string"
                                                                },
                                                                "tailNumber": {
                                                                    "description": "Aircraft tail number (airline assigned)",
                                                                    "type": "string"
                                                                },
                                                                "variant": {
                                                                    "description": "Aircraft model variant -- e.g. B735",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "airline": {
                                                    "description": "The airline CRUD event",
                                                    "properties": {
                                                        "airline": {
                                                            "description": "The writable properties for an airline",
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
                                                            "type": "object"
                                                        },
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "analyticAdjustment": {
                                                    "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                                    "properties": {
                                                        "analyticAdjustment": {
                                                            "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "adjustLifeLimit"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "amount": {
                                                                    "type": "number"
                                                                },
                                                                "assembly": {
                                                                    "description": "Assembly serial number",
                                                                    "type": "string"
                                                                },
                                                                "reason": {
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "assembly": {
                                                    "description": "The assembly event. Note that assetID is the assembly serial number",
                                                    "properties": {
                                                        "assembly": {
                                                            "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                            "properties": {
                                                                "arlsZone": {
                                                                    "description": "tbd",
                                                                    "type": "string"
                                                                },
                                                                "assemblyNumber": {
                                                                    "description": "Assembly type identifier",
                                                                    "type": "string"
                                                                },
                                                                "ataCode": {
                                                                    "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                                    "type": "string"
                                                                },
                                                                "lifeLimitInitial": {
                                                                    "description": "Initial assembly life limit.",
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
                                                        "common": {
                                                            "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                                "flight": {
                                                    "description": "flight event, assetID defines airplane against which the event occurred",
                                                    "properties": {
                                                        "flight": {
                                                            "description": "A takeoiff and a landing",
                                                            "properties": {
                                                                "aircraft": {
                                                                    "description": "Aircraft tail or serial number (tbd)",
                                                                    "type": "string"
                                                                },
                                                                "analyticHardlanding": {
                                                                    "description": "landing considered hard by analytics",
                                                                    "type": "boolean"
                                                                },
                                                                "atd": {
                                                                    "description": "actual time departure",
                                                                    "type": "string"
                                                                },
                                                                "flightnumber": {
                                                                    "description": "A flight number",
                                                                    "type": "string"
                                                                },
                                                                "from": {
                                                                    "description": "3 letter code of originating airport",
                                                                    "type": "string"
                                                                },
                                                                "gForce": {
                                                                    "description": "force incurred on landing",
                                                                    "type": "number"
                                                                },
                                                                "hardlanding": {
                                                                    "description": "landing considered hard by pilot or aircraft sensor",
                                                                    "type": "boolean"
                                                                },
                                                                "landingType": {
                                                                    "description": "code defining landing quality??",
                                                                    "type": "string"
                                                                },
                                                                "sta": {
                                                                    "description": "standard time arrival",
                                                                    "type": "string"
                                                                },
                                                                "std": {
                                                                    "description": "standard time departure",
                                                                    "type": "string"
                                                                },
                                                                "to": {
                                                                    "description": "3 letter code of terminating airport",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "inspection": {
                                                    "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                                    "properties": {
                                                        "inspection": {
                                                            "description": "indicates that an inspection has occured for this assembly",
                                                            "properties": {
                                                                "action": {
                                                                    "description": "inspection that has been performed",
                                                                    "enum": [
                                                                        "ACHECK",
                                                                        "BCHECK",
                                                                        "HARDLANDING"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "Aircraft serial number",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "type": "object"
                                                        }
                                                    },
                                                    "type": "object"
                                                },
                                                "maintenance": {
                                                    "description": "maintenance event",
                                                    "properties": {
                                                        "maintenance": {
                                                            "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                            "properties": {
                                                                "action": {
                                                                    "enum": [
                                                                        "commission",
                                                                        "install",
                                                                        "uninstall",
                                                                        "startMaintenance",
                                                                        "endMaintenance",
                                                                        "scrap"
                                                                    ],
                                                                    "type": "string"
                                                                },
                                                                "aircraft": {
                                                                    "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                                    "type": "string"
                                                                },
                                                                "assembly": {
                                                                    "description": "This assembly's serial number",
                                                                    "type": "string"
                                                                },
                                                                "note": {
                                                                    "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                                    "type": "string"
                                                                }
                                                            },
                                                            "required": [
                                                                "assembly",
                                                                "action"
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
                            "oneOf": {
                                "aircraft": {
                                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                    "properties": {
                                        "airline": {
                                            "description": "AssetID of airline that owns this airplane",
                                            "type": "string"
                                        },
                                        "code": {
                                            "description": "Aircraft code -- e.g. WN / SWA",
                                            "type": "string"
                                        },
                                        "dateOfBuild": {
                                            "description": "Aircraft build completed / in service date",
                                            "type": "string"
                                        },
                                        "mode-s": {
                                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                                            "type": "string"
                                        },
                                        "model": {
                                            "description": "Aircraft model -- e.g. 737-5H4",
                                            "type": "string"
                                        },
                                        "operator": {
                                            "description": "AssetID of operator that flies this airplane",
                                            "type": "string"
                                        },
                                        "serialNumber": {
                                            "description": "Aircraft serial number (manufacturer assigned)",
                                            "type": "string"
                                        },
                                        "tailNumber": {
                                            "description": "Aircraft tail number (airline assigned)",
                                            "type": "string"
                                        },
                                        "variant": {
                                            "description": "Aircraft model variant -- e.g. B735",
                                            "type": "string"
                                        }
                                    },
                                    "type": "object"
                                },
                                "airline": {
                                    "description": "The writable properties for an airline",
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
                                    "type": "object"
                                },
                                "assembly": {
                                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                    "properties": {
                                        "arlsZone": {
                                            "description": "tbd",
                                            "type": "string"
                                        },
                                        "assemblyNumber": {
                                            "description": "Assembly type identifier",
                                            "type": "string"
                                        },
                                        "ataCode": {
                                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                            "type": "string"
                                        },
                                        "lifeLimitInitial": {
                                            "description": "Initial assembly life limit.",
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
                                }
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
                    "items": {
                        "description": "True to redirect on missing asset to create.",
                        "properties": {
                            "setCreateOnUpdate": {
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
                    "items": {
                        "description": "A logging level for the fabric.",
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
        "updateAssetAircraft": {
            "description": "Update the state of an aircraft asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The aircraft CRUD event",
                        "properties": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAssetAircraft function",
                    "enum": [
                        "updateAssetAircraft"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetAirline": {
            "description": "Update the state of an airline asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The airline CRUD event",
                        "properties": {
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAssetAirline function",
                    "enum": [
                        "updateAssetAirline"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateAssetAssembly": {
            "description": "Update the state of an assembly asset. The one argument is a JSON encoded event. The 'assetID' property is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "The assembly event. Note that assetID is the assembly serial number",
                        "properties": {
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateAssetAssembly function",
                    "enum": [
                        "updateAssetAssembly"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateContractConfig": {
            "description": "Update the dynamic configuration for the contract. These can include thresholds and other values.",
            "properties": {
                "args": {
                    "maxItems": 1,
                    "minItems": 1,
                    "properties": {
                        "properties": {
                            "aCheckThreshold": {
                                "description": "Cycles threshold for the aCheck inspection alert.",
                                "type": "number"
                            },
                            "bCheckThreshold": {
                                "description": "Cycles threshold for the bCheck inspection alert.",
                                "type": "number"
                            }
                        },
                        "type": "object"
                    },
                    "type": "object"
                },
                "function": {
                    "description": "updateContractConfig function",
                    "enum": [
                        "updateContractConfig"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "aircraftEvent": {
            "description": "The aircraft CRUD event",
            "properties": {
                "aircraft": {
                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                    "properties": {
                        "airline": {
                            "description": "AssetID of airline that owns this airplane",
                            "type": "string"
                        },
                        "code": {
                            "description": "Aircraft code -- e.g. WN / SWA",
                            "type": "string"
                        },
                        "dateOfBuild": {
                            "description": "Aircraft build completed / in service date",
                            "type": "string"
                        },
                        "mode-s": {
                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                            "type": "string"
                        },
                        "model": {
                            "description": "Aircraft model -- e.g. 737-5H4",
                            "type": "string"
                        },
                        "operator": {
                            "description": "AssetID of operator that flies this airplane",
                            "type": "string"
                        },
                        "serialNumber": {
                            "description": "Aircraft serial number (manufacturer assigned)",
                            "type": "string"
                        },
                        "tailNumber": {
                            "description": "Aircraft tail number (airline assigned)",
                            "type": "string"
                        },
                        "variant": {
                            "description": "Aircraft model variant -- e.g. B735",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
        "aircraftState": {
            "description": "An aircraft state",
            "properties": {
                "adjustedCycles": {
                    "description": "Cycles plus analytic adjustments for this aircraft.",
                    "type": "integer"
                },
                "age": {
                    "description": "Aircraft age, computed as today's date minus DOB",
                    "type": "string"
                },
                "aircraft": {
                    "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                    "properties": {
                        "airline": {
                            "description": "AssetID of airline that owns this airplane",
                            "type": "string"
                        },
                        "code": {
                            "description": "Aircraft code -- e.g. WN / SWA",
                            "type": "string"
                        },
                        "dateOfBuild": {
                            "description": "Aircraft build completed / in service date",
                            "type": "string"
                        },
                        "mode-s": {
                            "description": "Aircraft transponder response -- e.g.  A68E4A",
                            "type": "string"
                        },
                        "model": {
                            "description": "Aircraft model -- e.g. 737-5H4",
                            "type": "string"
                        },
                        "operator": {
                            "description": "AssetID of operator that flies this airplane",
                            "type": "string"
                        },
                        "serialNumber": {
                            "description": "Aircraft serial number (manufacturer assigned)",
                            "type": "string"
                        },
                        "tailNumber": {
                            "description": "Aircraft tail number (airline assigned)",
                            "type": "string"
                        },
                        "variant": {
                            "description": "Aircraft model variant -- e.g. B735",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "assemblies": {
                    "description": "*Internal prefixed* assetIDs of assemblies that are mounted on this airplane",
                    "items": {
                        "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                        "type": "string"
                    },
                    "type": "array"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                "cycles": {
                    "description": "Total number of cycles for this aircraft",
                    "type": "number"
                }
            },
            "type": "object"
        },
        "airlineEvent": {
            "description": "The airline CRUD event",
            "properties": {
                "airline": {
                    "description": "The writable properties for an airline",
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
                    "type": "object"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
        "airlineState": {
            "description": "An airline state",
            "properties": {
                "airline": {
                    "description": "The writable properties for an airline",
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
                    "type": "object"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
        "analyticAdjustmentEvent": {
            "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
            "properties": {
                "analyticAdjustment": {
                    "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                    "properties": {
                        "action": {
                            "enum": [
                                "adjustLifeLimit"
                            ],
                            "type": "string"
                        },
                        "amount": {
                            "type": "number"
                        },
                        "assembly": {
                            "description": "Assembly serial number",
                            "type": "string"
                        },
                        "reason": {
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "assemblyEvent": {
            "description": "The assembly event. Note that assetID is the assembly serial number",
            "properties": {
                "assembly": {
                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                    "properties": {
                        "arlsZone": {
                            "description": "tbd",
                            "type": "string"
                        },
                        "assemblyNumber": {
                            "description": "Assembly type identifier",
                            "type": "string"
                        },
                        "ataCode": {
                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                            "type": "string"
                        },
                        "lifeLimitInitial": {
                            "description": "Initial assembly life limit.",
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
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
        "assemblyState": {
            "description": "The set of writable properties that define an assembly's state.",
            "properties": {
                "aCheckCounter": {
                    "description": "The total cycles since last reset. Used as a baseline against which to compare adjusted cycles based on analytic inputs.",
                    "type": "number"
                },
                "aCheckCounterAdjusted": {
                    "description": "ACheckCounter plus analytic adjustments. Can be larger or smaller than aCheckCounter. Used for rule calculations.",
                    "type": "number"
                },
                "adjustedCycles": {
                    "description": "Cycles plus analytic adjustments for this assembly.",
                    "type": "integer"
                },
                "aircraft": {
                    "description": "The assetID of the aircraft on which this assembly is mounted. Blank if removed for maintenance.",
                    "type": "string"
                },
                "assembly": {
                    "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                    "properties": {
                        "arlsZone": {
                            "description": "tbd",
                            "type": "string"
                        },
                        "assemblyNumber": {
                            "description": "Assembly type identifier",
                            "type": "string"
                        },
                        "ataCode": {
                            "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                            "type": "string"
                        },
                        "lifeLimitInitial": {
                            "description": "Initial assembly life limit.",
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
                "bCheckCounter": {
                    "description": "The total cycles since last reset. Used as a baseline against which to compare adjusted cycles based on analytic inputs.",
                    "type": "number"
                },
                "bCheckCounterAdjusted": {
                    "description": "BCheckCounter plus analytic adjustments. Can be larger or smaller than aCheckCounter. Used for rule calculations.",
                    "type": "number"
                },
                "common": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                "cycles": {
                    "description": "Lifetime cycle count for this assembly.",
                    "type": "integer"
                },
                "maintenance": {
                    "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                    "properties": {
                        "action": {
                            "enum": [
                                "commission",
                                "install",
                                "uninstall",
                                "startMaintenance",
                                "endMaintenance",
                                "scrap"
                            ],
                            "type": "string"
                        },
                        "aircraft": {
                            "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                            "type": "string"
                        },
                        "assembly": {
                            "description": "This assembly's serial number",
                            "type": "string"
                        },
                        "note": {
                            "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                            "type": "string"
                        }
                    },
                    "required": [
                        "assembly",
                        "action"
                    ],
                    "type": "object"
                },
                "status": {
                    "enum": [
                        "new",
                        "aircraft",
                        "inventory",
                        "maintenance",
                        "scrapped"
                    ],
                    "type": "string"
                }
            },
            "required": [
                "ATAcode",
                "name"
            ],
            "type": "object"
        },
        "assetName": {
            "description": "Defines an asset that is handled by this smart contract. The complete list is in the enum property. Assets follow the 'partial state as event' pattern, which includes deep merging into the asset state when updating.",
            "enum": [
                "airline",
                "aircraft",
                "assembly"
            ],
            "type": "string"
        },
        "contractConfig": {
            "properties": {
                "aCheckThreshold": {
                    "description": "Cycles threshold for the aCheck inspection alert.",
                    "type": "number"
                },
                "bCheckThreshold": {
                    "description": "Cycles threshold for the bCheck inspection alert.",
                    "type": "number"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of properties that constitute a message to the contract.",
            "properties": {
                "iotCommon": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                "oneOf": {
                    "aircraft": {
                        "description": "The aircraft CRUD event",
                        "properties": {
                            "aircraft": {
                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                "properties": {
                                    "airline": {
                                        "description": "AssetID of airline that owns this airplane",
                                        "type": "string"
                                    },
                                    "code": {
                                        "description": "Aircraft code -- e.g. WN / SWA",
                                        "type": "string"
                                    },
                                    "dateOfBuild": {
                                        "description": "Aircraft build completed / in service date",
                                        "type": "string"
                                    },
                                    "mode-s": {
                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                        "type": "string"
                                    },
                                    "model": {
                                        "description": "Aircraft model -- e.g. 737-5H4",
                                        "type": "string"
                                    },
                                    "operator": {
                                        "description": "AssetID of operator that flies this airplane",
                                        "type": "string"
                                    },
                                    "serialNumber": {
                                        "description": "Aircraft serial number (manufacturer assigned)",
                                        "type": "string"
                                    },
                                    "tailNumber": {
                                        "description": "Aircraft tail number (airline assigned)",
                                        "type": "string"
                                    },
                                    "variant": {
                                        "description": "Aircraft model variant -- e.g. B735",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "airline": {
                        "description": "The airline CRUD event",
                        "properties": {
                            "airline": {
                                "description": "The writable properties for an airline",
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
                                "type": "object"
                            },
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "analyticAdjustment": {
                        "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                        "properties": {
                            "analyticAdjustment": {
                                "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                "properties": {
                                    "action": {
                                        "enum": [
                                            "adjustLifeLimit"
                                        ],
                                        "type": "string"
                                    },
                                    "amount": {
                                        "type": "number"
                                    },
                                    "assembly": {
                                        "description": "Assembly serial number",
                                        "type": "string"
                                    },
                                    "reason": {
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "assembly": {
                        "description": "The assembly event. Note that assetID is the assembly serial number",
                        "properties": {
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "arlsZone": {
                                        "description": "tbd",
                                        "type": "string"
                                    },
                                    "assemblyNumber": {
                                        "description": "Assembly type identifier",
                                        "type": "string"
                                    },
                                    "ataCode": {
                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                        "type": "string"
                                    },
                                    "lifeLimitInitial": {
                                        "description": "Initial assembly life limit.",
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
                            "common": {
                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                    "flight": {
                        "description": "flight event, assetID defines airplane against which the event occurred",
                        "properties": {
                            "flight": {
                                "description": "A takeoiff and a landing",
                                "properties": {
                                    "aircraft": {
                                        "description": "Aircraft tail or serial number (tbd)",
                                        "type": "string"
                                    },
                                    "analyticHardlanding": {
                                        "description": "landing considered hard by analytics",
                                        "type": "boolean"
                                    },
                                    "atd": {
                                        "description": "actual time departure",
                                        "type": "string"
                                    },
                                    "flightnumber": {
                                        "description": "A flight number",
                                        "type": "string"
                                    },
                                    "from": {
                                        "description": "3 letter code of originating airport",
                                        "type": "string"
                                    },
                                    "gForce": {
                                        "description": "force incurred on landing",
                                        "type": "number"
                                    },
                                    "hardlanding": {
                                        "description": "landing considered hard by pilot or aircraft sensor",
                                        "type": "boolean"
                                    },
                                    "landingType": {
                                        "description": "code defining landing quality??",
                                        "type": "string"
                                    },
                                    "sta": {
                                        "description": "standard time arrival",
                                        "type": "string"
                                    },
                                    "std": {
                                        "description": "standard time departure",
                                        "type": "string"
                                    },
                                    "to": {
                                        "description": "3 letter code of terminating airport",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "inspection": {
                        "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                        "properties": {
                            "inspection": {
                                "description": "indicates that an inspection has occured for this assembly",
                                "properties": {
                                    "action": {
                                        "description": "inspection that has been performed",
                                        "enum": [
                                            "ACHECK",
                                            "BCHECK",
                                            "HARDLANDING"
                                        ],
                                        "type": "string"
                                    },
                                    "aircraft": {
                                        "description": "Aircraft serial number",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "type": "object"
                    },
                    "maintenance": {
                        "description": "maintenance event",
                        "properties": {
                            "maintenance": {
                                "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                "properties": {
                                    "action": {
                                        "enum": [
                                            "commission",
                                            "install",
                                            "uninstall",
                                            "startMaintenance",
                                            "endMaintenance",
                                            "scrap"
                                        ],
                                        "type": "string"
                                    },
                                    "aircraft": {
                                        "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                        "type": "string"
                                    },
                                    "assembly": {
                                        "description": "This assembly's serial number",
                                        "type": "string"
                                    },
                                    "note": {
                                        "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "assembly",
                                    "action"
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
        "eventName": {
            "description": "Defines an event that is related to an asset. The complete list is in the enum property.",
            "enum": [
                "flight",
                "inspection",
                "analyticAdjustment"
            ],
            "type": "string"
        },
        "flightEvent": {
            "description": "flight event, assetID defines airplane against which the event occurred",
            "properties": {
                "flight": {
                    "description": "A takeoiff and a landing",
                    "properties": {
                        "aircraft": {
                            "description": "Aircraft tail or serial number (tbd)",
                            "type": "string"
                        },
                        "analyticHardlanding": {
                            "description": "landing considered hard by analytics",
                            "type": "boolean"
                        },
                        "atd": {
                            "description": "actual time departure",
                            "type": "string"
                        },
                        "flightnumber": {
                            "description": "A flight number",
                            "type": "string"
                        },
                        "from": {
                            "description": "3 letter code of originating airport",
                            "type": "string"
                        },
                        "gForce": {
                            "description": "force incurred on landing",
                            "type": "number"
                        },
                        "hardlanding": {
                            "description": "landing considered hard by pilot or aircraft sensor",
                            "type": "boolean"
                        },
                        "landingType": {
                            "description": "code defining landing quality??",
                            "type": "string"
                        },
                        "sta": {
                            "description": "standard time arrival",
                            "type": "string"
                        },
                        "std": {
                            "description": "standard time departure",
                            "type": "string"
                        },
                        "to": {
                            "description": "3 letter code of terminating airport",
                            "type": "string"
                        }
                    },
                    "type": "object"
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
        "inspectionEvent": {
            "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
            "properties": {
                "inspection": {
                    "description": "indicates that an inspection has occured for this assembly",
                    "properties": {
                        "action": {
                            "description": "inspection that has been performed",
                            "enum": [
                                "ACHECK",
                                "BCHECK",
                                "HARDLANDING"
                            ],
                            "type": "string"
                        },
                        "aircraft": {
                            "description": "Aircraft serial number",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "maintenanceEvent": {
            "description": "maintenance event",
            "properties": {
                "maintenance": {
                    "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                    "properties": {
                        "action": {
                            "enum": [
                                "commission",
                                "install",
                                "uninstall",
                                "startMaintenance",
                                "endMaintenance",
                                "scrap"
                            ],
                            "type": "string"
                        },
                        "aircraft": {
                            "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                            "type": "string"
                        },
                        "assembly": {
                            "description": "This assembly's serial number",
                            "type": "string"
                        },
                        "note": {
                            "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                            "type": "string"
                        }
                    },
                    "required": [
                        "assembly",
                        "action"
                    ],
                    "type": "object"
                }
            },
            "type": "object"
        },
        "state": {
            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
            "properties": {
                "alerts": {
                    "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                    "properties": {
                        "active": {
                            "items": {
                                "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                "enum": [
                                    "ACHECK",
                                    "BCHECK",
                                    "HARDLANDING"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "cleared": {
                            "items": {
                                "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                "enum": [
                                    "ACHECK",
                                    "BCHECK",
                                    "HARDLANDING"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "raised": {
                            "items": {
                                "description": "These alerts signal a required inspection. Only the inspection event with an action of the same name can clear these alewrts.",
                                "enum": [
                                    "ACHECK",
                                    "BCHECK",
                                    "HARDLANDING"
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
                "iotCommon": {
                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                "lastEvent": {
                    "description": "function and string parameter that created this state object",
                    "properties": {
                        "arg": {
                            "description": "A set of properties that constitute a message to the contract.",
                            "properties": {
                                "iotCommon": {
                                    "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                "oneOf": {
                                    "aircraft": {
                                        "description": "The aircraft CRUD event",
                                        "properties": {
                                            "aircraft": {
                                                "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                                                "properties": {
                                                    "airline": {
                                                        "description": "AssetID of airline that owns this airplane",
                                                        "type": "string"
                                                    },
                                                    "code": {
                                                        "description": "Aircraft code -- e.g. WN / SWA",
                                                        "type": "string"
                                                    },
                                                    "dateOfBuild": {
                                                        "description": "Aircraft build completed / in service date",
                                                        "type": "string"
                                                    },
                                                    "mode-s": {
                                                        "description": "Aircraft transponder response -- e.g.  A68E4A",
                                                        "type": "string"
                                                    },
                                                    "model": {
                                                        "description": "Aircraft model -- e.g. 737-5H4",
                                                        "type": "string"
                                                    },
                                                    "operator": {
                                                        "description": "AssetID of operator that flies this airplane",
                                                        "type": "string"
                                                    },
                                                    "serialNumber": {
                                                        "description": "Aircraft serial number (manufacturer assigned)",
                                                        "type": "string"
                                                    },
                                                    "tailNumber": {
                                                        "description": "Aircraft tail number (airline assigned)",
                                                        "type": "string"
                                                    },
                                                    "variant": {
                                                        "description": "Aircraft model variant -- e.g. B735",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            },
                                            "common": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                    "airline": {
                                        "description": "The airline CRUD event",
                                        "properties": {
                                            "airline": {
                                                "description": "The writable properties for an airline",
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
                                                "type": "object"
                                            },
                                            "common": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                    "analyticAdjustment": {
                                        "description": "analytic adjustment event, assetid defines the assembly receiving the adjustment",
                                        "properties": {
                                            "analyticAdjustment": {
                                                "description": "An adjustment based on analytical analysis to the assembly's cycle counters, which translates to changes to life limit *used*. Positive number indicates that the assembly has used more of its life, negative number indicates that the assembly has been granted a bit more life based on conditions such as weather, landing gForces, runway roughness and so on.",
                                                "properties": {
                                                    "action": {
                                                        "enum": [
                                                            "adjustLifeLimit"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "amount": {
                                                        "type": "number"
                                                    },
                                                    "assembly": {
                                                        "description": "Assembly serial number",
                                                        "type": "string"
                                                    },
                                                    "reason": {
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "assembly": {
                                        "description": "The assembly event. Note that assetID is the assembly serial number",
                                        "properties": {
                                            "assembly": {
                                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                                "properties": {
                                                    "arlsZone": {
                                                        "description": "tbd",
                                                        "type": "string"
                                                    },
                                                    "assemblyNumber": {
                                                        "description": "Assembly type identifier",
                                                        "type": "string"
                                                    },
                                                    "ataCode": {
                                                        "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                                        "type": "string"
                                                    },
                                                    "lifeLimitInitial": {
                                                        "description": "Initial assembly life limit.",
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
                                            "common": {
                                                "description": "The set of common properties for any event to a contract that adheres to the IoT contract pattern 'partial state as event' for assets and that may have pure events that are *about* these assets.",
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
                                    "flight": {
                                        "description": "flight event, assetID defines airplane against which the event occurred",
                                        "properties": {
                                            "flight": {
                                                "description": "A takeoiff and a landing",
                                                "properties": {
                                                    "aircraft": {
                                                        "description": "Aircraft tail or serial number (tbd)",
                                                        "type": "string"
                                                    },
                                                    "analyticHardlanding": {
                                                        "description": "landing considered hard by analytics",
                                                        "type": "boolean"
                                                    },
                                                    "atd": {
                                                        "description": "actual time departure",
                                                        "type": "string"
                                                    },
                                                    "flightnumber": {
                                                        "description": "A flight number",
                                                        "type": "string"
                                                    },
                                                    "from": {
                                                        "description": "3 letter code of originating airport",
                                                        "type": "string"
                                                    },
                                                    "gForce": {
                                                        "description": "force incurred on landing",
                                                        "type": "number"
                                                    },
                                                    "hardlanding": {
                                                        "description": "landing considered hard by pilot or aircraft sensor",
                                                        "type": "boolean"
                                                    },
                                                    "landingType": {
                                                        "description": "code defining landing quality??",
                                                        "type": "string"
                                                    },
                                                    "sta": {
                                                        "description": "standard time arrival",
                                                        "type": "string"
                                                    },
                                                    "std": {
                                                        "description": "standard time departure",
                                                        "type": "string"
                                                    },
                                                    "to": {
                                                        "description": "3 letter code of terminating airport",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "inspection": {
                                        "description": "An inspection has been performed against a specific assembly. Will clear one or more alerts and reset their counters.",
                                        "properties": {
                                            "inspection": {
                                                "description": "indicates that an inspection has occured for this assembly",
                                                "properties": {
                                                    "action": {
                                                        "description": "inspection that has been performed",
                                                        "enum": [
                                                            "ACHECK",
                                                            "BCHECK",
                                                            "HARDLANDING"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "aircraft": {
                                                        "description": "Aircraft serial number",
                                                        "type": "string"
                                                    }
                                                },
                                                "type": "object"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "maintenance": {
                                        "description": "maintenance event",
                                        "properties": {
                                            "maintenance": {
                                                "description": "Maintenance consists of installation of an assembly onto an aircraft or uninstallation of same. When an assembly is not installed on an aircraft, it is said to be in inventory or in maintenance. Thus, there is a status on assemblies showing that.",
                                                "properties": {
                                                    "action": {
                                                        "enum": [
                                                            "commission",
                                                            "install",
                                                            "uninstall",
                                                            "startMaintenance",
                                                            "endMaintenance",
                                                            "scrap"
                                                        ],
                                                        "type": "string"
                                                    },
                                                    "aircraft": {
                                                        "description": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
                                                        "type": "string"
                                                    },
                                                    "assembly": {
                                                        "description": "This assembly's serial number",
                                                        "type": "string"
                                                    },
                                                    "note": {
                                                        "description": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event.",
                                                        "type": "string"
                                                    }
                                                },
                                                "required": [
                                                    "assembly",
                                                    "action"
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
                "oneOf": {
                    "aircraft": {
                        "description": "Writable properties for an aircraft. Note that assetID is the aircraft serial number.",
                        "properties": {
                            "airline": {
                                "description": "AssetID of airline that owns this airplane",
                                "type": "string"
                            },
                            "code": {
                                "description": "Aircraft code -- e.g. WN / SWA",
                                "type": "string"
                            },
                            "dateOfBuild": {
                                "description": "Aircraft build completed / in service date",
                                "type": "string"
                            },
                            "mode-s": {
                                "description": "Aircraft transponder response -- e.g.  A68E4A",
                                "type": "string"
                            },
                            "model": {
                                "description": "Aircraft model -- e.g. 737-5H4",
                                "type": "string"
                            },
                            "operator": {
                                "description": "AssetID of operator that flies this airplane",
                                "type": "string"
                            },
                            "serialNumber": {
                                "description": "Aircraft serial number (manufacturer assigned)",
                                "type": "string"
                            },
                            "tailNumber": {
                                "description": "Aircraft tail number (airline assigned)",
                                "type": "string"
                            },
                            "variant": {
                                "description": "Aircraft model variant -- e.g. B735",
                                "type": "string"
                            }
                        },
                        "type": "object"
                    },
                    "airline": {
                        "description": "The writable properties for an airline",
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
                        "type": "object"
                    },
                    "assembly": {
                        "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                        "properties": {
                            "arlsZone": {
                                "description": "tbd",
                                "type": "string"
                            },
                            "assemblyNumber": {
                                "description": "Assembly type identifier",
                                "type": "string"
                            },
                            "ataCode": {
                                "description": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                                "type": "string"
                            },
                            "lifeLimitInitial": {
                                "description": "Initial assembly life limit.",
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
                    }
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
        "stateFilter": {
            "description": "A filter that can be used to restrict a read operation to a subset of objects that meet a criteria.",
            "properties": {
                "filter": {
                    "properties": {
                        "entries": {
                            "description": "Array of qualified propery name and value pairs.",
                            "items": {
                                "properties": {
                                    "qprop": {
                                        "description": "A qualified property as dot separated levels terminated by a leaf node. An example would be 'common.assetID'.",
                                        "type": "string"
                                    },
                                    "value": {
                                        "description": "The value to be compared.",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "type": "array"
                        },
                        "matchhow": {
                            "enum": [
                                "matchall",
                                "matchany",
                                "matchnone"
                            ],
                            "type": "string"
                        }
                    },
                    "required": [
                        "entries",
                        "matchhow"
                    ],
                    "type": "object"
                }
            },
            "required": [
                "filter"
            ],
            "type": "object"
        }
    }
}`