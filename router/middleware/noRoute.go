package middleware

import (
	"github.com/coffee377/voc-admin/result"
	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(404, result.Of(nil, result.NotFound))
}
