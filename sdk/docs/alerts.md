# Alerts  
[`alerts.go`](../alerts.go "the contract alerts processing")

##Definition and Extension
The alerts module manages contract alert names and processing. Alerts are designated by short names such as `OVERTEMP` and are defined in a constant array as shown:

``` go
// Alerts
type Alerts int32

const (
    // AlertsOVERTEMP the over temperature alert 
    AlertsOVERTEMP    Alerts = 0

    // AlertsSIZE is to be maintained always as 1 greater than the last alert, giving a size  
	AlertsSIZE        Alerts = 1
)
```

It is common to assign the value `iota` so that enumerations are numbered automatically, but in this enumeration we use explicit numbering so that an alert can become deprecated but remain in the contract for historical purposes. 

Alerts_SIZE is the final entry and its index **does** change in order to provide an easily accessed value for the size needed to allocate a string array to hold the maximum possible alerts. 

EXAMPLE: You want to add UNDERTEMP and remove OVERTEMP (I know it makes no sense, but such is the nature of contrived examples), here are the wrong and then the right solutions:

###WRONG

``` go
type Alerts int32

const (
    // AlertsUNDERTEMP the over temperature alert 
    AlertsUNDERTEMP    Alerts = 0

    // AlertsSIZE is to be maintained always as 1 greater than the last alert, giving a size  
	AlertsSIZE        Alerts = 1
)
```

Very bad, as the algorithm for calculating the alerts status and thresholds for raised and cleared uses these values explicitly as array indexes.

###RIGHT

``` go
type Alerts int32

const (
    // AlertsOVERTEMP the over temperature alert 
    AlertsOVERTEMPDeprecated    Alerts = 0
    AlertsUNDERTEMP             Alerts = 1

    // AlertsSIZE is to be maintained always as 1 greater than the last alert, giving a size  
	AlertsSIZE                  Alerts = 2
)
```
These alerts are supported by protobuf-like functions to translate back and forth between names and values. Stringification is also supported.  

``` go
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
```

In a future version of this contract, these definitions may be generated from the schema, protobuf-style. Meanwhile, following this pattern makes adding new alerts a simple task.

##Raising and Clearing Alerts
The rules engine runs the individual rules and each of those is responsible to **ALWAYS** raise or clear the associated alert(s). It is extremely important that the rules have no previous knowledge of alert status, as that is the job of this module. The job of the rules is to simply state whether conditions exist for an active alert (raise it) or not (clear it).

The external view of alerts is defined in the schema and reflected in this module by the following declaration:

``` go
type AlertNameArray []string
type AlertStatus struct {
    Active  AlertNameArray  `json:"active"`
    Raised  AlertNameArray  `json:"raised"`
    Cleared AlertNameArray  `json:"cleared"`
}
```

Applications will see \["OVERTEMP"\] as the data in any of these if the alert is active, has been raised by this specific event, or has been cleared by this specific event. Simple monitoring UIs can display the alert values firectly from the JSON without translating it to some other string and application need not maintain a language-specific copy of the alerts constants or enum. 

This format is not, however, a convenient structure to process for the raising and clearing of events. For that, the alerts are converted to an internal format with boolean arrays instead of name arrays:

``` go
type AlertArrayInternal [Alerts_SIZE]bool
type AlertStatusInternal struct {
    Active  AlertArrayInternal
    Raised  AlertArrayInternal
    Cleared AlertArrayInternal
}
```

When looking at a single asset state, which defines the asset's status at any given point in time, it is not clear just by looking at the active alert status when the alert was raised or cleared. I.e. is this state just a continuation of a problem that happened earlier, or is this state the moment when the alert was raised? Similarly, if the alert is not active in this state, when was the alert actually cleared?

A monitor application could of course poll asset state constantly, maintaining the state of all alerts against all assets so that the moments of transition are easily identified. But this is a complex feature with significant potential for either missing some data or dragging the blockchain performance down by constant polling.

Instead, the extra information is embedded in the asset's state as the two rased and cleared arrays so that an application can ask for a specific historical period of states from the contract and process the returned list of states to find the moments of transition. 

This is possible by looking at individual states without writing algorithms that remember previous states etc because the contract knows the previous alert status and isolates the single raise or clear moment in the alertstatus at the one state transition where it occured.

An example: say that OVERTEMP occured at some point along a journey, and the frozen goods are spoilt upon arrival. A monitor application can perform a some lightweight analytics by requesting the historical states for all or part of the journey of the package or container, and then plot the moments when the temperature went out of spec and the time it stayed that way. Since the carrier property is remembered in every state since it was set upon assumption of the transport responsibility, finding the carrier that is responsible for the spoilage is an easy matter. Graphing the issue in a rudimentary fashion could look like:

|State->|1|2|3|4|5|6|7|8|9|10|
|-------|---|---|---|---|---|---|---|---|---|---|
|Active|.|.|.|T|T|T|.|.|.|.|
|Raised|.|.|.|T|.|.|.|.|.|.|
|Cleared|.|.|.|.|.|.|T|.|.|.|
|Carrier|1|1|2|2|2|2|2|3|3|3|

Carrier 2 is responsible and may be allocated penalties.

##Transition Processing
The functions for raising and clearing alerts receive the internal representation of the alert status and a single alert to be raised or cleared respectively. 

The calculation is straightforward, testing teh current active status and setting the three status values to match the appropriate thresholds.

For example, if the alert was already active, and we are in raiseAlert, then the active status is set to true, the raised status is set to false (because it was already active so this event could not be the transition point) and the cleared status is set to false (because an active alert could not possibly have been cleared).

This method is used for all four possible combinations, which are:

|Previous State||raiseAlert|||clearAlert||
|:-------:|:---:|:---:|:---:|:---:|:---:|:---:|
||active|raised|cleared|active|raised|cleared|
|Active|T|.|.|.|.|T|
|Inactive|T|T|.|.|.|.|

There is no need to change any of the code in the alerts module when extending the contract, only additional alert names need be added along with the changes described at the top of this document. The threshold tests, however, need to be implemented in the rules module, strictly following the pattern established in that file and in earlier comments this in document.