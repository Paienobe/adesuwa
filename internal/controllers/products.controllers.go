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
	"github.com/google/uuid"
)

func CreateNewProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("failed to parse form ", err)
		return
	}

	productName := r.Form.Get("name")
	productPrice, _ := strconv.ParseFloat(r.Form.Get("price"), 64)
	productAmountAvailable, _ := strconv.ParseInt(r.Form.Get("amount_available"), 0, 64)
	productCategory := r.Form.Get("category")
	productDiscount, _ := strconv.ParseInt(r.Form.Get("discount"), 0, 64)
	fileLength, _ := strconv.ParseInt(r.Form.Get("file_length"), 0, 64)
	imageURLs := []string{}

	for i := int64(0); i < fileLength; i++ {
		fileName := fmt.Sprintf("file_%v", i)
		file := utils.FileExtractor(w, r, fileName)

		imageURLs = append(imageURLs, cloudinary.UploadImage(file, uuid.NewString()))
	}

	product, err := apiCfg.DB.CreateProduct(r.Context(), database.CreateProductParams{
		ID:              uuid.New(),
		Name:            productName,
		Images:          imageURLs,
		Price:           productPrice,
		AmountAvailable: int32(productAmountAvailable),
		Category:        productCategory,
		VendorID:        vendor.ID,
		Discount:        int32(productDiscount),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create product", http.StatusBadRequest)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, utils.DbProductToProduct(product))

}
