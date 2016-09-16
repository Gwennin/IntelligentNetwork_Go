package controllers

import (
	"encoding/json"
	"github.com/Gwennin/IntelligentNetwork_Go/src/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func ListPublicSpaces(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	spaces := models.ListPublicSpaces()
	json.NewEncoder(w).Encode(spaces)
}

func AddSpace(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	space := new(models.Space)
	if err := json.NewDecoder(r.Body).Decode(space); err != nil {
		return
	}

	models.AddSpace(space)
}

func DeleteSpace(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	models.DeleteSpace(name)
}

func ListLinks(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	// TODO Change this to get it in session
	alias := "user"

	links := models.ListLinks(name, alias)
	json.NewEncoder(w).Encode(links)
}

func AddLink(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]

	link := new(models.Link)

	if err := json.NewDecoder(r.Body).Decode(link); err != nil {
		return
	}

	link.PostedIn = name
	// TODO change this to get it from session
	link.PostedBy = "user"
	link.PostedOn = time.Now()

	newLink := models.AddLink(link)
	json.NewEncoder(w).Encode(newLink)
}

func DeleteLink(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	name := vars["name"]
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	models.DeleteLink(id, name)
}

func SetLinkRead(w http.ResponseWriter, r *http.Request) {
	if !IsTokenValid(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	strId := vars["id"]
	id, _ := strconv.Atoi(strId)

	// TODO Change this to get it in session
	alias := "user"

	models.SetLinkRead(id, alias)
}
