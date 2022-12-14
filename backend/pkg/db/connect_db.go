package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var (
	dbContainerName = os.Getenv("MYSQL_CONTAINER")
	user            = os.Getenv("MYSQL_USER")
	password        = os.Getenv("MYSQL_ROOT_PASSWORD")
	dbname          = os.Getenv("MYSQL_DATABASE")
)

func ConnectToDB() {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dbContainerName, dbname)

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal("Error when trying to connect to DB. Err:", err)
	}

	DB = db

	CreateTables()

}
