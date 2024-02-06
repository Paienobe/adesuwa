-- +goose Up
CREATE TABLE product (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    images TEXT[] NOT NULL,
    price FLOAT NOT NULL,
    amount_available INT NOT NULL,
    category TEXT NOT NULL,
    discount INT DEFAULT 0 NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    vendor_id UUID NOT NULL REFERENCES vendor(id)  ON DELETE CASCADE
);
-- +goose Down
DROP TABLE product;