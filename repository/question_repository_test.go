package repository

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type QuestionSuiteTest struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    QuestionRepository
}

func (suite *QuestionSuiteTest) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewQusetionRepository(suite.mockDB)
}

func TestQuestionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(QuestionSuiteTest))
}

var mockQuestion = model.Question{
	Id: "1",
	ScheduleDetails: model.ScheduleDetails{
		Id:         "1",
		ScheduleId: "1",
		Trainer: model.User{
			Id:          "1",
			Name:        "jamal",
			Email:       "juned",
			PhoneNumber: "0893298219213",
			Username:    "jakalksa",
			Password:    "password",
			Age:         19,
			Address:     "indonseia timur",
			Gander:      "L",
			Role:        "Trainer",
			CreatedAt:   time.Now().Round(0),
			UpdatedAt:   time.Now().Round(0),
		},
		Stack: model.Stack{
			Id:        "1",
			Name:      "Golang",
			Status:    "Active",
			CreatedAt: time.Now().Round(0),
			UpdatedAt: time.Now().Round(0),
		},
		StartTime: time.Now().Round(0),
		EndTime:   time.Now().Round(0),
		CreatedAt: time.Now().Round(0),
		UpdatedAt: time.Now().Round(0),
	},
	StudentId: model.User{
		Id:          "1",
		Name:        "joko kendi;",
		Email:       "leskass@gmail.com",
		PhoneNumber: "08736273282",
		Username:    "jamal",
		Password:    "password",
		Age:         19,
		Address:     "Indramayu Barat",
		Gander:      "L",
		Role:        "Participant",
		CreatedAt:   time.Now().Round(0),
		UpdatedAt:   time.Now().Round(0),
	},
	Question:  "apakabar gaess",
	Status:    "Proccess",
	CreatedAt: time.Now().Round(0),
	UpdatedAt: time.Now().Round(0),
}

var mockDtoQuestion = dto.QuestionChangeDto{
	Id:       "1",
	Question: "apakabar gaess",
}

var mockDtoQuestionStatus = dto.QuestionChangeStatusDto{
	Id:     "1",
	Status: "Proccess",
}

func (s *QuestionSuiteTest) TestCreateQuestion_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(mockQuestion.Id, mockQuestion.CreatedAt, mockQuestion.UpdatedAt)

	s.mockSql.ExpectQuery("INSERT INTO student_questions").WillReturnRows(rows)

	actual, err := s.repo.CreateQuestion(mockQuestion)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionSuiteTest) TestCreateQuestion_Fail() {
	s.mockSql.ExpectQuery("INSERT INTO student_questions").WillReturnError(errors.New("error"))

	_, err := s.repo.CreateQuestion(mockQuestion)
	assert.Error(s.T(), err)
}

func (s *QuestionSuiteTest) TestDeleteQuestion_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_details_id", "student_id", "question", "status", "created_at", "updated_at"}).AddRow(mockQuestion.Id, mockQuestion.ScheduleDetails.Id, mockQuestion.StudentId.Id, mockQuestion.Question, mockQuestion.Status, mockQuestion.CreatedAt, mockQuestion.UpdatedAt)

	s.mockSql.ExpectQuery("Select (.*) From student_questions").WillReturnRows(rows)

	s.mockSql.ExpectExec("Delete From student_questions where id=\\$1").WithArgs(mockQuestion.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	actual, err := s.repo.DeleteQuestion(mockQuestion.Id)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionSuiteTest) TestDeleteQuestion_Fail() {
	s.mockSql.ExpectQuery("Select (.*) From student_questions").WillReturnError(errors.New("error"))

	_, err := s.repo.DeleteQuestion(mockQuestion.Id)
	assert.Error(s.T(), err)
}

