# How This Flow Works

In the config node, enter as many peer urls as you would like tracked. The
flow is set to run once, but you can set it to poll of course. Use periods 
of at least 1m, since this flow can timeout on each invocation of get block.

The flow sends the same command to all peers in parallel and checks every 2s
to see if all responses have arrived. It loops until 30s has passed and then
stops. If all peers respond, the flow continues.

If the chain heights are all the same, the hash comparison will proceed. Else 
it will print a message that it cannot proceed. Try again.

If the hash comparison shows divergence, the subflow for binary search is
launched. It will find the point of divergence quickly and will log results 
to the console. The chaincode ID and payload are Base64 decoded in the output.
Output is sorted by peer, using the naming convention vpx, where x if the zero
based position in your initial config array.

> Note: Watch your node-red console for progress reports and the final result.
The result can be hooked up to email etc in the usual node-red way, but I use
the console as well for convenience and debugging. Note that it is not possible
to debug with the usual debug panel in subflows, and most of this solution is
implemented in subflows. Hence, the heavy use of the console log.
