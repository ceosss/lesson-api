package modelhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateModel creates a new model
func CreateModel(response http.ResponseWriter, request *http.Request) {
	var model models.Model
	var err error
	err = json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	modelCollection := db.GetModelCollection(client)

	res, err := modelCollection.InsertOne(context.TODO(), model)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "Model Created Successfully"}`))
}

// AllModels ...
func AllModels(response http.ResponseWriter, request *http.Request) {
	var allModels []models.Model
	var err error

	client, err := db.ConnectToDB()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	modelCollection := db.GetModelCollection(client)

	cursor, err := modelCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cursor.Close(ctx)
	cancel()

	for cursor.Next(ctx) {
		var m models.Model
		cursor.Decode(&m)
		allModels = append(allModels, m)
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(allModels)
}

//SingleModel ...
func SingleModel(response http.ResponseWriter, request *http.Request) {
	var model models.Model
	var err error

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	modelCollection := db.GetModelCollection(client)
	filter := bson.M{"_id": id}
	modelCollection.FindOne(context.TODO(), filter).Decode(&model)

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model)
}

//DeleteModel ...
func DeleteModel(response http.ResponseWriter, request *http.Request) {
	var err error
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	client, err := db.ConnectToDB()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	modelCollection := db.GetModelCollection(client)

	filter := bson.M{"_id": id}

	res, err := modelCollection.DeleteOne(context.TODO(), filter)
	fmt.Printf("MODEL DELETED: %v", res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)
}

// UpdateModel ...
func UpdateModel(response http.ResponseWriter, request *http.Request) {
	var model models.Model
	var err error

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	err = json.NewDecoder(request.Body).Decode(&model)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	modelCollection := db.GetModelCollection(client)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": model}

	res, err := modelCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	fmt.Printf("MODEL UPDATED: %v", res)
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)

}
