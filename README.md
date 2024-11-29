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
└── main.go
