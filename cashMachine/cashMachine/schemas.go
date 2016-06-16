package main

var schemas = `
{
    "API": {
        "createCashMachine": {
            "description": "Create an asset. One argument, a JSON encoded event. AssetID is required with zero or more writable properties. Establishes an initial asset state.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                            "assetID": {
                                "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                                "type": "string"
                            },
                            "actiontype": {
                                "description": "A String with one of three values is expected: InitialBalance, Deposit or Withdraw",
                                "type": "string"
                            },
                            "amount": {
                                "description": "The transaction amount.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "Current timestamp. If not sent in, the transaction time is set",
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
                    "description": "createCashMachine function",
                    "enum": [
                        "createCashMachine"
                    ],
                    "type": "string"
                }
            },
            "type": "object"
        },
        "deleteCashMachine": {
            "description": "Delete an asset and its history. Argument is a JSON encoded string containing only an assetID.",
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
                    "description": "deleteCashMachine function",
                    "enum": [
                        "deleteCashMachine"
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
                            "version": {
                                "description": "The ID of a managed asset, the cash machine.",
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
        "readCashMachine": {
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
                    "description": "readCashMachine function",
                    "enum": [
                        "readCashMachine"
                    ],
                    "type": "string"
                },
                "result": {
                    "description": "A set of fields that constitute the complete asset state.",
                    "properties": {
                        "assetID": {
                            "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                            "type": "string"
                        },
                        "actiontype": {
                            "description": "The last transaction: InitialBalance, Deposit or Withdraw",
                            "type": "string"
                        },
                        "amount": {
                            "description": "The last transaction amount.",
                            "type": "number"
                        },
                        "balance": {
                            "description": "The current balance of the asset.",
                            "type": "number"
                        },
                        "timestamp": {
                            "description": "Current timestamp. If not sent in, the transaction time is set",
                            "type": "string"
                        }
                    },
                    "type": "object"
                }
            },
            "type": "object"
        },
        "readCashMachineHistory": {
            "description": "Requests a specified number of history states for an assets. Returns an array of states sorted with the most recent first. AssetID is required and count is optional. A missing count, a count of zero, or too large a count returns all existing history states.",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "Requested assetID",
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
                    "description": "readCashMachineHistory function",
                    "enum": [
                        "readCashMachineHistory"
                    ],
                    "type": "string"
                },
                "result": {
                    "description": "an array of states for one asset sorted by timestamp with the most recent entry first",
                    "items": {
                        "description": "A set of fields that constitute the complete asset state.",
                        "properties": {
                            "assetID": {
                            "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                            "type": "string"
                            },
                            "actiontype": {
                                "description": "The last transaction: InitialBalance, Deposit or Withdraw",
                                "type": "string"
                            },
                            "amount": {
                                "description": "The last transaction amount.",
                                "type": "number"
                            },
                            "balance": {
                                "description": "The current balance of the asset.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "Current timestamp. If not sent in, the transaction time is set",
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
        "updateCashMachine": {
            "description": "Update the state of an asset. The one argument is a JSON encoded event. AssetID is required along with one or more writable properties. Establishes the next asset state. ",
            "properties": {
                "args": {
                    "description": "args are JSON encoded strings",
                    "items": {
                        "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                        "properties": {
                           "assetID": {
                                "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                                "type": "string"
                            },
                            "actiontype": {
                                "description": "A String with one of three values is expected: InitialBalance, Deposit or Withdraw",
                                "type": "string"
                            },
                            "amount": {
                                "description": "The transaction amount.",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "Current timestamp. If not sent in, the transaction time is set",
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
                    "description": "updateCashMachine function",
                    "enum": [
                        "updateCashMachine"
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
                    "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                    "type": "string"
                },
                "actiontype": {
                    "description": "A String with one of three values is expected: InitialBalance, Deposit or Withdraw",
                    "type": "string"
                },
                "amount": {
                    "description": "The transaction amount.",
                    "type": "number"
                },
                "timestamp": {
                    "description": "Current timestamp. If not sent in, the transaction time is set",
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
                "assetID": {
                    "description": "The ID of a managed asset.In this case, the uniqie ID of the case machine.",
                    "type": "string"
                },
                "actiontype": {
                    "description": "A String with one of three values is expected: InitialBalance, Deposit or Withdraw",
                    "type": "string"
                },
                "amount": {
                    "description": "The transaction amount.",
                    "type": "number"
                },
                "balance": {
                    "description": "The balance on the asset.",
                    "type": "number"
                },
                "timestamp": {
                    "description": "Current timestamp. If not sent in, the transaction time is set",
                    "type": "string"
                }
            },
            "type": "object"
        }
    }
}`