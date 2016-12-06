# IoT Contract Platform Package

> WORK IN PROGRESS

This platform exists to help a developer create very small IoT-flavoured smart contracts 
that automatically support features such as:

- asset classifiers with segregated world state to allow for "read all assets" for a given class
- tracked assets, incoming events, and outgoing events as separate concepts
- queryable world state history linked to transactions on the blockchain
- recent state changes across all assets
- filters and date ranges for browsing history and reading all assets
- rules and alerts
- schema-driven API that supports automated integration with our test platform (named the monitoring UI) and the Watson IoT Platform
- built in development tools for every contract, including "read world state", "delete world state"
- built in production tools for every contract, including "set logging level", "create new on update"

-----------------

## Include the Platform

In your `main.go` file, use these imports:

``` go
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	iot "github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform"
)
```

## Include a Command to Process / Generate your schema

The following include should work for most people who are developing a Hyperledger based contract inside the Vagrant environment.
See the documents in the [hyperledger v0.6 docs folder](https://github.com/hyperledger/fabric/tree/v0.6/docs/Setup), most specifically the 
`chaincode_setup.md` file, to setup your development environment and work inside the Vagrant environment with chaincode etc.

``` go
// Update the path to match your configuration
//go:generate go run /local-dev/src/github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform/scripts/processSchema.go
```

The dependencies for your contract are best dealt with by copying the [`vendor`](../iotcontractminimalsample/vendor) folder from the (minimal contract) so that golang tools 
will find everything they need to operate. This platform will be buried under your vendor folder and you will thus be able to reference it using the path shown
above, as your development folder should be on the `GOPATH`.

The `go build` command can be quite slow when using vendored packages. If you prefer higher speed builds, clone the iot contract platform to your machine and make
sure that it is in your local development directory, mapped inside vagrant as `/local-dev`. In order to build your contract and include the cloned platform, your local dev 
directory must be on the `GOPATH`, which it is not by default.

Further, the instructions in the Hyperledger documenst presume that the GOPATH contains one location, 
which is mapped inside the vagrant environment to `/opt/gopath`. This manifests as instructions like `cd $GOPATH`, which are
unfortunately not going to work once you start working against a local clones of included packages.

Consider getting used to using commands like `cd /opt/gopath` instead. This will allow you to adjust your `GOPATH` to something that works with locally
cloned packages, which should also be located under your `/local-dev`. Your own contracts should be on your `/local-dev` path,
and to set this all up correctly, see the instructions for setting up your local development environment, including the setting of the environment variable `LOCALDEVDIR`. 
so when you issue the following command, `go generate` and `go build` will work properly with local dependencies in `/local-dev`.

> You will know that you have the issue when you see an error like this:

``` bash
vagrant@hyperledger-devenv:v0.0.11-b111ac5:/local-dev/src/github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractminimalsample$ go generate

Included schema file not found on GOPATH: /opt/gopath: github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractplatform/schema/IOTCPschema.json
    -- please ensure that you have cloned or fetched the platform to your GOPATH
    -- also, please ensure that you add /local-dev to your GOPATH using the command
       'export GOPATH=/opt/gopath:/local-dev' and then run 'go generate' again

exit status 1
main.go:27: running "go": exit status 1
vagrant@hyperledger-devenv:v0.0.11-b111ac5:/local-dev/src/github.com/ibm-watson-iot/blockchain-samples/contracts/platform/iotcontractminimalsample$ 
```
More to follow ....