package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Role     string             `json:"role" bson:"role,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
	FullName string             `json:"full_name" bson:"full_name,omitempty"`
	Bio      string             `json:"bio" bson:"bio,omitempty"`
	IsNew    bool               `json:"is_new" bson:"is_new,omitempty"`
}
