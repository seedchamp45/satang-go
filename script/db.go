package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"github.com/seedchamp45/satang-go/config"
	"gopkg.in/yaml.v3"
)

func main() {

	// Read the YAML file
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the YAML content
	var config config.DatabaseConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the database credentials
	dbHost := config.Database.Host
	dbPort := config.Database.Port
	dbName := config.Database.Name
	dbUser := config.Database.User
	dbPassword := config.Database.Password
	dbPGName := config.Database.DBName

	// Use the retrieved credentials as needed
	fmt.Println("Host:", dbHost)
	fmt.Println("Port:", dbPort)
	fmt.Println("Name:", dbName)
	fmt.Println("User:", dbUser)
	fmt.Println("Password:", dbPassword)

	// Connect to the PostgreSQL server
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", dbHost, dbPort, dbPGName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the database
	err = createDatabase(db, dbName)
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the newly created database
	db, err = sql.Open("postgres", fmt.Sprintf("%s dbname=%s", psqlInfo, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the transactions table
	err = createTransactionsTable(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database and table initialized successfully")
}

// createDatabase creates the PostgreSQL database
func createDatabase(db *sql.DB, dbName string) error {
	createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	fmt.Println(createDBQuery)
	fmt.Println(db)
	_, err := db.Exec(createDBQuery)
	if err != nil {
		return fmt.Errorf("failed to create database: %v", err)
	}
	return nil
}

// createTransactionsTable creates the transactions table in the database
func createTransactionsTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			hash VARCHAR(255),
			recipient VARCHAR(255),
			value VARCHAR(255),
			created_at TIMESTAMPTZ DEFAULT NOW()
		);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create transactions table: %v", err)
	}

	return nil
}