func (s *QuestionSuiteTest) TestGetQuestion_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_details_id", "student_id", "question", "status", "created_at", "updated_at", "name", "email", "phone_number", "username", "age", "address", "gender", "role", "created_at", "updated_at"}).AddRow(mockQuestion.Id, mockQuestion.ScheduleDetails.Id, mockQuestion.StudentId.Id, mockQuestion.Question, mockQuestion.Status, mockQuestion.CreatedAt, mockQuestion.UpdatedAt, mockQuestion.StudentId.Name, mockQuestion.StudentId.Email, mockQuestion.StudentId.PhoneNumber, mockQuestion.StudentId.Username, mockQuestion.StudentId.Age, mockQuestion.StudentId.Address, mockQuestion.StudentId.Gander, mockQuestion.StudentId.Role, mockQuestion.StudentId.CreatedAt, mockQuestion.UpdatedAt)

	s.mockSql.ExpectQuery("Select sq.id,sq.schedule_details_id,sq.student_id,sq.question,sq.status,sq.created_at,sq.updated_at,u.name,u.email,u.phone_number,u.username,u.age,u.address,u.gender,u.role,u.created_at,u.updated_at from student_questions sq join users u ON sq.student_id = u.id where sq.schedule_details_id=\\$1").
		WithArgs(mockQuestion.ScheduleDetails.Id).
		WillReturnRows(rows)

	actual, err := s.repo.GetQuestionByID(mockQuestion.ScheduleDetails.Id)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.ScheduleDetails.Id, actual[0].Id)
}

func (s *QuestionSuiteTest) TestGetQuestion_Fail() {
	s.mockSql.ExpectQuery("Select sq.id,sq.schedule_details_id,sq.student_id,sq.question,sq.status,sq.created_at,sq.updated_at,u.name,u.email,u.phone_number,u.username,u.age,u.address,u.gender,u.role,u.created_at,u.updated_at from student_questions sq join users u ON sq.student_id = u.id where sq.schedule_details_id=\\$1").
		WithArgs(mockQuestion.ScheduleDetails.Id).
		WillReturnError(errors.New("error"))

	_, err := s.repo.GetQuestionByID(mockQuestion.ScheduleDetails.Id)
	assert.Error(s.T(), err)
}

func (s *QuestionSuiteTest) TestUpdateQuestion_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_details_id", "student_id", "question", "status", "created_at", "updated_at"}).AddRow(mockQuestion.Id, mockQuestion.ScheduleDetails.Id, mockQuestion.StudentId.Id, mockQuestion.Question, mockQuestion.Status, mockQuestion.CreatedAt, mockQuestion.UpdatedAt)

	rows2 := sqlmock.NewRows([]string{"id"}).AddRow(mockQuestion.Id)

	s.mockSql.ExpectQuery("Select id From student_questions Where id=\\$1").
		WithArgs(mockQuestion.Id).
		WillReturnRows(rows2)

	s.mockSql.ExpectQuery("Update student_questions SET question=\\$1,updated_at=\\$2 WHERE id=\\$3 RETURNING id,schedule_details_id,student_id,question,status,created_at,updated_at").
		WithArgs(mockQuestion.Question, sqlmock.AnyArg(), mockQuestion.Id).
		WillReturnRows(rows)

	actual, err := s.repo.UpdateQuestion(mockDtoQuestion)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionSuiteTest) TestUpdateQuestion_Fail() {
	s.mockSql.ExpectQuery("Select id From student_questions Where id=\\$1").
		WillReturnError(errors.New("error"))

	_, err := s.repo.UpdateQuestion(mockDtoQuestion)
	assert.Error(s.T(), err)
}

func (s *QuestionSuiteTest) TestUpdateStatusQuestion_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_details_id", "student_id", "question", "status", "created_at", "updated_at"}).AddRow(mockQuestion.Id, mockQuestion.ScheduleDetails.Id, mockQuestion.StudentId.Id, mockQuestion.Question, mockQuestion.Status, mockQuestion.CreatedAt, mockQuestion.UpdatedAt)

	rows2 := sqlmock.NewRows([]string{"id"}).AddRow(mockQuestion.Id)

	s.mockSql.ExpectQuery("Select id From student_questions Where id=\\$1").
		WithArgs(mockQuestion.Id).
		WillReturnRows(rows2)

	s.mockSql.ExpectQuery("Update student_questions SET status=\\$1,updated_at=\\$2 WHERE id=\\$3 RETURNING id,schedule_details_id,student_id,question,status,created_at,updated_at").
		WithArgs(mockQuestion.Status, sqlmock.AnyArg(), mockQuestion.Id).
		WillReturnRows(rows)

	actual, err := s.repo.UpdateStatusQuestion(mockDtoQuestionStatus)

	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockQuestion.Id, actual.Id)
}

func (s *QuestionSuiteTest) TestUpdateStatusQuestion_Fail() {
	s.mockSql.ExpectQuery("Select id From student_questions Where id=\\$1").
		WillReturnError(errors.New("error"))

	_, err := s.repo.UpdateStatusQuestion(mockDtoQuestionStatus)
	assert.Error(s.T(), err)
}
