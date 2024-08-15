package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
	"user-rest-api/internal/config"
	"user-rest-api/internal/handlers"
	"user-rest-api/internal/repository"
	"user-rest-api/internal/repository/postgres"
	"user-rest-api/internal/service"
	"user-rest-api/pkg/dbclient"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	cfg        *config.Config
	handler    http.Handler
	service    *service.UsersService
	repository repository.Users
}

// NewApp initializes the whole application
func NewApp() App {
	logger.InitLogger()
	logger := logger.GetLogger()
	logger.Info("initializing application")

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

	return App{
		cfg:        cfg,
		handler:    router,
		service:    usersService,
		repository: repository,
	}
}

// startHTTP initializes and start http server
// TODO: Handle HTTP server shutdown gracefully
// by goroutine, channel and signals
func (a *App) startHTTP() error {
	logger := logger.GetLogger()
	logger.Info("HTTP server initializing")

	logger.Infof("bind application to %s:%s", a.cfg.Connection.BindIP, a.cfg.Connection.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Connection.BindIP, a.cfg.Connection.Port))
	if err != nil {
		return err
	}

	httpServer := &http.Server{
		Handler:      a.handler,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("application initialized and started")
	if err := httpServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

// Run starts the application.
// TODO: Handle HTTP server shutdown gracefully
// by goroutine, channel and signals
func Run() {
	a := NewApp()
	logger := logger.GetLogger()

	if err := a.startHTTP(); err != nil {
		logger.Fatalf("unable to start application due to error: %w", err)
	}
}
