#Aviation Sample Contract

> **This contract is derived from the [Generic IoT Sample Contract](../iot_sample_contract) to adapt it to an aviation scenario. 
 The generic IoT sample contract is described in a set of documents in the `docs` folder. See the [README](../docs/README.md) for more information. And please read the introduction and customization documents therein before proceeding through this code base.**

The aviation scenario is as follows:

- Three primary asset classes:
    - Airline -- owns zero to many aircraft
    - Aircraft -- encompasses zero ro many assemblies
    - Assembly -- landing gear, wing, etc. These are life limited parts.
- The CRUD features for each asset are abstracted in the module [assetCommon.go](./assetCommon.go), which makes heavy use of abstracted services in [crudCommon.go](./crudCommon.go)  
- each asset follows the *partial state as event* pattern, where the asset's writable properties make up its primary `event` to be passed to create and update
- additional events exist for:
    - flights that record a takeoff and landing sequence, which is known as a cycle
    - inspections against an assembly
    - analyticAdjustment events record calculated changes in wear and tear based on such analytics as weather patterns and runway conditions
    - maintenance events record the mounting and unmounting of assemblies with a full state machine

- rules exist to track cycles and hard landings
    - ACHECK rule compares adjusted (by analytics) cycle counters to a dynamically configurable threshold, raising an alert as necessary, typically short time
    - BCHECK rule compares adjusted (by analytics) cycle counters to a dynamically configurable threshold, raising an alert as necessary, typically long time
    - HARDLANDING rule raises an alert when a hard landing is known to have occured
    
- inspection events clear these alerts, note that bcheck clears both acheck and bcheck alerts

> Note that the usual common properties such as geolocation, extension, etc. are available in the `common` subsection of asset event and state.

Physical changes from the Generic IoT Contract include:

- Three assets tracked with full CRUD APIs: airline, aircraft, assembly
- Four event types handled with event APIs: flight, maintenance, inspection, analyticAdjustment
- 2-way inverted index maintains 1 to many relationship from aircraft to assemblies
- filters allow sophisticated queries and offer lightweight relationships between asset classes
    - a filter is a match type (all, any, none) and an array of k:v pairs with qualified property names and values
- contractConfig module supports static and dynamic configuration of contract
- new common layer for quick addition of a new asset class
- new common layer to support crud operations
- rules for acheck and bcheck (short and long term inspection cycles) and hard landing alerts 