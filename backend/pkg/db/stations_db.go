package db

import (
	"citybike/pkg/types"
	"database/sql"
	"fmt"
)

func GetStations(page int, limit int) ([]*types.Station, error) {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}
	rows, err := DB.Query(fmt.Sprintf(getStationsQuery, from, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []*types.Station

	for rows.Next() {
		station := scanStation(rows)

		counts, err := DB.Query(fmt.Sprintf(getRetAndDepCountQuery,
			station.ID, station.ID))
		if err != nil {
			return nil, err
		}
		defer counts.Close()

		counts.Next()
		counts.Scan(
			&station.Returns,
			&station.Departures)
		stations = append(stations, &station)
	}

	return stations, nil
}

func GetStationsByName(page int, limit int, name string) ([]*types.Station, error) {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}

	rows, err := DB.Query(fmt.Sprintf(getStationsByNameQuery, name, from, limit))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stations []*types.Station

	for rows.Next() {
		station := scanStation(rows)

		counts, err := DB.Query(fmt.Sprintf(getRetAndDepCountQuery,
			station.ID, station.ID))
		if err != nil {
			return nil, err
		}
		defer counts.Close()

		counts.Next()
		counts.Scan(
			&station.Returns,
			&station.Departures)

		stations = append(stations, &station)
	}

	return stations, nil
}

func AddStation(newStation *types.StationRequest) error {
	_, err := DB.Exec(fmt.Sprintf(
		insertStation,
		newStation.ID,
		newStation.Name,
		newStation.Address,
		newStation.City,
		newStation.Operator,
		newStation.Capacity,
		newStation.Coordinates.X,
		newStation.Coordinates.Y))
	if err != nil {
		return err
	}
	return nil
}

func scanStation(rows *sql.Rows) types.Station {
	var station types.Station
	var coordinates types.Coordinates
	rows.Scan(
		&station.FID,
		&station.ID,
		&station.Name,
		&station.Address,
		&station.City,
		&station.Operator,
		&station.Capacity,
		&coordinates.X,
		&coordinates.Y,
	)
	station.Coordinates = coordinates

	return station
}
