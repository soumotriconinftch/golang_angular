CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    username varchar(255) UNIQUE NOT NULL,
    password bytea NOT NULL
);
