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

type TeeamHandler struct {
	db *db.DB
}

func RegisterTeamHandler(r *http.ServeMux, db *db.DB) {
	teamHandler := &TeeamHandler{db}

	r.HandleFunc("GET /teams", teamHandler.GetTeams)
	r.HandleFunc("GET /teams/{id}", teamHandler.GetTeamById)
	r.HandleFunc("POST /teams", teamHandler.CreateTeam)
}

func (h *TeeamHandler) GetTeams(w http.ResponseWriter, r *http.Request) {
	ts := services.NewTeamService(h.db)
	teams, err := ts.GetTeams()
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "teams",
		Data:    teams,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *TeeamHandler) GetTeamById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		buildErrorResponse(w, "invalid team_id", http.StatusBadRequest)
		return
	}

	ts := services.NewTeamService(h.db)
	team, err := ts.GetTeamById(id)
	if err != nil {
		if errors.Is(err, services.ErrTeamNotFound) {
			buildErrorResponse(w, "team not found", http.StatusNotFound)
			return
		}

		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "team",
		Data:    team,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *TeeamHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	req, err := validateTeamRequest(r)
	if err != nil {
		buildErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts := services.NewTeamService(h.db)
	var id int64
	if id, err = ts.CreateTeam(models.Team{Name: req.Name}); err != nil {
		buildErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.Response{
		Status:  "success",
		Message: "team created",
		Data: models.Team{
			Id:   id,
			Name: req.Name,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func validateTeamRequest(r *http.Request) (*models.CreateTeamRequest, error) {
	var req models.CreateTeamRequest
	if err := decode(r, req); err != nil {
		return nil, err
	}

	if req.Name == "" {
		return nil, errors.New("name is required")
	}

	return &req, nil
}
