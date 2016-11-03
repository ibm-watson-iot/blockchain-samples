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
// RC 29 Jun 2016 Initial trade package
// RC 14 Jul 2016 Trade history attributes set up
// ************************************

package main

import (
)

//updating tradeBlock once it is passed in, creating new TradeBlock if one does not exist
//adding on the contract value if it does exist
//input is ledger map and returning an updated tradeblock or ledger
func (a *ArgsMap) updateTradeBlock(regCompany bool, tradeCredits string, tradePrice string, tradetimestamp string, tradeCompany string, tradeType string) (map[string]interface {}){
    var tradeBlockMap map[string]interface{} 
    //get the object from ledger if tradeHistory already exists
    tbytes, found := getObject(*a, "tradeHistory")
    //if found is false, then a tradeHistory does not exists and new struct needs to be created
    if found == false {
        tradeBlockMap = make(map[string]interface{})
        tradeBlockMap["credits"] = []string{tradeCredits}
        tradeBlockMap["price"] = []string{tradePrice}
        tradeBlockMap["timestamp"] = []string{tradetimestamp}
        if regCompany {
            tradeBlockMap["company"] = []string{tradeCompany}
            tradeBlockMap["buysell"] = []string{tradeType}
        }
    } else {
        tradeBlockMap = tbytes.(map[string]interface{})
        //appending all the new attributes
        tradeBlockMap["credits"] = append(tradeBlockMap["credits"].([]interface{}), tradeCredits)
        tradeBlockMap["price"] = append(tradeBlockMap["price"].([]interface{}), tradePrice)
        tradeBlockMap["timestamp"] = append(tradeBlockMap["timestamp"].([]interface{}), tradetimestamp)
        if regCompany {
            tradeBlockMap["company"] = append(tradeBlockMap["price"].([]interface{}), tradeCompany)
            tradeBlockMap["buysell"] = append(tradeBlockMap["buysell"].([]interface{}), tradeType)
        }
    }
    return tradeBlockMap
}