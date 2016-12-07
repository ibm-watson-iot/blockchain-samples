# Track and Trace Contract Sample

This is a small and very simple track and trace application to follow surgical kits from creation through their use in a hospital or clinic.

The most interesting new feature in this contract is geo-fencing. This is the act of tracing the movement of a surgical kit and fencing it into its final destination, which is
a hospital campus as defined by a center geolocation and a radius in meters. 

Once the kit enters that zone, an alert is thrown if its location is ever reported outside
of the fenced area. The alert only clears when the surgical kit is moved back into the fenced area, or the fenced area is updated or removed from the kit's world state.

This contract is based upon the [IoT Contract Platform](http://github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform), and is meant to demonstrate some of the features that the platform provides with little to no effort.