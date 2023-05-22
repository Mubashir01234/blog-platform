package db

import (
	"blog/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *BlogDBImpl) CheckEmailExistsDB(r *http.Request, collectionName, email string) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err == nil {
		return nil, fmt.Errorf("email already exists")
	}
	return &existingUser, nil
}

func (db *BlogDBImpl) CheckUsernameExistsDB(r *http.Request, collectionName, username string) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "username", Value: username}}).Decode(&existingUser)
	if err == nil {
		return nil, fmt.Errorf("username already exists")
	}
	return &existingUser, nil
}

func (db *BlogDBImpl) GetUserByEmailDB(r *http.Request, collectionName, email string) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("email doesn't exists")
		}
		return nil, err
	}
	return &existingUser, nil
}

func (db *BlogDBImpl) GetUserByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.User, error) {
	var existingUser models.User

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user doesn't exists")
		}
		return nil, err
	}
	return &existingUser, nil
}
