package main

import (
	"fmt"
	"testing"
	//"fmt"
)

// new struct for mocking
type MockCustomer struct{}

//declaration of structure for mocking configuration
var mockCustomer = MockCustomer{}

//using method for configuration substitution
//var TestCustomer = mockCustomer.CreateCustomer(TestConfiguration)

//func for customer substitution
// func (*MockCustomer) CreateCustomer(configuration *Configuration) Customer {
// 	var customer Customer
// 	_assignedAddress, _ := mockCustomer.GetNewAddress(TestConfiguration)
// 	customer.AssignedAddress = _assignedAddress
// 	fmt.Println("getter" + customer.AssignedAddress)
// 	customer.Deposit = 0
// 	customer.StartBlock = 182117
// 	customer.StatusPaid = false
// 	return customer
// }
func (*MockCustomer) GetNewAddress(configuration *Configuration) (string, error) {
	address := "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"
	return address, nil
}

func TestGetNewAddress(t *testing.T) {
	expectedLength := len("AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4")

	actualAddress, err := mockCustomer.GetNewAddress(TestConfiguration) //here we use mock of configuration

	fmt.Println("actualAddress" + actualAddress)
	if err != nil {
		t.Errorf("GetNewAddress returned error: %v", err)
	}
	actualLength := len(actualAddress)
	if actualLength != expectedLength {
		t.Errorf("GetNewAddress returned unexpected NEO Address: got %v want %v", actualLength, expectedLength)
	}
}

func TestCreateCustomer(t *testing.T) {

	var expectedBalance int64 = 0
	expectedStartBlock := GetCurrentBlockIndex(TestConfiguration)
	expectedStatusPaid := false
	//expectedAddress := "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"

	//here we use only mock of configuration, because we need the actual CurrentBlock for comparing
	actualCustomer := CreateCustomer(TestConfiguration, &mockCustomer)
	fmt.Println(actualCustomer.AssignedAddress)
	if actualCustomer.Deposit != expectedBalance {
		t.Errorf("CreateCustomer returned unexpected customer object balance : got %v want %v", actualCustomer.Deposit, expectedBalance)
	}
	if actualCustomer.StartBlock != expectedStartBlock {
		t.Errorf("CreateCustomer returned unexpected customer object startBlock : got %+v want %+v", actualCustomer.StartBlock, expectedStartBlock)
	}
	if actualCustomer.StatusPaid != expectedStatusPaid {
		t.Errorf("CreateCustomer returned unexpected customer object statusPaid : got %v want %v", actualCustomer.StatusPaid, expectedStatusPaid)
	}

	//if actualCustomer.AssignedAddress != expectedAddress {
	//	t.Errorf("CreateCustomer returned unexpected customer object AssignedAddress : got %v want %v", actualCustomer.AssignedAddress, expectedAddress)
	//}
	//}

	//func mock_GetNewAddress() (string, error) {
	//	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4", nil
}
