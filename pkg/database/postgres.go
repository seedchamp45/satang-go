package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"github.com/seedchamp45/satang-go/config"
	"github.com/seedchamp45/satang-go/model"
	"gopkg.in/yaml.v3"
)

// PostgreSQLDB represents the PostgreSQL database connection
type PostgreSQLDB struct {
	*sql.DB
}

// NewPostgreSQLDB creates a new PostgreSQLDB instance and establishes a connection to the database
func NewPostgreSQLDB() (*PostgreSQLDB, error) {
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

	// Connect to the PostgreSQL server
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return &PostgreSQLDB{DB: db}, nil
}

// Close closes the database connection
func (db *PostgreSQLDB) Close() {
	db.DB.Close()
}

// Store the transaction data in the database
func (db *PostgreSQLDB) StoreTransaction(transaction model.Transaction) error {
	// Prepare the INSERT statement
	stmt, err := db.Prepare(`INSERT INTO transactions (hash, recipient, value)
		VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.Exec(
		transaction.Hash,
		transaction.To,
		transaction.Value.String(),
	)
	if err != nil {
		return err
	}

	return nil
}
