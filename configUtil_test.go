package main

import (
	"testing"
)

//declaration of structure for mocking configuration
var mockConfiguration = &MockConfiguration{}

//using method for configuration substitution
var TestConfiguration = mockConfiguration.NewConfiguraion()

func TestNewConfiguraion(t *testing.T) {
	expectedURI := TestConfiguration.NodeURI
	actualConfiguration := NewConfiguraion()

	actualURI := actualConfiguration.NodeURI
	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}
