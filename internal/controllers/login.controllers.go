package controllers

import (
	"errors"
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

type handleLoginArgs struct {
	paramPassword string
	password      string
	email         string
	id            uuid.UUID
}

type loginSuccess[T any] struct {
	AccessToken string `json:"access_token"`
	Data        T      `json:"data"`
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

		args := handleLoginArgs{
			paramPassword: params.Password,
			password:      vendor.Password,
			email:         vendor.Email,
			id:            vendor.ID,
		}

		refreshToken, accessToken, err := handleLogin(w, args)
		if err != nil {
			log.Println(err)
			return
		}

		utils.SetRefreshCookie(w, refreshToken)

		utils.RespondWithJSON(w, http.StatusOK, loginSuccess[utils.Vendor]{AccessToken: accessToken,
			Data: utils.DbVendorToVendor(vendor)})

	}
}

func LoginCustomer(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := LoginParams{}

		utils.DecodeRequestBody(r, &params)

		customer, err := apiCfg.DB.FindCustomerByEmail(r.Context(), params.Email)

		if err != nil {
			log.Println(err)
			http.Error(w, "could not find buyer with provided email", http.StatusNotFound)
			return
		}

		args := handleLoginArgs{
			paramPassword: params.Password,
			password:      customer.Password,
			email:         customer.Email,
			id:            customer.ID,
		}

		refreshToken, accessToken, err := handleLogin(w, args)
		if err != nil {
			log.Println(err)
			return
		}

		utils.SetRefreshCookie(w, refreshToken)

		utils.RespondWithJSON(w, http.StatusOK, loginSuccess[utils.Customer]{AccessToken: accessToken, Data: utils.DbCustomerToCustomer(customer)})
	}
}

func handleLogin(w http.ResponseWriter, args handleLoginArgs) (string, string, error) {
	passwordIsCorrect := utils.CheckPasswordHash(args.paramPassword, args.password)
	if !passwordIsCorrect {
		http.Error(w, "password is incorrect", http.StatusUnauthorized)
		return "", "", errors.New("password is incorrect")
	}

	refreshToken, err := utils.GenerateJWT(args.email, args.id, true)
	if err != nil {
		log.Println("failed to generate JWT")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return "", "", errors.New("failed to generate JWT")
	}

	accessToken, err := utils.GenerateJWT(args.email, args.id, true)
	if err != nil {
		log.Println("failed to generate JWT")
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return "", "", errors.New("failed to generate JWT")
	}

	return refreshToken, accessToken, nil
}
