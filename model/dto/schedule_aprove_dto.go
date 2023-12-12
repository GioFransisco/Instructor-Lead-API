package dto

import "time"

type ScheduleAproveResponseDto struct {
	Id 					string		`json:"id"`
	ScheduleDetailsId 	string		`json:"scheduleDetailsId"`
	ScheduleAprove 		string		`json:"scheduleAprove"`
	CreatedAt 			time.Time	`json:"createdAt"`
	UpdatedAt 			time.Time	`json:"updatedAt"`
}