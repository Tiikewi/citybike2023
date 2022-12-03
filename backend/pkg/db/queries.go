package db

const getJourneysQuery = `SELECT j.*, s1.station_name_finnish AS departure_station, 
	s2.station_name_finnish AS return_station FROM journey j
	LEFT OUTER JOIN station s1 ON j.departure_station_id = s1.id
	LEFT OUTER JOIN station s2 ON return_station_id = s2.id
	ORDER BY %s
	LIMIT %d, %d`

const getStationsQuery = `SELECT fid, id, station_name_finnish, 
	address_finnish, city_name_finnish, 
	operator, capacity, x_coordinate, y_coordinate
	FROM station LIMIT %d,%d;`

const getStationsByNameQuery = `SELECT fid, id, station_name_finnish,
address_finnish, city_name_finnish,
operator, capacity, x_coordinate,
y_coordinate 
FROM station 
WHERE station_name_finnish 
LIKE "%s%%" LIMIT %d,%d;`
