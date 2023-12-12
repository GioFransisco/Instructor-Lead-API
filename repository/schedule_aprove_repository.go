package repository

import (
	"database/sql"
	"fmt"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type ScheduleApproveRepository interface {
	CreateApprove(payload model.ScheduleAprove) (model.ScheduleAprove, error)
	GetApproveById(schDetailID string) ([]dto.ScheduleAproveResponseDto, error)
}

type scheduleApproveRepository struct {
	db *sql.DB
}

func (sa *scheduleApproveRepository) GetApproveById(schDetailID string) (payload []dto.ScheduleAproveResponseDto, err error) {
	rows, err := sa.db.Query(utilsmodel.GetScheduleApproveById, schDetailID)

	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		schApprove := dto.ScheduleAproveResponseDto{}
		err := rows.Scan(&schApprove.Id, &schApprove.ScheduleDetailsId, &schApprove.ScheduleAprove, &schApprove.CreatedAt, &schApprove.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed scan, data not found")
		}
		payload = append(payload, schApprove)
	}

	return
}

func (sa *scheduleApproveRepository) CreateApprove(payload model.ScheduleAprove) (model.ScheduleAprove, error) {
	err := sa.db.QueryRow(utilsmodel.CreateScheduleApprove, payload.ScheduleDetails.Id, payload.ScheduleAprove, time.Now()).Scan(
		&payload.Id, &payload.CreatedAt, &payload.UpdatedAt,
	)
	if err != nil {
		return model.ScheduleAprove{}, err
	}
	return payload, nil
}

func NewScheduleApproveRepository(db *sql.DB) ScheduleApproveRepository {
	return &scheduleApproveRepository{db: db}
}
