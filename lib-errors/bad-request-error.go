package lib_errors

import "net/http"

type BadRequestError struct {
	status  int
	message string
}

func (e BadRequestError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    BAD_REQUEST,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e BadRequestError) ErrorStatus() int {
	return e.status
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{status: http.StatusBadRequest, message: message}
}
