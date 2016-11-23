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

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Rule stores a route for an asset class or event
type Rule struct {
	RuleName string
	Alerts   []AlertName
	Class    AssetClass
	Function func(stub shim.ChaincodeStubInterface, asset *Asset) error
}

// RuleFunc is the signature for all rule functions
type RuleFunc func(stub shim.ChaincodeStubInterface, asset *Asset) error

var rulerouter = make(map[AssetClass][]Rule, 0)
var compliancerouter = make(map[AssetClass]Rule, 0)

func findRule(c AssetClass, ruleName string) (Rule, bool) {
	rules, found := rulerouter[c]
	if !found || len(rules) == 0 {
		return Rule{}, false
	}
	for _, r := range rules {
		if r.RuleName == ruleName {
			return r, true
		}
	}
	return Rule{}, false
}

func classRules(c AssetClass) []Rule {
	r := rulerouter[c]
	if r == nil {
		return []Rule{}
	}
	return r
}

// AddRule allows a class register its rules, one at a time. Note that
// rules should be added in the required order of execution.
// functionName is the function that will appear in a rest or gRPC message
// alerts is a list of strings defining the alerts that the rule will manipulate
// class is the asset class that registered the route
// rule is the function to be executed when the rulerouter is triggered
func AddRule(ruleName string, class AssetClass, alerts []AlertName, rule RuleFunc) error {
	r, found := findRule(class, ruleName)
	if found {
		err := fmt.Errorf("AddRule: rule name %s attempt to register against class %s for alerts [%v] but is already registered against class %s for alerts %v",
			ruleName, class.Name, alerts, r.Class.Name, r.Alerts)
		log.Error(err)
		return err
	}
	r = Rule{
		RuleName: ruleName,
		Alerts:   alerts,
		Class:    class,
		Function: rule,
	}
	rulerouter[class] = append(rulerouter[class], r)
	log.Debugf("Class %s added rule %s with alerts %v", r.Class.Name, r.RuleName, r.Alerts)
	return nil
}

// AddComplianceRule allows a class to register a custom algorithm that calculates whether
// this state is compliant. If there is no registered compliance rule, then compliance is set
// to false when any alert is active. In order to disable compliance checking, register the
// AlwaysCompliant rule.
func AddComplianceRule(class AssetClass, rule RuleFunc) error {
	r, found := compliancerouter[class]
	if found {
		err := fmt.Errorf("AddRule: compliance rule attempt to register against class %s but is already registered", class.Name)
		log.Error(err)
		return err
	}
	r = Rule{
		RuleName: "compliance",
		Alerts:   nil,
		Class:    class,
		Function: rule,
	}
	compliancerouter[class] = r
	log.Debugf("Class %s added compliance rule", r.Class.Name)
	return nil
}

// ExecuteRules executes all registered rules for the Asset's class
func (a *Asset) ExecuteRules(stub shim.ChaincodeStubInterface) error {
	log.Debugf("Executing rules input: %+v", a.AlertsActive)
	rules := classRules(a.Class)
	for _, rule := range rules {
		err := rule.Function(stub, a)
		if err != nil {
			err := fmt.Errorf("Rule (%v) failed with error %s", rule, err)
			log.Error(err)
			return err
		}
	}
	crule, found := compliancerouter[a.Class]
	if found {
		err := crule.Function(stub, a)
		if err != nil {
			err := fmt.Errorf("Compliance rule for class %s failed with error %s", a.Class, err)
			log.Error(err)
			return err
		}
		return nil
	}
	// default when no compliance rule is registered
	a.Compliant = len(a.AlertsActive) == 0
	return nil
}
