package base

import "time"

type Update struct {
	UpdateBy   string    `json:"updateBy" orm:"null;column(UPDATE_BY);size(32);description(更新人)"` //更新人
	UpdateTime time.Time `json:"updateTime" orm:"null;column(UPDATE_TIME);description(更新时间)"`     //更新时间
}
