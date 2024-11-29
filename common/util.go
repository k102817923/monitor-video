package common

import (
	"monitor-video/logging"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate(ctx *gin.Context, params interface{}) error {
	if err := ctx.ShouldBindJSON(params); err != nil {
		ResponseError(ctx, INVALID_PARAMS)
		return err
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			logging.Error("Validation error: %s %s\n", err.Field(), err.Tag())
		}
		ResponseError(ctx, INVALID_PARAMS)
		return err
	}

	return nil
}
