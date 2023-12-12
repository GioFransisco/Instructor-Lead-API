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

type AbsencesRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    AbsencesRepository
}

func (suite *AbsencesRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewAbsencesRepository(suite.mockDb)
}

var mockAbsences = model.Absences{
	Id: "1",
	ScheduleDetails: model.ScheduleDetails{
		Id: "1",
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
	},
	StudentId: model.User{
		Id:          "2",
		Name:        "Dina",
		Email:       "dina@email.com",
		PhoneNumber: "089768758274",
		Username:    "dina",
		Age:         23,
		Address:     "Garut",
		Gander:      "P",
		Role:        "Participant",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	},
	Description: "Hadir",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

var mockGetAbsences = model.GetAbsences{
	Id: "1",
	ScheduleDetails: []model.GetScheduleDetails{
		{
			Id: "1",
			Schedule: model.Schedule{
				Id:           "1",
				Name:         "Test name",
				DateActivity: time.Now(),
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			},
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
		},
	},
	Description: "Hadir",
	CreatedAt:   time.Now(),
	UpdatedAt:   time.Now(),
}

func (suite *AbsencesRepositoryTestSuite) TestCreateAbsences_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockAbsences.Id, mockAbsences.CreatedAt, mockAbsences.UpdatedAt)

	suite.mockSql.ExpectQuery("INSERT INTO absences").WillReturnRows(rows)

	actual, err := suite.repo.Create(mockAbsences)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockAbsences.Id, actual.Id)
}

func (suite *AbsencesRepositoryTestSuite) TestCreateAbsences_Fail() {
	suite.mockSql.ExpectQuery("INSERT INTO absences").WillReturnError(errors.New("error"))

	_, err := suite.repo.Create(mockAbsences)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesRepositoryTestSuite) TestGetScheduleDetailId_Success() {
	rows := sqlmock.NewRows([]string{"a.id", "u.id", "u.name", "u.email", "u.phone_number", "u.username", "u.password", "u.age", "u.address", "u.gender", "u.role", "u.created_at", "u.updated_at", "a.description", "a.created_at", "a.updated_at"})
	rows.AddRow(
		mockAbsences.Id,
		mockAbsences.StudentId.Id,
		mockAbsences.StudentId.Name,
		mockAbsences.StudentId.Email,
		mockAbsences.StudentId.PhoneNumber,
		mockAbsences.StudentId.Username,
		mockAbsences.StudentId.Password,
		mockAbsences.StudentId.Age,
		mockAbsences.StudentId.Address,
		mockAbsences.StudentId.Gander,
		mockAbsences.StudentId.Role,
		mockAbsences.StudentId.CreatedAt,
		mockAbsences.StudentId.UpdatedAt,
		mockAbsences.Description,
		mockAbsences.CreatedAt,
		mockAbsences.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT (.+) FROM absences AS a JOIN users AS u ON u.id = a.student_id WHERE a.schedule_details_id = \\$1").
		WithArgs(mockAbsences.ScheduleDetails.Id).
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"sd.id", "s.id", "s.name", " s.date_activity", "s.created_at", "s.updated_at", "u.id", "u.name", "u.email", "u.phone_number", "u.username", "u.password", "u.age", "u.address", "u.gender", "u.role", "u.created_at", "u.updated_at", "st.id", "st.name", "st.status", "st.created_at", "st.updated_at", "sd.start_time", "sd.end_time", "sd.created_at", "sd.updated_at"})
	rows.AddRow(
		mockGetAbsences.ScheduleDetails[0].Id,
		mockGetAbsences.ScheduleDetails[0].Schedule.Id,
		mockGetAbsences.ScheduleDetails[0].Schedule.Name,
		mockGetAbsences.ScheduleDetails[0].Schedule.DateActivity,
		mockGetAbsences.ScheduleDetails[0].Schedule.CreatedAt,
		mockGetAbsences.ScheduleDetails[0].Schedule.UpdatedAt,
		mockGetAbsences.ScheduleDetails[0].Trainer.Id,
		mockGetAbsences.ScheduleDetails[0].Trainer.Name,
		mockGetAbsences.ScheduleDetails[0].Trainer.Email,
		mockGetAbsences.ScheduleDetails[0].Trainer.PhoneNumber,
		mockGetAbsences.ScheduleDetails[0].Trainer.Username,
		mockGetAbsences.ScheduleDetails[0].Trainer.Password,
		mockGetAbsences.ScheduleDetails[0].Trainer.Age,
		mockGetAbsences.ScheduleDetails[0].Trainer.Address,
		mockGetAbsences.ScheduleDetails[0].Trainer.Gander,
		mockGetAbsences.ScheduleDetails[0].Trainer.Role,
		mockGetAbsences.ScheduleDetails[0].Trainer.CreatedAt,
		mockGetAbsences.ScheduleDetails[0].Trainer.UpdatedAt,
		mockGetAbsences.ScheduleDetails[0].Stack.Id,
		mockGetAbsences.ScheduleDetails[0].Stack.Name,
		mockGetAbsences.ScheduleDetails[0].Stack.Status,
		mockGetAbsences.ScheduleDetails[0].Stack.CreatedAt,
		mockGetAbsences.ScheduleDetails[0].Stack.UpdatedAt,
		mockGetAbsences.ScheduleDetails[0].StartTime,
		mockGetAbsences.ScheduleDetails[0].EndTime,
		mockGetAbsences.ScheduleDetails[0].CreatedAt,
		mockGetAbsences.ScheduleDetails[0].UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT (.+) FROM schedule_details AS sd JOIN schedules AS s ON s.id = sd.schedule_id JOIN users AS u ON u.id = sd.trainer_id JOIN stacks AS st ON st.id = sd.stack_id WHERE sd.id = \\$1").
		WithArgs(mockAbsences.ScheduleDetails.Id).
		WillReturnRows(rows)

	actual, err := suite.repo.GetScheduleDetailId(mockAbsences.ScheduleDetails.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), mockStack.Id, actual.Id)
}

