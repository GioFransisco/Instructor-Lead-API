package usecase

import (
	"fmt"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
)

type ScheduleApproveUseCase interface {
	CreateNewScheduleAprove(payload model.ScheduleAprove) (dto.ScheduleAproveResponseDto, error)
	FindSchApproveById(schDetailID string) ([]dto.ScheduleAproveResponseDto, error)
}

type scheduleApproveUseCase struct {
	repo repository.ScheduleApproveRepository
	SDuc ScheduleUseCase
}

func (sa *scheduleApproveUseCase) FindSchApproveById(schDetailID string) ([]dto.ScheduleAproveResponseDto, error) {
	_, err := sa.SDuc.ScheduleDetailFindById(schDetailID)
	if err != nil {
		return []dto.ScheduleAproveResponseDto{}, fmt.Errorf("schedule detail with id %s not found", schDetailID)
	}

	return sa.repo.GetApproveById(schDetailID)
}

func (sa *scheduleApproveUseCase) CreateNewScheduleAprove(payload model.ScheduleAprove) (dto.ScheduleAproveResponseDto, error) {
	_, err := sa.SDuc.ScheduleDetailFindById(payload.ScheduleDetails.Id)
	if err != nil {
		return dto.ScheduleAproveResponseDto{}, fmt.Errorf("schedule detail with id %s not found", payload.ScheduleDetails.Id)
	}

	scheduleApprove, err := sa.repo.CreateApprove(payload)
	if err != nil {
		return dto.ScheduleAproveResponseDto{}, err
	}

	scheduleApproveDTO := dto.ScheduleAproveResponseDto{
		Id:                scheduleApprove.Id,
		ScheduleDetailsId: scheduleApprove.ScheduleDetails.Id,
		ScheduleAprove:    scheduleApprove.ScheduleAprove,
		CreatedAt:         scheduleApprove.CreatedAt,
		UpdatedAt:         scheduleApprove.UpdatedAt,
	}

	return scheduleApproveDTO, nil

}

func NewScheduleApproveUseCase(repo repository.ScheduleApproveRepository, SDuc ScheduleUseCase) ScheduleApproveUseCase {
	return &scheduleApproveUseCase{repo: repo, SDuc: SDuc}
}
