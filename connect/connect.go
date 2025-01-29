package connect

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func InitDB() error {
	var err error
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
