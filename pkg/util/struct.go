package util

import (
	"reflect"
)

// 结构体转 map
func Struct2Map(s interface{}) map[string]interface{} {

	var (
		m     = make(map[string]interface{})
		t     = reflect.TypeOf(s)
		v     = reflect.ValueOf(s)
		field reflect.StructField
		tag   reflect.StructTag
		key   string
	)

	for i := 0; i < t.NumField(); i++ {
		field = t.Field(i)
		tag = field.Tag
		keys := append([]string{}, tag.Get("yml")) // yml 名称
		keys = append(keys, tag.Get("yaml"))       //yaml 名称
		keys = append(keys, tag.Get("json"))       //json 名称
		keys = append(keys, t.Field(i).Name)       //原始名称

		for _, value := range keys {
			if value != "" {
				key = value
				break
			}
		}

		add(m, key, v.Field(i).Interface())
	}
	return m
}

// map 转结构体
func Map2Struct(m map[string]interface{}, s interface{}) {

}

func add(m map[string]interface{}, key string, value interface{}) {
	m[key] = value
}
