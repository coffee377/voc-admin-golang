package handler

import (
	"github.com/coffee377/voc-admin/result"
	"github.com/gin-gonic/gin"
)

// 业务逻辑函数
type ServiceFunc func() (interface{}, error)

func RestResponse(c *gin.Context, data interface{}, err error) {
	status, obj := result.Of(data, err).Response()
	c.JSON(status, obj)
}
