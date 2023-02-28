package errno

var (
	// 通用错误

	// OK Err OK
	OK = &Errno{Code: 0, Message: "OK"}
	// InternalServerError 服务错误
	InternalServerError = &Errno{Code: 10001, Message: "Internal server errno"}
	// ErrBind bind错误
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	// ErrValidation 验证错误
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	// ErrDatabase 数据库错误
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}
	// ErrToken token错误
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	// ErrTokenInvalid token验证错误
	ErrTokenInvalid = &Errno{Code: 20004, Message: "The token was invalid."}
	// ErrNotFound 未找到
	ErrNotFound = &Errno{Code: 20005, Message: "Not Found"}
)
