-- name: CreateProduct :one
INSERT INTO product (id, name, images, price, amount_available, category, discount, vendor_id)
VALUES ($1, $2, $3, $4, $5, $6, 0, $7)
RETURNING *;