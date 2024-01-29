package utils

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
)

func FileExtractor(w http.ResponseWriter, r *http.Request) multipart.File {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("failed to parse form ", err)
		return nil
	}

	json.NewDecoder(r.Body).Decode(r.Body)

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("failed to get file ", err)
		return nil
	}

	return file
}
