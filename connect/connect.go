package connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:root@12345@tcp(127.0.0.1:3306)/music")
	if err != nil {
		fmt.Println("Can't connect to Database: ", err)
	} else {
		fmt.Println("DB Connected successfully")
	}

}
