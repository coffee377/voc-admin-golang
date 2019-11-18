package router

import (
	"github.com/coffee377/voc-admin/api"
	"github.com/coffee377/voc-admin/handler/sd"
)

func init() {
	//if api.Server.RunMode != gin.ReleaseMode {
	//	log.Printf("路由初始化开始")
	//}
	//api.Server.Engine.Group("").g
	//api.
	g := api.Group("/sd")
	g.GET("/health/:id", sd.HealthCheck)
	g.GET("/health", sd.HealthCheck)
	//api.Router("GET", "sd/health/:id", sd.HealthCheck)
	//api.Router("GET", "sd/health", sd.HealthCheck)

	auth()

	//if api.Server.RunMode != gin.ReleaseMode {
	//	log.Printf("路由初始化结束")
	//}
}

// 认证接口
func auth() {
	//portal := api.Group("/portal")
	// POST /portal/login
	//portal.POST("/login")
	//POST /portal/logout
	//portal.POST("/logout")
}
