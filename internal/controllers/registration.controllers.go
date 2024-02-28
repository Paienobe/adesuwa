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

type vendorRegistrationParams struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Country      string `json:"country"`
	BusinessName string `json:"business_name"`
	PhoneNumber  string `json:"phone_number"`
}

type buyerRegistrationParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
}

type registrationSuccess[T any] struct {
	AccessToken string `json:"access_token"`
	Data        T      `json:"data"`
}

func RegisterVendor(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := vendorRegistrationParams{}

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&params)

		hashedPassword, err := utils.GeneratehashPassword(params.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong:", http.StatusBadRequest)
			return
		}

		vendor, err := apiCfg.DB.RegisterVendor(r.Context(), database.RegisterVendorParams{
			ID:           uuid.New(),
			FirstName:    params.FirstName,
			LastName:     params.LastName,
			BusinessName: params.BusinessName,
			Email:        params.Email,
			PhoneNumber:  params.PhoneNumber,
			Country:      params.Country,
			Password:     hashedPassword,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering vendor", http.StatusBadRequest)
			return
		}

		refreshToken, err := utils.GenerateJWT(vendor.Email, vendor.ID, false)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering vendor", http.StatusBadRequest)
			return
		}

		accessToken, err := utils.GenerateJWT(vendor.Email, vendor.ID, true)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering vendor", http.StatusBadRequest)
			return
		}

		utils.SetRefreshCookie(w, refreshToken)

		utils.RespondWithJSON(w, http.StatusCreated, registrationSuccess[utils.Vendor]{AccessToken: accessToken, Data: utils.DbVendorToVendor(vendor)})

	}
}

func RegisterCustomer(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := buyerRegistrationParams{}

		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&params)

		hashedPassword, err := utils.GeneratehashPassword(params.Password)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong:", http.StatusBadRequest)
			return
		}

		customer, err := apiCfg.DB.RegisterCustomer(r.Context(), database.RegisterCustomerParams{
			ID:        uuid.New(),
			FirstName: params.FirstName,
			LastName:  params.LastName,
			Email:     params.Email,
			Password:  hashedPassword,
			Country:   params.Country,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong", http.StatusBadRequest)
			return
		}

		refreshToken, err := utils.GenerateJWT(customer.Email, customer.ID, false)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering customer", http.StatusBadRequest)
			return
		}

		accessToken, err := utils.GenerateJWT(customer.Email, customer.ID, true)

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering customer", http.StatusBadRequest)
			return
		}

		utils.SetRefreshCookie(w, refreshToken)

		utils.RespondWithJSON(w, http.StatusCreated, registrationSuccess[utils.Customer]{AccessToken: accessToken, Data: utils.DbCustomerToCustomer(customer)})
	}
}
