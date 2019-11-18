package util

import (
	"encoding/json"
)

/* 对象序列化 */
func ToJSON(entity interface{}) []byte {
	b, err := json.Marshal(entity)
	if err != nil {
		return nil
	}
	return b
}

/* 反序列化 */
func PhraseJSON(data []byte, entity interface{}) {
	// Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构
	// 第二个参数必须是指针，否则无法接收解析的数据
	_ = json.Unmarshal(data, entity)
}
