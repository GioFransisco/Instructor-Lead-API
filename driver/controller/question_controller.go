package controller

import (
	"errors"
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type questionController struct {
	questionUC     usecase.QuestionUsecase
	authMiddleware middleware.JwtMiddleware
	rg             *gin.RouterGroup
}

func (c *questionController) createQuestion(ctx *gin.Context) {
	var payloadQuestion model.Question

	if err := ctx.ShouldBindJSON(&payloadQuestion); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	question, err := c.questionUC.CreateQuestion(payloadQuestion)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusCreated, question)
}

func (c *questionController) findQuestionByID(ctx *gin.Context) {
	id := ctx.Param("scheduleDetailId")

	question, err := c.questionUC.FindQuestionByID(id)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, question)
}

func (c *questionController) updateQuestion(ctx *gin.Context) {
	var payloadDto dto.QuestionChangeDto

	if err := ctx.ShouldBind(&payloadDto); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	payloadDto.Id = ctx.Param("id")

	question, err := c.questionUC.UpdateQuestion(payloadDto)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, question)
}

func (c *questionController) updateStatusQuestion(ctx *gin.Context) {
	var payloadDto dto.QuestionChangeStatusDto

	if err := ctx.ShouldBind(&payloadDto); err != nil {
		common.ResponseError(ctx, errors.New(config.ErrorDescriptionForInvalidData).Error(), http.StatusBadRequest)
		return
	}

	payloadDto.Id = ctx.Param("id")

	question, err := c.questionUC.UpdateStatusQuestion(payloadDto)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, question)
}

func (c *questionController) deleteQuestion(ctx *gin.Context) {
	id := ctx.Param("id")

	question, err := c.questionUC.DeleteQuestion(id)

	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusOK, question)
}

func (c *questionController) QuestionRouter() {
	v2 := c.rg.Group(config.QuestionCreatePath)
	v2.POST("", c.authMiddleware.AuthMiddleware(config.RolePartcipant), c.createQuestion)
	v2.GET(config.QuestionGetPath, c.authMiddleware.AuthMiddleware(config.RolePartcipant, config.RoleTrainer, config.RoleAdmin), c.findQuestionByID)
	v2.PUT(config.QuestionUpdatePath, c.authMiddleware.AuthMiddleware(config.RolePartcipant), c.updateQuestion)
	v2.PUT(config.QuestionUpdateStatusPath, c.authMiddleware.AuthMiddleware(config.RoleTrainer), c.updateStatusQuestion)
	v2.DELETE(config.QuestionDeletePath, c.authMiddleware.AuthMiddleware(config.RoleAdmin, config.RolePartcipant), c.deleteQuestion)
}

func NewQuestionController(questionUC usecase.QuestionUsecase, rg *gin.RouterGroup, authMiddleware middleware.JwtMiddleware) *questionController {
	return &questionController{
		questionUC:     questionUC,
		authMiddleware: authMiddleware,
		rg:             rg,
	}
}
