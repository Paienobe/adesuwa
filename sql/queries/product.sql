-- name: CreateProduct :one
INSERT INTO product (id, name, images, price, amount_available, category, discount, description, created_at, updated_at, vendor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING *;

-- name: UpdateProduct :one
UPDATE product
SET 
  name = $1, 
  images = COALESCE($2, images),  -- Use COALESCE to handle null values
  price = $3, 
  amount_available = $4, 
  discount = $5, 
  description = $6, 
  updated_at = CURRENT_TIMESTAMP
WHERE id = $7
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1;

-- name: GetAllVendorProducts :many
SELECT * FROM product
WHERE vendor_id = $1;