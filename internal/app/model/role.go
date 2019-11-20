package model

import (
	"github.com/coffee377/voc-admin/internal/app/model/base"
)

/* 角色 */
type Role struct {
	base.Primary        //主键
	Code         string `json:"code" orm:"column(CODE);size(32);unique;description(角色编码)"` //角色编码
	Name         string `json:"name" orm:"column(NAME);size(32);description(角色名称)"`        //角色名称
	base.Create         //创建
	base.Update         //更新
}

/* 自定义表名 */
func (r *Role) TableName() string {
	return "sys_role"
}
