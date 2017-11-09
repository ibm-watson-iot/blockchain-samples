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
*/


// ************************************
// Rules for Contract
// KL 16 Feb 2016 Initial rules package for contract v2.8
// KL 22 Feb 2016 Add compliance calculation
// KL 09 Mar 2016 Logging replaces printf for v3.1
// KL 12 Mar 2016 Conversion to externally present as alert names
// KL 29 Mar 2016 Fixed subtle bug in OVERTEMP discovered while
//                documenting the rules engine for 3.0.5
// ************************************

package main

import ()

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool) {
    log.Debugf("Executing rules input: %v", *alerts)
    var internal = (*alerts).asAlertStatusInternal()

    // rule 1 -- overtemp
    internal.overTempRule(a)
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

func (alerts *AlertStatusInternal) overTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
                return
            }
        }
    }
    alerts.clearAlert(AlertsOVERTEMP)
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