package lib_errors

import "net/http"

type NotFoundError struct {
	status  int
	message string
}

func (e NotFoundError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    NOT_FOUND,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e NotFoundError) ErrorStatus() int {
	return e.status
}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{status: http.StatusNotFound, message: "Not found"}
}
