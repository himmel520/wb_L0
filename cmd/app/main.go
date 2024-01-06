package main

import (
	"log"
	"net/http"

	"github.com/caarlos0/env/v10"
	"github.com/himmel520/wb_L0/internal/config"
	"github.com/himmel520/wb_L0/internal/handlers"
	"github.com/himmel520/wb_L0/internal/nats"
	"github.com/himmel520/wb_L0/internal/repository"
	"github.com/himmel520/wb_L0/internal/repository/postgres"
	"github.com/himmel520/wb_L0/internal/repository/redis"
	"github.com/himmel520/wb_L0/internal/services"
)

func main() {
	cfg := config.NewConfig()
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	consumer, err := nats.NewNats(&cfg.Nats)
	if err != nil {
		log.Fatalf("[Nats] %v", err)
	}

	db, err := postgres.NewPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("[Pg] %v", err)
	}
	defer db.Close()

	rdb, err := redis.NewRedis(cfg.Redis)
	if err != nil {
		log.Fatalf("[Redis] %v", err)
	}
	defer rdb.Close()

	repo := repository.NewRepository(db, rdb)
	services := services.NewService(repo)
	handlers := handlers.NewHandler(services, consumer)

	cc := handlers.InitNatsConsume()
	defer cc.Stop()

	log.Printf("Server is running on %v..\n", cfg.App.Addr)
	http.ListenAndServe(cfg.App.Addr, handlers.InitRoutes())
}
