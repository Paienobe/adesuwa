package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/Paienobe/adesuwa/internal/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type protectedVendorHandler func(w http.ResponseWriter, r *http.Request, vendor database.Vendor, apiCfg *models.ApiConfig)
type protectedBuyerHandler func(w http.ResponseWriter, r *http.Request, vendor database.Customer, apiCfg *models.ApiConfig)

func VendorIsAuthorized(apiCfg *models.ApiConfig, handler protectedVendorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := utils.GetBearerToken(r)
		if err != nil {
			log.Println(err)
			return
		}

		id, err := JwtValidation(w, r, tokenString)
		if err != nil {
			log.Println(err)
			http.Error(w, "Vendor not authorized", http.StatusUnauthorized)
		}

		vendor, err := apiCfg.DB.FindVendorById(r.Context(), id)
		if err != nil {
			log.Println("Error finding vendor: ", err)
			return
		}

		handler(w, r, vendor, apiCfg)
	}
}

func BuyerIsAuthorized(apiCfg *models.ApiConfig, handler protectedBuyerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := utils.GetBearerToken(r)
		if err != nil {
			log.Println(err)
			return
		}

		id, err := JwtValidation(w, r, tokenString)
		if err != nil {
			log.Println(err)
			http.Error(w, "Buyer not authorized", http.StatusUnauthorized)
		}

		buyer, err := apiCfg.DB.GetCustomerByID(r.Context(), id)
		if err != nil {
			log.Println("error finding vendor: ", err)
			return
		}

		handler(w, r, buyer, apiCfg)
	}
}

func JwtValidation(w http.ResponseWriter, r *http.Request, tokenString string) (uuid.UUID, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Println("JWT_SECRET does not exist in environment")
		return uuid.Nil, errors.New("something went wrong")
	}

	var mySigningKey = []byte(secretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		log.Printf("Failed to parse JWT %v", err)
		return uuid.Nil, errors.New("something went wrong")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := claims["id"].(string)
		parsedId, err := uuid.Parse(id)
		if err != nil {
			log.Println("failed to parse uuid", err)
		}
		return parsedId, nil
	}

	return uuid.Nil, errors.New("not authorized")
}
