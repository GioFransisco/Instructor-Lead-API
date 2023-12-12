package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScheduleRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ScheduleRepository
}

func (suite *ScheduleRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewScheduleRepository(suite.mockDb)
}

var mockSchedule = model.Schedule{
	Id:           "1",
	Name:         "Instructor Led test name",
	DateActivity: time.Now(),
	ScheduleDetails: []model.ScheduleDetails{
		{
			Id:         "1",
			ScheduleId: "1",
			Trainer: model.User{
				Id:          "1",
				Name:        "Yopi",
				Email:       "yopitn@email.com",
				PhoneNumber: "089768758274",
				Username:    "yopitn",
				Age:         23,
				Address:     "Garut",
				Gander:      "L",
				Role:        "Trainer",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			Stack: model.Stack{
				Id:        "1",
				Name:      "Golang",
				Status:    "Active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartTime: time.Now(),
			EndTime:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:         "2",
			ScheduleId: "1",
			Trainer: model.User{
				Id:          "2",
				Name:        "Kira",
				Email:       "kira@email.com",
				PhoneNumber: "089768758272",
				Username:    "kira",
				Age:         23,
				Address:     "Garut",
				Gander:      "P",
				Role:        "Trainer",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			Stack: model.Stack{
				Id:        "2",
				Name:      "Java",
				Status:    "Active",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			StartTime: time.Now(),
			EndTime:   time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockSchedulePayload = model.Schedule{
	Id:           "1",
	Name:         "Instructor Led test name",
	DateActivity: time.Now(),
}

var mockScheduleDetailPayload = model.ScheduleDetails{
	Id: "1",
	Trainer: model.User{
		Id: "1",
	},
	Stack: model.Stack{
		Id: "1",
	},
	StartTime: time.Now(),
	EndTime:   time.Now(),
}

func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectBegin()

	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnRows(rows)

	for _, v := range mockSchedule.ScheduleDetails {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
		rows.AddRow(v.Id, v.CreatedAt, v.UpdatedAt)

		suite.mockSql.ExpectPrepare("INSERT INTO schedule_details").
			ExpectQuery().WithArgs(mockSchedule.Id, v.Trainer.Id, v.Stack.Id, v.StartTime, v.EndTime, sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	suite.mockSql.ExpectCommit()

	actual, err := suite.repo.Create(mockSchedule)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Fail_OnInserSchedule() {
	suite.mockSql.ExpectBegin().WillReturnError(errors.New("error when begin transaction"))
	_, err := suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").WillReturnError(errors.New("error when prepare query"))
	_, err = suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnError(errors.New("error when insert into data"))

	_, err = suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Fail_OnInsertScheduleDetail_Prepare() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnRows(rows)

	for range mockSchedule.ScheduleDetails {
		suite.mockSql.ExpectPrepare("INSERT INTO schedule_details").WillReturnError(errors.New("error when prepare query"))
	}

	_, err := suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnRows(rows)

	for range mockSchedule.ScheduleDetails {
		suite.mockSql.ExpectPrepare("INSERT INTO schedule_details").
			ExpectQuery().WillReturnError(errors.New("error when insert into schedule_details"))
	}

	_, err = suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Fail_OnInsertScheduleDetail_Query() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnRows(rows)

	for range mockSchedule.ScheduleDetails {
		suite.mockSql.ExpectPrepare("INSERT INTO schedule_details").
			ExpectQuery().WillReturnError(errors.New("error when insert into schedule_details"))
	}

	_, err := suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestCreateSchedule_Fail_OnCommit() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectBegin()
	suite.mockSql.ExpectPrepare("INSERT INTO schedules").
		ExpectQuery().WithArgs(mockSchedule.Name, mockSchedule.DateActivity, sqlmock.AnyArg()).
		WillReturnRows(rows)

	for _, v := range mockSchedule.ScheduleDetails {
		rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
		rows.AddRow(v.Id, v.CreatedAt, v.UpdatedAt)

		suite.mockSql.ExpectPrepare("INSERT INTO schedule_details").
			ExpectQuery().WithArgs(mockSchedule.Id, v.Trainer.Id, v.Stack.Id, v.StartTime, v.EndTime, sqlmock.AnyArg()).
			WillReturnRows(rows)
	}

	suite.mockSql.ExpectCommit().WillReturnError(errors.New("errors when commit"))

	_, err := suite.repo.Create(mockSchedule)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestListSchedule_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "date_activity", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules ORDER BY created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "start_time", "end_time", "created_at", "updated_at", "id", "name", "email", "phone_number", "username", "age", "address", "gender", "role", "created_at", "updated_at", "id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(
		mockSchedule.ScheduleDetails[0].Id,
		mockSchedule.ScheduleDetails[0].StartTime,
		mockSchedule.ScheduleDetails[0].EndTime,
		mockSchedule.ScheduleDetails[0].CreatedAt,
		mockSchedule.ScheduleDetails[0].UpdatedAt,
		mockSchedule.ScheduleDetails[0].Trainer.Id,
		mockSchedule.ScheduleDetails[0].Trainer.Name,
		mockSchedule.ScheduleDetails[0].Trainer.Email,
		mockSchedule.ScheduleDetails[0].Trainer.PhoneNumber,
		mockSchedule.ScheduleDetails[0].Trainer.Username,
		mockSchedule.ScheduleDetails[0].Trainer.Age,
		mockSchedule.ScheduleDetails[0].Trainer.Address,
		mockSchedule.ScheduleDetails[0].Trainer.Gander,
		mockSchedule.ScheduleDetails[0].Trainer.Role,
		mockSchedule.ScheduleDetails[0].Trainer.CreatedAt,
		mockSchedule.ScheduleDetails[0].Trainer.UpdatedAt,
		mockSchedule.ScheduleDetails[0].Stack.Id,
		mockSchedule.ScheduleDetails[0].Stack.Name,
		mockSchedule.ScheduleDetails[0].Stack.Status,
		mockSchedule.ScheduleDetails[0].Stack.CreatedAt,
		mockSchedule.ScheduleDetails[0].Stack.UpdatedAt,
	)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedule_details sd JOIN users t ON sd.trainer_id = t.id JOIN stacks s on sd.stack_id = s.id WHERE sd.schedule_id = \\$1 AND t.id = '1' ORDER BY sd.created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual[0].Id)
}

func (suite *ScheduleRepositoryTestSuite) TestListSchedule_Fail_OnSchedule() {
	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules ORDER BY created_at DESC").WillReturnError(errors.New("error when prepare query for schedules"))

	_, err := suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules ORDER BY created_at DESC").
		ExpectQuery().WillReturnError(errors.New("error when select from schedules"))

	_, err = suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Error(suite.T(), err)

	rows := sqlmock.NewRows([]string{"id", "name", "date_activity", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules ORDER BY created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	_, err = suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestListSchedule_Fail_OnScheduleDetail_Prepare() {
	rows := sqlmock.NewRows([]string{"id", "name", "date_activity", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules ORDER BY created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedule_details sd JOIN users t ON sd.trainer_id = t.id JOIN stacks s on sd.stack_id = s.id WHERE sd.schedule_id = \\$1 AND t.id = '1' ORDER BY sd.created_at DESC").
		WillReturnError(errors.New("error when prepare query for schedule_details"))

	_, err := suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestListSchedule_Fail_OnScheduleDetail_Query() {
	scheduleRows := sqlmock.NewRows([]string{"id", "name", "date_activity", "created_at", "updated_at"})
	scheduleRows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules").
		ExpectQuery().WillReturnRows(scheduleRows)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedule_details").
		ExpectQuery().WithArgs(mockSchedule.Id).WillReturnError(errors.New("error when select from schedule_details"))

	_, err := suite.repo.List(mockSchedule.Id, mockSchedule.ScheduleDetails[0].Trainer.Role)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestGetSchedule_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "date_activity", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedules").
		ExpectQuery().WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "start_time", "end_time", "created_at", "updated_at", "id", "name", "email", "phone_number", "username", "age", "address", "gender", "role", "created_at", "updated_at", "id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(
		mockSchedule.ScheduleDetails[0].Id,
		mockSchedule.ScheduleDetails[0].StartTime,
		mockSchedule.ScheduleDetails[0].EndTime,
		mockSchedule.ScheduleDetails[0].CreatedAt,
		mockSchedule.ScheduleDetails[0].UpdatedAt,
		mockSchedule.ScheduleDetails[0].Trainer.Id,
		mockSchedule.ScheduleDetails[0].Trainer.Name,
		mockSchedule.ScheduleDetails[0].Trainer.Email,
		mockSchedule.ScheduleDetails[0].Trainer.PhoneNumber,
		mockSchedule.ScheduleDetails[0].Trainer.Username,
		mockSchedule.ScheduleDetails[0].Trainer.Age,
		mockSchedule.ScheduleDetails[0].Trainer.Address,
		mockSchedule.ScheduleDetails[0].Trainer.Gander,
		mockSchedule.ScheduleDetails[0].Trainer.Role,
		mockSchedule.ScheduleDetails[0].Trainer.CreatedAt,
		mockSchedule.ScheduleDetails[0].Trainer.UpdatedAt,
		mockSchedule.ScheduleDetails[0].Stack.Id,
		mockSchedule.ScheduleDetails[0].Stack.Name,
		mockSchedule.ScheduleDetails[0].Stack.Status,
		mockSchedule.ScheduleDetails[0].Stack.CreatedAt,
		mockSchedule.ScheduleDetails[0].Stack.UpdatedAt,
	)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM schedule_details").
		ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.Get(mockSchedule.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.Id, actual.Id)
}

func (suite *ScheduleRepositoryTestSuite) TestGetScheduleDetail_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_id", "trainer_id", "stack_id", "start_time", "end_time", "created_at", "updated_at"})
	rows.AddRow(
		mockSchedule.ScheduleDetails[0].Id,
		mockSchedule.ScheduleDetails[0].ScheduleId,
		mockSchedule.ScheduleDetails[0].Trainer.Id,
		mockSchedule.ScheduleDetails[0].Stack.Id,
		mockSchedule.ScheduleDetails[0].StartTime,
		mockSchedule.ScheduleDetails[0].EndTime,
		mockSchedule.ScheduleDetails[0].CreatedAt,
		mockSchedule.ScheduleDetails[0].UpdatedAt,
	)

	suite.mockSql.ExpectPrepare("SELECT").ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.GetScheduleDetail(mockSchedule.ScheduleDetails[0].Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedule.ScheduleDetails[0].Id, actual.Id)
}

func (suite *ScheduleRepositoryTestSuite) TestGetScheduleDetail_Fail() {
	suite.mockSql.ExpectPrepare("SELECT").WillReturnError(errors.New("error when prepare query"))

	_, err := suite.repo.GetScheduleDetail(mockSchedule.ScheduleDetails[0].Id)
	assert.Error(suite.T(), err)

	rows := sqlmock.NewRows([]string{"id", "trainer_id", "stack_id", "start_time", "end_time", "created_at", "updated_at"})
	rows.AddRow(
		mockSchedule.ScheduleDetails[0].Id,
		mockSchedule.ScheduleDetails[0].Trainer.Id,
		mockSchedule.ScheduleDetails[0].Stack.Id,
		mockSchedule.ScheduleDetails[0].StartTime,
		mockSchedule.ScheduleDetails[0].EndTime,
		mockSchedule.ScheduleDetails[0].CreatedAt,
		mockSchedule.ScheduleDetails[0].UpdatedAt,
	)

	suite.mockSql.ExpectPrepare("SELECT").ExpectQuery().WillReturnRows(rows)
	_, err = suite.repo.GetScheduleDetail(mockSchedule.ScheduleDetails[0].Id)
	assert.Error(suite.T(), err)
}

func (suite *ScheduleRepositoryTestSuite) TestUpdateSchedule_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "date_activity", "created_at", "updated_at"})
	rows.AddRow(mockSchedule.Id, mockSchedule.Name, mockSchedule.DateActivity, mockSchedule.CreatedAt, mockSchedule.UpdatedAt)

	suite.mockSql.ExpectPrepare("UPDATE").ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.UpdateSchedule(mockSchedulePayload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockSchedulePayload.Id, actual.Id)
}

func (suite *ScheduleRepositoryTestSuite) TestUpdateScheduleDetail_Success() {
	rows := sqlmock.NewRows([]string{"id", "trainer_id", "stack_id", "start_time", "end_time", "created_at", "updated_at"})
	rows.AddRow(
		mockSchedule.ScheduleDetails[0].Id,
		mockSchedule.ScheduleDetails[0].Trainer.Id,
		mockSchedule.ScheduleDetails[0].Stack.Id,
		mockSchedule.ScheduleDetails[0].StartTime,
		mockSchedule.ScheduleDetails[0].EndTime,
		mockSchedule.ScheduleDetails[0].CreatedAt,
		mockSchedule.ScheduleDetails[0].UpdatedAt,
	)

	suite.mockSql.ExpectPrepare("UPDATE").ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.UpdateScheduleDetail(mockScheduleDetailPayload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockScheduleDetailPayload.Id, actual.Id)
}

func TestScheduleRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleRepositoryTestSuite))
}
