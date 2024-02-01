// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: product.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO product (id, name, images, price, amount_available, category, discount, vendor_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, images, price, amount_available, category, discount, vendor_id, created_at, updated_at, description
`

type CreateProductParams struct {
	ID              uuid.UUID
	Name            string
	Images          []string
	Price           float64
	AmountAvailable int32
	Category        string
	Discount        int32
	VendorID        uuid.UUID
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.ID,
		arg.Name,
		pq.Array(arg.Images),
		arg.Price,
		arg.AmountAvailable,
		arg.Category,
		arg.Discount,
		arg.VendorID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Images),
		&i.Price,
		&i.AmountAvailable,
		&i.Category,
		&i.Discount,
		&i.VendorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM product WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE product
SET name = $2, images = $3, price = $4, amount_available = $5, discount = $6, description = $7, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, name, images, price, amount_available, category, discount, vendor_id, created_at, updated_at, description
`

type UpdateProductParams struct {
	ID              uuid.UUID
	Name            string
	Images          []string
	Price           float64
	AmountAvailable int32
	Discount        int32
	Description     string
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.Name,
		pq.Array(arg.Images),
		arg.Price,
		arg.AmountAvailable,
		arg.Discount,
		arg.Description,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.Images),
		&i.Price,
		&i.AmountAvailable,
		&i.Category,
		&i.Discount,
		&i.VendorID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
	)
	return i, err
}
