package model

import "time"

type Absences struct {
	Id              	string          	`json:"id"`
	ScheduleDetails 	ScheduleDetails 	`json:"scheduleDetail"`
	StudentId       	User            	`json:"student"`
	Description     	string          	`json:"description"`
	CreatedAt       	time.Time       	`json:"createdAt"`
	UpdatedAt       	time.Time       	`json:"updatedAt"`
}

type GetAbsences struct {
	Id              	string          		`json:"id"`
	Student       		User            		`json:"student"`
	ScheduleDetails 	[]GetScheduleDetails 	`json:"scheduleDetail"`
	Description     	string          		`json:"description"`
	CreatedAt       	time.Time       		`json:"createdAt"`
	UpdatedAt       	time.Time       		`json:"updatedAt"`
}