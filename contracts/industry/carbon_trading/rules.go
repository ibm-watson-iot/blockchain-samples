/*
Copyright (c) 2016 IBM Corporation and other Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.

Contributors:
Kim Letkeman - Initial Contribution
Risham Chokshi - Initial Contribution
*/


// ************************************
// Rules for Contract
// KL 23 June 2016 Initial rules package for giving alerts based on threshold
// RC 6 July 2016  Added all rules dealing with bought, sold, allotted and used credits
// ************************************

package main

import ()

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool) {
    log.Debugf("Executing rules input: %v", *alerts)
    var internal = (*alerts).asAlertStatusInternal()

    // rule 1 -- overtemp
    internal.overCarbRule(a)
    // rule 2 -- ???

    // now transform internal back to external in order to give the contract the
    // appropriate JSON to send externally
    *alerts = internal.asAlertStatus()
    log.Debugf("Executing rules output: %v", *alerts)

    // set compliance true means out of compliance
    compliant := internal.calculateContractCompliance(a)
    // returns true if anything at all is active (i.e. NOT compliant)
    return !compliant
}

//***********************************
//**           RULES               **
//***********************************

func (alerts *AlertStatusInternal) overCarbRule (a *ArgsMap) {
    var carbonThreshold  float64 = 0 // (inclusive good value)
    var allot float64 = 0 //value of alloted
    var soldCred float64 = 0 //value for soldCredits from the ledger
    var boughtCred float64 = 0 //value for boughtCredits from the ledger
    tbytes, found := getObject(*a, "reading")
    if found {
        carbonUsed, found := tbytes.(float64)
        if found { 
            //checking for the threshold that needs to match
            tbytes, found := getObject(*a, "threshold")
            if found {
                carbonThreshold, found = tbytes.(float64)
                if found {
                    //getting the value from attributes
                    tbytes, found = getObject(*a, "allottedCredits")
                    allot, found = tbytes.(float64)
                    tbytes, found = getObject(*a, "soldCredits")
                    soldCred, found = tbytes.(float64)
                    tbytes, found = getObject(*a, "boughtCredits")
                    boughtCred, found = tbytes.(float64)
                    //calculating total number of credits used
                    carbonUsed = carbonUsed + soldCred
                    //carbon threshold calculated
                    carbonThreshold = (carbonThreshold/100.0) * (allot + boughtCred)
                    if carbonUsed >= carbonThreshold {
                        alerts.raiseAlert(AlertsOVERCARBON)
                        return 
                    }
                } else {
                    //error should be thrown if threshold was not found
                    log.Errorf("Alerts could not be generated, threshold value is not correct type")
                    return 
                }
            } else {
                //error should be thrown if threshold was not found
                log.Errorf("Alerts could not be generated, threshold value was not found")
                return 
            } 
        } else {
            //error should be thrown if threshold was not found
            log.Errorf("Alerts could not be generated, carbon reading accumulated value is not correct type")
            return 
        }
    } else {
        //error should be thrown if threshold was not found
        log.Errorf("Alerts could not be generated, carbon reading accumulated value was not found")
        return
    }
    alerts.clearAlert(AlertsOVERCARBON)
    return
}

//***********************************
//**         COMPLIANCE            **
//***********************************

func (alerts *AlertStatusInternal) calculateContractCompliance (a *ArgsMap) (bool) {
    // a simplistic calculation for this particular contract, but has access
    // to the entire state object and can thus have at it
    // compliant is no alerts active
    return alerts.NoAlertsActive()
    // NOTE: There could still a "cleared" alert, so don't go
    //       deleting the alerts from the ledger just on this status.
}
