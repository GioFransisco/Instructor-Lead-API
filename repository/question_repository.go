package repository

import (
	"database/sql"
	"errors"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type QuestionRepository interface {
	CreateQuestion(model.Question) (model.Question, error)
	GetQuestionByID(string) ([]dto.QuestionResponseGET, error)
	UpdateQuestion(dto.QuestionChangeDto) (dto.QuestionResponseUpdate, error)
	UpdateStatusQuestion(dto.QuestionChangeStatusDto) (dto.QuestionResponseUpdate, error)
	DeleteQuestion(string) (dto.QuestionResponseUpdate, error)
}

type questionRepository struct {
	db *sql.DB
}

func NewQusetionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db}
}

// CreateQuestion implements QuestionRepository.
func (r *questionRepository) CreateQuestion(payload model.Question) (model.Question, error) {
	err := r.db.QueryRow(utilsmodel.QuestionCreate, payload.ScheduleDetails.Id, payload.StudentId.Id, payload.Question, payload.Status, time.Now()).Scan(&payload.Id, &payload.CreatedAt, &payload.UpdatedAt)

	if err != nil {
		return model.Question{}, err
	}

	return payload, nil
}

// DeleteQuestion implements QuestionRepository.
func (r *questionRepository) DeleteQuestion(id string) (payloadDto dto.QuestionResponseUpdate, err error) {
	err = r.db.QueryRow(utilsmodel.SelectQuestionById, id).Scan(&payloadDto.Id, &payloadDto.ScheduleDetails, &payloadDto.StudentId, &payloadDto.Question, &payloadDto.Status, &payloadDto.CreatedAt, &payloadDto.UpdatedAt)

	if err != nil {
		return dto.QuestionResponseUpdate{}, errors.New("failed scan, data not found")
	}

	_, err = r.db.Exec(utilsmodel.DeleteQuestion, id)

	return
}

// GetQuestionByID implements QuestionRepository.
func (r *questionRepository) GetQuestionByID(schedulDId string) (payloadQuestion []dto.QuestionResponseGET, err error) {
	rows, err := r.db.Query(utilsmodel.GetQuestionByIDScheduleDetail, schedulDId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		question := dto.QuestionResponseGET{}

		err = rows.Scan(&question.Id, &question.ScheduleDetails, &question.StudentId.Id, &question.Question, &question.Status, &question.CreatedAt, &question.UpdatedAt, &question.StudentId.Name, &question.StudentId.Email, &question.StudentId.PhoneNumber, &question.StudentId.Username, &question.StudentId.Age, &question.StudentId.Address, &question.StudentId.Gander, &question.StudentId.Role, &question.StudentId.CreatedAt, &question.StudentId.UpdatedAt)

		if err != nil {
			return nil, errors.New("failed scan, data not found")
		}

		payloadQuestion = append(payloadQuestion, question)
	}

	return
}

// UpdateQuestion implements QuestionRepository.
func (r *questionRepository) UpdateQuestion(payload dto.QuestionChangeDto) (response dto.QuestionResponseUpdate, err error) {
	err = r.db.QueryRow(utilsmodel.SelectIdQuestionById, payload.Id).Scan(&response.Id)

	if err != nil {
		return dto.QuestionResponseUpdate{}, errors.New("failed scan, data not found")
	}

	err = r.db.QueryRow(utilsmodel.UpdateQuestion, payload.Question, time.Now(), payload.Id).Scan(&response.Id, &response.ScheduleDetails, &response.StudentId, &response.Question, &response.Status, &response.CreatedAt, &response.UpdatedAt)

	return
}

// UpdateStatusQuestion implements QuestionRepository.
func (r *questionRepository) UpdateStatusQuestion(payload dto.QuestionChangeStatusDto) (response dto.QuestionResponseUpdate, err error) {
	err = r.db.QueryRow(utilsmodel.SelectIdQuestionById, payload.Id).Scan(&response.Id)

	if err != nil {
		return dto.QuestionResponseUpdate{}, errors.New("scan value failed. make sure the id you entered is correct")
	}

	err = r.db.QueryRow(utilsmodel.UpdateStatusQuestion, payload.Status, time.Now(), payload.Id).Scan(&response.Id, &response.ScheduleDetails, &response.StudentId, &response.Question, &response.Status, &response.CreatedAt, &response.UpdatedAt)

	return
}
