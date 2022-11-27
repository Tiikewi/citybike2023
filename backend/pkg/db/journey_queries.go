package db

import (
	"citybike/pkg/types"
	"log"
)

func GetJourneys(page int, limit int) []*types.Journey {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}
	rows, err := DOT.Query(DB, "get-journeys", from, limit)
	if err != nil {
		log.Fatal(err)
	}

	var journeys []*types.Journey

	for rows.Next() {
		var journey types.Journey
		rows.Scan(
			&journey.ID,
			&journey.DepTime,
			&journey.RetTime,
			&journey.DepStationId,
			&journey.DepStationName,
			&journey.RetStationId,
			&journey.RetStationName,
			&journey.Distance,
			&journey.Duration)
		journeys = append(journeys, &journey)
	}

	return journeys
}

func GetJourneyCount() int {

	row, err := DOT.QueryRow(DB, "get-journey-count")
	if err != nil {
		log.Fatal(err)
	}

	var count int

	if err := row.Scan(&count); err != nil {
		panic("Error when fetching amount of journeys!")
	}

	return count
}
