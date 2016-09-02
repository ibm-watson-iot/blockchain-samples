# Rules Engine

[`rules.go`](rules.go "the contract rules engine")

The rules engine in the contract template is small and compact but very easily expanded. The function `executeRules` is called by each of the contract's event processing functions that can change asset state (except of course for the deletions, which wipe out state completely).

The rules engine is meant to be a place where validation and alert calculations are performed. Validation rules typically test the quality or existence of data in the merged state and can return an error that will cause the entire transaction to fail and all of the state deltas that may have been calculated to be rolled back.

Alert calculations are meant to signal application-or contract specific status for an asset.

## When to Call the Rules Engine

Every time the state is changed, the rules engine must be called just before the state is put into the ledger. This is so that the rules see all of the state's data after the event has been processed, but before the state has been written. The alerts are injected into or removed from the state based on whether there is anything to see. For example, nothing raised and nothing active, but one alert cleared by this event means that the alerts will appear in the state to show what happened to the alert status.

Examples of state-changing functions in this pattern are `createAsset`, `updateAsset` and `deletePropertiesFromAsset`, but any added state-changing functions or setters would want to execute the rules engine at the appropriate point in the code.

## The Engine's Structure

The engine is simplicity itself. Each rule is called in order and the final result is recorded in the alerts structure that is passed by reference.

``` go

func (a *ArgsMap) executeRules(alerts *AlertStatus) (bool) {
    log.Debugf("Executing rules input: %v", *alerts)
    // rule 1 -- overtemp
    alerts.OverTempRule(a)
    // rule 2 -- ???
    log.Debugf("Executing rules output: %v", *alerts)

```

The state object in the form of the ArgsMap type is the receiver for this function. A generic map cannot be a receiver, which is why ArgsMap exists as a separate type. The generalized object / map shape applies to both events and states in this pattern. The create code sends in the original event because there is no existing state with which the event is merged. The update code would send in the merged state instead.

## Receivers versus Parameters

The rules engine has its receiver and parameter as pointers so that the engine can safely update the alerts in place. The engine returns the compliance status as a boolean value so that the calling code is compact.

Note that the receiver and parameter relationship is reversed in the rules, as they operate mainly on the alerts, but still need access to all properties in the final state. 

The reality is that the engine and rules do not really need receivers, as there is no real polymorphism involved, so they could also be implemented to simply accept both parameters as pointers. 

## Rules and State Changes

Since pointer receivers and parameters allow the rules engine efficient access to the entire state, and since any property could be in scope for the rules engine, there is no reason why a rule cannot exist to calculate another state property rather than raise or clear alerts.

There is in fact no rule (pun intended) that says that a rule cannot change the state. That might be the whole point of a specific rule in the first place.

And of course rules can now pass an error back to the calling event process to fail the whole transaction. Use this with caution of course. 

## Compliance

The state in this pattern contains a property called `compliant`, which documents whether the contract considers this specific asset to be in compliance with agreed-to rules (which of course are implemented in code in the rules engine).

``` go

    // set compliance true means out of compliance
    compliant := alerts.CalculateContractCompliance(a)

    // returns true if anything at all is active
    return !compliant

```

The sample compliance calculation `CalculateContractCompliance` determines whether there are any active alerts and returns true if there are **no** alerts active. This is in fact calculated in a single line:

``` go

    return alerts.NoAlertsActive()

```

But the rules engine **reverses** that value to notify the caller that there is an **out of compliance* asset. Looking again at the return statement, it returns true if the contract is *not* compliant, as in:

``` go

    return !compliant

```

This could go either way, but for now the contract calls the rules engine in an if statement and in the *if body* it handles non-compliance. The *else body* handles compliance.

## Writing a Rule

The basic rule pattern for an alert is:

- get the property or properties to be tested
    - if not found, clear the alert 
- assert the type
    - if the wrong type, clear the alert
- test the threshold or perform whatever other logic determines alert status
    - if good, clear the alert
    - if bad, raise the alert

And here is the sample alert for `OVERTEMP`:

``` go

func (alerts *AlertStatus) OverTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if !found {
            alerts.clearAlert(AlertsOVERTEMP)
        } else {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
            } else {
                alerts.clearAlert(AlertsOVERTEMP)
            }
        }
    }
}

```

In writing this document, I actually discovered a bug in that code. Do you see it?

Tick tock, tick tock ...

The bug is in the very first step. If temperature is found, the rule happens. If not, nothing happens. This violates the statement:

```

    - if not found, clear the alert 

```

So the correct rule implemention is:

``` go

func (alerts *AlertStatus) OverTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if !found {
            alerts.clearAlert(AlertsOVERTEMP)
        } else {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
            } else {
                alerts.clearAlert(AlertsOVERTEMP)
            }
        }
    } else {
        alerts.clearAlert(AlertsOVERTEMP)
    }
}

```

Had I not reversed the logic up top by handling !found first, this would have been more obvious. There is tension between early exit programming and nested programming. I prefer the former in most cases, but the latter works well in this case ... but preferably done this way:

``` go

func (alerts *AlertStatus) OverTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
            } else {
                alerts.clearAlert(AlertsOVERTEMP)
            }
        } else {
            alerts.clearAlert(AlertsOVERTEMP)
        }
    } else {
        alerts.clearAlert(AlertsOVERTEMP)
    }
}

```

Or how about the more compact version?

``` go

func (alerts *AlertStatus) OverTempRule (a *ArgsMap) {
    const temperatureThreshold  float64 = 0 // (inclusive good value)

    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if found {
            if t > temperatureThreshold {
                alerts.raiseAlert(AlertsOVERTEMP)
                return
            }
        }
    }
    alerts.clearAlert(AlertsOVERTEMP)
}

```

That's dead clear now and matches the pattern, now expressed the same way:

- get the property or properties to be tested
    - if found, assert the type
        - if the right type, test the threshold or perform whatever other logic determines alert status
        - return
- clear the alert
- return

If this all seems a bit basic, I'm just pointing out that programming style matters a lot when implementing such short bits of code, because the tendency to take the code in at a glance actually works against the details of getting it **exactly right**. It is worth writing the rules down in the pseudo-code style used for the pattern just to be sure that alerts are raised and cleared as necessary.

> And do remember that alerts must **ALWAYS** be raised or cleared by a rule with **no exceptions**.

## Why Raise and Clear Without Regard to Existing Alert Status?

By now, it will be obvious that the rule does not care what its alert status was at point of entry. This is by design. It simply has no bearing on the current alert status and thus should be left out. That knowledge belongs inside the alerts module. 

The demonstration of confusion in the shortest snippet of code above should make the rule writer leery about adding any potential complexity. And one look at how the 
alerts module calculates the active, raised and cleared states for each alert will tell you that proliferating that calculation out to all rules is a very bad idea.

A rule's job is to tell the alerts module whether an alert is active or not by **always** calling `raiseAlert(alert)` or `clearAlert(alert)` respectively. The alerts module then calculates the exact alert state and thresholds for that specific alert.

See the [`alerts module`](alerts.md "calculates active, raised and cleared status based on inputs from rules") for more information on how the calculation works, but *please make certain that every rule __raises or clears every alert with which it is concerned__ every time it is executed.* 