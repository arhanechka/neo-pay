package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configurator interface {
	NewConfiguraion() Configuration
}

//from config.json
type Configuration struct {
	NodeURI     string
	WaitTimeSec int
}

// new struct for mocking
type MockConfiguration struct{}

//func for configuration substitution
func (*MockConfiguration) NewConfiguraion() *Configuration {
	var configuration Configuration
	configuration.NodeURI = "http://localhost:10332"
	return &configuration
}

func NewConfiguraion() *Configuration {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("something wrong with TestConfiguration ", err)
		os.Exit(3)
	}
	return &configuration
}
