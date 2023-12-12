package dto

import "time"

type StackRequestDto struct {
	Id        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
