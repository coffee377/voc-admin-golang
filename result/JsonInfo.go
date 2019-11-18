package result

/* JSON信息接口 */
type JsonInfo interface {
	// HTTP 响应状态码
	HttpStatus() int

	// 响应数据字节码
	Bytes() []byte

	//WrapperBytes() []byte
}

// 支持 gin context.Result 接口
type JSONResponse interface {
	Response() (status int, data interface{})
}
