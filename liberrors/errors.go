package libErrors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"net/http"
)

type ErrorTypes string

const (
	NOT_FOUND          ErrorTypes = "NOT_FOUND"
	BAD_REQUEST        ErrorTypes = "BAD_REQUEST"
	NOT_AUTHORIZED     ErrorTypes = "NOT_AUTHORIZED"
	INVALID_ROLE       ErrorTypes = "INVALID_ROLE"
	INTERNAL           ErrorTypes = "INTERNAL"
	EVENT_PUBLISH      ErrorTypes = "EVENT_PUBLISH"
	EVENT_CONSUMER     ErrorTypes = "EVENT_CONSUMER"
	PAYLOAD_VALIDATION ErrorTypes = "PAYLOAD_VALIDATION"
	PAYLOAD_ENCRYPTION ErrorTypes = "PAYLOAD_ENCRYPTION"
	ENV                ErrorTypes = "ENV"
)

type CustomError struct {
	ErrorType ErrorTypes
	Status    int
	Message   []string
}

type ErrorObject struct {
	Status int        `json:"status"`
	Type   ErrorTypes `json:"type"`
	Error  []string   `json:"error"`
}

func (e CustomError) GetErrors() ErrorObject {
	var errorObj ErrorObject
	for _, msg := range e.Message {
		errorObj.Error = append(errorObj.Error, msg)
	}
	errorObj.Type = e.ErrorType
	errorObj.Status = e.Status

	return errorObj
}

func (e CustomError) ErrorStatus() int {
	return e.Status
}

func (e CustomError) Type() ErrorTypes {
	return e.ErrorType
}

func NewBadRequestError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: BAD_REQUEST,
		Status:    http.StatusBadRequest,
		Message:   m,
	}
}

func NewDatabaseError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: INTERNAL,
		Status:    http.StatusInternalServerError,
		Message:   m,
	}
}

func NewRabbitConnectionError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: INTERNAL,
		Status:    http.StatusInternalServerError,
		Message:   m,
	}
}

func NewElevatedAuthError() *CustomError {
	return &CustomError{
		ErrorType: INVALID_ROLE,
		Status:    http.StatusUnauthorized,
		Message:   []string{"Invalid user role"},
	}
}

func NewEventPublisherError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: EVENT_PUBLISH,
		Status:    http.StatusInternalServerError,
		Message:   m,
	}
}

func NewEventConsumerError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: EVENT_CONSUMER,
		Status:    http.StatusInternalServerError,
		Message:   m,
	}
}

func NewNotAuthorizedError() *CustomError {
	return &CustomError{
		ErrorType: NOT_AUTHORIZED,
		Status:    http.StatusUnauthorized,
		Message:   []string{"Not unauthorized"},
	}
}

func NewNotFoundError(msg string) *CustomError {
	m := []string{msg}
	return &CustomError{
		ErrorType: NOT_FOUND,
		Status:    http.StatusNotFound,
		Message:   m,
	}
}

func NewPayloadEncryptionError() *CustomError {
	return &CustomError{
		ErrorType: PAYLOAD_ENCRYPTION,
		Status:    http.StatusBadRequest,
		Message:   []string{"JWT payload encryption error"},
	}
}

func NewPayloadValidationError(validationErrors validator.ValidationErrors) *CustomError {
	ce := CustomError{}
	ce.ErrorType = PAYLOAD_VALIDATION
	ce.Status = http.StatusBadRequest
	for _, validationError := range validationErrors {
		ce.Message = append(ce.Message, validationError.Error())
	}
	return &ce
}

func NewEnvError() *CustomError {
	return &CustomError{
		ErrorType: ENV,
		Status:    http.StatusNotImplemented,
		Message:   []string{"Missing required env variable"},
	}
}

func GrpcTranslator(grpcStatus error) *CustomError {
	respErr, ok := status.FromError(grpcStatus)
	fmt.Printf("Details: %v", respErr.Details())
	if ok == true {
		switch eTYpe := respErr.Code(); eTYpe {
		case codes.AlreadyExists, codes.Unknown, codes.InvalidArgument:
			return NewBadRequestError(respErr.Message())
		case codes.NotFound:
			return NewNotFoundError(respErr.Message())
		case codes.PermissionDenied:
			return NewNotAuthorizedError()
		case codes.Internal:
			return NewDatabaseError(respErr.Message())
		default:
			return NewBadRequestError(respErr.Message())
		}
	} else {
		return NewBadRequestError(grpcStatus.Error())
	}
}
