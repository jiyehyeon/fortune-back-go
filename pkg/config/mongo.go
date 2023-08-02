package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	URI     	string
	Database 	string
	// Username string `json:"username"`
	// Password string `json:"password"`
}

func GetMongoDBConfig() *MongoDBConfig{
	return &MongoDBConfig{
		URI: os.Getenv("MONGO_URI"),
		Database: os.Getenv("MONGO_DATABASE"),
	}
}

func (config *MongoDBConfig) Connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(config.URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Database)
	return db, nil
}