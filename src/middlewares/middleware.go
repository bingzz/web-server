package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error Handler requests
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// defer makes the function run at the end
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				fmt.Println(err)
				// utils.ErrorLog(err)

				// Send an Error Response back to the Client
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()

		// Proceed to the next function request
		c.Next()

	}
}
