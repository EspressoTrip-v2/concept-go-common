package logging

type LogMsg struct {
	Service    string `json:"service"`
	LogContext string `json:"logContext"`
	Origin     string `json:"origin"`
	Message    string `json:"message"`
	Details    string `json:"details"`
	Date       string `json:"date"`
}
