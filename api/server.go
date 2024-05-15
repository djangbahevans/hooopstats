package api

import (
	"log"
	"net/http"

	v1 "github.com/djangbahevans/hooopstats/api/v1"
	"github.com/djangbahevans/hooopstats/db"
	"github.com/djangbahevans/hooopstats/middleware"
)

type APIServer struct {
	addr string
	db   *db.DB
}

func NewAPIServer(addr string, db *db.DB) *APIServer {
	return &APIServer{addr, db}
}

func (s *APIServer) Run() error {
	r := http.NewServeMux()
	v1.RegisterV1Handlers(r, s.db)

	api := http.NewServeMux()
	api.Handle("/api/", http.StripPrefix("/api", r))

	server := http.Server{
		Addr:    s.addr,
		Handler: middleware.RequestLoggerMiddleware(api),
	}

	log.Printf("Server listening on %s", s.addr)

	return server.ListenAndServe()
}
