package main
import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"

//peer address group
//bankx
//banky
//regtech
//embassy
//customer


//custom data models
type Customer struct {
		ID						int 		`json:"id"`
        Salutation				String		`json:"Salutation"`
        FirstName				String		`json:"FirstName"`
        MiddleName				String		`json:"MiddleName"`
        LastName				String		`json:"LastName"`
        DateofBirth				String		`json:"DateofBirth"`
        SocialSecurityNumber	String		`json:"SocialSecurityNumber"`
        Gender					String		`json:"Gender"`
}
type Address struct {
		Customer ID
        Street Address 1
        Street Address 2
        City
        State
        Country
        Zip Code
        is it current? 
        Valid From
        Valid To
}

type ContactInformation struct {
		 Home Phone
        Work Phone
        Cellular Phone
        Email Address
}

type EmploymentInformation struct {
		 Customer ID
        Employee ID
        Company Name
        Company Address 1
        Company Address 2
        City
        State
        Country
        Zip Code
        Employment Start Date
        Employment End Date
        Annual Gross Salary
}

type BankAccounts struct {
		  Account Number
        Account Type
        Bank Name
        Bank Branch
        Bank Branch Address 1
        Bank Branch Address 2
        City
        State
        Country
        Zip Code
        Routing Number
        Bank Balance
}
type BankTransactions struct {
		Date
        Description
        Amount
        Credit/Debit
}
//as part of kyc or here
type Document struct {
		Photograph
    Signature
    Biometric
}

type KYC struct {
 kYCApplicationId int `json:"kYCApplicationId"`
 BankTransactions BankTransactions `json:"bankTransactions"`
 Document Document `json:"document"`
 
}


func CreateKYCApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 fmt.Println("Entering CreateKYCApplication")
 if len(args) < 2 {
 fmt.Println("Invalid number of args")
 //NEED TO FIX ARGUMENTS
 return nil, errors.New("Expected at least two arguments for KYC application creation")
 }
 var kycApplicationId = args[0]
 var kycApplicationInput = args[1]
 err := stub.PutState(kycApplicationId, []byte(kycApplicationInput))
 if err != nil {
 fmt.Println("Could not save kyc application to ledger", err)
 return nil, err
 }
 fmt.Println("Successfully saved kyc application")
 return nil, nil
}

func GetKYCApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 fmt.Println("Entering GetKYCApplication")
 if len(args) < 1 {
 fmt.Println("Invalid number of arguments")
 return nil, errors.New("Missing KYC application ID")
 }
 var kYCApplicationId = args[0]
 bytes, err := stub.GetState(kYCApplicationId)
 if err != nil {
 fmt.Println("Could not fetch kyc application with id "+kYCApplicationId+" from ledger", err)
 return nil, err
 }
 re
type SampleChaincode struct {
}
//called when block chain 1st executed
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 
 return nil, nil
}


//all read query here
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 
 return nil, nil

}
//all transaction update delete here
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte,
 error) {
 
 return nil, nil
}

func main() {
//starts the chain
 err := shim.Start(new(SampleChaincode))
 if err != nil {
 fmt.Println("Could not start SampleChaincode")
 } else {
 fmt.Println("SampleChaincode successfully started")
 }
}