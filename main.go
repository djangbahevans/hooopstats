package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/djangbahevans/hooopstats/api"
	"github.com/djangbahevans/hooopstats/db"
)

func main() {
	database := db.NewDB()
	defer database.Close()

	server := api.NewAPIServer(":8000", database)
	server.Run()

	go func() {
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
