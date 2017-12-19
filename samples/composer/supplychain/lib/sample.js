/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

function emitTraceEvent(namespace, eventname, message, payload) {
    var evt = getFactory().newEvent(namespace, eventname);
    evt.message = message;
    try {
        evt.payload = "object: " + typeof(payload) + " " + JSON.stringify(payload);
    } catch (err) {
        evt.payload = "resource: " + JSON.stringify(serializer.toJSON(payload,{validate:false})); // yes, seriously
    }
    emit(evt);
}

function emitEvent(namespace, eventname, tx) {
    var event = getFactory().newEvent(namespace, eventname);
    event.deviceId = tx.deviceId;
    event.description = tx.description;
    event.date = tx.date;
    event.time = tx.time;
    event.altitude = tx.altitude;
    event.accelX = tx.accelX;
    event.accelY = tx.accelY;
    event.accelZ = tx.accelZ;
    event.latitude = tx.latitude;
    event.longitude = tx.longitude;
    event.temperature = tx.temperature;
    event.humidity = tx.humidity;
    event.eventCode = tx.eventCode;
    event.shipmentId = tx.shipmentId;
    event.participant = tx.participant;

    try {
        emit(event);
    } catch (err) {
        emitTraceEvent(namespace, "LogInfoEvent", "Emit failed with: ", err);
    }
}

/**
 * Sample transaction processor function.
 * @param {org.poc.scbn.CreateIoTDevice} tx The sample transaction instance.
 * @transaction
 */
function CreateIoTDevice(tx) {
    // Get the asset registry for IoTDevices
    return getAssetRegistry('org.poc.scbn.IoTDevice')
        .then(function (assetRegistry) {
            // Create an asset to hold the incoming values
            var factory = getFactory();
            var IoTdevice = factory.newResource('org.poc.scbn', 'IoTDevice', tx.deviceId);
            IoTdevice.deviceId = tx.deviceId;
            IoTdevice.description = tx.description;
            IoTdevice.date = tx.date;
            IoTdevice.time = tx.time;
            IoTdevice.altitude = tx.altitude;
            IoTdevice.accelX = tx.accelX;
            IoTdevice.accelY = tx.accelY;
            IoTdevice.accelZ = tx.accelZ;
            IoTdevice.latitude = tx.latitude;
            IoTdevice.longitude = tx.longitude;
            IoTdevice.temperature = tx.temperature;
            IoTdevice.humidity = tx.humidity;
            IoTdevice.eventCode = tx.eventCode;
            IoTdevice.shipmentId = tx.shipmentId;
            IoTdevice.participant = tx.participant;
      
            // Add the asset to the asset registry
            return assetRegistry.add(IoTdevice);
        })
        .then(function() {
            // Emit an event for the created asset
            emitEvent('org.poc.scbn', 'IoTDeviceCreatedEvent', tx);
        });
}

/**
 * Sample transaction processor function.
 * @param {org.poc.scbn.IoTDeviceReading} tx The sample transaction instance.
 * @transaction
 */
function IoTDeviceReading(tx) {
    // Get the asset registry for IoTDevices
    return getAssetRegistry('org.poc.scbn.IoTDevice')
        .then(function (assetRegistry) {
            // Create an asset to hold the incoming values
            var factory = getFactory();
            var IoTdevice = factory.newResource('org.poc.scbn', 'IoTDevice', tx.deviceId);
            IoTdevice.deviceId = tx.deviceId;
            IoTdevice.description = tx.description;
            IoTdevice.date = tx.date;
            IoTdevice.time = tx.time;
            IoTdevice.altitude = tx.altitude;
            IoTdevice.accelX = tx.accelX;
            IoTdevice.accelY = tx.accelY;
            IoTdevice.accelZ = tx.accelZ;
            IoTdevice.latitude = tx.latitude;
            IoTdevice.longitude = tx.longitude;
            IoTdevice.temperature = tx.temperature;
            IoTdevice.humidity = tx.humidity;
            IoTdevice.eventCode = tx.eventCode;
            IoTdevice.shipmentId = tx.shipmentId;
            IoTdevice.participant = tx.participant;

            // Update the asset in the asset registry.
            return assetRegistry.update(IoTdevice);
        })
        .then(function () {
            // Emit an event for the modified asset.
            emitEvent('org.poc.scbn', 'IoTUpdateEvent', tx);
        });
}
