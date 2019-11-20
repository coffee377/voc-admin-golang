package app

import (
	"context"
)

// 配置
type options struct {
	ConfigFile string
	ModelFile  string
	WWWDir     string
	SwaggerDir string
	MenuFile   string
	Version    string
}

// 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// 设定 casbin 模型配置文件
func SetModelFile(s string) Option {
	return func(o *options) {
		o.ModelFile = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	//err := config.LoadGlobal(o.ConfigFile)
	//handleError(err)

	//cfg := config.Global()
	//logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程：%d", "cfg.RunMode", o.Version, os.Getpid())

	//if v := o.ModelFile; v != "" {
	//	cfg.Casbin.Model = v
	//}
	//if v := o.WWWDir; v != "" {
	//	cfg.WWW = v
	//}
	//if v := o.SwaggerDir; v != "" {
	//	cfg.Swagger = v
	//}
	//if v := o.MenuFile; v != "" {
	//	cfg.Menu.Data = v
	//}

	loggerCall, err := InitLogger()
	handleError(err)

	_ = InitDB()
	//err = InitMonitor()
	//if err != nil {
	//	logger.Errorf(ctx, err.Error())
	//}

	//// 初始化图形验证码
	//InitCaptcha()
	//
	//// 创建依赖注入容器
	//container, containerCall := BuildContainer()

	//// 初始化数据
	//err = InitData(ctx, container)
	//handleError(err)

	// 初始化HTTP服务
	httpCall := InitHTTPServer(ctx)
	return func() {
		if httpCall != nil {
			httpCall()
		}
		//if containerCall != nil {
		//	containerCall()
		//}
		if loggerCall != nil {
			loggerCall()
		}
	}
}

// APP 内部错误处理
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
