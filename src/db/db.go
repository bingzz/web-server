package db

import (
	"database/sql"
	"fmt"
	"net/http"
	"web-server/src/schema"
	"web-server/src/utils"
	"web-server/src/validators"

	_ "github.com/lib/pq"
)

var db *sql.DB

// PostgreSQL Connect
func InitializeDB() {
	fmt.Println("Connecting to Database...")
	dbUser, dbPass, dbHost, dbPort, dbName := validators.ValidateDBEnv()

	// Connection String
	psqlConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	// Open DB
	var err error
	db, err = sql.Open("postgres", psqlConn)
	if err != nil {
		response := schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to connect database: %s", err.Error()),
		}

		utils.ErrorLog(response)
		panic(err.Error())
	}

	// Close DB
	// defer db.Close()

	// Ping DB
	if err = db.Ping(); err != nil {
		defer db.Close()
		response := schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    fmt.Sprintf("Failed to ping database: %s", err.Error()),
		}

		utils.ErrorLog(response)
		panic(err.Error())
	}

	fmt.Println("Database successfully connected")
}

// Function to execute a query (POST, PUT, DELETE, PATCH)
func Execute(querystring string, arguments ...interface{}) (sql.Result, error) {
	// Verify if database is connected
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	// Execute query
	result, err := db.Exec(querystring, arguments...)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return result, err
}

// Function to request (GET)
func Request(querystring string, arguments ...interface{}) (*sql.Rows, error) {
	// Verify if database is connected
	if db == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	// Call query
	rows, err := db.Query(querystring, arguments...)

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	// defer rows.Close()
	return rows, err
}

// Postgre ORM??
