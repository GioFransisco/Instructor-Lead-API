package usecasemock

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"

	"github.com/stretchr/testify/mock"
)

type NoteUseCaseMock struct {
	mock.Mock
}

func (m *NoteUseCaseMock) Create(payload model.Note) (dto.NoteDTO, error) {
	args := m.Called(payload)
	return args.Get(0).(dto.NoteDTO), args.Error(1)
}

func (m *NoteUseCaseMock) Update(id string, note model.Note) (dto.NoteDTO, error) {
	args := m.Called(id, note)
	return args.Get(0).(dto.NoteDTO), args.Error(1)
}

func (m *NoteUseCaseMock) FindAll() ([]dto.NoteDTO, error) {
	args := m.Called()
	return args.Get(0).([]dto.NoteDTO), args.Error(1)
}

func (m *NoteUseCaseMock) FindByID(id string) (dto.NoteDTO, error) {
	args := m.Called(id)
	return args.Get(0).(dto.NoteDTO), args.Error(1)
}

func (m *NoteUseCaseMock) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
