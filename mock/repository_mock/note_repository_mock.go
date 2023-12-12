package repositorymock

import (
	
	"github.com/stretchr/testify/mock"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"

)

type NoteRepoMock struct {
	mock.Mock
}

func (m *NoteRepoMock) Create(payload dto.NoteDTO) (dto.NoteDTO, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.NoteDTO), args.Error(1)
}

func (m *NoteRepoMock) Update(id string, note model.Note) (model.Note, error) {
	args := m.Called(id, note)
	return args.Get(0).(model.Note), args.Error(1)
}

func (m *NoteRepoMock) List() ([]dto.NoteDTO, error) {
	args := m.Called()
	return args.Get(0).([]dto.NoteDTO), args.Error(1)
}

func (m *NoteRepoMock) FindByID(id string) (dto.NoteDTO, error) {
	args := m.Called(id)
	return args.Get(0).(dto.NoteDTO), args.Error(1)
}

func (m *NoteRepoMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func NewNoteRepoMock() *NoteRepoMock {
	return &NoteRepoMock{}
}
