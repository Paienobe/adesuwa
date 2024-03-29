// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	Country      string
	ProfileImage sql.NullString
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CustomerOrder struct {
	ID              uuid.UUID
	CustomerID      uuid.UUID
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Status          string
	ShippingAddress string
	PaymentMethod   string
	PaymentStatus   string
	TotalSpent      float64
}

type OrderItem struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	ProductID uuid.UUID
	VendorID  uuid.UUID
	Quantity  int32
	Price     float64
}

type Product struct {
	ID              uuid.UUID
	Name            string
	Images          []string
	Price           float64
	AmountAvailable int32
	Category        string
	Discount        int32
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	VendorID        uuid.UUID
}

type Vendor struct {
	ID           uuid.UUID
	FirstName    string
	LastName     string
	BusinessName string
	Email        string
	PhoneNumber  string
	Country      string
	ProfileImage sql.NullString
	BannerImage  sql.NullString
	Description  sql.NullString
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
