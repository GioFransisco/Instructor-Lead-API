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

type AuthTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    AuthRepository
}

func (suite *AuthTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewAuthRepository(suite.mockDB)
}

func TestAuthRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (suite *AuthTestSuite) TestAuthRepo_Success() {
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

	rows := sqlmock.NewRows([]string{"id", "password", "email", "role"}).AddRow(mockUser.Id, mockUser.Password, mockUser.Email, mockUser.Role)

	suite.mockSql.ExpectQuery("Select id,password,email,role From users Where username=\\$1").WithArgs(mockUser.Username).
		WillReturnRows(rows)

	actual, err := suite.repo.Login(mockUser.Username)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser.Id, actual.Id)
}

func (suite *AuthTestSuite) TestAuthRepo_Fail() {
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

	rows := sqlmock.NewRows([]string{"id", "password", "email", "role"}).AddRow(mockUser.Id, mockUser.Password, mockUser.Email, mockUser.Role)

	suite.mockSql.ExpectQuery("Select id,password,email,role From users Where username=\\$1").WithArgs(mockUser.Username).
		WillReturnError(errors.New("user not found"))

	_, err := suite.repo.Login(mockUser.Username)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectQuery("Select id,password,email,role From users Where username=\\$1").WithArgs(mockUser.Username).
		WillReturnRows(rows)
}

func (suite *AuthTestSuite) TestAuthRepoRegist_Success() {
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

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(mockUser.Id, mockUser.CreatedAt, mockUser.UpdatedAt)

	suite.mockSql.ExpectQuery("Insert Into users").
		WillReturnRows(rows)

	actual, err := suite.repo.Register(mockUser)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockUser.Id, actual.Id)
}

func (suite *AuthTestSuite) TestAuthRepoRegist_Fail() {
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

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(mockUser.Id, mockUser.CreatedAt, mockUser.UpdatedAt)

	suite.mockSql.ExpectQuery("Insert Into users").
		WillReturnError(errors.New("make sure all data is filled in correctly"))

	_, err := suite.repo.Register(mockUser)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectQuery("Insert Into users").
		WillReturnRows(rows)
}
