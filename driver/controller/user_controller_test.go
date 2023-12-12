package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	middlewaremock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/middleware_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserController struct {
	suite.Suite
	aum *usecasemock.UserUCMock
	rg  *gin.RouterGroup
	amm *middlewaremock.AuthMiddlewareMock
}

func (s *UserController) SetupTest() {
	s.aum = new(usecasemock.UserUCMock)
	rg := gin.Default()
	s.rg = rg.Group("api/v1")
	s.amm = new(middlewaremock.AuthMiddlewareMock)
}

func TestBillUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserController))
}

var mockTokenJwtUser = model.TokenModel{
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJncm91cDEiLCJleHAiOjE3MDAzMjM0MDksImlhdCI6MTcwMDMxNjIwOSwidXNlcklkIjoiMjlmYmEzZTktZDAzYS00ODY0LWFkYTctZDViYjdkMzA4YWFkIiwiZW1haWwiOiJha2JhcnJhdzA5QGdtYWlsLmNvbSIsInJvbGUiOiJBZG1pbiJ9.voTa_GJFgyGbRvthdyDJ1DvLGpUQ1tD6OPJ0ZIOjFro",
}

var updateDto = dto.UserUpdateDto{
	Id:          "kdjksooaieoi",
	Name:        "kdjksooaieoi",
	Email:       "kdjksooaieoi",
	PhoneNumber: "kdjksooaieoi",
	Username:    "kdjksooaieoi",
	Age:         18,
	Address:     "kdjksooaieoild Kol",
	Gender:      "",
}

func (s *UserController) TestChangePaswordUser_Success() {
	s.aum.On("ChangePaswordUser", mockAuth.Password, mockAuth.Id).Return(mockAuth, nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/password", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("userId", mockAuth.Id)
	userController.changePaswordUser(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *UserController) TestChangePaswordUser_Fail() {
	s.aum.On("ChangePaswordUser", mockAuth.Password, mockAuth.Id).Return(mockAuth, nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()
	// Simulasi mengirim sebuah paylaod dalam bentuk JSON
	// mockPayloadJson, err := json.Marshal(mockAuth)
	// assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/password", nil)
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("userId", mockAuth.Id)
	userController.changePaswordUser(ctx)
	assert.Equal(s.T(), http.StatusBadRequest, record.Code)
}

func (s *UserController) TestChangePaswordUser_FailChange() {
	s.aum.On("ChangePaswordUser", mockAuth.Password, mockAuth.Id).Return(model.User{}, errors.New("error"))
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()
	//Simulasi mengirim sebuah paylaod dalam bentuk JSON
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/users/password", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("userId", mockAuth.Id)
	userController.changePaswordUser(ctx)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *UserController) TestUserDelete_Success() {
	s.aum.On("FindById", "").Return(mockAuth, nil)
	s.aum.On("DeleteUserById", "").Return(nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/users/02910wkdwj", nil)
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	userController.deleteUser(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *UserController) TestUserUpdate_Success() {
	s.aum.On("UpdateUser", updateDto).Return(mockAuth, nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()
	mockPayloadJson, err := json.Marshal(mockAuth)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/users", bytes.NewBuffer(mockPayloadJson))
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Set("userId", mockAuth.Id)
	ctx.Request = req
	userController.updateUser(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *UserController) TestFindUserByEmail_Success() {
	s.aum.On("FindUserByEmail", "").Return(mockAuth, nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/akbradhjs", nil)
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	userController.findUserByEmail(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *UserController) TestFindUserId() {
	s.aum.On("FindById", "").Return(mockAuth, nil)
	userController := NewUserController(s.aum, s.rg, s.amm)
	userController.RouterUser()

	record := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/users/akbradhjs", nil)
	assert.NoError(s.T(), err)
	req.Header.Set("Authorization", "Bearer "+mockTokenJwtUser.Token)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	userController.findUserId(ctx)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}
