package lessonhandlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/helper/initializemodels"
	"github.com/ceosss/lesson-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLesson - Creates a new Lesson with the name provided
func CreateLesson(response http.ResponseWriter, request *http.Request) {
	type n struct {
		Name string
	}
	var name n
	var err error

	err = json.NewDecoder(request.Body).Decode(&name)

	if err != nil {
		fmt.Printf("ERROR: %v", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	if len(name.Name) < 3 {
		fmt.Printf("ERROR: Invalid Name length")
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "Invalid Name length"}`))
		return
	}

	lesson := initializemodels.NewLesson(name.Name)

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

	id := res.InsertedID.(primitive.ObjectID).Hex()

	fmt.Println(id)

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(201)
	response.Write([]byte(`{  "response": "` + id + `"}`))
}

// AllLessons - Fetches all the Lessons
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()

	for cursor.Next(ctx) {
		var l models.Lesson
		cursor.Decode(&l)
		allLessons = append(allLessons, l)
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(200)
	json.NewEncoder(response).Encode(allLessons)
}

//SingleLesson - Fetches a specific lesson
func SingleLesson(response http.ResponseWriter, request *http.Request) {
	var lesson models.Lesson
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

	lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}
	lessonCollection.FindOne(context.TODO(), filter).Decode(&lesson)

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(lesson)
}

//DeleteLesson - Deletes a specific lesson
func DeleteLesson(response http.ResponseWriter, request *http.Request) {
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

	lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}

	res, err := lessonCollection.DeleteOne(context.TODO(), filter)
	fmt.Printf("LESSON DELETED: %v", res)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)
}

// UpdateLesson - Updates a specific lesson
func UpdateLesson(response http.ResponseWriter, request *http.Request) {
	var lesson models.Lesson
	var err error

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}

	err = json.NewDecoder(request.Body).Decode(&lesson)

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

	lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": lesson}

	res, err := lessonCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "` + err.Error() + `"}`))
		return
	}
	fmt.Printf("LESSON UPDATED: %v", res)
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(204)

}
