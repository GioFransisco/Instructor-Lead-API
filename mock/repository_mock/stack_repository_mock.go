package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/mock"
)

type StackRepoMock struct {
	mock.Mock
}

func (s *StackRepoMock) Create(payload model.Stack) (model.Stack, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackRepoMock) Get(id string) (model.Stack, error) {
	args := s.Called(id)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackRepoMock) List() ([]model.Stack, error) {
	args := s.Called()
	return args.Get(0).([]model.Stack), args.Error(1)
}

func (s *StackRepoMock) FindByID(id string) (model.Stack, error) {
	args := s.Called(id)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackRepoMock) Update(id string, payload model.Stack) (model.Stack, error) {
	args := s.Called(id, payload)
	return args.Get(0).(model.Stack), args.Error(1)
}

func (s *StackRepoMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
