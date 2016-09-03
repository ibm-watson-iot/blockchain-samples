# Assets and Events

## Introduction

### **Asset**

The primary focus of an IoT smart contract is often referred to as an asset. Some tangible thing that tracks the state of something important. An example might be the warranty status on a car. The primary asset could be the warranty itself or the car. When tracking of only one asset is desired, as in a contract derived from the [`iot_sample_contract`](../iot_sample_contract), the choice must be made as to which asset is the one most likely to be referenced. 

With the example of a warranty contract, it is easy to imagine that the car's VIN will always be readily accessible when the owner or a concerned 3rd party such as an authorized repair center would like to check on the warranty status. It is obvious from a rudimentary discussion of the primary scenarios that the VIN is the right choice for the lookup key (which we call the `assetID`) and thus the car itself is the primary contract asset.

Once the primary asset is selected, the API is designed to manipulate that asset class. As discussed in the [Introduction to Hyperledger Smart Contracts for IoT, Best Practices and Patterns](HyperledgerContractsIntroBestPracticesPatterns.html) document, these samples use a pattern based on a simplified CRUD (create, read update, delete) API for manipulating assets. The IoT sample, for example, has this delegation for its asset CRUD functions:

``` go

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if function == "createAsset" {
		return t.createAsset(stub, args)
	} else if function == "updateAsset" {
		return t.updateAsset(stub, args)
	} else if function == "deleteAsset" {
		return t.deleteAsset(stub, args)
	} else if function == "deleteAllAssets" {
		return t.deleteAllAssets(stub, args)
	} else if function == "deletePropertiesFromAsset" {
		return t.deletePropertiesFromAsset(stub, args)
	} else if function == "setLoggingLevel" {
		return nil, t.setLoggingLevel(stub, args)
	} else if function == "setCreateOnUpdate" {
		return nil, t.setCreateOnUpdate(stub, args)
	}
	err := fmt.Errorf("Invoke received unknown invocation: %s", function)
    log.Warning(err)
	return nil, err
}

```

There is little reason to change the crud pattern described above to more specific asset class names (e.g. warranty, automobile) in a single asset contract, which we often call these recorder contracts as they are mainly concerned with recording asset status. The default names (i.e. boiler plate) in the IoT sample contract implementation work perfectly fine for a single asset class, so calling it "asset" throughout the contract causes no confusion.

So as a general rule for single-asset-class contracts, generic names work and have the property of reducing the need for rework when a sample is extended or adapted into other domains. However, generic names do not differentiate asset classes in an *advanced multi-asset* contract such as the aviation sample contract, so these contracts would adopt the more explicit naming convention for assets. Like this:

``` go

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	// asset CRUD API
	if function == "createAssetAirline" {
		return t.createAssetAirline(stub, args)
	} else if function == "createAssetAircraft" {
		return t.createAssetAircraft(stub, args)
	} else if function == "createAssetAssembly" {
		return t.createAssetAssembly(stub, args)
	} else if function == "updateAssetAirline" {
		return t.updateAssetAirline(stub, args)
	} else if function == "updateAssetAircraft" {
		return t.updateAssetAircraft(stub, args)
	} else if function == "updateAssetAssembly" {
		return t.updateAssetAssembly(stub, args)
	} else if function == "deleteAssetAirline" {
		return t.deleteAssetAirline(stub, args)
	} else if function == "deleteAssetAircraft" {
		return t.deleteAssetAircraft(stub, args)
	} else if function == "deleteAssetAssembly" {
		return t.deleteAssetAssembly(stub, args)
	} else if function == "deleteAllAssetsAirline" {
		return t.deleteAllAssetsAirline(stub, args)
	} else if function == "deleteAllAssetsAircraft" {
		return t.deleteAllAssetsAircraft(stub, args)
	} else if function == "deleteAllAssetsAssembly" {
		return t.deleteAllAssetsAssembly(stub, args)
	} else if function == "deletePropertiesFromAssetAirline" {
		return t.deletePropertiesFromAssetAirline(stub, args)
	} else if function == "deletePropertiesFromAssetAircraft" {
		return t.deletePropertiesFromAssetAircraft(stub, args)
	} else if function == "deletePropertiesFromAssetAssembly" {
		return t.deletePropertiesFromAssetAssembly(stub, args)

		// event API
	} else if function == "eventFlight" {
		return eventFlight(stub, args)
	} else if function == "eventInspection" {
		return eventInspection(stub, args)
	} else if function == "eventAnalyticAdjustment" {
		return eventAnalyticAdjustment(stub, args)
	} else if function == "eventMaintenance" {
		return eventMaintenance(stub, args)

		// contract dynamic config API
	} else if function == "updateContractConfig" {
		return nil, updateContractConfig(stub, args)

		// contract state / behavior API
	} else if function == "setLoggingLevel" {
		return nil, t.setLoggingLevel(stub, args)
	} else if function == "setCreateOnUpdate" {
		return nil, t.setCreateOnUpdate(stub, args)

		// debugging API
	} else if function == "deleteWorldState" {
		return nil, t.deleteWorldState(stub)
	}
	err := fmt.Errorf("Invoke received unknown invocation: %s", function)
	log.Warning(err)
	return nil, err
}

```

As you can see, API proliferation is unaviodable in such a substantial smart contract, so it is necessary to follow clear naming conventions (patterns) and to not deviate at any point. The CRUD pattern continues to serve us well for assets, and so the adaptation is to simply add the asset's name to the standard name like this:

|Single Asset|Multi-Asset|
|:------------:|:-----------:|
|createAsset|createAssetAircraft|
|readAsset|readAssetAircraft|
|updateAsset|updateAssetAircraft|
|deleteAsset|deleteAssetAircraft|

> **Caution**: *Do not get carried away with name adaptation for asset classes. Only the asset CRUD functions should be renamed. It is **not** appropriate to rename the service function `readAssetSchemas`, which is used for integration into applications such as the [Watson IoT Platform](http://www.ibm.com/internet-of-things/iot-solutions/watson-iot-platform/). This function is a convention that cannot change in order to preserve smooth interoperation with the platform and with tools such as the generic [monitoring_UI](../../monitoring_UI).*

### Event

The word event is defined as:

> **event**  
noun  
a thing that happens, especially one of importance.  
"one of the main political events of the late 20th century"

And for a smart contract, events are important indeed. Events drive the contract state to change, as they do in any computer program, especially those that handle events from physical hardware like mice and keyboards. Each click, movement, keypress and so on is an event as far as the program is concerned.

It is much the same with smart contracts. But since event is one of the most overloaded words in the history of programming, we will draw some distinctions between *kinds* of events.

#### Device Events

The Internet of Things is about, well, *things*. In our contracts, we call the big things assets. They move, or encapsulate some important functions, or cost a lot of money and must be protected or whatever.

There are many smaller things that can tell us what is happening to the bigger things (assets), and we call the smaller things *devices*.

