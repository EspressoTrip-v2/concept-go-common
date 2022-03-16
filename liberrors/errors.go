package libErrors

import (
	"fmt"
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
	PAYLOAD_VALIDATION ErrorTypes = "PAYLOAD_VALIDATION"
	PAYLOAD_ENCRYPTION ErrorTypes = "PAYLOAD_ENCRYPTION"
	ENV                ErrorTypes = "ENV"
)

type CustomError struct {
	ErrorType ErrorTypes
	Status    int
	Message   string
}

type Error struct {
	Type    ErrorTypes `json:"type"`
	Message string     `json:"message"`
	Status  int        `json:"status"`
}
type ErrorObject struct {
	Error Error `json:"error"`
}

func (e CustomError) GetErrors() []ErrorObject {
	return []ErrorObject{{Error: Error{
		Type:    e.ErrorType,
		Message: e.Message,
		Status:  e.Status,
	}}}
}

func (e CustomError) ErrorStatus() int {
	return e.Status
}

func (e CustomError) Type() ErrorTypes {
	return e.ErrorType
}

func NewBadRequestError(msg string) *CustomError {
	return &CustomError{
		ErrorType: BAD_REQUEST,
		Status:    http.StatusBadRequest,
		Message:   msg,
	}
}

func NewDatabaseError(msg string) *CustomError {
	return &CustomError{
		ErrorType: INTERNAL,
		Status:    http.StatusInternalServerError,
		Message:   msg,
	}
}

func NewRabbitConnectionError(msg string) *CustomError {
	return &CustomError{
		ErrorType: INTERNAL,
		Status:    http.StatusInternalServerError,
		Message:   msg,
	}
}

func NewElevatedAuthError() *CustomError {
	return &CustomError{
		ErrorType: INVALID_ROLE,
		Status:    http.StatusUnauthorized,
		Message:   "Invalid user role",
	}
}

func NewEventPublisherError(msg string) *CustomError {
	return &CustomError{
		ErrorType: EVENT_PUBLISH,
		Status:    http.StatusInternalServerError,
		Message:   msg,
	}
}

func NewNotAuthorizedError() *CustomError {
	return &CustomError{
		ErrorType: NOT_AUTHORIZED,
		Status:    http.StatusUnauthorized,
		Message:   "Not unauthorized",
	}
}

func NewNotFoundError(msg string) *CustomError {
	return &CustomError{
		ErrorType: NOT_FOUND,
		Status:    http.StatusNotFound,
		Message:   msg,
	}
}

func NewPayloadEncryptionError() *CustomError {
	return &CustomError{
		ErrorType: PAYLOAD_ENCRYPTION,
		Status:    http.StatusBadRequest,
		Message:   "JWT payload encryption error",
	}
}

func NewPayloadValidationError() *CustomError {
	return &CustomError{
		ErrorType: PAYLOAD_VALIDATION,
		Status:    http.StatusBadRequest,
		Message:   "Validation error",
	}
}

func NewEnvError() *CustomError {
	return &CustomError{
		ErrorType: ENV,
		Status:    http.StatusNotImplemented,
		Message:   "Missing required env variable",
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
