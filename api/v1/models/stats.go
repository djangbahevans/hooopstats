package models

// Stats is a struct that represents stats
type Stats struct {
	Points        int `json:"points"`
	Rebounds      int `json:"rebounds"`
	Assists       int `json:"assists"`
	Steals        int `json:"steals"`
	Blocks        int `json:"blocks"`
	Fouls         int `json:"fouls"`
	Turnovers     int `json:"turnovers"`
	MinutesPlayed int `json:"minutes_played"`
}

type CreateStatsRequest struct {
	*Stats
	PlayerID int `json:"player_id"`
}

type StatsResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	PlayerStats Stats  `json:"player_stats"`
}
