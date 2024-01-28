-- +goose Up
CREATE TABLE product (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    images TEXT[] NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    amount_available INT NOT NULL,
    category TEXT NOT NULL,
    discount INT DEFAULT 0 NOT NULL,
    vendor_id UUID REFERENCES vendor(id) NOT NULL
);
-- +goose Down
DROP TABLE product;