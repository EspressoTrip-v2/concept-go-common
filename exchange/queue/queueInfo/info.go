// Package queueInfo constant queue names
package queueInfo

type QueueInfo string

const (
	AUTH_ERROR  QueueInfo = "error:auth"
	CREATE_USER QueueInfo = "create:auth"
	UPDATE_USER QueueInfo = "update:auth"
	DELETE_USER QueueInfo = "delete:auth"

	CREATE_EMPLOYEE QueueInfo = "create:employee"
	UPDATE_EMPLOYEE QueueInfo = "update:employee"
	DELETE_EMPLOYEE QueueInfo = "delete:employee"

	CREATE_TASK QueueInfo = "create:task"
	DELETE_TASK QueueInfo = "delete:task"
	UPDATE_TASK QueueInfo = "update:task"

	CREATE_DIVISION QueueInfo = "create:division"
	DELETE_DIVISION QueueInfo = "delete:division"
	UPDATE_DIVISION QueueInfo = "update:division"

	SERVICE_ERROR QueueInfo = "service:error"

	LOG_EVENT QueueInfo = "log:event"
)
