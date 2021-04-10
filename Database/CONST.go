package Database

const (
	DB                        = "EMDB"
	User                      = "user"
	EMAILSUBJECT              = "Xác minh email Error Monitoring"
	EMAILTEXT                 = "Xin chào,\n\nChúng tôi đã nhận được yêu cầu thay đổi mật khẩu của bạn.\n\nVui lòng xác minh bằng cách dán mã bên dưới:"
	DATABASE_ERROR  ErrorCode = -1
	NO_ERROR        ErrorCode = 0
	DUPLICATE_EMAIL ErrorCode = 1
	UNKNOWN_ERROR   ErrorCode = 2
)

type ErrorCode int
