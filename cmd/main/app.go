package main

import (
	"net/http"
	"user-rest-api/internal/user"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger.InitLogger()
	logger := logger.GetLogger()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("handler initializing")
	handler := user.NewHandler(logger)
	handler.Register(router)

	logger.Info("server running...")
	http.ListenAndServe(":8080", router)
}
