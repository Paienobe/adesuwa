package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginVendor(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := LoginParams{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&params)
		if err != nil {
			log.Printf("Failed to decode body: %v", err)
			return
		}

		vendor, err := apiCfg.DB.FindVendorByEmail(r.Context(), params.Email)

		if err != nil {
			log.Println(err)
			http.Error(w, "could not find vendor with provided email", http.StatusNotFound)
			return
		}

		passwordIsCorrect := utils.CheckPasswordHash(params.Password, vendor.Password)
		if !passwordIsCorrect {
			http.Error(w, "password is incorrect", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateJWT(vendor.Email, vendor.ID)
		if err != nil {
			log.Println("failed to generate JWT")
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}

		type vendorLogin struct {
			Token string `json:"token"`
		}

		utils.RespondWithJSON(w, 200, vendorLogin{Token: tokenString})
	}
}
