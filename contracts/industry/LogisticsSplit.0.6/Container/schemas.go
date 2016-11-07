package main

var schemas = `
{
    "API": {
        "createContainerLogistics": {
            "description": "Create an asset. One argument, a JSON encoded event. Container No is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. Container No is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "containerno": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
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
                            },
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "humidity": {
                                "description": "Humidity in percentage.",
                                "type": "number"
                            },
                            "light": {
                                "description": "Light in candela.",
                                "type": "number"
                            },
                             "acceleration": {
                                "description": "acceleration -gforce / shock.",
                                "type": "number"
                            },
                            "airquality": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "oxygen": {
                                        "type": "number"
                                    },
                                    "carbondioxide": {
                                        "type": "number"
                                    },
                                    "ethylene": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "containerno"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "createContainerLogistics function",
                    "enum": [
                        "createContainerLogistics"
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
        "readContainerCurrentStatus": {
            "description": "Returns the state an asset. Argument is a JSON encoded string. Container No is the only accepted property.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "An object containing only an Container No for use as an argument to read or delete.",
                        "properties": {
                            "containerno": {
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
                    "description": "readContainerCurrentStatus function",
                    "enum": [
                        "readContainerCurrentStatus"
                    ],
                    "type": "string"
                },
                "method": "query",
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "containerno": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "carrier": {
                            "description": "transport entity currently in possession of asset",
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
                        },
                        "temperature": {
                            "description": "Temperature of the asset in CELSIUS.",
                            "type": "number"
                        },
                        "humidity": {
                            "description": "Humidity in percentage.",
                            "type": "number"
                        },
                        "light": {
                            "description": "Light in candela.",
                            "type": "number"
                        },
                            "acceleration": {
                            "description": "acceleration -gforce / shock.",
                            "type": "number"
                        },
                        "airquality": {
                            "description": "A geographical coordinate",
                            "properties": {
                                "oxygen": {
                                    "type": "number"
                                },
                                "carbondioxide": {
                                    "type": "number"
                                },
                                "ethylene": {
                                    "type": "number"
                                }
                            },
                            "type": "object"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readContainerLogisitcsSchemas": {
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
                    "description": "readContainerLogisitcsSchemas function",
                    "enum": [
                        "readContainerLogisitcsSchemas"
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
        "updateContainerLogistics": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. Container No is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. Container No is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "containerno": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "carrier": {
                                "description": "transport entity currently in possession of asset",
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
                            },
                            "temperature": {
                                "description": "Temperature of the asset in CELSIUS.",
                                "type": "number"
                            },
                            "humidity": {
                                "description": "Humidity in percentage.",
                                "type": "number"
                            },
                            "light": {
                                "description": "Light in candela.",
                                "type": "number"
                            },
                             "acceleration": {
                                "description": "acceleration -gforce / shock.",
                                "type": "number"
                            },
                            "airquality": {
                                "description": "A geographical coordinate",
                                "properties": {
                                    "oxygen": {
                                        "type": "number"
                                    },
                                    "carbondioxide": {
                                        "type": "number"
                                    },
                                    "ethylene": {
                                        "type": "number"
                                    }
                                },
                                "type": "object"
                            }
                        },
                        "required": [
                            "containerno"
                        ],
                        "type": "object"
                    },
                    "maxItems": 1,
                    "minItems": 1,
                    "type": "array"
                },
                "function": {
                    "description": "updateContainerLogistics function",
                    "enum": [
                        "updateContainerLogistics"
                    ],
                    "type": "string"
                },
                 "method": "invoke"
            },
            "type": "object"
        }
    },
    "objectModelSchemas": {
        "containernoKey": {
            "description": "An object containing only an Container No for use as an argument to read or delete.",
            "properties": {
                "containerno": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                }
            },
            "type": "object"
        },
        "event": {
            "description": "A set of fields that constitute the writable fields in an asset's state. Container no. is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
            "properties": {
                "containerno": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
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
                },
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                },
                "humidity": {
                    "description": "Humidity in percentage.",
                    "type": "number"
                },
                "light": {
                    "description": "Light in candela.",
                    "type": "number"
                },
                    "acceleration": {
                    "description": "acceleration -gforce / shock.",
                    "type": "number"
                },
                "airquality": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "oxygen": {
                            "type": "number"
                        },
                        "carbondioxide": {
                            "type": "number"
                        },
                        "ethylene": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                }
            },
            "required": [
                "containerno"
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
                "containerno": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "carrier": {
                    "description": "transport entity currently in possession of asset",
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
                },
                "temperature": {
                    "description": "Temperature of the asset in CELSIUS.",
                    "type": "number"
                },
                "humidity": {
                    "description": "Humidity in percentage.",
                    "type": "number"
                },
                "light": {
                    "description": "Light in candela.",
                    "type": "number"
                },
                    "acceleration": {
                    "description": "acceleration -gforce / shock.",
                    "type": "number"
                },
                "airquality": {
                    "description": "A geographical coordinate",
                    "properties": {
                        "oxygen": {
                            "type": "number"
                        },
                        "carbondioxide": {
                            "type": "number"
                        },
                        "ethylene": {
                            "type": "number"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        }
    }
}`
