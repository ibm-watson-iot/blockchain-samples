#IoT Blockchain - Map Utilities  
[`mapUtils.go`](../mapUtils.go "a set of mapping utilities to be used by contracts")
##Introduction
As mentioned in 
[the tutorial on smart contract patterns](TutorialSmartContractPatterns.md "in this folder"),
there are several API patterns that can work within the general JSON-RPC like infrastructure to which all contracts must adhere. This module offers help for those contracts that choose to use the CRUD pattern with the partial state event pattern and with maps as opposed to structs. 

> Note that structs don't really work with partial states unless protobufs are used to define the messaging payload schema. We explored that method and JSON Schema 4 had advantages that we could not ignore. For example, compiled language-specoific protobuf manipulation is not necessary, which alleviates a lot of upgrade cycles as contracts evolve.

##JSON-RPC and Case
Hyperledger has introduced a JSON-RPC 2.0 envelope in messages to the fabric (and thus to the contract) with protobuf definitions and formal adherence to the specification. 

This module's concern is with the JSON-RPC payload that is intended to execute as contract API. The JSON RPC specification formally defines JSON tags as **case sensitive**, but this sample contract is capable of working with case independent JSON tags. This feature is controlled at *compile time* and so be aware that there is a constant in mapUtils that must be changed if strict adherence to case dependence is desired.

>In some environments, flexibility with the case of tags might be useful. An environment in which end point applications are written in many different languages with differing levels of formal schema support can offer faster development in the more flexible mode. 

``` go
var CASESENSITIVEMODE bool = false 
``` 

This affects every access to a property and so can become a nightmare of guard code without a canonical implementation, as exists in this utility module. It is therefore unnecessary to write case sensitive code if these mapUtils functions are used for all access to the properties in events and states. 

###Examples of Relaxed Case Sensitivity
These are all the same identifiers
- `assetID`, `AssetID`, `AssetId`, `AsSeTiD`
- `averylongidentifierthathasnobusinessbeinghere`, `AVeryLongIdentifierThatHasNoBusinessBeingHere`

##Finding Objects in Events and State
Events and state travel as JSON objects on the wire, and so they are modeled in contract code using the generic `map[string]interface{}` type definition. There is a typed version of this with the less than perfect name `ArgsMap` that can act as a receiver. This is discussed more deeply in the [alerts module document](alerts.md "an alerts management mechanism").

Three functions do the bulk of the work in isolating the key matching away from mainline contract code:

###`findMatchingKey (objIn interface{}, key string) (string, bool)`   
This function takes a map object and returns the matching key, either sensitive or insensitive to case, based on the state of the constant showed previously. It only looks at the current layer of the map and thus does not do a recursive search. The actual algorithms will be obvious -- direct access in case sensitive mode since it is of course a map, and iterate over the keys in the map's first level, comparing lowercase versions of the incoming key and the map key to handle case insensitivity. 

It returns the found key so that functions in the mainline contract can always use the original key for access and storage and not change it back and forth, which results in the potential proliferation of phantom properties. Key names can also be used as part of a key for storage of ancillary asset data and so property names are best not changed once created. Note that this does not imply that property names, once written into state, have to all be of uniform case. That's only an issue if clients themselves use case sensitivity while the contract is case-insensitive. 

###`findObjectByKey (objIn interface{}, key string) (interface{}, bool)`   
This function will return an object with a key that matches the incoming key. It uses  findMatchingKey and so is case aware. It too looks only at the current level of the incoming map.

###`getObject (objIn interface{}, qname string) (interface{}, bool)`   
This function searches a map recursively, following a **qualified** key. Dot notation separates levels in the key, an example being `location.latitude`. 

In this TRADELANE contract the data model is mostly flat, but contracts can be arbitrarily complex, and the extension field can hold arbitrarily nested objects as application-managed state. And context matters for properties, an example being temperatures that can be tracked for several components and can thus appear in multiple locations in the state. 

The extension property can also be used to extend the contract with arbitrarily deep object nesting. Objects at each level could have different capitalization from different event sources, and so this function provides the independence to correctly find the target object.

####Example of getObject Used by a Rule
The rule engine has a very simple format -- each rule exists to service a specific alert. Since there is only one sample alert in the schema at this time (`OVERTEMP`), there is also only a single rule in engine, called `OverTempRule` as would be expected.

For this module's purpose, what matters is the mechanism for access to the required data. The aforementioned `getObject` function is used to accomplish that. The JSON tag for temperature in the schema is `temperature`, and it lives at the top level of the event object. That means that the qualified name to get the temperature is just `"temperature"`.

Although discussed in the [rules document](rules.go "a simple rule engine to set the current alert status"), it bears repeating here to be clear on how property access works.

``` go
    tbytes, found := getObject(*a, "temperature")
    if found {
        t, found := tbytes.(float64)
        if !found {
            alerts.clearAlert(Alerts_OVERTEMP)
        } else {
            if t > temperatureThreshold {
                alerts.raiseAlert(Alerts_OVERTEMP)
            } else {
                alerts.clearAlert(Alerts_OVERTEMP)
            }
        }
    }
```  

The getObject call returns a generic interface{} object and a boolean that signifies whether the temperature was found. The returned value must have its type asserted, which is pretty easy since we know that the temperature is always a JSON number on the wire, and that maps in Go to a `float64`. After asserting the `float64`, test against the threshold and raise or clear the alert. See the rules document for more information. 

If the rule looks at a nested property - a contrived example being that the package may not enter the southern hemisphere, the getObject call could look like:

``` go
    tbytes, found := getObject(*a, "location.latitude")
    if found {
        t, found := tbytes.(float64)
        if !found {
            alerts.clearAlert(Alerts_GOINGSOUTH)
            and so on ...
```

###`contains(arr interface{}, val interface{}) bool`
Contains does what the name suggests for arrays of built in JSON equivalent types: string, int, float64 and interface{}.

The incoming `arr` is searched for a matching `val` and a boolean is returned to signify whether it is found. Go does not implement this function.

###`deepMerge(srcIn interface{}, dstIn interface{}) (map[string]interface{})`
Since state is always an object with potentially many nested levels of contained objects, and event must have a matching structure, there exists a need to look for matching properties anywhere in the event object structure and to overwrite the equivalent property in the state. In case-insensitive mode, the utilities documented above really start to earn their keep.

Deep merge must match the exact key in both modes to avoid proliferating properties of the same spelling but different case when in case-insensitive mode. 

The basic algorithm is to crawl the source object, which is an event as partial state, descending all levels recursively. For each leaf node in every level, the equivalent property is replaced in or copied into the destination object, which obviously must share a schema.

When the property being copied is an array, the two arrays are unioned. 

###`prettyPrint(m interface{}) (string)`
This function returns a string with a representation of whatever is passed in. A JSON MarshalIndent returned if successful, and if not then a simple SPrintf with %#v is returned. This is very useful for debugging.