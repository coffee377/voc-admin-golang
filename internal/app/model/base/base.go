package base

import (
	"fmt"
	"time"
)

type UnixTime time.Time

// MarshalJSON implements json.Marshaler.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).UnixNano()/1e6)
	return []byte(stamp), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
//func (t *UnixTime) UnmarshalJSON(data []byte) error {
//	// Ignore null, like in the main JSON package.
//	if string(data) == "null" {
//		return nil
//	}
//	// Fractional seconds are handled implicitly by Parse.
//	var err error
//	*t, err = Parse(`"`+RFC3339+`"`, string(data))
//	return err
//}

/* 主键 */
type Primary struct {
	Id string `json:"id" orm:"column(ID);size(32);pk;description(主键ID)"` //主键ID
}

type Created struct {
	CreatedBy string    `json:"createdBy" orm:"null;column(CREATE_BY);size(32);description(创建人)"` //创建人
	CreatedAt time.Time `json:"createdAt" orm:"null;column(CREATE_TIME);description(创建时间)"`       //创建时间
}

type Updated struct {
	UpdatedBy string    `json:"updateBy" orm:"null;column(UPDATE_BY);size(32);description(更新人)"` //更新人
	UpdatedAt time.Time `json:"updateTime" orm:"null;column(UPDATE_TIME);description(更新时间)"`     //更新时间
}

type Deleted struct {
	DeletedBy string
	DeletedAt time.Time
}

type Base struct {
}
