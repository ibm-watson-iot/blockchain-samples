
//Global variables
var clicks=0;
var countdown=0;
var nIntervId;
var sMeterid ;
var sMeterNode;
var ratePerSec=0.025;
var amount;
var setFree =true;
var sRate ='$0.25 / 10 sec' ;
var thisMeterid = "";


//Array of parking meters. Currently hardcoded for PKM-001 thru PKM-010
var arr = [ "PKM-001", "PKM-002", "PKM-003", "PKM-004", "PKM-005", "PKM-006", "PKM-007", "PKM-008", "PKM-009", "PKM-010" ];
var nodes = '{ "noderoutes" : [' +
'{ "meterid":"PKM-001" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-002" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-003" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-004" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-005" , "link":"http://mbedpkmflow.mybluemix.net/" },' +
'{ "meterid":"PKM-006" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-007" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-008" , "link":"http://mbedpkmflow.mybluemix.net/" },' +
'{ "meterid":"PKM-009" , "link":"https://my-watson-nodered-app.mybluemix.net/" },' +
'{ "meterid":"PKM-010" , "link":"http://mbedpkmflow.mybluemix.net/" }]}';


// Change title on portrait / landscape views
jQuery(window).bind('orientationchange', function(e) {
  switch ( window.orientation ) {
    case 0:
        $('#meterid').html(thisMeterid);
    break;
    case 90:
    case -90:
        $('#meterid').html("The Car is at Pay Location "+thisMeterid);
    break;
    default: {}
  }
  
});
///////////////// Blockchain calls //////////////////////
//  Get the device's rate and duration
function getDeviceRates() {
  var chaincodeID="043ab3fb66cb2a0886f5cc535b014024a4ba2cc1f80bf8b9cb9e6cce17a01f54cc20fda32a2d59010a06d510f0edb949b5a897d9275b9224233665d944baf929";
  var fabricPeer = "https://808accbe-4cdc-42e7-8c69-1908716c7890_vp1.us.blockchain.ibm.com:443/chaincode";
  var jsonString = '{ "jsonrpc": "2.0","method": "query","params": {"type": 1,"chaincodeID": '+ 
  '{"name": "'+ chaincodeID+ '"}, "ctorMsg": { "function": "readDevice", '+ 
  '"args": ["{\\\"deviceid\\\": \\\"'+sMeterid+'\\\"}"]},"secureContext": "user_type1_760ffab3fa"},'+
  '"id": 0 }';
   console.log(jsonString);
   $.ajax({
      type: "POST",
      url: fabricPeer,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
          obj=JSON.parse(data.result.message);
          minCost = Number(obj.minimumusagecost);
          minDur = Number(obj.minimumusagetime);
          if (obj.available==false)
            {
              // If device is unavailable
              sUnavilMsg = 'Meter ' + sMeterid + '  is unavailable';
              $('#meterid').html(sUnavilMsg);
              nope();
            } else {
              yup();
            }
          if (mincost<=0)
          {
            ratePerSec=0.0;
            sRate = 0.0;
            $('#amount').html("$0");
          } else {
            ratePerSec=minCost/minDur;
            sRate = obj.minimumusagecost+" / "+obj.minimumusagetime+" secs";
          }

      },
      error: function (jqXHR, textStatus, errorThrown)
        {
          console.log(errorThrown);
        }
  });
}

