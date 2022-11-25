package csv

import (
	"log"
	"os"
)

var FILE_NAME = "pkg/csv/data/2021-05.csv"

func ProcessFile() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	journeys := validateFile(f)
	writeJourneyCSV(journeys)
}
