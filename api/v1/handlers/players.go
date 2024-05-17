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

type PlayerHandler struct {
	db *db.DB
}

func RegisterPlayerHandler(r *http.ServeMux, db *db.DB) {
	playerHandler := &PlayerHandler{db}

	r.HandleFunc("GET /players", playerHandler.GetPlayers)
	r.HandleFunc("GET /players/{id}", playerHandler.GetPlayerById)
	r.HandleFunc("POST /players", playerHandler.CreatePlayer)
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	ps := services.NewPlayerService(h.db)
	players, err := ps.GetPlayers()
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "players",
		Data:    players,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *PlayerHandler) GetPlayerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		buildErrorResponse(w, "invalid player_id", http.StatusBadRequest)
		return
	}

	ps := services.NewPlayerService(h.db)
	player, err := ps.GetPlayerById(id)
	if err != nil {
		if errors.Is(err, services.ErrPlayerNotFound) {
			buildErrorResponse(w, "player not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "player",
		Data:    player,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	req, err := validatePlayerRequest(r)
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ps := services.NewPlayerService(h.db)
	player, err := ps.CreatePlayer(models.Player{FirstName: req.FirstName, LastName: req.LastName, TeamID: req.TeamID})
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "player created",
		Data: models.Player{
			Id:        player,
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func validatePlayerRequest(r *http.Request) (*models.CreatePlayerRequest, error) {
	var req models.CreatePlayerRequest
	if err := decode(r, &req); err != nil {
		return nil, err
	}

	if req.FirstName == "" {
		return nil, errors.New("first_name is required")
	}

	if req.LastName == "" {
		return nil, errors.New("last_name is required")
	}

	if req.TeamID == 0 {
		return nil, errors.New("team_id is required")
	}

	return &req, nil
}
