package services

import (
	"github.com/djangbahevans/hooopstats/api/v1/models"
	"github.com/djangbahevans/hooopstats/db"
)

type StatsService struct {
	db *db.DB
}

func NewStatsService(db *db.DB) *StatsService {
	return &StatsService{db}
}

func (s *StatsService) GetStats() ([]models.Stats, error) {
	result, err := s.db.Query("SELECT * FROM stats")
	if err != nil {
		return nil, err
	}

	var stats []models.Stats
	for result.Next() {
		var stat models.Stats
		err = result.Scan(&stat.Points, &stat.Rebounds, &stat.Assists, &stat.Steals, &stat.Blocks, &stat.Fouls, &stat.Turnovers, &stat.MinutesPlayed)
		if err != nil {
			return nil, err
		}

		stats = append(stats, stat)
	}

	return stats, nil
}

func (s *StatsService) GetStatsByPlayerId(playerId int64) (models.Stats, error) {
	result := s.db.QueryRow("SELECT * FROM stats WHERE player_id = $1", playerId)

	var stats models.Stats
	err := result.Scan(
		&stats.Points,
		&stats.Rebounds,
		&stats.Assists,
		&stats.Steals,
		&stats.Blocks,
		&stats.Fouls,
		&stats.Turnovers,
		&stats.MinutesPlayed,
	)
	if err != nil {
		return models.Stats{}, ErrStatsNotFound
	}

	return stats, nil
}

func (s *StatsService) CreateStats(playerId int64, data models.Stats) (int64, error) {
	row := s.db.QueryRow(`
	INSERT INTO stats (
    points,
    rebounds,
    assists,
    steals,
    blocks,
    fouls,
    turnovers,
    minutes_played,
    player_id,
		game_number
  )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
		data.Points, data.Rebounds, data.Assists,
		data.Steals, data.Blocks, data.Fouls,
		data.Turnovers, data.MinutesPlayed, playerId,
		data.GameNumber,
	)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *StatsService) GetTeamStatsAverage(teamid int64) (models.StatsAverage, error) {
	results := s.db.QueryRow(`
	SELECT 
		SUM(s.points) as total_points,
		SUM(s.rebounds) as total_rebounds,
		SUM(s.assists) as total_assists,
		SUM(s.steals) as total_steals,
		SUM(s.blocks) as total_blocks,
		SUM(s.fouls) as total_fouls,
		SUM(s.turnovers) as total_turnovers,
		SUM(s.minutes_played) as total_minutes_played,
		COUNT(DISTINCT s.game_number) as games_played
	FROM stats s
	JOIN players p ON s.player_id = p.id
	WHERE p.team_id = $1
	GROUP BY p.team_id
	`, teamid)

	var stats models.StatsAverage
	var totalGames int
	err := results.Scan(
		&stats.Points,
		&stats.Rebounds,
		&stats.Assists,
		&stats.Steals,
		&stats.Blocks,
		&stats.Fouls,
		&stats.Turnovers,
		&stats.MinutesPlayed,
		&totalGames,
	)
	if err != nil {
		return models.StatsAverage{}, err
	}

	if totalGames == 0 {
		return models.StatsAverage{}, ErrStatsNotFound
	}

	stats.Points /= float64(totalGames)
	stats.Rebounds /= float64(totalGames)
	stats.Assists /= float64(totalGames)
	stats.Steals /= float64(totalGames)
	stats.Blocks /= float64(totalGames)
	stats.Fouls /= float64(totalGames)
	stats.Turnovers /= float64(totalGames)
	stats.MinutesPlayed /= float64(totalGames)

	return stats, nil
}

func (s *StatsService) GetPlayerStatsAverage(playerId int64) (models.Stats, error) {
	result := s.db.QueryRow(`
	SELECT AVG(points),
		AVG(rebounds),
		AVG(assists),
		AVG(steals),
		AVG(blocks),
		AVG(fouls),
		AVG(turnovers),
		AVG(minutes_played)
	FROM stats
	WHERE player_id = $1
`, playerId)

	var stats models.Stats
	err := result.Scan(
		&stats.Points,
		&stats.Rebounds,
		&stats.Assists,
		&stats.Steals,
		&stats.Blocks,
		&stats.Fouls,
		&stats.Turnovers,
		&stats.MinutesPlayed,
	)
	if err != nil {
		return models.Stats{}, err
	}

	return stats, nil
}
