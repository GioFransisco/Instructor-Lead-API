package controller

import (
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type StackController struct {
	uc             usecase.StackUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.JwtMiddleware
}

func (s *StackController) createHandler(ctx *gin.Context) {
	var payload dto.StackRequestDto

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)
		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}

	stack, err := s.uc.RegisterNewStack(payload)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusCreated, stack)
}

func (s *StackController) listHandler(ctx *gin.Context) {
	stacks, err := s.uc.FindAll()
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

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, stacks)
}

func NewStackController(uc usecase.StackUseCase, rg *gin.RouterGroup, authMiddleware middleware.JwtMiddleware) *StackController {
	return &StackController{uc: uc, rg: rg, authMiddleware: authMiddleware}
}

func (s *StackController) getHandler(ctx *gin.Context) {
	stackID := ctx.Param("id")
	stack, err := s.uc.FindByID(stackID)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, stack)
}

func (s *StackController) updateHandler(ctx *gin.Context) {
	stackID := ctx.Param("id")

	var payload model.Stack
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)
		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}

	stack, err := s.uc.UpdateStack(stackID, payload)
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

	common.ResponseSuccess(ctx, "Ok", http.StatusOK, stack)
}

func (s *StackController) deleteHandler(ctx *gin.Context) {
	stackID := ctx.Param("id")
	err := s.uc.DeleteStack(stackID)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "Ok", http.StatusNoContent, nil)
}

func (s *StackController) Route() {
	stackGroup := s.rg.Group(config.StackGroupPath, s.authMiddleware.AuthMiddleware(("Admin")))

	stackGroup.POST(config.StackCreatePath, s.createHandler)
	stackGroup.GET(config.StackGetPath, s.listHandler)
	stackGroup.GET(config.StackGetByIdPath, s.getHandler)
	stackGroup.PUT(config.StackUpdatePath, s.updateHandler)
	stackGroup.DELETE(config.StackDeletePath, s.deleteHandler)
}
