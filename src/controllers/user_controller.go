package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
	"github.com/gorilla/mux"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := models.ListUsers()
	json.NewEncoder(w).Encode(users)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := new(models.NewUser)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return
	}

	username := models.AddUser(user)

	space := new(models.Space)
	space.Name = "u_" + username
	space.Owner = username

	models.AddPrivateSpace(space)

	json.NewEncoder(w).Encode(username)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	models.DeleteUser(alias)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	password := buf.String()

	models.ChangePassword(alias, password)
}

func AddUserSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]
	space := vars["space"]

	userSpace := new(models.UserSpace)
	userSpace.UserId = alias
	userSpace.SpaceId = space

	models.AddUserSpace(userSpace)
}

func DeleteUserSpace(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]
	space := vars["space"]

	userSpace := new(models.UserSpace)
	userSpace.UserId = alias
	userSpace.SpaceId = space

	models.DeleteUserSpace(userSpace)
}

func ListUserSpaces(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	spaces := models.ListUserSpaces(alias)
	json.NewEncoder(w).Encode(spaces)
}

func ListOwnedSpaces(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	spaces := models.ListOwnedSpaces(alias)
	json.NewEncoder(w).Encode(spaces)
}
