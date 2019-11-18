package base

import "time"

type Create struct {
	CreateBy   string    `json:"createBy" orm:"null;column(CREATE_BY);size(32);description(创建人)"` //创建人
	CreateTime time.Time `json:"createTime" orm:"null;column(CREATE_TIME);description(创建时间)"`     //创建时间
}
