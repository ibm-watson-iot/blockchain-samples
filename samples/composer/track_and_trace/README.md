# Track and Trace Network

This business network defines a contract between a shipper and a hospital regarding surgical kit shipment. The contract sets a maximum force and tilt value that a surgical kit can withstand during shipment. Kits that exceed the maximum force or tilt value are not taken by the hospital. 

This business network defines:

**Participants**
`Shipper` `Hospital`

**Assets**
`Contract` `SurgicalKit`

**Transactions**
`TiltReading` `ForceReading` `ShipmentReceived` `KitMovement`
