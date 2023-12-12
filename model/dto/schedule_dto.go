package dto

import (
	"time"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-12-golang/final-project/group-1/instructor-led-apps/model"
)

type ScheduleCreateRequestDto struct {
	Name            string                           `json:"name" binding:"required"`
	DateActivity    string                           `json:"dateActivity" binding:"required"`
	ScheduleDetails []ScheduleDetailCreateRequestDto `json:"scheduleDetails" binding:"required,dive,required"`
}

type ScheduleDetailCreateRequestDto struct {
	TrainerId string `json:"trainerId" binding:"required"`
	StackId   string `json:"stackId" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime   string `json:"endTime" binding:"required"`
}

type ScheduleUpdateRequestDto struct {
	Name         string `json:"name"`
	DateActivity string `json:"dateActivity"`
}

type ScheduleDetailUpdateRequestDto struct {
	TrainerId string `json:"trainerId"`
	StackId   string `json:"stackId"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type ScheduleResponseDto struct {
	Id              string                      `json:"id"`
	Name            string                      `json:"name"`
	DateActivity    string                      `json:"dateActivity"`
	ScheduleDetails []ScheduleDetailResponseDto `json:"scheduleDetails,omitempty"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}

type ScheduleDetailResponseDto struct {
	Id         string      `json:"id"`
	ScheduleId string      `json:"scheduleId,omitempty"`
	Trainer    model.User  `json:"trainer"`
	Stack      model.Stack `json:"stack"`
	StartTime  string      `json:"startTime"`
	EndTime    string      `json:"endTime"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}
