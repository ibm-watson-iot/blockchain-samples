import sys #for arguments
from flask import Flask
from flask_cors import CORS, cross_origin
import pandas as pd #for creating DataFrame
import requests #querying for information
import json #for posting to get all the information from blockchain
from sklearn.ensemble import RandomForestRegressor #the regression model
import numpy as np #standard library to convert 1d array to 2d arra
import json

app = Flask(__name__)

prediction = 0

@app.route("/", methods=['GET','OPTIONS'])
@cross_origin(origin='*')
def hello():

    #sys.argv = ['analysis.py', 'Will']
    #exec(open("analysis.py").read())

    #bring in all the information
    #getting request from the URL
    url = 'https://bb01c9bc-8a22-4329-94cb-fa722fd3bdce_vp0.us.blockchain.ibm.com:443/chaincode'
    body = {
            "jsonrpc": "2.0",
            "method": "query",
            "params": {
            "type": 1,
            "chaincodeID":{
                "name":"89684ecf448f90c8fcbf0232aab899aec47e9ac5530db4d6956fc0033a775c48aa73572253ebddfee29d449536fda6f8353570e81e026dee777832d002702521"
            },
            "ctorMsg": {
                "function":"readAsset",
                "args":["{\"assetID\":\""+"Will"+"\"}"]
            },
            "secureContext": "user_type1_fc806186e6"
        },
        "id":1234
        }
    bodyStr = json.dumps(body)
    headerReq = {'Content-Type': 'application/json', 'Accept':'application/json'}
    res = requests.post(url, bodyStr, headers=headerReq)
    columns = ["Temperature Celsius", "Temperature Fahrenheit", "Wind Speed", "Wind Gust Speed", "Wind Degrees", "Precipitation", "Carbon Reading"] #column title
    #DATAFRAME
    df = pd.DataFrame(columns=columns)
    df.fillna(0) # with 0s rather than NaNs
    if "result" in res.json() and res.json()["result"]["status"] == "OK":
        JsonResponse = json.loads(res.json()["result"]["message"])
        #check if the fields exists in the response given back
        if "reading" in JsonResponse:
            #add all the fields to dataframe
            sensorReading = JsonResponse["sensorWeatherHistory"]["sensorReading"]
            precipitation = JsonResponse["sensorWeatherHistory"]["precipitation"]
            tempCel = JsonResponse["sensorWeatherHistory"]["temperatureCelcius"]
            tempFah = JsonResponse["sensorWeatherHistory"]["temperatureFahrenheit"]
            windDegrees = JsonResponse["sensorWeatherHistory"]["windDegrees"]
            windGustSp = JsonResponse["sensorWeatherHistory"]["windGustSpeed"]
            windSp = JsonResponse["sensorWeatherHistory"]["windSpeed"]
            #adding rows to dataframe
            for i in range(len(sensorReading)):
                df.loc[len(df)] = [tempCel[i],tempFah[i],windSp[i],windGustSp[i],windDegrees[i],precipitation[i],sensorReading[i]]
            #convert dataframe to csv
            df.to_csv("output.csv", sep=',', encoding='utf-8')
            #trying to make an expected maximum likehood model 
            columnsPredict = [c for c in columns if c not in ["Carbon Reading"]]
            #what are we trying to predict
            target = "Carbon Reading"
            # Initialize the model with some parameters.
            model = RandomForestRegressor(n_estimators=100, min_samples_leaf=1, random_state=1)
            # Fit the model to the data.
            model.fit(df[columnsPredict], df[target])
            # Make predictions.
            #test = df.loc[len(df)-1][columnsPredict]
            WeatherURL = 'http://api.wunderground.com/api/62493c160d2ce863/forecast10day/q/TX/Austin.json'
            weatherRes = requests.get(WeatherURL)
            try:
                weather_res = weatherRes.json()
                #make an dataFrame
                testCol = ["Temperature Celsius", "Temperature Fahrenheit", "Wind Speed", "Wind Gust Speed", "Wind Degrees", "Precipitation"] #column title
                #DATAFRAME
                weatherDF = pd.DataFrame(columns=testCol)
                weatherDF.fillna(0) # with 0s rather than NaNs
                for i in weather_res["forecast"]["simpleforecast"]["forecastday"]:
                    cH = i["high"]["celsius"]
                    cL = i["low"]["celsius"]
                    fH = i["high"]["fahrenheit"]
                    fL = i["low"]["fahrenheit"]
                    if type(i["high"]["celsius"]) is str:
                        cH = float(cH)
                        cL = float(cL)
                    if type(i["high"]["fahrenheit"]) is str:
                        fH = float(fH)
                        fL = float(fL)
                    weatherDF.loc[len(weatherDF)] = [(cH + cL)/2, (fL + fH)/2, i["avewind"]["kph"], i["maxwind"]["kph"], (i["avewind"]["degrees"] + i["maxwind"]["degrees"])/2, i["qpf_allday"]["mm"]]        
                    #test = np.array(df.loc[len(df)-1][columnsPredict]).reshape((1, -1))
                    predictions = model.predict(weatherDF)
                #add all 10 day forecast values
                totalValue = 0
                for val in predictions:
                    totalValue = totalValue + val
                print(totalValue)
                prediction = totalValue
            except ValueError:
                print('error')

#    data = request.get_json(force=True)
    return json.dumps({"prediction": str(prediction)})


if __name__ == "__main__":
    app.run(port=2000)
