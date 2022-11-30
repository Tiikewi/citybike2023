package db

import (
	"citybike/pkg/types"
	"fmt"
	"log"
)

func GetStations(page int, limit int) []*types.Station {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}
	rows, err := DOT.Query(DB, "get-stations", from, limit)
	if err != nil {
		log.Fatal(err)
	}

	var stations []*types.Station

	for rows.Next() {
		var station types.Station
		var coordinates types.Coordinates
		rows.Scan(
			&station.FID,
			&station.ID,
			&station.NameFinnish,
			&station.NameSwedish,
			&station.NameEnglish,
			&station.Address,
			&station.AddressSwedish,
			&station.City,
			&station.CitySwedish,
			&station.Operator,
			&station.Capacity,
			&coordinates.X,
			&coordinates.Y,
		)
		station.Coordinates = coordinates
		stations = append(stations, &station)
	}

	return stations
}

func GetStationsByName(page int, limit int, name string) []*types.Station {
	var from int
	if page == 1 {
		from = 0
	} else {
		from = (page * limit) - limit
	}

	// % for searching substrings.
	rows, err := DOT.Query(DB, "get-stations-by-name", name+"%", from, limit)
	if err != nil {
		log.Fatal(err)
	}

	var stations []*types.Station

	for rows.Next() {
		var station types.Station
		var coordinates types.Coordinates
		rows.Scan(
			&station.FID,
			&station.ID,
			&station.NameFinnish,
			&station.NameSwedish,
			&station.NameEnglish,
			&station.Address,
			&station.AddressSwedish,
			&station.City,
			&station.CitySwedish,
			&station.Operator,
			&station.Capacity,
			&coordinates.X,
			&coordinates.Y,
		)
		fmt.Println("STATION: ")
		station.Coordinates = coordinates
		stations = append(stations, &station)
	}

	return stations
}
