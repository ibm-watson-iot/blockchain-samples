package main

var samples = `
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
}`
