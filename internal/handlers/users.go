package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"user-rest-api/internal/apperror"
	"user-rest-api/internal/domain"
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
	r.GET(usersURL, UserMiddleware(h.GetUsersList))
	r.GET(usersParamsURL, UserMiddleware(h.GetUserByUUID))
	r.POST(usersURL, UserMiddleware(h.CreateUser))
	r.PUT(usersParamsURL, UserMiddleware(h.UpdateUser))
	r.DELETE(usersParamsURL, UserMiddleware(h.DeleteUser))
}

func (h *usersHandler) GetUsersList(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("GET all users")
	w.Header().Set("Content-Type", "application/json")

	h.logger.Debug("fetching all users")
	users, err := h.service.GetAllUsers(r.Context())
	if err != nil {
		return err
	}

	h.logger.Debug("marshalling users")
	userBytes, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshall user. error: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(userBytes)
	return nil
}

func (h *usersHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("GET user by UUID")
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
	h.logger.Info("CREATE user")
	w.Header().Set("Content-Type", "application/json")

	h.logger.Debug("decode create user dto")
	var dto domain.CreateUserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.ErrBadRequest
	}

	h.logger.Debug("creating user with UUID")
	userUUID, err := h.service.CreateUser(r.Context(), dto)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("%s/%s", usersURL, userUUID))
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *usersHandler) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("DELETE user")
	w.Header().Set("Content-Type", "application/json")

	userUUID := p.ByName("uuid")

	h.logger.Debug("deleting user with uuid")
	err := h.service.DeleteUser(r.Context(), userUUID)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h *usersHandler) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	h.logger.Info("UPDATE user")
	w.Header().Set("Content-Type", "application/json")

	userUUID := p.ByName("uuid")

	h.logger.Debug("decode update user dto")
	var dto domain.UpdateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.ErrBadRequest
	}

	h.logger.Debug("updating user with uuid")
	err := h.service.UpdateUser(r.Context(), userUUID, dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
