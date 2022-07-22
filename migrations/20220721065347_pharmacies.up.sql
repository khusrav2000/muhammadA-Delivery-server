CREATE TABLE pharmacies (
    id bigserial not null primary key,
    name varchar not null,
    geog GEOGRAPHY(Point)
)