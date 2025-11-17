package app

import (
	"elestial/config"
	dbpac "elestial/internal/db"
	"elestial/internal/handler"
	"elestial/internal/logger"
	"elestial/internal/repository"
	"elestial/internal/server"
	"elestial/internal/service"
	"log"
)

func Run(cfg *config.Config) {

	logger.InitLogger()

	db, err := dbpac.NewPool(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	if err := dbpac.RunMigrations(cfg.PG.URL); err != nil {
		log.Fatalf("Unable to run migrations: %v", err)
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository, cfg)

	handler := handler.NewHandler(service)

	server := new(server.Server)

	if err := server.Run(cfg.HTTP.Port, handler.InitRoutes()); err != nil {
		log.Fatalf("Error running server: %s\n", err)
	}
}
