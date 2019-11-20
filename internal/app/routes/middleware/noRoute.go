package middleware

import (
	"github.com/coffee377/voc-admin/internal/app/errors"
	"github.com/coffee377/voc-admin/internal/app/result"
	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(404, result.Of(nil, errors.NotFound))
}
