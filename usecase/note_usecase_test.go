package usecase

import (
	
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	repositorymock "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/mock/repository_mock"
	
	
)

func TestNoteUseCase_Create(t *testing.T) {
	noteRepoMock := new(repositorymock.NoteRepoMock)
	noteUseCase := NewNoteUseCase(noteRepoMock)

	payload := model.Note{
		ScheduleDetails: model.ScheduleDetails{Id: "1b39443e-9bd4-4307-a301-2696f267117f"},
		UserEmail:       "user@example.com",
		Note:            "This is a note",
	}

	dtoNote := dto.NoteDTO{
		ScheduleID: "1b39443e-9bd4-4307-a301-2696f267117f",
		UserEmail:  "user@example.com",
		Note:       "This is a note",
	}

	noteRepoMock.On("Create", mock.Anything).Return(dtoNote, nil)

	createdNote, err := noteUseCase.Create(payload)
	assert.NoError(t, err)
	assert.Equal(t, dtoNote, createdNote)
}
