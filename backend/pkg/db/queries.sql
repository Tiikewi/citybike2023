-- name: get-version
SELECT VERSION()

-- name: create-journey-table
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
);

-- name: create-station-table
CREATE TABLE IF NOT EXISTS station (
	fid                    INT NOT NULL,
    id                     INT NOT NULL PRIMARY KEY,
    station_name_finnish   VARCHAR(100) NOT NULL,
    station_name_swedish   VARCHAR(100) NOT NULL,
    station_name_english   VARCHAR(100) NOT NULL,
    address_finnish        VARCHAR(100) NOT NULL,
    address_swedish        VARCHAR(100) NOT NULL,
    city_name_finnish      VARCHAR(100) NOT NULL,
    city_name_swedish      VARCHAR(100) NOT NULL,
    operator               VARCHAR(100),
    capacity               INT NOT NULL,
    x_coordinate           FLOAT NOT NULL,
    y_coordinate           FLOAT NOT NULL
);

-- -- name: get-journeys
-- SELECT j.*, s1.station_name_finnish AS departure_station, 
-- 	s2.station_name_finnish AS return_station FROM journey j
-- LEFT OUTER JOIN station s1 ON j.departure_station_id = s1.id
-- LEFT OUTER JOIN station s2 ON return_station_id = s2.id
-- ORDER BY ?
-- LIMIT ?, ?;

-- -- name: get-journey-count
-- SELECT count(*) AS exact_count FROM journey;


-- --name: get-stations
-- SELECT fid, id, station_name_finnish, 
-- address_finnish, city_name_finnish, 
-- operator, capacity, x_coordinate, y_coordinate
-- FROM station LIMIT ?,?;


-- --name: get-stations-by-name
-- SELECT fid, id, station_name_finnish,
-- address_finnish, city_name_finnish,
-- operator, capacity, x_coordinate,
-- y_coordinate 
-- FROM station 
-- WHERE station_name_finnish 
-- LIKE ? LIMIT ?,?;


--name: get-ret-and-dep-count
-- SELECT (
--     SELECT COUNT(*)
--     FROM   journey j
--     WHERE j.return_station_id = ?
-- ) AS tot_returns,
-- (
-- SELECT  (
--     SELECT COUNT(*)
--     FROM   journey j
--     WHERE j.departure_station_id = ?) 
-- ) AS tot_departures

--name: get-departure-count
-- SELECT count(*) from journey WHERE departure_station_id = ?;