// Block the device in the blockhain for the specified period of time
function createUsageRecord() {
  var chaincodeID="043ab3fb66cb2a0886f5cc535b014024a4ba2cc1f80bf8b9cb9e6cce17a01f54cc20fda32a2d59010a06d510f0edb949b5a897d9275b9224233665d944baf929";
  var fabricPeer = "https://808accbe-4cdc-42e7-8c69-1908716c7890_vp1.us.blockchain.ibm.com:443/chaincode";

   var jsonString = '{ "jsonrpc": "2.0","method": "invoke","params": {"type": 1,"chaincodeID": '+ 
  '{"name": "'+ chaincodeID+ '"}, "ctorMsg": { "function": "createUsage", '+ 
  '"args": ["{\\\"deviceid\\\":\\\"'+sMeterid+'\\\", \\\"starttime\\\":\\\"'+getDateTime()+'\\\", \\\"duration\\\":'+ clicks + '}"]},'+
  '"secureContext": "user_type1_760ffab3fa"},'+
  '"id": 0 }';
   console.log(jsonString);
   $.ajax({
      type: "POST",
      url: fabricPeer,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
         console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
        console.log(errorThrown);
      }
      
  });

 }
 // Once the timer runs out, make the device available again
function makeMeterAvailable() {

  var chaincodeID="043ab3fb66cb2a0886f5cc535b014024a4ba2cc1f80bf8b9cb9e6cce17a01f54cc20fda32a2d59010a06d510f0edb949b5a897d9275b9224233665d944baf929";
  var fabricPeer = "https://808accbe-4cdc-42e7-8c69-1908716c7890_vp1.us.blockchain.ibm.com:443/chaincode";

  var jsonString = '{ "jsonrpc": "2.0","method": "invoke","params": {"type": 1,"chaincodeID": '+ 
  '{"name": "'+ chaincodeID+ '"}, "ctorMsg": { "function": "updateDeviceAsAvailable", '+
  '"args": ["{\\\"deviceid\\\":\\\"'+sMeterid+'\\\", \\\"available\\\":true}"]},"secureContext": "user_type1_760ffab3fa"},'+
  '"id": 0 }';
   console.log(jsonString);
   $.ajax({
      type: "POST",
      url: fabricPeer,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
        console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
        console.log(errorThrown);
      }
  });
 }
///////////////// node-red calls //////////////////////
// Once the timer runs out, make the device available again
function startCounter() {
  var jsonString = '{ "timestamp": "' + new Date() + '", "meterid": "'+sMeterid+'"}'
  //alert (jsonString);
  sStart = sMeterNode+'start'+sMeterid;
   $.ajax({
      type: "POST",
      url: sStart,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
         console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
        console.log(errorThrown);
      }
      
  });
 }
//Set the parking meter free : toggle dollar sign
 function setFreeMeter() {
  var jsonString = '{ "timestamp": "' + new Date() + '"}'
  sFree = sMeterNode+'free'+sMeterid;
  //alert (sFree);
   $.ajax({
      type: "POST",
      url: sFree,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
         console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {

      console.log(errorThrown);
      }
      
  });
}
//Set meter back to paid: toggle dollar sign
function setPaidMeter() {
  var jsonString = '{ "timestamp": "' + new Date() + '"}'
  sFree = sMeterNode+'paid'+sMeterid;
   $.ajax({
      type: "POST",
      url: sFree,
      data: jsonString,// now data come in this function,
      dataType:"json",
      success: function(data, textStatus, jqXHR)
      {
         console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
      console.log(errorThrown);

      }
      
  });
 
 }
