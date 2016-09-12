package managers

import "net/http"

func StartServer() {
	getRoutes()

	http.ListenAndServe(":3000", nil)
}
