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
// Alerts package
// KL 16 Feb 2016 Initial alert package
// KL 22 Feb 2016 add AllClear method and associated constant
// KL 18 Apr 2016 Adapt new external JSON alerts (names instead of booleans) from orig 3.1/4.0
// ************************************

package main

import (
)

// Alerts exists so that strict type checking can be applied
type Alerts int32

const (
    // AlertsOVERTEMP the over temperature alert 
    AlertsOVERTEMP    Alerts = 0

    // AlertsSIZE is to be maintained always as 1 greater than the last alert, giving a size  
	AlertsSIZE        Alerts = 1
)

// AlertsName is a map of ID to name
var AlertsName = map[int]string{
	0: "OVERTEMP",
}

// AlertsValue is a map of name to ID
var AlertsValue = map[string]int32{
	"OVERTEMP": 0,
}

func (x Alerts) String() string {
	return AlertsName[int(x)]
}

// AlertArrayInternal is used to store the list of active, raised or cleared alerts
// for internal processing
type AlertArrayInternal [AlertsSIZE]bool
// AlertNameArray is used for external alerts in JSON
type AlertNameArray []string

// NOALERTSACTIVEINTERNAL is the zero value of an internal alerts array (bools)
var NOALERTSACTIVEINTERNAL = AlertArrayInternal{}
// NOALERTSACTIVE is the zero value of an external alerts array (string names)
var NOALERTSACTIVE = AlertNameArray{}

// AlertStatusInternal contains the three possible statuses for alerts
type AlertStatusInternal struct {
    Active  AlertArrayInternal  
    Raised  AlertArrayInternal  
    Cleared AlertArrayInternal  
}

type AlertStatus struct {
    Active  AlertNameArray  `json:"active"`
    Raised  AlertNameArray  `json:"raised"`
    Cleared AlertNameArray  `json:"cleared"`
}

// convert from external representation with slice of names
// to full length array of bools 
func (a *AlertStatus) asAlertStatusInternal() (AlertStatusInternal) {
    var aOut = AlertStatusInternal{}
    for i := range a.Active {
        aOut.Active[AlertsValue[a.Active[i]]] = true
    }
    for i := range a.Raised {
        aOut.Raised[AlertsValue[a.Raised[i]]] = true
    }
    for i := range a.Cleared {
        aOut.Cleared[AlertsValue[a.Cleared[i]]] = true
    }
    return aOut
}

// convert from internal representation with full length array of bools  
// to slice of names
func (a *AlertStatusInternal) asAlertStatus() (AlertStatus) {
    var aOut = newAlertStatus()
    for i := range a.Active {
        if a.Active[i] {
            aOut.Active = append(aOut.Active, AlertsName[i])
        }
    }
    for i := range a.Raised {
        if a.Raised[i] {
            aOut.Raised = append(aOut.Raised, AlertsName[i])
        }
    }
    for i := range a.Cleared {
        if a.Cleared[i] {
            aOut.Cleared = append(aOut.Cleared, AlertsName[i])
        }
    }
    return aOut
}

func (a *AlertStatusInternal) raiseAlert (alert Alerts) {
    if a.Active[alert] {
        // already raised
        // this is tricky, should not say this event raised an
        // active alarm, as it makes it much more difficult to track
        // the exact moments of transition
        a.Active[alert] = true
        a.Raised[alert] = false
        a.Cleared[alert] = false
    } else {
        // raising it
        a.Active[alert] = true
        a.Raised[alert] = true
        a.Cleared[alert] = false
    }
}

func (a *AlertStatusInternal) clearAlert (alert Alerts) {
    if a.Active[alert] {
        // clearing alert
        a.Active[alert] = false
        a.Raised[alert] = false
        a.Cleared[alert] = true
    } else {
        // was not active
        a.Active[alert] = false
        a.Raised[alert] = false
        // this is tricky, should not say this event cleared an
        // inactive alarm, as it makes it much more difficult to track
        //  the exact moments of transition
        a.Cleared[alert] = false
    }
}

func newAlertStatus() (AlertStatus) {
    var a AlertStatus
    a.Active = make([]string, 0, AlertsSIZE)
    a.Raised = make([]string, 0, AlertsSIZE)
    a.Cleared = make([]string, 0, AlertsSIZE)
    return a
}

func (a *AlertStatus) alertStatusFromMap (aMap map[string]interface{}) () {
    a.Active.copyFrom(aMap["active"].([]interface{}))
    a.Raised.copyFrom(aMap["raised"].([]interface{}))
    a.Cleared.copyFrom(aMap["cleared"].([]interface{}))
} 

func (arr *AlertNameArray) copyFrom (s []interface{}) {
    // a conversion like this must assert type at every level
    for i := 0; i < len(s); i++ {
        *arr = append(*arr, s[i].(string))
    }
}

// NoAlertsActive returns true when no alerts are active in the asset's status at this time
func (arr *AlertStatusInternal) NoAlertsActive() (bool) {
    return (arr.Active == NOALERTSACTIVEINTERNAL)
}

// AllClear returns true when no alerts are active, raised or cleared in the asset's status at this time
func (arr *AlertStatusInternal) AllClear() (bool) {
    return  (arr.Active == NOALERTSACTIVEINTERNAL) &&
            (arr.Raised == NOALERTSACTIVEINTERNAL) &&
            (arr.Cleared == NOALERTSACTIVEINTERNAL) 
}

// NoAlertsActive returns true when no alerts are active in the asset's status at this time
func (a *AlertStatus) NoAlertsActive() (bool) {
    return len(a.Active) == 0
}

// AllClear returns true when no alerts are active, raised or cleared in the asset's status at this time
func (a *AlertStatus) AllClear() (bool) {
    return  len(a.Active) == 0 &&
            len(a.Raised) == 0 &&
            len(a.Cleared) == 0 
}