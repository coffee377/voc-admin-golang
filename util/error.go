package util

import (
	"net/http"
	"regexp"
)

/* 错误编码转响应状态码 */
func ErrorCode2Status(code string) int {
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
