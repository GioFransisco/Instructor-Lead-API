package model

import "time"

type ScheduleAprove struct {
	Id              string          `json:"id"`
	ScheduleDetails ScheduleDetails `json:"scheduleDetail"`
	ScheduleAprove  string          `json:"scheduleAprove"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}
