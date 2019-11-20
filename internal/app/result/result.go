package result

import (
	errors2 "github.com/coffee377/voc-admin/internal/app/errors"
	"github.com/coffee377/voc-admin/pkg/util"
)

/* API 响应数据格式 */
type Result struct {
	Code    string      `json:"code,omitempty"`    // 响应编码
	Message string      `json:"message,omitempty"` // 响应信息
	Status  int         `json:"-"`                 // 响应转态码
	Data    interface{} `json:"data,omitempty"`    // 响应数据
}

func (result *Result) Bytes() []byte {
	return util.ToJSON(result)
}

func (result *Result) Response() (status int, data interface{}) {
	return result.Status, result
}

func Of(data interface{}, err error) Result {
	e := errors2.DecodeError(err)
	return Result{Code: e.Code, Message: e.Message, Status: e.StatusCode, Data: data}
}
