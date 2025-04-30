package router

import (
	"web-server/src/functions"
	"web-server/src/middlewares"

	"github.com/gin-gonic/gin"
)

func InitializeAPI() *gin.Engine {
	// Create Gin Router
	router := gin.Default()

	// Set middlewares
	router.Use(middlewares.ErrorHandler())
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Set Endpoints
	router.GET("/albums", functions.GetAlbums)
	router.POST("/newAlbum", functions.AddAlbum)
	router.GET("/albums/:id", functions.GetSpecificAlbum)
	router.PATCH("/album", functions.UpdateAlbum)
	router.DELETE("/albums/:id", functions.DeleteAlbum)

	return router
}
