package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/djangbahevans/hooopstats/db"
)

type StatsHandler struct {
	db *db.DB
}

func RegisterStatHandler(r *http.ServeMux, db *db.DB) {
	statsHandler := &StatsHandler{db}

	r.HandleFunc("GET /stats/{id}", statsHandler.GetStats)
	r.HandleFunc("GET /stats/", statsHandler.CreateStats)
}

func (h *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := struct {
		TotalUsers int `json:"total_users"`
	}{
		TotalUsers: 30,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (h *StatsHandler) CreateStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Stats created")
}
