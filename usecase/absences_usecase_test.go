package usecase

import (
	"errors"
	"testing"
	"time"

	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AbsencesUseCaseTestSuite struct {
	suite.Suite
	absencesRepoMock    *repositorymock.AbsencesRepoMock
	absencesUseCaseMock *usecasemock.AbsencesUseCaseMock
	scheduleUseCaseMock *usecasemock.ScheduleUseCaseMock
	userUseCaseMock     *usecasemock.UserUCMock
	absencesUseCase     AbsencesUseCase
}

func (suite *AbsencesUseCaseTestSuite) SetupTest() {
	suite.absencesRepoMock = new(repositorymock.AbsencesRepoMock)
	suite.absencesUseCaseMock = new(usecasemock.AbsencesUseCaseMock)
	suite.scheduleUseCaseMock = new(usecasemock.ScheduleUseCaseMock)
	suite.userUseCaseMock = new(usecasemock.UserUCMock)
	suite.absencesUseCase = NewAbsencesUseCase(suite.absencesRepoMock, suite.scheduleUseCaseMock, suite.userUseCaseMock)
}

var mockAbsences = model.Absences{
	Id: "1",
	ScheduleDetails: model.ScheduleDetails{
		Id: "1",
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
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
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

func (suite *AbsencesUseCaseTestSuite) TestCreateAbsences_Success() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", mockAbsences.ScheduleDetails.Id).Return(mockAbsences.ScheduleDetails, nil)
	suite.userUseCaseMock.On("FindById", mockAbsences.StudentId.Id).Return(mockAbsences.StudentId, nil)
	suite.absencesRepoMock.On("Create", mockAbsences).Return(mockAbsences, nil)

	actual, err := suite.absencesUseCase.CreateNewAbsence(mockAbsences)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), mockAbsences.Id, actual.Id)
}

func (suite *AbsencesUseCaseTestSuite) TestCreateAbsences_Fail_OnScheduleDetailFindById() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", mockAbsences.ScheduleDetails.Id).Return(model.ScheduleDetails{}, errors.New("error"))

	_, err := suite.absencesUseCase.CreateNewAbsence(mockAbsences)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesUseCaseTestSuite) TestCreateAbsences_Fail_OnUserFindById() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", mockAbsences.ScheduleDetails.Id).Return(mockAbsences.ScheduleDetails, nil)
	suite.userUseCaseMock.On("FindById", mockAbsences.StudentId.Id).Return(model.User{}, errors.New("error"))

	_, err := suite.absencesUseCase.CreateNewAbsence(mockAbsences)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesUseCaseTestSuite) TestCreateAbsences_Fail_OnEmptyDescription() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", mockAbsences.ScheduleDetails.Id).Return(mockAbsences.ScheduleDetails, nil)
	suite.userUseCaseMock.On("FindById", mockAbsences.StudentId.Id).Return(mockAbsences.StudentId, nil)

	mockPayloadAbsences := mockAbsences
	mockPayloadAbsences.Description = ""
	suite.absencesRepoMock.On("Create", mockPayloadAbsences).Return(model.Absences{}, errors.New("error"))

	_, err := suite.absencesUseCase.CreateNewAbsence(mockPayloadAbsences)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesUseCaseTestSuite) TestCreateAbsences_Fail_OnCreate() {
	suite.scheduleUseCaseMock.On("ScheduleDetailFindById", mockAbsences.ScheduleDetails.Id).Return(mockAbsences.ScheduleDetails, nil)
	suite.userUseCaseMock.On("FindById", mockAbsences.StudentId.Id).Return(mockAbsences.StudentId, nil)
	suite.absencesRepoMock.On("Create", mockAbsences).Return(model.Absences{}, errors.New("error"))

	_, err := suite.absencesUseCase.CreateNewAbsence(mockAbsences)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesUseCaseTestSuite) TestFindAbsenceById_Success() {
	suite.absencesRepoMock.On("GetScheduleDetailId", mockGetAbsences.ScheduleDetails[0].Id).Return(mockGetAbsences, nil)
	actual, err := suite.absencesUseCase.FindAbsenceById(mockGetAbsences.ScheduleDetails[0].Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), mockGetAbsences.ScheduleDetails[0].Id, actual.Id)
}

func (suite *AbsencesUseCaseTestSuite) TestFindAbsenceById_Fail() {
	suite.absencesRepoMock.On("GetScheduleDetailId", mockGetAbsences.ScheduleDetails[0].Id).Return(model.GetAbsences{}, errors.New("error"))
	_, err := suite.absencesUseCase.FindAbsenceById(mockGetAbsences.ScheduleDetails[0].Id)
	assert.Error(suite.T(), err)
}

func TestAbsenccesUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AbsencesUseCaseTestSuite))
}
