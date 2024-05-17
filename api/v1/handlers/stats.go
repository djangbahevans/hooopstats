package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/djangbahevans/hooopstats/api/v1/models"
	"github.com/djangbahevans/hooopstats/db"
	"github.com/djangbahevans/hooopstats/services"
)

type StatsHandler struct {
	db *db.DB
}

func RegisterStatHandler(r *http.ServeMux, db *db.DB) {
	statsHandler := &StatsHandler{db}

	r.HandleFunc("GET /stats/{id}", statsHandler.GetStatsByPlayerId)
	r.HandleFunc("POST /stats", statsHandler.CreateStats)
	r.HandleFunc("GET /team/{id}/stats/average", statsHandler.GetTeamStatsAverage)
	r.HandleFunc("GET /player/{id}/stats/average", statsHandler.GetPlayerStatsAverage)
}

func (h *StatsHandler) GetStatsByPlayerId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		buildErrorResponse(w, "invalid player_id", http.StatusBadRequest)
		return
	}

	ss := services.NewStatsService(h.db)
	stats, err := ss.GetStatsByPlayerId(id)
	if err != nil {
		if errors.Is(err, services.ErrStatsNotFound) {
			buildErrorResponse(w, "stats not found", http.StatusNotFound)
			return
		}

		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "player stats",
		Data:    stats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *StatsHandler) CreateStats(w http.ResponseWriter, r *http.Request) {
	req, err := validateStatsRequest(r)
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ss := services.NewStatsService(h.db)
	var id int64
	if id, err = ss.CreateStats(req.PlayerID, models.Stats(*req)); err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "stats created",
		Data: models.Stats{
			Points:        req.Points,
			Rebounds:      req.Rebounds,
			Assists:       req.Assists,
			Steals:        req.Steals,
			Blocks:        req.Blocks,
			Fouls:         req.Fouls,
			Turnovers:     req.Turnovers,
			MinutesPlayed: req.MinutesPlayed,
			PlayerID:      req.PlayerID,
			Id:            id,
			GameNumber:    req.GameNumber,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *StatsHandler) GetTeamStatsAverage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		buildErrorResponse(w, "invalid team_id", http.StatusBadRequest)
		return
	}
	ss := services.NewStatsService(h.db)
	stats, err := ss.GetTeamStatsAverage(id)
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "average stats",
		Data:    stats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *StatsHandler) GetPlayerStatsAverage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		buildErrorResponse(w, "invalid player_id", http.StatusBadRequest)
		return
	}

	ss := services.NewStatsService(h.db)
	stats, err := ss.GetPlayerStatsAverage(id)
	if err != nil {
		if errors.Is(err, services.ErrStatsNotFound) {
			buildErrorResponse(w, "stats not found", http.StatusNotFound)
			return
		}

		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "player average stats",
		Data:    stats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func validateStatsRequest(r *http.Request) (*models.CreateStatsRequest, error) {
	var req models.CreateStatsRequest
	if err := decode(r, &req); err != nil {
		return nil, err
	}

	if req.PlayerID == 0 {
		return nil, errors.New("missing player_id")
	}

	if req.GameNumber == 0 {
		return nil, errors.New("missing game_number")
	}

	if req.Points < 0 ||
		req.Rebounds < 0 ||
		req.Assists < 0 ||
		req.Steals < 0 ||
		req.Blocks < 0 ||
		req.Turnovers < 0 {
		return nil, errors.New("invalid stats")
	}

	if req.Fouls < 0 || req.Fouls > 6 {
		return nil, errors.New("invalid fouls")
	}

	if req.MinutesPlayed < 0 || req.MinutesPlayed > 48 {
		return nil, errors.New("invalid minutes played")
	}

	return &req, nil
}
