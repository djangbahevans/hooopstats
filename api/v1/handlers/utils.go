package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/djangbahevans/hooopstats/api/v1/models"
)

func buildErrorResponse(w http.ResponseWriter, message string, status int) {
	resp := models.Response{
		Status:  "error",
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

func decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
