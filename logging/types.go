package logging

type LogCodes string

const (
	CREATED LogCodes = "created"
	DELETED LogCodes = "deleted"
	UPDATED LogCodes = "updated"
	ERROR   LogCodes = "error"
	INFO    LogCodes = "info"
)

type LogMsg struct {
	Service    string `json:"service"`
	LogContext string `json:"logContext"`
	Origin     string `json:"origin"`
	Message    string `json:"message"`
	Details    string `json:"details"`
	Date       string `json:"date"`
}
