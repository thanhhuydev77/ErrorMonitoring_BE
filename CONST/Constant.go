package CONST

const (
	DB                        = "EMDB"
	User                      = "user"
	MAILSMTP                  = "smtp.gmail.com"
	Project                   = "project"
	EMAILSUBJECT              = "[EM] Password Reset Request"
	DATABASE_ERROR  ErrorCode = -1
	NO_ERROR        ErrorCode = 0
	DUPLICATE_EMAIL ErrorCode = 1
	UNKNOWN_ERROR   ErrorCode = 2
	TIMEFORMAT                = "2012-11-01T22:08:41+00:00"
	PROCESSING                = "processing"
	UNRESOLVED                = "unresolved"
)

type ErrorCode int
type Role int
type Env int
type Status int
