package oath

type SignInTypes string

const (
	GITHUB  SignInTypes = "github"
	GOOGLE  SignInTypes = "google"
	LOCAL   SignInTypes = "local"
	UNKNOWN SignInTypes = "unknown"
)
