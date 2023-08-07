package main

import (
	"github.com/Shakhrik/video_task/api"
	"github.com/Shakhrik/video_task/config"
	"github.com/Shakhrik/video_task/pkg/logger"
	"github.com/Shakhrik/video_task/service"
	"github.com/Shakhrik/video_task/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.Load()
	logger := logger.New(cfg.LogLevel, "safia_catalogue_service")

	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	// initialize storage
	storage := storage.NewStorage(db)
	// initialize service
	service := service.NewService(storage)
	// initialize apiServer
	apiServer := api.New(&api.RouterOptions{
		Log: logger,
		// Cfg:     &cfg,
		Service: service,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		panic(err)
	}

}
