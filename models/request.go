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

type CreateBlogRequest struct {
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}

type UpdateBlogRequest struct {
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}
