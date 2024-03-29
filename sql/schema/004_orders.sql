-- +goose Up
CREATE TABLE customer_order (
    id UUID PRIMARY KEY NOT NULL,
    customer_id UUID NOT NULL REFERENCES customer(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    status VARCHAR(25) NOT NULL,
    shipping_address TEXT NOT NULL,
    payment_method TEXT NOT NULL,
    payment_status TEXT NOT NULL,
    total_spent FLOAT NOT NULL
);
-- +goose Down
DROP TABLE customer_order;