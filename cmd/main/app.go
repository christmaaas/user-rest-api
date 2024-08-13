package main

import (
	"user-rest-api/internal/config"
	"user-rest-api/internal/handlers"
	"user-rest-api/internal/server"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger.InitLogger()
	logger := logger.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("handler initializing")
	handler := handlers.NewUserHandler(logger)
	handler.Register(router)

	logger.Info("application running")
	server.Run(router, cfg)
}
