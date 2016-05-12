1. go build in the simple_contract_hyperledger folder

2. Start peer: ./peer peer --peer-chaincodedev

3. Register to the shim in the samdbox, inside the cotnract folder: CORE_CHAINCODE_ID_NAME=simhyp CORE_PEER_ADDRESS=0.0.0.0:30303 ./simple_contract_hyperledger

4. Run the tests (from  /opt/gopath/src/github.com/hyperledger/fabric/peer)

a. deploy:

./peer chaincode deploy -n simhyp -c '{"Function":"Init", "Args": ["{\"version\":\"1.0\"}"]}'

b. query schema: 

./peer chaincode query -l golang -n simhyp  -c '{"Function":"readAssetSchemas", "Args":[]}'

c. create / update and read (invoke and query calls)

create:
./peer chaincode invoke -l golang -n simhyp -c '{"Function":"createAsset", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":34, \"longitude\":23},  \"carrier\":\"ABCCARRIER\"}"]}'

read:
./peer chaincode query -l golang -n simhyp -c '{"Function":"readAsset", "Args":["{\"assetID\":\"CON123\"}"]}'

update:
./peer chaincode invoke -l golang -n simhyp -c '{"Function":"updateAsset", "Args":["{\"assetID\":\"CON123\", \"location\":{\"latitude\":34, \"longitude\":25}},\"temperature\":14,  \"carrier\":\"ABCCARRIER\"}"]}'

A new call to 'read' should give you the updated asset record.

delete:
./peer chaincode invoke -l golang -n simhyp -c '{"Function":"deleteAsset" , "Args":["{\"assetID\":\"CON123\"}"]}' 





