package repository

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
)

type ScheduleRepository interface {
	Create(payload model.Schedule) (model.Schedule, error)
	List(userId, userRole string) ([]model.Schedule, error)
	Get(id string) (model.Schedule, error)
	GetSchedule(id string) (model.Schedule, error)
	GetScheduleDetail(id string) (model.ScheduleDetails, error)
	UpdateSchedule(payload model.Schedule) (model.Schedule, error)
	UpdateScheduleDetail(payload model.ScheduleDetails) (model.ScheduleDetails, error)
}

type scheduleRepository struct {
	db *sql.DB
}

func (s *scheduleRepository) Create(payload model.Schedule) (model.Schedule, error) {
	var schedule model.Schedule
	var scheduleDetails []model.ScheduleDetails

	tx, err := s.db.Begin()
	if err != nil {
		return model.Schedule{}, err
	}

	scheduleStmt, err := tx.Prepare(utilsmodel.ScheduleCreate)

	if err != nil {
		return model.Schedule{}, tx.Rollback()
	}

	defer scheduleStmt.Close()

	err = scheduleStmt.QueryRow(payload.Name, payload.DateActivity, time.Now()).Scan(&schedule.Id, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		return model.Schedule{}, tx.Rollback()
	}

	for _, v := range payload.ScheduleDetails {
		var scheduleDetail model.ScheduleDetails
		scheduleDetailStmt, err := tx.Prepare(utilsmodel.ScheduleDetailCreate)

		if err != nil {
			return model.Schedule{}, tx.Rollback()
		}

		if err := scheduleDetailStmt.QueryRow(schedule.Id, v.Trainer.Id, v.Stack.Id, v.StartTime, v.EndTime, time.Now()).Scan(&scheduleDetail.Id, &scheduleDetail.CreatedAt, &scheduleDetail.UpdatedAt); err != nil {
			return model.Schedule{}, tx.Rollback()
		}

		scheduleDetail.ScheduleId = v.ScheduleId
		scheduleDetail.Trainer = v.Trainer
		scheduleDetail.StartTime = v.StartTime
		scheduleDetail.EndTime = v.EndTime
		scheduleDetail.Stack = v.Stack

		scheduleDetails = append(scheduleDetails, scheduleDetail)
	}

	schedule.Name = payload.Name
	schedule.DateActivity = payload.DateActivity
	schedule.ScheduleDetails = scheduleDetails

	if err := tx.Commit(); err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

func (s *scheduleRepository) List(userId, userRole string) ([]model.Schedule, error) {
	var schedules []model.Schedule

	scheduleQry := utilsmodel.ScheduleGet
	scheduleStmt, err := s.db.Prepare(scheduleQry)
	if err != nil {
		return []model.Schedule{}, err
	}

	defer scheduleStmt.Close()

	scheduleRows, err := scheduleStmt.Query()
	if err != nil {
		return []model.Schedule{}, err
	}

	defer scheduleRows.Close()

	scheduleDetailQry := utilsmodel.ScheduleDetailGet
	if userRole == "Trainer" {
		scheduleDetailQry += fmt.Sprintf("AND t.id = '%s' ", userId)
	}

	scheduleDetailQry += "ORDER BY sd.created_at DESC"

	for scheduleRows.Next() {
		var schedule model.Schedule

		if err := scheduleRows.Scan(&schedule.Id, &schedule.Name, &schedule.DateActivity, &schedule.CreatedAt, &schedule.UpdatedAt); err != nil {
			return []model.Schedule{}, err
		}

		var scheduleDetails []model.ScheduleDetails

		scheduleDetailStmt, err := s.db.Prepare(scheduleDetailQry)
		if err != nil {
			return []model.Schedule{}, err
		}

		defer scheduleDetailStmt.Close()

		scheduleDetailRows, err := scheduleDetailStmt.Query(schedule.Id)
		if err != nil {
			return []model.Schedule{}, err
		}

		defer scheduleDetailRows.Close()

		for scheduleDetailRows.Next() {
			var scheduleDetail model.ScheduleDetails

			err := scheduleDetailRows.Scan(
				&scheduleDetail.Id,
				&scheduleDetail.StartTime,
				&scheduleDetail.EndTime,
				&scheduleDetail.CreatedAt,
				&scheduleDetail.UpdatedAt,
				&scheduleDetail.Trainer.Id,
				&scheduleDetail.Trainer.Name,
				&scheduleDetail.Trainer.Email,
				&scheduleDetail.Trainer.PhoneNumber,
				&scheduleDetail.Trainer.Username,
				&scheduleDetail.Trainer.Age,
				&scheduleDetail.Trainer.Address,
				&scheduleDetail.Trainer.Gander,
				&scheduleDetail.Trainer.Role,
				&scheduleDetail.Trainer.CreatedAt,
				&scheduleDetail.Trainer.UpdatedAt,
				&scheduleDetail.Stack.Id,
				&scheduleDetail.Stack.Name,
				&scheduleDetail.Stack.Status,
				&scheduleDetail.Stack.CreatedAt,
				&scheduleDetail.Stack.UpdatedAt,
			)

			if err != nil {
				return []model.Schedule{}, err
			}

			scheduleDetails = append(scheduleDetails, scheduleDetail)
		}

		schedule.ScheduleDetails = scheduleDetails

		schedules = append(schedules, schedule)
	}

	return schedules, nil
}

func (s *scheduleRepository) Get(id string) (model.Schedule, error) {
	schedule, err := s.GetSchedule(id)
	if err != nil {
		return model.Schedule{}, err
	}

	scheduleDetailQry := utilsmodel.ScheduleDetailGet
	scheduleDetailQry += "ORDER BY sd.created_at DESC"

	db, err := s.db.Prepare(scheduleDetailQry)
	if err != nil {
		return model.Schedule{}, err
	}

	defer db.Close()

	scheduleDetailRows, err := db.Query(schedule.Id)
	if err != nil {
		return model.Schedule{}, err
	}

	var scheduleDetails []model.ScheduleDetails
	for scheduleDetailRows.Next() {
		var scheduleDetail model.ScheduleDetails

		err := scheduleDetailRows.Scan(
			&scheduleDetail.Id,
			&scheduleDetail.StartTime,
			&scheduleDetail.EndTime,
			&scheduleDetail.CreatedAt,
			&scheduleDetail.UpdatedAt,
			&scheduleDetail.Trainer.Id,
			&scheduleDetail.Trainer.Name,
			&scheduleDetail.Trainer.Email,
			&scheduleDetail.Trainer.PhoneNumber,
			&scheduleDetail.Trainer.Username,
			&scheduleDetail.Trainer.Age,
			&scheduleDetail.Trainer.Address,
			&scheduleDetail.Trainer.Gander,
			&scheduleDetail.Trainer.Role,
			&scheduleDetail.Trainer.CreatedAt,
			&scheduleDetail.Trainer.UpdatedAt,
			&scheduleDetail.Stack.Id,
			&scheduleDetail.Stack.Name,
			&scheduleDetail.Stack.Status,
			&scheduleDetail.Stack.CreatedAt,
			&scheduleDetail.Stack.UpdatedAt,
		)

		if err != nil {
			return model.Schedule{}, err
		}

		scheduleDetails = append(scheduleDetails, scheduleDetail)
	}

	schedule.ScheduleDetails = scheduleDetails

	return schedule, nil
}

func (s *scheduleRepository) GetSchedule(id string) (model.Schedule, error) {
	var schedule model.Schedule

	db, err := s.db.Prepare(utilsmodel.ScheduleGetById)
	if err != nil {
		return model.Schedule{}, err
	}

	defer db.Close()

	err = db.QueryRow(id).Scan(&schedule.Id, &schedule.Name, &schedule.DateActivity, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

func (s *scheduleRepository) GetScheduleDetail(id string) (model.ScheduleDetails, error) {
	var scheduleDetail model.ScheduleDetails

	db, err := s.db.Prepare(utilsmodel.ScheduleDetailGetById)
	if err != nil {
		return model.ScheduleDetails{}, err
	}

	defer db.Close()

	err = db.QueryRow(id).Scan(
		&scheduleDetail.Id,
		&scheduleDetail.ScheduleId,
		&scheduleDetail.Trainer.Id,
		&scheduleDetail.Stack.Id,
		&scheduleDetail.StartTime,
		&scheduleDetail.EndTime,
		&scheduleDetail.CreatedAt,
		&scheduleDetail.UpdatedAt,
	)

	if err != nil {
		return model.ScheduleDetails{}, err
	}

	return scheduleDetail, nil
}

func (s *scheduleRepository) UpdateSchedule(payload model.Schedule) (model.Schedule, error) {
	var schedule model.Schedule

	index := 1
	var value []any
	qry := utilsmodel.ScheduleUpdate

	if payload.Name != "" {
		qry += "name = $" + strconv.Itoa(index)
		value = append(value, payload.Name)
		index++
	}

	if !payload.DateActivity.IsZero() {
		if index > 1 {
			qry += ", date_activity = $" + strconv.Itoa(index)
		} else {
			qry += "date_activity = $" + strconv.Itoa(index)
		}

		value = append(value, payload.DateActivity)
		index++
	}

	if index > 1 {
		qry += ", updated_at = $" + strconv.Itoa(index)
	} else {
		qry += "updated_at = $" + strconv.Itoa(index)
	}

	value = append(value, time.Now())
	index++

	qry += " WHERE id = $" + strconv.Itoa(index) + utilsmodel.ScheduleUpdateReturning
	value = append(value, payload.Id)

	db, err := s.db.Prepare(qry)
	if err != nil {
		return model.Schedule{}, err
	}

	defer db.Close()

	err = db.QueryRow(value...).Scan(&schedule.Id, &schedule.Name, &schedule.DateActivity, &schedule.CreatedAt, &schedule.UpdatedAt)
	if err != nil {
		return model.Schedule{}, err
	}

	return schedule, nil
}

func (s *scheduleRepository) UpdateScheduleDetail(payload model.ScheduleDetails) (model.ScheduleDetails, error) {
	var scheduleDetail model.ScheduleDetails

	index := 1
	var value []any
	qry := utilsmodel.ScheduleDetailUpdate

	if payload.Trainer.Id != "" {
		qry += "trainer_id = $" + strconv.Itoa(index)
		value = append(value, payload.Trainer.Id)
		index++
	}

	if payload.Stack.Id != "" {
		if index > 1 {
			qry += ", stack_id = $" + strconv.Itoa(index)
		} else {
			qry += "stack_id = $" + strconv.Itoa(index)
		}

		value = append(value, payload.Stack.Id)
		index++
	}

	if !payload.StartTime.IsZero() {
		if index > 1 {
			qry += ", start_time = $" + strconv.Itoa(index)
		} else {
			qry += "start_time = $" + strconv.Itoa(index)
		}

		value = append(value, payload.StartTime)
		index++
	}

	if !payload.EndTime.IsZero() {
		if index > 1 {
			qry += ", end_time = $" + strconv.Itoa(index)
		} else {
			qry += "end_time = $" + strconv.Itoa(index)
		}

		value = append(value, payload.EndTime)
		index++
	}

	if index > 1 {
		qry += ", updated_at = $" + strconv.Itoa(index)
	} else {
		qry += "updated_at = $" + strconv.Itoa(index)
	}

	value = append(value, time.Now())
	index++

	qry += " WHERE id = $" + strconv.Itoa(index) + utilsmodel.ScheduleDetailUpdateReturning
	value = append(value, payload.Id)

	db, err := s.db.Prepare(qry)
	if err != nil {
		return model.ScheduleDetails{}, err
	}

	defer db.Close()

	err = db.QueryRow(value...).Scan(
		&scheduleDetail.Id,
		&scheduleDetail.Trainer.Id,
		&scheduleDetail.Stack.Id,
		&scheduleDetail.StartTime,
		&scheduleDetail.EndTime,
		&scheduleDetail.CreatedAt,
		&scheduleDetail.UpdatedAt,
	)

	if err != nil {
		return model.ScheduleDetails{}, err
	}

	return scheduleDetail, nil
}

func NewScheduleRepository(db *sql.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}
