package db

import (
	"blog/config"
	"blog/models"
	"context"
	"log"
	"net/http"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogDB interface {
	CheckEmailExistsDB(r *http.Request, collectionName, email string) (*models.User, error)
	CheckUsernameExistsDB(r *http.Request, collectionName, username string) (*models.User, error)
	RegisterDB(r *http.Request, collectionName string, user models.User) (*mongo.InsertOneResult, error)
	GetUserByEmailDB(r *http.Request, collectionName, email string) (*models.User, error)
	GetUserByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.User, error)
	UpdateUserDB(r *http.Request, collectionName string, user models.User) error
	DeleteProfileDB(r *http.Request, collectionName string, id primitive.ObjectID) error
	CreateBlogDB(r *http.Request, collectionName string, blog models.Blog) (*mongo.InsertOneResult, error)
	GetBlogByIdDB(r *http.Request, collectionName string, id primitive.ObjectID) (*models.Blog, error)
	UpdateBlogDB(r *http.Request, collectionName string, blog models.Blog) error
}

type BlogDBImpl struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

func ConnectDB() *BlogDBImpl {

	clientOptions := options.Client().ApplyURI(config.Cfg.MongoURL)
	var err error

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Connection Failed to Database")
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection Failed to Connect Mongo Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")

	collections := loadCollection(client)

	return &BlogDBImpl{
		Client:      client,
		Collections: collections,
	}
}

func loadCollection(mongoConn *mongo.Client) map[string]*mongo.Collection {
	collections := make(map[string]*mongo.Collection, 3)
	collections["users"] = colHelper(mongoConn, "users")
	collections["roles"] = colHelper(mongoConn, "roles")
	collections["blogs"] = colHelper(mongoConn, "blogs")
	return collections
}

func colHelper(db *mongo.Client, collectionName string) *mongo.Collection {
	return db.Database("blog").Collection(collectionName)
}

var _ BlogDB = &BlogDBImpl{}
