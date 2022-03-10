package lib_errors

import "net/http"

type NotAuthorizedError struct {
	status  int
	message string
}

func (e NotAuthorizedError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    NOT_AUTHORIZED,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e NotAuthorizedError) ErrorStatus() int {
	return e.status
}

func NewNotAuthorizedError() *NotAuthorizedError {
	return &NotAuthorizedError{status: http.StatusUnauthorized, message: "Not authorized"}
}
