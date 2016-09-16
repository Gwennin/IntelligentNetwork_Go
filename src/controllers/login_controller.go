package controllers

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/helpers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"net/http"
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

func Logout(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := helpers.ExtractToken(r)

	if token != nil {
		managers.CloseSession(*token)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
