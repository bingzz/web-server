package utils

import (
	"fmt"
	"time"
)

// Get the date in YYYY-MM-DD format
func YYYYMMDD() string {
	currentDateTime := time.Now()

	var year int = currentDateTime.Year()

	// Use padding of 2 digits
	var month string = fmt.Sprintf("%02d", int(currentDateTime.Month()))
	var day string = fmt.Sprintf("%02d", currentDateTime.Day())

	return fmt.Sprintf("%d-%s-%s", year, month, day)
}
