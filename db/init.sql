CREATE DATABASE velostat
    WITH
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8';

\c velostat;

-- Statistics
CREATE TABLE IF NOT EXISTS station_statistics (
    stat_id SERIAL PRIMARY KEY,
    station VARCHAR(4) NOT NULL,
    address TEXT NOT NULL,
    lon REAL, -- Maybe it will be better to make a geopoint in the future :)
    lat REAL,
    total_places SMALLINT NOT NULL,
    free_places SMALLINT NOT NULL,
    is_locked BOOLEAN NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);

COMMENT ON TABLE station_statistics
    IS 'Store simple denormalized station statistics';

CREATE INDEX ON station_statistics (lower(station));
