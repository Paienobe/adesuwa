package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Paienobe/adesuwa/internal/controllers"
	"github.com/Paienobe/adesuwa/internal/database"
	"github.com/Paienobe/adesuwa/internal/middleware"
	"github.com/Paienobe/adesuwa/internal/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// func (g database.Buyer) em() {}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT does not exist in environment")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL does not exist in environment")
	}

	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Failed to open database")
	}

	db := database.New(dbConn)
	apiCfg := models.ApiConfig{
		DB: db,
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// ===================== vendor routes =====================
	router.Post("/register-vendor", controllers.RegisterVendor(&apiCfg))
	router.Post("/login-vendor", controllers.LoginVendor(&apiCfg))
	router.Post("/create-product", middleware.VendorIsAuthorized(&apiCfg, controllers.CreateNewProduct))
	router.Post("/update-profile-pic", middleware.VendorIsAuthorized(&apiCfg, controllers.UpdateVendorProfilePic))
	router.Post("/update-banner", middleware.VendorIsAuthorized(&apiCfg, controllers.UpdateVendorBanner))
	router.Post("/update-description", middleware.VendorIsAuthorized(&apiCfg, controllers.UpdateVendorDescription))

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	fmt.Println("Server starting on PORT ", portString)
	log.Fatal(server.ListenAndServe())
}
