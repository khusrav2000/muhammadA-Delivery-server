CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    encrypted_password varchar not null,
    name varchar not null,
    surname varchar not null,
    register_at timestamp without time zone
);