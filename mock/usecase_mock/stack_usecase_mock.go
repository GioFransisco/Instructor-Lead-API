package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type StackUseCaseMock struct {
	mock.Mock
}

func (s *StackUseCaseMock) RegisterNewStack(payload dto.StackRequestDto) (model.Stack, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackUseCaseMock) FindAll() ([]model.Stack, error) {
	args := s.Called()
	return args.Get(0).([]model.Stack), args.Error(1)
}

func (s *StackUseCaseMock) FindByID(id string) (model.Stack, error) {
	args := s.Called(id)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackUseCaseMock) UpdateStack(id string, payload model.Stack) (model.Stack, error) {
	args := s.Called(id, payload)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackUseCaseMock) DeleteStack(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
