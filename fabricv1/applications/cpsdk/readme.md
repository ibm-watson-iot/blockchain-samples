# Contract Platform SDK

A work in progress.

Intent is to make it easy for applications to find a fabric and to work with it using transactions and queries. 

More advanced apps can create and join channels, and install and instantiate chaincode. 

This is derived from the Hyperledger e2e test suite. The order in which operations must be done is:

- create channel by asking a peer to do it
- join peer to channel
- install chaincode on a peer that needs to access a channel
- instantiate the chaincode in the channel
- send invokes and queries to the chaincode on a peer (which will use the channel on which it is instantiated)

