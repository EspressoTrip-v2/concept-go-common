package mongodb

type DatabaseNames string

const (
	TASK_DB      DatabaseNames = "task"
	EMPLOYEE_DB  DatabaseNames = "employee"
	DIVISION_DB  DatabaseNames = "division"
	USER_DB      DatabaseNames = "user"
	DASHBOARD_DB DatabaseNames = "dashboard"
)
