package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tapan176/go-practice/config"
	"github.com/Tapan176/go-practice/internal"
	routes "github.com/Tapan176/go-practice/webApiV1"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.GetConfig()

	db, err := internal.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	defer db.Close()

	mainRouter := http.NewServeMux()

	mainRouter.Handle("/v1/", http.StripPrefix("/v1", routes.IndexRouter(db)))

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port),
		Handler: mainRouter,
	}

	log.Printf("Server listening to http://%s:%d/v1\n", cfg.Http.Host, cfg.Http.Port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
}
