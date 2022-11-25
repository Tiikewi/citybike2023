package db

import (
	"log"
	"os"

	"github.com/qustavo/dotsql"
)

var DOT *dotsql.DotSql

func getQueryFile() *dotsql.DotSql {
	// Queries in separete file for ease of maintenance.
	// get enviroment, because when testing, path changes.
	env := os.Getenv("ENVIROMENT")
	var path string
	if env == "dev" {
		path = "pkg/db/queries.sql"
	} else {
		path = "../db/queries.sql"
	}

	dot, err := dotsql.LoadFromFile(path)
	if err != nil {
		log.Fatal("Cannot load queries.sql")
	}

	return dot
}

// func GetVersion() string {
// 	dot := getQueryFile()

// 	row, err := dot.QueryRow(DB, "get-version")
// 	if err != nil {
// 		log.Fatal("Error when trying to connect to DB. Err:", err)
// 	}
// 	var version string
// 	err = row.Scan(&version)
// 	if err != nil {
// 		log.Fatal("Error: ", err)
// 	}

// 	return version
// }

func CreateTables() {
	DOT = getQueryFile()

	_, err := DOT.Exec(DB, "create-journey-table")
	if err != nil {
		log.Fatal("error when creating journey table. Error: ", err)
	}
}
