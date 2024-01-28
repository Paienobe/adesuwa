-- name: RegisterVendor :one
INSERT INTO vendor (id, name, email, password)
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: FindVendorById :one
SELECT * FROM vendor WHERE id = $1; 

-- name: FindVendorByEmail :one
SELECT * FROM vendor WHERE email = $1;