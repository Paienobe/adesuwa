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
}

type buyerRegistrationParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Country   string `json:"country"`
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

		http.SetCookie(w, &http.Cookie{
			Name:     "adesuwa_refresh",
			Value:    refreshToken,
			Path:     "/",
			Expires:  time.Now().Add(time.Minute * 262800), //262800 minutes make 6 months
			HttpOnly: true,
		})

		type regSuc struct {
			AccessToken string       `json:"access_token"`
			Vendor      utils.Vendor `json:"vendor"`
		}

		utils.RespondWithJSON(w, http.StatusCreated, regSuc{AccessToken: accessToken, Vendor: utils.DbVendorToVendor(vendor)})

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

		_, err = apiCfg.DB.RegisterCustomer(r.Context(), database.RegisterCustomerParams{
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

		utils.RespondWithJSON(w, http.StatusCreated, models.SuccessResponse{Success: true, Msg: "Buyer account created!"})
	}
}
