# Quick set-up 

## Use Case

Track cash machine data  - asset id, ActionType (Initial Balance, Deposit or Withdraw), Amount (amount of a particular transaction), Balance (total balance of the Cash Machine) and Timestamp. On create, an initial balance is set for the machine. On update, Deposit or Withdraw action take place and the balance is updated. The current state of the machine, retrieved on read, will have the asset id, the last Action type, the amount for the action type, the current balance and the last action's timestamp. The contract will track history and will return history in descending order. The functions are createCashMachine, updateCashMachine, readCashMachine, readCashMachineHistory and deleteCashMachine. Get monitoring ui to work with the same.

The ask was to be able to run the code on a Sandbox environment, so the instructions below are around how to achieve it. This code is compatible with a newer version of Hyperledger than the one in BlueMix today. Once BlueMix is updated, it can be run on BlueMix.

> Assumption : Hyperledger set up and built. 

Download the cashMachine folder and save it in the go code path . For example: /Users/spnair/go/src/github.com/hyperledger/fabric/examples/chaincode/go

## Local Machine

LOCALDEVDIR will be something similar to /Users/spnair/go/src/github.com/hyperledger/fabric/devenv
Its not mandatory that this be set. 

### 1. Navigate to the folder with the Vagrant file (LOCALDEVDIR). Bring up vagrant

``` bash

FORWARD_DOCKER_PORTS='true' vagrant up

```

## Inside Vagrant Ubuntu

There are two key folders you will need to use: the folder where the peer executable is, and the chain code folder. For convenience, you could add these as environment variables in bashrc

``` bash

vi ~/.bashrc
Towards the end of the file, add the following:
#path to chaincode folder - could vary based on the path where the code is saved
CCPATH="$GOPATH/src/github.com/hyperledger/fabric/examples/chaincode/go"
export CCPATH
#path to peer
PRPATH="$GOPATH/src/github.com/hyperledger/fabric/peer"
export PRPATH

```

The above is not mandatory. The following steps _do not assume_ that this has been done, instead, it shows the full path of the concerned folders.

### You will need **3 terminal windows**

