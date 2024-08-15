package error

import "encoding/json"

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message"`
	DeveloperMessage string `json:"developer_message"`
	Code             string `json:"code"`
}

func NewAppError(err error, message, develoeprMessage, code string) *AppError {
	return &AppError{
		Err:              err,
		Message:          message,
		DeveloperMessage: develoeprMessage,
		Code:             code,
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
