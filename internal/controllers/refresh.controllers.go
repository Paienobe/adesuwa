package controllers

import (
	"log"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/middleware"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
)

func RefreshUser(apiCfg *models.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("adesuwa_refresh")
		if err != nil {
			log.Println("failed to get cookie ", err)
		}

		refreshToken := cookie.Value
		id, err := middleware.JwtValidation(w, r, refreshToken)
		if err != nil {
			log.Println(err)
			http.Error(w, "Invalid user", http.StatusUnauthorized)
		}

		vendor, err := apiCfg.DB.FindVendorById(r.Context(), id)
		if err != nil {
			log.Println("no vendor", err)
			return
		} else {
			access_token, err := utils.GenerateJWT(vendor.Email, vendor.ID, true)
			if err != nil {
				log.Println("err generating access token ", err)
				return
			}

			type VendorRefresh struct {
				AccessToken string       `json:"access_token"`
				UserType    string       `json:"user_type"`
				UserData    utils.Vendor `json:"user_data"`
				ExpiresIn   int          `json:"expires_in"`
			}

			parsedVendor := utils.DbVendorToVendor(vendor)

			utils.RespondWithJSON(w, 200, VendorRefresh{AccessToken: access_token, UserType: "vendor", ExpiresIn: 900, UserData: parsedVendor})
		}

		// customer, err := apiCfg.DB.GetCustomerByID(r.Context(), id)
		// if err != nil {
		// 	log.Println("no customer", err)
		// 	return
		// } else {
		// 	utils.RespondWithJSON(w, 200, customer)
		// }
	}

}
