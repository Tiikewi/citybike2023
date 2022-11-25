-- name: get-version
SELECT VERSION()

-- JOURNEY

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

-- name: get-journeys
SELECT * FROM journey WHERE id >= ? ORDER BY id LIMIT ?;

-- name: get-journey-count
SELECT count(*) AS exact_count FROM journey;

