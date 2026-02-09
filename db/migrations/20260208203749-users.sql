-- +migrate Up
CREATE TYPE roles AS ENUM ('admin', 'it_support');

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    code_name VARCHAR(10) NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS users;
DROP TYPE IF EXISTS roles;
