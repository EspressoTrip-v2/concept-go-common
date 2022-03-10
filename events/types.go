package events

type ExchangeNames string
type ExchangeType string
type QueueInfo string
type MicroserviceNames string

const (
	AUTH           ExchangeNames = "exchange:auth"
	EMPLOYEE       ExchangeNames = "exchange:employee"
	TASK           ExchangeNames = "exchange:task"
	DIVISION       ExchangeNames = "exchange:division"
	SERVICE_ERRORS ExchangeNames = "exchange:service-errors"
	LOG            ExchangeNames = "exchange:log"

	FAN_OUT ExchangeType = "fanout"
	DIRECT  ExchangeType = "direct"
	TOPIC   ExchangeType = "topic"

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

	ANALYTIC_API          MicroserviceNames = "ANALYTIC_API"
	ANALYTIC_SERVICE      MicroserviceNames = "ANALYTIC_SERVICE"
	AUTH_SERVICE          MicroserviceNames = "AUTH_SERVICE"
	AUTH_API              MicroserviceNames = "AUTH_API"
	EMPLOYEE_API          MicroserviceNames = "EMPLOYEE_API"
	EMPLOYEE_SERVICE      MicroserviceNames = "EMPLOYEE_SERVICE"
	TASK_API              MicroserviceNames = "TASK_API"
	TASK_SERVICE          MicroserviceNames = "TASK_SERVICE"
	DIVISION_API          MicroserviceNames = "DIVISION_API"
	DIVISION_SERVICE      MicroserviceNames = "DIVISION_SERVICE"
	EMPLOYEE_DASH_API     MicroserviceNames = "EMPLOYEE_DASH_API"
	EMPLOYEE_DASH_SERVICE MicroserviceNames = "EMPLOYEE_DASH_SERVICE"
	LOG_SERVICE           MicroserviceNames = "LOG_SERVICE"
)
