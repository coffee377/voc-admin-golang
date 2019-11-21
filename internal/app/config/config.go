package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"strings"
)
import "github.com/spf13/viper"

type Map = map[string]interface{}

type Http struct {
	Enable bool   `json:"enable"` //是否启用
	Host   string `json:"host"`   // Host
	Port   int    `json:"port"`   // 端口
}

type Https struct {
	Enable            bool     `json:"enable"`              // 是否启用
	AutoTLS           bool     `json:"auto_tls"`            //
	Domains           []string `json:"domains"`             //
	TLSCacheDir       string   `json:"tls_cache_dir"`       //
	EnableMutualHTTPS bool     `json:"enable_mutual_https"` //
	Host              string   `json:"host"`                // Host
	Port              int      `json:"port"`                // 端口
	CertFile          string   `json:"cert_file"`           //
	KeyFile           string   `json:"key_file"`            //
	TrustFile         string   `json:"trust_file"`          //
}

type DataBase struct {
	Adapter  string                 `json:"adapter"`
	Dir      string                 `json:"dir"` // 数据库文件存放目录
	Prefix   string                 `json:"prefix"`
	Username string                 `json:"username"`
	Password string                 `json:"password"`
	Host     string                 `json:"host"`
	Port     string                 `json:"port"`
	Protocol string                 `json:"protocol"`
	Name     string                 `json:"name"`
	Params   map[string]interface{} `json:"params"`
}

// 日志配置
//@see https://github.com/lexkong/log
type LogConfig struct {
	Writers        string `json:"writers"`          //输出位置，有2个可选项：file,stdout
	LoggerLevel    string `json:"logger_level"`     // 日志级别，DEBUG, INFO, WARN, ERROR, FATAL
	LoggerFile     string `json:"logger_file"`      // 日志文件
	LogFormatText  bool   `json:"log_format_text"`  // 日志的输出格式，json或者plaintext，true会输出成json格式，false会输出成非json格式
	RollingPolicy  string `json:"rollingPolicy"`    // rotate依据，可选的有：daily, size。如果选daily则根据天进行转存，如果是size则根据大小进行转存
	LogRotateDate  int    `json:"log_rotate_date"`  // rotate转存时间，配合rollingPolicy: daily使用
	LogRotateSize  int    `json:"log_rotate_size"`  // rotate转存大小，配合rollingPolicy: size使用
	LogBackupCount int    `json:"log_backup_count"` // 当日志文件达到转存标准时，log系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数。
}

type WebConfig struct {
}

// Listen holds for http and https related config
type Listen struct {
	Graceful      bool  // Graceful means use graceful module to start the server
	ServerTimeOut int64 `json:"server_time_out"` // 服务超时时间
	ListenTCP4    bool
	//EnableHTTP        bool `json:"enable_http"`
	HTTPAddr    string
	HTTPPort    int
	AutoTLS     bool
	Domains     []string
	TLSCacheDir string
	//EnableHTTPS       bool
	EnableMutualHTTPS bool
	//HTTPSAddr         string
	//HTTPSPort         int
	//HTTPSCertFile     string
	//HTTPSKeyFile      string
	TrustCaFile string
	//EnableAdmin       bool
	//AdminAddr         string
	//AdminPort         int
	//EnableFcgi        bool
	//EnableStdIo       bool // EnableStdIo works with EnableFcgi Use FCGI via standard I/O
}

type Configuration struct {
	ConfigFile string `json:"-"`        // 配置文件名称
	AppName    string `json:"app_name"` //
	ServerName string `json:"name"`     // 服务名称
	RunMode    string `json:"runmode"`  //
	//TimeZone     string    `json:"time_zone"`     // 时区
	RecoverPanic bool      `json:"recover_panic"` //
	Http         Http      `json:"http"`
	Https        Https     `json:"https"`
	Log          LogConfig `json:"log"`
	Web          WebConfig `json:"web"`
	DB           DataBase  `json:"database"`
	//Listen       Listen    `json:"listen"`
}

var (
	// 全局配置
	Global = &Configuration{}
)

// 默认配置
func (c *Configuration) defaultConfig() {
	//viper.SetDefault("time_zone", "Asia/Shanghai")
	viper.SetDefault("runmode", gin.ReleaseMode)
	// HTTP
	viper.SetDefault("http", Map{
		"enable": true,
		"host":   "",
		"port":   80,
	})

	// HTTPS
	viper.SetDefault("https", Map{
		"enable":              false,
		"auto_tls":            false,
		"domains":             "",
		"tls_cache_dir":       "",
		"enable_mutual_https": "",
		"host":                "",
		"port":                443,
		"cert_file":           "conf/cert.pem",
		"key_file":            "conf/key.pem",
		"trust_file":          "",
	})

	// LOG
	viper.SetDefault("log", Map{
		"writers":          "stdout",
		"logger_level":     "DEBUG",
		"logger_file":      "logs/log.log",
		"log_format_text":  true,
		"rollingPolicy":    "size",
		"log_rotate_date":  1,
		"log_rotate_size":  1,
		"log_backup_count": 7,
	})

	// DataBase
	viper.SetDefault("database", Map{
		"adapter":  "sqlite3",
		"dir":      "./data",
		"prefix":   "",
		"username": "",
		"Password": "",
		"host":     "127.0.0.1",
		"port":     "3306",
		"protocol": "tcp",
		"name":     "test",
		"params": Map{
			"charset":   "utf8",
			"parseTime": true,
			"loc":       "Local",
		},
	})
}

// 初始化配置
func (c *Configuration) initConfig() {
	if c.ConfigFile != "" {
		viper.SetConfigFile(c.ConfigFile) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yml") // 设置配置文件格式为YAML
	viper.AutomaticEnv()       // 读取匹配的环境变量
	viper.SetEnvPrefix("API")  // 读取环境变量的前缀为 API
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		//return err
	}
	//return nil
}

// 监控配置文件变化并热加载程序
func (c *Configuration) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		c.unmarshal()
		//log.Printf("Configuration file changed: %s", e.Name)
	})
}

func (c *Configuration) unmarshal() {

	err := viper.Unmarshal(&Global, func(m *mapstructure.DecoderConfig) {
		m.TagName = "json"
	})
	if err != nil {
		//log.Fatalf("unable to decode into struct, %v", err)
	}
}

func Init(configFile string) {
	Global.ConfigFile = configFile

	// 初始化配置文件
	Global.initConfig()

	// 默认配置
	Global.defaultConfig()

	// 构建正式配置
	Global.unmarshal()

	// 监控配置文件变化并热加载程序
	Global.watchConfig()

}
