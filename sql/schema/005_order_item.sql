-- +goose Up
CREATE TABLE order_item (
    id UUID PRIMARY KEY NOT NULL,
    order_id UUID NOT NULL REFERENCES customer_order(id) ON DELETE CASCADE,
    product_id UUID  NOT NULL REFERENCES product(id) ON DELETE CASCADE,
    vendor_id UUID  NOT NULL REFERENCES vendor(id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    price FLOAT NOT NULL
);
-- +goose Down
DROP TABLE order_item;