package model

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id              *uuid.UUID      `json:"id"`
	ScheduleDetails ScheduleDetails `json:"scheduleDetails"`
	UserEmail       string          `json:"userEmail"`
	Note            string          `json:"note"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}
