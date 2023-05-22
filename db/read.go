package db

import (
	"blog/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *BlogDBImpl) GetUserDataByEmailDB(r *http.Request, collectionName, email string) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err == nil {
		return nil, fmt.Errorf("email already exists")
	}
	return &existingUser, nil
}

func (db *BlogDBImpl) GetUserDataByUsernameDB(r *http.Request, collectionName, username string) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "username", Value: username}}).Decode(&existingUser)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}
	return &existingUser, nil
}
