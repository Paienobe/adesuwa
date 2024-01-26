-- name: RegisterVendor :one
INSERT INTO vendor (id, name, email, password)
VALUES($1, $2, $3, $4)
RETURNING *;