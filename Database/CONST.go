package Database

const (
	DB                        = "EMDB"
	User                      = "user"
	EMAILSUBJECT              = "[EM] Password Reset Request"
	EMAILTEXT                 = "Hi Sir/Madam,\nWe have received a request to change your password.\nPlease verify using the following code : "
	DATABASE_ERROR  ErrorCode = -1
	NO_ERROR        ErrorCode = 0
	DUPLICATE_EMAIL ErrorCode = 1
	UNKNOWN_ERROR   ErrorCode = 2
)

type ErrorCode int
