package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandler(t *testing.T) {
	// Create a request to pass to handler
	r, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new responserecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)

	// Handlers satisfy http.Handler, so can call their ServeHTTP method directly to pass in the request and responserecorder.
	handler.ServeHTTP(rr, r)

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	expectedCustomer := TestCustomer
	var testCustomer Customer

	error := json.Unmarshal(body, &testCustomer)
	if error != nil {
		t.Errorf("HandlerFunc(StatusHandler) returned wrong JSON structure ", error)
	}

	if rr.Code != http.StatusOK {
		t.Errorf("HandlerFunc(StatusHandler) returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	if len(testCustomer.AssignedAddress) != len(expectedCustomer.AssignedAddress) {
		t.Errorf("HandlerFunc(StatusHandler) returned wrong customer AssignedAddress length: got %v want %v", len(testCustomer.AssignedAddress), len(expectedCustomer.AssignedAddress))
	}

}

func TestNewAddressHandler(t *testing.T) {
	// Create a request to pass to handler
	r, err := http.NewRequest("GET", "/address", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new responserecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewAddressHandler)

	// Handlers satisfy http.Handler, so can call their ServeHTTP method directly to pass in the request and responserecorder.
	handler.ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		t.Errorf("HandlerFunc(NewAddressHandler) returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
}
