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
// KL 30 Jun 2016 Copy new function call to automatically maintain
//                raised and cleared flags. From aviation contract v4.2sa
// v4.5 KL October 2016 Ping Pong contract to demonstrate two-way communication with
//                 devices using the Hyperledger event infrastructure
// ************************************

package main

import (
	"errors"
)

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool, error) {
	log.Debugf("Executing rules input: %+v", *alerts)
	// transform external to internal for easy alert status processing
	var internal = (*alerts).asAlertStatusInternal()

	internal.clearRaisedAndClearedStatus()

	// ------ validation rules
	// rule 1 -- incoming assetID must be "PING" or "PONG"
	err := internal.isPingOrPongRule(a)
	// return true if noncompliant
	if err != nil {
		return true, err
	}
	// rule 2 -- ???

	// ------ alert rules
	// rule 1 -- count the pings
	err = internal.countPingsRule(a)
	if err != nil {
		return true, err
	}
	// rule 2 -- count the errors
	err = internal.countPongsRule(a)
	if err != nil {
		return true, err
	}

	// transform for external consumption
	*alerts = internal.asAlertStatus()
	log.Debugf("Executing rules output: %+v", *alerts)

	// set compliance true means out of compliance
	compliant, err := internal.calculateContractCompliance(a)
	if err != nil {
		return true, err
	}
	// returns true if anything at all is active (i.e. NOT compliant)
	return !compliant, nil
}

//***********************************
//**     VALIDATION RULES          **
//***********************************

func (alerts *AlertStatusInternal) isPingOrPongRule(a *ArgsMap) error {
	assetID, found := getObjectAsString(*a, "assetID")
	if found {
		if assetID != "PING" && assetID != "PONG" {
			err := errors.New("assetID illegal value, MUST be PING or PONG")
			return err
		}
	}
	return nil
}

//***********************************
//**        ALERT RULES            **
//***********************************

func (alerts *AlertStatusInternal) countPingsRule(a *ArgsMap) error {
	assetID, found := getObjectAsString(*a, "assetID")
	if found && assetID == "PING" {
		pingcount, found := getObjectAsInteger(*a, "pingcount")
		if found {
			_, ok := putObject(*a, "pingcount", pingcount+1)
			if !ok {
				err := errors.New("could not update ping count")
				return err
			}
		} else {
			_, ok := putObject(*a, "pingcount", 1)
			if !ok {
				err := errors.New("could not update ping count")
				return err
			}
		}
	}
	return nil
}

func (alerts *AlertStatusInternal) countPongsRule(a *ArgsMap) error {
	assetID, found := getObjectAsString(*a, "assetID")
	if found && assetID == "PONG" {
		pongcount, found := getObjectAsInteger(*a, "pongcount")
		if found {
			_, ok := putObject(*a, "pongcount", pongcount+1)
			if !ok {
				err := errors.New("could not update pong count")
				return err
			}
		} else {
			_, ok := putObject(*a, "pongcount", 1)
			if !ok {
				err := errors.New("could not update pong count")
				return err
			}
		}
	}
	return nil
}

//***********************************
//**         COMPLIANCE            **
//***********************************

func (alerts *AlertStatusInternal) calculateContractCompliance(a *ArgsMap) (bool, error) {
	// a simplistic calculation for this particular contract, but has access
	// to the entire state object and can thus have at it
	// compliant is no alerts active
	return alerts.NoAlertsActive(), nil
	// NOTE: There could still a "cleared" alert, so don't go
	//       deleting the alerts from the ledger just on this status.
}
