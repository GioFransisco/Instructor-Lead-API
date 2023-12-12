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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AbsencesControllerTestSuite struct {
	suite.Suite
	rg                 *gin.RouterGroup
	absenceUseCaseMock *usecasemock.AbsencesUseCaseMock
	authMiddleware     *middlewaremock.AuthMiddlewareMock
	absencesController *absenceController
	record             *httptest.ResponseRecorder
}

var mockGetAbsences = model.GetAbsences{
	Id: "1",
	ScheduleDetails: []model.GetScheduleDetails{
		{
			Id: "1",
			Schedule: model.Schedule{
				Id:           "1",
				Name:         "Test name",
				DateActivity: time.Now(),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
			Trainer: model.User{
				Id:          "1",
				Name:        "Yopi",
				Email:       "yopitn@email.com",
				PhoneNumber: "089768758274",
				Username:    "yopitn",
				Age:         23,
				Address:     "Garut",
				Gander:      "L",
				Role:        "Trainer",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			Stack: model.Stack{
				Id:        "1",
				Name:      "Golang",
				Status:    "Active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	},
	Description: "Hadir",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

var mockAbsence = dto.AbsencesResponseDto{
	Id:                "1",
	ScheduleDetailsId: "1",
	StudentId: model.User{
		Id:          "2",
		Name:        "Dina",
		Email:       "dina@email.com",
		PhoneNumber: "089768758274",
		Username:    "dina",
		Age:         23,
		Address:     "Garut",
		Gander:      "P",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	Description: "Hadir",
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
}

var mockPayloadAbsences = model.Absences{
	Id: "1",
	ScheduleDetails: model.ScheduleDetails{
		Id: "1",
	},
	StudentId: model.User{
		Id: "2",
	},
	Description: "Hadir",
}

func (suite *AbsencesControllerTestSuite) SetupTest() {
	engine := gin.Default()
	suite.rg = engine.Group("/api/v1")
	suite.absenceUseCaseMock = new(usecasemock.AbsencesUseCaseMock)
	suite.authMiddleware = new(middlewaremock.AuthMiddlewareMock)
	suite.absencesController = NewAbsencesController(suite.absenceUseCaseMock, suite.rg, suite.authMiddleware)
	suite.record = httptest.NewRecorder()

	suite.absencesController.Route()
}

func (suite *AbsencesControllerTestSuite) TestCreateHandler_Success() {
	suite.absenceUseCaseMock.On("CreateNewAbsence", mockPayloadAbsences).Return(mockAbsence, nil)

	mockPayloadJSON, err := json.Marshal(mockPayloadAbsences)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/absences", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.absencesController.createAbsenceHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, suite.record.Code)
}

func (suite *AbsencesControllerTestSuite) TestCreateHandler_Fail_OnBind() {
	suite.absenceUseCaseMock.On("CreateNewAbsence", mockPayloadAbsences).Return(mockAbsence, nil)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/absences", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.absencesController.createAbsenceHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *AbsencesControllerTestSuite) TestCreateHandler_Fail_OnCreate() {
	suite.absenceUseCaseMock.On("CreateNewAbsence", mockPayloadAbsences).Return(dto.AbsencesResponseDto{}, errors.New("error"))

	mockPayloadJSON, err := json.Marshal(mockPayloadAbsences)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/absences", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.absencesController.createAbsenceHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *AbsencesControllerTestSuite) TestGetAbsenceHandler_Success() {
	req, err := http.NewRequest(http.MethodPost, "/api/v1/absences/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleId := ctx.Param("id")
	suite.absenceUseCaseMock.On("FindAbsenceById", scheduleId).Return(mockGetAbsences, nil)

	suite.absencesController.getAbsenceHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *AbsencesControllerTestSuite) TestGetAbsenceHandler_Fail() {
	req, err := http.NewRequest(http.MethodPost, "/api/v1/absences/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	scheduleId := ctx.Param("id")
	suite.absenceUseCaseMock.On("FindAbsenceById", scheduleId).Return(model.GetAbsences{}, errors.New("error"))

	suite.absencesController.getAbsenceHandler(ctx)
	assert.Equal(suite.T(), http.StatusNotFound, suite.record.Code)
}

func TestAbsencesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AbsencesControllerTestSuite))
}
