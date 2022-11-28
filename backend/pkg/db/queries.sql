-- name: get-version
SELECT VERSION()

-- name: create-journey-table
CREATE TABLE IF NOT EXISTS journey (
	id                     INT auto_increment NOT NULL PRIMARY KEY,
	departure_time         DATETIME NOT NULL,
	return_time            DATETIME NOT NULL,
    departure_station_id   INT NOT NULL,
    departure_station_name VARCHAR(100) NOT NULL,
    return_station_id      INT NOT NULL,
    return_station_name    VARCHAR(100) NOT NULL,
    distance               INT NOT NULL,
    duration               INT NOT NULL
);

-- name: create-station-table
CREATE TABLE IF NOT EXISTS station (
	fid                    INT NOT NULL,
    id                     INT NOT NULL PRIMARY KEY,
    station_name_finnish   VARCHAR(100) NOT NULL,
    station_name_swedish   VARCHAR(100) NOT NULL,
    address_finnish        VARCHAR(100) NOT NULL,
    address_swedish        VARCHAR(100) NOT NULL,
    city_name_finnish      VARCHAR(100) NOT NULL,
    city_name_swedish      VARCHAR(100) NOT NULL,
    operator               VARCHAR(100),
    capacity               INT NOT NULL,
    x_coordinate           FLOAT NOT NULL,
    y_coordinate           FLOAT NOT NULL
);

-- name: get-journeys
SELECT * FROM journey WHERE id > ? ORDER BY id LIMIT ?;

-- name: get-journey-count
SELECT count(*) AS exact_count FROM journey;


--name: get-stations
SELECT * FROM station WHERE id > ? ORDER BY id LIMIT ?;
