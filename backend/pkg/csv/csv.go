package csv

import (
	"bufio"
	"citybike/pkg/types"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var FILE_NAME = "pkg/csv/2021-05.csv"
var LINE_LENGTH = 8

func ProcessFile() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	journeys := validateFile(f)

	// TODO something with journeys *
	fmt.Println("JOURNEYS: ", journeys)
}

func validateFile(file *os.File) []*types.Journey {
	s := bufio.NewScanner(file)
	s.Scan() // Skips header row
	journeys := []*types.Journey{}

	// store invalid lines for more info and possible future use..
	var invalidLines []string

	for s.Scan() {
		line := strings.Trim(s.Text(), " ")
		lineArray := strings.Split(line, ",")
		var journey types.Journey

		if len(lineArray) < LINE_LENGTH || len(lineArray) > LINE_LENGTH {
			invalidLines = append(invalidLines, line)
			continue
		}

		// Indexes 2, 4, 6 and 7 should be integers.
		var stringInts []string
		stringInts = append(stringInts, lineArray[2], lineArray[4], lineArray[6], lineArray[7])

		ints, err := convertInts(stringInts)
		if err != nil {
			invalidLines = append(invalidLines, line)
			continue
		}

		journey.DepStationId = ints[0]
		journey.RetStationId = ints[1]
		journey.Distance = ints[2]
		journey.Duration = ints[3]

		journey.DepTime = lineArray[0]
		journey.RetTime = lineArray[1]
		journey.DepStationName = lineArray[3]
		journey.RetStationName = lineArray[5]

		// don't get journey if it lasts less than 10 second
		// or is less than 10 meters.
		if journey.Distance < 10 || journey.Duration < 10 {
			invalidLines = append(invalidLines, line)
			continue
		}

		if !validateTime(journey.DepTime) && validateTime(journey.RetTime) {
			invalidLines = append(invalidLines, line)
			continue
		}

		journeys = append(journeys, &journey)
	}
	fmt.Println("Omitted, invalid lines: ", len(invalidLines))
	return journeys
}

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
