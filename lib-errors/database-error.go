package lib_errors

import "net/http"

type DatabaseError struct {
	status  int
	message string
}

func (e DatabaseError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    INTERNAL,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e DatabaseError) ErrorStatus() int {
	return e.status
}

func NewDatabaseError(message string) *DatabaseError {
	return &DatabaseError{status: http.StatusInternalServerError, message: message}
}
