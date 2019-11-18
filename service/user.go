package service

import (
	"github.com/coffee377/voc-admin/model/account"
	"github.com/coffee377/voc-admin/result"
)

func Login(credentials account.Credentials) (interface{}, error) {
	return nil, result.BizError{Code: "100001", Message: "用户不存在"}
}

func Logout() (interface{}, error) {
	return nil, nil
}
