package functions

import (
	"fmt"
	"net/http"
	"web-server/src/data"
	"web-server/src/schema"
	"web-server/src/utils"

	"github.com/gin-gonic/gin"
)

// Basic CRUD using Arrays and Basic Loop

// Get the Albums from the array
func GetAlbums(c *gin.Context) {
	var httpResponse schema.HTTPResponse
	var errMsg string

	// Check if album list is empty
	if len(data.Albums) == 0 {
		errMsg = "No Albums"
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    errMsg,
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Fetched Albums",
		Data:       data.Albums,
	}

	c.JSON(httpResponse.StatusCode, httpResponse) // Serializes the response into JSON format
}

// Add a new Album
func AddAlbum(c *gin.Context) {
	var newAlbum schema.Album
	var httpResponse schema.HTTPResponse
	var errMsg string

	// Bind JSON to the new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		errMsg = fmt.Sprintf("Failed to add Album: %s", err.Error())
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
			// Data will be set to null if not provided
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Add new Album to the slice (not recommended this is not thread safe)
	data.DataStore.Mu.Lock()
	data.Albums = append(data.Albums, newAlbum) // similar to Array.push() in javascript
	data.DataStore.Mu.Unlock()

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusCreated,
		Message:    "New Album has been added",
		Data:       newAlbum,
	}

	// WARNING: we recommend using this only for development purposes since printing pretty JSON is more CPU and bandwidth consuming. Use Context.JSON() instead.
	c.JSON(httpResponse.StatusCode, httpResponse)
}

// Get a specific Album based on the ID
func GetSpecificAlbum(c *gin.Context) {
	var httpResponse schema.HTTPResponse
	var errMsg string

	// Get the ID from the request parameters URL
	id := c.Param("id")

	// Find the Album in the array
	for _, a := range data.Albums {
		if a.ID == id {
			httpResponse = schema.HTTPResponse{
				StatusCode: http.StatusFound,
				Message:    "Album Found",
				Data:       a,
			}

			// Send response if found
			c.JSON(httpResponse.StatusCode, httpResponse)
			return
		}
	}

	// Album is not found
	errMsg = "Album not found"
	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusNotFound,
		Message:    errMsg,
	}

	utils.ErrorLog(httpResponse)
	c.JSON(httpResponse.StatusCode, httpResponse)
}

// Update the price of the specific Album
func UpdateAlbum(c *gin.Context) {
	var updateAlbum schema.Album
	var httpResponse schema.HTTPResponse
	var errMsg string

	if err := c.BindJSON(&updateAlbum); err != nil {
		errMsg = fmt.Sprintf("Failed to update Album: %s", err.Error())
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errMsg,
			// Data will be set to null if not provided
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Find the Album from the list
	for i, a := range data.Albums {
		if a.ID == updateAlbum.ID {
			// Update the price of this album
			data.Albums[i].Price = updateAlbum.Price

			httpResponse = schema.HTTPResponse{
				StatusCode: http.StatusOK,
				Message:    "Successfully updated Album Price",
				Data:       data.Albums,
			}

			c.JSON(httpResponse.StatusCode, httpResponse)
			return
		}
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusNotFound,
		Message:    "Album not found",
	}

	utils.ErrorLog(httpResponse)
	c.JSON(httpResponse.StatusCode, httpResponse)
}

// Delete a specific Album based on the ID
func DeleteAlbum(c *gin.Context) {
	var httpResponse schema.HTTPResponse
	var index int = -1

	// Get the ID from the request parameters URL
	id := c.Param("id")

	// Find the Album from the list
	for i, a := range data.Albums {
		if a.ID == id {
			index = i
		}
	}

	// If the album is not found, send the client "not found"
	if index == -1 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Album not found",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Delete the specific album
	data.DataStore.Mu.Lock()
	data.Albums = append(data.Albums[:index], data.Albums[index+1:]...) // Method of removing an item from an array
	data.DataStore.Mu.Unlock()

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Album Removed",
		Data:       data.Albums,
	}

	c.JSON(httpResponse.StatusCode, httpResponse)
}
