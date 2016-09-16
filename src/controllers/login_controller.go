package controllers

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/helpers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()

	if ok {
		authenticator := managers.GetAuthenticator()

		if authenticator != nil {
			authenticated, err := authenticator.Authenticate(username, password)

			if err != nil {
				helpers.WriteResponseError(err, w)
			} else if authenticated {
				token := managers.OpenSession(username)
				w.Write([]byte(token))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		} else {
			err := errors.FatalError(3, "No authenticator found")
			helpers.WriteResponseError(err, w)
		}

		return
	}
	err := errors.NewError(2, "No authorization header found.")
	helpers.WriteResponseError(err, w)
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
		err := errors.NewError(2, "No authorization header found.")
		helpers.WriteResponseError(err, w)
	}
}
