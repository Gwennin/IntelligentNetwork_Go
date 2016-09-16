package server

import (
	"github.com/Gwennin/IntelligentNetwork_Go/src/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func getRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	router.HandleFunc("/users", controllers.ListUsers).Methods("GET")
	router.HandleFunc("/users/add", controllers.AddUser).Methods("POST")
	router.HandleFunc("/users/{alias}/delete", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{alias}/change/password", controllers.ChangePassword).Methods("PUT")
	router.HandleFunc("/users/{alias}/spaces", controllers.ListUserSpaces).Methods("GET")

	router.HandleFunc("/users/{alias}/spaces/{space}/add", controllers.AddUserSpace).Methods("POST")
	router.HandleFunc("/users/{alias}/spaces/{space}/delete", controllers.DeleteUserSpace).Methods("DELETE")
	router.HandleFunc("/users/{alias}/spaces/owned", controllers.ListOwnedSpaces).Methods("GET")

	router.HandleFunc("/spaces", controllers.ListPublicSpaces).Methods("GET")
	router.HandleFunc("/spaces/add", controllers.AddSpace).Methods("POST")
	router.HandleFunc("/spaces/{name}/delete", controllers.DeleteSpace).Methods("DELETE")

	router.HandleFunc("/spaces/{name}", controllers.ListLinks).Methods("GET")
	router.HandleFunc("/spaces/{name}/add/link", controllers.AddLink).Methods("POST")
	router.HandleFunc("/spaces/{name}/delete/link/{id:[0-9]+}", controllers.DeleteLink).Methods("DELETE")
	router.HandleFunc("/spaces/{name}/read/{id:[0-9]+}", controllers.SetLinkRead).Methods("PUT")

	http.Handle("/", router)
}
