package db

import (
	"citybike/pkg/types"
)

func GetJourneys(page int, limit int) ([]*types.Journey, error) {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}
	rows, err := DOT.Query(DB, "get-journeys", from, limit)
	if err != nil {
		return nil, err
	}

	var journeys []*types.Journey
	for rows.Next() {
		var journey types.Journey
		rows.Scan(
			&journey.ID,
			&journey.DepTime,
			&journey.RetTime,
			&journey.DepStationId,
			&journey.RetStationId,
			&journey.Distance,
			&journey.Duration,
			&journey.DepStationName,
			&journey.RetStationName)
		journeys = append(journeys, &journey)
	}

	return journeys, nil
}

func GetJourneyCount() (int, error) {

	row, err := DOT.QueryRow(DB, "get-journey-count")
	if err != nil {
		return -1, err
	}

	var count int

	if err := row.Scan(&count); err != nil {
		panic("Error when fetching amount of journeys!")
	}

	return count, nil
}
