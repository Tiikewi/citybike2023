package db

import (
	"log"
)

func CreateTables() {
	_, err := DB.Exec(createJourneyTableQuery)
	if err != nil {
		log.Fatal("error when creating journey table. Error: ", err)
	}
	_, err = DB.Exec(createStationTableQuery)
	if err != nil {
		log.Fatal("error when creating station table. Error: ", err)
	}
}
