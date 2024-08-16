package apperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ErrInternalSystem = NewAppError(http.StatusInternalServerError, "00100", "internal system error")
	ErrBadRequest     = NewAppError(http.StatusBadRequest, "00101", "bad request")
	ErrValidation     = NewAppError(http.StatusBadRequest, "00102", "validation error")
	ErrNotFound       = NewAppError(http.StatusNotFound, "00103", "not found")
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "00104", "unauthorized")
	ErrForbidden      = NewAppError(http.StatusForbidden, "00105", "access forbidden")
	ErrConflict       = NewAppError(http.StatusConflict, "00106", "resource conflict")
)

type AppError struct {
	Err           error  `json:"-"`
	Message       string `json:"message"`
	Code          string `json:"code"`
	TransportCode int    `json:"transport_code"`
}

func NewAppError(transportCode int, code, message string) *AppError {
	return &AppError{
		Err:           fmt.Errorf(message),
		Message:       message,
		Code:          code,
		TransportCode: transportCode,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bytes
}
