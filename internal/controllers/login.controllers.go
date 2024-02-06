package controllers

import (
	"log"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/google/uuid"
)

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginVendor(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := LoginParams{}

		utils.DecodeRequestBody(r, &params)

		vendor, err := apiCfg.DB.FindVendorByEmail(r.Context(), params.Email)

		if err != nil {
			log.Println(err)
			http.Error(w, "could not find vendor with provided email", http.StatusNotFound)
			return
		}

		handleLogin(w, params.Password, vendor.Password, vendor.Email, vendor.ID)
	}
}

func LoginBuyer(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := LoginParams{}

		utils.DecodeRequestBody(r, &params)

		buyer, err := apiCfg.DB.FindCustomerByEmail(r.Context(), params.Email)

		if err != nil {
			log.Println(err)
			http.Error(w, "could not find buyer with provided email", http.StatusNotFound)
			return
		}

		handleLogin(w, params.Password, buyer.Password, buyer.Email, buyer.ID)
	}
}

func handleLogin(w http.ResponseWriter, paramPassword, userPassword, userEmail string, userID uuid.UUID) {
	passwordIsCorrect := utils.CheckPasswordHash(paramPassword, userPassword)
	if !passwordIsCorrect {
		http.Error(w, "password is incorrect", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.GenerateJWT(userEmail, userID)
	if err != nil {
		log.Println("failed to generate JWT")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	type login struct {
		Token string `json:"token"`
	}

	utils.RespondWithJSON(w, 200, login{Token: tokenString})
}
