package v1

import (
	"net/http"

	"github.com/djangbahevans/hooopstats/api/v1/handlers"
	"github.com/djangbahevans/hooopstats/db"
)

func RegisterV1Handlers(r *http.ServeMux, db *db.DB) {
	stats := http.NewServeMux()
	handlers.RegisterStatHandler(stats, db)
	
	r.Handle("/v1/", http.StripPrefix("/v1", stats))
}
