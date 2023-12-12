package repositorymock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/stretchr/testify/mock"
)

type ScheduleApproveRepoMock struct {
	mock.Mock
}

func (s *ScheduleApproveRepoMock) CreateApprove(payload model.ScheduleAprove) (model.ScheduleAprove, error) {
	args := s.Called(payload)
	return args.Get(0).(model.ScheduleAprove), args.Error(1)
}

func (s *ScheduleApproveRepoMock) GetApproveById(schDetailID string) ([]dto.ScheduleAproveResponseDto, error) {
	args := s.Called(schDetailID)
	return args.Get(0).([]dto.ScheduleAproveResponseDto), args.Error(1)
}
