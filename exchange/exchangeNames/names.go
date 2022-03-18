// Package exchangeNames constant exchange names
package exchangeNames

type ExchangeNames string

const (
	DELETE_AUTH ExchangeNames = "exchange:delete-auth"
	UPDATE_AUTH ExchangeNames = "exchange:update-auth"
	CREATE_AUTH ExchangeNames = "exchange:create-auth"
	ERROR_AUTH  ExchangeNames = "exchange:error-auth"
	INFO_AUTH   ExchangeNames = "exchange:info-auth"

	DELETE_EMPLOYEE ExchangeNames = "exchange:delete-employee"
	UPDATE_EMPLOYEE ExchangeNames = "exchange:update-employee"
	CREATE_EMPLOYEE ExchangeNames = "exchange:create-employee"
	ERROR_EMPLOYEE  ExchangeNames = "exchange:error-employee"
	INFO_EMPLOYEE   ExchangeNames = "exchange:info-employee"

	DELETE_TASK ExchangeNames = "exchange:delete-task"
	UPDATE_TASK ExchangeNames = "exchange:update-task"
	CREATE_TASK ExchangeNames = "exchange:create-task"
	ERROR_TASK  ExchangeNames = "exchange:error-task"
	INFO_TASK   ExchangeNames = "exchange:info-task"

	DELETE_DIVISION ExchangeNames = "exchange:delete-division"
	UPDATE_DIVISION ExchangeNames = "exchange:update-division"
	CREATE_DIVISION ExchangeNames = "exchange:create-division"
	ERROR_DIVISION  ExchangeNames = "exchange:error-division"
	INFO_DIVISION   ExchangeNames = "exchange:info-division"

	SERVICE_ERRORS ExchangeNames = "exchange:service-errors"
	LOG            ExchangeNames = "exchange:log"
)
