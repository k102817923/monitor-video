package main

import (
	"context"
	"fmt"
	"log"
	"monitor-video/logging"
	"monitor-video/routes"
	"monitor-video/setting"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	router := routes.InitRouter()

	// Go 的标准 HTTP 服务器结构。它封装了一个 HTTP 服务，配置了服务器的一些基本参数
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 使用 goroutine 启动 HTTP 服务器
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logging.Info("Listen: %s\n", err)
		}
	}()

	// 使用 buffered channel 来避免死锁
	quit := make(chan os.Signal, 1) // 设置容量为 1 的缓冲通道
	signal.Notify(quit, os.Interrupt)

	// 等待信号
	<-quit

	logging.Info("Shutdown Server...")

	// 创建带有超时的 context，优雅关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	logging.Info("Server exiting")
}
