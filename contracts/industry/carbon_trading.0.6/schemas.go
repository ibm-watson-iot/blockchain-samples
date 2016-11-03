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
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "allottedCredits": {
                                "description": "defines how much a company can spend",
                                "type": "number"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "boughtCredits": {
                                "description": "Total number of credits bought from other companies",
                                "type": "number"
                            },
                            "creditsForSale": {
                                "description": "Total credits which are going to be put on sale by a company",
                                "type": "number"
                            },
                            "creditsRequestBuy": {
                                "description": "Total credits requested to buy from the market",
                                "type": "number"
                            },
                            "email": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "iconUrl": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                            "notificationRead": {
                                "description": "Value will be true if company saw their weather notification, false otherwise",
                                "type": "boolean"
                            },
                            "phoneNum": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "precipitation": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "pricePerCredit": {
                                "description": "Price set for every credit that is put on sale",
                                "type": "number"
                            },
                            "priceRequestBuy": {
                                "description": "Price put per credit requested to buy from the market",
                                "type": "number"
                            },
                            "reading": {
                                "description": "defines one reading for a sensor",
                                "type": "number"
                            },
                            "sensorID": {
                                "description": "defines one sensor in a company",
                                "type": "number"
                            },
                            "sensorlocation": {
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
                            "soldCredits": {
                                "description": "Total number of credits sold to other companies",
                                "type": "number"
                            },
                            "temperatureCelsius": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "temperatureFahrenheit": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "threshold": {
                                "description": "limit on credit consumption before it alerts",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            },
                            "tradeBuySell": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCompany": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCredits": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradePrice": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeTimestamp": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "updateBuyCredits": {
                                "description": "Credits that need to be updated in buy list",
                                "type": "number"
                            },
                            "updateBuyIndex": {
                                "description": "Index of the buy list array that needs to be updated",
                                "type": "number"
                            },
                            "updateSellCredits": {
                                "description": "Credits that need to be updated in sell list",
                                "type": "number"
                            },
                            "updateSellIndex": {
                                "description": "Index of the sell list array that needs to be updated",
                                "type": "number"
                            },
                            "windDegrees": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windGustSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "allottedCredits": {
                                "description": "defines how much a company can spend",
                                "type": "number"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "boughtCredits": {
                                "description": "Total number of credits bought from other companies",
                                "type": "number"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "contactInformation": {
                                "description": "",
                                "properties": {
                                    "email": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    },
                                    "phoneNum": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "creditsBuyList": {
                                "description": "List of credits requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "creditsForSale": {
                                "description": "Total credits which are going to be put on sale by a company",
                                "type": "number"
                            },
                            "creditsRequestBuy": {
                                "description": "Total credits requested to buy from the market",
                                "type": "number"
                            },
                            "creditsSellList": {
                                "description": "List of credits company willing to sell",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "email": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "iconUrl": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
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
                            "notificationRead": {
                                "description": "Value will be true if company saw their weather notification, false otherwise",
                                "type": "boolean"
                            },
                            "phoneNum": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "precipitation": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "priceBuyList": {
                                "description": "List of price requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "pricePerCredit": {
                                "description": "Price set for every credit that is put on sale",
                                "type": "number"
                            },
                            "priceRequestBuy": {
                                "description": "Price put per credit requested to buy from the market",
                                "type": "number"
                            },
                            "priceSellList": {
                                "description": "List of price for every credit put on sell by a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "reading": {
                                "description": "defines one reading for a sensor",
                                "type": "number"
                            },
                            "sensorID": {
                                "description": "defines one sensor in a company",
                                "type": "number"
                            },
                            "sensorWeatherHistory": {
                                "description": "sensorReading means history of all the carbon readings from the sensor, timestamp refers to time it was recorded and all the other fields refers to weather data",
                                "properties": {
                                    "iconUrl": {
                                        "type": "string"
                                    },
                                    "precipitation": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "sensorReading": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempCelsius": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempFahrenheit": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windDegrees": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windGustSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "sensorlocation": {
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
                            "soldCredits": {
                                "description": "Total number of credits sold to other companies",
                                "type": "number"
                            },
                            "temperatureCelsius": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "temperatureFahrenheit": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "threshold": {
                                "description": "limit on credit consumption before it alerts",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            },
                            "tradeBuySell": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCompany": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCredits": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeHistory": {
                                "description": "Sold means how many credits were traded and Price refers for how much per credit. Company means institution trade was made to. Timestamp means what time trade occured. BuySell attribute is to indicate if it was a buy or sell ",
                                "properties": {
                                    "buysell": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "company": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "credits": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "price": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "tradePrice": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeTimestamp": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "updateBuyCredits": {
                                "description": "Credits that need to be updated in buy list",
                                "type": "number"
                            },
                            "updateBuyIndex": {
                                "description": "Index of the buy list array that needs to be updated",
                                "type": "number"
                            },
                            "updateSellCredits": {
                                "description": "Credits that need to be updated in sell list",
                                "type": "number"
                            },
                            "updateSellIndex": {
                                "description": "Index of the sell list array that needs to be updated",
                                "type": "number"
                            },
                            "windDegrees": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windGustSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                    "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                    "properties": {
                        "alerts": {
                            "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                            "properties": {
                                "active": {
                                    "items": {
                                        "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                        "enum": [
                                            "OVERCARBONEMISSION"
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
                                            "OVERCARBONEMISSION"
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
                                            "OVERCARBONEMISSION"
                                        ],
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                }
                            },
                            "type": "object"
                        },
                        "allottedCredits": {
                            "description": "defines how much a company can spend",
                            "type": "number"
                        },
                        "assetID": {
                            "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                            "type": "string"
                        },
                        "boughtCredits": {
                            "description": "Total number of credits bought from other companies",
                            "type": "number"
                        },
                        "compliant": {
                            "description": "A contract-specific indication that this asset is compliant.",
                            "type": "boolean"
                        },
                        "contactInformation": {
                            "description": "",
                            "properties": {
                                "email": {
                                    "description": "Contact information of the company will be stored here",
                                    "type": "string"
                                },
                                "phoneNum": {
                                    "description": "Contact information of the company will be stored here",
                                    "type": "string"
                                }
                            },
                            "type": "object"
                        },
                        "creditsBuyList": {
                            "description": "List of credits requested to buy from a company",
                            "items": {
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "creditsForSale": {
                            "description": "Total credits which are going to be put on sale by a company",
                            "type": "number"
                        },
                        "creditsRequestBuy": {
                            "description": "Total credits requested to buy from the market",
                            "type": "number"
                        },
                        "creditsSellList": {
                            "description": "List of credits company willing to sell",
                            "items": {
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "email": {
                            "description": "Contact information of the company will be stored here",
                            "type": "string"
                        },
                        "extension": {
                            "description": "Application-managed state. Opaque to contract.",
                            "properties": {},
                            "type": "object"
                        },
                        "iconUrl": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
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
                        "notificationRead": {
                            "description": "Value will be true if company saw their weather notification, false otherwise",
                            "type": "boolean"
                        },
                        "phoneNum": {
                            "description": "Contact information of the company will be stored here",
                            "type": "string"
                        },
                        "precipitation": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
                        },
                        "priceBuyList": {
                            "description": "List of price requested to buy from a company",
                            "items": {
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "pricePerCredit": {
                            "description": "Price set for every credit that is put on sale",
                            "type": "number"
                        },
                        "priceRequestBuy": {
                            "description": "Price put per credit requested to buy from the market",
                            "type": "number"
                        },
                        "priceSellList": {
                            "description": "List of price for every credit put on sell by a company",
                            "items": {
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "reading": {
                            "description": "defines one reading for a sensor",
                            "type": "number"
                        },
                        "sensorID": {
                            "description": "defines one sensor in a company",
                            "type": "number"
                        },
                        "sensorWeatherHistory": {
                            "description": "sensorReading means history of all the carbon readings from the sensor, timestamp refers to time it was recorded and all the other fields refers to weather data",
                            "properties": {
                                "iconUrl": {
                                    "type": "string"
                                },
                                "precipitation": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "sensorReading": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "tempCelsius": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "tempFahrenheit": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "timestamp": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "windDegrees": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "windGustSpeed": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "windSpeed": {
                                    "items": {
                                        "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                }
                            },
                            "type": "object"
                        },
                        "sensorlocation": {
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
                        "soldCredits": {
                            "description": "Total number of credits sold to other companies",
                            "type": "number"
                        },
                        "temperatureCelsius": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
                        },
                        "temperatureFahrenheit": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
                        },
                        "threshold": {
                            "description": "limit on credit consumption before it alerts",
                            "type": "number"
                        },
                        "timestamp": {
                            "description": "RFC3339nanos formatted timestamp.",
                            "type": "string"
                        },
                        "tradeBuySell": {
                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                            "type": "string"
                        },
                        "tradeCompany": {
                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                            "type": "string"
                        },
                        "tradeCredits": {
                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                            "type": "string"
                        },
                        "tradeHistory": {
                            "description": "Sold means how many credits were traded and Price refers for how much per credit. Company means institution trade was made to. Timestamp means what time trade occured. BuySell attribute is to indicate if it was a buy or sell ",
                            "properties": {
                                "buysell": {
                                    "items": {
                                        "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "company": {
                                    "items": {
                                        "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "credits": {
                                    "items": {
                                        "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "price": {
                                    "items": {
                                        "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                },
                                "timestamp": {
                                    "items": {
                                        "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                        "type": "string"
                                    },
                                    "minItems": 0,
                                    "type": "array"
                                }
                            },
                            "type": "object"
                        },
                        "tradePrice": {
                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                            "type": "string"
                        },
                        "tradeTimestamp": {
                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                            "type": "string"
                        },
                        "txntimestamp": {
                            "description": "Transaction timestamp matching that in the blockchain.",
                            "type": "string"
                        },
                        "txnuuid": {
                            "description": "Transaction UUID matching that in the blockchain.",
                            "type": "string"
                        },
                        "updateBuyCredits": {
                            "description": "Credits that need to be updated in buy list",
                            "type": "number"
                        },
                        "updateBuyIndex": {
                            "description": "Index of the buy list array that needs to be updated",
                            "type": "number"
                        },
                        "updateSellCredits": {
                            "description": "Credits that need to be updated in sell list",
                            "type": "number"
                        },
                        "updateSellIndex": {
                            "description": "Index of the sell list array that needs to be updated",
                            "type": "number"
                        },
                        "windDegrees": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
                        },
                        "windGustSpeed": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                            "type": "string"
                        },
                        "windSpeed": {
                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "allottedCredits": {
                                "description": "defines how much a company can spend",
                                "type": "number"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "boughtCredits": {
                                "description": "Total number of credits bought from other companies",
                                "type": "number"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "contactInformation": {
                                "description": "",
                                "properties": {
                                    "email": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    },
                                    "phoneNum": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "creditsBuyList": {
                                "description": "List of credits requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "creditsForSale": {
                                "description": "Total credits which are going to be put on sale by a company",
                                "type": "number"
                            },
                            "creditsRequestBuy": {
                                "description": "Total credits requested to buy from the market",
                                "type": "number"
                            },
                            "creditsSellList": {
                                "description": "List of credits company willing to sell",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "email": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "iconUrl": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
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
                            "notificationRead": {
                                "description": "Value will be true if company saw their weather notification, false otherwise",
                                "type": "boolean"
                            },
                            "phoneNum": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "precipitation": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "priceBuyList": {
                                "description": "List of price requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "pricePerCredit": {
                                "description": "Price set for every credit that is put on sale",
                                "type": "number"
                            },
                            "priceRequestBuy": {
                                "description": "Price put per credit requested to buy from the market",
                                "type": "number"
                            },
                            "priceSellList": {
                                "description": "List of price for every credit put on sell by a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "reading": {
                                "description": "defines one reading for a sensor",
                                "type": "number"
                            },
                            "sensorID": {
                                "description": "defines one sensor in a company",
                                "type": "number"
                            },
                            "sensorWeatherHistory": {
                                "description": "sensorReading means history of all the carbon readings from the sensor, timestamp refers to time it was recorded and all the other fields refers to weather data",
                                "properties": {
                                    "iconUrl": {
                                        "type": "string"
                                    },
                                    "precipitation": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "sensorReading": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempCelsius": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempFahrenheit": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windDegrees": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windGustSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "sensorlocation": {
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
                            "soldCredits": {
                                "description": "Total number of credits sold to other companies",
                                "type": "number"
                            },
                            "temperatureCelsius": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "temperatureFahrenheit": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "threshold": {
                                "description": "limit on credit consumption before it alerts",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            },
                            "tradeBuySell": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCompany": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCredits": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeHistory": {
                                "description": "Sold means how many credits were traded and Price refers for how much per credit. Company means institution trade was made to. Timestamp means what time trade occured. BuySell attribute is to indicate if it was a buy or sell ",
                                "properties": {
                                    "buysell": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "company": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "credits": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "price": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "tradePrice": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeTimestamp": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "updateBuyCredits": {
                                "description": "Credits that need to be updated in buy list",
                                "type": "number"
                            },
                            "updateBuyIndex": {
                                "description": "Index of the buy list array that needs to be updated",
                                "type": "number"
                            },
                            "updateSellCredits": {
                                "description": "Credits that need to be updated in sell list",
                                "type": "number"
                            },
                            "updateSellIndex": {
                                "description": "Index of the sell list array that needs to be updated",
                                "type": "number"
                            },
                            "windDegrees": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windGustSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                        "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
                        "properties": {
                            "alerts": {
                                "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                                "properties": {
                                    "active": {
                                        "items": {
                                            "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                            "enum": [
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
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
                                                "OVERCARBONEMISSION"
                                            ],
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "allottedCredits": {
                                "description": "defines how much a company can spend",
                                "type": "number"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "boughtCredits": {
                                "description": "Total number of credits bought from other companies",
                                "type": "number"
                            },
                            "compliant": {
                                "description": "A contract-specific indication that this asset is compliant.",
                                "type": "boolean"
                            },
                            "contactInformation": {
                                "description": "",
                                "properties": {
                                    "email": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    },
                                    "phoneNum": {
                                        "description": "Contact information of the company will be stored here",
                                        "type": "string"
                                    }
                                },
                                "type": "object"
                            },
                            "creditsBuyList": {
                                "description": "List of credits requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "creditsForSale": {
                                "description": "Total credits which are going to be put on sale by a company",
                                "type": "number"
                            },
                            "creditsRequestBuy": {
                                "description": "Total credits requested to buy from the market",
                                "type": "number"
                            },
                            "creditsSellList": {
                                "description": "List of credits company willing to sell",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "email": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "iconUrl": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
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
                            "notificationRead": {
                                "description": "Value will be true if company saw their weather notification, false otherwise",
                                "type": "boolean"
                            },
                            "phoneNum": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "precipitation": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "priceBuyList": {
                                "description": "List of price requested to buy from a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "pricePerCredit": {
                                "description": "Price set for every credit that is put on sale",
                                "type": "number"
                            },
                            "priceRequestBuy": {
                                "description": "Price put per credit requested to buy from the market",
                                "type": "number"
                            },
                            "priceSellList": {
                                "description": "List of price for every credit put on sell by a company",
                                "items": {
                                    "type": "string"
                                },
                                "minItems": 0,
                                "type": "array"
                            },
                            "reading": {
                                "description": "defines one reading for a sensor",
                                "type": "number"
                            },
                            "sensorID": {
                                "description": "defines one sensor in a company",
                                "type": "number"
                            },
                            "sensorWeatherHistory": {
                                "description": "sensorReading means history of all the carbon readings from the sensor, timestamp refers to time it was recorded and all the other fields refers to weather data",
                                "properties": {
                                    "iconUrl": {
                                        "type": "string"
                                    },
                                    "precipitation": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "sensorReading": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempCelsius": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "tempFahrenheit": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windDegrees": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windGustSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "windSpeed": {
                                        "items": {
                                            "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "sensorlocation": {
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
                            "soldCredits": {
                                "description": "Total number of credits sold to other companies",
                                "type": "number"
                            },
                            "temperatureCelsius": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "temperatureFahrenheit": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "threshold": {
                                "description": "limit on credit consumption before it alerts",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            },
                            "tradeBuySell": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCompany": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCredits": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeHistory": {
                                "description": "Sold means how many credits were traded and Price refers for how much per credit. Company means institution trade was made to. Timestamp means what time trade occured. BuySell attribute is to indicate if it was a buy or sell ",
                                "properties": {
                                    "buysell": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "company": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "credits": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "price": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    },
                                    "timestamp": {
                                        "items": {
                                            "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                            "type": "string"
                                        },
                                        "minItems": 0,
                                        "type": "array"
                                    }
                                },
                                "type": "object"
                            },
                            "tradePrice": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeTimestamp": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "txntimestamp": {
                                "description": "Transaction timestamp matching that in the blockchain.",
                                "type": "string"
                            },
                            "txnuuid": {
                                "description": "Transaction UUID matching that in the blockchain.",
                                "type": "string"
                            },
                            "updateBuyCredits": {
                                "description": "Credits that need to be updated in buy list",
                                "type": "number"
                            },
                            "updateBuyIndex": {
                                "description": "Index of the buy list array that needs to be updated",
                                "type": "number"
                            },
                            "updateSellCredits": {
                                "description": "Credits that need to be updated in sell list",
                                "type": "number"
                            },
                            "updateSellIndex": {
                                "description": "Index of the sell list array that needs to be updated",
                                "type": "number"
                            },
                            "windDegrees": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windGustSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                        "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
                        "properties": {
                            "allottedCredits": {
                                "description": "defines how much a company can spend",
                                "type": "number"
                            },
                            "assetID": {
                                "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                                "type": "string"
                            },
                            "boughtCredits": {
                                "description": "Total number of credits bought from other companies",
                                "type": "number"
                            },
                            "creditsForSale": {
                                "description": "Total credits which are going to be put on sale by a company",
                                "type": "number"
                            },
                            "creditsRequestBuy": {
                                "description": "Total credits requested to buy from the market",
                                "type": "number"
                            },
                            "email": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "extension": {
                                "description": "Application-managed state. Opaque to contract.",
                                "properties": {},
                                "type": "object"
                            },
                            "iconUrl": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                            "notificationRead": {
                                "description": "Value will be true if company saw their weather notification, false otherwise",
                                "type": "boolean"
                            },
                            "phoneNum": {
                                "description": "Contact information of the company will be stored here",
                                "type": "string"
                            },
                            "precipitation": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "pricePerCredit": {
                                "description": "Price set for every credit that is put on sale",
                                "type": "number"
                            },
                            "priceRequestBuy": {
                                "description": "Price put per credit requested to buy from the market",
                                "type": "number"
                            },
                            "reading": {
                                "description": "defines one reading for a sensor",
                                "type": "number"
                            },
                            "sensorID": {
                                "description": "defines one sensor in a company",
                                "type": "number"
                            },
                            "sensorlocation": {
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
                            "soldCredits": {
                                "description": "Total number of credits sold to other companies",
                                "type": "number"
                            },
                            "temperatureCelsius": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "temperatureFahrenheit": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "threshold": {
                                "description": "limit on credit consumption before it alerts",
                                "type": "number"
                            },
                            "timestamp": {
                                "description": "RFC3339nanos formatted timestamp.",
                                "type": "string"
                            },
                            "tradeBuySell": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCompany": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeCredits": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradePrice": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "tradeTimestamp": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "updateBuyCredits": {
                                "description": "Credits that need to be updated in buy list",
                                "type": "number"
                            },
                            "updateBuyIndex": {
                                "description": "Index of the buy list array that needs to be updated",
                                "type": "number"
                            },
                            "updateSellCredits": {
                                "description": "Credits that need to be updated in sell list",
                                "type": "number"
                            },
                            "updateSellIndex": {
                                "description": "Index of the sell list array that needs to be updated",
                                "type": "number"
                            },
                            "windDegrees": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windGustSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "windSpeed": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
            "description": "The set of writable properties that define an asset's state. For asset creation, the only mandatory property is the 'assetID'. Updates should include at least one other writable property. This exemplifies the IoT contract pattern 'partial state as event'.",
            "properties": {
                "allottedCredits": {
                    "description": "defines how much a company can spend",
                    "type": "number"
                },
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "boughtCredits": {
                    "description": "Total number of credits bought from other companies",
                    "type": "number"
                },
                "creditsForSale": {
                    "description": "Total credits which are going to be put on sale by a company",
                    "type": "number"
                },
                "creditsRequestBuy": {
                    "description": "Total credits requested to buy from the market",
                    "type": "number"
                },
                "email": {
                    "description": "Contact information of the company will be stored here",
                    "type": "string"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {},
                    "type": "object"
                },
                "iconUrl": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
                "notificationRead": {
                    "description": "Value will be true if company saw their weather notification, false otherwise",
                    "type": "boolean"
                },
                "phoneNum": {
                    "description": "Contact information of the company will be stored here",
                    "type": "string"
                },
                "precipitation": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "pricePerCredit": {
                    "description": "Price set for every credit that is put on sale",
                    "type": "number"
                },
                "priceRequestBuy": {
                    "description": "Price put per credit requested to buy from the market",
                    "type": "number"
                },
                "reading": {
                    "description": "defines one reading for a sensor",
                    "type": "number"
                },
                "sensorID": {
                    "description": "defines one sensor in a company",
                    "type": "number"
                },
                "sensorlocation": {
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
                "soldCredits": {
                    "description": "Total number of credits sold to other companies",
                    "type": "number"
                },
                "temperatureCelsius": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "temperatureFahrenheit": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "threshold": {
                    "description": "limit on credit consumption before it alerts",
                    "type": "number"
                },
                "timestamp": {
                    "description": "RFC3339nanos formatted timestamp.",
                    "type": "string"
                },
                "tradeBuySell": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeCompany": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeCredits": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradePrice": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeTimestamp": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "updateBuyCredits": {
                    "description": "Credits that need to be updated in buy list",
                    "type": "number"
                },
                "updateBuyIndex": {
                    "description": "Index of the buy list array that needs to be updated",
                    "type": "number"
                },
                "updateSellCredits": {
                    "description": "Credits that need to be updated in sell list",
                    "type": "number"
                },
                "updateSellIndex": {
                    "description": "Index of the sell list array that needs to be updated",
                    "type": "number"
                },
                "windDegrees": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "windGustSpeed": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "windSpeed": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
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
            "description": "A set of properties that constitute a complete asset state. Includes event properties and any other calculated properties such as compliance related alerts.",
            "properties": {
                "alerts": {
                    "description": "Active means that the alert is in force in this state. Raised means that the alert became active as the result of the event that generated this state. Cleared means that the alert became inactive as the result of the event that generated this state.",
                    "properties": {
                        "active": {
                            "items": {
                                "description": "Alerts are triggered or cleared by rules that are run against incoming events. This contract considers any active alert to created a state of non-compliance.",
                                "enum": [
                                    "OVERCARBONEMISSION"
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
                                    "OVERCARBONEMISSION"
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
                                    "OVERCARBONEMISSION"
                                ],
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "allottedCredits": {
                    "description": "defines how much a company can spend",
                    "type": "number"
                },
                "assetID": {
                    "description": "The ID of a managed asset. The resource focal point for a smart contract.",
                    "type": "string"
                },
                "boughtCredits": {
                    "description": "Total number of credits bought from other companies",
                    "type": "number"
                },
                "compliant": {
                    "description": "A contract-specific indication that this asset is compliant.",
                    "type": "boolean"
                },
                "contactInformation": {
                    "description": "",
                    "properties": {
                        "email": {
                            "description": "Contact information of the company will be stored here",
                            "type": "string"
                        },
                        "phoneNum": {
                            "description": "Contact information of the company will be stored here",
                            "type": "string"
                        }
                    },
                    "type": "object"
                },
                "creditsBuyList": {
                    "description": "List of credits requested to buy from a company",
                    "items": {
                        "type": "string"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "creditsForSale": {
                    "description": "Total credits which are going to be put on sale by a company",
                    "type": "number"
                },
                "creditsRequestBuy": {
                    "description": "Total credits requested to buy from the market",
                    "type": "number"
                },
                "creditsSellList": {
                    "description": "List of credits company willing to sell",
                    "items": {
                        "type": "string"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "email": {
                    "description": "Contact information of the company will be stored here",
                    "type": "string"
                },
                "extension": {
                    "description": "Application-managed state. Opaque to contract.",
                    "properties": {},
                    "type": "object"
                },
                "iconUrl": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
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
                "notificationRead": {
                    "description": "Value will be true if company saw their weather notification, false otherwise",
                    "type": "boolean"
                },
                "phoneNum": {
                    "description": "Contact information of the company will be stored here",
                    "type": "string"
                },
                "precipitation": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "priceBuyList": {
                    "description": "List of price requested to buy from a company",
                    "items": {
                        "type": "string"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "pricePerCredit": {
                    "description": "Price set for every credit that is put on sale",
                    "type": "number"
                },
                "priceRequestBuy": {
                    "description": "Price put per credit requested to buy from the market",
                    "type": "number"
                },
                "priceSellList": {
                    "description": "List of price for every credit put on sell by a company",
                    "items": {
                        "type": "string"
                    },
                    "minItems": 0,
                    "type": "array"
                },
                "reading": {
                    "description": "defines one reading for a sensor",
                    "type": "number"
                },
                "sensorID": {
                    "description": "defines one sensor in a company",
                    "type": "number"
                },
                "sensorWeatherHistory": {
                    "description": "sensorReading means history of all the carbon readings from the sensor, timestamp refers to time it was recorded and all the other fields refers to weather data",
                    "properties": {
                        "iconUrl": {
                            "type": "string"
                        },
                        "precipitation": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "sensorReading": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "tempCelsius": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "tempFahrenheit": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "timestamp": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "windDegrees": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "windGustSpeed": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "windSpeed": {
                            "items": {
                                "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "sensorlocation": {
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
                "soldCredits": {
                    "description": "Total number of credits sold to other companies",
                    "type": "number"
                },
                "temperatureCelsius": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "temperatureFahrenheit": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "threshold": {
                    "description": "limit on credit consumption before it alerts",
                    "type": "number"
                },
                "timestamp": {
                    "description": "RFC3339nanos formatted timestamp.",
                    "type": "string"
                },
                "tradeBuySell": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeCompany": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeCredits": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeHistory": {
                    "description": "Sold means how many credits were traded and Price refers for how much per credit. Company means institution trade was made to. Timestamp means what time trade occured. BuySell attribute is to indicate if it was a buy or sell ",
                    "properties": {
                        "buysell": {
                            "items": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "company": {
                            "items": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "credits": {
                            "items": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "price": {
                            "items": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        },
                        "timestamp": {
                            "items": {
                                "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                                "type": "string"
                            },
                            "minItems": 0,
                            "type": "array"
                        }
                    },
                    "type": "object"
                },
                "tradePrice": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "tradeTimestamp": {
                    "description": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
                    "type": "string"
                },
                "txntimestamp": {
                    "description": "Transaction timestamp matching that in the blockchain.",
                    "type": "string"
                },
                "txnuuid": {
                    "description": "Transaction UUID matching that in the blockchain.",
                    "type": "string"
                },
                "updateBuyCredits": {
                    "description": "Credits that need to be updated in buy list",
                    "type": "number"
                },
                "updateBuyIndex": {
                    "description": "Index of the buy list array that needs to be updated",
                    "type": "number"
                },
                "updateSellCredits": {
                    "description": "Credits that need to be updated in sell list",
                    "type": "number"
                },
                "updateSellIndex": {
                    "description": "Index of the sell list array that needs to be updated",
                    "type": "number"
                },
                "windDegrees": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "windGustSpeed": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                },
                "windSpeed": {
                    "description": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
                    "type": "string"
                }
            },
            "type": "object"
        }
    }
}`