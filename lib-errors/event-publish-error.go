package lib_errors

import "net/http"

type EventPublishError struct {
	status  int
	message string
}

func (e EventPublishError) GetError() []ErrorObject {
	errObj := ErrorObject{
		error: Error{
			Type:    EVENT_PUBLISH,
			Message: e.message,
			Status:  e.status,
		},
	}
	return []ErrorObject{errObj}
}

func (e EventPublishError) ErrorStatus() int {
	return e.status
}

func NewEventPublishError(message string) *EventPublishError {
	return &EventPublishError{status: http.StatusInternalServerError, message: message}
}
