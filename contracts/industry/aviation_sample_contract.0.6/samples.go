package main

var samples = `
{
    "aircraftEvent": {
        "aircraft": {
            "airline": "AssetID of airline that owns this airplane",
            "code": "Aircraft code -- e.g. WN / SWA",
            "dateOfBuild": "Aircraft build completed / in service date",
            "mode-s": "Aircraft transponder response -- e.g.  A68E4A",
            "model": "Aircraft model -- e.g. 737-5H4",
            "operator": "AssetID of operator that flies this airplane",
            "serialNumber": "Aircraft serial number (manufacturer assigned)",
            "tailNumber": "Aircraft tail number (airline assigned)",
            "variant": "Aircraft model variant -- e.g. B735"
        },
        "common": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "references": [
                "carpe noctem"
            ],
            "timestamp": "2016-10-05T05:35:44.57831513Z"
        }
    },
    "airlineEvent": {
        "airline": {
            "code": "The airline 3 letter code.",
            "name": "The name of the airline."
        },
        "common": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "references": [
                "carpe noctem"
            ],
            "timestamp": "2016-10-05T05:35:44.578237395Z"
        }
    },
    "analyticAdjustmentEvent": {
        "analyticAdjustment": {
            "action": "adjustLifeLimit",
            "amount": 123.456,
            "assembly": "Assembly serial number",
            "reason": "carpe noctem"
        }
    },
    "assemblyEvent": {
        "assembly": {
            "arlsZone": "tbd",
            "ataCode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
            "lifeLimitInitial": 789,
            "name": "The assembly name.",
            "serialNumber": "Assembly identifier assigned by manufacturer"
        },
        "common": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "references": [
                "carpe noctem"
            ],
            "timestamp": "2016-10-05T05:35:44.578334201Z"
        }
    },
    "flightEvent": {
        "flight": {
            "aircraft": "Aircraft tail or serial number (tbd)",
            "analyticHardlanding": true,
            "atd": "actual time departure",
            "flightnumber": "A flight number",
            "from": "3 letter code of originating airport",
            "gForce": 123.456,
            "hardlanding": true,
            "landingType": "code defining landing quality??",
            "sta": "standard time arrival",
            "std": "standard time departure",
            "to": "3 letter code of terminating airport"
        }
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "inspectionEvent": {
        "inspection": {
            "action": "BCHECK",
            "assembly": "assembly serial number"
        }
    },
    "maintenanceEvent": {
        "maintenance": {
            "action": "install",
            "aircraft": "The serial number of the aircraft to / from which the assembly has been installed / uninstalled.",
            "assembly": "This assembly's serial number",
            "note": "Maintenance note for this action. Overwritten whenever a new note property is inserted into the maintenance sub-event."
        }
    },
    "state": {
        "alerts": {
            "active": [
                "ACHECK",
                "BCHECK",
                "HARDLANDING"
            ],
            "cleared": [
                "ACHECK",
                "BCHECK",
                "HARDLANDING"
            ],
            "raised": [
                "ACHECK",
                "BCHECK",
                "HARDLANDING"
            ]
        },
        "compliant": true,
        "iotCommon": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": [
                {}
            ],
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "references": [
                "carpe noctem"
            ],
            "timestamp": "2016-10-05T05:35:44.57837156Z"
        },
        "lastEvent": {
            "arg": {
                "iotCommon": {
                    "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "extension": [
                        {}
                    ],
                    "location": {
                        "latitude": 123.456,
                        "longitude": 123.456
                    },
                    "references": [
                        "carpe noctem"
                    ],
                    "timestamp": "2016-10-05T05:35:44.578383473Z"
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
                                    },
                                    "serialNumber": {
                                        "description": "Assembly identifier assigned by manufacturer",
                                        "type": "string"
                                    }
                                },
                                "required": [
                                    "serialNumber",
                                    "ataCode",
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
                                    "assembly": {
                                        "description": "assembly serial number",
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
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
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
                    },
                    "serialNumber": {
                        "description": "Assembly identifier assigned by manufacturer",
                        "type": "string"
                    }
                },
                "required": [
                    "serialNumber",
                    "ataCode",
                    "name"
                ],
                "type": "object"
            }
        },
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    },
    "stateFilter": {
        "entries": [
            {
                "qprop": "A qualified property as dot separated levels terminated by a leaf node. An example would be 'common.assetID'.",
                "value": "The value to be compared."
            }
        ],
        "matchmode": "matchany"
    }
}`