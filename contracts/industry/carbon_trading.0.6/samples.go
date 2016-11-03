package main

var samples = `
{
    "contractState": {
        "activeAssets": [
            "The ID of a managed asset. The resource focal point for a smart contract."
        ],
        "nickname": "TRADELANE",
        "version": "The version number of the current contract"
    },
    "event": {
        "allottedCredits": 123.456,
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "boughtCredits": 123.456,
        "creditsForSale": 123.456,
        "creditsRequestBuy": 123.456,
        "email": "Contact information of the company will be stored here",
        "extension": {},
        "iconUrl": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "notificationRead": true,
        "phoneNum": "Contact information of the company will be stored here",
        "precipitation": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "pricePerCredit": 123.456,
        "priceRequestBuy": 123.456,
        "reading": 123.456,
        "sensorID": 123.456,
        "sensorlocation": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "soldCredits": 123.456,
        "temperatureCelsius": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "temperatureFahrenheit": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "threshold": 123.456,
        "timestamp": "2016-08-09T15:49:18.021728986-05:00",
        "tradeBuySell": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeCompany": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeCredits": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradePrice": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeTimestamp": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "updateBuyCredits": 123.456,
        "updateBuyIndex": 123.456,
        "updateSellCredits": 123.456,
        "updateSellIndex": 123.456,
        "windDegrees": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "windGustSpeed": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "windSpeed": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
    },
    "initEvent": {
        "nickname": "TRADELANE",
        "version": "The ID of a managed asset. The resource focal point for a smart contract."
    },
    "state": {
        "alerts": {
            "active": [
                "OVERCARBONEMISSION"
            ],
            "cleared": [
                "OVERCARBONEMISSION"
            ],
            "raised": [
                "OVERCARBONEMISSION"
            ]
        },
        "allottedCredits": 123.456,
        "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
        "boughtCredits": 123.456,
        "compliant": true,
        "contactInformation": {
            "email": "Contact information of the company will be stored here",
            "phoneNum": "Contact information of the company will be stored here"
        },
        "creditsBuyList": [
            "carpe noctem"
        ],
        "creditsForSale": 123.456,
        "creditsRequestBuy": 123.456,
        "creditsSellList": [
            "carpe noctem"
        ],
        "email": "Contact information of the company will be stored here",
        "extension": {},
        "iconUrl": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "lastEvent": {
            "args": [
                "parameters to the function, usually args[0] is populated with a JSON encoded event object"
            ],
            "function": "function that created this state object",
            "redirectedFromFunction": "function that originally received the event"
        },
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "notificationRead": true,
        "phoneNum": "Contact information of the company will be stored here",
        "precipitation": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "priceBuyList": [
            "carpe noctem"
        ],
        "pricePerCredit": 123.456,
        "priceRequestBuy": 123.456,
        "priceSellList": [
            "carpe noctem"
        ],
        "reading": 123.456,
        "sensorID": 123.456,
        "sensorWeatherHistory": {
            "iconUrl": "carpe noctem",
            "precipitation": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "sensorReading": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "tempCelsius": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "tempFahrenheit": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "timestamp": [
                "2016-08-09T15:49:18.022023884-05:00"
            ],
            "windDegrees": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "windGustSpeed": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ],
            "windSpeed": [
                "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
            ]
        },
        "sensorlocation": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "soldCredits": 123.456,
        "temperatureCelsius": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "temperatureFahrenheit": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "threshold": 123.456,
        "timestamp": "2016-08-09T15:49:18.022033248-05:00",
        "tradeBuySell": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeCompany": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeCredits": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeHistory": {
            "buysell": [
                "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies"
            ],
            "company": [
                "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies"
            ],
            "credits": [
                "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies"
            ],
            "price": [
                "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies"
            ],
            "timestamp": [
                "2016-08-09T15:49:18.022017853-05:00"
            ]
        },
        "tradePrice": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "tradeTimestamp": "Trade values are triggered for every trade which is processed. This contract stores every trade which will be made between two companies",
        "txntimestamp": "Transaction timestamp matching that in the blockchain.",
        "txnuuid": "Transaction UUID matching that in the blockchain.",
        "updateBuyCredits": 123.456,
        "updateBuyIndex": 123.456,
        "updateSellCredits": 123.456,
        "updateSellIndex": 123.456,
        "windDegrees": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "windGustSpeed": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value",
        "windSpeed": "Sensor and weather value will be stored in a string. So sensorWeatherData object could refer to this definition to store its value"
    }
}`