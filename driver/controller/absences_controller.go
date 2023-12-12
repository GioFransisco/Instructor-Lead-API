package controller

import (
	"net/http"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type absenceController struct {
	absencesUC     usecase.AbsencesUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.JwtMiddleware
}

func (a *absenceController) createAbsenceHandler(ctx *gin.Context) {
	var payload model.Absences
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		errMsg := common.ValidationErrors(err)
		common.ResponseError(ctx, errMsg, http.StatusBadRequest)
		return
	}
	absence, err := a.absencesUC.CreateNewAbsence(payload)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	common.ResponseSuccess(ctx, "OK", http.StatusCreated, absence)
}

func (a *absenceController) getAbsenceHandler(ctx *gin.Context) {
	scheduleId := ctx.Param("id")
	absence, err := a.absencesUC.FindAbsenceById(scheduleId)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusNotFound)
		return
	}
	common.ResponseSuccess(ctx, "OK", http.StatusOK, absence)
}

func (a *absenceController) Route() {
	v2 := a.rg.Group(config.AbsencesCreatePath)
	v2.POST("", a.authMiddleware.AuthMiddleware("Trainer"), a.createAbsenceHandler)
	v2.GET(":id", a.authMiddleware.AuthMiddleware("Trainer", "Admin"), a.getAbsenceHandler)
}

func NewAbsencesController(absencesUC usecase.AbsencesUseCase, rg *gin.RouterGroup, authMiddleware middleware.JwtMiddleware) *absenceController {
	return &absenceController{
		absencesUC:     absencesUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}
