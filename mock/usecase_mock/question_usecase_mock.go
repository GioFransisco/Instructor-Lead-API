package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type QuestionUsecaseMock struct {
	mock.Mock
}

func (m *QuestionUsecaseMock) CreateQuestion(questionPayload model.Question) (dto.QuestionResponseGET, error) {
	args := m.Called(questionPayload)

	return args.Get(0).(dto.QuestionResponseGET), args.Error(1)
}

func (m *QuestionUsecaseMock) FindQuestionByID(id string) ([]dto.QuestionResponseGET, error) {
	args := m.Called(id)

	return args.Get(0).([]dto.QuestionResponseGET), args.Error(1)
}

func (m *QuestionUsecaseMock) UpdateQuestion(dtoQuestion dto.QuestionChangeDto) (dto.QuestionResponseUpdate, error) {
	args := m.Called(dtoQuestion)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}

func (m *QuestionUsecaseMock) UpdateStatusQuestion(dtoQuestionStatus dto.QuestionChangeStatusDto) (dto.QuestionResponseUpdate, error) {
	args := m.Called(dtoQuestionStatus)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}

func (m *QuestionUsecaseMock) DeleteQuestion(id string) (dto.QuestionResponseUpdate, error) {
	args := m.Called(id)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}
