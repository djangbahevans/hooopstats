package services

import "errors"

// Errors
var (
	ErrStatsNotFound  = errors.New("stats not found")
	ErrPlayerNotFound = errors.New("player not found")
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrTeamNotFound   = errors.New("team not found")
)
