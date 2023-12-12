package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type ScheduleUseCaseMock struct {
	mock.Mock
}

func (s *ScheduleUseCaseMock) RegisterNewSchedule(payload dto.ScheduleCreateRequestDto) (dto.ScheduleResponseDto, error) {
	args := s.Called(payload)
	return args.Get(0).(dto.ScheduleResponseDto), args.Error(1)
}

func (s *ScheduleUseCaseMock) FindAll(userId, userRole string) ([]dto.ScheduleResponseDto, error) {
	args := s.Called(userId, userRole)
	return args.Get(0).([]dto.ScheduleResponseDto), args.Error(1)
}

func (s *ScheduleUseCaseMock) FindById(id string) (model.Schedule, error) {
	args := s.Called(id)
	return args.Get(0).(model.Schedule), args.Error(1)
}

func (s *ScheduleUseCaseMock) ScheduleDetailFindById(id string) (model.ScheduleDetails, error) {
	args := s.Called(id)

	return args.Get(0).(model.ScheduleDetails), args.Error(1)
}

func (s *ScheduleUseCaseMock) UpdateSchedule(id string, payload dto.ScheduleUpdateRequestDto) (dto.ScheduleResponseDto, error) {
	args := s.Called(id, payload)
	return args.Get(0).(dto.ScheduleResponseDto), args.Error(1)
}

func (s *ScheduleUseCaseMock) UpdateScheduleDetail(id string, payload dto.ScheduleDetailUpdateRequestDto) (dto.ScheduleDetailResponseDto, error) {
	args := s.Called(id, payload)
	return args.Get(0).(dto.ScheduleDetailResponseDto), args.Error(1)
}
