package db

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *BlogDBImpl) DeleteProfileDB(r *http.Request, collectionName string, id primitive.ObjectID) error {
	res, err := db.Collections[collectionName].DeleteOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("user doesn't exists")
	}
	return nil
}

func (db *BlogDBImpl) DeleteBlogDB(r *http.Request, collectionName string, id primitive.ObjectID) error {
	res, err := db.Collections[collectionName].DeleteOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("blog doesn't exists")
	}
	return nil
}
