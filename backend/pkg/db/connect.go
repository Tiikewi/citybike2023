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
	// host            = os.Getenv("MYSQL_HOST")
	dbContainerName = "mariadb"
	user            = os.Getenv("MYSQL_USER")
	password        = os.Getenv("MYSQL_PASSWORD")
	dbname          = os.Getenv("MYSQL_DATABASE")
)

func ConnectToDB() {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, dbContainerName, dbname)

	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Fatal("Error when trying to connect to DB. Err:", err)
	}

	DB = db

	// Print version of mariadb.
	fmt.Println("Connected to MariaDB version: ", GetVersion())
}
