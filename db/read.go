package db

import (
	"blog/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (db *BlogDBImpl) GetBlogByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.Blog, error) {
	var existedBlog models.Blog

	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&existedBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("blog doesn't exists")
		}
		return nil, err
	}
	return &existedBlog, nil
}

func (db *BlogDBImpl) GetBlogsByUsernameDB(r *http.Request, collectionName, username string) ([]*models.GetBlogResp, error) {
	var blogs []*models.GetBlogResp

	opts := options.Find().SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	cursor, err := db.Collections[collectionName].Find(r.Context(), bson.D{primitive.E{Key: "username", Value: username}}, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	for cursor.Next(r.Context()) {
		var blog models.GetBlogResp
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}
	return blogs, nil
}
