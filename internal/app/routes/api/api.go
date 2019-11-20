package api

import (
	"github.com/coffee377/voc-admin/internal/app/routes/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine) error {
	g := app.Group("/api")

	// 用户身份授权
	//g.Use(middleware.UserAuthMiddleware(a,
	//	middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	//))

	// casbin权限校验中间件
	//g.Use(middleware.CasbinMiddleware(e,
	//	middleware.AllowPathPrefixSkipper("/api/v1/pub"),
	//))

	// 请求频率限制中间件
	//g.Use(middleware.RateLimiterMiddleware())
	v1 := g.Group("/v1")
	{
		portal := v1.Group("/portal")
		// 注册/api/v1/portal/login
		login := portal.Group("login")
		{
			//login.GET("captcha")
			login.GET("", handler.Login)
			//login.POST("exit")
		}
	}

	return nil
}
