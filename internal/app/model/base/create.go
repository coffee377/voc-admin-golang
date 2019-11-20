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
type Create struct {
	CreateBy   string   `json:"createBy" orm:"null;column(CREATE_BY);size(32);description(创建人)"` //创建人
	CreateTime UnixTime `json:"createTime" orm:"null;column(CREATE_TIME);description(创建时间)"`     //创建时间
}
