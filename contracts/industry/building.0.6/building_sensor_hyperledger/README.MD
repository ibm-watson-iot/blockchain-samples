# A Hyperledger Sample Contract with Building sensor data scenario

> **This smart contract is based on the tradelane contract /  IoT Blockchain Demo Smart Contract at 

https://github.com/ibm-watson-iot/blockchain-samples/tree/master/trade_lane_contract_hyperledger, as on July 1, 2016.**

Please refer to the documentation for the *trade_lane-contract_hyperledger* (link above) to understand the features offered by the contract. This Readme file only touches upon the differences

1. Since multiple asset types are being handled simultaneously, the contract uses "assetid"_"type" as the key, instead of just the asset id.

2. assettype is not sent in at present, instead it is inferred from the asset name.

3. The following alert scenarios exist:

   a) If creation time is later than modification time, it raises `"CREATE_TIME_GREATER_THAN_MODIFY_TIME"` and

   b) For the variable speed motor, if the rpm is less than 20% of the maximum rpm, it implies that the motor is likely to over-heat. So it raises a `"RPM_LESS_THAN_20PERCENT"` alert. These alerts are mere examples intended to illustrate alert scenarios for the devices used with this contract
   
4. The schemas and samples are unchanged from the Tradelane scenario. So this example as it stands today is not completely compatible with the monitoring_ui, if you are using it.(You can still see the blocks, run readAllAssets etc). The schemas and sample generation approach is being overhauled and this code may get updated once that work is completed. There is no projected timeline on when this change will be in place. Users are welcome to tweak these to get them working to their satisfaction.

Note: This is sample WIP code and can change. Please use and modify your copy as per your discretion.
