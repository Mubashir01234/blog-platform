package db

import (
	"blog/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func (db *BlogDBImpl) RegisterDB(r *http.Request, collectionName string, user models.User) (*mongo.InsertOneResult, error) {
	result, err := db.Collections[collectionName].InsertOne(r.Context(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (db *BlogDBImpl) CreateBlogDB(r *http.Request, collectionName string, blog models.Blog) (*mongo.InsertOneResult, error) {
	result, err := db.Collections[collectionName].InsertOne(r.Context(), blog)
	if err != nil {
		return nil, err
	}
	return result, nil
}
