-- name: RegisterCustomer :one
INSERT into customer (id, first_name, last_name, email, country, password, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCustomerByID :one
SELECT * FROM customer WHERE id = $1;

-- name: FindCustomerByEmail :one
SELECT * FROM customer WHERE email = $1;