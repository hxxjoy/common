// pkg/errors/errors.go
package errors

import (
    "fmt"
    "net/http"
)

// ErrorCode 错误码
type ErrorCode int

// Error 自定义错误
type Error struct {
    Code    ErrorCode
    Message string
    Err     error
}

func (e *Error) Error() string {
    return fmt.Sprintf("code: %d, message: %s, error: %v", e.Code, e.Message, e.Err)
}

// New 创建错误
func New(code ErrorCode, message string) *Error {
    return &Error{
        Code:    code,
        Message: message,
    }
}

// Wrap 包装错误
func Wrap(err error, code ErrorCode, message string) *Error {
    return &Error{
        Code:    code,
        Message: message,
        Err:     err,
    }
}