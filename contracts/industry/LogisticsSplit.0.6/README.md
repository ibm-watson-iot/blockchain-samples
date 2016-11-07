#Update - commandline steps

##Registration (only in sandbox)
CORE_CHAINCODE_ID_NAME=comp CORE_PEER_ADDRESS=0.0.0.0:7051 ./compliance
CORE_CHAINCODE_ID_NAME=cont CORE_PEER_ADDRESS=0.0.0.0:7051 ./container
CORE_CHAINCODE_ID_NAME=blReg CORE_PEER_ADDRESS=0.0.0.0:7051 ./billoflading

##Deploy
peer chaincode deploy -n comp -c '{"function":"Init", "args":["{\"Version\":\"2.0.0\"}"]}'
peer chaincode deploy -n cont -c '{"function":"Init", "args":["{\"Version\":\"2.0.0\", \"compliancecc\":\"comp\"}"]}'
peer chaincode deploy -n blReg -c '{"function":"Init", "args":["{\"Version\":\"2.0.0\", \"containercc\":\"cont\", \"compliancecc\":\"comp\"}"]}'

##Create Bill of Lading
peer chaincode invoke -n blReg -c '{"function":"registerBillOfLading", "args":["{\"blno\":\"10203040\", \"containernos\":\"CONT1000,CONT2000\", \"hazmat\":false, \"mintemperature\":-10, \"maxtemperature\":30, \"minhumidity\":0, \"maxhumidity\":50, \"minlight\":0, \"maxlight\":30, \"minacceleration\":0.01, \"maxacceleration\":2}"]}'

##Query
###Use the below to get bill of lading registration data
peer chaincode query -n blReg  -c '{"function":"getBillOfLadingRegistration",  "args":["{\"blno\":\"10203040\"}"]}'
Query Result: 
```
{"blno":"10203040","containernos":"CONT1000,CONT2000","mintemperature":-10,"maxtemperature":30,"maxhumidity":50,"maxlight":30,"minacceleration":0.01,"maxacceleration":2,"timestamp":"2016-11-04 23:50:54.922599827 +0000 UTC"}
```
###Use the below to read current container status
peer chaincode query -n cont -c '{"function":"readContainerCurrentStatus", "args":["{\"containerno\":\"CONT1000\"}"]}'
Query Result: 
```
{"containerno":"CONT1000","blno":"10203040","location":{},"timestamp":"2016-11-04 23:50:54.922599827 +0000 UTC","airquality":{}}
```

###Use the below to find the last compliance violation raised
peer chaincode query -n comp -c '{"function":"readCurrentComplianceState", "args":["{\"blno\":\"10203040\"}"]}'
Query Result: 
```
{"blno":"10203040","type":"SHIPPING","compliance":false,"assetalerts":{"CONT1000":"{\"tempalert\":\"above\"}"},"active":true,"timestamp":"2016-11-04 23:54:02.693662104 +0000 UTC"}
```
###Use the below to get the compliance history for a bill of lading
peer chaincode query -n comp -c '{"function":"readComplianceHistory", "args":["{\"blno\":\"10203040\"}"]}'
Query Result: 
```
{"comphistory":["{\"blno\":\"10203040\",\"type\":\"SHIPPING\",\"compliance\":false,\"assetalerts\":{\"CONT100\":\"{\\\"tempalert\\\":\\\"above\\\",\\\"dooralert\\\":true}\"},\"active\":true,\"timestamp\":\"2016-11-05 00:22:42.532116207 +0000 UTC\"}","{\"blno\":\"10203049\",\"type\":\"SHIPPING\",\"compliance\":false,\"assetalerts\":{\"CONT100\":\"{\\\"tempalert\\\":\\\"above\\\"}\"},\"active\":true,\"timestamp\":\"2016-11-05 00:21:53.569390543 +0000 UTC\"}","{\"blno\":\"10203049\",\"type\":\"SHIPPING\",\"compliance\":true,\"assetalerts\":null,\"active\":true,\"timestamp\":\"2016-11-05 00:20:51.93096927 +0000 UTC\"}"]}
```
##Update - updateContainerLogistics
peer chaincode invoke -n cont -c '{"function":"updateContainerLogistics", "args":["{\"containerno\":\"CONT1000\",\"location\":{\"latitude\":10, \"longitude\":9}, \"temperature\":41, \"carrier\":\"ARAMEX\", \"humidity\":20, \"light\":10, \"acceleration\":1, \"doorclosed\":true, \"airquality\":{\"oxygen\":1, \"carbondioxide\":1, \"ethylene\":1}}"]}'

peer chaincode query -n cont -c '{"function":"readContainerCurrentStatus", "args":["{\"containerno\":\"CONT1000\"}"]}'
##### (notice below that an alert is attached now to the container record)
Query Result: 
```
{"containerno":"CONT1000","blno":"10203040","location":{"latitude":10,"longitude":9},"carrier":"ARAMEX","timestamp":"2016-11-04 23:54:02.693662104 +0000 UTC","temperature":41,"humidity":20,"light":10,"acceleration":1,"doorclosed":true,"airquality":{"oxygen":1,"carbondioxide":1,"ethylene":1},"alerts":"{\"tempalert\":\"above\"}"}
```

###Other updateContainerLogistics examples

####Door Open:
peer chaincode invoke -n cont -c '{"function":"updateContainerLogistics", "args":["{\"containerno\":\"CONT1000\",\"location\":{\"latitude\":10, \"longitude\":9}, \"temperature\":41, \"carrier\":\"ARAMEX\", \"humidity\":20, \"light\":10, \"acceleration\":1, \"doorclosed\":false, \"airquality\":{\"oxygen\":1, \"carbondioxide\":1, \"ethylene\":1}}"]}'

####No compliance violations:
peer chaincode invoke -n cont -c '{"function":"updateContainerLogistics", "args":["{\"containerno\":\"CONT100\",\"location\":{\"latitude\":10, \"longitude\":9}, \"temperature\":4, \"carrier\":\"ARAMEX\", \"humidity\":20, \"light\":10, \"acceleration\":1, \"doorclosed\":true, \"airquality\":{\"oxygen\":1, \"carbondioxide\":1, \"ethylene\":1}}"]}'
##### In the above case, if you query the container record, it won't have an alert attached

