package db

// CREATE TABLES

const createJourneyTableQuery = `
CREATE TABLE IF NOT EXISTS journey (
	id                     INT auto_increment NOT NULL PRIMARY KEY,
	departure_time         DATETIME NOT NULL,
	return_time            DATETIME NOT NULL,
    departure_station_id   INT NOT NULL,
    return_station_id      INT NOT NULL,
    distance               INT NOT NULL,
    duration               INT NOT NULL,

    FOREIGN KEY(departure_station_id)
    REFERENCES station(id),
    FOREIGN KEY(return_station_id)
    REFERENCES station(id)
)`

const createStationTableQuery = `
CREATE TABLE IF NOT EXISTS station (
	fid                    INT DEFAULT 0,
    id                     INT NOT NULL PRIMARY KEY,
    station_name_finnish   VARCHAR(100) NOT NULL UNIQUE,
    station_name_swedish   VARCHAR(100),
    station_name_english   VARCHAR(100),
    address_finnish        VARCHAR(100) NOT NULL,
    address_swedish        VARCHAR(100),
    city_name_finnish      VARCHAR(100) NOT NULL,
    city_name_swedish      VARCHAR(100),
    operator               VARCHAR(100),
    capacity               INT,
    x_coordinate           FLOAT NOT NULL,
    y_coordinate           FLOAT NOT NULL
)`

// JOURNEYS

const getJourneysQuery = `SELECT j.*, s1.station_name_finnish AS departure_station, 
	s2.station_name_finnish AS return_station FROM journey j
	LEFT OUTER JOIN station s1 ON j.departure_station_id = s1.id
	LEFT OUTER JOIN station s2 ON return_station_id = s2.id
	ORDER BY %s
	LIMIT %d, %d`

// STATIONS
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

const getRetAndDepCountQuery = `SELECT (
    SELECT COUNT(*)
    FROM   journey j
    WHERE j.return_station_id = %d
) AS tot_returns,
(
SELECT  (
    SELECT COUNT(*)
    FROM   journey j
    WHERE j.departure_station_id = %d) 
) AS tot_departures`

const insertStation = `INSERT INTO station 
    (id, station_name_finnish, address_finnish, 
    city_name_finnish, operator, capacity, 
    x_coordinate, y_coordinate)
    VALUES (%d, '%s', '%s', '%s', '%s', %d, %f, %f);`

const insertJourney = `INSERT INTO journey 
    (departure_time, return_time, departure_station_id, 
    return_station_id, distance, duration)
    VALUES ('%s', '%s', %d, %d, %d, %d);`
