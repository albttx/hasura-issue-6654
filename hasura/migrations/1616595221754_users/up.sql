CREATE TABLE users (
    id         SERIAL UNIQUE PRIMARY KEY,
    username   VARCHAR(128) NOT NULL,
    password   VARCHAR(128) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
