package main

import (
	"github.com/TemurMannonov/crawler_task/api"
	"github.com/TemurMannonov/crawler_task/config"

	"github.com/TemurMannonov/crawler_task/pkg/logger"
)

var (
	log logger.Logger
	cfg config.Config
)

func main() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "crawler")

	server := api.New(api.Config{
		Logger: log,
		Cfg:    cfg,
	})

	err := server.Run(cfg.HttpPort)
	if err != nil {
		panic(err)
	}
}
