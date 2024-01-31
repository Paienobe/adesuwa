package utils

import (
	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/google/uuid"
)

type Product struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Images          []string  `json:"images"`
	Price           float64   `json:"price"`
	AmountAvailable int32     `json:"amount_available"`
	Category        string    `json:"category"`
	Discount        int32     `json:"discount"`
	VendorID        uuid.UUID `json:"vendor_id"`
}

func DbProductToProduct(dbProduct database.Product) Product {
	return Product{
		ID:              dbProduct.ID,
		Name:            dbProduct.Name,
		Images:          dbProduct.Images,
		Price:           dbProduct.Price,
		AmountAvailable: dbProduct.AmountAvailable,
		Category:        dbProduct.Category,
		Discount:        dbProduct.Discount,
		VendorID:        dbProduct.VendorID,
	}
}
