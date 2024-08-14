package handlers

import (
	"encoding/json"
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
	r.GET(usersURL, h.GetUsersList)
	r.GET(usersParamsURL, h.GetUserByUUID)
	r.POST(usersURL, h.CreateUser)
	r.PUT(usersParamsURL, h.UpdateUser)
	r.DELETE(usersParamsURL, h.DeleteUser)
}

func (h *usersHandler) GetUsersList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("GetUsersList"))
}

func (h *usersHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h.logger.Info("GET USER BY UUID")
	w.Header().Set("Content-Type", "application/json")

	userUUID := p.ByName("uuid")

	h.logger.Debug("fetching user by UUID")
	user, err := h.service.GetOneUser(r.Context(), userUUID)
	if err != nil {
		h.logger.Fatal("asd")
	}

	h.logger.Debug("marshalling user")
	userBytes, err := json.Marshal(user)
	if err != nil {
		h.logger.Fatal("asd")
	}

	w.WriteHeader(http.StatusOK)
	_, writeErr := w.Write(userBytes)
	if writeErr != nil {
		h.logger.Fatal("asd")
	}
}

func (h *usersHandler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("CreateUser"))
}

func (h *usersHandler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("DeleteUser"))
}

func (h *usersHandler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("UpdateUser"))
}
