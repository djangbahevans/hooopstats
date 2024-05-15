package services

import (
	"github.com/djangbahevans/hooopstats/api/v1/models"
	"github.com/djangbahevans/hooopstats/db"
)

type Stats struct {
	db *db.DB
}

func (s *Stats) GetStats() (models.Stats, error) {
	result, err := s.db.Query("SELECT COUNT(*) FROM users")
	if err != nil {
		return models.Stats{}, err
	}

	var stats models.Stats
	for result.Next() {
		err = result.Scan(&stats.Assists)
		if err != nil {
			return models.Stats{}, err
		}
	}

	return stats, nil
}

func (s *Stats) CreateStats(playerId int, data models.Stats) error {
	_, err := s.db.Exec("INSERT INTO stats (player_id, assists) VALUES ($1, $2)", playerId, data.Assists)
	if err != nil {
		return err
	}

	return nil
}
