package user

import (
	"net/http"
	"user-rest-api/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	// TODO Logger
	// TODO Service
}

const (
	usersURL       = "/users"
	usersParamsURL = "/users/:uuid"
)

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(r *httprouter.Router) {
	r.GET(usersURL, h.GetUsersList)
	r.GET(usersParamsURL, h.GetUserByUUID)
	r.POST(usersURL, h.CreateUser)
	r.PUT(usersParamsURL, h.UpdateUser)
	r.DELETE(usersParamsURL, h.DeleteUser)
}

func (h *handler) GetUsersList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("GetUsersList"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("GetUserByUUID"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("CreateUser"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("DeleteUser"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("UpdateUser"))
}
