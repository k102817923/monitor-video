# 赛事监控服务

基于七牛云 SKD 开发的赛事监控服务，主要功能：音视频监控、录屏监控、视频转码等。

monitor-video
├── common
│ └── code.go 错误码
│ └── qiniu.go 七牛云 SDK
│ └── response.go 通用响应
│ └── util.go 工具类
├── config 配置文件
│ └── config.default.ini
├── controller
│ └── monitor.go
├── logging 日志
│ └── file.go
│ └── log.go
├── middleware
│ └── cors.go 处理跨域
├── routes 路由
│ └── router.go
├── runtime 应用运行时数据
│ └── logs
├── service 业务逻辑
│ └── qiniu.go
├── setting 初始化配置
│ └── setting.go
├── tmp 临时文件
├── .air.toml Air 配置文件
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── main.go

# Running

1. `go mod download` 下载 go.mod 文件中指明的所有依赖
2. `go mod tidy` 整理现有的依赖
3. `go run main.go` 运行项目

# Hot Reload

1. `go get -u github.com/cosmtrek/air` 下载 Air 自动重新加载工具
2. `air` 运行项目
