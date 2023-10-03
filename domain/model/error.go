package model

import (
	"encoding/json"
	"fmt"
)

const (
	Unauthorized        = 1
	BadRequest          = 2
	InternalServerError = 3
	ApiError            = 4
	ClientError         = 5
	Repository          = 6
	Database            = 7
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	StatusCode       int    `json:"status_code,omitempty"`
}

func NewAppError(message, developerMessage string, statusCode int) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Message:          message,
		DeveloperMessage: developerMessage,
		StatusCode:       statusCode,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
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
