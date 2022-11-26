package csv

import (
	"log"
	"os"
)

// TODO
var FILE_NAME = "pkg/csv/data/test-2021-05.csv"

func ProcessFile() {
	var fileName string
	if os.Getenv("ENVIROMENT") == "test" {
		fileName = "data/test-2021-05.csv"
	} else {
		fileName = FILE_NAME
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	journeys := validateFile(f)

	writeJourneyCSV(journeys)

}
