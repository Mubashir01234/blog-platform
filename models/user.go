package models

type UserData struct {
	UserId      *int64  `json:"user_id" db:"user_id"`
	Name        *string `json:"name" db:"name"`
	Email       *string `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
}

func NewUserData() *UserData {
	return &UserData{
		UserId:      new(int64),
		Name:        new(string),
		Email:       new(string),
		PhoneNumber: new(string),
	}
}
