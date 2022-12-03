package db

import (
	"citybike/pkg/types"
	"fmt"
)

const (
	DepartureName = iota + 1
	DepartureID
	ReturnName
	ReturnID
	Distance
	Duration
)

func GetJourneys(page int, limit int, sort int) ([]*types.Journey, error) {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}

	// Determine sort
	var sortStr string
	switch sort {
	case DepartureName:
		sortStr = "departure_station"
	case DepartureID:
		sortStr = "departure_station_id"
	case ReturnName:
		sortStr = "return_station"
	case Distance:
		sortStr = "distance"
	case Duration:
		sortStr = "duration"
	default:
		sortStr = "departure_station"
	}

	rows, err := DB.Query(fmt.Sprintf(getJourneysQuery, sortStr, from, limit))
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

// func GetJourneyCount() (int, error) {

// 	row, err := DOT.QueryRow(DB, "get-journey-count")
// 	if err != nil {
// 		return -1, err
// 	}

// 	var count int

// 	if err := row.Scan(&count); err != nil {
// 		panic("Error when fetching amount of journeys!")
// 	}

// 	return count, nil
// }
