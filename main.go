package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"example.com/web-service-gin/connect"
	"example.com/web-service-gin/model"
)

func main() {
	router := gin.Default()
	err := connect.InitDB()
	if err != nil {
		return
	}

	var albumObject model.Album
	router.GET("/albums", albumObject.GetAlbums)
	router.GET("/albums/:id", albumObject.GetAlbumByID)
	router.POST("/albums", albumObject.PostAlbums)
	router.PUT("/albums/:id", albumObject.PutAlbumsByID)
	router.DELETE("/albums/:id", albumObject.DeleteAlbumByID)
	err2 := router.Run("localhost:8080")
	if err2 != nil {
		fmt.Println("Error in starting server!")
		return
	}
}
