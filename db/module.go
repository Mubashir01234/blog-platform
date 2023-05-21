package db

import (
	"blog/config"
	"context"
	"log"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogDB interface {
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
	collections["user"] = colHelper(mongoConn, "user")
	collections["role"] = colHelper(mongoConn, "role")
	collections["blog"] = colHelper(mongoConn, "blog")
	return collections
}

func colHelper(db *mongo.Client, collectionName string) *mongo.Collection {
	return db.Database("blog").Collection(collectionName)
}

var _ BlogDB = &BlogDBImpl{}
