package CONST

const (
	DB                        = "EMDB"
	User                      = "user"
	MAILSMTP                  = "smtp.gmail.com"
	Project                   = "project"
	EMAILSUBJECT              = "[EM] Password Reset Request"
	EMAILTEXT                 = "Hi Sir/Madam,\nWe have received a request to change your password.\nPlease verify using the following code : "
	DATABASE_ERROR  ErrorCode = -1
	NO_ERROR        ErrorCode = 0
	DUPLICATE_EMAIL ErrorCode = 1
	UNKNOWN_ERROR   ErrorCode = 2
	TIMEFORMAT                = "2012-11-01T22:08:41+00:00"
	Admin           Role      = 0
	Editor          Role      = 1
	Viewer          Role      = 2
	PROCESSING                = "processing"
	UNRESOLVED                = "unresolved"
	Development     Env       = 0
	Staging         Env       = 1
	Production      Env       = 2
)

type ErrorCode int
type Role int
type Env int
