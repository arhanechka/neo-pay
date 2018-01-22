package main

import (
	"testing"
)

// new struct for mocking
type MockConfiguration struct{}

//func for configuration substitution
func (*MockConfiguration) NewConfiguraion() *Configuration {
	var configuration Configuration
	configuration.NodeURI = "http://localhost:20332"
	return &configuration
}

//declaration of structure for mocking configuration
var mockConfiguration = MockConfiguration{}

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
