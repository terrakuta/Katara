package database

import (
	"Katara/internal/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func MongoLoad(cfg *config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(cfg.MONGO_URI))

	if err != nil {
		return nil, err
	}

	pingErr := client.Ping(ctx, readpref.Primary())

	if pingErr != nil {
		return nil, pingErr
	}

	log.Println("Successfully Connected to db")

	return client, nil

}
