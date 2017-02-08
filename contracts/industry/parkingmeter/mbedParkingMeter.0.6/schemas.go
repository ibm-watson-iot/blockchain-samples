package main

var schemas = `
{
    "API": {
        "createDevice": {
            "description": "Create one or more  parking meter device. One argument, a JSON encoded event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Parking meter device registration / update.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "minimumusagecost": {
                                "description": "minimum cost for using the meter",
                                "type": "number"
                            },
                            "minimumusagetime": {
                                "description": "minimum duration of use",
                                "type": "integer"
                            },
                            "overtimeusagecost": {
                                "description": "overtime cost.",
                                "type": "number"
                            },
                            "overtimeusagetime": {
                                "description": "Overtime duration.",
                                "type": "integer"
                            },
                            "latitude": {
                                "description": "Location:latitude",
                                "type": "number"
                            },
                            "longitude": {
                                "description": "Location:longitude",
                                "type": "number"
                            },
                            "address": {
                                "description": "Abbreviated street address",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                        },
                        "required": [
                            "deviceid",
                            "minimumusagecost",
                            "minimumusagetime",
                            "overtimeusagecost",
                            "overtimeusagetime",
                            "latitude",
                            "longitude",
                            "address",
                            "available"
                        ],
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createDevice function",
                    "enum": [
                        "createDevice"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "createUsage": {
            "description": "Create  parking meter usage record. One argument, a JSON encoded event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Parking meter device registration / update.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "starttime": {
                                "description": "start time of usage",
                                "type": "string"
                            },
                            "endtime": {
                                "description": "usage end time. Computed from start time and duration",
                                "type": "string"
                            },
                            "duration": {
                                "description": "Usage duration. Difference between start and end times. Duration is passed in and end time computed",
                                "type": "number"
                            },
                            "usagecost": {
                                "description": "Usage cost. Based on duration and rates defined for device.",
                                "type": "number"
                            },
                            "actualendtime": {
                                "description": "actual end time. Provision for overtime scenario",
                                "type": "string"
                            },
                            "overtimecost": {
                                "description": "Cost incurred for overtime use. Provision for overtime scenario",
                                "type": "number"
                            },
                             "totalcost": {
                                "description": "Total Cost incurred including overtime use. Provision for overtime scenario",
                                "type": "number"
                            }
                        },
                        "required": [
                            "deviceid"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createUsage function",
                    "enum": [
                        "createUsage"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        },
        "updateDeviceAsAvailable": {
            "description": "Flag device as available.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Parking meter device registration / update.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                        },
                        "required": [
                            "deviceid"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateDeviceAsAvailable function",
                    "enum": [
                        "updateDeviceAsAvailable"
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
                                "default": "SIMPLE",
                                "description": "The nickname of the current contract",
                                "type": "string"
                            },
                            "version": {
                                "description": "The ID of a managed device. The resource focal point for a smart contract.",
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
        "readDevice": {
            "description": "Returns the state an device. Argument is a JSON encoded string. Device id is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an Device Id for use as an argument to read or delete.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a managed device. The resource focal point for a smart contract.",
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
                    "description": "readDevice function",
                    "enum": [
                        "readDevice"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete device state.",
                    "properties": {
                         "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "minimumusagecost": {
                                "description": "minimum cost for using the meter",
                                "type": "number"
                            },
                            "minimumusagetime": {
                                "description": "minimum duration of use",
                                "type": "integer"
                            },
                            "overtimeusagecost": {
                                "description": "overtime cost.",
                                "type": "number"
                            },
                            "overtimeusagetime": {
                                "description": "Overtime duration.",
                                "type": "integer"
                            },
                            "latitude": {
                                "description": "Location:latitude",
                                "type": "number"
                            },
                            "longitude": {
                                "description": "Location:longitude",
                                "type": "number"
                            },
                            "address": {
                                "description": "Abbreviated street address",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readAssetSchemas": {
            "description": "Returns a string generated from the schema containing APIs and Objects as specified in generate.json in the scripts folder.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readAssetSchemas function",
                    "enum": [
                        "readAssetSchemas"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "JSON encoded object containing selected schemas",
                    "type": "string"
                }
            },
            "type": "object"
        },
         "readUsage": {
            "description": "Returns the state an device. Argument is a JSON encoded string. Device Id is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an Device Id for use as an argument to read or delete.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a managed device. The resource focal point for a smart contract.",
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
                    "description": "readUsage function",
                    "enum": [
                        "readUsage"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete device state.",
                    "properties": {
                         "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "minimumusagecost": {
                                "description": "minimum cost for using the meter",
                                "type": "number"
                            },
                            "minimumusagetime": {
                                "description": "minimum duration of use",
                                "type": "integer"
                            },
                            "overtimeusagecost": {
                                "description": "overtime cost.",
                                "type": "number"
                            },
                            "overtimeusagetime": {
                                "description": "Overtime duration.",
                                "type": "integer"
                            },
                            "latitude": {
                                "description": "Location:latitude",
                                "type": "number"
                            },
                            "longitude": {
                                "description": "Location:longitude",
                                "type": "number"
                            },
                            "address": {
                                "description": "Abbreviated street address",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
         "readUsageHistory": {
            "description": "Returns the history of an device. Argument is a JSON encoded string. deviceid is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an Device Id for use as an argument to read or delete.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a managed device. The resource focal point for a smart contract.",
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
                    "description": "readUsageHistory function",
                    "enum": [
                        "readUsageHistory"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete device state.",
                    "properties": {
                         "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "minimumusagecost": {
                                "description": "minimum cost for using the meter",
                                "type": "number"
                            },
                            "minimumusagetime": {
                                "description": "minimum duration of use",
                                "type": "integer"
                            },
                            "overtimeusagecost": {
                                "description": "overtime cost.",
                                "type": "number"
                            },
                            "overtimeusagetime": {
                                "description": "Overtime duration.",
                                "type": "integer"
                            },
                            "latitude": {
                                "description": "Location:latitude",
                                "type": "number"
                            },
                            "longitude": {
                                "description": "Location:longitude",
                                "type": "number"
                            },
                            "address": {
                                "description": "Abbreviated street address",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
         "readDeviceList": {
            "description": "Returns the history of an device. Argument is a JSON encoded string. deviceid is the only accepted property.",
            "properties": {
                "args": {
                    "description": "accepts no arguments",
                    "items": {},
                    "maxItems": 0,
                    "minItems": 0,
                    "type": "array"
                },
                "function": {
                    "description": "readDeviceList function",
                    "enum": [
                        "readDeviceList"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete device state.",
                    "properties": {
                         "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            },
                            "minimumusagecost": {
                                "description": "minimum cost for using the meter",
                                "type": "number"
                            },
                            "minimumusagetime": {
                                "description": "minimum duration of use",
                                "type": "integer"
                            },
                            "overtimeusagecost": {
                                "description": "overtime cost.",
                                "type": "number"
                            },
                            "overtimeusagetime": {
                                "description": "Overtime duration.",
                                "type": "integer"
                            },
                            "latitude": {
                                "description": "Location:latitude",
                                "type": "number"
                            },
                            "longitude": {
                                "description": "Location:longitude",
                                "type": "number"
                            },
                            "address": {
                                "description": "Abbreviated street address",
                                "type": "string"
                            },
                            "available": {
                                "description": "Light in candela.",
                                "type": "boolean"
                            }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "deleteDevice": {
            "description": "Delete a parking meter device. One argument, a JSON encoded event.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Parking meter device registration / update.",
                        "properties": {
                            "deviceid": {
                                "description": "The ID of a meter.",
                                "type": "string"
                            }
                        },
                        "required": [
                            "deviceid"
                        ],
                        "type": "object"
                    },
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "deleteDevice function",
                    "enum": [
                        "deleteDevice"
                    ],
                    "type": "string"
                },
                "method": "invoke"
            },
            "type": "object"
        }
        
    },
    "objectModelSchemas": {
        "deviceidKey": {
            "description": "An object containing only an Device Id  for use as an argument to read or delete.",
            "properties": {
                "deviceid": {
                    "description": "The ID of a managed device. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an device's state. Device Id  is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "deviceid": {
                    "description": "The ID of a meter.",
                    "type": "string"
                },
                "minimumusagecost": {
                    "description": "minimum cost for using the meter",
                    "type": "number"
                },
                "minimumusagetime": {
                    "description": "minimum duration of use",
                    "type": "integer"
                },
                "overtimeusagecost": {
                    "description": "overtime cost.",
                    "type": "number"
                },
                "overtimeusagetime": {
                    "description": "Overtime duration.",
                    "type": "integer"
                },
                "latitude": {
                    "description": "Location:latitude",
                    "type": "number"
                },
                "longitude": {
                    "description": "Location:longitude",
                    "type": "number"
                },
                "address": {
                    "description": "Abbreviated street address",
                    "type": "string"
                },
                "available": {
                    "description": "Light in candela.",
                    "type": "boolean"
                }
            },
            "required": [
                "deviceid"
            ],
            "type": "object"
        },
        "initEvent": {
            "description": "event sent to init on deployment",
            "properties": {
                "nickname": {
                    "default": "SIMPLE",
                    "description": "The nickname of the current contract",
                    "type": "string"
                },
                "version": {
                    "description": "The version of the contract.",
                    "type": "string"
                }
            },
            "required": [
                "version"
            ],
            "type": "object"
        },
        "state": {
            "description": "A set of fields that constitute the complete device state.",
            "properties": {
                "deviceid": {
                    "description": "The ID of a meter.",
                    "type": "string"
                },
                "minimumusagecost": {
                    "description": "minimum cost for using the meter",
                    "type": "number"
                },
                "minimumusagetime": {
                    "description": "minimum duration of use",
                    "type": "integer"
                },
                "overtimeusagecost": {
                    "description": "overtime cost.",
                    "type": "number"
                },
                "overtimeusagetime": {
                    "description": "Overtime duration.",
                    "type": "integer"
                },
                "latitude": {
                    "description": "Location:latitude",
                    "type": "number"
                },
                "longitude": {
                    "description": "Location:longitude",
                    "type": "number"
                },
                "address": {
                    "description": "Abbreviated street address",
                    "type": "string"
                },
                "available": {
                    "description": "Light in candela.",
                    "type": "boolean"
                }
            },
            "type": "object"
        }
    }
}`
