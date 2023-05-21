package models

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type UserDataResponse struct {
	UserId      *int64  `json:"user_id" db:"user_id"`
	Name        *string `json:"name" db:"name"`
	Email       *string `json:"email" db:"email"`
	PhoneNumber *string `json:"phone_number" db:"phone_number"`
}

func NewUserDataResponse() *UserDataResponse {
	return &UserDataResponse{
		UserId:      new(int64),
		Name:        new(string),
		Email:       new(string),
		PhoneNumber: new(string),
	}
}
