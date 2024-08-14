package app

import (
	"context"
	"user-rest-api/internal/config"
	"user-rest-api/internal/handlers"
	"user-rest-api/internal/repository"
	"user-rest-api/internal/repository/postgres"
	"user-rest-api/pkg/dbclient"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	cfg     *config.Config
	handler handlers.Handler
	//service service.Service
	repository repository.Users
}

func NewApp() (App, error) {
	logger.InitLogger()
	logger := logger.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	postgreSQLClient := dbclient.NewClient(context.TODO(), cfg.Storage)
	repository := postgres.NewUsersRepo(postgreSQLClient, logger)

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("handler initializing")
	handler := handlers.NewUserHandler(logger)
	handler.Register(router)

	return App{
		cfg:        cfg,
		handler:    handler,
		repository: repository,
	}, nil
}

// Run initializes whole application.
func (a *App) Run(ctx context.Context) error {
	// TODO
	return nil
}
