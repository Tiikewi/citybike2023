package csv

import (
	"citybike/pkg/types"
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

func validateTime(timeString string) bool {
	// check if times are in correct format
	_, err := time.Parse("2006-01-02T15:04:05", timeString)
	return err == nil
}

// convert strings to int
func convertInts(numbers []string) ([]int, error) {
	var ints []int

	for _, num := range numbers {
		if kk, err := strconv.Atoi(num); err != nil {
			return ints, err
		} else if kk < 0 {
			return ints, errors.New("negative integer")
		} else {
			ints = append(ints, kk)
		}
	}

	return ints, nil
}

func writeJourneyCSV(journeys []*types.Journey) {
	env := os.Getenv("ENVIROMENT")
	var path string
	if env == "test" {
		path = "data/test-valitated-journeys.csv"
	} else {
		path = "pkg/csv/data/valitated-journeys.csv"
	}
	csvFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	var strJourney [][]string
	for _, journey := range journeys {
		strJourney = append(strJourney, []string{
			journey.DepTime,
			journey.RetTime,
			strconv.Itoa(journey.DepStationId),
			journey.DepStationName,
			strconv.Itoa(journey.RetStationId),
			journey.RetStationName,
			strconv.Itoa(journey.Distance),
			strconv.Itoa(journey.Duration),
		})
	}

	writer := csv.NewWriter(csvFile)
	err = writer.WriteAll(strJourney)
	if err != nil {
		log.Fatal("Error when writing new csv")
	}
}
