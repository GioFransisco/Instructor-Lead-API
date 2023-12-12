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

type QuestionControllerSuite struct {
	suite.Suite
	qum *usecasemock.QuestionUsecaseMock
	rg  *gin.RouterGroup
	amm *middlewaremock.AuthMiddlewareMock
}

func (s *QuestionControllerSuite) SetupTest() {
	s.qum = new(usecasemock.QuestionUsecaseMock)
	rg := gin.Default()
	s.rg = rg.Group("api/v1")
	s.amm = new(middlewaremock.AuthMiddlewareMock)
}

func TestQuestionController(t *testing.T) {
	suite.Run(t, new(QuestionControllerSuite))
}

var mockQuestion = model.Question{
	Id: "1",
	ScheduleDetails: model.ScheduleDetails{
		Id:         "1",
		ScheduleId: "1",
		Trainer: model.User{
			Id:          "1",
			Name:        "jamal",
			Email:       "juned",
			PhoneNumber: "0893298219213",
			Username:    "jakalksa",
			Password:    "password",
			Age:         19,
			Address:     "indonseia timur",
			Gander:      "L",
			Role:        "Trainer",
			CreatedAt:   time.Now().Round(0),
			UpdatedAt:   time.Now().Round(0),
		},
		Stack: model.Stack{
			Id:        "1",
			Name:      "Golang",
			Status:    "Active",
			CreatedAt: time.Now().Round(0),
			UpdatedAt: time.Now().Round(0),
		},
		StartTime: time.Now().Round(0),
		EndTime:   time.Now().Round(0),
		CreatedAt: time.Now().Round(0),
		UpdatedAt: time.Now().Round(0),
	},
	StudentId: model.User{
		Id:          "1",
		Name:        "joko kendi;",
		Email:       "leskass@gmail.com",
		PhoneNumber: "08736273282",
		Username:    "jamal",
		Password:    "password",
		Age:         19,
		Address:     "Indramayu Barat",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now().Round(0),
		UpdatedAt:   time.Now().Round(0),
	},
	Question:  "apakabar gaess",
	Status:    "Proccess",
	CreatedAt: time.Now().Round(0),
	UpdatedAt: time.Now().Round(0),
}

var dtoMockResponse = dto.QuestionResponseUpdate{
	Id:              "1",
	ScheduleDetails: "1",
	StudentId:       "1",
	Question:        "apakabar gaess",
	Status:          "Proccess",
	CreatedAt:       time.Now().Round(0),
	UpdatedAt:       time.Now().Round(0),
}

var dtoMockResponseGetSlice = []dto.QuestionResponseGET{
	{
		Id:              "1",
		ScheduleDetails: "1",
		StudentId: model.User{
			Id:          "1",
			Name:        "joko kendi;",
			Email:       "leskass@gmail.com",
			PhoneNumber: "08736273282",
			Username:    "jamal",
			Password:    "password",
			Age:         19,
			Address:     "Indramayu Barat",
			Gander:      "L",
			Role:        "Participant",
			CreatedAt:   time.Now().Round(0),
			UpdatedAt:   time.Now().Round(0),
		},
		Question:  "apakabar gaess",
		Status:    "Proccess",
		CreatedAt: time.Now().Round(0),
		UpdatedAt: time.Now().Round(0),
	},
}

var dtoMockResponseGet = dto.QuestionResponseGET{
	Id:              "1",
	ScheduleDetails: "1",
	StudentId: model.User{
		Id:          "1",
		Name:        "joko kendi;",
		Email:       "leskass@gmail.com",
		PhoneNumber: "08736273282",
		Username:    "jamal",
		Password:    "password",
		Age:         19,
		Address:     "Indramayu Barat",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now().Round(0),
		UpdatedAt:   time.Now().Round(0),
	},
	Question:  "apakabar gaess",
	Status:    "Proccess",
	CreatedAt: time.Now().Round(0),
	UpdatedAt: time.Now().Round(0),
}

var dtoMockQuestionChangeStatus = dto.QuestionChangeStatusDto{
	Id:     "",
	Status: "",
}

var dtoMockQuestionChange = dto.QuestionChangeDto{
	Id:       "",
	Question: "",
}

func (s *QuestionControllerSuite) TestCreateQuestion_Success() {
	s.qum.On("CreateQuestion", mockQuestion).Return(dtoMockResponseGet, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockQuestion)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/questions", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.createQuestion(ctx)
	assert.Equal(s.T(), http.StatusCreated, record.Code)
}

func (s *QuestionControllerSuite) TestCreateQuestion_FailReq() {
	s.qum.On("CreateQuestion", mockQuestion).Return(dtoMockResponseGet, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/questions", nil)
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.createQuestion(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *QuestionControllerSuite) TestCreateQuestion_FailCreate() {
	s.qum.On("CreateQuestion", mockQuestion).Return(dto.QuestionResponseGET{}, errors.New("error"))
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockQuestion)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/questions", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.createQuestion(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *QuestionControllerSuite) TestFindQuestionByID_Success() {
	s.qum.On("FindQuestionByID", "").Return(dtoMockResponseGetSlice, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/questions/ksjdaoio2", nil)

	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.findQuestionByID(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *QuestionControllerSuite) TestFindQuestionByID_Fail() {
	s.qum.On("FindQuestionByID", "").Return([]dto.QuestionResponseGET{}, errors.New("error"))
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/questions/ksjdaoio2", nil)
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.findQuestionByID(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *QuestionControllerSuite) TestUpdateStatusQuestion_Success() {
	s.qum.On("UpdateStatusQuestion", dtoMockQuestionChangeStatus).Return(dtoMockResponse, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()
	dataJson, err := json.Marshal(dtoMockQuestionChangeStatus)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/questions/1/status", bytes.NewBuffer(dataJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.updateStatusQuestion(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *QuestionControllerSuite) TestUpdateStatusQuestion_FailReq() {
	s.qum.On("UpdateStatusQuestion", dtoMockQuestionChangeStatus).Return(dtoMockResponse, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodPut, "/api/v1/questions/1/status", nil)
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.updateStatusQuestion(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *QuestionControllerSuite) TestUpdateStatusQuestion_FailUpdate() {
	s.qum.On("UpdateStatusQuestion", dtoMockQuestionChangeStatus).Return(dto.QuestionResponseUpdate{}, errors.New("error"))
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()
	dataJson, err := json.Marshal(dtoMockQuestionChangeStatus)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/questions/1/status", bytes.NewBuffer(dataJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.updateStatusQuestion(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *QuestionControllerSuite) TestUpdateQuestion() {
	s.qum.On("UpdateQuestion", dtoMockQuestionChange).Return(dtoMockResponse, nil)
	questionController := NewQuestionController(s.qum, s.rg, s.amm)
	questionController.QuestionRouter()

	record := httptest.NewRecorder()
	dataJson, err := json.Marshal(dtoMockQuestionChange)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/questions/1/question", bytes.NewBuffer(dataJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	questionController.updateQuestion(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}
