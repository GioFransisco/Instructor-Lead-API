package repository

import (
	"database/sql"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
)

type NoteRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    NoteRepository
}

func (suite *NoteRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)

	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewNoteRepository(suite.mockDB)
}

var mockPayloadNote = dto.NoteDTO{
	ScheduleID: "1b39443e-9bd4-4307-a301-2696f267117f",
	UserEmail:  "test@example.com",
	Note:       "Test note",
}

func (suite *NoteRepositoryTestSuite) TestCreate_Success() {
	rows := sqlmock.NewRows([]string{"id", "schedule_details_id", "email", "note", "created_at", "updated_at"})
	rows.AddRow("a8fbf953-0b6d-4a76-8345-90a03cdb8f3b", mockPayloadNote.ScheduleID, mockPayloadNote.UserEmail, mockPayloadNote.Note, time.Now(), time.Now())

	suite.mockSql.ExpectQuery("INSERT INTO notes").
		WithArgs(sqlmock.AnyArg(), mockPayloadNote.ScheduleID, mockPayloadNote.UserEmail, mockPayloadNote.Note, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	actual, err := suite.repo.Create(mockPayloadNote)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), mockPayloadNote.ScheduleID, actual.ScheduleID)
}

func TestNoteRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(NoteRepositoryTestSuite))
}
