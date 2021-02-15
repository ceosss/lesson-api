package lessonhandlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ceosss/lesson-api/helper/cookiehandler"
	"github.com/ceosss/lesson-api/helper/customerror"
	"github.com/ceosss/lesson-api/helper/db"
	"github.com/ceosss/lesson-api/helper/initializemodels"
	"github.com/ceosss/lesson-api/helper/successresponse"
	"github.com/ceosss/lesson-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLesson - Creates a new Lesson with the name provided
func CreateLesson(response http.ResponseWriter, request *http.Request) {

	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	type n struct {
		Name string
	}
	var name n

	err = json.NewDecoder(request.Body).Decode(&name)

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	if len(name.Name) < 3 {
		customerror.BadRequest(&response, errors.New("Invalid Name"))
		return
	}

	lesson := initializemodels.NewLesson(name.Name)

	// client, err := db.ConnectToDB()

	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// lessonCollection := db.GetLessonCollection(client)

	res, err := lessonCollection.InsertOne(context.TODO(), lesson)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	fmt.Println(id)

	successresponse.Created(&response)
	response.Write([]byte(`{  "response": "` + id + `"}`))
}

// AllLessons - Fetches all the Lessons
func AllLessons(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var allLessons []models.Lesson

	// client, err := db.ConnectToDB()

	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// lessonCollection := db.GetLessonCollection(client)

	cursor, err := lessonCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()

	for cursor.Next(ctx) {
		var l models.Lesson
		cursor.Decode(&l)
		allLessons = append(allLessons, l)
	}

	successresponse.OK(&response)
	json.NewEncoder(response).Encode(allLessons)
}

//SingleLesson - Fetches a specific lesson
func SingleLesson(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var lesson models.Lesson

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	// client, err := db.ConnectToDB()

	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}
	lessonCollection.FindOne(context.TODO(), filter).Decode(&lesson)

	successresponse.OK(&response)
	json.NewEncoder(response).Encode(lesson)
}

//DeleteLesson - Deletes a specific lesson
func DeleteLesson(response http.ResponseWriter, request *http.Request) {
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

	lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}

	res, err := lessonCollection.DeleteOne(context.TODO(), filter)
	fmt.Printf("LESSON DELETED: %v", res)
	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	successresponse.NoContent(&response)
}

// UpdateLesson - Updates a specific lesson
func UpdateLesson(response http.ResponseWriter, request *http.Request) {
	var err error

	err = cookiehandler.VerifyCookie(response, request)

	if err != nil {
		return
	}

	var lesson models.Lesson

	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])

	if err != nil {
		customerror.BadRequest(&response, err)
		return
	}

	err = json.NewDecoder(request.Body).Decode(&lesson)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}

	// client, err := db.ConnectToDB()

	// if err != nil {
	// 	customerror.InternalServerError(&response, err)
	// 	return
	// }

	// lessonCollection := db.GetLessonCollection(client)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": lesson}

	res, err := lessonCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		customerror.InternalServerError(&response, err)
		return
	}
	fmt.Printf("LESSON UPDATED: %v", res)
	successresponse.NoContent(&response)

}
