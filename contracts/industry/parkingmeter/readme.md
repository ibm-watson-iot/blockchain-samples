#Parking Meter Demo


##Quick Overview.

The Parking meter is FRDM-K64 and Cordio Boards combined. The FRDM board acts as the meter and the cordio acts as the BLE Beacon. The beacon transmits short urls in Eddystone standard and this points to the bluemix page *https://mbed-parkingmeter.mybluemix.net.* This page supports pakring meters PKM-001 through PKM-010. The code for the UI is inside *mbedParkingMeterUI*. 

The UI makes a call to the IBM commerce system to initiate payment and the chaincode (code availalbe in *mbedParkingMeter / mbedParkingMeter.0.6*) for recording parking meter usage data.

It makes http calls to the node-red flow - *NodeFlow.json* -for setting the Parking meter to free (beacon stops emitting, UI says 'Free Parking') or paid (beacon starts emitting, UI says 'Paid parking') and also initiating the countdown once payment is made. 
