package api

import (
	"github.com/coffee377/voc-admin/router/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	// BeeApp is an application instance
	server *App
)

func init() {
	server = NewApp()
}

func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return server.Engine.Group(relativePath, handlers...)
}

func Router(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return server.Engine.Handle(httpMethod, relativePath, handlers...)
}

func Run(params ...string) {
	if server.RunMode != gin.ReleaseMode {
		log.Printf("开始运行")
	}
	server.Run(middleware.NoRoute, middleware.AuthMiddleware)
}
