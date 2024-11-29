# 使用官方的 Golang 基础镜像
FROM golang:1.23 AS build

# 设置工作目录
WORKDIR /app
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制项目源代码到容器中
COPY . .

# 构建应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o monitor-video .

# 使用轻量的 alpine 安装 git 后的镜像上传到腾讯云镜像仓库作为基础镜像
FROM tatfook-registry.tencentcloudcr.com/devops/alpine-git:latest

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制二进制文件到最终镜像
COPY --from=build /app/monitor-video .
COPY --from=build /app/config/config.default.ini ./config/config.default.ini

# 设置环境变量
ENV PORT=80
# 添加配置文件路径环境变量
ENV CONFIG_PATH="/root/config/config.default.ini"

# 暴露端口
EXPOSE $PORT

# 启动应用程序
ENTRYPOINT ["/bin/sh", "-c", "./monitor-video -config ${CONFIG_PATH}"]
