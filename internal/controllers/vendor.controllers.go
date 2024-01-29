package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Paienobe/adesuwa/internal/cloudinary"
	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/google/uuid"
)

type imageUploadSuccess struct {
	Url string `json:"url"`
}

func UpdateVendorProfilePic(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	file := utils.FileExtractor(w, r)

	url := cloudinary.UploadImage(file, uuid.NewString())

	savedUrl, err := apiCfg.DB.UpdateVendorProfilePicture(r.Context(), database.UpdateVendorProfilePictureParams{ProfileImage: sql.NullString{String: url, Valid: true}, ID: vendor.ID})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to upload image", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, 200, imageUploadSuccess{Url: savedUrl.String})

}

func UpdateVendorBanner(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	file := utils.FileExtractor(w, r)

	url := cloudinary.UploadImage(file, uuid.NewString())

	savedUrl, err := apiCfg.DB.UpdateVendorBanner(r.Context(), database.UpdateVendorBannerParams{BannerImage: sql.NullString{String: url, Valid: true}, ID: vendor.ID})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to upload image", http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, 200, imageUploadSuccess{Url: savedUrl.String})
}

func UpdateVendorDescription(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig) {
	type parameters struct {
		Description string `json:"description"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&params)

	description, err := apiCfg.DB.UpdateVendorDescription(r.Context(), database.UpdateVendorDescriptionParams{
		ID:          vendor.ID,
		Description: sql.NullString{String: params.Description, Valid: true},
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "failed to update description", http.StatusBadRequest)
		return
	}

	type descriptionUpdate struct {
		Description string `json:"description"`
	}

	utils.RespondWithJSON(w, 200, descriptionUpdate{Description: description.String})
}
