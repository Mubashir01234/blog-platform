package models

type RegisterRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Username string `json:"username" bson:"username,omitempty"`
	FullName string `json:"full_name" bson:"full_name,omitempty"`
	Role     string `json:"role" bson:"role,omitempty"`
	Bio      string `json:"bio" bson:"bio,omitempty"`
}
