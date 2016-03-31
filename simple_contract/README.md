This is a simple contract (chaincode in IBM Blockchain language) for a basic Trade Lane scenario. 

Asset data (smart IoT devices or smart scanners sending asset data) comprising of a unique id for the asset, location (Latitude and Longitude), temperature and carrier comes in from IBM IoT Platform and gets stored in the blockchain. The contract can be used to manage multiple assets, one at a time. 

This contract has a JSON 4 compatible schema implementation that exposes the asset data structure and signature of functions supported by the contract for consumption by IoT platform's mapping component. 

Detailed documentation can be found [here] (https://www.ibm.com/developerworks/community/groups/service/html/communityoverview?communityUuid=cf921b26-ee7f-4c10-a445-89d51d432fa1#fullpageWidgetId=W6b3a7a280b13_46e4_94a7_cabb83baaa81)