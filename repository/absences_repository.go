package repository

import (
	"database/sql"
	"log"
	"time"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type AbsencesRepository interface {
	Create(payload model.Absences) (model.Absences, error)
	GetScheduleDetailId(scheduleDetailId string) (model.GetAbsences, error)
}

type absencesRepository struct {
	db *sql.DB
}

func (a *absencesRepository) Create(payload model.Absences) (model.Absences, error) {
	err := a.db.QueryRow(utilsmodel.CreateAbsences, payload.ScheduleDetails.Id, payload.StudentId.Id, payload.Description, time.Now()).Scan(
		&payload.Id, &payload.CreatedAt, &payload.UpdatedAt,
	)
	if err != nil {
		return model.Absences{}, err
	}
	return payload, nil
}

func (a *absencesRepository) GetScheduleDetailId(scheduleDetailid string) (model.GetAbsences, error) {
	var absences model.GetAbsences
	err := a.db.QueryRow(utilsmodel.GetAbsences, scheduleDetailid).Scan(
		&absences.Id, &absences.Student.Id, &absences.Student.Name, &absences.Student.Email, 
		&absences.Student.PhoneNumber, &absences.Student.Username, &absences.Student.Password, 
		&absences.Student.Age, &absences.Student.Address, &absences.Student.Gander, &absences.Student.Role,
		&absences.Student.CreatedAt, &absences.Student.UpdatedAt, &absences.Description, &absences.CreatedAt, &absences.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error retrieving absences: %v", err)
		return model.GetAbsences{}, err
	}

	var getScheduleDetails []model.GetScheduleDetails
	rows, err := a.db.Query(utilsmodel.GetScheduleDetailAbsenceById, scheduleDetailid)

	if err != nil {
		log.Printf("Error retrieving schedule detail: %v", err)
		return model.GetAbsences{}, err
	}

	for rows.Next() {
		var getScheduleDetail model.GetScheduleDetails
		rows.Scan(
			&getScheduleDetail.Id, &getScheduleDetail.Schedule.Id, &getScheduleDetail.Schedule.Name, 
			&getScheduleDetail.Schedule.DateActivity, &getScheduleDetail.Schedule.CreatedAt, 
			&getScheduleDetail.Schedule.UpdatedAt, &getScheduleDetail.Trainer.Id, &getScheduleDetail.Trainer.Name, 
			&getScheduleDetail.Trainer.Email, &getScheduleDetail.Trainer.PhoneNumber, &getScheduleDetail.Trainer.Username, 
			&getScheduleDetail.Trainer.Password, &getScheduleDetail.Trainer.Age, &getScheduleDetail.Trainer.Address, 
			&getScheduleDetail.Trainer.Gander, &getScheduleDetail.Trainer.Role, &getScheduleDetail.Trainer.CreatedAt, 
			&getScheduleDetail.Trainer.UpdatedAt, &getScheduleDetail.Stack.Id, &getScheduleDetail.Stack.Name, 
			&getScheduleDetail.Stack.Status, &getScheduleDetail.Stack.CreatedAt, &getScheduleDetail.Stack.UpdatedAt,
			&getScheduleDetail.StartTime, &getScheduleDetail.EndTime, &getScheduleDetail.CreatedAt, &getScheduleDetail.UpdatedAt,
		)
		getScheduleDetails = append(getScheduleDetails, getScheduleDetail)
	}
	absences.ScheduleDetails = getScheduleDetails
	return absences, nil
}

func NewAbsencesRepository(db *sql.DB) AbsencesRepository {
	return &absencesRepository{db: db}
}
