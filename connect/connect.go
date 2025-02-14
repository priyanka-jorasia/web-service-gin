package connect

import (
	"database/sql"
	"errors"
	"fmt"

	"example.com/web-service-gin/model"
)

var Db *sql.DB
var err error

func InitDB() error {
	Db, err = sql.Open("mysql", "root:root@12345@tcp(127.0.0.1:3306)/music")
	if err != nil {
		fmt.Println("Could not connect to the database error in the dnc: ", err)
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Could not connect to the database (Db.Ping()),err")
		return err
	}
	fmt.Println("DB connected successfully")
	return nil
}

func GetAllAlbums() ([]model.Album, error) {
	var fetchedAlbums []model.Album
	query := "Select * from music.albums;"
	var rows *sql.Rows
	rows, err = Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var newAlbum model.Album

		err = rows.Scan(&newAlbum.ID, &newAlbum.Title, &newAlbum.Artist, &newAlbum.Price)
		if err != nil {
			return nil, err
		}
		fetchedAlbums = append(fetchedAlbums, newAlbum)
	}
	return fetchedAlbums, nil
}

func GetAlbumById(id string) (model.Album, error) {
	query := "Select * from music.albums where Album_ID=?;"
	var fetchedAlbum model.Album
	row := Db.QueryRow(query, id)
	err = row.Scan(&fetchedAlbum.ID, &fetchedAlbum.Title, &fetchedAlbum.Artist, &fetchedAlbum.Price)
	if err != nil {
		return fetchedAlbum, err
	}

	return fetchedAlbum, nil
}

func AddAlbum(album model.Album) error {
	query := "insert into music.albums values(?,?,?,?);"
	result, err := Db.Exec(query, album.ID, album.Title, album.Artist, album.Price)

	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows were added")
	}
	return nil
}

func UpdateAlbum(id string, album model.Album) error {
	query := "Update music.albums set Title=?,Artist=?,Price=? where Album_ID=?;"

	result, err := Db.Exec(query, album.Title, album.Artist, album.Price, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		err = errors.New("no rows were added")
		fmt.Println("Error,", err)
		return err
	}
	return nil
}

func DeleteAlbum(id string) error {
	query := "Delete from music.albums where Album_ID=?;"

	result, err := Db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		err = errors.New("ID not found")
		fmt.Println("Error,", err)
		return err
	}
	return nil
}
