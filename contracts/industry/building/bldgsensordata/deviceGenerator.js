
var fs = require("fs");

///// Read data from configration file
 var config = fs.readFileSync("configGenerator.json");
 var configData = JSON.parse(config);
// Get details required for device data generation
 var startDate = configData.startDate;
 var endDate = configData.endDate;
 var genData = configData.generateData;
 var noOfDevices = configData.numberofNewDevices;
// If generate data flag is set to true, i.e., value is 'Y'
// generate new device data
// Get device data from devices file
var devices= fs.readFileSync("devices.json");
var jsonContent = JSON.parse(devices);
var items = jsonContent.devices;
var devData=[];
for (var i = 0; i < items.length; ++i) {
     //if (i>0) {break;}
     var deviceRaw = items[i];
     var deviceStr = JSON.stringify(deviceRaw).trim();
     devData.push(deviceStr);
}

for (var newItems = 1; newItems <= noOfDevices; newItems ++) {
    var newDeviceRecord = "";
    
    if (toss()==0) {
        newDeviceRecord=generateSmartPlug(newItems, startDate, endDate);
    } else {
        newDeviceRecord=generateVariableMotor(newItems, startDate, endDate);
    }
    //var newDev = JSON.stringify(newDeviceRecord);
    //console.log("new device data is ", newDeviceRecord);
    devData.push(newDeviceRecord);
}

//jsonStr = JSON.stringify(devices);
console.log("New devices data \n ", devData.toString());
var sNewDev = "{\"devices\": [ " + devData.toString() + "]}";
var devNewFile = JSON.parse(sNewDev);
var devToFlatFile = JSON.stringify(devNewFile,null, 4);
console.log("New devices data parsed \n ",devToFlatFile);
console.log('before unlink: ');

fs.unlink('devices.json', function(err){

    // Ignore error if no file already exists
    if (err && err.code !== 'ENOENT')
        throw err;

    var options = { flag : 'w' };

    fs.writeFile('devices.json', devToFlatFile, options, function(err) {
        if (err) throw err;
        console.log('file saved');
    });
});

/////////////////generateSmartPlug//////////////
function generateSmartPlug (counter, start, end) {
    var j = "";
    /*****************************
    Sample data:  
    {
			"name": "Smart Plug", 
			"description": "Smart Plug working in XBee network", 
			"assetid" :"77ce3e45-a5d1-416a-9701-4383aa4bd00b",
			"create_date": 1417484201541,
			"last_mod_date": 1417484970337,
			"last_sample_date": 1417484981903, 
			"schema_version": 1,
			"indicators": { 
				"on_off_state": false
			} 
	    },
    *************************/
    var name = "Smart Plug - New device "+ counter;
    var desc = "Smart Plug working in XBee network - New device" + counter;
    var state =  toss()==1 ? "true}}" : "false}}";
    
     j= "{\"name\": \"" + name + "\" , \"description\" : \""+ desc+ "\", \"assetid\": \""+makeId()
    +"\" , \"create_date\" : " + randomDate(start, end) + ", \"last_mod_date\": " + randomDate(start, end) + ", \"last_sample_date\": "+
    randomDate(start, end) + ", \"schema_version\": 1, \"indicators\": { \"on_off_state\": "+state ; 
   var newSmartPlug = JSON.stringify(j); 
   var newSmartPlug = j;
    //console.log("new smart plug record is ", newSmartPlug );
    return newSmartPlug;

}
////////////////////generateVariableMotor/////////////////
function generateVariableMotor(counter, start, end) {
    var j = "";
/******************************* 
 {
    "name": "Variable Speed Motor",
    "description": "Variable Speed Motor managed by Motor Driver through Modbus",
    "assetid": "66fad0c0-1cfb-11e4-8c21-0800200c9a66",
    "create_date": 1417484200223,
    "last_mod_date": 1417484681910,
    "last_sample_date": 1417484981895,
    "schema_version": 1,
    "indicators": {
        "acceleration_time": 6.0,
        "on_off_state": true,
        "direction": false,
        "stop_method": 1,
        "base_rpm": 1525,
        "max_rpm": 3600,
        "rpm": 2500,
        "output_current": 10,
        "output_voltage": 115,
        "output_frequency": 34.0
    }
}
*******************************/
    var name = "Variable Speed Motor - New device "+ counter;
    var desc = "Variable Speed Motor managed by Motor Driver through Modbu - New device" + counter;
    var accelTime = Math.floor(Math.random() * 10)+Math.round( Math.random() * 10 ) / 10;
    var state =  toss()==1 ? "true" : "false";
    var direction =  toss()==1 ? "true" : "false";
    var stopMethod = toss();
    var baserpm = Math.floor(1000 + Math.random() * 9000);
    var maxrpm = Math.floor(1000 + Math.random() * 9000);
    if (maxrpm < baserpm) {
        maxrpm = baserpm + 1000;
    }
    var rpm = Math.floor(1000 + Math.random() * 9000);
    if (rpm > maxrpm) {
        rpm = Math.round(rpm/10);
    }
    var current = Math.floor(Math.random() * 90 + 10);
    var voltage = Math.floor(Math.random() * 90 + 10);
    var frequency = Math.floor(Math.random() * 90 + 10)+ Math.round( Math.random() * 10 ) / 10;
    j= "{\"name\": \"" + name + "\" , \"description\" : \""+ desc+ "\", \"assetid\": \""+makeId()
     +"\" , \"create_date\" : " + randomDate(start, end) + ", \"last_mod_date\": " + randomDate(start, end) + ", \"last_sample_date\": "+
    randomDate(start, end) + ", \"schema_version\": 1, \"indicators\": { \"acceleration_time\": " + accelTime 
    + ",\"on_off_state\": "+state+",\"direction\": "+direction+",\"stop_method\": "+stopMethod
    + ",\"base_rpm\": "+baserpm+",\"max_rpm\": "+maxrpm+",\"rpm\": "+rpm
    + ",\"output_current\": "+current+",\"output_voltage\": "+voltage+",\"output_frequency\": "+frequency+"}}";
     //var newVariableMotor = JSON.stringify(j); 
     var newVariableMotor = j;
    //console.log("new variable motor record is ", newVariableMotor );
    return newVariableMotor;
}

////////////////randomDate////////////////////////
// generate random date in millisecond time
function randomDate(start, end) {
    var startDate = new Date(start)
    var endDate = new Date(end)
    var date = new Date(+startDate + Math.random() * (endDate - startDate));
    var msTime = date.getTime(); 
    //console.log("ms time is ", msTime)
    return msTime;
}
/////////////////////toss/////////////////////////
// Binary toss, to decide which device record to send.
function toss() {
    answer = Math.round(Math.random());
    return answer;
}
///////////makeSmartPlugId/////////////////////////
function makeId() {
    var newID = "";
    var text = "";
    var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for( var i=0; i < 8; i++ ) {
        newID += possible.charAt(Math.floor(Math.random() * possible.length));
    }
    
    newID += "-";
    for (var j=0; j<3;j++) {
        for( var i=0; i < 4; i++ ) {
            newID += possible.charAt(Math.floor(Math.random() * possible.length));
        }
        newID += "-";
    }
    for( var i=0; i < 12; i++ ) {
        newID += possible.charAt(Math.floor(Math.random() * possible.length));
    }
    return newID;
}
