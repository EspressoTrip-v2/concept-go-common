package bindkeys

type BindKey string

const (
	CREATE BindKey = "key:create"
	DELETE BindKey = "key:delete"
	UPDATE BindKey = "key:update"
	ERROR  BindKey = "key:error"
	INFO   BindKey = "key:info"
	LOG    BindKey = "key:log"
)
