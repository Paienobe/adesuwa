package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	res, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Failed to marshal data")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}
