package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"example.com/web-service-gin/api"
	"example.com/web-service-gin/connect"
)

func main() {

	port := os.Getenv("PORT")

	s := api.NewServer("localhost", port, "dev")

	router := gin.Default()
	err := connect.InitDB()
	if err != nil {
		return
	}

	router.GET("/albums", s.GetAlbums)
	router.GET("/albums/:id", s.GetAlbumByID)
	router.POST("/albums", s.PostAlbums)
	router.PUT("/albums/:id", s.PutAlbumsByID)
	router.DELETE("/albums/:id", s.DeleteAlbumByID)
	err2 := router.Run("localhost:" + s.Port)
	if err2 != nil {
		fmt.Println("Error in starting server!")
		return
	}
}
