package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type App struct {
	Server  *http.Server
	Engine  *gin.Engine
	RunMode string
}

func NewApp() *App {
	//gin.SetMode(gin.ReleaseMode) 初始化配置

	app := &App{Engine: gin.New(), Server: &http.Server{}, RunMode: gin.Mode()}
	return app
}

func (app *App) Run(middleware ...gin.HandlerFunc) {
	app.Engine.Use(middleware...)
	app.Server.Handler = app.Engine
	log.Printf("Start to listening the incoming requests on http address: %s", "8080")
	app.Server.Addr = ":8080"
	log.Printf(app.Server.ListenAndServe().Error())
}

func (app *App) Group(relativePath string, handlers ...gin.HandlerFunc) *App {
	app.Engine.Group(relativePath, handlers...)
	return app
}

func (app *App) Router(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return app.Engine.Handle(httpMethod, relativePath, handlers...)
}

//func (app *App) GetRouter(relativePath string, handlers ...gin.HandlerFunc) {
//	app.Engine.GET(relativePath, handlers...)
//}
