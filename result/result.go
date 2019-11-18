package result

import (
	"github.com/coffee377/voc-admin/util"
)

/* API 响应数据格式 */
type Result struct {
	Code    string      `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (json *Result) HttpStatus() int {
	return util.ErrorCode2Status(json.Code)
}

func (json *Result) Bytes() []byte {
	return util.ToJSON(json)
}

func (json *Result) Response() (status int, data interface{}) {
	return json.HttpStatus(), json
}

func Of(data interface{}, err error) *Result {
	e := DecodeError(err)
	return &Result{Code: e.Code, Message: e.Message, Data: data}
}
