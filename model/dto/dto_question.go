package dto

import (
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
)

type QuestionChangeDto struct {
	Id       string `json:"id"`
	Question string `json:"question"`
}

type QuestionChangeStatusDto struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type QuestionResponseUpdate struct {
	Id              string    `json:"id"`
	ScheduleDetails string    `json:"scheduleDetailId"`
	StudentId       string    `json:"studentId"`
	Question        string    `json:"question"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type QuestionResponseGET struct {
	Id              string     `json:"id"`
	ScheduleDetails string     `json:"scheduleDetailId"`
	StudentId       model.User `json:"user"`
	Question        string     `json:"question"`
	Status          string     `json:"status"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}
