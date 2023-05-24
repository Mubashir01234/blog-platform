package db

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"           // Importing the BSON package from MongoDB driver
	"go.mongodb.org/mongo-driver/bson/primitive" // Importing the primitive package from MongoDB driver
)

// DeleteProfileDB deletes a user profile from the specified collection by ID.
func (db *BlogDBImpl) DeleteProfileDB(r *http.Request, collectionName string, id primitive.ObjectID) error {
	res, err := db.Collections[collectionName].DeleteOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("user doesn't exist")
	}
	return nil
}

// DeleteBlogDB deletes a blog from the specified collection by ID.
func (db *BlogDBImpl) DeleteBlogDB(r *http.Request, collectionName string, id primitive.ObjectID) error {
	res, err := db.Collections[collectionName].DeleteOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("blog doesn't exist")
	}
	return nil
}
