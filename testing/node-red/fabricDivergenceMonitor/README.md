# A Simple Blockchain Hash Divergence Monitor

## How This Flow Works

The four VPx nodes contain a URL property 
that you will have to adjust. It assumes 4 peers
because that is a minimum viable fabric, and Bluemix
creates fabrics with 4 peers as of October 2016.

The flow will poll your four peers every 15 seconds 
using the <URL:port>/chain REST endpoint. The 
responses from the peers are decoded and sent 
to two join nodes, each of which creates an array
of a specific property -- chain height on the top
join, and currentBlockHash on the bottom join.

Timeouts are set so that height processing slightly 
precedes block hash processing so that the result
of the height processing is joined with the result 
of the initial block hash join. 

If the chain heights are all the same, then the 
hash comparison will proceed. Else it will print
a message that it cannot proceed.

When hashes diverge, the outgoing message has the
property { "hashsame" : "N" }, which can be 
forwarded to another flow for further processing, 
stored in a database, or otherwise used to signal 
chain divergence.

This is useful for debugging the behavior of a
smart contract under load. Non-deterministic 
behaviors will lead to block hash divergence, 
which will permanently corrupt a blockchain.

## Configuring This Flow

Double click each validating peer HTML node and
add your peer URL for that specific node. You 
should be following the standard naming pattern
"vp[0,1,2,3]" so that your fabric logs and reports from
tools such as this one make sense.

## Caveats

This is not a complete solution, and it is not 
perfectly reliable. However, it reports incomplete
data sets and it handles timeouts relatively 
smoothly. Typical errors appear in the outgoing
messages so that a downstream flow can use only
known good data.