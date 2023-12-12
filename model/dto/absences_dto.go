package dto

import (
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
)

type AbsencesResponseDto struct {
	Id                string     `json:"id"`
	ScheduleDetailsId string     `json:"scheduleDetailsId"`
	StudentId         model.User `json:"userId"`
	Description       string     `json:"description"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}
