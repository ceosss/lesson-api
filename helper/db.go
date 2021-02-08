package helper

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return nil, err
	}
	return client, nil
}

func getModelCollection(client *mongo.Client) *mongo.Collection {
	modelCollection := client.Database("lesson-api").Collection("model")
	return modelCollection
}
