package modelhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/models"
)

// CreateModel creates a new model
func CreateModel(response http.ResponseWriter, request *http.Request) {
	var model models.Model
	var err error
	err = json.NewDecoder(request.Body).Decode(&model)
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	client, err := db.ConnectToDB()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	modelCollection := db.GetModelCollection(client)

	res, err := modelCollection.InsertOne(context.TODO(), model)
	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "Model Created Successfully"}`))
}
