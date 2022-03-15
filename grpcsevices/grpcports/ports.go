package grpcports

type GrpcServicePortDns string

const (
	AUTH_SERVICE_DNS          GrpcServicePortDns = "auth-service-srv:8080"
	ANALYTIC_SERVICE_DNS      GrpcServicePortDns = "analytic-service-srv:8080"
	EMPLOYEE_SERVICE_DNS      GrpcServicePortDns = "employee-service-srv:8080"
	TASK_SERVICE_DNS          GrpcServicePortDns = "task-service-srv:8080"
	DIVISION_SERVICE_DNS      GrpcServicePortDns = "division-service-srv:8080"
	EMPLOYEE_DASH_SERVICE_DNS GrpcServicePortDns = "employee-dash-service-srv:8080"
)
