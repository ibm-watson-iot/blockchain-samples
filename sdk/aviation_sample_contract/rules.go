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
// KL 29 Aug 2016 Add HARDLANDING rule for aviation v4.4
// ************************************

package main

import (
	//"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func (state *ArgsMap) executeRules(stub *shim.ChaincodeStub, eventName string, alerts *AlertStatus, event ArgsMap) (bool, error) {
	log.Debugf("Executing rules input: %+v", *alerts)
	// transform external to internal for easy alert status processing
	var internal = (*alerts).asAlertStatusInternal()

	internal.clearRaisedAndClearedStatus()

	dynamicConfig, err := getDynamicConfigFromLedger(stub)
	if err != nil {
		return true, err
	}

	// ------ validation and state machine rules

	// ------ alert rules
	// rule 1 -- inspections to clear alerts acheck, bcheck, hardlanding
	err = state.inspectionsRule(dynamicConfig, &internal, event)
	if err != nil {
		return true, err
	}
	// rule 2 -- short cycle count acheck
	err = state.acheckRule(dynamicConfig, &internal, event)
	if err != nil {
		return true, err
	}
	// rule 3 -- long cycle count bcheck
	err = state.bcheckRule(dynamicConfig, &internal, event)
	if err != nil {
		return true, err
	}
	// rule 4 -- hard landings check
	err = state.hardlandingRule(dynamicConfig, &internal, event)
	if err != nil {
		return true, err
	}

	// transform for external consumption
	*alerts = internal.asAlertStatus()
	log.Debugf("Executing rules output: %+v", *alerts)

	// set compliance true means out of compliance
	compliant, err := state.calculateContractCompliance(&internal, event)
	if err != nil {
		return true, err
	}
	// returns true if anything at all is active (i.e. NOT compliant)
	// TODO improve on this
	return !compliant, nil
}

//****************************************
//**        VALIDATION RULES            **
//****************************************

//***********************************
//**        ALERT RULES            **
//***********************************

// Inspection actions are processed here:
// ACHECK -- clears the ACHECK (short cycle count) alert
// BCHECK -- clears the BCHECK (long cycle count) alert, also clears ACHECK
// HARDLANDING -- clears the hardlanding inspection alert
func (state *ArgsMap) inspectionsRule(config DynamicContractConfig, alerts *AlertStatusInternal, event ArgsMap) error {
	if _, found := getObject(event, "inspection"); !found {
		// this is an inspections rule, and this is not an inspection event
		return nil
	}
	if _, found := getObject(*state, "assembly"); !found {
		// inspections are on assemblies only
		return nil
	}
	insp, found := getObjectAsString(event, "inspection.action")
	if found {
		if insp == "ACHECK" {
			// clears acheck and resets
			_, ok := putObject(*state, "aCheckCounter", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to aCheckCounter for state %+v", state)
			}
			_, ok = putObject(*state, "aCheckCounterAdjusted", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to aCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsACHECK)
		} else if insp == "BCHECK" {
			// clears acheck and bcheck and resets both
			_, ok := putObject(*state, "bCheckCounter", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to bCheckCounter for state %+v", state)
			}
			_, ok = putObject(*state, "bCheckCounterAdjusted", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to bCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsBCHECK)
			_, ok = putObject(*state, "aCheckCounter", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to aCheckCounter for state %+v", state)
			}
			_, ok = putObject(*state, "aCheckCounterAdjusted", 0)
			if !ok {
				return fmt.Errorf("inspection rule: cannot put 0 to aCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsACHECK)
		} else if insp == "HARDLANDING" {
			// clears hardlanding, nothing to reset
			alerts.clearAlert(AlertsHARDLANDING)
		}
	}
	return nil
}

// ACHECK alert handled by this rule.
func (state *ArgsMap) acheckRule(config DynamicContractConfig, alerts *AlertStatusInternal, event ArgsMap) error {
	//log.Debugf("\n\n**** ACHECKRULE\n\nEVENT: %T || %+v\n\nSTATE: %T || %+v", event, event, state, state)

	_, flightFound := getObject(event, "flight")
	_, analyticAdjustmentFound := getObject(event, "analyticAdjustment")

	if !flightFound && !analyticAdjustmentFound {
		// neither flight nor analyticAdjustment event
		return nil
	}

	if _, found := getObject(*state, "assembly"); !found {
		// acheck on assemblies only
		return nil
	}

	// use the adjusted counter for threshold check
	accadjusted, accadjustedfound := getObjectAsNumber(*state, "aCheckCounterAdjusted")
	if accadjustedfound && accadjusted >= config.ACheckThreshold {
		alerts.raiseAlert(AlertsACHECK)
	}

	return nil
}

// BCHECK alert handled by this rule.
func (state *ArgsMap) bcheckRule(config DynamicContractConfig, alerts *AlertStatusInternal, event ArgsMap) error {

	_, flightFound := getObject(event, "flight")
	_, analyticAdjustmentFound := getObject(event, "analyticAdjustment")

	if !flightFound && !analyticAdjustmentFound {
		// neither flight nor analyticAdjustment event
		return nil
	}

	if _, found := getObject(*state, "assembly"); !found {
		// acheck on assemblies only
		return nil
	}

	// use the adjusted counter for threshold check
	bccadjusted, bccadjustedfound := getObjectAsNumber(*state, "bCheckCounterAdjusted")
	if bccadjustedfound && bccadjusted >= config.BCheckThreshold {
		alerts.raiseAlert(AlertsBCHECK)
	}

	return nil
}

// HARDLANDING alert handled by this rule:
func (state *ArgsMap) hardlandingRule(config DynamicContractConfig, alerts *AlertStatusInternal, event ArgsMap) error {
	// alert raised on 2nd consecutive hard landing for testing

	if _, found := getObject(*state, "assembly"); !found {
		// hard landing alerts on assemblies only
		return nil
	}

	hlevent, hlfound := getObjectAsBoolean(event, "flight.hardlanding")
	ahlevent, ahlfound := getObjectAsBoolean(event, "flight.analyticHardlanding")
	if (hlfound && hlevent) || (ahlfound && ahlevent) {
		// it was definitely a hard landing
		alerts.raiseAlert(AlertsHARDLANDING)
	}

	return nil
}

//***********************************
//**         COMPLIANCE            **
//***********************************

func (state *ArgsMap) calculateContractCompliance(alerts *AlertStatusInternal, event ArgsMap) (bool, error) {
	// a simplistic calculation for this particular contract, but has access
	// to the entire state object and can thus have at it
	// compliant is no alerts active
	return alerts.NoAlertsActive(), nil
	// NOTE: There could still a "cleared" alert, so don't go
	//       deleting the alerts from the ledger just on this status.
}
