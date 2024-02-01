-- name: CreateCustomerOrder :one
INSERT INTO customer_order (id, buyer_id, created_at, status, shipping_address, payment_method, payment_status, total_spent)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;