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
	defer rows.Close()

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

func AddJourney(newJourney *types.JourneyRequest) error {
	_, err := DB.Exec(fmt.Sprintf(insertJourney,
		newJourney.DepTime,
		newJourney.RetTime,
		newJourney.DepStationId,
		newJourney.RetStationId,
		newJourney.Distance,
		newJourney.Duration))

	if err != nil {
		return err
	}
	return nil
}
