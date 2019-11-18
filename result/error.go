package result

import (
	"fmt"
)

/* 业务异常定义 */
type BizError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (err BizError) Error() string {
	return fmt.Sprintf("BizError - code: %s, message: %s", err.Code, err.Message)
}

func DecodeError(err error) BizError {
	if err == nil {
		return OK
	}

	// 错误类型判断
	switch typed := err.(type) {

	case BizError:
		return NewBizError(typed.Code, typed.Message)
	default:
		return NewBizError(InternalServerError.Code, err.Error())
	}
}

func NewBizError(code string, message string) BizError {
	return BizError{code, message}
}
