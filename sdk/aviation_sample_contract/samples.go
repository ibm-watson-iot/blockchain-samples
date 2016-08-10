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
            "tailNumber": "Designated asset ID. Aircraft tail number (airline assigned)",
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
            "timestamp": "2016-08-09T16:43:36.617660285Z"
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
            "timestamp": "2016-08-09T16:43:36.617601589Z"
        }
    },
    "analyticAdjustmentEvent": {
        "analyticAdjustment": {
            "assembly": "Assembly serial number",
            "increaselifelimit": {
                "reason": "carpe noctem",
                "reduction": 123.456
            },
            "reducelifelimit": {
                "reason": "carpe noctem",
                "reduction": 123.456
            }
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
            "timestamp": "2016-08-09T16:43:36.617727645Z"
        }
    },
    "assemblyEvent": {
        "assembly": {
            "aircraft": "The assetID of the aircraft on which this assembly is mounted. Blank if removed for maintenance.",
            "arlsZone": "tbd",
            "assemblyNumber": "Assembly type identifier",
            "ataCode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
            "lifeLimitInitial": 789,
            "name": "The assembly name."
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
            "timestamp": "2016-08-09T16:43:36.617674539Z"
        }
    },
    "flightEvent": {
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
            "timestamp": "2016-08-09T16:43:36.617701092Z"
        },
        "flight": {
            "aircraft": "Aircraft tail or serial number (tbd)",
            "atd": "actual time departure",
            "flightnumber": "A flight number",
            "from": "3 letter code of originating airport",
            "gForce": 123.456,
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
            "timestamp": "2016-08-09T16:43:36.617715347Z"
        },
        "inspection": {
            "aircraft": "Aircraft tail or serial number (tbd)",
            "enum": "UNKNOWN ARRAY OBJECT",
            "type": "ACHECK or BCHECK inspection has been performed and will clear the alert of the same name"
        }
    },
    "state": {
        "alerts": {
            "active": [
                "NONE",
                "ACHECK",
                "BCHECK"
            ],
            "cleared": [
                "NONE",
                "ACHECK",
                "BCHECK"
            ],
            "raised": [
                "NONE",
                "ACHECK",
                "BCHECK"
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
            "timestamp": "2016-08-09T16:43:36.617739664Z"
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
                    "timestamp": "2016-08-09T16:43:36.617749726Z"
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
                                    "tailNumber": {
                                        "description": "Designated asset ID. Aircraft tail number (airline assigned)",
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
                                "description": "an adjustment to the assembly's life limit based on analytical analysis; can be positive or negative with the former indicating that the part is granted additional life and the latter indicating that the part has its lime limit reduced by rough runway landings etc",
                                "properties": {
                                    "assembly": {
                                        "description": "Assembly serial number",
                                        "type": "string"
                                    },
                                    "increaselifelimit": {
                                        "description": "increase of life limit",
                                        "properties": {
                                            "reason": {
                                                "type": "string"
                                            },
                                            "reduction": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
                                    },
                                    "reducelifelimit": {
                                        "description": "reduction of life limit",
                                        "properties": {
                                            "reason": {
                                                "type": "string"
                                            },
                                            "reduction": {
                                                "type": "number"
                                            }
                                        },
                                        "type": "object"
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
                    "assembly": {
                        "description": "The assembly event. Note that assetID is the assembly serial number",
                        "properties": {
                            "assembly": {
                                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                                "properties": {
                                    "aircraft": {
                                        "description": "The assetID of the aircraft on which this assembly is mounted. Blank if removed for maintenance.",
                                        "type": "string"
                                    },
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
                            "flight": {
                                "description": "A takeoiff and a landing",
                                "properties": {
                                    "aircraft": {
                                        "description": "Aircraft tail or serial number (tbd)",
                                        "type": "string"
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
                        "description": "inspection event, assetid defines the airplane against which the inspection occurred",
                        "properties": {
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
                            "inspection": {
                                "description": "indicates that an inspection has occured for this aircraft",
                                "properties": {
                                    "aircraft": {
                                        "description": "Aircraft tail or serial number (tbd)",
                                        "type": "string"
                                    },
                                    "enum": [
                                        "select an inspection type",
                                        "ACHECK",
                                        "BCHECK"
                                    ],
                                    "type": {
                                        "description": "ACHECK or BCHECK inspection has been performed and will clear the alert of the same name",
                                        "type": "string"
                                    }
                                },
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
                    "tailNumber": {
                        "description": "Designated asset ID. Aircraft tail number (airline assigned)",
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
                    "aircraft": {
                        "description": "The assetID of the aircraft on which this assembly is mounted. Blank if removed for maintenance.",
                        "type": "string"
                    },
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
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain."
    }
}`