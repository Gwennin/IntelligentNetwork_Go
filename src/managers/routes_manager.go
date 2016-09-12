package managers

import (
	"encoding/json"
	"net/http"
)

func getRoutes() {
	http.HandleFunc("/", helloWorld)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
