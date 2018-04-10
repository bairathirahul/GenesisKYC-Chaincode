package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//peer address group
//bankx
//banky
//regtech
//embassy
//customer

// Chaincode Interface
type GenesisChainCode struct{}

// Data Models
type BasicInfo struct {
	salutation  string
	firstName   string
	middleName  string
	lastName    string
	dateofBirth uint64
	gender      string
	ssn         string
	verified    bool
}

type Address struct {
	street1  string
	street2  string
	city     string
	state    string
	country  string
	zip      string
	verified bool
	active   bool
}

type Contact struct {
	contactType  string
	phoneNumber  string
	emailAddress string
	verified     bool
	active       bool
}

type Employment struct {
	employmentType string
	companyName    string
	street1        string
	street2        string
	city           string
	state          string
	country        string
	zip            string
	designation    string
	startDate      uint64
	endDate        uint64
	isCurrent      bool
	grossSalary    int
	verified       bool
	active         bool
}

type BankAccounts struct {
	accountNo      string
	bankName       string
	bankBranchName string
	street1        string
	street2        string
	city           string
	state          string
	country        string
	zip            string
	active         bool
}

type CustomerDocument struct {
	documentType string
	documentId   string
	active       bool
}

type BankTransaction struct {
	transactionId   string
	transactionDate uint32
	transactionType string
	description     string
	amount          int
}

type Customer struct {
	id               string
	basicInfo        BasicInfo
	addresses        []Address
	contacts         []Contact
	documents        []CustomerDocument
	bankAccounts     []BankAccounts
	bankTransactions []BankTransaction
}

// =====================================
// Main
// =====================================
func main() {
	err := shim.Start(new(GenesisChainCode))
	if err != nil {
		fmt.Printf("Error occurred in starting chaincode: %s", err)
	} else {
		fmt.Printf("Chaincode started successfully")
	}
}

// =====================================
// Initializes Chaincode
// =====================================
func (t *GenesisChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// =====================================
// Invoke Chaincode method
// =====================================
func (t *GenesisChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()

	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "registerCustomer" {
		return t.registerCustomer(stub, args)
	} else if function == "queryCustomer" {
		return t.queryCustomer(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// =====================================
// Register Customer
// This method will be executed when the customer first registers in the platform
// It will generate a Unique ID for the customer and fills Basic Information,
// Current Address and Personal Contact Information. These fields are mandatory
// ones on the registration form
// =====================================
func (t *GenesisChainCode) registerCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Customer basic information, address and contact must be provided for registration")
	}

	var basicInfo BasicInfo
	var address Address
	var contact Contact

	json.Unmarshal([]byte(args[0]), &basicInfo)
	json.Unmarshal([]byte(args[1]), &address)
	json.Unmarshal([]byte(args[2]), &contact)

	customer := Customer{basicInfo: basicInfo, addresses: []Address{address}, contacts: []Contact{contact}}
	customerAsBytes, _ := json.Marshal(customer)
	stub.PutState(args[0], customerAsBytes)

	return shim.Success(nil)
}

// =====================================
// Update Customer
// This method will update the data of the customer. It must be made flexible
// to accomodate any type of modification, instead of making separate method for
// each modification type
// =====================================
func (t *GenesisChainCode) updateCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	carAsBytes, _ := stub.GetState(args[0])
	return shim.Success(carAsBytes)
}

// =====================================
// Query Customer
// This method will return the data of the customer. Apart from the customer ID
// the type of requested information must also be provided. It will only return
// the type of information requested
// =====================================
func (t *GenesisChainCode) queryCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}
