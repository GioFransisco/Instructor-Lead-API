package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type QuestionMock struct {
	mock.Mock
}

func (q *QuestionMock) CreateQuestion(question model.Question) (model.Question, error) {
	args := q.Called(question)

	return args.Get(0).(model.Question), args.Error(1)
}

func (q *QuestionMock) GetQuestionByID(id string) ([]dto.QuestionResponseGET, error) {
	args := q.Called(id)

	return args.Get(0).([]dto.QuestionResponseGET), args.Error(1)
}

func (q *QuestionMock) UpdateQuestion(dtoQuestion dto.QuestionChangeDto) (dto.QuestionResponseUpdate, error) {
	args := q.Called(dtoQuestion)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}

func (q *QuestionMock) UpdateStatusQuestion(dtoQuestion dto.QuestionChangeStatusDto) (dto.QuestionResponseUpdate, error) {
	args := q.Called(dtoQuestion)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}

func (q *QuestionMock) DeleteQuestion(id string) (dto.QuestionResponseUpdate, error) {
	args := q.Called(id)

	return args.Get(0).(dto.QuestionResponseUpdate), args.Error(1)
}
