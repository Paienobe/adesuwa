package controllers

import (
	"fmt"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
)

func CreateNewProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	fmt.Println("create product")
}
