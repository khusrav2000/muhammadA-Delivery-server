CREATE TABLE IF NOT EXISTS pharmacies (
    id bigserial not null primary key,
    name varchar not null,
    address varchar not null,
    geog Point not null,
    add_at timestamp without time zone,
    description varchar
);