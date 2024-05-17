package models

type Player struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TeamID    int64  `json:"team_id"`
}

type CreatePlayerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TeamID    int64  `json:"team_id"`
}
