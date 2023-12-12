package common

import (
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
	"github.com/gin-gonic/gin"
)

func ResponseSuccess(ctx *gin.Context, des string, code int, data any) {
	ctx.JSON(code, utilsmodel.ResponseSuccess{
		Response: utilsmodel.Response{
			Code:        code,
			Description: des,
		},
		Data: data,
	})
}

func ResponseError(ctx *gin.Context, des string, code int) {
	ctx.JSON(code, utilsmodel.ResponseError{
		Response: utilsmodel.Response{
			Code:        code,
			Description: des,
		},
	})
}
