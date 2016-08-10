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
- additional events exist to record, well, events in the environment
    - flight events record a takeoff and landing sequence, which is known as a cycle
    - inspection events record an inspection against an assembly
    - analyticAdjustment events record calculated changes in life limits based on such analytics as weather patterns and runway conditions
    - maintenance events record the mounting and unmounting of assemblies on the aircraft (with unmounted assemblies being attached to an inventory asset)

- rules exist to track cycles and hard landings
    - ACHECK rule follows cycles against assemblies and when an assembly crosses the threshold, it raises the ACHECK inspection alert
    - BCHECK rule follows consecutive hard landings and when an assembly crosses the threshold for consecutive hard landings, it raises the BCHECK alert
    
- inspection events clear alerts via rules

> Note that the usual common properties such as geolocation, extension, etc. are available in the `common` subsection of event and state.

