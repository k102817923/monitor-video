package controller

import (
	"monitor-video/common"
	"monitor-video/logging"
	"monitor-video/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitStorage(ctx *gin.Context) {
	var params struct {
		ExpireDay int `json:"expiredDay" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	if err := service.InitStorage(params.ExpireDay); err != nil {
		logging.Error("Failed to init storage: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, "Init storage success")
}

func GetRoomToken(ctx *gin.Context) {
	var params struct {
		RoomName   string `json:"roomName" validate:"required"`
		UserId     int    `json:"userId" validate:"required"`
		Permission string `json:"permission" validate:"required,oneof=user admin"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	token, err := service.GetRoomToken(params.RoomName, strconv.Itoa(params.UserId), params.Permission)
	if err != nil {
		logging.Error("Failed to get room token: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, token)
}

func CloseRoom(ctx *gin.Context) {
	var params struct {
		RoomName string `json:"roomName" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	if err := service.CloseRoom(params.RoomName); err != nil {
		logging.Error("Failed to close room: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, "Room closed success")
}

func RecordTask(ctx *gin.Context) {
	var params struct {
		RoomName string `json:"roomName" validate:"required"`
		UserId   int    `json:"userId" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	info, err := service.RecordTask(params.RoomName, strconv.Itoa(params.UserId))
	if err != nil {
		logging.Error("Failed to record task: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, info)
}

func KickUser(ctx *gin.Context) {
	var params struct {
		RoomName string `json:"roomName" validate:"required"`
		UserId   int    `json:"userId" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	if err := service.KickUser(params.RoomName, strconv.Itoa(params.UserId)); err != nil {
		logging.Error("Failed to kick user: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, "User kicked success")
}

func SaveVideo(ctx *gin.Context) {
	var params struct {
		RoomName string `json:"roomName" validate:"required"`
		UserId   int    `json:"userId" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	if err := service.SaveVideo(params.RoomName, strconv.Itoa(params.UserId)); err != nil {
		logging.Error("Failed to save video: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, "Video saved success")
}

func VideoToMp4(ctx *gin.Context) {
	var params struct {
		VideoName string `json:"videoName" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	persistentId, err := service.VideoToMp4(params.VideoName)
	if err != nil {
		logging.Error("Failed to video to mp4: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, persistentId)
}

func CheckMp4Status(ctx *gin.Context) {
	var params struct {
		PersistentId string `json:"persistentId" validate:"required"`
	}

	if err := common.BindAndValidate(ctx, &params); err != nil {
		return
	}

	ret, err := service.CheckMp4Status(params.PersistentId)
	if err != nil {
		logging.Error("Failed to check mp4 status: %v", err)
		common.ResponseError(ctx, common.ERROR)
		return
	}

	common.ResponseSuccess(ctx, ret)
}
