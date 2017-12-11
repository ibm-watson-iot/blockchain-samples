# IoT Blockchain - Mapping Device Events to Contract Events

[`payloadSchema.json`](../payloadSchema.json "the contract schema")

## Introduction

The Internet Of Things consists of billions of devices, each sending data in a device-specific format to a target IP address and port. Some application will then apply a context to this *device event* and process it in an application-specific way. Distributed applications have component chains (or layers, or both) that make successive transformations to data on its way from a source to a target.

For generic blockchain-based applications, there may be a component whose job it is to apply context to the device event, for example a bar code scan that indicates arrival at a port, and send it on to the application component that logs such arrivals in some database.

There is a great deal of development cost invested in applications, with schemas (or more likely hand-written encode/decode software) everywhere and it is easy to imagine device-specific application components proliferating like software rabbits.

## IoT Platform

Enter the IoT Platform (IoTP).

In a great many cases, an event from a device already has context based on its location or can trivially have the context applied. Perhaps the device can be programmed to send a bit of text that signifies an action (e.g. arrived) with a geolocation or a port name to indicate where it has arrived. There is often enough information to fully describe what the event is about.

With IoTP, device events can be received directly and the can be mapped to an application-specific blockchain contract event that then stores the information in an asset's state. The device scanned a bar code in our example, and it is that bar code that is sent forward to indicate which asset has arrived. That could be a package or a contain, and the contract need not care.

## Mapping Using Contract Schema

Say that the device event looks like this on the wire:

```json

{
    "ID": "xyz",
    "loc": {
        "lat": 49,
        "long": -97
    },
    "scancode": "abcd",
    "meta1": "HLS"
}

```

This simplistic representation of a device's output allows us to map onto a contract directly as there is plenty of information.

There are fields here that are not in the contract. To be more illustrative of how mapping should work, the shape of the event structure is augmented here to include some separation of the common fields that all crud-style contracts must have and the custom fields that contains customizations of the basic pattern. Here is a sample of such a schema from the output of the `go generate` script.

``` json

    "event": {
        "common": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": {},
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "timestamp": "2016-03-13T18:30:50.386636281Z"
        },
        "custom": {
            "carrier": "transport entity currently in possession of asset",
            "gForce": 123.456,
            "temperature": 123.456
        }
    },

```

Expanding the schema for `event` to include the extra fields is a trivial exercise.
The new schema would look like this:

``` json

    "event": {
        "common": {
            "assetID": "The ID of a managed asset. The resource focal point for a smart contract.",
            "extension": {},
            "location": {
                "latitude": 123.456,
                "longitude": 123.456
            },
            "timestamp": "2016-03-13T18:30:50.386636281Z"
        },
        "custom": {
            "carrier": "transport entity currently in possession of asset",
            "gForce": 123.456,
            "temperature": 123.456,
            "device": "a devicie-specific ID, often a GUID",
            "port": "sent by some devices as metadata"
        }
    },

```

Now we can map the device data to the contract data. We use qualified field names in dot-notation to reference the hierarchy of properties (hence why an advanced version of the schema is shown).

Mapping from the device to the contract is best performed conceptually in tabular form, as this is the natural way to express a mapping. Do note that the initial level `event` is a constant when using the IoT mapper and so is not needed in the qualified names.

|Device|Contract|
|:------|:--------|
|ID|custom.device|
|loc.lat|common.location.latitude|
|loc.long|common.location.longitude|
|meta1|custom.port|
|scancode|common.assetID|

The mapper can now build a new event from scratch that will look like this:

``` json

{
    "chaincodeSpec":{
        "type": "GOLANG",
        "chaincodeID":{
            "name":"123JKKJH87KL563123JKKJH87KL563123JKKJH87KL563123JKKJH87KL563"
        },
        "ctorMsg":{
            "function":"updateAsset",
            "args":["{\"common\": {\"assetID\":\"abcd\",
                                   \"location\":{\"latitude\":49,\"longitude\":-97}}
                      \"custom\": {\"port\":\"HLS\",
                                   \"device\":\"xyz\"}}"]
        }
    }
}

```

*Note: Name is a UUID that was returned when the contract was deployed. IoTP remembers that detail for you.*

*Note2: The args to a contract are always strings, so quotes are escaped. Line breaks are shown for the reader's convenience.*

## Mapping Using The Extension Mechanism

But what if the contract is already running in the cloud in a form that does not contain the additional data? And what happens if you cannot redeploy for some reason? How can the device data that is not in the schema be represented?

This is where the `extension` object comes into play. It is an opaque, application-managed object that is invisible to the contract's rules engine (for now). It is handled safely by the deep merge in mapUtils so that events can contain one or more application-specific properties that act additively as demanded by the *partial state as event* pattern. 

So these build up in the asset's state just as schema-based properties do. Think of these as *sidecar* properties that ride along in the event so that appropriate associations are maintained in the blockchain. For example, data stored elsewhere can be hashed and that hash stored in an extension property, even when the need went unanticipated until the contract was deployed. 

The mechanism works particularly well because of the fluid nature of JSON encoding. IoT patterns do not overuse the *required* schema property, thus leaving events to be as rich or as sparse as is necessary for the devices in use by the application. (Again, we call this the *partial state as event* pattern.)

The mapping for this mechanism changes only slightly.

|Device|Contract|
|:------|:--------|
|ID|common.extension.device|
|loc.lat|common.location.latitude|
|loc.long|common.location.longitude|
|meta1|common.extension.port|
|scancode|common.assetID|

And the message changes to:

``` json

{
    "chaincodeSpec":{
        "type": "GOLANG",
        "chaincodeID":{
            "name":"123JKKJH87KL563123JKKJH87KL563123JKKJH87KL563123JKKJH87KL563"
        },
        "ctorMsg":{
            "function":"updateAsset",
            "args":["{\"common\": {\"assetID\":\"abcd\",
                                   \"location\":{\"latitude\":49,\"longitude\":-97},
                                   \"extension\": {\"port\":\"HLS\",
                                                   \"device\":\"xyz\"}}"]
        }
    }
}

```

## Creating Assets from Events

The function name sent in the contract payload is called `updateAsset` in this pattern. By default, it is able to create the asset if it is not already. So by programming even one device to send an event with an assetID, device registration can take place.

On the other hand, there may be cases where this behavior is undesirable. In such situations, updateAsset events should fail if the asset is unknown. The `setCreateOnUpdate` function can change this behavior while the contract is running with its boolean parameter.

The message for that looks like:

``` json

{
    "chaincodeSpec":{
        "type": "GOLANG",
        "chaincodeID":{
            "name":"mycc"
        },
        "ctorMsg":{
            "function":"setCreateOnUpdate",
            "args":["{\"setCreateOnUpdate\": false}"]
        }
    }
}

```