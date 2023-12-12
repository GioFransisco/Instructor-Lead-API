package usecase

import (
	"fmt"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
)

type ScheduleUseCase interface {
	RegisterNewSchedule(payload dto.ScheduleCreateRequestDto) (dto.ScheduleResponseDto, error)
	FindAll(userId, userRole string) ([]dto.ScheduleResponseDto, error)
	FindById(id string) (model.Schedule, error)
	ScheduleDetailFindById(id string) (model.ScheduleDetails, error)
	UpdateSchedule(id string, payload dto.ScheduleUpdateRequestDto) (dto.ScheduleResponseDto, error)
	UpdateScheduleDetail(id string, payload dto.ScheduleDetailUpdateRequestDto) (dto.ScheduleDetailResponseDto, error)
}

type scheduleUseCase struct {
	repo      repository.ScheduleRepository
	trainerUC UserUC
	stackUC   StackUseCase
}

// ScheduleDetailFindById implements ScheduleUseCase.
func (s *scheduleUseCase) ScheduleDetailFindById(id string) (model.ScheduleDetails, error) {
	return s.repo.GetScheduleDetail(id)
}

func (s *scheduleUseCase) RegisterNewSchedule(payload dto.ScheduleCreateRequestDto) (dto.ScheduleResponseDto, error) {
	var newSchedule model.Schedule
	var newScheduleDetails []model.ScheduleDetails

	dateActivity, err := time.Parse("2006-01-02", payload.DateActivity)
	if err != nil {
		return dto.ScheduleResponseDto{}, common.InvalidError{Message: "invalid format date, make sure dateActivity using format YYYY-MM-DD"}
	}

	newSchedule.Name = payload.Name
	newSchedule.DateActivity = dateActivity

	if len(payload.ScheduleDetails) < 1 {
		return dto.ScheduleResponseDto{}, common.InvalidError{Message: "scheduleDetails requires at least one piece of data"}
	}

	for _, v := range payload.ScheduleDetails {
		var scheduleDetail model.ScheduleDetails

		trainer, err := s.trainerUC.FindById(v.TrainerId)
		if err != nil {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: err.Error()}
		}

		if trainer.Role != "Trainer" {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: fmt.Sprintf("invalid trainer, make sure the user with ID  %s is trainer", v.TrainerId)}
		}

		stack, err := s.stackUC.FindByID(v.StackId)
		if err != nil {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: err.Error()}
		}

		if stack.Status == "Inactive" {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: fmt.Sprintf("make sure the stack with ID %s is active", v.StackId)}
		}

		startTime, err := time.Parse("15:04", v.StartTime)
		if err != nil {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: "invalid format date, make sure startTime using format HH:MM"}
		}

		endTime, err := time.Parse("15:04", v.EndTime)
		if err != nil {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: "invalid format date, make sure endTime using format HH:MM"}
		}

		scheduleDetail.Trainer = trainer
		scheduleDetail.Stack = stack
		scheduleDetail.StartTime = startTime
		scheduleDetail.EndTime = endTime

		newScheduleDetails = append(newScheduleDetails, scheduleDetail)
	}

	newSchedule.ScheduleDetails = newScheduleDetails

	schedule, err := s.repo.Create(newSchedule)

	var responseSchedule dto.ScheduleResponseDto
	responseSchedule.Id = schedule.Id
	responseSchedule.Name = schedule.Name
	responseSchedule.DateActivity = schedule.DateActivity.Format("2006-01-02")
	responseSchedule.CreatedAt = schedule.CreatedAt
	responseSchedule.UpdatedAt = schedule.UpdatedAt

	var scheduleDetails []dto.ScheduleDetailResponseDto
	for _, v := range schedule.ScheduleDetails {
		var scheduleDetail dto.ScheduleDetailResponseDto
		scheduleDetail.Id = v.Id
		scheduleDetail.Trainer = v.Trainer
		scheduleDetail.Stack = v.Stack
		scheduleDetail.StartTime = v.StartTime.Format("15:04")
		scheduleDetail.EndTime = v.EndTime.Format("15:04")
		scheduleDetail.CreatedAt = v.CreatedAt
		scheduleDetail.UpdatedAt = v.UpdatedAt

		scheduleDetails = append(scheduleDetails, scheduleDetail)
	}

	responseSchedule.ScheduleDetails = scheduleDetails

	return responseSchedule, nil
}

func (s *scheduleUseCase) FindAll(userId, userRole string) ([]dto.ScheduleResponseDto, error) {
	var scheduleResponse []dto.ScheduleResponseDto

	schedules, err := s.repo.List(userId, userRole)
	if err != nil {
		return []dto.ScheduleResponseDto{}, err
	}

	if len(schedules) < 1 {
		return []dto.ScheduleResponseDto{}, common.InvalidError{Message: "data not found"}
	}

	for _, vSchedule := range schedules {
		var schedule dto.ScheduleResponseDto

		schedule.Id = vSchedule.Id
		schedule.Name = vSchedule.Name
		schedule.DateActivity = vSchedule.DateActivity.Format("2006-01-02")
		schedule.CreatedAt = vSchedule.CreatedAt
		schedule.UpdatedAt = vSchedule.UpdatedAt

		var scheduleDetails []dto.ScheduleDetailResponseDto
		for _, vScheduleDetail := range vSchedule.ScheduleDetails {
			var scheduleDetail dto.ScheduleDetailResponseDto

			scheduleDetail.Id = vScheduleDetail.Id
			scheduleDetail.Trainer = vScheduleDetail.Trainer
			scheduleDetail.Stack = vScheduleDetail.Stack
			scheduleDetail.StartTime = vScheduleDetail.StartTime.Format("15:04")
			scheduleDetail.EndTime = vScheduleDetail.EndTime.Format("15:04")
			scheduleDetail.CreatedAt = vScheduleDetail.CreatedAt
			scheduleDetail.UpdatedAt = vScheduleDetail.UpdatedAt

			scheduleDetails = append(scheduleDetails, scheduleDetail)
		}

		schedule.ScheduleDetails = scheduleDetails
		scheduleResponse = append(scheduleResponse, schedule)
	}

	return scheduleResponse, nil
}

