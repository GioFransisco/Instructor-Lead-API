package usecase

import (
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/repository"
	utils "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"

)

type NoteUseCase interface {
	FindAll() ([]dto.NoteDTO, error)
	Create(payload model.Note) (dto.NoteDTO, error)
	Update(id string, payload model.Note) (dto.NoteDTO, error)
	FindByID(id string) (dto.NoteDTO, error)
	Delete(id string) error
}

type noteUseCase struct {
	noteRepo repository.NoteRepository
}

func NewNoteUseCase(noteRepo repository.NoteRepository) NoteUseCase {
	return &noteUseCase{noteRepo: noteRepo}
}

func (n *noteUseCase) Create(payload model.Note) (dto.NoteDTO, error) {
	dtoNote := dto.ConvertNoteToDTO(payload)
	createdNote, err := n.noteRepo.Create(dtoNote)
	if err != nil {
		return dto.NoteDTO{},  utils.InternalServerError
	}
	return createdNote, nil
}

func (n *noteUseCase) Update(id string, payload model.Note) (dto.NoteDTO, error) {
	updatedNote, err := n.noteRepo.Update(id, payload)
	if err != nil {
		return dto.NoteDTO{}, err
	}
	return dto.ConvertUpdatedNoteToDTO(updatedNote), nil
}

func (n *noteUseCase) FindAll() ([]dto.NoteDTO, error) {
	return n.noteRepo.List()
}

func (n *noteUseCase) FindByID(id string) (dto.NoteDTO, error) {
	return n.noteRepo.FindByID(id)
}

func (n *noteUseCase) Delete(id string) error {
	return n.noteRepo.Delete(id)
}
