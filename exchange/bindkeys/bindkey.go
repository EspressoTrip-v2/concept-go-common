package bindkeys

type BindKey string

const (
	AUTH_CREATE BindKey = "auth:create"
	AUTH_DELETE BindKey = "auth:delete"
	AUTH_UPDATE BindKey = "auth:update"
	AUTH_ERROR  BindKey = "auth:error"

	EMPLOYEE_CREATE BindKey = "employee:create"
	EMPLOYEE_DELETE BindKey = "employee:delete"
	EMPLOYEE_UPDATE BindKey = "employee:update"

	TASK_CREATE BindKey = "task:create"
	TASK_DELETE BindKey = "task:delete"
	TASK_UPDATE BindKey = "task:update"

	TASK_EMP_CREATE BindKey = "task_emp:create"
	TASK_EMP_DELETE BindKey = "task_emp:delete"
	TASK_EMP_UPDATE BindKey = "task_emp:update"

	INFO BindKey = "all:info"
	LOG  BindKey = "logger:log"
)
