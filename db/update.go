package db

import (
	"blog/models" // Importing the models package which likely contains data models for users and blogs
	"fmt"         // Importing the fmt package for error handling
	"net/http"    // Importing the net/http package for handling HTTP requests

	"go.mongodb.org/mongo-driver/bson"           // Importing the BSON package from MongoDB driver
	"go.mongodb.org/mongo-driver/bson/primitive" // Importing the primitive package from MongoDB driver
)

// UpdateUserDB updates a user document in the specified MongoDB collection.
func (db *BlogDBImpl) UpdateUserDB(r *http.Request, collectionName string, user models.User) error {
	// Update the user document in the collection
	res, err := db.Collections[collectionName].UpdateOne(
		r.Context(), // MongoDB context
		bson.D{primitive.E{Key: "_id", Value: user.Id}}, // Filter for matching user document by ID
		bson.D{ // Fields to be updated using the $set operator
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "username", Value: user.Username},
				primitive.E{Key: "full_name", Value: user.FullName},
				primitive.E{Key: "role", Value: user.Role},
				primitive.E{Key: "bio", Value: user.Bio},
				primitive.E{Key: "updated_at", Value: user.UpdatedAt},
				primitive.E{Key: "is_new", Value: false},
			}},
		},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		// No user found with the provided ID
		return fmt.Errorf("user doesn't exist")
	}

	return nil
}

// UpdateBlogDB updates a blog document in the specified MongoDB collection.
func (db *BlogDBImpl) UpdateBlogDB(r *http.Request, collectionName string, blog models.Blog) error {
	// Update the blog document in the collection
	res, err := db.Collections[collectionName].UpdateOne(
		r.Context(), // MongoDB context
		bson.D{primitive.E{Key: "_id", Value: blog.Id}}, // Filter for matching blog document by ID
		bson.D{ // Fields to be updated using the $set operator
			primitive.E{Key: "$set", Value: bson.D{
				primitive.E{Key: "title", Value: blog.Title},
				primitive.E{Key: "description", Value: blog.Description},
				primitive.E{Key: "updated_at", Value: blog.UpdatedAt},
			}},
		},
	)
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		// No blog found with the provided ID
		return fmt.Errorf("blog doesn't exist")
	}
	// Return nil if no error found
	return nil
}
