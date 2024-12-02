package middleware

import (
	"monitor-video/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求头中的 x-api-key
		apiKey := ctx.GetHeader("x-api-key")

		// 校验 x-api-key 是否正确
		if apiKey != setting.LoadStringParam("self", "INTERNAL_API_KEY") {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 继续处理请求
		ctx.Next()
	}
}
