package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

// conform error interface
func (e AppError) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppError{Code: http.StatusNotFound, Message: message}
}

func NewUnexpectedError() error {
	return AppError{Code: http.StatusInternalServerError, Message: "unexpected error"}
}

func NewValidationError(message string) error {
	return AppError{Code: http.StatusUnprocessableEntity, Message: message}
}
