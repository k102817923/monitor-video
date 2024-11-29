package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	})
	ctx.Abort()
}

func ResponseError(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  GetMsg(code),
		Data: nil,
	})
	ctx.Abort()
}
