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

type StackRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    StackRepository
}

func (suite *StackRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewStackRepository(suite.mockDb)
}

var mockStack model.Stack = model.Stack{
	Id:        "1",
	Name:      "Golang",
	Status:    "Active",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockPayloadStack model.Stack = model.Stack{
	Name:   "Golang",
	Status: "Active",
}

func (suite *StackRepositoryTestSuite) TestCreateStack_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectPrepare("INSERT INTO stacks (.*) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at, updated_at").
		ExpectQuery().WithArgs(mockStack.Name, mockStack.Status, sqlmock.AnyArg()).
		WillReturnRows(rows)

	actual, err := suite.repo.Create(mockPayloadStack)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockStack.Id, actual.Id)
}

func (suite *StackRepositoryTestSuite) TestCreateStack_Fail() {
	suite.mockSql.ExpectPrepare("INSERT INTO stacks \\(name, status, updated_at\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at, updated_at").
		WillReturnError(errors.New("error when prepare query"))

	_, err := suite.repo.Create(mockStack)
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectPrepare("INSERT INTO stacks \\(name, status, updated_at\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id, created_at, updated_at").
		ExpectQuery().WithArgs(mockStack.Name, mockStack.Status, sqlmock.AnyArg()).
		WillReturnError(errors.New("error when insert data"))

	_, err = suite.repo.Create(mockStack)
	assert.Error(suite.T(), err)
}

func (suite *StackRepositoryTestSuite) TestListStack_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, mockStack.Name, mockStack.Status, mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT (.*) FROM stacks ORDER BY created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	actual, err := suite.repo.List()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockStack.Id, actual[0].Id)
}

func (suite *StackRepositoryTestSuite) TestListStack_Fail() {
	suite.mockSql.ExpectPrepare("SELECT \\*\\ FROM stacks ORDER BY created_at DESC").
		WillReturnError(errors.New("error when prepare query"))

	_, err := suite.repo.List()
	assert.Error(suite.T(), err)

	suite.mockSql.ExpectPrepare("SELECT \\*\\ FROM stacks ORDER BY created_at DESC").
		ExpectQuery().WillReturnError(errors.New("error when query"))

	_, err = suite.repo.List()
	assert.Error(suite.T(), err)

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, mockStack.Name, mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectPrepare("SELECT \\*\\ FROM stacks ORDER BY created_at DESC").
		ExpectQuery().WillReturnRows(rows)

	_, err = suite.repo.List()
	assert.Error(suite.T(), err)
}

func (suite *StackRepositoryTestSuite) TestFindByIdStack_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, mockStack.Name, mockStack.Status, mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectQuery("SELECT (.*) FROM stacks WHERE id = \\$1").
		WillReturnRows(rows)

	actual, err := suite.repo.FindByID(mockStack.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockStack.Id, actual.Id)
}

func (suite *StackRepositoryTestSuite) TestFindByIdStack_Fail() {
	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, mockStack.Name, mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectQuery("SELECT \\*\\ FROM stacks WHERE id = \\$1").
		WillReturnRows(rows)

	_, err := suite.repo.FindByID(mockStack.Id)
	assert.Error(suite.T(), err)
}

func (suite *StackRepositoryTestSuite) TestUpdateStack_Success() {
	payload := model.Stack{
		Name:   "Java",
		Status: "Inactive",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, "Java", "Inactive", mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectQuery("UPDATE stacks SET (.*) RETURNING id, name, status, created_at, updated_at").
		WillReturnRows(rows)

	actual, err := suite.repo.Update(mockStack.Id, payload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), payload.Name, actual.Name)

	payload = model.Stack{
		Status: "Inactive",
	}

	rows = sqlmock.NewRows([]string{"id", "name", "status", "created_at", "updated_at"})
	rows.AddRow(mockStack.Id, "Java", "Inactive", mockStack.CreatedAt, mockStack.UpdatedAt)

	suite.mockSql.ExpectQuery("UPDATE stacks SET (.*) RETURNING id, name, status, created_at, updated_at").
		WillReturnRows(rows)

	actual, err = suite.repo.Update(mockStack.Id, payload)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), payload.Status, actual.Status)
}

func (suite *StackRepositoryTestSuite) TestUpdateStack_Fail() {
	suite.mockSql.ExpectQuery("UPDATE stacks SET (.*) RETURNING id, name, status, created_at, updated_at").
		WillReturnError(errors.New("error when query"))

	_, err := suite.repo.Update(mockStack.Id, mockPayloadStack)
	assert.Error(suite.T(), err)
}

func (suite *StackRepositoryTestSuite) TestDeleteStack_Fail() {
	suite.mockSql.ExpectExec("DELETE FROM stacks WHERE id = $1").
		WithArgs(mockStack.Id).
		WillReturnError(errors.New("error when query"))

	err := suite.repo.Delete(mockStack.Id)
	assert.Error(suite.T(), err)
}

func TestStackRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(StackRepositoryTestSuite))
}
