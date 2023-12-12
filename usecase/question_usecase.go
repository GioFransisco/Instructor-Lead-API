package usecase

import (
	"errors"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
)

type QuestionUsecase interface {
	CreateQuestion(model.Question) (dto.QuestionResponseGET, error)
	FindQuestionByID(string) ([]dto.QuestionResponseGET, error)
	UpdateQuestion(dto.QuestionChangeDto) (dto.QuestionResponseUpdate, error)
	UpdateStatusQuestion(dto.QuestionChangeStatusDto) (dto.QuestionResponseUpdate, error)
	DeleteQuestion(string) (dto.QuestionResponseUpdate, error)
}

type questionUsecase struct {
	repo   repository.QuestionRepository
	sDUC   ScheduleUseCase
	userUC UserUC
}

// CreateQuestion implements QuestionUsecase.
func (u *questionUsecase) CreateQuestion(payload model.Question) (dto.QuestionResponseGET, error) {
	payload.Status = "Proccess"

	_, err := u.sDUC.ScheduleDetailFindById(payload.ScheduleDetails.Id)

	if err != nil {
		return dto.QuestionResponseGET{}, errors.New("schedule detail not found")
	}

	user, err := u.userUC.FindById(payload.StudentId.Id)

	if err != nil {
		return dto.QuestionResponseGET{}, errors.New("user not found")
	}

	if payload.Question == "" {
		return dto.QuestionResponseGET{}, errors.New("questions cannot be empty")
	}

	payload.StudentId = user

	question, err := u.repo.CreateQuestion(payload)

	if err != nil {
		return dto.QuestionResponseGET{}, err
	}

	dtoQusetion := dto.QuestionResponseGET{
		Id:              question.Id,
		ScheduleDetails: question.ScheduleDetails.Id,
		StudentId:       question.StudentId,
		Question:        question.Question,
		Status:          question.Status,
		CreatedAt:       question.CreatedAt,
		UpdatedAt:       question.UpdatedAt,
	}

	return dtoQusetion, nil
}

// DeleteQuestion implements QuestionUsecase.
func (u *questionUsecase) DeleteQuestion(id string) (dto.QuestionResponseUpdate, error) {
	if id == "" {
		return dto.QuestionResponseUpdate{}, errors.New("id required")
	}

	return u.repo.DeleteQuestion(id)
}

// FindQuestionByID implements QuestionUsecase.
func (u *questionUsecase) FindQuestionByID(sDId string) ([]dto.QuestionResponseGET, error) {
	if _, err := u.sDUC.ScheduleDetailFindById(sDId); err != nil {
		return nil, errors.New("schedule details not found")
	}

	return u.repo.GetQuestionByID(sDId)
}

// UpdateQuestion implements QuestionUsecase.
func (u *questionUsecase) UpdateQuestion(dtoPayload dto.QuestionChangeDto) (dto.QuestionResponseUpdate, error) {
	if dtoPayload.Id == "" || dtoPayload.Question == "" {
		return dto.QuestionResponseUpdate{}, errors.New("id and question required")
	}

	return u.repo.UpdateQuestion(dtoPayload)
}

// UpdateStatusQuestion implements QuestionUsecase.
func (u *questionUsecase) UpdateStatusQuestion(dtoPayload dto.QuestionChangeStatusDto) (dto.QuestionResponseUpdate, error) {
	question := &model.Question{
		Status: dtoPayload.Status,
	}

	if dtoPayload.Id == "" || question.Status == "" {
		return dto.QuestionResponseUpdate{}, errors.New("id and status required")
	}

	if !question.IsValidate() {
		return dto.QuestionResponseUpdate{}, errors.New("question status is not valid")
	}

	return u.repo.UpdateStatusQuestion(dtoPayload)
}

func NewQusetionUsecase(repo repository.QuestionRepository, sDUC ScheduleUseCase, userUC UserUC) QuestionUsecase {
	return &questionUsecase{repo: repo, sDUC: sDUC, userUC: userUC}
}
