CREATE TABLE product (
        price INTEGER NOT NULL CHECK (price >0),
        name TEXT NOT NULL,
)

CREATE TABLE customers (
        name TEXT NOT NULL,
        tel TEXT NOT NULL UNIQUE,
)

CREATE TABLE bill (
        product_name TEXT REFERENCES product
        coustomers_name TEXT REFERENCES customer
)