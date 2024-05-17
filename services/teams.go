package services

import (
	"github.com/djangbahevans/hooopstats/api/v1/models"
	"github.com/djangbahevans/hooopstats/db"
)

type TeamService struct {
	db *db.DB
}

func NewTeamService(db *db.DB) *TeamService {
	return &TeamService{db}
}

func (s *TeamService) GetTeams() ([]models.Team, error) {
	result, err := s.db.Query("SELECT * FROM teams")
	if err != nil {
		return nil, err
	}

	teams := []models.Team{}
	for result.Next() {
		var team models.Team
		err = result.Scan(&team.Id, &team.Name)
		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	return teams, nil
}

func (s *TeamService) GetTeamById(teamId int) (models.Team, error) {
	result := s.db.QueryRow("SELECT * FROM teams WHERE id = $1", teamId)

	var team models.Team
	err := result.Scan(&team.Id, &team.Name)
	if err != nil {
		return models.Team{}, ErrTeamNotFound
	}

	return team, nil
}

func (s *TeamService) CreateTeam(data models.Team) (int64, error) {
	row := s.db.QueryRow(`
	INSERT INTO teams (name)
	VALUES ($1)
	RETURNING id`, data.Name)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
