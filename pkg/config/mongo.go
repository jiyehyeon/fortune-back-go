package config

import (
	"context"
	"fmt"
	"fortune-back-go/pkg/utils"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	Host     	string
	Port     	int
	Database 	string
	// Username string `json:"username"`
	// Password string `json:"password"`
}

func GetMongoDBConfig() *MongoDBConfig{
	return &MongoDBConfig{
		Host: os.Getenv("MONGO_HOST"),
		Port: utils.StrToInt(os.Getenv("MONGO_PORT")),
		Database: os.Getenv("MONGO_DATABASE"),
	}
}

func (config *MongoDBConfig) Connect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port))

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(config.Database)
	return db, nil
}