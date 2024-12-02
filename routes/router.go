package routes

import (
	"monitor-video/controller"
	"monitor-video/middleware"
	"monitor-video/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	router.Use(middleware.CORS())

	router.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "Hello, world!",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"msg": "404",
		})
	})

	api := router.Group("/v0")
	api.Use(middleware.Auth())
	{
		// 存储配置
		api.POST("initStorage", controller.InitStorage)
		// 生成 Room Token
		api.POST("getRoomToken", controller.GetRoomToken)
		// 关闭房间
		api.POST("closeRoom", controller.CloseRoom)
		// 录制任务
		api.POST("recordTask", controller.RecordTask)
		// 生成回放
		api.POST("saveVideo", controller.SaveVideo)
		// 音视频转码
		api.POST("videoToMp4", controller.VideoToMp4)
		// 查询音视频转码进度
		api.POST("checkMp4Status", controller.CheckMp4Status)
	}

	return router
}
