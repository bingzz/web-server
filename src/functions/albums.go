package functions

import (
	"fmt"
	"net/http"
	"strconv"
	"web-server/src/db"
	"web-server/src/schema"
	"web-server/src/utils"

	"github.com/gin-gonic/gin"
)

// Basic CRUD using Arrays and Basic Loop

// Get the Albums from the array
func GetAlbums(c *gin.Context) {
	var httpResponse schema.HTTPResponse

	selectQuery := `SELECT a.id, a.title, b.id as author_id, b.name as author_name, a.price FROM albums a INNER JOIN authors b ON b.id = a.author`
	results, err := db.Request(selectQuery)
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		defer results.Close()
		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}
	defer results.Close()

	counter := 0
	var albums = make([]schema.Album, 0, 100) // list of albums
	fmt.Println(results)
	for results.Next() {
		counter++
		var album schema.Album // specific album
		err := results.Scan(&album.ID, &album.Title, &album.AuthorID, &album.AuthorName, &album.Price)
		if err != nil {
			httpResponse = schema.HTTPResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}

			utils.ErrorLog(httpResponse)
			c.JSON(httpResponse.StatusCode, httpResponse)
			return
		}

		albums = append(albums, album)
	}

	if counter == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "No albums",
			Data:       albums,
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Fetched Albums",
		Data:       albums,
	}

	c.JSON(httpResponse.StatusCode, httpResponse) // Serializes the response into JSON format
}

// Add a new Album
func AddAlbum(c *gin.Context) {
	var httpResponse schema.HTTPResponse
	var newAlbum schema.Album

	// Bind JSON to the new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    fmt.Sprintf("Failed to add Album: %s", err.Error()),
			// Data will be set to null if not provided
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Validate title
	if newAlbum.Title == "" {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Title is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Validate price
	if newAlbum.Price <= 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Price is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// Validate author
	if newAlbum.AuthorID == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Author is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	insertQuery := `INSERT INTO "albums" (title, author, price) VALUES ($1, $2, $3)`
	if _, err := db.Execute(insertQuery, newAlbum.Title, newAlbum.AuthorID, newAlbum.Price); err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusCreated,
		Message:    "New Album has been added",
	}

	// WARNING: we recommend using this only for development purposes since printing pretty JSON is more CPU and bandwidth consuming. Use Context.JSON() instead.
	c.JSON(httpResponse.StatusCode, httpResponse)
}

// Get a specific Album based on the ID
func GetSpecificAlbum(c *gin.Context) {
	var httpResponse schema.HTTPResponse

	// Get the ID from the request parameters URL
	id := c.Param("id")

	if id == "" {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Please search for a album",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	sqlQuery := `SELECT * FROM albums where id = $1`
	result, err := db.Request(sqlQuery, id)
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		defer result.Close()
		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}
	defer result.Close()

	var album schema.Album // album
	counter := 0
	for result.Next() {
		counter++
		err = result.Scan(&album.ID, &album.Title, &album.AuthorID, &album.Price)
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

	// album not found
	if counter == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "album not found",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "album found",
		Data:       album,
	}

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

	// validate album
	if updateAlbum.ID == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Album ID is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// validate album price
	if updateAlbum.Price <= 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Album ID is required",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	updateQuery := `UPDATE albums SET price = $1 WHERE id = $2`
	result, err := db.Execute(updateQuery, updateAlbum.Price, updateAlbum.ID)
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// check the number of rows affected
	ra, err := result.RowsAffected()
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// row not updated
	if ra == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Album failed to update: Not found",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Album has been updated",
	}

	c.JSON(httpResponse.StatusCode, httpResponse)
}

// Delete a specific Album based on the ID
func DeleteAlbum(c *gin.Context) {
	var httpResponse schema.HTTPResponse

	id := c.Param("id")

	// validate parameter
	if id == "" {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusNotFound,
			Message:    "Please search for an album",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// validate if id is an int
	if _, err := strconv.Atoi(id); err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid album ID",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	deleteQuery := `DELETE FROM albums WHERE id = $1`
	result, err := db.Execute(deleteQuery, id)
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	// check the number of rows affected
	ra, err := result.RowsAffected()
	if err != nil {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
	}

	// row not updated
	if ra == 0 {
		httpResponse = schema.HTTPResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Album failed to remove: Not found",
		}

		utils.ErrorLog(httpResponse)
		c.JSON(httpResponse.StatusCode, httpResponse)
		return
	}

	httpResponse = schema.HTTPResponse{
		StatusCode: http.StatusOK,
		Message:    "Album successfully removed",
	}

	c.JSON(httpResponse.StatusCode, httpResponse)
}