Vagrant needs to be running on all windows. (don't forget to vagrant ssh)

### 2. Navigate to the ‘peer’ folder

``` bash

cd $GOPATH/src/github.com/hyperledger/fabric/peer OR cd $PRPATH (If you set the bashprofile)

```

### 3. Start the peer in dev mode

``` bash

./peer node start --peer-chaincodedev

```

### 4. In the second tab, navigate to the folder where you have the chaincode. for example,

``` bash

cd /opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go/cashMachine 
OR 
cd $CCPATH/cashMachine (If you set the bashprofile)

```

### 5. build the chaincode

``` bash

go build

```

If there are build errors, they will be crearly reported and listed with line number, so it should be easy to identify and fix them.

### 6. register

In Sandbox environment, we can register the chain code with a name we specify. This step applies __only to sandbox__. In DevNet, the name would be assigned (long 128 b alphanumeric string) when you deploy the chain code

``` bash

CORE_CHAINCODE_ID_NAME=cash CORE_PEER_ADDRESS=0.0.0.0:30303 ./cashMachine

```

### 7. On the third terminal window, run the chaincode commands from the peer folder

``` bash

cd $GOPATH/src/github.com/hyperledger/fabric/peer 
OR 
cd $PRPATH (If you set the bashprofile)

```

#### a.  deploy 

``` bash

./peer chaincode deploy -n cash  -c '{"Function":"Init", "Args": ["{\"version\":\"1.0\"}"]}'

```

This returns the name of the chaincode - in this case, 'cash' In a DevNet environment, the new chaincode id would be created and returned.

#### b. Read a sample to see what data should be sent in

``` bash

./peer chaincode query -n cash -c '{"Function":"readCashMachineSamples", "Args":[]}'

```

This returns:

``` json

{
    "event": {
        "assetID": "The ID of a managed asset. In this case, the cash machine's unique id wrt monetary transactions.For query operations, only assetID needs to be sent in.",
        "ActionType": "One of three actions are expected: InitialBalance, Deposit or Withdraw",
        "Amount": "The amount that needs to be transacted. eg. 123.05"
        "Timestamp": "A string with timestamp. If not sent in, it is set to the transaction time in the fabric"
    },
    "initEvent": {
        "version": "The version number of the contract. This version expects 1.0"
    },
    "state": {
        "assetID": "String with The ID of a managed asset. In this case, the cash machine's unique id wrt monetary transactions.",
        "ActionType": "A String with one of three values is expected: InitialBalance, Deposit or Withdraw",
        "Amount": "The amount that needs to be transacted. eg. 123.05"
        "Balance": "This is a computed field. Don't send it in, it will be overwritten. eg. 234.56"
        "Timestamp": "A string with timestamp. If not sent in, it is set to the transaction time in the fabric"
    }
}

```

The 'event' section indicates what needs to be sent in for an invoke or query call. The 'initEvent' section tells you what to send in while initializing (deploying) the chaincode. The 'state' section tells you what to expect as a response to query calls (the 'read' calls - except for readCashMachineSamples which is explained here)

#### c. invoke : createCashMachine - create an id for the cashmachine with the initial balance

``` bash

./peer chaincode invoke -n cash  -c '{"Function":"createCashMachine", "Args":["{\"assetID\":\"CM01\",\"actiontype\":\"InitialBalance\", \"amount\":1000, \"timestamp\":\"2016-06-07 10:01:01 UTC\"}"]}'

```

#### d. query: readCashMachine - check the most recent status of the machine

``` bash

./peer chaincode query -n cash   -c '{"Function":"readCashMachine", "Args":["{\"assetID\":\"CM01\"}"]}'

```

#### e. invoke : updateCashMachine - Update the cash machine balance - operations could be "Deposit" or "Withdraw". Right now it assumes that if the ActionType is not a 'Deposit", it mush be a withdrawal.

``` bash

./peer chaincode invoke -n cash -c '{"Function":"updateCashMachine", "Args":["{\"assetID\":\"CM01\",\"actiontype\":\"Deposit\",  \"amount\":10}"]}'
OR
./peer chaincode invoke -n cash -c '{"Function":"updateCashMachine", "Args":["{\"assetID\":\"CM01\",\"actiontype\":\"Withdraw\",  \"amount\":20}"]}'

```

#### f. query : readCashMachineHistory - see all transactions so far by Asset ID.

``` bash

./peer chaincode query -n cash -c '{"Function":"readCashMachineHistory", "Args":["{\"assetID\":\"CM01\"}"]}'

```

This returns all the transactions in descending order

``` json

{"cashhistory":["{\"assetid\":\"CM01\",\"actiontype\":\"Withdraw\",\"amount\":20,\"balance\":990,\"timestamp\":\"2016-06-07 13:24:31.752701888 +0000 UTC\"}","{\"assetid\":\"CM01\",\"actiontype\":\"Deposit\",\"amount\":10,\"balance\":1010,\"timestamp\":\"2016-06-07 13:23:58.336522709 +0000 UTC\"}","{\"assetid\":\"CM01\",\"actiontype\":\"InitialBalance\",\"amount\":1000,\"balance\":1000,\"timestamp\":\"2016-06-07 13:22:52.732470093 +0000 UTC\"}"]}

```

#### g.invoke: deleteCashMachine - delete the asset record

``` bash

./peer chaincode invoke -n cash -c '{"Function":"deleteCashMachine", "Args":["{\"assetID\":\"CM01\"}"]}'

```

If you run the readCashMachine command for the same asset now, it will return an error
_Error: Error querying chaincode: rpc error: code = 2 desc = "Error:Transaction or query returned with failure: Unable to get asset state from ledger"_

## Testing on Postman

Test scripts are in this folder as 'CashMachine.postman_collection'

## Testing with the monitoring UI

https://www.youtube.com/watch?v=p8SSAnuMZec
https://www.youtube.com/watch?v=cPuIgOygxzo
https://www.youtube.com/watch?v=30jE2xLlq_k

The changes mentioned in the videos are implemented in the `cashmachine_UI` as derived from the `monitoring_UI` in this folder.