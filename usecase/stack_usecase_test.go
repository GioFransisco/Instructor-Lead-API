package usecase

import (
	"errors"
	"testing"
	"time"

	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StackUseCaseTestSuite struct {
	suite.Suite
	stackRepoMock    *repositorymock.StackRepoMock
	stackUseCaseMock *usecasemock.StackUseCaseMock
	stackUseCase     StackUseCase
}

func (suite *StackUseCaseTestSuite) SetupTest() {
	suite.stackRepoMock = new(repositorymock.StackRepoMock)
	suite.stackUseCaseMock = new(usecasemock.StackUseCaseMock)
	suite.stackUseCase = NewStackUseCase(suite.stackRepoMock)
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

func (suite *StackUseCaseTestSuite) TestRegisterNewStack_Success() {
	var stack model.Stack
	stack.Name = mockPayloadStack.Name

	suite.stackRepoMock.On("Create", stack).Return(mockStack, nil)
	_, err := suite.stackUseCase.RegisterNewStack(mockPayloadStack)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestFindAll_Success() {
	var mockStacks = []model.Stack{
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

	suite.stackRepoMock.On("List").Return(mockStacks, nil)
	_, err := suite.stackUseCase.FindAll()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestFindAll_Fail() {
	suite.stackRepoMock.On("List").Return([]model.Stack{}, errors.New("error"))
	_, err := suite.stackUseCase.FindAll()
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestFindByID_Success() {
	suite.stackRepoMock.On("FindByID", mockStack.Id).Return(mockStack, nil)
	_, err := suite.stackUseCase.FindByID(mockStack.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestFindByID_Fail_NotFound() {
	suite.stackRepoMock.On("FindByID", mockStack.Id).Return(model.Stack{}, errors.New("uuid"))
	_, err := suite.stackUseCase.FindByID(mockStack.Id)
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestFindByID_Fail() {
	suite.stackRepoMock.On("FindByID", mockStack.Id).Return(model.Stack{}, errors.New("errors"))
	_, err := suite.stackUseCase.FindByID(mockStack.Id)
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestUpdate_Success() {
	var stack model.Stack
	stack.Name = mockPayloadStack.Name
	stack.Status = mockPayloadStack.Status

	suite.stackRepoMock.On("Update", mockStack.Id, stack).Return(mockStack, nil)
	_, err := suite.stackUseCase.UpdateStack(mockStack.Id, stack)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestUpdate_Success_Fail() {
	var stack model.Stack
	stack.Name = mockPayloadStack.Name
	stack.Status = "S"

	suite.stackRepoMock.On("Update", mockStack.Id, stack).Return(model.Stack{}, common.InvalidError{Message: "error"})
	_, err := suite.stackUseCase.UpdateStack(mockStack.Id, stack)
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestUpdate_Success_Fail_NotFound() {
	var stack model.Stack
	stack.Name = mockPayloadStack.Name
	stack.Status = mockPayloadStack.Status

	suite.stackRepoMock.On("Update", mockStack.Id, stack).Return(model.Stack{}, common.InvalidError{Message: "uuid"})
	_, err := suite.stackUseCase.UpdateStack(mockStack.Id, stack)
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestUpdate_Success_Fail_500() {
	var stack model.Stack
	stack.Name = mockPayloadStack.Name
	stack.Status = mockPayloadStack.Status

	suite.stackRepoMock.On("Update", mockStack.Id, stack).Return(model.Stack{}, errors.New("error"))
	_, err := suite.stackUseCase.UpdateStack(mockStack.Id, stack)
	assert.Error(suite.T(), err)
}

func (suite *StackUseCaseTestSuite) TestDelete_success() {
	suite.stackRepoMock.On("Delete", mockStack.Id).Return(nil)
	err := suite.stackUseCase.DeleteStack(mockStack.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func TestStackUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(StackUseCaseTestSuite))
}
