package usecase

import (
	"fmt"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
)

type AbsencesUseCase interface {
	CreateNewAbsence(payload model.Absences) (dto.AbsencesResponseDto, error)
	FindAbsenceById(scheduleDetailId string) (model.GetAbsences, error)
}

type absencesUseCase struct {
	repo   repository.AbsencesRepository
	SDuc   ScheduleUseCase
	userUC UserUC
}

func (a *absencesUseCase) FindAbsenceById(scheduleDetailId string) (model.GetAbsences, error) {
	absence, err := a.repo.GetScheduleDetailId(scheduleDetailId)
	if err != nil {
		return model.GetAbsences{}, fmt.Errorf("schedule detail id with id %s not found", scheduleDetailId)
	}
	return absence, nil
}

// CreateNewAbsence implements AbsencesUseCase.
func (a *absencesUseCase) CreateNewAbsence(payload model.Absences) (dto.AbsencesResponseDto, error) {
	_, err := a.SDuc.ScheduleDetailFindById(payload.ScheduleDetails.Id)
	if err != nil {
		return dto.AbsencesResponseDto{}, fmt.Errorf("schedule detail with id %s not found", payload.ScheduleDetails.Id)
	}
	user, err := a.userUC.FindById(payload.StudentId.Id)
	if err != nil {
		return dto.AbsencesResponseDto{}, fmt.Errorf("user not found")
	}
	if payload.Description == "" {
		return dto.AbsencesResponseDto{}, fmt.Errorf("description can't be empty")
	}
	payload.StudentId = user

	absences, err := a.repo.Create(payload)
	if err != nil {
		return dto.AbsencesResponseDto{}, err
	}

	payloadDTO := dto.AbsencesResponseDto{
		Id:                absences.Id,
		ScheduleDetailsId: absences.ScheduleDetails.Id,
		StudentId:         absences.StudentId,
		Description:       absences.Description,
		CreatedAt:         absences.CreatedAt,
		UpdatedAt:         absences.UpdatedAt,
	}

	return payloadDTO, nil

}

func NewAbsencesUseCase(repo repository.AbsencesRepository, SDuc ScheduleUseCase, userUC UserUC) AbsencesUseCase {
	return &absencesUseCase{repo: repo, SDuc: SDuc, userUC: userUC}
}
