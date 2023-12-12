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

type AuthController struct {
	suite.Suite
	aum *usecasemock.AuthUserMock
	rg  *gin.RouterGroup
	amm *middlewaremock.AuthMiddlewareMock
}

func (s *AuthController) SetupTest() {
	s.aum = new(usecasemock.AuthUserMock)
	rg := gin.Default()
	s.rg = rg.Group("api/v1")
	s.amm = new(middlewaremock.AuthMiddlewareMock)
}

func TestAuthController(t *testing.T) {
	suite.Run(t, new(AuthController))
}

var mockAuth = model.User{
	Id:          "kdjksooaieoi",
	Name:        "kdjksooaieoi",
	Email:       "kdjksooaieoi",
	PhoneNumber: "kdjksooaieoi",
	Username:    "kdjksooaieoi",
	Password:    "kdjksooaieoi",
	Age:         18,
	Address:     "kdjksooaieoild Kol",
	Gander:      "L",
	Role:        "Participant",
	CreatedAt:   time.Now().Round(0),
	UpdatedAt:   time.Now().Round(0),
}

var mockUSerDto = dto.UserLoginDto{
	Username: "kdjksooaieoi",
	Password: "kdjksooaieoi",
}

var mockTokenJwt = model.TokenModel{
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJncm91cDEiLCJleHAiOjE3MDAzMjM0MDksImlhdCI6MTcwMDMxNjIwOSwidXNlcklkIjoiMjlmYmEzZTktZDAzYS00ODY0LWFkYTctZDViYjdkMzA4YWFkIiwiZW1haWwiOiJha2JhcnJhdzA5QGdtYWlsLmNvbSIsInJvbGUiOiJBZG1pbiJ9.voTa_GJFgyGbRvthdyDJ1DvLGpUQ1tD6OPJ0ZIOjFro",
}

func (s *AuthController) TestFindByUsername_Success() {
	s.aum.On("FindByUsername", mockUSerDto).Return(mockTokenJwt, nil)
	authController := NewAuthController(s.aum, s.rg, s.amm)
	authController.Router()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	authController.findByUsername(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *AuthController) TestFindByUsername_FailReq() {
	s.aum.On("FindByUsername", mockUSerDto).Return(model.TokenModel{}, errors.New("error"))
	authController := NewAuthController(s.aum, s.rg, s.amm)
	authController.Router()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON

	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	authController.findByUsername(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *AuthController) TestFindByUsername_FailFind() {
	s.aum.On("FindByUsername", mockUSerDto).Return(model.TokenModel{}, errors.New("error"))
	authController := NewAuthController(s.aum, s.rg, s.amm)
	authController.Router()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	dataJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(dataJson))
	assert.NoError(s.T(), err)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	authController.findByUsername(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *AuthController) TestRegistuser_Success() {
	s.aum.On("FindByUsername", mockUSerDto).Return(mockTokenJwt, nil)
	s.aum.On("CreateNewUser", mockAuth).Return(mockAuth, nil)
	authController := NewAuthController(s.aum, s.rg, s.amm)
	authController.Router()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/regist", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	authController.createNewUser(ctx)
	assert.Equal(s.T(), http.StatusCreated, record.Code)
}

func (s *AuthController) TestRegistuser_Fail() {
	s.aum.On("FindByUsername", mockUSerDto).Return(mockTokenJwt, nil)
	s.aum.On("CreateNewUser", mockAuth).Return(model.User{}, errors.New("error"))
	authController := NewAuthController(s.aum, s.rg, s.amm)
	authController.Router()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/auth/regist", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwt.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	authController.createNewUser(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}
