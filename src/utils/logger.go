package utils

import (
	"fmt"
	"os"
	"time"
	"web-server/src/constants"
	"web-server/src/schema"
)

// Logs an error in the logs folder
func ErrorLog(errMsg schema.HTTPResponse) {

	var fileDate string = YYYYMMDD()
	var fileName string = fmt.Sprintf("src/logs/%s.txt", fileDate)

	// Write Error message
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // Creates a file if file does not exist, appends the message to the existing file
	if err != nil {
		panic(err)
	}

	dt := time.Now().Format(time.RFC1123) // Datetime of the error occurred
	if _, err := f.Write([]byte(fmt.Sprintf("%s => %+v\n\n", dt, errMsg))); err != nil {
		panic(err)
	}

	fmt.Printf("%sError: %s => %+v\n%s", constants.Red, dt, errMsg, constants.Reset)
	if err := f.Close(); err != nil {
		panic(err)
	}
}
