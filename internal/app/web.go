package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/coffee377/voc-admin/internal/app/config"
	"github.com/coffee377/voc-admin/internal/app/routes/api"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"golang.org/x/crypto/acme/autocert"
	"io/ioutil"
	"net/http"
	"time"
)

// 初始化 WEB 引擎
func initWeb() *gin.Engine {
	gin.SetMode(config.Global.RunMode)

	app := gin.New()

	//app.NoMethod(middleware.NoMethodHandler())
	//app.NoRoute(middleware.NoRouteHandler())
	//
	//apiPrefixes := []string{"/api/"}
	//
	//// 跟踪ID
	//app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))
	//
	//// 访问日志
	//app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(apiPrefixes...)))
	//
	//// 崩溃恢复
	//app.Use(middleware.RecoveryMiddleware())
	//
	//// 跨域请求
	//if cfg.CORS.Enable {
	//	app.Use(middleware.CORSMiddleware())
	//}

	// 注册/api路由
	err := api.RegisterRouter(app)
	handleError(err)

	//// swagger文档
	//if dir := cfg.Swagger; dir != "" {
	//	app.Static("/swagger", dir)
	//}
	//
	//// 静态站点
	//if dir := cfg.WWW; dir != "" {
	//	app.Use(middleware.WWWMiddleware(dir))
	//}

	return app
}

//var server = &http.Server{}
var (
	//err        error
	//l          net.Listener
	handler    = initWeb()
	endRunning = make(chan bool, 1)
)

// 初始化http服务
func InitHTTPServer(ctx context.Context) func() {
	//beego.Run()

	//server.Handler = initWeb()
	//server.ReadTimeout = time.Duration(60) * time.Second
	//server.WriteTimeout = time.Duration(60) * time.Second
	//server.IdleTimeout = 15 * time.Second
	//server.ErrorLog =

	// HTTP
	go startHttpServer(config.Global.Http)

	// HTTPS
	go startHttpsServer(config.Global.Https)

	//time.Sleep(100 * time.Microsecond)
	<-endRunning
	return func() {
		//	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		//	defer cancel()
		//
		//	srv.SetKeepAlivesEnabled(false)
		//	if err := srv.Shutdown(ctx); err != nil {
		//		logger.Errorf(ctx, err.Error())
		//	}
	}

}

func startHttpServer(cfg config.Http) {
	if cfg.Enable {
		server := http.Server{
			Addr:              fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:           handler,
			TLSConfig:         nil,
			ReadTimeout:       0,
			ReadHeaderTimeout: 0,
			WriteTimeout:      0,
			IdleTimeout:       0,
			MaxHeaderBytes:    0,
			TLSNextProto:      nil,
			ConnState:         nil,
			ErrorLog:          nil,
			BaseContext:       nil,
			ConnContext:       nil,
		}
		log.Infof("http server Running on http://%s", server.Addr)

		if err := server.ListenAndServe(); err != nil {
			log.Error("ListenAndServe", err)
			time.Sleep(100 * time.Microsecond)
			endRunning <- true
		}
	}
}

func startHttpsServer(cfg config.Https) {
	if cfg.Enable {
		time.Sleep(1000 * time.Microsecond)
		server := http.Server{
			Handler: handler,
		}
		if cfg.Port != 0 {
			server.Addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		}

		log.Infof("https server Running on https://%s", server.Addr)

		if cfg.AutoTLS {
			m := autocert.Manager{
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(cfg.Domains...),
				Cache:      autocert.DirCache(cfg.TLSCacheDir),
			}
			server.TLSConfig = &tls.Config{GetCertificate: m.GetCertificate}
			cfg.CertFile, cfg.KeyFile = "", ""
		} else if cfg.EnableMutualHTTPS {
			pool := x509.NewCertPool()
			data, err := ioutil.ReadFile(cfg.TrustFile)
			if err != nil {
				log.Warnf("MutualHTTPS should provide TrustCaFile")
				return
			}
			pool.AppendCertsFromPEM(data)
			server.TLSConfig = &tls.Config{
				ClientCAs:  pool,
				ClientAuth: tls.RequireAndVerifyClientCert,
			}
		}

		if err := server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil {
			log.Error("ListenAndServeTLS: ", err)
			time.Sleep(100 * time.Microsecond)
			endRunning <- true
		}
	}
}
