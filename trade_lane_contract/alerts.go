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
// ************************************

package main // sitting beside the main file for now

import (
)

// Alerts
type Alerts int32

const (
    Alerts_OVERTEMP    Alerts = 0
	Alerts_SIZE        Alerts = 1 // maintained as 1 higher than the last entry for array sizing
)

var Alerts_name = map[int32]string{
	0: "OVERTEMP",
	1: "TBD",
}
var Alerts_value = map[string]int32{
	"OVERTEMP":       0,
	"TBD": 1,
}

func (x Alerts) String() string {
	return Alerts_name[int32(x)]
}

type AlertArray [Alerts_SIZE]bool

var NOALERTSACTIVE AlertArray = AlertArray{}

// Alerts is a struct with arrays indexed by the enum Alerts from protobuf
type AlertStatus struct {
    Active  AlertArray  `json:"active"`
    Raised  AlertArray  `json:"raised"`
    Cleared AlertArray  `json:"cleared"`
}

func (a *AlertStatus) raiseAlert (alert Alerts) {
    if a.Active[alert] {
        // already raised
        // this is tricky, should not say this event raised an
        // active alarm, as it makes it much more difficult to track
        //  the exact moments of transition
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

func (a *AlertStatus) clearAlert (alert Alerts) {
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

func (a *AlertStatus) alertStatusFromMap (aMap map[string]interface{}) () {
    a.Active.copyFrom(aMap["active"])
    a.Raised.copyFrom(aMap["raised"])
    a.Cleared.copyFrom(aMap["cleared"])
} 

// slice can never be smaller
func (arr *AlertArray) copyFrom (s interface{}) {
    // a conversion like this must assert type at every level
    slice := s.([]interface{})
    for i := 0; i < len(slice); i++ {
        if i > len(arr) {
            // what has probably happened is that someone did not properly
            // deprecate the use of an alert and just shrunk the list of 
            // alerts, which changes numbers in the contract's storage and
            // breaks compatibility. FIX IT!
            log.Warning("INCOMING STORED LIST OF ALERTS SHOULD NEVER BE LONGER THAN CURRENT LIST OF ALERTS!!!!")
            break
        }
        alert := slice[i].(bool)
        arr[i] = alert
    }
} 

func (arr *AlertStatus) NoAlertsActive() (bool) {
    return (arr.Active == NOALERTSACTIVE)
}

func (arr *AlertStatus) AllClear() (bool) {
    return  (arr.Active == NOALERTSACTIVE) &&
            (arr.Raised == NOALERTSACTIVE) &&
            (arr.Cleared == NOALERTSACTIVE) 
}