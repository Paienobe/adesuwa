-- name: CreateProduct :one
INSERT INTO product (id, name, images, price, amount_available, category, discount, vendor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: UpdateProduct :one
UPDATE product
SET name = $2, images = $3, price = $4, amount_available = $5, discount = $6, description = $7, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;