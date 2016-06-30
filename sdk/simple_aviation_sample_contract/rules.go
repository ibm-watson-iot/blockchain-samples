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
// KL 28 Jun 2016 Remove OVERTEMP and add ACHECK and BCHECK rules for simple
//                aviation contract v4.2sa
// ************************************

package main

import (
    // "errors"
)

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool, error) {
    log.Debugf("Executing rules input: %+v", *alerts)
    // transform external to internal for easy alert status processing
    var internal = (*alerts).asAlertStatusInternal()

    internal.clearRaisedAndClearedStatus()

    // ------ alert rules
    // rule 1 -- inspections to clear alerts acheck, bcheck
    err := internal.inspectionsRule(a)
    if err != nil {return true, err}
    // rule 2 -- cycle counts acheck
    err = internal.cyclecountsRule(a)
    if err != nil {return true, err}
    // rule 3 -- hard landings bcheck
    err = internal.hardlandingsRule(a)
    if err != nil {return true, err}

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
//**        ALERT RULES            **
//***********************************

// There are two inspection types handled by this rule: 
// ACHECK -- clears the ACHECK (10 cycle counts) alert
// BCHECK -- clears the BCHECK (2 consecutive hard landings)
// NOTE: This rule is the only place in which an active ACHECK or BCHECK 
//       alert can be cleared.
func (alerts *AlertStatusInternal) inspectionsRule (a *ArgsMap) error {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    apbytes, found := getObject(*a, "airplane")
    if found {
        ap, found := apbytes.(map[string]interface{})
        if found {
            // first we deal with inspections to clear existing alerts
            insp, found := getObjectAsString(*a, "lastEvent.arg.inspection")
            if found {
                if insp == "ACHECK" {
                    // 1) ACHECK was performed, reset the cycle counter
                    ap["cyclecounter"] = 0
                    alerts.clearAlert(AlertsACHECK)
                } else if insp == "BCHECK" {
                    // 2) BCHECK was performed, reset the hard landing bit
                    ap["hardlanding"] = false
                    alerts.clearAlert(AlertsBCHECK)
                }
            }
        }
    }
    return nil
}

// There is one alert handled by this rule: 
// ACHECK -- if the state's cyclecounter has reached the threshold (10 cycle counts
//           by default), the ACHECK alert is raised.
// NOTE: This rule cannot clear an active ACHECK for any reason. Only an ACHECK
//       inspection can do that.
func (alerts *AlertStatusInternal) cyclecountsRule (a *ArgsMap) error {
    const acheckThreshold  float64 = 5 // alert raised on 5th cycle for testing

    apbytes, found := getObject(*a, "airplane")
    if found {
        ap, found := apbytes.(map[string]interface{})
        if found {
            tbytes, found := getObject(*a, "lastEvent.arg.flight")
            if found {
                _, found := tbytes.(map[string]interface{})
                if found {
                    // this is a flight event
                    ccount, found := getObjectAsNumber(ap, "cyclecounter")
                    if found {
                        ccount++
                    } else {
                        ccount = 1
                    }
                    if ccount >= acheckThreshold {
                        alerts.raiseAlert(AlertsACHECK)
                    }
                    // store it back
                    ap["cyclecounter"] = ccount 
                }
            }
        }
    }
    return nil
}

// There is one alert handled by this rule: 
// BCHECK -- if the state's hardlanding when another hardlanding is received
//           (two in a row), the BCHECK alert is raised.
// NOTE: This rule cannot clear an active BCHECK for any reason. Only a BCHECK
//       inspection can do that.
func (alerts *AlertStatusInternal) hardlandingsRule (a *ArgsMap) error {
    apbytes, found := getObject(*a, "airplane")
    if found {
        ap, found := apbytes.(map[string]interface{})
        if found {
            fbytes, found := getObject(*a, "lastEvent.arg.flight")
            if found {
                f, found := fbytes.(map[string]interface{})
                if found {
                    // this is a flight event
                    hlevent, found := getObjectAsBoolean(f, "hardlanding")
                    if found {
                        // hardlanding boolean present in incoming event
                        if hlevent {
                            // hard landing boolean is true
                            hlstate, found := getObjectAsBoolean(ap, "hardlanding")
                            if found && hlstate {
                                // have a previous hard landing in the airplane state
                                alerts.raiseAlert(AlertsBCHECK)
                            }
                        } 
                    } else {
                        // hard landing not present in event
                        hlevent = false
                    } 
                    // propagate event value to state
                    ap["hardlanding"] = hlevent
                }
            }
        }
    }
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