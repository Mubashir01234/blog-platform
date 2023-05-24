package db

import (
	"blog/models"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo" // Importing mongo from the MongoDB driver
)

// RegisterDB inserts a user document into the specified MongoDB collection.
func (db *BlogDBImpl) RegisterDB(r *http.Request, collectionName string, user models.User) (*mongo.InsertOneResult, error) {
	// Insert the user document into the collection
	result, err := db.Collections[collectionName].InsertOne(r.Context(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateBlogDB inserts a blog document into the specified MongoDB collection.
func (db *BlogDBImpl) CreateBlogDB(r *http.Request, collectionName string, blog models.Blog) (*mongo.InsertOneResult, error) {
	// Insert the blog document into the collection
	result, err := db.Collections[collectionName].InsertOne(r.Context(), blog)
	if err != nil {
		return nil, err
	}
	return result, nil
}
