package logcodes

type LogCodes string

const (
	CREATED LogCodes = "created"
	DELETED LogCodes = "deleted"
	UPDATED LogCodes = "updated"
	ERROR   LogCodes = "error"
	INFO    LogCodes = "info"
)
