/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Risham Chokshi - Initial Contribution
*/

// ************************************
// Trade package
// RC 25 Jul 2016 Initial sensor History package
// ************************************

package main

import (
)

//updating sensorWeatherHistory once it is passed in, creating new sensorWeatherHistory attribute if one does not exist
//adding on the agrument value if it does exist
//input is ledger map and returning an updated sensorWeatherHistory or ledger
func (a *ArgsMap) updateSensorWeatherBlock(sensorVal string, timestamp string, temperatureCelsius string, temperatureFahrenheit string, windSpeed string, windGustSpeed string, precipitation string, windDegrees string, URL string) (map[string]interface {}){
    var sWBlockMap map[string]interface{} 
    //get the object from ledger if sensorWeatherHsitory already exists
    tbytes, found := getObject(*a, "sensorWeatherHistory")
    //if found is false, then a sensorWeatherHistroy does not exists and new struct needs to be created
    if found == false {
        sWBlockMap = make(map[string]interface{})
        sWBlockMap["sensorReading"] = []string{sensorVal}
        sWBlockMap["timestamp"] = []string{timestamp}
        sWBlockMap["temperatureFahrenheit"] = []string{temperatureFahrenheit}
        sWBlockMap["temperatureCelsius"] = []string{temperatureCelsius}
        sWBlockMap["windSpeed"] = []string{windSpeed}
        sWBlockMap["windGustSpeed"] = []string{windGustSpeed}
        sWBlockMap["precipitation"] = []string{precipitation}
        sWBlockMap["windDegrees"] = []string{windDegrees}
        sWBlockMap["iconUrl"] = URL
    } else {
        sWBlockMap = tbytes.(map[string]interface{})
        //appending all the new attributes
        sWBlockMap["sensorReading"] = append(sWBlockMap["sensorReading"].([]interface{}), sensorVal)
        sWBlockMap["temperatureCelsius"] = append(sWBlockMap["temperatureCelsius"].([]interface{}), temperatureCelsius)
        sWBlockMap["timestamp"] = append(sWBlockMap["timestamp"].([]interface{}), timestamp)
        sWBlockMap["temperatureFahrenheit"] = append(sWBlockMap["temperatureFahrenheit"].([]interface{}), temperatureFahrenheit)
        sWBlockMap["windSpeed"] = append(sWBlockMap["windSpeed"].([]interface{}), windSpeed)
        sWBlockMap["windGustSpeed"] = append(sWBlockMap["windGustSpeed"].([]interface{}), windGustSpeed)
        sWBlockMap["precipitation"] = append(sWBlockMap["precipitation"].([]interface{}), precipitation)
        sWBlockMap["windDegrees"] = append(sWBlockMap["windDegrees"].([]interface{}), windDegrees)
        sWBlockMap["iconUrl"] = URL
    }
    return sWBlockMap
}