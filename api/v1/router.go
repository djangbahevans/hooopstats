package v1

import (
	"net/http"

	"github.com/djangbahevans/hooopstats/api/v1/handlers"
	"github.com/djangbahevans/hooopstats/db"
)

func RegisterV1Handlers(r *http.ServeMux, db *db.DB) {
	v1 := http.NewServeMux()
	handlers.RegisterStatHandler(v1, db)
	handlers.RegisterTeamHandler(v1, db)
	handlers.RegisterPlayerHandler(v1, db)
	
	r.Handle("/v1/", http.StripPrefix("/v1", v1))
}
