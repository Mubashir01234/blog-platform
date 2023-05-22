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

// CheckEmailExistsDB checks if an email already exists in the specified collection
// Returns the existing user if found, otherwise returns an error
func (db *BlogDBImpl) CheckEmailExistsDB(r *http.Request, collectionName, email string) (*models.User, error) {
	var existingUser models.User

	// Find a document with the provided email
	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err == nil {
		// Email already exists
		return nil, fmt.Errorf("email already exists")
	}
	// Email doesn't exist
	return &existingUser, nil
}

// CheckUsernameExistsDB checks if a username already exists in the specified collection
// Returns the existing user if found, otherwise returns an error
func (db *BlogDBImpl) CheckUsernameExistsDB(r *http.Request, collectionName, username string) (*models.User, error) {
	var existingUser models.User

	// Find a document with the provided username
	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "username", Value: username}}).Decode(&existingUser)
	if err == nil {
		// Username already exists
		return nil, fmt.Errorf("username already exists")
	}
	// Username doesn't exist
	return &existingUser, nil
}

// GetUserByEmailDB retrieves a user by their email from the specified collection
// Returns the user if found, otherwise returns an error
func (db *BlogDBImpl) GetUserByEmailDB(r *http.Request, collectionName, email string) (*models.User, error) {
	var existingUser models.User

	// Find a document with the provided email
	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "email", Value: email}}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Email doesn't exist
			return nil, fmt.Errorf("email doesn't exist")
		}
		// Other error occurred
		return nil, err
	}
	// User found
	return &existingUser, nil
}

func (db *BlogDBImpl) GetUserByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.User, error) {
	var existingUser models.User

	// Find a document with the provided ID
	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// User doesn't exist
			return nil, fmt.Errorf("user doesn't exist")
		}
		// Other error occurred
		return nil, err
	}
	// User found
	return &existingUser, nil
}

func (db *BlogDBImpl) GetBlogByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.Blog, error) {
	var existedBlog models.Blog

	// Find a document with the provided ID
	err := db.Collections[collectionName].FindOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&existedBlog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Blog doesn't exist
			return nil, fmt.Errorf("blog doesn't exist")
		}
		// Other error occurred
		return nil, err
	}
	// Blog found
	return &existedBlog, nil
}

func (db *BlogDBImpl) GetBlogsByUsernameDB(r *http.Request, collectionName, username string) ([]*models.GetBlogResp, error) {
	var blogs []*models.GetBlogResp

	// Set the options for sorting by "created_at" in descending order
	opts := options.Find().SetSort(bson.D{primitive.E{Key: "created_at", Value: -1}})

	// Find all documents with the provided username and apply the options
	cursor, err := db.Collections[collectionName].Find(r.Context(), bson.D{primitive.E{Key: "username", Value: username}}, opts)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No blogs found for the username
			return nil, nil
		}
		// Other error occurred
		return nil, err
	}

	// Iterate over the cursor and decode each document into a blog
	for cursor.Next(r.Context()) {
		var blog models.GetBlogResp
		err := cursor.Decode(&blog)
		if err != nil {
			return nil, err
		}

		blogs = append(blogs, &blog)
	}
	// Blogs found
	return blogs, nil
}
