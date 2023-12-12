package model

import "time"

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	Username    string    `json:"username"`
	Password    string    `json:"password,omitempty"`
	Age         int       `json:"age"`
	Address     string    `json:"address"`
	Gander      string    `json:"gander"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (u *User) IsValidRole() bool {
	return u.Role == "Admin" || u.Role == "Participant" || u.Role == "Trainer"
}

func (u *User) IsValidGender() bool {
	return u.Gander == "L" || u.Gander == "P"
}
