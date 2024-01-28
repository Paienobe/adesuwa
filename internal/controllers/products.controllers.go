package controllers

import (
	"fmt"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/database"
)

func CreateNewProduct(w http.ResponseWriter, r *http.Request, vendor database.Vendor) {
	fmt.Println("create product")
}