///////////////////Button click events ////////////
// Add time
$(document).on("click", "#add", function(){
     clicks=clicks+10; 
      $('#timer').html(pad(clicks,2));
      $( "#timer" ).css('font-size', '40px');
       $( "#timer" ).css('color', '#66a3ff');
       //$("#subtract").prop('disabled', false);
       $("#subtract").attr("disabled", false);
      $("#submit").attr("disabled", false);
      if (ratePerSec ==0.0)
      {
        amount = 0;
      }
      else {
        amount=clicks* ratePerSec;
      }
      
      $('#amount').html("$"+amount);
      $( "#amount" ).css('color', '#339966');
      
});
//remove time
$(document).on("click", "#subtract", function(){
      if (clicks >=10 ) {
        clicks=clicks-10; 
        $('#timer').html(pad(clicks,2));
        $( "#timer" ).css('font-size', '40px');
         $( "#timer" ).css('color', '#66a3ff');
         if (ratePerSec ==0.0)
        {
          amount = 0;
        } else {
          amount=clicks*ratePerSec;
        }
        $('#amount').html("$"+amount);
        $( "#amount" ).css('color', '#339966');
        $( "#timer" ).css('color', '#66a3ff');

      }
     
});
//Click refresh button
$(document).on("click touchstart", "#refbtn", function(){
    location.reload();
});
//Click free button - toggles free and paid
$(document).on("click touchstart", "#freebtn", function(){
  if (setFree) {
    $('#amount').html("$0");
    $('#rateinfo').html("Free Parking!");
    setFreeMeter();
    nope();
    setFree = false;
  } else {
    $('#amount').html("$0");
    $('#rateinfo').html(sRate);
    setFree = true;
    setPaidMeter();
    yup();
  }
});
//Click the submit button
$(document).on("click", "#submit", function(){
  if (clicks >=10)
  {
    $("#refbtn").attr("disabled", "disabled");
    createUsageRecord(); // Temporarily commenting expensive rest calls for testing
    startCounter();
    makePayment();
     changeColor();
      $("#subtract").attr("disabled", "disabled");
      $("#add").attr("disabled", "disabled");
      //$("#this").attr("disabled", "disabled");
      $(this).css('background-color','lightgrey');
      //alert('before make meter available');
    }
      
});
// click credit card
////////////// IBM Payment integration /////////////
/*
$(document).on("click", "#paylink", function(){
    //alert('button clicked');
    $.ajax({
        type: 'POST',
        url: '/pymt'
    });
});
*/
//////////////////////Others////////////////////////
//Timestamp in UTC
function getDateTime() {
  var currentdate = new Date(); 
  todaysDate = currentdate.getUTCDate()<10? "0"+currentdate.getUTCDate():currentdate.getUTCDate();
  todaysMonth = (currentdate.getUTCMonth()+1)<10? "0"+(currentdate.getUTCMonth()+1):(currentdate.getUTCMonth()+1);
  thisHour = currentdate.getUTCHours() <10? "0"+currentdate.getUTCHours():currentdate.getUTCHours();
  thisMin = currentdate.getUTCMinutes() <10? "0"+currentdate.getUTCMinutes():currentdate.getUTCMinutes();
  thisSec = currentdate.getUTCSeconds() <10? "0"+currentdate.getUTCSeconds():currentdate.getUTCSeconds();

  var datetime = currentdate.getUTCFullYear()+"-"+ todaysMonth  + "-" +todaysDate+ " "
                + thisHour + ":" + thisMin + ":" + thisSec;
  //alert (datetime);
  return datetime;
}
//Computing the timer 
function pad(num, size) {
    var s = num+"";
    if (s>=60) {
      min = Math.floor(s/60);
      sec=s-(min*60);
      if (sec==0)
      {
        secs=": 00";
      }
      if (sec<10)
        {
        secs=":0" + sec ;
      } else {
        secs=":"+sec;
      }
      s=min+secs+ " m";
    } else if(s<=0) {
      $("#subtract").attr("disabled", "disabled");
      s="00:00";
    } else {
       s=s+":00 s";
    }
    return s;
}

//timer countdown functions
 function changeColor() {
      nIntervId = setInterval(flashText, 1000);     
}
 
function flashText() {
  //$( "#timer" ).css('color',  ("red" ? "blue" : "red"));
  if (clicks==10)
  {
    warntime=5;
  } else {
    warntime=10;
  }
  var oElem = document.getElementById("timer");
  oElem.style.color = oElem.style.color == "chocolate" ? "darkcyan" : "chocolate";
  countdown++;
    clicks--;
  $('#timer').html(pad(clicks,2));
  if (clicks==0) 
    {
        $('#timer').html("Expired!");
        $( "#timer" ).css('color', 'red');
         //Update Parking meter to set it as available
        clearInterval(nIntervId);
          makeMeterAvailable(); // Temporarily commenting expensive rest calls for testing
          $("#refbtn").attr("disabled", false);
      return;
    }
}

