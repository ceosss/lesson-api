package modelhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/cookiehandler"
	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateModel - Creates a new model
func CreateModel(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var model models.Model

	err = json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	v := validator.New()

	err = v.Struct(model)

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	modelCollection := db.GetModelCollection(client)

	res, err := modelCollection.InsertOne(context.TODO(), model)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "Model Created Successfully"}`))
}

// AllModels - Fetches all the models
func AllModels(response http.ResponseWriter, request *http.Request) {

	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var allModels []models.Model

	client, err := db.ConnectToDB()

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	modelCollection := db.GetModelCollection(client)

	cursor, err := modelCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
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

//SingleModel - Fetches a specific model
func SingleModel(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var model models.Model

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	modelCollection := db.GetModelCollection(client)
	filter := bson.M{"_id": id}
	modelCollection.FindOne(context.TODO(), filter).Decode(&model)

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(model)
}

//DeleteModel - Deletes a specific model
func DeleteModel(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	client, err := db.ConnectToDB()
	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	modelCollection := db.GetModelCollection(client)

	filter := bson.M{"_id": id}

	res, err := modelCollection.DeleteOne(context.TODO(), filter)
	fmt.Printf("MODEL DELETED: %v", res)
	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)
}

// UpdateModel - Updates a specific model
func UpdateModel(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var model models.Model

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&model)

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	v := validator.New()

	err = v.Struct(model)

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	modelCollection := db.GetModelCollection(client)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": model}

	res, err := modelCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}
	fmt.Printf("MODEL UPDATED: %v", res)
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)

}
