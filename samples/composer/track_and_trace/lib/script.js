/**
 * A tilt reading has been received for a surgical kit
 * @param {org.acme.model.TiltReading} tiltReading - the TiltReading transaction
 * @transaction
 */
function tiltReading(tiltReading) {

    var surgicalKit = tiltReading.surgicalKit;
  
    // set kit status to 'in transit'
    surgicalKit.status = 'IN_TRANSIT';
    console.log('Adding tilt ' + tiltReading.tilt + ' to surgicl kit ' + surgicalKit.$identifier);

    if (surgicalKit.tiltReadings) {
        surgicalKit.tiltReadings.push(tiltReading);
    } else {
        surgicalKit.tiltReadings = [tiltReading];
    }

    return getAssetRegistry('org.acme.model.SurgicalKit')
        .then(function (surgicalKitRegistry) {
            // add the tilt reading to the shipment
            return surgicalKitRegistry.update(surgicalKit);
        });
}


/**
 * A force reading has been received for a surgical kit
 * @param {org.acme.model.ForceReading} forceReading - the ForceReading transaction
 * @transaction
 */
function forceReading(forceReading) {

    var surgicalKit = forceReading.surgicalKit;
  
    // set kit status to 'in transit'
    surgicalKit.status = 'IN_TRANSIT';       
    console.log('Adding force ' + forceReading.force + ' to surgical kit ' + surgicalKit.$identifier);
  
     if (surgicalKit.forceReadings) {
       surgicalKit.forceReadings.push(forceReading);        
    } else {
        surgicalKit.forceReadings = [forceReading];
    }

    return getAssetRegistry('org.acme.model.SurgicalKit')
        .then(function (surgicalKitRegistry) {
            // add the force reading to the shipment
            return surgicalKitRegistry.update(surgicalKit);
        });
}

/**
 * A surgical kit has been received by a shipper
 * The contract will be checked.
 * The location of the surgical kit within the hospital will be checked here as well
 * @param {org.acme.model.ShipmentReceived} shipmentReceived - the ShipmentReceived transaction
 * @transaction
 */
function addKit(shipmentReceived) {

    var contract = shipmentReceived.surgicalKit.contract;
    var surgicalKit = shipmentReceived.surgicalKit;
    var addKit = 1;

    console.log('Received at: ' + shipmentReceived.timestamp);
    console.log('Contract arrivalDateTime: ' + contract.arrivalDateTime);

    // set the status of the shipment
    surgicalKit.status = 'ARRIVED';

    // if the shipment did not arrive on time do not take the kit
    if (shipmentReceived.timestamp > contract.arrivalDateTime) {
        addKit = 0;
        console.log('Late shipment');
      
    } else {
      
        // find the highest tilt reading
        if (surgicalKit.tiltReadings) {  
          
          //sort the readings
          surgicalKit.tiltReadings.sort(function(a,b) {
            return (b.tilt - a.tilt);
          });
          
         var highestTilt = surgicalKit.tiltReadings[0];                
            console.log('Highest tilt reading: ' + highestTilt.tilt);
          
            // does the highest tilt violate the contract?             
           if (highestTilt.tilt > contract.maxTilt) {
                addKit = 0;
                console.log('Tilt reading violates the contract.');
             
            } else {
              
              // find the highest force reading
              if (surgicalKit.forceReadings) {
                
                // sort the force readings
                surgicalKit.forceReadings.sort(function(a,b) {
                  return (b.force - a.force);
                });
                
                var highestForce = surgicalKit.forceReadings[0];                
                console.log('Highest force reading: ' + highestForce.force);   
                
                //does the highest force violate the contract?
                if (highestForce.force > contract.maxForce) {
                  addKit = 0;
                  console.log('Force reading violates the contract.');
                }
              }
            }            
          }
       contract.shipper.numOfKits = contract.shipper.numOfKits - addKit; 
       contract.hospital.numOfKits = contract.hospital.numOfKits + addKit;
    }


    console.log('Shipper: ' + contract.shipper.$identifier + ' new number of kits: ' + contract.shipper.numOfKits);
    console.log('Hospital: ' + contract.hospital.$identifier + ' new number of kits: ' + contract.hospital.numOfKits);

    return getParticipantRegistry('org.acme.model.Shipper')
        .then(function (shipperRegistry) {
            // update the shipper's balance
            return shipperRegistry.update(contract.shipper);
        })
        .then(function () {
            return getParticipantRegistry('org.acme.model.Hospital');
        })
        .then(function (hospitalRegistry) {
            // update the hospital's balance
            return hospitalRegistry.update(contract.hospital);
        })
        .then(function () {
            return getAssetRegistry('org.acme.model.SurgicalKit');
        })
        .then(function (surgicalKitRegistry) {
            // update the state of the shipment
            return surgicalKitRegistry.update(surgicalKit);
        });
}

/**
* A surgical kit is at the hospital and is moving. 
* This function alerts the hospital if the kit moves out of the geofence.
* @param {org.acme.model.KitMovement} kitMovement - the KitMovement transaction
* @transaction
*/

function kitMovement(kitMovement) {
  
  var contract = kitMovement.surgicalKit.contract;
  var surgicalKit = kitMovement.surgicalKit;
  var geoLat = kitMovement.surgicalKit.contract.hospital.geoLat;
  var geoLong = kitMovement.surgicalKit.contract.hospital.geoLong;
  var latitude = kitMovement.latitude;
  var longitude = kitMovement.longitude;
  
  // Check if the kit is in the hospital
  if (surgicalKit.status != 'ARRIVED') {
    throw new Error('Cannot perform transaction, kit is not in the hospital.');
  }
  
 var dLong = longitude - geoLong;
  var dLat = latitude- geoLat;
  var distance = Math.sqrt(Math.pow(dLong, 2) + Math.pow(dLat, 2));
  
  // check if kit is outside of geofence
  if (distance > kitMovement.surgicalKit.contract.hospital.geoRadius) {
    throw new Error('Cannot move surgical kit outside of permited area!');
  }
}
