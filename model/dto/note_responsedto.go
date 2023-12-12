
package dto

import (
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
	"github.com/google/uuid"
)

type NoteDTO struct {
	ID         *uuid.UUID `json:"id"`
	ScheduleID string     `json:"scheduleDetailId"`
	UserEmail  string     `json:"userEmail"`
	Note       string     `json:"note"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
}

func ConvertNoteToDTO(note model.Note) NoteDTO {
	return NoteDTO{
		ID:         note.Id,
		ScheduleID: note.ScheduleDetails.Id,
		UserEmail:  note.UserEmail,
		Note:       note.Note,
		CreatedAt:  note.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  note.UpdatedAt.Format(time.RFC3339),
	}
}


func ConvertUpdatedNoteToDTO(note model.Note) NoteDTO {
	return ConvertNoteToDTO(note)
}
