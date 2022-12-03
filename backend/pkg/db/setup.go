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

func CreateTables() {
	DOT = getQueryFile()

	_, err := DB.Exec(createJourneyTableQuery)
	if err != nil {
		log.Fatal("error when creating journey table. Error: ", err)
	}
	_, err = DB.Exec(createStationTableQuery)
	if err != nil {
		log.Fatal("error when creating station table. Error: ", err)
	}
}
