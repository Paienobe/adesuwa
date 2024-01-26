// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Buyer struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	ProfileImage sql.NullString
	Password     string
}

type Vendor struct {
	ID           uuid.UUID
	Name         string
	Email        string
	ProfileImage sql.NullString
	BannerImage  sql.NullString
	Description  sql.NullString
	Password     string
}