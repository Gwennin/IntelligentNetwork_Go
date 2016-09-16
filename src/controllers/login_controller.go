package controllers

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if ok {
		authenticator := managers.GetAuthenticator()

		if authenticator != nil && authenticator.Authenticate(username, password) {
			token := managers.OpenSession(username)
			w.Write([]byte(token))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}

		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func IsTokenValid(r *http.Request) bool {
	token := ExtractToken(r)

	if token != nil {
		openedOn := managers.GetSessionOpenedDate(*token)
		if openedOn != nil {
			sessionTimeout := os.Getenv("SESSION_TIMEOUT")
			seconds, _ := strconv.Atoi(sessionTimeout)
			expireDate := openedOn.Add(time.Duration(seconds) * time.Second)

			if time.Now().Before(expireDate) {
				return true
			}

			managers.CloseSession(*token)
		}
	}

	return false
}

func ExtractToken(r *http.Request) *string {
	authHeader := r.Header.Get("Authorization")

	const bearerPrefix = "Bearer "
	if authHeader != "" && strings.HasPrefix(authHeader, bearerPrefix) {
		token := strings.TrimPrefix(authHeader, bearerPrefix)
		return &token
	}
	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := ExtractToken(r)

	if token != nil {
		managers.CloseSession(*token)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
