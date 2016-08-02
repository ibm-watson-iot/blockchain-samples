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
            "timestamp": "2016-07-27T13:33:19.900240086Z"
        }
    },
    "aircraftState": {
        "aCheckCounter": 123.456,
        "age": "Aircraft age, computed as today's date minus DOB",
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
        "assemblies": [
            "NO TYPE PROPERTY"
        ],
        "bCheckCounter": 123.456,
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
            "timestamp": "2016-07-27T13:33:19.900252144Z"
        },
        "cycles": 123.456
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
            "timestamp": "2016-07-27T13:33:19.900178375Z"
        }
    },
    "airlineState": {
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
            "timestamp": "2016-07-27T13:33:19.90022825Z"
        }
    },
    "assemblyEvent": {
        "assembly": {
            "airplane": "The assetID of the airplane on which this assembly is mounted. Blank if removed for maintenance.",
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
            "timestamp": "2016-07-27T13:33:19.9002655Z"
        }
    },
    "assemblyState": {
        "aCheckCounter": 123.456,
        "assembly": {
            "airplane": "The assetID of the airplane on which this assembly is mounted. Blank if removed for maintenance.",
            "arlsZone": "tbd",
            "assemblyNumber": "Assembly type identifier",
            "ataCode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
            "lifeLimitInitial": 789,
            "name": "The assembly name."
        },
        "bCheckCounter": 123.456,
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
            "timestamp": "2016-07-27T13:33:19.900275617Z"
        },
        "cycles": 789,
        "lifeLimitAdjusted": 789
    },
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "TRADELANE",
        "version": "The version number of the current contract"
    },
    "event": {
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
            "timestamp": "2016-07-27T13:33:19.900287362Z"
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
            "assembly": {
                "description": "The set of writable properties that define an assembly. Note that assetID is the assembly serial number",
                "properties": {
                    "airplane": {
                        "description": "The assetID of the airplane on which this assembly is mounted. Blank if removed for maintenance.",
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
        }
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
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
            "timestamp": "2016-07-27T13:33:19.900364639Z"
        },
        "lastEvent": {
            "arg": {
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
                "airline": {
                    "code": "The airline 3 letter code.",
                    "name": "The name of the airline."
                },
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
                "assembly": {
                    "airplane": "The assetID of the airplane on which this assembly is mounted. Blank if removed for maintenance.",
                    "arlsZone": "tbd",
                    "assemblyNumber": "Assembly type identifier",
                    "ataCode": "The ATA code defining the assembly type, e.g. 32=landing gear, 32-50=steering.",
                    "lifeLimitInitial": 789,
                    "name": "The assembly name."
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
                },
                "inspection": {
                    "aircraft": "Aircraft tail or serial number (tbd)",
                    "enum": "UNKNOWN ARRAY OBJECT",
                    "type": "ACHECK or BCHECK inspection has been performed and will clear the alert of the same name"
                },
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
                    "timestamp": "2016-07-27T13:33:19.900323187Z"
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
                    "airplane": {
                        "description": "The assetID of the airplane on which this assembly is mounted. Blank if removed for maintenance.",
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