package main

import (
	"context"
	"user-rest-api/internal/config"
	"user-rest-api/internal/handlers"
	"user-rest-api/internal/repository/postgres"
	"user-rest-api/internal/server"
	"user-rest-api/internal/service"
	"user-rest-api/pkg/dbclient"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger.InitLogger()
	logger := logger.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	postgreSQLClient := dbclient.NewClient(context.TODO(), cfg.Storage)
	repository := postgres.NewUsersRepo(postgreSQLClient, logger)

	usersService := service.NewUsersService(logger, repository)

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("handler initializing")
	handler := handlers.NewUserHandler(logger, usersService)
	handler.Register(router)

	logger.Info("application running")
	server.Run(router, cfg)
}
