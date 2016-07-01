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
   //"time"
   // "strconv"
    )

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool) {
    log.Debugf("Executing rules input: %v", *alerts)
    var internal = (*alerts).asAlertStatusInternal()

    // rule 1 -- Create and mod time check
    internal.timeCheck(a)
    // rule 2 --RPM check : if motor is running at 20% or below, it will likely overheat
    internal.rpmCheck(a)
    // rule 3 -- HVAC Check. If the HVAC is not running, that is an alert scenario
    //internal.hvacCheck(a)
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

func (alerts *AlertStatusInternal) timeCheck (a *ArgsMap) {
//var createTime time.Time
//var modTime time.Time
    /*
    now := time.Now()
    unixNano := now.UnixNano()                                                                      
    umillisec := unixNano / 1000000  */
    crTime, found := getObject(*a, "create_date")
    mdTime, found2 := getObject(*a, "last_mod_date")
    if found && found2 {
        //modTime= time.Unix(0, msInt*int64(time.Millisecond))
        if crTime.(float64) > mdTime.(float64) {
            alerts.raiseAlert(AlertsTIMEERROR)
        return
        }
        alerts.clearAlert(AlertsTIMEERROR)
    }
}
// Need to modify so that for motor, this ic called first 
func (alerts *AlertStatusInternal) rpmCheck (a *ArgsMap) {
//Reference : http://www.vfds.in/be-aware-of-vfd-running-in-low-speed-frequency-655982.html
    maxRPM, found := getObject(*a, "max_rpm")
    if found {
        curRPM, found2 := getObject(*a, "rpm")
        if found2 {
            percRPM := (curRPM.(float64)/maxRPM.(float64))*100
            if percRPM <=30 {
                alerts.raiseAlert(AlertsRPMERROR)
                return
            }
        }
    }
    alerts.clearAlert(AlertsRPMERROR)
}
/*
func (alerts *AlertStatusInternal) hvacCheck (a *ArgsMap) {
    hvacMode, found := getObject(*a, "hvac_mode")
    if found {
        tgtTemp, found2 := getObject(*a, "target_temperature_c")
        if found2 {
            ambTemp, found3 := getObject(*a, "ambient_temperature_c")
            if found3 {
                if (ambTemp.(float64) >tgtTemp.(float64) && hvacMode =="heat") {
                    alerts.raiseAlert(AlertsHVACOVERHEAT)
                    return
                }
                alerts.clearAlert(AlertsHVACOVERHEAT)
                if (ambTemp.(float64) <tgtTemp.(float64) && hvacMode =="cool") {
                    alerts.raiseAlert(AlertsHVACOVERCOOL)
                    return
                }
                alerts.clearAlert(AlertsHVACOVERCOOL)
            }
        }
    }
    alerts.clearAlert(AlertsHVACOVERHEAT)
    alerts.clearAlert(AlertsHVACOVERCOOL)
}
*/
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