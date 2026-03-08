package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(statusCode int, message string) *AppError {
	return &AppError{StatusCode: statusCode, Message: message}
}

func Newf(statusCode int, format string, args ...any) *AppError {
	return &AppError{StatusCode: statusCode, Message: fmt.Sprintf(format, args...)}
}

var (
	ErrBadRequest      = New(http.StatusBadRequest, "Bad Request")
	ErrUnauthorized    = New(http.StatusUnauthorized, "Unauthorized")
	ErrForbidden       = New(http.StatusForbidden, "Forbidden")
	ErrNotFound        = New(http.StatusNotFound, "Entity Not Found")
	ErrDuplicated      = New(http.StatusConflict, "Entity Already Exists")
	ErrInvalidFileType = New(http.StatusUnsupportedMediaType, "Invalid file type")
)
