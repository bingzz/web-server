package validators

import (
	"fmt"
	"net/http"
	"web-server/src/constants"
	"web-server/src/schema"
	"web-server/src/utils"

	"github.com/gofor-little/env"
)

// Get the ENV file data
func ValidateEnv() (string, string, string) {
	// Validate if .env exists in the directory
	if err := env.Load(".env"); err != nil {
		panic("No ENV provided")
	}

	// Validate if ENV credentials are available
	environment := checkEnvVar("ENVIRONMENT")
	host := checkEnvVar("URL")
	port := checkEnvVar("PORT")

	fmt.Println("Environment Set...")
	return environment, host, port
}

// Get the ENV Database Data
func ValidateDBEnv() (string, string, string, string, string) {
	// Validate if .env exists in the directory
	if err := env.Load(".env"); err != nil {
		panic("No ENV provided")
	}

	// Validate if DB variables is provided
	dbUser := checkEnvVar("DBUSER")
	dbPass := checkEnvVar("DBPASSWORD")
	dbHost := checkEnvVar("DBHOST")
	dbPort := checkEnvVar("DBPORT")
	dbName := checkEnvVar("DBNAME")

	return dbUser, dbPass, dbHost, dbPort, dbName
}

// Function to check a specific ENV variable
func checkEnvVar(vbl string) string {
	envVar, err := env.MustGet(vbl)

	// Log an error if there are missing variables
	if err != nil {

		response := schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		utils.ErrorLog(response)
		panic(fmt.Sprintf("%sMissing ENV variable: \"%s\"%s", constants.Red, vbl, constants.Reset))
	}

	return envVar
}
