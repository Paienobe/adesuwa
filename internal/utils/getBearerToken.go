package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(r *http.Request) (string, error) {
	authorization := r.Header.Get("Authorization")
	authValues := strings.Split(authorization, " ")
	if authValues[0] == "" {
		return "", errors.New("no auth found")
	}

	if authValues[0] != "Bearer" {
		return "", errors.New("malformed auth")
	}

	return authValues[1], nil

}
