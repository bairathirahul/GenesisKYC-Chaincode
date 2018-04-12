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
	Salutation  string
	FirstName   string
	MiddleName  string
	LastName    string
	DateOfBirth uint64
	Gender      string
	SSN         string
	Verified    bool
}

type Address struct {
	Street1  string
	Street2  string
	City     string
	State    string
	Country  string
	Zip      string
	Verified bool
	Active   bool
}

type Contact struct {
	ContactType  string
	PhoneNumber  string
	EmailAddress string
	Verified     bool
	Active       bool
}

type Employment struct {
	EmploymentType string
	CompanyName    string
	Street1        string
	Street2        string
	City           string
	State          string
	Country        string
	Zip            string
	Designation    string
	StartDate      uint64
	EndDate        uint64
	IsCurrent      bool
	GrossSalary    int
	Verified       bool
	Active         bool
}

type BankAccounts struct {
	AccountNo      string
	BankName       string
	BankBranchName string
	Street1        string
	Street2        string
	City           string
	State          string
	Country        string
	Zip            string
	Active         bool
}

type CustomerDocument struct {
	DocumentType string
	DocumentId   string
	Active       bool
}

type BankTransaction struct {
	TransactionID   string
	TransactionDate uint32
	TransactionType string
	Description     string
	Amount          int
}

type Customer struct {
	ID               string
	BasicInfo        BasicInfo
	Addresses        []Address
	Contacts         []Contact
	Documents        []CustomerDocument
	BankAccounts     []BankAccounts
	BankTransactions []BankTransaction
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
	} else if function == "updateCustomer" {
		return t.updateCustomer(stub, args)
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

	if len(args) != 4 {
		return shim.Error("Customer basic information, address and contact must be provided for registration")
	}

	basicInfo := BasicInfo{}
	address := Address{}
	contact := Contact{}

	json.Unmarshal([]byte(args[1]), &basicInfo)
	json.Unmarshal([]byte(args[2]), &address)
	json.Unmarshal([]byte(args[3]), &contact)

	customer := Customer{BasicInfo: basicInfo, Addresses: []Address{address}, Contacts: []Contact{contact}}
	customerAsBytes, _ := json.Marshal(customer)

	fmt.Println(string(customerAsBytes))

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

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	info := args[0]
	operation := args[1]

	customerJSONAsBytes, err := stub.GetState(args[2])
	if err != nil {
		//jsonResp := "{\"Error\":\"Failed to get state for " + args[2] + "\"}"
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	customer := Customer{}
	json.Unmarshal(customerJSONAsBytes, &customer)
	fmt.Println("CHAINCODE: After Unmarshalling customer")

	switch {
	case info == "BasicInfo":

		basicInfo := BasicInfo{}
		json.Unmarshal([]byte(args[3]), &basicInfo)
		fmt.Println("CHAINCODE: After Unmarshalling basicInfo")

		customer.BasicInfo = basicInfo
		fmt.Println("CHAINCODE: Writing customer back to ledger")
		jsonAsBytes, _ := json.Marshal(customer)
		err = stub.PutState(args[2], jsonAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	case info == "Address":
		addresses := []Address{}
		json.Unmarshal([]byte(args[3]), &addresses)
		fmt.Println("CHAINCODE: After Unmarshalling basicInfo")

		if operation == "add" {

			customer.Addresses = append(customer.Addresses, addresses...)
		} else {
			//customer.addresses = addresses
		}

		fmt.Println("CHAINCODE: Writing customer back to ledger")
		jsonAsBytes, _ := json.Marshal(customer)
		err = stub.PutState(customer.ID, jsonAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}

	case info == "Contact":
		fmt.Println("Good afternoon.")
	case info == "BankAccounts":
		fmt.Println("Good afternoon.")
	case info == "CustomerDocument":
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	return shim.Success(nil)
}

// =====================================
// Query Customer
// This method will return the data of the customer. Apart from the customer ID
// the type of requested information must also be provided. It will only return
// the type of information requested
// =====================================
func (t *GenesisChainCode) queryCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	customerJSONAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println(string(customerJSONAsBytes))
	return shim.Success(nil)
}
