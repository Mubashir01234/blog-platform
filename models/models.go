package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct is used in user endpoints and they shows all the entities of the user.
type User struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username,omitempty"`
	Role      string             `json:"role" bson:"role,omitempty"`
	Email     string             `json:"email" bson:"email,omitempty"`
	Password  string             `json:"password" bson:"password,omitempty"`
	FullName  string             `json:"full_name" bson:"full_name,omitempty"`
	Bio       string             `json:"bio" bson:"bio,omitempty"`
	IsNew     bool               `json:"is_new" bson:"is_new,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}

// Blog struct is used in blogs endpoint and they also shows all the entities of the blog.
type Blog struct {
	Id          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserId      string             `json:"user_id" bson:"user_id,omitempty"`
	Username    string             `json:"username" bson:"username,omitempty"`
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt   time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}
