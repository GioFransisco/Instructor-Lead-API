package controller

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"math/rand"
	"time"
	"fmt"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/config"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/driver/middleware"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/usecase"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type scheduleApproveController struct {
	scApproveUC    usecase.ScheduleApproveUseCase
	rg             *gin.RouterGroup
	authMiddleware middleware.JwtMiddleware
}

func (sa *scheduleApproveController) createScheduleApproveHandler(ctx *gin.Context) {
	var payload model.ScheduleAprove
	schedule := ctx.PostForm("schedule")
	json.Unmarshal([]byte(schedule), &payload)

	// upload file code
	file, header, _ := ctx.Request.FormFile("photo")

	if err := common.ValidateUploadFile(file, header);err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	
	newFileName := fmt.Sprintf("%v_photo%s", rand.New(rand.NewSource(time.Now().UTC().UnixNano())).Int(), filepath.Ext(header.Filename))
	fileLocation := filepath.Join("asset/image", newFileName)
	os.MkdirAll("asset/image", os.ModePerm)

	payload.ScheduleAprove = header.Filename
	
	scApprove, err := sa.scApproveUC.CreateNewScheduleAprove(payload)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ctx.SaveUploadedFile(header, fileLocation);err != nil{
		common.ResponseError(ctx, "error when saving upload file", http.StatusInternalServerError)
		return
	}

	common.ResponseSuccess(ctx, "OK", http.StatusCreated, scApprove)

}

func (sa *scheduleApproveController) getSCheduleApproveByIdHandler(ctx *gin.Context) {
	schDetailID := ctx.Param("schDetailID")

	schApprove, err := sa.scApproveUC.FindSchApproveById(schDetailID)
	if err != nil {
		common.ResponseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	common.ResponseSuccess(ctx, "OK", http.StatusOK, schApprove)
}

func (sa *scheduleApproveController) Route() {
	v2 := sa.rg.Group(config.ScheduleAprovePath)
	v2.POST("", sa.authMiddleware.AuthMiddleware("Trainer"), sa.createScheduleApproveHandler)
	v2.GET(":schDetailID", sa.authMiddleware.AuthMiddleware("Admin"), sa.getSCheduleApproveByIdHandler)
}

func NewScheduleApproveController(scApproveUC usecase.ScheduleApproveUseCase, rg *gin.RouterGroup, authMiddleware middleware.JwtMiddleware) *scheduleApproveController {
	return &scheduleApproveController{
		scApproveUC:    scApproveUC,
		rg:             rg,
		authMiddleware: authMiddleware,
	}
}