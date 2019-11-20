package errors

import (
	"net/http"
	"regexp"
)

// 错误编码与状态码映射关系
type Code2StatusMap = map[string]int

// 定义错误
var (
	Mapping         = Code2StatusMap{"0": 200}
	OK              = NewException("0", "OK", nil)
	BadRequest      = NewException("400", "请求发生错误", nil)
	Unauthorized    = NewException("401", "用户未通过认证", nil)
	InvalidToken    = NewException("9999", "令牌失效（认证过期）", nil)
	Forbidden       = NewException("403", "无访问权限", nil)
	NotFound        = NewException("404", "资源不存在", nil)
	MethodNotAllow  = NewException("405", "方法不被允许", nil)
	TooManyRequests = NewException("429", "请求过于频繁", nil)
	InternalServer  = NewException("500", "服务器发生错误", nil)

	//InternalServerError = NewBizError("500", "Internal server error.")
	//ErrInvalidParent           = New400Response("无效的父级节点")
	//ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	//ErrNotAllowDelete          = New400Response("资源不允许删除")
	//ErrInvalidUserName         = New400Response("无效的用户名")
	//ErrInvalidPassword         = New400Response("无效的密码")
	//ErrInvalidUser             = New400Response("无效的用户")
	//ErrUserDisable             = New400Response("用户被禁用，请联系管理员")
	//
)

// 添加映射关系
func add(code string, status int) {
	Mapping[code] = status
}

func init() {
	add("0", 200)
}

/* 错误编码转响应状态码 */
func Code2Status(code string) int {
	// 先从码表中获取
	if s, ok := Mapping[code]; ok && s > 0 {
		return s
	}
	// 未查找到使用匹配规则返回
	if match, _ := regexp.MatchString("401", code); match {
		return http.StatusUnauthorized
	} else if match, _ := regexp.MatchString("403", code); match {
		return http.StatusForbidden
	} else if match, _ := regexp.MatchString("404", code); match {
		return http.StatusNotFound
	} else if match, _ := regexp.MatchString("^4", code); match {
		return http.StatusBadRequest
	} else if match, _ := regexp.MatchString("^5", code); match {
		return http.StatusInternalServerError
	} else {
		return http.StatusOK
	}
}
