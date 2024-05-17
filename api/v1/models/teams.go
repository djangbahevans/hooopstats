package models

type Team struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateTeamRequest struct {
	Name string `json:"name"`
}
