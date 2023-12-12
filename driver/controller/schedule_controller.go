package controller

import (
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	uc             usecase.ScheduleUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.JwtMiddleware
}

func (s *ScheduleController) createHandler(ctx *gin.Context) {
	var payload dto.ScheduleCreateRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)

		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}

	schedule, err := s.uc.RegisterNewSchedule(payload)
	if err != nil {
		switch err.(type) {
		case common.InvalidError:
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		default:
			common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusCreated, schedule)
}

func (s *ScheduleController) listHandler(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(string)
	userRole := ctx.MustGet("userRole").(string)

	schedules, err := s.uc.FindAll(userId, userRole)
	if err != nil {
		switch err.(type) {
		case common.InvalidError:
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		default:
			common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, schedules)
}

func (s *ScheduleController) updateScheduleHandler(ctx *gin.Context) {
	var payload dto.ScheduleUpdateRequestDto

	scheduleId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)

		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}

	schedule, err := s.uc.UpdateSchedule(scheduleId, payload)
	if err != nil {
		switch err.(type) {
		case common.InvalidError:
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		default:
			common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, schedule)
}

func (s *ScheduleController) updateScheduleDetailHandler(ctx *gin.Context) {
	var payload dto.ScheduleDetailUpdateRequestDto

	scheduleDetailId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)

		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}

	scheduleDetail, err := s.uc.UpdateScheduleDetail(scheduleDetailId, payload)
	if err != nil {
		switch err.(type) {
		case common.InvalidError:
			common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
			return
		default:
			common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, scheduleDetail)
}

func (s *ScheduleController) Route() {
	s.rg.POST(config.ScheduleCreatePath, s.authMiddleware.AuthMiddleware("Admin"), s.createHandler)
	s.rg.GET(config.ScheduleGetPath, s.authMiddleware.AuthMiddleware("Admin", "Trainer", "Participant"), s.listHandler)
	s.rg.PUT(config.ScheduleUpdatePath, s.authMiddleware.AuthMiddleware("Admin"), s.updateScheduleHandler)
	s.rg.PUT(config.ScheduleDetailUpdatePath, s.authMiddleware.AuthMiddleware("Admin"), s.updateScheduleDetailHandler)
}

func NewScheduleController(uc usecase.ScheduleUseCase, rg *gin.RouterGroup, authMiddleware middleware.JwtMiddleware) *ScheduleController {
	return &ScheduleController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}
