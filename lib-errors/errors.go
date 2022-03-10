package lib_errors

const (
	NOT_FOUND          = "NOT_FOUND"
	BAD_REQUEST        = "BAD_REQUEST"
	NOT_AUTHORIZED     = "NOT_AUTHORIZED"
	INVALID_ROLE       = "INVALID_ROLE"
	UNKNOWN            = "UNKNOWN"
	INTERNAL           = "INTERNAL"
	EVENT_PUBLISH      = "EVENT_PUBLISH"
	PAYLOAD_VALIDATION = "PAYLOAD_VALIDATION"
	PAYLOAD_ENCRYPTION = "PAYLOAD_ENCRYPTION"
	ENV                = "ENV"
)

type CustomError interface {
	GetError() []ErrorObject
	ErrorStatus() int
}

type Error struct {
	Type    string
	Message string
	Status  int
}
type ErrorObject struct {
	error Error
}
