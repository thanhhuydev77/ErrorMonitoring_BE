package Database

const (
	DB   = "EMDB"
	User = "user"

	DATABASE_ERROR  ErrorCode = -1
	NO_ERROR        ErrorCode = 0
	DUPLICATE_EMAIL ErrorCode = 1
	UNKNOWN_ERROR   ErrorCode = 2
)

type ErrorCode int
