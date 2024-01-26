-- name: RegisterBuyer :one
INSERT into buyer (id, first_name, last_name, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;