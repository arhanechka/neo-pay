package main

import (
	"log"

	"github.com/CityOfZion/neo-go-sdk/neo"
	"github.com/CityOfZion/neo-go-sdk/neo/models"
	//"fmt"
	//"os"
	//"encoding/json"
	//"strconv"
)

// type ClientServise interface {
// 	GetCurrentBlockIndex(configuration *Configuration) int64
// 	GetBlockByIndex(index int64) *models.Block
// 	initClient(configuration *Configuration)
// }

// type MockClient struct{}

// func (m *MockClient) MockGetCurrentBlockIndex(configuration *Configuration) int64 {
// 	return 1840533
// }

// func (m *MockClient) MockGetBlockByIndex(int64) {

// }

// func (m *MockClient) MockInitClient(*Configuration) {

// }

var Client neo.Client

func GetCurrentBlockIndex(configuration *Configuration) int64 {
	if Client.NodeURI == "" {
		initClient(configuration)
	}
	currentBlockIndex, err := Client.GetBlockCount()
	if err != nil {
		log.Println(err)
		return -1
	}
	return currentBlockIndex - 1
}

func GetBlockByIndex(index int64) *models.Block {
	block, err := Client.GetBlockByIndex(index)
	if err != nil {
		log.Println(err)
		return nil
	}
	return block
}
func initClient(configuration *Configuration) {
	Client = neo.NewClient(configuration.NodeURI)
	ok := Client.Ping()
	if !ok {
		log.Fatal("Unable to connect to NEO node")
	}
}
