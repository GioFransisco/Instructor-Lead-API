package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	middlewaremock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/middleware_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScheduleControllerTestSuite struct {
	suite.Suite
	rg                  *gin.RouterGroup
	scheduleUseCaseMock *usecasemock.ScheduleUseCaseMock
	trainerUseCaseMock  *usecasemock.UserUCMock
	stackUseCaseMock    *usecasemock.StackUseCaseMock
	authMiddleware      *middlewaremock.AuthMiddlewareMock
	scheduleController  *ScheduleController
	record              *httptest.ResponseRecorder
}

func (suite *ScheduleControllerTestSuite) SetupTest() {
	engine := gin.Default()
	suite.rg = engine.Group("/api/v1")
	suite.scheduleUseCaseMock = new(usecasemock.ScheduleUseCaseMock)
	suite.authMiddleware = new(middlewaremock.AuthMiddlewareMock)
	suite.scheduleController = NewScheduleController(suite.scheduleUseCaseMock, suite.rg, suite.authMiddleware)
	suite.record = httptest.NewRecorder()

	suite.scheduleController.Route()
}

var mockSchedule = dto.ScheduleResponseDto{
	Id:           "1",
	Name:         "Test tanggal 19",
	DateActivity: "2023-11-19",
	ScheduleDetails: []dto.ScheduleDetailResponseDto{
		{
			Trainer: model.User{
				Id:        "1",
				Name:      "Yopi",
				Role:      "Admin",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Stack: model.Stack{
				Id:        "1",
				Name:      "Golang",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartTime: "19:00",
			EndTime:   "20:00",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayloadSchedule = dto.ScheduleCreateRequestDto{
	Name:         "Test tanggal 19",
	DateActivity: "2023-11-19",
	ScheduleDetails: []dto.ScheduleDetailCreateRequestDto{
		{
			TrainerId: "1",
			StackId:   "1",
			StartTime: "19:00",
			EndTime:   "20:00",
		},
	},
}

func (suite *ScheduleControllerTestSuite) TestCreateHandler_Success() {
	mockPayloadJSON, err := json.Marshal(mockPayloadSchedule)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/schedules", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleUseCaseMock.On("RegisterNewSchedule", mockPayloadSchedule).Return(mockSchedule, nil)

	suite.scheduleController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestCreateHandler_Fail() {
	req, err := http.NewRequest(http.MethodPost, "/api/v1/schedules", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	mockPayloadJSON, err := json.Marshal(mockPayloadSchedule)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/schedules", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ = gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleUseCaseMock.On("RegisterNewSchedule", mockPayloadSchedule).Return(dto.ScheduleResponseDto{}, common.InvalidError{Message: "error"})

	suite.scheduleController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestCreateHandler_Fail_500() {
	mockPayloadJSON, err := json.Marshal(mockPayloadSchedule)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/schedules", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleUseCaseMock.On("RegisterNewSchedule", mockPayloadSchedule).Return(dto.ScheduleResponseDto{}, errors.New("error"))

	suite.scheduleController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestListHandler_Success() {
	mockSchedules := []dto.ScheduleResponseDto{mockSchedule}

	req, err := http.NewRequest(http.MethodGet, "/api/v1/schedules", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	ctx.Set("userId", mockSchedule.ScheduleDetails[0].Trainer.Id)
	ctx.Set("userRole", mockSchedule.ScheduleDetails[0].Trainer.Role)
	suite.scheduleUseCaseMock.On("FindAll", "1", "Admin").Return(mockSchedules, nil)

	suite.scheduleController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestListHandler_Fail() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/schedules", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	ctx.Set("userId", mockSchedule.ScheduleDetails[0].Trainer.Id)
	ctx.Set("userRole", mockSchedule.ScheduleDetails[0].Trainer.Role)
	suite.scheduleUseCaseMock.On("FindAll", "1", "Admin").Return([]dto.ScheduleResponseDto{}, errors.New("error"))

	suite.scheduleController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleHandler_Success() {
	mockPayload := dto.ScheduleUpdateRequestDto{
		Name:         mockPayloadSchedule.Name,
		DateActivity: mockPayloadSchedule.DateActivity,
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedules/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateSchedule", scheduleId, mockPayload).Return(dto.ScheduleResponseDto{}, nil)

	suite.scheduleController.updateScheduleHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleHandler_Fail() {
	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedules/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleController.updateScheduleHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	mockPayload := dto.ScheduleUpdateRequestDto{
		Name:         mockPayloadSchedule.Name,
		DateActivity: mockPayloadSchedule.DateActivity,
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodPut, "/api/v1/schedules/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ = gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateSchedule", scheduleId, mockPayload).Return(dto.ScheduleResponseDto{}, common.InvalidError{Message: "error"})

	suite.scheduleController.updateScheduleHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleHandler_Fail_500() {
	mockPayload := dto.ScheduleUpdateRequestDto{
		Name:         mockPayloadSchedule.Name,
		DateActivity: mockPayloadSchedule.DateActivity,
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedules/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateSchedule", scheduleId, mockPayload).Return(dto.ScheduleResponseDto{}, errors.New("error"))

	suite.scheduleController.updateScheduleHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleDetailHandler_Success() {
	mockPayload := dto.ScheduleDetailUpdateRequestDto{
		TrainerId: "1",
		StackId:   "1",
		StartTime: "19:00",
		EndTime:   "20:00",
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedule-details/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleDetailId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateScheduleDetail", scheduleDetailId, mockPayload).Return(dto.ScheduleDetailResponseDto{}, nil)

	suite.scheduleController.updateScheduleDetailHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleDetailHandler_Fail() {
	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedule-details/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.scheduleController.updateScheduleDetailHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	mockPayload := dto.ScheduleDetailUpdateRequestDto{
		TrainerId: "1",
		StackId:   "1",
		StartTime: "19:00",
		EndTime:   "20:00",
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodPut, "/api/v1/schedule-details/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ = gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleDetailId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateScheduleDetail", scheduleDetailId, mockPayload).Return(dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: "error"})

	suite.scheduleController.updateScheduleDetailHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *ScheduleControllerTestSuite) TestUpdateScheduleDetailHandler_Fail_500() {
	mockPayload := dto.ScheduleDetailUpdateRequestDto{
		TrainerId: "1",
		StackId:   "1",
		StartTime: "19:00",
		EndTime:   "20:00",
	}

	mockPayloadJSON, err := json.Marshal(mockPayload)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/schedule-details/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleDetailId := ctx.Param("id")
	suite.scheduleUseCaseMock.On("UpdateScheduleDetail", scheduleDetailId, mockPayload).Return(dto.ScheduleDetailResponseDto{}, errors.New("error"))

	suite.scheduleController.updateScheduleDetailHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func TestScheduleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleControllerTestSuite))
}
