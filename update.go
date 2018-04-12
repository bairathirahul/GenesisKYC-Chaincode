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
			jsonResp := "{\"Error\":\"Failed to get state for " + args[2] + "\"}"
			return nil, errors.New(jsonResp)
		}
		customer := Customer{}
		json.Unmarshal(customerJSONAsBytes, &customer)
		fmt.Println("CHAINCODE: After Unmarshalling customer")
		
switch {
	case info == "BasicInfo":
		
		basicInfo := BasicInfo{}
		json.Unmarshal([]byte(args[3]), &basicInfo)
		fmt.Println("CHAINCODE: After Unmarshalling basicInfo")
		
		customer.basicInfo=basicInfo
		fmt.Println("CHAINCODE: Writing customer back to ledger")
		jsonAsBytes, _ := json.Marshal(customer)
		err = stub.PutState(customer.id, jsonAsBytes)
		if err != nil {
			return nil, err
		}

		

		

	case info == "Address":
		addresses := Address{}
		json.Unmarshal([]byte(args[3]), &addresses)
		fmt.Println("CHAINCODE: After Unmarshalling basicInfo")
	
		if(oper=="add") {
		
			v.addresses=append(customer.addresses, addresses...)
		} else {
			v.addresses=addresses
		}
		
		fmt.Println("CHAINCODE: Writing customer back to ledger")
		jsonAsBytes, _ := json.Marshal(customer)
		err = stub.PutState(customer.id, jsonAsBytes)
		if err != nil {
			return nil, err
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
	return nil, nil
	//carAsBytes, _ := stub.GetState(args[0])
	//return shim.Success(carAsBytes)
}