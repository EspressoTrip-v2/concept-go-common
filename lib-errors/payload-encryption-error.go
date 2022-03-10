package lib_errors

import "net/http"

type PayloadEncryptionError struct {
	status  int
	message string
}

func (e PayloadEncryptionError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    PAYLOAD_ENCRYPTION,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e PayloadEncryptionError) ErrorStatus() int {
	return e.status
}

func NewPayloadEncryptionError() *PayloadEncryptionError {
	return &PayloadEncryptionError{status: http.StatusBadRequest, message: "JWT payload encryption error"}
}
