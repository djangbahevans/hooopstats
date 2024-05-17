package services

import (
	"github.com/djangbahevans/hooopstats/api/v1/models"
	"github.com/djangbahevans/hooopstats/db"
)

type PlayerService struct {
	db *db.DB
}

func NewPlayerService(db *db.DB) *PlayerService {
	return &PlayerService{db}
}

func (s *PlayerService) GetPlayers() ([]models.Player, error) {
	result, err := s.db.Query("SELECT * FROM players")
	if err != nil {
		return nil, err
	}

	players := []models.Player{}
	for result.Next() {
		var player models.Player
		err = result.Scan(&player.Id, &player.FirstName, &player.LastName, &player.TeamID)
		if err != nil {
			return nil, err
		}

		players = append(players, player)
	}

	return players, nil
}

func (s *PlayerService) GetPlayerById(playerId int) (models.Player, error) {
	result := s.db.QueryRow("SELECT * FROM players WHERE id = $1", playerId)

	var player models.Player
	err := result.Scan(&player.Id, &player.FirstName, &player.LastName, &player.TeamID)
	if err != nil {
		return models.Player{}, ErrPlayerNotFound
	}

	return player, nil
}

func (s *PlayerService) CreatePlayer(data models.Player) (int64, error) {
	row := s.db.QueryRow(`
	INSERT INTO players (first_name, last_name, team_id)
	VALUES ($1, $2, $3)
	RETURNING id`, data.FirstName, data.LastName, data.TeamID)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
