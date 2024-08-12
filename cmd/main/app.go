package main

import (
	"user-rest-api/internal/config"
	"user-rest-api/internal/server"
	"user-rest-api/internal/user"
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
	handler := user.NewHandler(logger)
	handler.Register(router)

	logger.Info("application running")
	server.Run(router, cfg)
}
