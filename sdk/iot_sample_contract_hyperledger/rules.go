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

import (
    "errors"
)

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool, error) {
    log.Debugf("Executing rules input: %+v", *alerts)
    // transform external to internal for easy alert status processing
    var internal = (*alerts).asAlertStatusInternal()

    // ------ validation rules
    // rule 1 -- test validation failure
    err := internal.testValidationRule(a)
    // return value is not used, return true, which means noncompliant
    if err != nil {return true, err}
    // rule 2 -- ???

    // ------ alert rules
    // rule 1 -- overtemp
    err = internal.overTempRule(a)
    if err != nil {return true, err}
    // rule 2 -- ???

    // transform for external consumption
    *alerts = internal.asAlertStatus()
    log.Debugf("Executing rules output: %+v", *alerts)

    // set compliance true means out of compliance
    compliant, err := internal.calculateContractCompliance(a) 
    if err != nil {return true, err} 
    // returns true if anything at all is active (i.e. NOT compliant)
    return !compliant, nil
}

//***********************************
//**     VALIDATION RULES          **
//***********************************

func (alerts *AlertStatusInternal) testValidationRule (a *ArgsMap) error {
    tbytes, found := getObject(*a, "testValidation")
    if found {
        t, found := tbytes.(bool)
        if found {
            if t {
                err := errors.New("testValidation property found and is true") 
                return err
            }
        }
    }
    return nil
}

//***********************************
//**        ALERT RULES            **
//***********************************

func (alerts *AlertStatusInternal) overTempRule (a *ArgsMap) error {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
                return nil
            }
        } else {
            log.Warning("overTempRule: temperature not type JSON Number, alert status not changed")
            // do nothing to the alerts status
            return nil 
        }
    }
    alerts.clearAlert(AlertsOVERTEMP)
    return nil
}

//***********************************
//**         COMPLIANCE            **
//***********************************

func (alerts *AlertStatusInternal) calculateContractCompliance (a *ArgsMap) (bool, error) {
    // a simplistic calculation for this particular contract, but has access
    // to the entire state object and can thus have at it
    // compliant is no alerts active
    return alerts.NoAlertsActive(), nil
    // NOTE: There could still a "cleared" alert, so don't go
    //       deleting the alerts from the ledger just on this status.
}