// Page load 

window.onload = function() {
  showLoc();
  getnodeFlowLink();
  getPaymentLink();
};

function showLoc() {

  var oLocation = location, aLog = ["Property (Typeof): Value", "location (" + (typeof oLocation) + "): " + oLocation ];
  for (var sProp in oLocation){
    aLog.push(sProp + " (" + (typeof oLocation[sProp]) + "): " +  (oLocation[sProp] || "n/a"));
    if(sProp=="search") {
      sTemp = (oLocation[sProp] || "n/a");
      sKey=sTemp.substring(0,4);
      sValue = sTemp.substring(4,sTemp.length);
      if (sKey=="?id=") {
        sMeterid=sValue;
      } else {
        sMeterid = sTemp;
      }
    }
  }
  iValidMeter=jQuery.inArray( sMeterid, arr );
   if(iValidMeter==-1) {
    var sWelcome = "No meter id or unknown meter!";
//       $( "#meterid" ).css('font-size', '25px');
//       $( "#meterid" ).css('color', '#FF0000');
//       $( "#meterid" ).text( sWelcome);
       $("#submit").attr("disabled", "disabled");
       nope();   
  }

  else {
     var sWelcome = "Meter "+sMeterid  ;
       $( "#meterid" ).text(sWelcome);
//       $( "#meterid" ).css('font-size', '25px');
//       $( "#meterid" ).css('color', '#000000 ');
//       $("#meterid").css("background-color","#65a7f5 ");
        thisMeterid=sWelcome;
       $("#submit").attr("disabled", false);
       getDeviceRates(); // Temporarily commenting expensive rest calls for testing
       //$("#meterid").css("color", blue);
  }
     
}

function nope(){
   $("#submit").attr("disabled", "disabled");
//   $( "#meterid" ).css('font-size', '25px');
//   $( "#meterid" ).css('color', '#000000 ');
//   $("#meterid").css("background-color","#65a7f5 ");
   $("#add").attr("disabled", "disabled");
   $("#subtract").attr("disabled", "disabled");
}

function yup(){
   $("#submit").attr("disabled", false);
//   $("#meterid").css("background-color","#65a7f5");
   $("#add").attr("disabled", false);
   $("#subtract").attr("disabled", false);

}

function getPaymentLink() {
  $.ajax({
      type: "POST",
      url: '/session',
      data: '',
      success: function(data, textStatus, jqXHR)
      {
        //alert('success');
         //alert('url received is : ' +JSON.stringify(data));
         console.log(JSON.stringify(data));
         //alert (data.sessionPath);
         $("a[href='http://www.example.com/']").attr('href', data.sessionPath);
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
        console.log(errorThrown);
      }
  });
}

function makePayment() {
  var amtString = amount.toString(); 
  var data = {};
  data.title = "money";
  data.message = "amtString";
  //alert('amount is '+amtString);
  //var amtData = {
  //  amount: amtString}
  $.ajax({
      type: "POST",
      url: '/pymt',
      data: JSON.stringify(data),
      contentType: 'application/json',
      success: function(data, textStatus, jqXHR)
      {
        //alert(JSON.stringify(data));
         console.log(JSON.stringify(data));
      },
      error: function (jqXHR, textStatus, errorThrown)
      {
        //alert(errorThrown);
        console.log(errorThrown);
      }
  });
}
 
function getnodeFlowLink() {
  var obj = JSON.parse(nodes);
  for (i = 0; i < obj.noderoutes.length; i++) {
    if (obj.noderoutes[i].meterid == sMeterid) {
      sMeterNode = obj.noderoutes[i].link;
      break;
    }
  }
}

