package result

/* 数据字节码接口 */
type BytesResult interface {

	// 响应数据字节码
	Bytes() []byte
}

// 支持 gin context.Result 接口
type IResponse interface {
	Response() (status int, data interface{})
}
