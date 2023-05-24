package models

// RegisterRequest is used for the registration request
type RegisterRequest struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest is used for login a user
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateProfileRequest is used for updating the profile of a user
type UpdateProfileRequest struct {
	Username string `json:"username" bson:"username,omitempty"`
	FullName string `json:"full_name" bson:"full_name,omitempty"`
	Role     string `json:"role" bson:"role,omitempty"`
	Bio      string `json:"bio" bson:"bio,omitempty"`
}

// CreateBlogRequest is used to create a new blog
type CreateBlogRequest struct {
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}

// UpdateBlogRequest is used to update the blog
type UpdateBlogRequest struct {
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}
