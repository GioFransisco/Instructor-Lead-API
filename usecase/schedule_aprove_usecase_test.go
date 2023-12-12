package usecase

import (
	"testing"
	"time"

	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScheduleApproveUseCaseTestSuite struct {
	suite.Suite
	scheduleApproveRepoMock *repositorymock.ScheduleApproveRepoMock
	scheduleUseCaseMock     *usecasemock.ScheduleUseCaseMock
	scheduleApproveUseCase  ScheduleApproveUseCase
}

func (suite *ScheduleApproveUseCaseTestSuite) SetupTest() {
	suite.scheduleApproveRepoMock = new(repositorymock.ScheduleApproveRepoMock)
	suite.scheduleUseCaseMock = new(usecasemock.ScheduleUseCaseMock)
	suite.scheduleApproveUseCase = NewScheduleApproveUseCase(suite.scheduleApproveRepoMock, suite.scheduleUseCaseMock)
}

func (suite *ScheduleApproveUseCaseTestSuite) TestFindSchApproveById_Success() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", "1").Return(model.ScheduleDetails{}, nil)
	suite.scheduleApproveRepoMock.On("GetApproveById", "1").Return([]dto.ScheduleAproveResponseDto{}, nil)
	_, err := suite.scheduleApproveUseCase.FindSchApproveById("1")
	assert.Nil(suite.T(), err)
}

func (suite *ScheduleApproveUseCaseTestSuite) TestCreate_Success() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", "1").Return(model.ScheduleDetails{}, nil)

	payload := model.ScheduleAprove{
		ScheduleDetails: model.ScheduleDetails{
			Id: "1",
		},
		ScheduleAprove: "/test.png",
		UpdatedAt:      time.Now(),
	}

	suite.scheduleApproveRepoMock.On("CreateApprove", payload).Return(payload, nil)
	_, err := suite.scheduleApproveUseCase.CreateNewScheduleAprove(payload)
	assert.Nil(suite.T(), err)
}

func TestScheduleApproveUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleApproveUseCaseTestSuite))
}
