package repository

import (
	"database/sql"
	"testing"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ScheduleApproveRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ScheduleApproveRepository
}

func (suite *ScheduleApproveRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewScheduleApproveRepository(suite.mockDb)
}

func (suite *ScheduleApproveRepositoryTestSuite) TestFindById_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_detail_id", "schedule_approve", "created_at", "updated_at"})
	rows.AddRow("1", "1", "/images/test.png", time.Now(), time.Now())

	suite.mockSql.ExpectQuery("SELECT").
		WillReturnRows(rows)

	actual, err := suite.repo.GetApproveById("1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "1", actual[0].Id)
}

func (suite *ScheduleApproveRepositoryTestSuite) TestCreate_Success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"})
	rows.AddRow("1", time.Now(), time.Now())

	suite.mockSql.ExpectQuery("INSERT INTO schedule_approve").
		WillReturnRows(rows)

	payload := model.ScheduleAprove{
		ScheduleDetails: model.ScheduleDetails{
			Id: "1",
		},
		ScheduleAprove: "/test.png",
		UpdatedAt:      time.Now(),
	}

	actual, err := suite.repo.CreateApprove(payload)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "1", actual.Id)
}

func TestScheduleApproveRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleApproveRepositoryTestSuite))
}
