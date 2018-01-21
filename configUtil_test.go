package main

import (
	"testing"
)

var mockConfiguration = &MockConfiguration{}
var TestConfiguration = mockConfiguration.NewConfiguraion()

func TestNewConfiguraion(t *testing.T) {
	configuratin := TestConfiguration
	expectedURI := configuratin.NodeURI
	actualConfiguration := NewConfiguraion()

	actualURI := actualConfiguration.NodeURI
	if actualURI != expectedURI {
		t.Errorf("initConfig returned unexpected confiuration: got %v want %v", actualURI, expectedURI)
	}

}