func (suite *AbsencesRepositoryTestSuite) TestGetScheduleDetailId_Fail_OnAbsences() {
	suite.mockSql.ExpectQuery("SELECT (.+) FROM absences AS a JOIN users AS u ON u.id = a.student_id WHERE a.schedule_details_id = \\$1").
		WithArgs(mockAbsences.ScheduleDetails.Id).
		WillReturnError(errors.New("error"))

	_, err := suite.repo.GetScheduleDetailId(mockAbsences.ScheduleDetails.Id)
	assert.Error(suite.T(), err)
}

func (suite *AbsencesRepositoryTestSuite) TestGetScheduleDetailId_Fail_OnScheduleDetailQuery() {
	rows := sqlmock.NewRows([]string{"a.id", "u.id", "u.name", "u.email", "u.phone_number", "u.username", "u.password", "u.age", "u.address", "u.gender", "u.role", "u.created_at", "u.updated_at", "a.description", "a.created_at", "a.updated_at"})
	rows.AddRow(
		mockAbsences.Id,
		mockAbsences.StudentId.Id,
		mockAbsences.StudentId.Name,
		mockAbsences.StudentId.Email,
		mockAbsences.StudentId.PhoneNumber,
		mockAbsences.StudentId.Username,
		mockAbsences.StudentId.Password,
		mockAbsences.StudentId.Age,
		mockAbsences.StudentId.Address,
		mockAbsences.StudentId.Gander,
		mockAbsences.StudentId.Role,
		mockAbsences.StudentId.CreatedAt,
		mockAbsences.StudentId.UpdatedAt,
		mockAbsences.Description,
		mockAbsences.CreatedAt,
		mockAbsences.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT (.+) FROM absences AS a JOIN users AS u ON u.id = a.student_id WHERE a.schedule_details_id = \\$1").
		WithArgs(mockAbsences.ScheduleDetails.Id).
		WillReturnRows(rows)

	suite.mockSql.ExpectQuery("SELECT (.+) FROM schedule_details AS sd JOIN schedules AS s ON s.id = sd.schedule_id JOIN users AS u ON u.id = sd.trainer_id JOIN stacks AS st ON st.id = sd.stack_id WHERE sd.id = \\$1").
		WithArgs(mockAbsences.ScheduleDetails.Id).
		WillReturnError(errors.New("error"))

	_, err := suite.repo.GetScheduleDetailId(mockAbsences.ScheduleDetails.Id)
	assert.Error(suite.T(), err)
}

func TestAbsencesRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AbsencesRepositoryTestSuite))
}
