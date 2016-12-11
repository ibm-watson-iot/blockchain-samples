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

// v0.1 KL -- new iot chaincode platform

package iotcontractplatform

import "sort"

// AlertNameArray is a string that represents an alert
type AlertNameArray []AlertName

// AlertName is a string that represents an alert
type AlertName string

// RaiseAlert adds an alertname to the active alerts array
func RaiseAlert(a *Asset, alert AlertName) {
	if a.AlertsActive == nil {
		a.AlertsActive = make(AlertNameArray, 0)
		a.AlertsActive = append(a.AlertsActive, alert)
	} else if !Contains(a.AlertsActive, alert) {
		a.AlertsActive = append(a.AlertsActive, alert)
	}
	sort.Sort(a.AlertsActive)
	return
}

// ClearAlert removes an alertname from the active alerts array
func ClearAlert(a *Asset, alert AlertName) {
	posn := -1
	for i, a := range a.AlertsActive {
		if a == alert {
			posn = i
			break
		}
	}
	if posn >= 0 {
		a.AlertsActive[posn] = a.AlertsActive[len(a.AlertsActive)-1]
		a.AlertsActive = a.AlertsActive[:len(a.AlertsActive)-1]
	}
	sort.Sort(a.AlertsActive)
	return
}

// GetAlertsAndDeltas takes two alert name arrays and returns a map with "raised" and "cleared" lists
func GetAlertsAndDeltas(alertsInOld AlertNameArray, alertsInNew AlertNameArray) map[string]interface{} {
	deltas := make(map[string]interface{})
	raised := AlertNameArray{}
	cleared := AlertNameArray{}
	for _, alert := range alertsInOld {
		if !Contains(alertsInNew, alert) {
			cleared = append(cleared, alert)
		}
	}
	for _, alert := range alertsInNew {
		if !Contains(alertsInOld, alert) {
			raised = append(raised, alert)
		}
	}
	if len(raised) > 0 {
		deltas["alertsRaised"] = raised
	}
	if len(cleared) > 0 {
		deltas["alertsCleared"] = cleared
	}
	if len(alertsInNew) > 0 {
		deltas["activeAlerts"] = alertsInNew
	}
	if len(deltas) > 0 {
		return deltas
	}
	return nil
}

func (aa AlertNameArray) Len() int           { return len(aa) }
func (aa AlertNameArray) Swap(i, j int)      { aa[i], aa[j] = aa[j], aa[i] }
func (aa AlertNameArray) Less(i, j int) bool { return aa[i] < aa[j] }
