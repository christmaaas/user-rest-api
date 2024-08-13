package handlers

import (
	"net/http"
	"user-rest-api/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

type usersHandler struct {
	logger logger.Logger
	// TODO UserService
}

const (
	usersURL       = "/users"
	usersParamsURL = "/users/:uuid"
)

func NewUserHandler(logger logger.Logger) Handler {
	return &usersHandler{
		logger: logger,
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
	w.Write([]byte("GetUserByUUID"))
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
