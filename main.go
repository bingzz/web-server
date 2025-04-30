package main

import (
	"fmt"
	"web-server/src/constants"
	"web-server/src/router"
	"web-server/src/validators"
	// Alias with import location
	// This is using the package name from this specific file
)

func main() {
	// ENV
	environment, host, port := validators.ValidateEnv()

	listener := router.InitializeAPI()

	fmt.Printf("%s======\nEnvironment: %s\nHost: %s:%s\n======\n\n%s", constants.Green, environment, host, port, constants.Reset)
	listener.Run(fmt.Sprintf("%s:%s", host, port))
}
