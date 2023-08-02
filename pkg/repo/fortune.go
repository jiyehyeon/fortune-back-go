package repo

import (
	"context"
	"fortune-back-go/pkg/config"
	"fortune-back-go/pkg/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FortuneRepo struct {
	db *mongo.Database
}

func NewFortuneRepo() *FortuneRepo {
	mongoConfig := config.GetMongoDBConfig()

	db, err := mongoConfig.Connect()

	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	return &FortuneRepo{
		db: db,
	}
}

func (r *FortuneRepo) GetFortune(ganji string) (*model.Fortune, error) {
	filter := bson.M{
		"ganji": ganji,
}

	var fortune model.Fortune
	err := r.db.Collection("fortunes").FindOne(context.Background(), filter).Decode(&fortune)

	if err != nil {
		log.Println("Failed to get fortune:", err)
		return nil, err
	}

	return &fortune, nil
}