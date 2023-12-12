package dto

type UserUpdateDto struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	Gender      string `json:"gender"`
}

type UserLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserChangePassDto struct {
	Password string `json:"password"`
}
