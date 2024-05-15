package db

import (
	"github.com/djangbahevans/hooopstats/api/v1/models"
)

type StatsRepository interface {
	CreateStats(stats *models.CreateStatsRequest) (*models.StatsResponse, error)
	GetStats() (*models.Stats, error)
	GetStatsByPlayerID(playerID int) (*models.Stats, error)
}

type StatsRepo struct {
	db DB
}

func NewStatsRepo(db DB) *StatsRepo {
	return &StatsRepo{db}
}

func (r *StatsRepo) CreateStats(stats *models.CreateStatsRequest) (*models.StatsResponse, error) {
	_, err := r.db.Exec(`
	INSERT INTO stats (
    points,
    rebounds,
    assists,
    steals,
    blocks,
    fouls,
    turnovers,
    minutes_played,
    player_id
		)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		stats.Points, stats.Rebounds, stats.Assists,
		stats.Steals, stats.Blocks, stats.Fouls,
		stats.Turnovers, stats.MinutesPlayed, stats.PlayerID)
	if err != nil {
		return nil, err
	}

	return &models.StatsResponse{
		Status:  "success",
		Message: "Stats created",
		PlayerStats: models.Stats{
			Points:        stats.Points,
			Rebounds:      stats.Rebounds,
			Assists:       stats.Assists,
			Steals:        stats.Steals,
			Blocks:        stats.Blocks,
			Fouls:         stats.Fouls,
			Turnovers:     stats.Turnovers,
			MinutesPlayed: stats.MinutesPlayed,
		},
	}, nil
}
