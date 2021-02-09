package lessonhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateLesson ...
func CreateLesson(response http.ResponseWriter, request *http.Request) {
	var lesson models.Lesson
	var err error

	err = json.NewDecoder(request.Body).Decode(&lesson)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusBadRequest)
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

	lessonCollection := db.GetLessonCollection(client)

	res, err := lessonCollection.InsertOne(context.TODO(), lesson)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	res.InsertedID = ""

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "Lesson Created Successfully"}`))
}

// AllLessons ...
func AllLessons(response http.ResponseWriter, request *http.Request) {
	var allLessons []models.Lesson
	var err error

	client, err := db.ConnectToDB()

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	lessonCollection := db.GetLessonCollection(client)

	cursor, err := lessonCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var l models.Lesson
		cursor.Decode(&l)
		allLessons = append(allLessons, l)
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(allLessons)
}
