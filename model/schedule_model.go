package model

import "time"

type Schedule struct {
	Id              string            `json:"id"`
	Name            string            `json:"name"`
	DateActivity    time.Time         `json:"dateActivity"`
	ScheduleDetails []ScheduleDetails `json:"scheduleDetails,omitempty"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}



type ScheduleDetails struct {
	Id         string    `json:"id"`
	ScheduleId string    `json:"scheduleId,omitempty"`
	Trainer    User      `json:"trainer"`
	Stack      Stack     `json:"stack"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type GetScheduleDetails struct {
	Id         string    	`json:"id"`
	Schedule   Schedule  	`json:"scheduleId,omitempty"`
	Trainer    User      	`json:"trainer"`
	Stack      Stack     	`json:"stack"`
	StartTime  time.Time 	`json:"startTime"`
	EndTime    time.Time 	`json:"endTime"`
	CreatedAt  time.Time 	`json:"createdAt"`
	UpdatedAt  time.Time 	`json:"updatedAt"`
}
