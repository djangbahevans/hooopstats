package models

// Stats is a struct that represents stats
type Stats struct {
	Id            int64   `json:"id"`
	Points        int     `json:"points"`
	Rebounds      int     `json:"rebounds"`
	Assists       int     `json:"assists"`
	Steals        int     `json:"steals"`
	Blocks        int     `json:"blocks"`
	Fouls         int     `json:"fouls"`
	Turnovers     int     `json:"turnovers"`
	MinutesPlayed float64 `json:"minutes_played"`
	PlayerID      int64   `json:"player_id"`
	GameNumber    int64   `json:"game_number"`
}

type CreateStatsRequest struct {
	Id            int64   `json:"id"`
	Points        int     `json:"points"`
	Rebounds      int     `json:"rebounds"`
	Assists       int     `json:"assists"`
	Steals        int     `json:"steals"`
	Blocks        int     `json:"blocks"`
	Fouls         int     `json:"fouls"`
	Turnovers     int     `json:"turnovers"`
	MinutesPlayed float64 `json:"minutes_played"`
	PlayerID      int64   `json:"player_id"`
	GameNumber    int64   `json:"game_number"`
}

type StatsAverage struct {
	Points        float64 `json:"points"`
	Rebounds      float64 `json:"rebounds"`
	Assists       float64 `json:"assists"`
	Steals        float64 `json:"steals"`
	Blocks        float64 `json:"blocks"`
	Fouls         float64 `json:"fouls"`
	Turnovers     float64 `json:"turnovers"`
	MinutesPlayed float64 `json:"minutes_played"`
}
