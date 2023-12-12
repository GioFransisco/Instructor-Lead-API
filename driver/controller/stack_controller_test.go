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

type StackControllerTestSuite struct {
	suite.Suite
	rg               *gin.RouterGroup
	stackUseCaseMock *usecasemock.StackUseCaseMock
	authMiddleware   *middlewaremock.AuthMiddlewareMock
	stackController  *StackController
	record           *httptest.ResponseRecorder
}

func (suite *StackControllerTestSuite) SetupTest() {
	engine := gin.Default()
	suite.rg = engine.Group("/api/v1")
	suite.stackUseCaseMock = new(usecasemock.StackUseCaseMock)
	suite.authMiddleware = new(middlewaremock.AuthMiddlewareMock)
	suite.stackController = NewStackController(suite.stackUseCaseMock, suite.rg, suite.authMiddleware)
	suite.record = httptest.NewRecorder()

	suite.stackController.Route()
}

var mockStack = model.Stack{
	Id:        "1",
	Name:      "Golang",
	Status:    "Active",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayloadStack = dto.StackRequestDto{
	Name:   "Golang",
	Status: "Active",
}

func (suite *StackControllerTestSuite) TestCreateHandler_Success() {
	suite.stackUseCaseMock.On("RegisterNewStack", mockPayloadStack).Return(mockStack, nil)

	mockPayloadJSON, err := json.Marshal(mockPayloadStack)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/stacks", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestCreateHandler_Fail() {
	req, err := http.NewRequest(http.MethodPost, "/api/v1/stacks", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	suite.stackUseCaseMock.On("RegisterNewStack", mockPayloadStack).Return(model.Stack{}, errors.New("error when creating stack"))

	mockPayloadJSON, err := json.Marshal(mockPayloadStack)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodPost, "/api/v1/stacks", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx.Request = req

	suite.stackController.createHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestListHandler_Success() {
	mockStacks := []model.Stack{
		{
			Id:        "1",
			Name:      "Golang",
			Status:    "Active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        "2",
			Name:      "Java",
			Status:    "Active",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	suite.stackUseCaseMock.On("FindAll").Return(mockStacks, nil)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/stacks", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestListHandler_Fail_400() {
	suite.stackUseCaseMock.On("FindAll").Return([]model.Stack{}, common.InvalidError{Message: "error"})

	req, err := http.NewRequest(http.MethodGet, "/api/v1/stacks", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestListHandler_Fail_500() {
	suite.stackUseCaseMock.On("FindAll").Return([]model.Stack{}, errors.New("error when get data from FindAll usecase"))

	req, err := http.NewRequest(http.MethodGet, "/api/v1/stacks", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestGetHandler_Success() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/stacks/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("FindByID", stackId).Return(mockStack, nil)

	suite.stackController.getHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestGetHandler_Fail() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/stacks/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("FindByID", stackId).Return(model.Stack{}, errors.New("errors get stack by ID"))

	suite.stackController.getHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestUpdateHandler_Success() {
	payload := model.Stack{
		Name:   mockPayloadStack.Name,
		Status: mockPayloadStack.Status,
	}

	mockPayloadJSON, err := json.Marshal(mockPayloadStack)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/stacks/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("UpdateStack", stackId, payload).Return(mockStack, nil)

	suite.stackController.updateHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestUpdateHandler_Fail_400() {
	payload := model.Stack{
		Name:   mockPayloadStack.Name,
		Status: mockPayloadStack.Status,
	}

	req, err := http.NewRequest(http.MethodPut, "/api/v1/stacks/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.stackController.updateHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	mockPayloadJSON, err := json.Marshal(mockPayloadStack)
	assert.NoError(suite.T(), err)

	req, err = http.NewRequest(http.MethodPut, "/api/v1/stacks/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("UpdateStack", stackId, payload).Return(model.Stack{}, common.InvalidError{Message: "error"})

	suite.stackController.updateHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestUpdateHandler_Fail_500() {
	payload := model.Stack{
		Name:   mockPayloadStack.Name,
		Status: mockPayloadStack.Status,
	}

	mockPayloadJSON, err := json.Marshal(mockPayloadStack)
	assert.NoError(suite.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/stacks/1", bytes.NewBuffer(mockPayloadJSON))
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("UpdateStack", stackId, payload).Return(model.Stack{}, errors.New("error when update stack data"))

	suite.stackController.updateHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestDeleteHandler_Success() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/stacks/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("DeleteStack", stackId).Return(nil)

	suite.stackController.deleteHandler(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, suite.record.Code)
}

func (suite *StackControllerTestSuite) TestDeleteHandler_Fail() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/stacks/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	stackId := ctx.Param("id")
	suite.stackUseCaseMock.On("DeleteStack", stackId).Return(errors.New("errors when deleting stack"))

	suite.stackController.deleteHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func TestStackControllerTestSuite(t *testing.T) {
	suite.Run(t, new(StackControllerTestSuite))
}
