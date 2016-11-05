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
// KL 02 Nov 2016 Genericize and move to ctalerts package
//                Routing by asset class
// ************************************

package ctalerts

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	as "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctasset"
	cf "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctconfig"
	st "github.com/ibm-watson-iot/blockchain-samples/iotbase/ctstate"
	"github.com/op/go-logging"
)

// Logger for the ctrules package
var log = logging.MustGetLogger("cmpl")

// router
type AssetRule func (a* AssetClass) rule ([]string, err)
type AssetRuleArray []AssetRule

var ruleRouter map[AssetClass]AssetRuleArray = make(map[string]AssetRuleArray, 0)

// AddRule
func (c AssetClass) AddRule

// ExecuteRules executes all registered rules, returning true is there is an active alert
func ExecuteRules(stub *shim.ChaincodeStub, a as.Asset) (map[string]interface{}, bool, error) {
	log.Debugf("Executing rules input: %+v", *alerts)
	// transform external to internal for easy alert status processing
	var internal = (*alerts).asAlertStatusInternal()
	internal.clearRaisedAndClearedStatus()
	dynamicConfig, err := cf.GetDynamicConfigFromLedger(stub)
	if err != nil {
		return nil, true, err
	}
	mstate, ok := st.AsMap(state)
	if !ok {
		return mstate, false, fmt.Errorf("ExecuteRules: state is not a map shape %+v", state)
	}
	// ------ validation and state machine rules

	// ------ alert rules
	// rule 1 -- inspections to clear alerts acheck, bcheck, hardlanding
	mstate, err = inspectionsRule(dynamicConfig, &internal, mstate, event)
	if err != nil {
		return nil, false, err
	}
	// rule 2 -- short cycle count acheck
	mstate, err = acheckRule(dynamicConfig, &internal, mstate, event)
	if err != nil {
		return nil, false, err
	}
	// rule 3 -- long cycle count bcheck
	mstate, err = bcheckRule(dynamicConfig, &internal, mstate, event)
	if err != nil {
		return nil, false, err
	}
	// rule 4 -- hard landings check
	mstate, err = hardlandingRule(dynamicConfig, &internal, mstate, event)
	if err != nil {
		return nil, false, err
	}

	// transform for external consumption
	*alerts = internal.asAlertStatus()
	log.Debugf("Executing rules output: %+v", *alerts)

	// set compliance true means out of compliance
	compliant, err := calculateContractCompliance(&internal, mstate, event)
	if err != nil {
		return nil, false, err
	}
	// returns true if anything at all is active (i.e. NOT compliant)
	// TODO improve on this
	return mstate, !compliant, nil
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
func inspectionsRule(config cf.DynamicContractConfig, alerts *AlertStatusInternal, state map[string]interface{}, event interface{}) (map[string]interface{}, error) {
	if _, found := st.GetObject(event, "inspection"); !found {
		// this is an inspections rule, and this is not an inspection event
		return state, nil
	}
	if _, found := st.GetObject(state, "assembly"); !found {
		// inspections are on assemblies only
		return state, nil
	}
	insp, found := st.GetObjectAsString(event, "inspection.action")
	if found {
		if insp == "ACHECK" {
			// clears acheck and resets
			state, ok := st.PutObject(state, "aCheckCounter", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to aCheckCounter for state %+v", state)
			}
			state, ok = st.PutObject(state, "aCheckCounterAdjusted", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to aCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsACHECK)
		} else if insp == "BCHECK" {
			// clears acheck and bcheck and resets both
			state, ok := st.PutObject(state, "bCheckCounter", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to bCheckCounter for state %+v", state)
			}
			state, ok = st.PutObject(state, "bCheckCounterAdjusted", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to bCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsBCHECK)
			state, ok = st.PutObject(state, "aCheckCounter", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to aCheckCounter for state %+v", state)
			}
			state, ok = st.PutObject(state, "aCheckCounterAdjusted", 0)
			if !ok {
				return nil, fmt.Errorf("inspection rule: cannot put 0 to aCheckCounterAdjusted for state %+v", state)
			}
			alerts.clearAlert(AlertsACHECK)
		} else if insp == "HARDLANDING" {
			// clears hardlanding, nothing to reset
			alerts.clearAlert(AlertsHARDLANDING)
		}
	}
	return state, nil
}

// ACHECK alert handled by this rule.
func acheckRule(config cf.DynamicContractConfig, alerts *AlertStatusInternal, state map[string]interface{}, event interface{}) (map[string]interface{}, error) {
	//log.Debugf("\n\n**** ACHECKRULE\n\nEVENT: %T || %+v\n\nSTATE: %T || %+v", event, event, state, state)

	_, flightFound := st.GetObject(event, "flight")
	_, analyticAdjustmentFound := st.GetObject(event, "analyticAdjustment")

	if !flightFound && !analyticAdjustmentFound {
		// neither flight nor analyticAdjustment event
		return state, nil
	}

	if _, found := st.GetObject(state, "assembly"); !found {
		// acheck on assemblies only
		return state, nil
	}

	// use the adjusted counter for threshold check
	accadjusted, accadjustedfound := st.GetObjectAsNumber(state, "aCheckCounterAdjusted")
	if accadjustedfound && accadjusted >= config.ACheckThreshold {
		alerts.raiseAlert(AlertsACHECK)
	}

	return state, nil
}

// BCHECK alert handled by this rule.
func bcheckRule(config cf.DynamicContractConfig, alerts *AlertStatusInternal, state map[string]interface{}, event interface{}) (map[string]interface{}, error) {

	_, flightFound := st.GetObject(event, "flight")
	_, analyticAdjustmentFound := st.GetObject(event, "analyticAdjustment")

	if !flightFound && !analyticAdjustmentFound {
		// neither flight nor analyticAdjustment event
		return state, nil
	}

	if _, found := st.GetObject(state, "assembly"); !found {
		// acheck on assemblies only
		return state, nil
	}

	// use the adjusted counter for threshold check
	bccadjusted, bccadjustedfound := st.GetObjectAsNumber(state, "bCheckCounterAdjusted")
	if bccadjustedfound && bccadjusted >= config.BCheckThreshold {
		alerts.raiseAlert(AlertsBCHECK)
	}

	return state, nil
}

// HARDLANDING alert handled by this rule:
func hardlandingRule(config cf.DynamicContractConfig, alerts *AlertStatusInternal, state map[string]interface{}, event interface{}) (map[string]interface{}, error) {
	// alert raised on 2nd consecutive hard landing for testing

	if _, found := st.GetObject(state, "assembly"); !found {
		// hard landing alerts on assemblies only
		return state, nil
	}

	hlevent, hlfound := st.GetObjectAsBoolean(event, "flight.hardlanding")
	ahlevent, ahlfound := st.GetObjectAsBoolean(event, "flight.analyticHardlanding")
	if (hlfound && hlevent) || (ahlfound && ahlevent) {
		// it was definitely a hard landing
		alerts.raiseAlert(AlertsHARDLANDING)
	}

	return state, nil
}

//***********************************
//**         COMPLIANCE            **
//***********************************

func calculateContractCompliance(alerts *AlertStatusInternal, state map[string]interface{}, event interface{}) (bool, error) {
	// a simplistic calculation for this particular contract, but has access
	// to the entire state object and can thus have at it
	// compliant is no alerts active
	return alerts.NoAlertsActive(), nil
	// NOTE: There could still a "cleared" alert, so don't go
	//       deleting the alerts from the ledger just on this status.
}
