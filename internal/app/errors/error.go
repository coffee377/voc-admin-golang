package errors

import (
	"fmt"
)

// Exception 定义响应错误
type Exception struct {
	Code       string `json:"code"`    // 错误码
	Message    string `json:"message"` // 错误消息
	StatusCode int    `json:"status"`  // 响应状态码
	ERR        error  `json:"-"`       // 响应错误
}

func NewException(code string, message string, err error, status ...int) *Exception {
	res := &Exception{
		Code:    code,
		Message: message,
		ERR:     err,
	}
	if len(status) > 0 {
		res.StatusCode = status[0]
	} else {
		res.StatusCode = Code2Status(code)
	}
	return res
}

func (e *Exception) Error() string {
	if e.ERR != nil {
		return e.ERR.Error()
	}
	return fmt.Sprintf("Exception - code: %s, message: %s", e.Code, e.Message)
}

func DecodeError(err error) *Exception {
	if err == nil {
		return OK
	}

	// 错误类型判断
	switch typed := err.(type) {

	case *Exception:
		return NewException(typed.Code, typed.Message, typed.ERR, typed.StatusCode)
	default:
		return InternalServer
	}
}
