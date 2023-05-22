package db

import (
	"blog/models"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *BlogDBImpl) UpdateUserDB(r *http.Request, collectionName string, user models.User) error {

	res, err := db.Collections[collectionName].UpdateOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: user.Id}}, bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "username", Value: user.Username},
				primitive.E{Key: "full_name", Value: user.FullName},
				primitive.E{Key: "role", Value: user.Role},
				primitive.E{Key: "bio", Value: user.Bio},
				primitive.E{Key: "updated_at", Value: user.UpdatedAt},
				primitive.E{Key: "is_new", Value: false},
			},
		},
	})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("user doesn't exist")
	}
	return nil
}

func (db *BlogDBImpl) UpdateBlogDB(r *http.Request, collectionName string, blog models.Blog) error {

	res, err := db.Collections[collectionName].UpdateOne(r.Context(), bson.D{primitive.E{Key: "_id", Value: blog.Id}}, bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "title", Value: blog.Title},
				primitive.E{Key: "description", Value: blog.Description},
				primitive.E{Key: "updated_at", Value: blog.UpdatedAt},
			},
		},
	})
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("blog doesn't exist")
	}
	return nil
}
