package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

type ScheduleAproveControllerTestSuite struct {
	suite.Suite
	rg                         *gin.RouterGroup
	scheduleApproveUseCaseMock *usecasemock.ScheduleApproveUseCaseMock
	authMiddleware             *middlewaremock.AuthMiddlewareMock
	scheduleApproveController  *scheduleApproveController
	record                     *httptest.ResponseRecorder
}

func (suite *ScheduleAproveControllerTestSuite) SetupTest() {
	engine := gin.Default()
	suite.rg = engine.Group("/api/v1")
	suite.scheduleApproveUseCaseMock = new(usecasemock.ScheduleApproveUseCaseMock)
	suite.authMiddleware = new(middlewaremock.AuthMiddlewareMock)
	suite.scheduleApproveController = NewScheduleApproveController(suite.scheduleApproveUseCaseMock, suite.rg, suite.authMiddleware)
	suite.record = httptest.NewRecorder()

	suite.scheduleApproveController.Route()
}

func (suite *ScheduleAproveControllerTestSuite) TestCreateHandler_Success() {
	payload := model.ScheduleAprove{
		Id: "1",
		ScheduleDetails: model.ScheduleDetails{
			Id: "1",
		},
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	mockPayloadJSON, err := json.Marshal(payload)
	assert.NoError(suite.T(), err)

	scheduleField, err := writer.CreateFormField("schedule")
	assert.NoError(suite.T(), err)
	scheduleField.Write(mockPayloadJSON)

	fileContent := []byte("test file content")
	tmpFile, err := os.CreateTemp("", "test*.png")
	assert.NoError(suite.T(), err)

	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	tmpFile.Write(fileContent)

	photoField, err := writer.CreateFormFile("photo", filepath.Base(tmpFile.Name()))
	assert.NoError(suite.T(), err)

	file, err := os.Open(tmpFile.Name())
	assert.NoError(suite.T(), err)
	defer file.Close()

	io.Copy(photoField, file)
	writer.Close()

	req, err := http.NewRequest(http.MethodPost, "/api/v1/schedule-aprove", body)
	assert.NoError(suite.T(), err)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	payload.ScheduleAprove = filepath.Base(tmpFile.Name())

	response := dto.ScheduleAproveResponseDto{
		Id:                payload.Id,
		ScheduleDetailsId: payload.ScheduleDetails.Id,
		ScheduleAprove:    payload.ScheduleAprove,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	suite.scheduleApproveUseCaseMock.On("CreateNewScheduleAprove", payload).Return(response, nil)

	suite.scheduleApproveController.createScheduleApproveHandler(ctx)
	assert.Equal(suite.T(), http.StatusCreated, suite.record.Code)
}

func (suite *ScheduleAproveControllerTestSuite) TestGetHandler_Success() {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/schedule-aprove/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	schDetailID := ctx.Param("schDetailID")
	suite.scheduleApproveUseCaseMock.On("FindSchApproveById", schDetailID).Return([]dto.ScheduleAproveResponseDto{}, nil)

	suite.scheduleApproveController.getSCheduleApproveByIdHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func TestScheduleAproveControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleAproveControllerTestSuite))
}
