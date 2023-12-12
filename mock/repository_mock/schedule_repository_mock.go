package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/stretchr/testify/mock"
)

type ScheduleRepoMock struct {
	mock.Mock
}

func (s *ScheduleRepoMock) Create(payload model.Schedule) (model.Schedule, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Schedule), args.Error(1)
}

func (s *ScheduleRepoMock) List(userId, userRole string) ([]model.Schedule, error) {
	args := s.Called(userId, userRole)
	return args.Get(0).([]model.Schedule), args.Error(1)
}

func (s *ScheduleRepoMock) Get(id string) (model.Schedule, error) {
	args := s.Called(id)
	return args.Get(0).(model.Schedule), args.Error(1)
}

func (s *ScheduleRepoMock) GetSchedule(id string) (model.Schedule, error) {
	args := s.Called(id)
	return args.Get(0).(model.Schedule), args.Error(1)
}

func (s *ScheduleRepoMock) GetScheduleDetail(id string) (model.ScheduleDetails, error) {
	args := s.Called(id)
	return args.Get(0).(model.ScheduleDetails), args.Error(1)
}

func (s *ScheduleRepoMock) UpdateSchedule(payload model.Schedule) (model.Schedule, error) {
	args := s.Called(payload)
	return args.Get(0).(model.Schedule), args.Error(1)
}

func (s *ScheduleRepoMock) UpdateScheduleDetail(payload model.ScheduleDetails) (model.ScheduleDetails, error) {
	args := s.Called(payload)
	return args.Get(0).(model.ScheduleDetails), args.Error(1)
}
