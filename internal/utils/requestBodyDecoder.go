package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeRequestBody[T any](r *http.Request, params *T) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(params)
	if err != nil {
		log.Printf("Failed to decode body: %v", err)
		return
	}
}
