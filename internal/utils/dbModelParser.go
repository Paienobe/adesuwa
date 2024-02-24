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
	Description     string    `json:"description"`
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
		Description:     dbProduct.Description,
	}
}

func DbProductsToProducts(dbProducts []database.Product) []Product {
	parsedProducts := []Product{}
	for _, product := range dbProducts {
		parsedProducts = append(parsedProducts, DbProductToProduct(product))
	}
	return parsedProducts
}

type Vendor struct {
	ID           uuid.UUID `json:"id"`
	BusinessName string    `json:"business_name"`
	Email        string    `json:"email"`
	Country      string    `json:"country"`
	ProfileImage string    `json:"profile_image"`
	BannerImage  string    `json:"banner_image"`
	Description  string    `json:"description"`
}

func DbVendorToVendor(dbVendor database.Vendor) Vendor {
	return Vendor{
		ID:           dbVendor.ID,
		BusinessName: dbVendor.BusinessName,
		Email:        dbVendor.Email,
		ProfileImage: dbVendor.ProfileImage.String,
		BannerImage:  dbVendor.BannerImage.String,
		Description:  dbVendor.Description.String,
	}
}
