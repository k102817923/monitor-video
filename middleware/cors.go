package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 设置跨域请求头
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                // 允许所有域名
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 允许的方法
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")                               // 允许的自定义头

		// 如果是预检请求，直接返回状态码 204
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 继续处理请求
		ctx.Next()
	}
}
