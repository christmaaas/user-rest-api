package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-rest-api/internal/service"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type usersHandler struct {
	logger  logger.Logger
	service *service.UsersService
}

const (
	usersURL       = "/api/users"
	usersParamsURL = "/api/users/:uuid"
)

func NewUserHandler(logger logger.Logger, service *service.UsersService) Handler {
	return &usersHandler{
		logger:  logger,
		service: service,
	}
}

func (h *usersHandler) Register(r *httprouter.Router) {
	r.GET(usersURL, AppErrorMiddleware(h.GetUserByUUID))
	r.GET(usersParamsURL, AppErrorMiddleware(h.GetUserByUUID))
	r.POST(usersURL, AppErrorMiddleware(h.CreateUser))
	r.PUT(usersParamsURL, AppErrorMiddleware(h.UpdateUser))
	r.DELETE(usersParamsURL, AppErrorMiddleware(h.DeleteUser))
}

func (h *usersHandler) GetUsersList(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Write([]byte("GetUsersList"))
	return nil
}

func (h *usersHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("get user by UUID")
	w.Header().Set("Content-Type", "application/json")

	userUUID := p.ByName("uuid")

	h.logger.Debug("fetching user by UUID")
	user, err := h.service.GetOneUser(r.Context(), userUUID)
	if err != nil {
		return err
	}

	h.logger.Debug("marshalling user")
	userBytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshall user. error: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
	return nil
}

func (h *usersHandler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Write([]byte("CreateUser"))
	return nil
}

func (h *usersHandler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Write([]byte("DeleteUser"))
	return nil
}

func (h *usersHandler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	w.Write([]byte("UpdateUser"))
	return nil
}
