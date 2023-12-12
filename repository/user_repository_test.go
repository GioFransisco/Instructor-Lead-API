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

type UserTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (s *UserTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(s.T(), err)
	s.mockDB = db
	s.mockSql = mock
	s.repo = NewUserRepository(s.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

func (s *UserTestSuite) TestUserGetRepo_Success() {
	mockUser := model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar77@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayu",
		Gander:      "L",
		Role:        "Trainer",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "username", "password", "age", "address", "gender", "role", "created_at", "updated_at"}).AddRow(mockUser.Id, mockUser.Name, mockUser.Email, mockUser.PhoneNumber, mockUser.Username, mockUser.Password, mockUser.Age, mockUser.Address, mockUser.Gander, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt)

	s.mockSql.ExpectQuery("SELECT (.*) FROM users WHERE id = \\$1").
		WillReturnRows(rows)

	actual, err := s.repo.Get(mockUser.Username)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockUser.Id, actual.Id)
}

func (s *UserTestSuite) TestUserGetRepo_Fail() {
	mockUser := model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar77@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayu",
		Gander:      "L",
		Role:        "Trainer",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "username", "password", "age", "address", "gender", "role", "created_at", "updated_at"}).AddRow(mockUser.Id, mockUser.Name, mockUser.Email, mockUser.PhoneNumber, mockUser.Username, mockUser.Password, mockUser.Age, mockUser.Address, mockUser.Gander, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt)

	s.mockSql.ExpectQuery("SELECT (.*) FROM users WHERE id = \\$1").
		WillReturnError(errors.New("user not found"))

	_, err := s.repo.Get(mockUser.Username)
	assert.Error(s.T(), err)

	s.mockSql.ExpectQuery("SELECT (.*) FROM users WHERE id = \\$1").
		WillReturnRows(rows)

}

func (s *UserTestSuite) TestUserUpdateRepo_Success() {
	mockUser := model.User{
		Id:          "54dac87d-da45-48b1-8182-b15178be6e56",
		Name:        "amar",
		Email:       "amar77@gmail.com",
		PhoneNumber: "0811775468198",
		Username:    "amarholo",
		Password:    "password",
		Age:         18,
		Address:     "Indramayu",
		Gander:      "L",
		Role:        "Trainer",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.mockSql.ExpectBegin()
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "username", "age", "gender", "address", "created_at", "updated_at"}).AddRow(mockUser.Id, mockUser.Name, mockUser.Email, mockUser.PhoneNumber, mockUser.Username, mockUser.Age, mockUser.Gander, mockUser.Address, mockUser.CreatedAt, mockUser.UpdatedAt)

	s.mockSql.ExpectQuery("Update users").WillReturnRows(rows)
	s.mockSql.ExpectCommit()
	actual, err := s.repo.UpdateUser(mockUser)
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), mockUser.Id, actual.Id)
}
