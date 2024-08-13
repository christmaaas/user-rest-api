package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
	"user-rest-api/internal/config"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

const network = "tcp"

func Run(router *httprouter.Router, cfg *config.Config) {
	logger := logger.GetLogger()

	logger.Infof("bind application to %s:%s", cfg.Connection.BindIP, cfg.Connection.Port)

	listener, err := net.Listen(network, fmt.Sprintf("%s:%s", cfg.Connection.BindIP, cfg.Connection.Port))
	if err != nil {
		logger.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("app initialized and started")

	if err := server.Serve(listener); err != nil {
		logger.Fatal(err)
	}
}
