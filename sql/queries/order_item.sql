-- name: CreateOrderItem :exec
INSERT INTO order_item (id, order_id, product_id, vendor_id, quantity, price) 
VALUES($1, $2, $3, $4, $5, $6);