package repository

import (
	"database/sql"
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model/dto"
	utilsmodel "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/utils_model"
	"github.com/google/uuid"
	utils "git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/utils/common"
)

type NoteRepository interface {
	Create(payload dto.NoteDTO) (dto.NoteDTO, error)
	List() ([]dto.NoteDTO, error)
	Update(id string, payload model.Note) (model.Note, error)
	FindByID(id string) (dto.NoteDTO, error)
	Delete(id string) error
}

type noteRepository struct {
	db *sql.DB
}

func (n *noteRepository) Create(payload dto.NoteDTO) (dto.NoteDTO, error) {
	if payload.ID == nil {

		newID := uuid.New()
		payload.ID = &newID
	}

	query := utilsmodel.NoteCreate
	err := n.db.QueryRow(query, payload.ID, payload.ScheduleID, payload.UserEmail, payload.Note, time.Now(), time.Now()).Scan(&payload.ID, &payload.ScheduleID, &payload.UserEmail, &payload.Note, &payload.CreatedAt, &payload.UpdatedAt)

	if err != nil {
		return dto.NoteDTO{}, utils.BadRequestError
	}

	return payload, nil

	
}

func (n *noteRepository) Update(id string, note model.Note) (model.Note, error) {
	var updatedNote model.Note

	query := utilsmodel.NoteUpdate
	err := n.db.QueryRow(query, note.Note, time.Now(), id).
		Scan(&updatedNote.Id, &updatedNote.UserEmail, &updatedNote.Note, &updatedNote.ScheduleDetails.Id, &updatedNote.CreatedAt, &updatedNote.UpdatedAt)

	if err != nil {
		return model.Note{}, utils.NotFoundError
	}

	return updatedNote, nil
}

func (n *noteRepository) List() ([]dto.NoteDTO, error) {
	rows, err := n.db.Query(utilsmodel.NoteList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var noteResponses []dto.NoteDTO
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.Id, &note.ScheduleDetails.Id, &note.Note, &note.UserEmail, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, err
		}

		dtoNote := dto.ConvertNoteToDTO(note)
		noteResponses = append(noteResponses, dtoNote)
	}

	return noteResponses, nil
}

func NewNoteRepository(db *sql.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (s *noteRepository) FindByID(id string) (dto.NoteDTO, error) {
	var note model.Note

	row := s.db.QueryRow(utilsmodel.NoteFindById, id)
	err := row.Scan(&note.Id, &note.ScheduleDetails.Id, &note.Note, &note.UserEmail, &note.CreatedAt, &note.UpdatedAt)

	if err != nil {
		return dto.NoteDTO{}, utils.NotFoundErrorByID
	}

	return dto.ConvertNoteToDTO(note), nil
}

func (n *noteRepository) Delete(id string) error {
	_, err := n.db.Exec(utilsmodel.DeleteNoteById, id)
	return err
}
