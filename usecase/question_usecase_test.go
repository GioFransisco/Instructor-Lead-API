package usecase

import (
	"errors"
	"testing"
	"time"

	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QuestionUsecaseTestSuite struct {
	suite.Suite
	repoQuestion *repositorymock.QuestionMock
	schMock      *usecasemock.ScheduleUseCaseMock
	userMock     *usecasemock.UserUCMock
	questionUc   QuestionUsecase
}

func (s *QuestionUsecaseTestSuite) SetupTest() {
	s.repoQuestion = new(repositorymock.QuestionMock)
	s.userMock = new(usecasemock.UserUCMock)
	s.schMock = new(usecasemock.ScheduleUseCaseMock)
	s.questionUc = NewQusetionUsecase(s.repoQuestion, s.schMock, s.userMock)
}

func TestQuestionUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionUsecaseTestSuite))
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

var dtoMockResponseGet = []dto.QuestionResponseGET{
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

func (s *QuestionUsecaseTestSuite) TestCreateQuestion_Success() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(mockQuestion.ScheduleDetails, nil)

	s.userMock.On("FindById", mockQuestion.StudentId.Id).Return(mockQuestion.StudentId, nil)

	s.repoQuestion.On("CreateQuestion", mockQuestion).Return(mockQuestion, nil)

	actual, err := s.questionUc.CreateQuestion(mockQuestion)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionUsecaseTestSuite) TestCreateQuestion_Fail() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(model.ScheduleDetails{}, errors.New("error"))

	_, err := s.questionUc.CreateQuestion(mockQuestion)
	assert.Error(s.T(), err)
}

func (s *QuestionUsecaseTestSuite) TestCreateQuestionFindUser_Fail() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(mockQuestion.ScheduleDetails, nil)

	s.userMock.On("FindById", mockQuestion.StudentId.Id).Return(model.User{}, errors.New("error"))

	_, err := s.questionUc.CreateQuestion(mockQuestion)
	assert.Error(s.T(), err)
}

func (s *QuestionUsecaseTestSuite) TestCreateQuestionEmpty_Fail() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(mockQuestion.ScheduleDetails, nil)

	s.userMock.On("FindById", mockQuestion.StudentId.Id).Return(mockQuestion.StudentId, nil)

	s.repoQuestion.On("CreateQuestion", mockQuestion).Return(model.Question{}, errors.New("error"))

	_, err := s.questionUc.CreateQuestion(mockQuestion)
	assert.Error(s.T(), err)
}

func (s *QuestionUsecaseTestSuite) TestDeleteQuestion_Success() {
	s.repoQuestion.On("DeleteQuestion", mockQuestion.Id).Return(dtoMockResponse, nil)

	actual, err := s.questionUc.DeleteQuestion(mockQuestion.Id)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionUsecaseTestSuite) TestDeleteQuestion_Fail() {
	s.repoQuestion.On("DeleteQuestion", mockQuestion.Id).Return(dtoMockResponse, nil)

	_, err := s.questionUc.DeleteQuestion("")
	assert.Error(s.T(), err)
}

func (s *QuestionUsecaseTestSuite) TestFindQuestionByID_Success() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(mockQuestion.ScheduleDetails, nil)

	s.repoQuestion.On("GetQuestionByID", mockQuestion.ScheduleDetails.Id).Return(dtoMockResponseGet, nil)

	actual, err := s.questionUc.FindQuestionByID(mockQuestion.ScheduleDetails.Id)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual[0].Id)
}

func (s *QuestionUsecaseTestSuite) TestFindQuestionByID_Fail() {
	s.schMock.On("ScheduleDetailFindById", mockQuestion.ScheduleDetails.Id).Return(model.ScheduleDetails{}, errors.New("error"))

	s.repoQuestion.On("GetQuestionByID", mockQuestion.ScheduleDetails.Id).Return(dtoMockResponseGet, nil)

	_, err := s.questionUc.FindQuestionByID(mockQuestion.ScheduleDetails.Id)
	assert.Error(s.T(), err)
}

func (s *QuestionUsecaseTestSuite) TestUpdateStatusQuestion() {

}
