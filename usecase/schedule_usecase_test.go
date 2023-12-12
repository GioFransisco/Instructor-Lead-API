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

type ScheduleUseCaseTestSuite struct {
	suite.Suite
	scheduleRepoMock    *repositorymock.ScheduleRepoMock
	scheduleUseCaseMock *usecasemock.ScheduleUseCaseMock
	trainerUseCaseMock  *usecasemock.UserUCMock
	stackUseCaseMock    *usecasemock.StackUseCaseMock
	scheduleUseCase     ScheduleUseCase
}

var mockSchedule = model.Schedule{
	Id:           "1",
	Name:         "Instructor Led test name",
	DateActivity: time.Now(),
	ScheduleDetails: []model.ScheduleDetails{
		{
			Id:         "1",
			ScheduleId: "1",
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
			StartTime: time.Now(),
			EndTime:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:         "2",
			ScheduleId: "1",
			Trainer: model.User{
				Id:          "2",
				Name:        "Kira",
				Email:       "kira@email.com",
				PhoneNumber: "089768758272",
				Username:    "kira",
				Age:         23,
				Address:     "Garut",
				Gander:      "P",
				Role:        "Trainer",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			Stack: model.Stack{
				Id:        "2",
				Name:      "Java",
				Status:    "Active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartTime: time.Now(),
			EndTime:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayloadSchedule = dto.ScheduleCreateRequestDto{
	Name:         "Instructor Led test name",
	DateActivity: "2023-11-20",
	ScheduleDetails: []dto.ScheduleDetailCreateRequestDto{
		{
			TrainerId: "1",
			StackId:   "1",
			StartTime: "19:00",
			EndTime:   "20:00",
		},
		{
			TrainerId: "2",
			StackId:   "2",
			StartTime: "19:00",
			EndTime:   "20:00",
		},
	},
}

func (suite *ScheduleUseCaseTestSuite) SetupTest() {
	suite.scheduleRepoMock = new(repositorymock.ScheduleRepoMock)
	suite.scheduleUseCaseMock = new(usecasemock.ScheduleUseCaseMock)
	suite.trainerUseCaseMock = new(usecasemock.UserUCMock)
	suite.stackUseCaseMock = new(usecasemock.StackUseCaseMock)
	suite.scheduleUseCase = NewScheduleUseCase(suite.scheduleRepoMock, suite.trainerUseCaseMock, suite.stackUseCaseMock)
}

func (suite *ScheduleUseCaseTestSuite) TestRegisterNewSchedule_Success() {
	var newSchedule model.Schedule
	var newScheduleDetails []model.ScheduleDetails

	dateActivity, _ := time.Parse("2006-01-02", mockPayloadSchedule.DateActivity)

	newSchedule.Name = mockPayloadSchedule.Name
	newSchedule.DateActivity = dateActivity

	for i, v := range mockPayloadSchedule.ScheduleDetails {
		var scheduleDetail model.ScheduleDetails

		suite.trainerUseCaseMock.On("FindById", v.TrainerId).Return(mockSchedule.ScheduleDetails[i].Trainer, nil)
		trainer, _ := suite.trainerUseCaseMock.FindById(v.TrainerId)

		suite.stackUseCaseMock.On("FindByID", v.StackId).Return(mockSchedule.ScheduleDetails[i].Stack, nil)
		stack, _ := suite.stackUseCaseMock.FindByID(v.StackId)

		startTime, _ := time.Parse("15:04", v.StartTime)
		endTime, _ := time.Parse("15:04", v.EndTime)

		scheduleDetail.Trainer = trainer
		scheduleDetail.Stack = stack
		scheduleDetail.StartTime = startTime
		scheduleDetail.EndTime = endTime

		newScheduleDetails = append(newScheduleDetails, scheduleDetail)
	}

	newSchedule.ScheduleDetails = newScheduleDetails

	suite.scheduleRepoMock.On("Create", newSchedule).Return(mockSchedule, nil)
	actual, err := suite.scheduleUseCase.RegisterNewSchedule(mockPayloadSchedule)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleUseCaseTestSuite) TestFindAllSchedules_Success() {
	mockSchedules := []model.Schedule{
		mockSchedule,
	}

	suite.scheduleRepoMock.On("List", mockSchedule.ScheduleDetails[0].Trainer.Id, mockSchedule.ScheduleDetails[0].Trainer.Role).Return(mockSchedules, nil)
	actual, err := suite.scheduleUseCase.FindAll(mockSchedule.ScheduleDetails[0].Trainer.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual[0].Id)
}

func (suite *ScheduleUseCaseTestSuite) TestFindByIdSchedule_Success() {
	suite.scheduleRepoMock.On("GetScheduleDetail", mockSchedule.ScheduleDetails[0].Id).Return(mockSchedule.ScheduleDetails[0], nil)
	actual, err := suite.scheduleUseCase.ScheduleDetailFindById(mockSchedule.ScheduleDetails[0].Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleUseCaseTestSuite) TestFindByIdScheduleDetail_Success() {
	suite.scheduleRepoMock.On("Get", mockSchedule.Id).Return(mockSchedule, nil)
	actual, err := suite.scheduleUseCase.FindById(mockSchedule.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleUseCaseTestSuite) TestUpdateSchedule_Success() {
	mockPayloadUpdateSchedule := dto.ScheduleUpdateRequestDto{
		Name:         "Tanggal 20",
		DateActivity: "2023-11-18",
	}

	dateActivity, err := time.Parse("2006-01-02", mockPayloadUpdateSchedule.DateActivity)
	assert.NoError(suite.T(), err)

	mockSchedule := model.Schedule{
		Id:           "1",
		Name:         mockPayloadUpdateSchedule.Name,
		DateActivity: dateActivity,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	suite.scheduleRepoMock.On("GetSchedule", mockSchedule.Id).Return(mockSchedule, nil)
	updateSchedule, err := suite.scheduleRepoMock.GetSchedule(mockSchedule.Id)
	assert.NoError(suite.T(), err)

	updateSchedule.Name = mockPayloadUpdateSchedule.Name
	updateSchedule.DateActivity = mockSchedule.DateActivity

	suite.scheduleRepoMock.On("UpdateSchedule", updateSchedule).Return(mockSchedule, nil)
	actual, err := suite.scheduleUseCase.UpdateSchedule(mockSchedule.Id, mockPayloadUpdateSchedule)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleUseCaseTestSuite) TestUpdateScheduleDetail_Success() {
	mockPayload := dto.ScheduleDetailUpdateRequestDto{
		TrainerId: "1",
		StackId:   "1",
		StartTime: "19:00",
		EndTime:   "20:00",
	}

	startTime, _ := time.Parse("15:04", mockPayload.StartTime)
	endTime, _ := time.Parse("15:04", mockPayload.EndTime)

	mockScheduleDetail := model.ScheduleDetails{
		Id:         "1",
		ScheduleId: mockSchedule.Id,
		Trainer:    mockSchedule.ScheduleDetails[0].Trainer,
		Stack:      mockSchedule.ScheduleDetails[0].Stack,
		StartTime:  startTime,
		EndTime:    endTime,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	suite.scheduleRepoMock.On("GetScheduleDetail", mockScheduleDetail.Id).Return(mockScheduleDetail, nil)
	updateScheduleDetail, err := suite.scheduleRepoMock.GetScheduleDetail(mockScheduleDetail.Id)
	assert.NoError(suite.T(), err)

	suite.scheduleRepoMock.On("UpdateScheduleDetail", updateScheduleDetail).Return(mockScheduleDetail, nil)
	suite.trainerUseCaseMock.On("FindById", mockScheduleDetail.Trainer.Id).Return(mockScheduleDetail.Trainer, nil)
	suite.stackUseCaseMock.On("FindByID", mockScheduleDetail.Stack.Id).Return(mockScheduleDetail.Stack, nil)

	actual, err := suite.scheduleUseCase.UpdateScheduleDetail(mockScheduleDetail.Id, mockPayload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockScheduleDetail.Id, actual.Id)
}

func TestScheduleUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleUseCaseTestSuite))
}
