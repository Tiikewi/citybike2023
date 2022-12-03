package db

const getJourneysQuery = `SELECT j.*, s1.station_name_finnish AS departure_station, 
	s2.station_name_finnish AS return_station FROM journey j
	LEFT OUTER JOIN station s1 ON j.departure_station_id = s1.id
	LEFT OUTER JOIN station s2 ON return_station_id = s2.id
	ORDER BY %s
	LIMIT %d, %d`
