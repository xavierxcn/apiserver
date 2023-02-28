package errno

import (
	"fmt"
)

// Errno 错误码
type Errno struct {
	Code    int
	Message string
}

// Error err
func (err Errno) Error() string {
	return err.Message
}

// Err err
type Err struct {
	Code    int
	Message string
	Err     error
}

// New 新建错误
func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

// Add 添加错误文字
func (err *Err) Add(message string) *Err {
	err.Message += " " + message
	return err
}

// Addf 错误文字
func (err *Err) Addf(format string, args ...interface{}) *Err {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

// Error error接口
func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, errno: %s", err.Code, err.Message, err.Err)
}

// DecodeErr 解析错误的文字信息
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message + " " + typed.Err.Error()
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
