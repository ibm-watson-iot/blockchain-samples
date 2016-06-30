#A Hyperledger Sample Contract with Simple Aviation Schema

> **This smart contract is described in a set of documents in the `/docs` folder. See the [README](./docs/README.md) for more information. And please read the introduction and customization documents therein before proceeding through the code base.**

This contract is derived from and contributed to the [Generic IoT Sample Contract](../iot_sample_contract_hyperledger) to adapt it to a simple aviation scenario. 

The scenario is as follows:

- A simple airplane object is the tracked asset
- using the *partial state as event* pattern, the `airplane` section of the event and state in the schema contains properties describing the airplane
- additionally, in the state object, `airplane` tracks a cyclecounter that is used to determine when to raise the ACHECK alert, which requires an inspection of the landing gear for wear on tires, etc
    - this counter does not appear in the `airplane` event, but rather it is incremented when a `flight` subevent is received
- the `flight` subevent carries only the flight number and a boolean property documenting whether there was a hard landing
    - two consecutive hard landings will raise the BCHECK alert, which requires an inspection of the landing gear for damage
- a second subevent called `inspection` can be sent at any time, indicating that the ACHECK or BCHECK inspection has been performed, which clears the associated alert
    - an inspection is the __*only way*__ to clear one of these alerts

> Note that the usual common properties such as geolocation, extension, etc. are available in the new iotCommon subsection of event and state.

Improvements over version 4.2 of the generic iot contract include:
- Store the JSON map of the argument in the lastEvent. This offers access to the incoming event for more complex rules. (See the simple aviation rules engine.)
    - Make deepcopy of args into stateOut in the `createAsset` function to avoid infinite loop when args are attached into lastEvent. 
- Store TXNUUID and timestamp in state only. It was polluting the lastEvent object for no good reason.
- Indicate position where setter or command subevents (if they are in use) are deleted from state, while leaving recorder subevents intact. 
    >- Partial state as event relies on the deepmerge of recorder style property sections. But in addition, setter or command sections can be sent to the contract through the single `updateAsset` API and then deleted from state just before committing it to the ledger. See the simple aviation contract for an example of this with hard landings and the cycle counter.
