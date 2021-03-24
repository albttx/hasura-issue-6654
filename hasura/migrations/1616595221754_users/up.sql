CREATE TABLE users (
    id         SERIAL UNIQUE PRIMARY KEY,
    first_name VARCHAR(128) NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
