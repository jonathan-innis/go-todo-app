package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionStringTemplate = "mongodb://%s:%s@%s/%s"
)

func ConnectDB(ctx context.Context, config *Config) *mongo.Database {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf(connectionStringTemplate, config.username, config.password, config.endpoint, config.name))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client.Database(config.name)
}
