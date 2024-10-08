package handlers

import (
	"errors"
	"net/http"
	"user-rest-api/internal/apperror"

	"github.com/julienschmidt/httprouter"
)

type appHandle func(http.ResponseWriter, *http.Request, httprouter.Params) error

func UserMiddleware(h appHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := h(w, r, p)
		if err != nil {
			var appErr *apperror.AppError
			if errors.As(err, &appErr) {
				switch {
				case errors.Is(err, apperror.ErrNotFound):
					w.WriteHeader(http.StatusNotFound)
					w.Write(apperror.ErrNotFound.Marshal())

				case errors.Is(err, apperror.ErrValidation):
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write(apperror.ErrValidation.Marshal())

				case errors.Is(err, apperror.ErrUnauthorized):
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(apperror.ErrUnauthorized.Marshal())

				case errors.Is(err, apperror.ErrForbidden):
					w.WriteHeader(http.StatusForbidden)
					w.Write(apperror.ErrForbidden.Marshal())

				case errors.Is(err, apperror.ErrConflict):
					w.WriteHeader(http.StatusConflict)
					w.Write(apperror.ErrConflict.Marshal())

				default:
					w.WriteHeader(http.StatusBadRequest)
					w.Write(appErr.Marshal())
				}
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(apperror.ErrInternalSystem.Marshal())
		}
	}
}
