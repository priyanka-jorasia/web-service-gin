package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"example.com/web-service-gin/connect"
	"example.com/web-service-gin/model"
)

func main() {
	router := gin.Default()
	connect.InitDB()
	var albumObject model.Album
	router.GET("/albums", albumObject.GetAlbums)
	router.GET("/albums/:id", albumObject.GetAlbumByID)  //200 OK
	router.POST("/albums", albumObject.PostAlbums)       // 201 Created
	router.PUT("/albums/:id", albumObject.PutAlbumsByID) // Update album
	router.DELETE("/albums/:id", albumObject.DeleteAlbumByID)
	router.Run("localhost:8080")
}
