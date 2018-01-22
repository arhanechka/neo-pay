package main

import (
	"testing"
	//"fmt"
)

//declaration of structure for mocking configuration
var mockCustomer = &MockCustomer{}

//using method for configuration substitution
var TestCustomer = mockCustomer.CreateCustomer(*TestConfiguration)

func TestGetNewAddress(t *testing.T) {
	expectedLength := len("AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4")

	actualAddress, err := GetNewAddress(TestConfiguration) //here we use mock of configuration
	if err != nil {
		t.Errorf("GetNewAddress returned error: %v", err)
	}
	actualLength := len(actualAddress)
	if actualLength != expectedLength {
		t.Errorf("GetNewAddress returned unexpected NEO Address: got %v want %v", actualLength, expectedLength)
	}
}

func TestCreateCustomer(t *testing.T) {

	//var expectedBalance int64 = 0
	expectedStartBlock := GetCurrentBlockIndex(TestConfiguration)
	//expectedStatusPaid := false
	//expectedAddress := "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4"

	//here we use only mock of configuration, because we need the actual CurrentBlock for comparing
	actualCustomer := CreateCustomer(TestConfiguration)

	if actualCustomer.Deposit != TestCustomer.Deposit {
		t.Errorf("CreateCustomer returned unexpected customer object balance : got %v want %v", actualCustomer.Deposit, expectedBalance)
	}
	if actualCustomer.StartBlock != expectedStartBlock {
		t.Errorf("CreateCustomer returned unexpected customer object startBlock : got %+v want %+v", actualCustomer.StartBlock, expectedStartBlock)
	}
	if actualCustomer.StatusPaid != TestCustomer.StatusPaid {
		t.Errorf("CreateCustomer returned unexpected customer object statusPaid : got %v want %v", actualCustomer.StatusPaid, expectedStatusPaid)
	}

	//if actualCustomer.AssignedAddress != expectedAddress {
	//	t.Errorf("CreateCustomer returned unexpected customer object AssignedAddress : got %v want %v", actualCustomer.AssignedAddress, expectedAddress)
	//}
}

//func mock_GetNewAddress() (string, error) {
//	return "AcbUNbdFMdYLBronyM3cHBzi49WKEwJWD4", nil
//}
