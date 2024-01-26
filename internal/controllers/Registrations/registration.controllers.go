package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/google/uuid"
)

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

		_, err = apiCfg.DB.RegisterVendor(r.Context(), database.RegisterVendorParams{
			ID:       uuid.New(),
			Name:     params.Name,
			Email:    params.Email,
			Password: hashedPassword,
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Error registering vendor", http.StatusBadRequest)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.SuccessResponse{Success: true, Msg: "Vendor account created!"})

	}
}

func RegisterBuyer(apiCfg *models.ApiConfig) http.HandlerFunc {
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

		_, err = apiCfg.DB.RegisterBuyer(r.Context(), database.RegisterBuyerParams{
			ID:        uuid.New(),
			FirstName: params.FirstName,
			LastName:  params.LastName,
			Email:     params.Email,
			Password:  hashedPassword,
		})

		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong", http.StatusBadRequest)
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.SuccessResponse{Success: true, Msg: "Buyer account created!"})
	}
}
