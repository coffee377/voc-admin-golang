package middleware

import (
	"github.com/coffee377/voc-admin/handler"
	"github.com/coffee377/voc-admin/result"
	"github.com/coffee377/voc-admin/util"
	"github.com/gin-gonic/gin"
	"regexp"
)

// 用户认证（验证用户身份）
func Authentication() {

}

// 用户访问控制（授权）
func Authorization() {

}

func AuthMiddleware(c *gin.Context) {
	// 非登录请求接口
	if match, _ := regexp.MatchString("/portal/login", c.Request.URL.Path); !match {
		authToken := c.Request.Header.Get("x-auth-token")
		if valid, _ := util.ValidToken(authToken); !valid {
			handler.RestResponse(c, nil, result.Unauthorized)
			c.Abort()
		}
	}
	c.Next()
}
