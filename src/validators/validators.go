package validators

import (
	"fmt"

	"github.com/gofor-little/env"
)

// Get the ENV file data
func ValidateEnv() (string, string, string) {
	// Validate if .env exists in the directory
	if err := env.Load(".env"); err != nil {
		panic("No ENV provided")
	}

	// Validate if the environment has been provided
	environment, err := env.MustGet("ENVIRONMENT")
	if err != nil {
		panic("No Environment selected")
	}

	// Validate Environment type
	switch environment {
	case "DEV", "STAGING", "PRODUCTION":
	default:
		panic("No Environment Selected")
	}

	// Validate if hostname is provided
	host, err := env.MustGet("URL")
	if err != nil {
		panic("No Host provided")
	}

	// Validate if the PORT is provided
	port, err := env.MustGet("PORT")
	if err != nil {
		panic("No Port provided")
	}

	fmt.Println("Environment Set...")
	return environment, host, port
}
