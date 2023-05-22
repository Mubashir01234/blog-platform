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

// BlogDB defines the interface for interacting with the database
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
	DeleteBlogDB(r *http.Request, collectionName string, id primitive.ObjectID) error
	GetBlogsByUsernameDB(r *http.Request, collectionName, username string) ([]*models.GetBlogResp, error)
}

// BlogDBImpl is an implementation of the BlogDB interface
type BlogDBImpl struct {
	Client      *mongo.Client
	Collections map[string]*mongo.Collection
}

// ConnectDB establishes a connection to the MongoDB database and returns a BlogDBImpl instance
func ConnectDB() *BlogDBImpl {
	// Configure MongoDB client options
	clientOptions := options.Client().ApplyURI(config.Cfg.MongoURL)
	var err error

	// Connect to MongoDB
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

	// Load collections
	collections := loadCollection(client)

	return &BlogDBImpl{
		Client:      client,
		Collections: collections,
	}
}

// loadCollection initializes the map of collection names to collection instances
func loadCollection(mongoConn *mongo.Client) map[string]*mongo.Collection {
	collections := make(map[string]*mongo.Collection, 2)
	collections["users"] = colHelper(mongoConn, "users")
	collections["blogs"] = colHelper(mongoConn, "blogs")
	return collections
}

// colHelper helps in retrieving a specific collection from the MongoDB database using the provided collection name
func colHelper(db *mongo.Client, collectionName string) *mongo.Collection {
	return db.Database("blog").Collection(collectionName)
}

// Ensure BlogDBImpl implements the BlogDB interface
var _ BlogDB = &BlogDBImpl{}
