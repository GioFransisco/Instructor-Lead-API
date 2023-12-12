package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type ScheduleApproveUseCaseMock struct {
	mock.Mock
}

func (s *ScheduleApproveUseCaseMock) CreateNewScheduleAprove(payload model.ScheduleAprove) (dto.ScheduleAproveResponseDto, error) {
	args := s.Called(payload)
	return args.Get(0).(dto.ScheduleAproveResponseDto), args.Error(1)
}

func (s *ScheduleApproveUseCaseMock) FindSchApproveById(schDetailID string) ([]dto.ScheduleAproveResponseDto, error) {
	args := s.Called(schDetailID)
	return args.Get(0).([]dto.ScheduleAproveResponseDto), args.Error(1)
}
