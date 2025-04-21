package main

import (
	"fmt"
	"net/http"
	"web-server/config/routes"

	// Alias with import location
	// This is using the package name from this specific file

	"github.com/gofor-little/env"
)

func main() {
	// ENV
	environment, host, port := validateEnv()

	// CORS
	// server.SetServerConfiguration()

	// Middleware
	// middleware.SetMiddleware()

	// API endpoints
	routes.InitializeAPI()

	// Logger

	// Listener
	// Check if the server can be started
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		fmt.Println("Failed to Start Server")
		panic(err.Error())
	}

	fmt.Printf("======\nEnvironment: %s\nHost: %s:%s\n======\n\n", environment, host, port)
}

// Get the ENV file data
func validateEnv() (string, string, string) { // func fn() (return arguments)

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
