package main

import (
	"crudTestTask/env"
	"crudTestTask/internal/repository"
	"crudTestTask/server"
	"log"
)

func main() {
	cfg, err := env.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	rep, err := repository.New(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	httpServer := server.New(rep)
	err = httpServer.Start(cfg)
	if err != nil {
		log.Fatalf("Error starting http server: %v", err)
	}
}
