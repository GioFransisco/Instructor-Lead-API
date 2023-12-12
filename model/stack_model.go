package model

import "time"

type Stack struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (s *Stack) IsValidStatus() bool {
	return s.Status == "Active" || s.Status == "Inactive"
}
