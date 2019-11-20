
## 项目结构概览

```
.
├── cmd
│   └── server：主服务（程序入口）
├── configs：配置文件目录(包含运行配置参数及casbin模型配置)
├── docs：文档目录
│   └── swagger：swagger静态文件目录
├── internal：内部应用
│   └── app：主应用目录
│       ├── bll：业务逻辑层接口
│       │   └── impl：业务逻辑层的接口实现
│       ├── config：配置参数（与配置文件一一映射）
│       ├── context：统一上下文
│       ├── errors：统一的错误处理
│       ├── ginplus：gin的扩展函数库
│       ├── middleware：gin中间件
│       ├── model：存储层接口
│       │   └── impl：存储层接口实现
│       │       └── gorm：基于gorm的存储层实现
│       ├── routers：路由层
│       │   └── api：/api路由模块
│       │       └── ctl：/api路由模块对应的控制器层
│       ├── schema：对象模型
│       └── test：针对接口的单元测试
├── pkg：公共模块
│   ├── auth：认证模块
│   │   └── jwtauth：JWT认证模块实现
│   ├── logger：日志模块
│   └── util：工具库
└── scripts：执行脚本
```