package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	middlewaremock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/middleware_mock"
	usecasemock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/usecase_mock"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type NoteControllerTestSuite struct {
	suite.Suite
	rg              *gin.RouterGroup
	noteUseCaseMock *usecasemock.NoteUseCaseMock
	authMiddleware  *middlewaremock.AuthMiddlewareMock
	noteController  *NoteController
	record          *httptest.ResponseRecorder
}

func (suite *NoteControllerTestSuite) SetupTest() {
	engine := gin.Default()
	suite.rg = engine.Group("/api/v1")
	suite.noteUseCaseMock = new(usecasemock.NoteUseCaseMock)
	suite.authMiddleware = new(middlewaremock.AuthMiddlewareMock)

	suite.noteController = NewNoteController(suite.noteUseCaseMock, suite.rg, suite.authMiddleware)
	suite.record = httptest.NewRecorder()

	suite.noteController.Route()
}

var mockNoteDTO = dto.NoteDTO{
	ScheduleID: "1",
	UserEmail:  "user@example.com",
	Note:       "This is a note",
	CreatedAt:  "2023-01-01T12:00:00Z",
	UpdatedAt:  "2023-01-01T12:30:00Z",
}

var mockPayloadNoteDTO = dto.NoteDTO{
	ScheduleID: "1",
	UserEmail:  "user@example.com",
	Note:       "This is a note",
	CreatedAt:  "2023-01-01T12:00:00Z",
	UpdatedAt:  "2023-01-01T12:30:00Z",
}

// func (suite *NoteControllerTestSuite) TestCreateHandler_Success() {
// 	suite.noteUseCaseMock.On("Create", mockPayloadNoteDTO).Return(mockNoteDTO, nil)

// 	mockPayloadJSON, err := json.Marshal(mockPayloadNoteDTO)
// 	assert.NoError(suite.T(), err)

// 	req, err := http.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewBuffer(mockPayloadJSON))
// 	assert.NoError(suite.T(), err)

// 	ctx, _ := gin.CreateTestContext(suite.record)
// 	ctx.Request = req

// 	suite.noteController.createHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusCreated, suite.record.Code)
// }

// func (suite *NoteControllerTestSuite) TestCreateHandler_Fail() {
// 	req, err := http.NewRequest(http.MethodPost, "/api/v1/notes", nil)
// 	assert.NoError(suite.T(), err)

// 	ctx, _ := gin.CreateTestContext(suite.record)
// 	ctx.Request = req

// 	suite.noteController.createHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

// 	suite.noteUseCaseMock.On("Create", mockPayloadNoteDTO).Return(dto.NoteDTO{}, errors.New("error when creating note"))

// 	mockPayloadJSON, err := json.Marshal(mockPayloadNoteDTO)
// 	assert.NoError(suite.T(), err)

// 	req, err = http.NewRequest(http.MethodPost, "/api/v1/notes", bytes.NewBuffer(mockPayloadJSON))
// 	assert.NoError(suite.T(), err)

// 	ctx.Request = req

// 	suite.noteController.createHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
// }

func (suite *NoteControllerTestSuite) TestListHandler_Success() {
	mockNoteDTOs := []dto.NoteDTO{
		{

			ScheduleID: "1b39443e-9bd4-4307-a301-2696f267117f",
			UserEmail:  "user@example.com",
			Note:       "This is a note",
			CreatedAt:  "2023-01-01T12:00:00Z",
			UpdatedAt:  "2023-01-01T12:30:00Z",
		},
		{

			ScheduleID: "1b39443e-9bd4-4307-a301-2696f267117i",
			UserEmail:  "user@example.com",
			Note:       "Another note",
			CreatedAt:  "2023-01-01T12:15:00Z",
			UpdatedAt:  "2023-01-01T12:45:00Z",
		},
	}

	suite.noteUseCaseMock.On("FindAll").Return(mockNoteDTOs, nil)
	req, err := http.NewRequest(http.MethodGet, "/api/v1/notes/1b39443e-9bd4-4307-a301-2696f267117i", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.noteController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
}

func (suite *NoteControllerTestSuite) TestListHandler_Fail() {
	suite.noteUseCaseMock.On("FindAll").Return([]dto.NoteDTO{}, errors.New("error when get data from ListNotes usecase"))

	req, err := http.NewRequest(http.MethodGet, "/api/v1/notes/:id", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.noteController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

// func (suite *NoteControllerTestSuite) TestUpdateHandler_Success() {
// 	payload := model.Note{
// 		ScheduleDetails: model.ScheduleDetails{Id: mockNoteDTO.ScheduleID},
// 		UserEmail:       mockNoteDTO.UserEmail,
// 		Note:            mockNoteDTO.Note,
// 	}

// 	mockPayloadJSON, err := json.Marshal(mockPayloadNoteDTO)
// 	assert.NoError(suite.T(), err)

// 	req, err := http.NewRequest(http.MethodPut, "/api/v1/notes/1b39443e-9bd4-4307-a301-2696f267117", bytes.NewBuffer(mockPayloadJSON))
// 	assert.NoError(suite.T(), err)

// 	ctx, _ := gin.CreateTestContext(suite.record)
// 	ctx.Request = req

// 	noteId := ctx.Param("id")
// 	suite.noteUseCaseMock.On("Update", noteId, payload).Return(model.Note{}, nil)

// 	suite.noteController.updateHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusOK, suite.record.Code)
// }

func (suite *NoteControllerTestSuite) TestUpdateHandler_Fail() {
	// payload := model.Note{
	// 	ScheduleDetails: model.ScheduleDetails{Id: mockNoteDTO.ScheduleID},
	// 	UserEmail:       mockNoteDTO.UserEmail,
	// 	Note:            mockNoteDTO.Note,
	// }

	req, err := http.NewRequest(http.MethodPut, "/api/v1/notes/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	suite.noteController.updateHandler(ctx)
	assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)

	// mockPayloadJSON, err := json.Marshal(mockPayloadNoteDTO)
	// assert.NoError(suite.T(), err)

	// req, err = http.NewRequest(http.MethodPut, "/api/v1/notes/1", bytes.NewBuffer(mockPayloadJSON))
	// assert.NoError(suite.T(), err)

	// ctx.Request = req

	// noteId := ctx.Param("id")
	// suite.noteUseCaseMock.On("Update", noteId, payload).Return(model.Note{}, errors.New("error when update note data"))

	// suite.noteController.updateHandler(ctx)
	// assert.Equal(suite.T(), http.StatusBadRequest, suite.record.Code)
}

func (suite *NoteControllerTestSuite) TestDeleteHandler_Success() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/notes/1", nil)
	assert.NoError(suite.T(), err)
	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req
	noteId := ctx.Param("id")
	suite.noteUseCaseMock.On("Delete", noteId).Return(nil)

	suite.noteController.deleteHandler(ctx)
	assert.Equal(suite.T(), http.StatusNoContent, suite.record.Code)
}

func (suite *NoteControllerTestSuite) TestDeleteHandler_Fail() {
	req, err := http.NewRequest(http.MethodDelete, "/api/v1/notes/1", nil)
	assert.NoError(suite.T(), err)

	ctx, _ := gin.CreateTestContext(suite.record)
	ctx.Request = req

	noteId := ctx.Param("id")
	suite.noteUseCaseMock.On("Delete", noteId).Return(errors.New("errors when deleting note"))

	suite.noteController.deleteHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, suite.record.Code)
}

func TestNoteControllerTestSuite(t *testing.T) {
	suite.Run(t, new(NoteControllerTestSuite))
}
