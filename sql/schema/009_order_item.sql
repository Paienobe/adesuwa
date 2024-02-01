-- +goose Up
CREATE TABLE order_item (
    id UUID PRIMARY KEY NOT NULL,
    order_id UUID REFERENCES customer_order(id) NOT NULL,
    product_id UUID REFERENCES product(id) NOT NULL,
    vendor_id UUID REFERENCES vendor(id) NOT NULL,
    quantity INT NOT NULL,
    price FLOAT NOT NULL
);
-- +goose Down
DROP TABLE order_item;