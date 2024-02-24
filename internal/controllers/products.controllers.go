package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Paienobe/adesuwa/internal/cloudinary"
	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type ProductData struct {
	id               string
	name             string
	price            float64
	amount_available int64
	category         string
	description      string
	discount         int64
	imageURLs        []string
}

func CreateNewProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {

	parsedProduct, err := parseProductForm(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create product", http.StatusBadRequest)
		return
	}

	product, err := apiCfg.DB.CreateProduct(r.Context(), database.CreateProductParams{
		ID:              uuid.New(),
		Name:            parsedProduct.name,
		Images:          parsedProduct.imageURLs,
		Price:           parsedProduct.price,
		AmountAvailable: int32(parsedProduct.amount_available),
		Category:        parsedProduct.category,
		VendorID:        vendor.ID,
		Discount:        int32(parsedProduct.discount),
		Description:     parsedProduct.description,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create product", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, utils.DbProductToProduct(product))

}

func UpdateProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	parsedProduct, err := parseProductForm(w, r)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update product", http.StatusBadRequest)
		return
	}

	parsedProductImages := parsedProduct.imageURLs
	if len(parsedProductImages) < 1 {
		parsedProductImages = nil
	}

	updatedProduct, err := apiCfg.DB.UpdateProduct(r.Context(), database.UpdateProductParams{
		ID:              uuid.MustParse(parsedProduct.id),
		Name:            parsedProduct.name,
		Images:          parsedProductImages,
		Price:           parsedProduct.price,
		AmountAvailable: int32(parsedProduct.amount_available),
		Discount:        int32(parsedProduct.discount),
		Description:     parsedProduct.description,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update product", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DbProductToProduct(updatedProduct))
}

func DeleteProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	productID := uuid.MustParse(chi.URLParam(r, "product_id"))
	err := apiCfg.DB.DeleteProduct(r.Context(), productID)
	if err != nil {
		log.Println("failed to delete item ", err)
		http.Error(w, "failed to delete product", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.SuccessResponse{Success: true, Msg: fmt.Sprintf("Product with id %s deleted", productID)})
}

func GetAllVendorProducts(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	products, err := apiCfg.DB.GetAllVendorProducts(r.Context(), vendor.ID)
	if err != nil {
		log.Println("failed to fetch products ", err)
		http.Error(w, "failed to fetch products", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.DbProductsToProducts(products))
}

func parseProductForm(w http.ResponseWriter, r *http.Request) (ProductData, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("failed to parse form ", err)
		return ProductData{}, fmt.Errorf("failed to parse form %v", err)
	}

	productId := r.Form.Get("id")
	productName := r.Form.Get("name")
	productPrice, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	productAmountAvailable, _ := strconv.ParseInt(r.Form.Get("amount_available"), 0, 64)
	productCategory := r.Form.Get("category")
	productDescription := r.Form.Get("description")
	productDiscount, _ := strconv.ParseInt(r.Form.Get("discount"), 0, 64)
	fileLength, _ := strconv.ParseInt(r.Form.Get("file_length"), 0, 64)
	imageURLs := []string{}

	for i := int64(0); i < fileLength; i++ {
		fileName := fmt.Sprintf("file_%v", i)
		file := utils.FileExtractor(w, r, fileName)

		imageURLs = append(imageURLs, cloudinary.UploadImage(file, uuid.NewString()))
	}

	return ProductData{
		id:               productId,
		name:             productName,
		price:            productPrice,
		amount_available: productAmountAvailable,
		category:         productCategory,
		description:      productDescription,
		discount:         productDiscount,
		imageURLs:        imageURLs,
	}, nil
}