func (s *scheduleUseCase) FindById(id string) (model.Schedule, error) {
	schedule, err := s.repo.Get(id)
	if err != nil {
		return model.Schedule{}, fmt.Errorf("schedule with ID %s not found", id)
	}

	return schedule, nil
}

func (s *scheduleUseCase) UpdateSchedule(id string, payload dto.ScheduleUpdateRequestDto) (dto.ScheduleResponseDto, error) {
	updateSchedule, err := s.repo.GetSchedule(id)
	if err != nil {
		return dto.ScheduleResponseDto{}, common.InvalidError{Message: fmt.Sprintf("schedule with ID %s not found", id)}
	}

	var dateActivity time.Time
	if payload.DateActivity != "" {
		dateActivity, err = time.Parse("2006-01-02", payload.DateActivity)
		if err != nil {
			return dto.ScheduleResponseDto{}, common.InvalidError{Message: "invalid format date, make sure dateActivity using format YYYY-MM-DD"}
		}
	}

	updateSchedule.Name = payload.Name
	updateSchedule.DateActivity = dateActivity

	var responseSchedule dto.ScheduleResponseDto
	schedule, err := s.repo.UpdateSchedule(updateSchedule)
	if err != nil {
		return dto.ScheduleResponseDto{}, err
	}

	responseSchedule.Id = schedule.Id
	responseSchedule.Name = schedule.Name
	responseSchedule.DateActivity = schedule.DateActivity.Format("2006-01-02")
	responseSchedule.CreatedAt = schedule.CreatedAt

	return responseSchedule, nil
}

func (s *scheduleUseCase) UpdateScheduleDetail(id string, payload dto.ScheduleDetailUpdateRequestDto) (dto.ScheduleDetailResponseDto, error) {
	var trainer model.User
	var stack model.Stack

	updateScheduleDetail, err := s.repo.GetScheduleDetail(id)
	if err != nil {
		return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: fmt.Sprintf("schedule detail with ID %s not found", id)}
	}

	if payload.TrainerId != "" {
		trainer, err = s.trainerUC.FindById(payload.TrainerId)
		if err != nil {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: fmt.Sprintf("trainer with ID %s not found", payload.TrainerId)}
		}

		if trainer.Role != "Trainer" {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: fmt.Sprintf("invalid trainer, make sure the user with ID %s is trainer", payload.TrainerId)}
		}
	}

	if payload.StackId != "" {
		stack, err = s.stackUC.FindByID(payload.StackId)
		if err != nil {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: fmt.Sprintf("stack with ID %s not found", payload.StackId)}
		}

		if stack.Status == "Inactive" {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: fmt.Sprintf("make sure the stack with ID %s is active", payload.StackId)}
		}
	}

	var startTime time.Time
	var endTime time.Time

	if payload.StartTime != "" {
		startTime, err = time.Parse("15:04", payload.StartTime)
		if err != nil {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: "invalid format date, make sure startTime using format HH:MM"}
		}
	}

	if payload.EndTime != "" {
		endTime, err = time.Parse("15:04", payload.EndTime)
		if err != nil {
			return dto.ScheduleDetailResponseDto{}, common.InvalidError{Message: "invalid format date, make sure endTime using format HH:MM"}
		}
	}

	updateScheduleDetail.Trainer.Id = payload.TrainerId
	updateScheduleDetail.Stack.Id = payload.StackId
	updateScheduleDetail.StartTime = startTime
	updateScheduleDetail.EndTime = endTime

	scheduleDetail, err := s.repo.UpdateScheduleDetail(updateScheduleDetail)

	trainer, _ = s.trainerUC.FindById(scheduleDetail.Trainer.Id)
	stack, _ = s.stackUC.FindByID(scheduleDetail.Stack.Id)

	var scheduleDetailResponse dto.ScheduleDetailResponseDto
	scheduleDetailResponse.Id = scheduleDetail.Id
	scheduleDetailResponse.Trainer = trainer
	scheduleDetailResponse.Stack = stack
	scheduleDetailResponse.StartTime = scheduleDetail.StartTime.Format("15:04")
	scheduleDetailResponse.EndTime = scheduleDetail.EndTime.Format("15:04")
	scheduleDetailResponse.CreatedAt = scheduleDetail.CreatedAt
	scheduleDetailResponse.UpdatedAt = scheduleDetail.UpdatedAt

	return scheduleDetailResponse, nil
}

func NewScheduleUseCase(repo repository.ScheduleRepository, trainerUC UserUC, stackUC StackUseCase) ScheduleUseCase {
	return &scheduleUseCase{repo: repo, trainerUC: trainerUC, stackUC: stackUC}
}
