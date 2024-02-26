package utils

import (
	"net/http"
	"time"
)

func SetRefreshCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "adesuwa_refresh",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * 262800), //262800 minutes make 6 months
		HttpOnly: true,
	})
}
