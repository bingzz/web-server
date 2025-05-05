package functions

import (
	"fmt"
	"net/http"
	"web-server/src/db"
	"web-server/src/schema"
	"web-server/src/utils"

	"github.com/gin-gonic/gin"
)

// Login user
func Login(c *gin.Context) {
	var httpResponse schema.HTTPResponse
	var userLogin schema.UserLogin

	// Bind JSON request body
	if err := c.BindJSON(&userLogin); err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to login: %s", err.Error()),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Validate user credentials
	if userLogin.Username == "" || userLogin.Password == "" {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Username and Password is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// sql query to the db
	sqlQuery := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	result, err := db.Request(sqlQuery, userLogin.Username)
	if err != nil { // Other errors
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		defer result.Close()
		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	defer result.Close() // Close the connection query

	var user schema.User // user from the db
	counter := 0
	for result.Next() {
		counter++
		err = result.Scan(&user.ID, &user.Username, &user.Password, &user.Name, &user.Status)
		if err != nil {
			httpResponse = schema.HTTPResponse{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}

			utils.ErrorLog(httpResponse)
			c.JSON(httpResponse.StatusCode, httpResponse)
			return
		}
	}

	// user not found
	if counter == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "User not found",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// incorrect password (bcrypt)
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
	// 	httpResponse = schema.HTTPResponse{
	// 		StatusCode: http.StatusUnauthorized,
	// 		Message:    "Incorrect password",
	// 	}

	// 	utils.ErrorLog(httpResponse)
	// 	c.JSON(httpResponse.StatusCode, httpResponse)
	// 	return
	// }

	// incorrect password (test)
	if user.Password != userLogin.Password {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Incorrect password",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// inactive status
	if !user.Status {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "User has been disabled",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully logged in",
		Data:       user,
	}

	c.JSON(httpResponse.StatusCode, httpResponse)
}
