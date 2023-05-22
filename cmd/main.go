package main

import (
	"io/ioutil"
	"log"

	"github.com/seedchamp45/satang-go/config"
	"github.com/seedchamp45/satang-go/pkg/database"
	"github.com/seedchamp45/satang-go/pkg/ethereum"
	"gopkg.in/yaml.v3"
)

func main() {
	// Initialize the database connection
	db, err := database.NewPostgreSQLDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Initialize the Ethereum client
	ethClient, err := ethereum.NewEthereumClient()
	if err != nil {
		log.Fatal("Failed to connect to the Ethereum client:", err)
	}

	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the YAML content
	var addresses config.Addresses
	err = yaml.Unmarshal(yamlFile, &addresses)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the database credentials
	address := addresses.Address
	log.Println(address)
	// Start monitoring transactions
	ethereum.StartTransactionMonitoring(ethClient, db, address)
}
