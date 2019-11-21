package model

import (
	"github.com/coffee377/voc-admin/internal/app/model/base"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

/* 用户 */
type User struct {
	base.Primary
	//Id         string    `json:"id" orm:"column(ID);size(32);pk;description(用户ID)"`                 //用户ID
	Username string `json:"username" orm:"column(USER_NAME);size(64);unique;description(用户名)"` //用户名
	Password string `json:"password" orm:"column(PASSWORD);size(128);description(密码)"`         //密码
	RealName string `json:"realName" orm:"null;column(REAL_NAME);size(8);description(真实姓名)"`   //真实姓名
	Email    string `json:"email" orm:"null;column(EMAIL);size(128);description(邮箱)"`          //邮箱
	Phone    string `json:"phone" orm:"null;column(PHONE);size(11);description(手机)"`           //手机
	base.Created
	//CreateBy   string    `json:"create_by" orm:"column(CREATE_BY);size(32);description(创建人)"`       //创建人
	//CreateTime time.Time `json:"create_time" orm:"column(CREATE_TIME);description(创建时间)"`           //创建时间
	base.Updated
	//UpdateBy   string    `json:"update_by" orm:"column(UPDATE_BY);size(32);description(更新人)"`       //更新人
	//UpdateTime time.Time `json:"update_time" orm:"column(UPDATE_TIME);description(更新时间)"`           //更新时间
	Status int `json:"status" orm:"null;column(STATUS);default(1);description(状态)"` //状态
}

/* 自定义表名 */
func (u *User) TableName() string {
	return "sys_user"
}
