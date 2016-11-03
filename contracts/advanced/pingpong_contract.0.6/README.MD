# PingPong Contract

> Updated for Hyperledger 0.6

This is a simple contract that has been derived from the IoT Sample contract. Please read the introductory articles in the [docs folder](../../../docs) if you are not familiar with how these contracts function.

This contract adds the ability to send outgoing events from an advanced IoT contract. Three events are registered, as clipped from the code:

``` go
// EVTINVOKEERR is sent out whenever there is an error in the chaincode
const EVTINVOKEERR = "EVTINVOKEERR"

// EVTPONG is sent out whenever a PING assetID is received
const EVTPONG = "EVTPONG"

// EVTPING is sent out whenever a PONG assetID is received
const EVTPING = "EVTPING"
```

See the [event listener application](../../../applcations/event_listener) README and code to understand how the client registers interest in specific events and then catches them in a gRPB stream (modeled in Go as channels).

