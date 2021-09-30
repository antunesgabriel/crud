CREATE SCHEMA IF NOT EXISTS go_crud;

\c go_crud;

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description TEXT,
    price DECIMAL,
    amount INTEGER
);