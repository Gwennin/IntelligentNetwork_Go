package controllers

import (
	"encoding/json"
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"github.com/Gwennin/IntelligentNetwork_Go/src/helpers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/managers"
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func ListPublicSpaces(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	spaces, err := models.ListPublicSpaces()
	if err == nil {
		json.NewEncoder(w).Encode(spaces)
	} else {
		helpers.WriteResponseError(err, w)
	}
}

func AddSpace(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	space := new(models.Space)
	if decodeErr := json.NewDecoder(r.Body).Decode(space); decodeErr != nil {
		err := errors.NewError(4, "Unable to decode body content.")
		helpers.WriteResponseError(err, w)
		return
	}

	addErr := models.AddSpace(space)
	if addErr != nil {
		helpers.WriteResponseError(addErr, w)
	}
}

func DeleteSpace(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	err := models.DeleteSpace(name)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func ListLinks(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	// Token presence was checked before (IsTokenValid)
	// So, we are able to unwrap the value pointed by token
	token := helpers.ExtractToken(r)
	alias := *managers.GetSessionUser(*token)

	links, err := models.ListLinks(name, alias)
	if err == nil {
		json.NewEncoder(w).Encode(links)
	} else {
		helpers.WriteResponseError(err, w)
	}
}

func AddLink(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	link := new(models.Link)

	if decodeErr := json.NewDecoder(r.Body).Decode(link); decodeErr != nil {
		err := errors.NewError(4, "Unable to decode body content.")
		helpers.WriteResponseError(err, w)
		return
	}

	// Token presence was checked before (IsTokenValid)
	// So, we are able to unwrap the value pointed by token
	token := helpers.ExtractToken(r)
	alias := *managers.GetSessionUser(*token)

	link.PostedIn = name
	link.PostedBy = alias
	link.PostedOn = time.Now()

	newLink, err := models.AddLink(link)
	if err == nil {
		json.NewEncoder(w).Encode(newLink)
	} else {
		helpers.WriteResponseError(err, w)
	}
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	// The route have some regexp check, so we could expect a correct format & conversion
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	err := models.DeleteLink(id, name)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}

func SetLinkRead(w http.ResponseWriter, r *http.Request) {
	if !helpers.IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)

	// The route have some regexp check, so we could expect a correct format & conversion
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	// Token presence was checked before (IsTokenValid)
	// So, we are able to unwrap the value pointed by token
	token := helpers.ExtractToken(r)
	alias := *managers.GetSessionUser(*token)

	err := models.SetLinkRead(id, alias)
	if err != nil {
		helpers.WriteResponseError(err, w)
	}
}
