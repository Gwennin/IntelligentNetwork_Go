package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/helpers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
	"github.com/gorilla/mux"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	users, err := models.ListUsers()
	if err == nil {
		json.NewEncoder(w).Encode(users)
	} else {
		helpers.WriteResponseError(err, w)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := new(models.NewUser)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		err := errors.NewError(4, "Unable to decode body content.")
		helpers.WriteResponseError(err, w)
		return
	}

	username, userErr := models.AddUser(user)
	if userErr != nil {
		helpers.WriteResponseError(userErr, w)
		return
	}

	space := new(models.Space)
	space.Name = "u_" + username
	space.Owner = username

	spaceErr := models.AddPrivateSpace(space)

	if spaceErr == nil {
		json.NewEncoder(w).Encode(username)
	} else {
		helpers.WriteResponseError(spaceErr, w)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := models.DeleteUser(alias)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	password := buf.String()

	err := models.ChangePassword(alias, password)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func AddUserSpace(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]
	space := vars["space"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userSpace := new(models.UserSpace)
	userSpace.UserId = alias
	userSpace.SpaceId = space

	err := models.AddUserSpace(userSpace)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func DeleteUserSpace(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]
	space := vars["space"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	userSpace := new(models.UserSpace)
	userSpace.UserId = alias
	userSpace.SpaceId = space

	err := models.DeleteUserSpace(userSpace)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func ListUserSpaces(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	spaces, err := models.ListUserSpaces(alias)
	if err == nil {
		json.NewEncoder(w).Encode(spaces)
	} else {
		helpers.WriteResponseError(err, w)
	}
}

func ListOwnedSpaces(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	token := helpers.ExtractToken(r)
	currentUser := *managers.GetSessionUser(*token)

	if currentUser != alias {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	spaces, err := models.ListOwnedSpaces(alias)
	if err == nil {
		json.NewEncoder(w).Encode(spaces)
	} else {
		helpers.WriteResponseError(err, w)
	}
}
