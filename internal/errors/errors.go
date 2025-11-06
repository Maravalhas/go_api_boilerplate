package errors

import (
	systemError "errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type Error uint8

const (
	BadRequest Error = iota
	Unauthorized
	Forbidden
	NotFound
	Duplicated
	InternalError
)

var errors = map[Error]ResponseError{
	BadRequest: {
		StatusCode: http.StatusBadRequest,
		Message:    "Bad Request",
	},
	Unauthorized: {
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
	},
	Forbidden: {
		StatusCode: http.StatusForbidden,
		Message:    "Forbidden",
	},
	NotFound: {
		StatusCode: http.StatusNotFound,
		Message:    "Entity Not Found",
	},
	Duplicated: {
		StatusCode: http.StatusConflict,
		Message:    "Entity Already Exists",
	},
	InternalError: {
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	},
}

func New(error Error) error {
	return systemError.New(errors[error].Message)
}

func Equal(err error, err2 Error) bool {
	return err.Error() == errors[err2].Message
}

func NewResponseError(error Error) (int, map[string]any) {
	err := errors[error]
	return err.StatusCode, gin.H{"status": err.StatusCode, "message": err.Message}
}

func NewResponseErrorWithMessage(error Error, message string) (int, map[string]any) {
	err := errors[error]
	err.Message = fmt.Sprintf("%s: %s", err.Message, message)
	return err.StatusCode, gin.H{"status": err.StatusCode, "message": err.Message}
}
