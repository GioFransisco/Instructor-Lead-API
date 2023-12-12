package model

import "time"

type Question struct {
	Id              string          `json:"id"`
	ScheduleDetails ScheduleDetails `json:"scheduleDetail"`
	StudentId       User            `json:"student"`
	Question        string          `json:"question"`
	Status          string          `json:"status"`
	CreatedAt       time.Time       `json:"createdAt"`
	UpdatedAt       time.Time       `json:"updatedAt"`
}

func (q *Question) IsValidate() bool {
	return q.Status == "Finish" || q.Status == "Proccess"
}
