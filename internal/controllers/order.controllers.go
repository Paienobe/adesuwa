package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/google/uuid"
)

func CreateOrder(w http.ResponseWriter, r *http.Request, customer database.Customer, apiCfg *models.ApiConfig) {

	type orderedItem struct {
		ProductID string  `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
		VendorID  string  `json:"vendor_id"`
	}

	type parameters struct {
		ShippingAddress string        `json:"shipping_address"`
		PaymentMethod   string        `json:"payment_method"`
		PaymentStatus   string        `json:"payment_status"`
		TotalSpent      float64       `json:"total_spent"`
		Items           []orderedItem `json:"items"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	order, err := apiCfg.DB.CreateCustomerOrder(r.Context(), database.CreateCustomerOrderParams{
		ID:              uuid.New(),
		CustomerID:      customer.ID,
		CreatedAt:       time.Now(),
		Status:          "Processing",
		ShippingAddress: params.ShippingAddress,
		PaymentMethod:   params.PaymentMethod,
		PaymentStatus:   params.PaymentStatus,
		TotalSpent:      params.TotalSpent,
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to create order", http.StatusBadRequest)
		return
	}

	for _, item := range params.Items {
		err = apiCfg.DB.CreateOrderItem(r.Context(), database.CreateOrderItemParams{
			ID:        uuid.New(),
			OrderID:   order.ID,
			ProductID: uuid.MustParse(item.ProductID),
			VendorID:  uuid.MustParse(item.VendorID),
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		})

		if err != nil {
			log.Println("failed to create order item", err)
			http.Error(w, "failed to create order", http.StatusBadRequest)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, order)
	}

}
