package helpers

import (
	"encoding/json"
	"github.com/Gwennin/IntelligentNetwork_Go/src/errors"
	"net/http"
)

func WriteResponseError(err *errors.INError, w http.ResponseWriter) {

	if err.Fatal {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(err)
}
