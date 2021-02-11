package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ConnectToDB - Creates a connection to DB
func ConnectToDB() (*mongo.Client, error) {
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

//GetModelCollection - Connects to the Model Collection
func GetModelCollection(client *mongo.Client) *mongo.Collection {
	modelCollection := client.Database("lesson-api").Collection("model")
	return modelCollection
}

//GetLessonCollection - Connects to the Lesson Collection
func GetLessonCollection(client *mongo.Client) *mongo.Collection {
	lessonCollection := client.Database("lesson-api").Collection("lesson")
	return lessonCollection
}

//GetUserCollection - Connects to the User Collection
func GetUserCollection(client *mongo.Client) *mongo.Collection {
	userCollection := client.Database("lesson-api").Collection("user")
	return userCollection
}
