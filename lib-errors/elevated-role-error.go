package lib_errors

import "net/http"

type ElevatedRoleError struct {
	status  int
	message string
}

func (e ElevatedRoleError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    INVALID_ROLE,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e ElevatedRoleError) ErrorStatus() int {
	return e.status
}

func NewElevatedRoleError() *ElevatedRoleError {
	return &ElevatedRoleError{status: http.StatusUnauthorized, message: "Invalid user role"}
}
