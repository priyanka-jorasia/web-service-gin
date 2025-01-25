package model

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"example.com/web-service-gin/connect"
)

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func (a Album) GetAlbums(c *gin.Context) {

	rows, err := connect.Db.Query("Select * from music.albums")
	if err != nil {
		fmt.Println("Error in executing Select * Query (GET):", err)
	}
	var albums []Album

	for rows.Next() {
		var id string
		var title string
		var artist string
		var price float64

		err = rows.Scan(&id, &title, &artist, &price)
		if err != nil {
			fmt.Println("Error in scanning databse rows:", err)
			return
		}
		a.ID = id
		a.Title = title
		a.Artist = artist
		a.Price = price

		albums = append(albums, a)

	}
	c.IndentedJSON(http.StatusOK, albums)
}

func (a Album) GetAlbumByID(c *gin.Context) {

	id := c.Param("id")

	query := "Select * from music.albums where Album_ID =" + id + ";"

	row := connect.Db.QueryRow(query)

	var title string
	var artist string
	var price float64

	err := row.Scan(&id, &title, &artist, &price)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"response": "ID not found"})
		return
	}
	a.ID = id
	a.Title = title
	a.Artist = artist
	a.Price = price
	fmt.Println("Fetched Album", a)
	c.IndentedJSON(http.StatusOK, a)
}

func (a Album) PostAlbums(c *gin.Context) {

	err := c.BindJSON(&a) //insert JSON value into newAlbum (a) object
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	query := "Insert into music.albums values (?,?,?,?);"
	_, err = connect.Db.Exec(query, a.ID, a.Title, a.Artist, a.Price)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Input"})
	} else {
		c.IndentedJSON(http.StatusCreated, a)
	}
}

func (a Album) PutAlbumsByID(c *gin.Context) { // PUT = remove the values which are not available
	newId := c.Param("id")

	err := c.BindJSON(&a)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid input", "Error": err})
		return
	}
	var emptyString string
	if a.ID != emptyString {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please provide ID as path parameter only"})
		return
	}

	query := "update music.albums set Title=?,Artist=?,Price=? where Album_ID=?;"
	result, err := connect.Db.Exec(query, a.Title, a.Artist, a.Price, newId)

	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Error in executing query:"})
		return
	}

	noOfRowsAffected, _ := result.RowsAffected()
	if noOfRowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID not found"})
		return

	}
	a.ID = newId
	c.IndentedJSON(http.StatusOK, a)
}

func (a Album) DeleteAlbumByID(c *gin.Context) {
	newId := c.Param("id")
	query := "Delete from music.albums where Album_ID=?"

	result, err := connect.Db.Exec(query, newId)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "Can't execute Delete query", "Error": err})
	}

	noOfRowsAffected, _ := result.RowsAffected()
	if noOfRowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Record Deleted with ID=" + newId})
}
