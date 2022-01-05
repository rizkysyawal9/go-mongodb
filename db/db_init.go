package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resource struct {
	DB *mongo.Database
}

func InitResource() (*Resource, error) {
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	dbName := os.Getenv("MONGO_DBNAME")
	dbUser := os.Getenv("MONGO_USER")
	dbPassword := os.Getenv("MONGO_PASSWORD")

	credential := options.Credential{
		Username: dbUser,
		Password: dbPassword,
	}

	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri).SetAuth(credential)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// If we don't call cancel it will cause a memory leak
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	return &Resource{
		DB: client.Database(dbName),
	}, nil
}
