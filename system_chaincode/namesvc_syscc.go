import (
    "errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//************* init *******************
//The name service is a system contract and is initialized when the peer comes up
// Parameters are set in the core.yaml. 
// Note: Clean the db and rebuild. The system chaincode, in its curent avatar
// lives on the genesis block
// This is being implemented as a key-value pair. At its core, it is just that.

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var key, val string    // Entities

	if len(args) != 2 {
		return nil, errors.New("need 2 args (key and a value).")
	}

	// Initialize the chaincode
	key = args[0]
	val = args[1]
	// Write the state to the ledger
	err := stub.PutState(key, []byte(val))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// As a system chaincode can easily be accessed by all participants of the system, a generic key-value storage makes sense. 
// It could be used as a naming service, or something else too.
// A wrapper will be needed to resolve the name, and metadata can be stored there as needed.

func (t *SampleSysCC) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	
	// The objective of this chaincode is to provide a simplistic name service. 
	// Once a contract is deployed, the name has to be explicitly registered with this system chaincode
	if function !="setKey" {
		return nil, errors.New("Invoke function 'setKey' expected")
	}
	if len(args) != 2 {
		return nil, errors.New("need 2 args (key and a value).")
	}

	// Initialize the chaincode
	key = args[0]
	val = args[1]

	_, err := stub.GetState(key)
	if err != nil {
		// This implies that the key does not exist in the stub. 
		err = stub.PutState(key, []byte(val))
			if err != nil {
				return nil, errors.New("Unable to put the new key-value pair to the stub")
			}
	}else {
		// The key already exists in the state
		return nil, errors.New("Key already exists. Please use a different one.")
	}

	return nil, nil
}

func (t *SampleSysCC) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("need 1 arg - the key .")
	}	
	key = args[0]	
	val, err := stub.GetState(key)
	if err != nil {
		erMsg := "{\"Error\":\"Failed to get val for " + key + "\"}"
		return nil, errors.New(erMsg)
	}
	if val == nil {
		erMsg = "{\"Error\":\"Nil val for " + key + "\"}"
		return nil, errors.New(erMsg)
	}

	return val, nil
	
