CREATE SCHEMA IF NOT EXISTS go_crud;

\c go_crud;

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description TEXT,
    price DECIMAL,
    amount INTEGER
);

INSERT INTO products (
    name,
    description,
    price,
    amount
) VALUES (
    'Apple iPhone XR',
    '64GB Preto 6,1” 12MP iOS',
    3161.07,
    42
);